package engine

import (
	cx "context"
	"errors"
	"github.com/ZenLiuCN/fn"
	. "github.com/dop251/goja"
	"io"
	"path/filepath"
	"strings"
	"time"
)

type Engine struct {
	*Runtime
	*EventLoop
	Resources         map[io.Closer]struct{}
	require           *Require
	scriptRootHandler func(p string)
	Debug             bool
	SourceMap         SourceMapping
}

// Register register mods
func (s *Engine) Register(mods ...Mod) {
	for _, module := range mods {
		if im, ok := module.(InitializeMod); ok {
			m := im.Initialize(s)
			fn.Panic(s.Runtime.Set(m.Name(), m))
		} else if tm, ok := module.(TopMod); ok {
			tm.Register(s)
		} else {
			fn.Panic(s.Runtime.Set(module.Name(), module))
		}

	}
}

// DisableModules disable load modules for current engine, modules are the full import path
func (s *Engine) DisableModules(modules ...string) bool {
	r := s.Get("_$require")
	if !s.IsNullish(r) {
		if x, ok := r.Export().(*Require); ok {
			x.AddDisabled(modules...)
			return true
		}
	}
	return false
}

// Free recycle this engine
func (s *Engine) Free() {
	debug("free engine,stop event loop")
	s.EventLoop.StopEventLoopNoWait()
	if s.Resources != nil && len(s.Resources) > 0 {
		for resource := range s.Resources {
			if resource != nil {
				_ = resource.Close()
			}
		}
		s.Resources = nil
	}
}

// region direct

// SetScriptPath define current script path to use for imports
func (s *Engine) SetScriptPath(p string) {
	if s.scriptRootHandler != nil {
		s.scriptRootHandler(p)
	}
}

// RunString execute raw javascript (es5|es6 without import) code.
// this not support import and some polyfill features.
func (s *Engine) RunString(src string) (v Value, err error) {
	s.TryStartEventLoop()
	defer func() {
		if r := recover(); r != nil {
			err = s.parse(r)
		}
	}()
	v, err = s.Runtime.RunString(src)
	if err != nil {
		panic(err)
	}
	return
}

// RunTs execute typescript code. Should manual control the execution, for a automatic timeout control see  RunTsTimeout.
func (s *Engine) RunTs(src string) (v Value, err error) {
	s.TryStartEventLoop()
	defer func() {
		if r := recover(); r != nil {
			err = s.parse(r)
		}
		s.SourceMap = nil
	}()
	if s.Debug {
		var src string
		src, s.SourceMap = CompileTsWithMapping(src, true)
		v, err = s.Runtime.RunString(src)
	} else {
		v, err = s.Runtime.RunString(CompileTs(src, true))
	}
	if err != nil {
		panic(err)
	}
	return
}

// RunJs execute javascript code. Should manual control the execution, for a automatic timeout control see  RunJsTimeout.
func (s *Engine) RunJs(src string) (v Value, err error) {
	s.TryStartEventLoop()
	defer func() {
		if r := recover(); r != nil {
			err = s.parse(r)
		}
		s.SourceMap = nil
	}()
	if s.Debug {
		var src string
		src, s.SourceMap = CompileJsWithMapping(src, true)
		v, err = s.Runtime.RunString(src)
	} else {
		v, err = s.Runtime.RunString(CompileJs(src, true))
	}
	if err != nil {
		panic(err)
	}
	return
}

// RunCode execute compiled code. The execution time should control manually, for an automatic timeout control see  RunCodeTimeout.
func (s *Engine) RunCode(code *Code) (v Value, err error) {
	s.SetScriptPath(computeCodeDir(code.Path))
	s.TryStartEventLoop()
	defer func() {
		if r := recover(); r != nil {
			err = s.parse(r)
		}
	}()
	v, err = s.Runtime.RunProgram(code.Program)
	if err != nil {
		panic(err)
	}
	return
}
func (s *Engine) RunCodeWithMapping(code *Code, mapping SourceMapping) (v Value, err error) {
	s.SetScriptPath(computeCodeDir(code.Path))
	s.TryStartEventLoop()
	defer func() {
		if r := recover(); r != nil {
			err = s.parse(r)
		}
		s.SourceMap = nil
	}()
	s.SourceMap = mapping
	v, err = s.Runtime.RunProgram(code.Program)
	if err != nil {
		panic(err)
	}
	return
}

//endregion

//region context

// RunCodeContext run code
// with context. If context closed early, the value will be HaltJobs, the error will be ErrTimeout.
func (s *Engine) RunCodeContext(code *Code, warm time.Duration, ctx cx.Context) (v Value, err error) {
	s.SetScriptPath(computeCodeDir(code.Path))
	return s.awaiting(code.Program, warm, ctx)
}

// RunJsContext run js source
// with context. If context closed early, the value will be HaltJobs, the error will be ErrTimeout.
func (s *Engine) RunJsContext(src string, warm time.Duration, ctx cx.Context) (v Value, err error) {
	if s.Debug {
		code, mapping := CompileSourceWithMapping(src, false, true)
		s.SetScriptPath(computeCodeDir(code.Path))
		s.SourceMap = mapping
		return s.awaiting(code.Program, warm, ctx)
	}
	code := CompileSource(src, false, true)
	s.SetScriptPath(computeCodeDir(code.Path))
	return s.awaiting(code.Program, warm, ctx)
}

// RunTsContext run Ts source
// with context. If context closed early, the value will be HaltJobs, the error will be ErrTimeout.
func (s *Engine) RunTsContext(src string, warm time.Duration, ctx cx.Context) (v Value, err error) {
	if s.Debug {
		code, mapping := CompileSourceWithMapping(src, true, true)
		s.SetScriptPath(computeCodeDir(code.Path))
		s.SourceMap = mapping
		return s.awaiting(code.Program, warm, ctx)
	}
	code := CompileSource(src, true, true)
	s.SetScriptPath(computeCodeDir(code.Path))
	return s.awaiting(code.Program, warm, ctx)
}

//endregion

func execute(e *Engine, ch chan<- Maybe[Value], act *Program) {
	defer func() {
		if r := recover(); r != nil {
			er := e.parse(r)
			ch <- Maybe[Value]{Error: er}
		}
	}()
	rr, err := e.Runtime.RunProgram(act)
	if err != nil {
		panic(err)
	}
	ch <- Maybe[Value]{Value: rr}
}

func (s *Engine) awaiting(act *Program, warm time.Duration, ctx cx.Context) (v Value, err error) {
	s.TryStartEventLoop()
	defer func() {
		s.ClearInterrupt()
		if r := recover(); r != nil {
			err = s.parse(r)
		}
	}()
	ch := make(chan Maybe[Value])
	defer close(ch)
	go execute(s, ch, act)
	time.Sleep(warm)
	j := s.AwaitWithContext(ctx)
	if !j.IsZero() {
		s.Interrupt(ErrTimeout)
	}
	r := <-ch
	if !j.IsZero() {
		return s.ToValue(j), ErrTimeout
	}
	if r.Error != nil {
		panic(r.Error)
	}
	return r.Value, r.Error
}
func computeCodeDir(s string) string {
	if s == "" {
		return ""
	}
	b := filepath.Base(s)
	if strings.LastIndexByte(b, '.') > 0 {
		return filepath.Dir(s)
	}
	return b
}

var (
	ErrTimeout = errors.New("execution timeout")
)

func NewEngine(modules ...Mod) (r *Engine) {
	r = &Engine{Runtime: New(), Resources: make(map[io.Closer]struct{})}
	r.Runtime.SetFieldNameMapper(EngineFieldMapper{})
	r.EventLoop = NewEventLoop(r)
	r.Register(Mods()...)
	r.Register(modules...)
	r.Set("registerResource", func(closer io.Closer) io.Closer {
		r.RegisterResources(closer)
		return closer
	})
	r.Set("removeResource", func(closer io.Closer) io.Closer {
		r.RemoveResources(closer)
		return closer
	})
	return
}
func NewRawEngine(modules ...Mod) (r *Engine) {
	r = &Engine{Runtime: New(), Resources: make(map[io.Closer]struct{})}
	r.Runtime.SetFieldNameMapper(EngineFieldMapper{})
	r.EventLoop = NewEventLoop(r)
	r.Register(modules...)
	r.Set("autoClose", func(closer io.Closer) io.Closer {
		r.RegisterResources(closer)
		return closer
	})
	return
}
