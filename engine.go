package engine

import (
	cx "context"
	"errors"
	"github.com/ZenLiuCN/fn"
	. "github.com/dop251/goja"
	"time"
)

type Engine struct {
	*Runtime
	*EventLoop
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
}

// region direct

// RunString execute raw javascript (es5|es6 without import) code.
// this not support import and some polyfill features.
func (s *Engine) RunString(src string) (v Value, err error) {
	s.TryStartEventLoop()
	defer func() {
		if r := recover(); r != nil {
			err = s.parse(r)
		}
	}()
	return s.Runtime.RunString(src)
}

// RunTs execute typescript code. Should manual control the execution, for a automatic timeout control see  RunTsTimeout.
func (s *Engine) RunTs(src string) (v Value, err error) {
	s.TryStartEventLoop()
	defer func() {
		if r := recover(); r != nil {
			err = s.parse(r)
		}
	}()
	return s.Runtime.RunString(CompileTs(src, true))
}

// RunJs execute javascript code. Should manual control the execution, for a automatic timeout control see  RunJsTimeout.
func (s *Engine) RunJs(src string) (v Value, err error) {
	s.TryStartEventLoop()
	defer func() {
		if r := recover(); r != nil {
			err = s.parse(r)
		}
	}()
	return s.Runtime.RunString(CompileJs(src, true))
}

// RunCode execute compiled code. The execution time should control manually, for an automatic timeout control see  RunCodeTimeout.
func (s *Engine) RunCode(code *Code) (v Value, err error) {
	s.TryStartEventLoop()
	defer func() {
		if r := recover(); r != nil {
			err = s.parse(r)
		}
	}()
	return s.Runtime.RunProgram(code.Program)
}

//endregion

//region context

// RunCodeContext run code
// with context. If context closed early, the value will be HaltJobs, the error will be ErrTimeout.
func (s *Engine) RunCodeContext(code *Code, warm time.Duration, ctx cx.Context) (v Value, err error) {
	return s.awaiting(code.Program, warm, ctx)
}

// RunJsContext run js source
// with context. If context closed early, the value will be HaltJobs, the error will be ErrTimeout.
func (s *Engine) RunJsContext(src string, warm time.Duration, ctx cx.Context) (v Value, err error) {
	return s.awaiting(CompileSource(src, false, true).Program, warm, ctx)
}

// RunTsContext run Ts source
// with context. If context closed early, the value will be HaltJobs, the error will be ErrTimeout.
func (s *Engine) RunTsContext(src string, warm time.Duration, ctx cx.Context) (v Value, err error) {
	return s.awaiting(CompileSource(src, true, true).Program, warm, ctx)
}

//endregion

func execute(e *Engine, ch chan<- Maybe[Value], act *Program) {
	defer func() {
		if r := recover(); r != nil {
			er := e.parse(r)
			ch <- Maybe[Value]{Error: er}
		}
	}()
	rr, er := e.Runtime.RunProgram(act)
	ch <- Maybe[Value]{Value: rr, Error: er}
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
	return r.Value, r.Error
}

var (
	ErrTimeout = errors.New("execution timeout")
)

func NewEngine(modules ...Mod) (r *Engine) {
	r = &Engine{Runtime: New()}
	r.Runtime.SetFieldNameMapper(EngineFieldMapper{})
	r.EventLoop = NewEventLoop(r)
	r.Register(Mods()...)
	r.Register(modules...)
	return
}
func NewRawEngine(modules ...Mod) (r *Engine) {
	r = &Engine{Runtime: New()}
	r.Runtime.SetFieldNameMapper(EngineFieldMapper{})
	r.EventLoop = NewEventLoop(r)
	r.Register(modules...)
	return
}
