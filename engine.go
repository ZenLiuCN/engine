package engine

import (
	cx "context"
	"errors"
	"fmt"
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

// region Value helper
func (s *Engine) parse(r any) (err error) {
	switch e := r.(type) {
	case *Exception:
		err = &ScriptError{Err: e, Stack: nil}
	case error:
		stack := s.CaptureCallStack(20, nil)
		b := GetBytesBuffer()
		for _, s := range stack {
			s.Write(b)
		}
		err = &ScriptError{Err: fmt.Errorf("%w", e), Stack: b}
	default:
		stack := s.CaptureCallStack(20, nil)
		b := GetBytesBuffer()
		for _, s := range stack {
			s.Write(b)
		}
		err = &ScriptError{Err: fmt.Errorf("%s", e), Stack: b}
	}
	return
}

// CallFunction invoke a function (without this)
func (s *Engine) CallFunction(fn Value, values ...any) (v Value, err error) {
	s.TryStartEventLoop()
	defer func() {
		if r := recover(); r != nil {
			err = s.parse(r)
		}
	}()
	if f, ok := AssertFunction(fn); ok {
		v := make([]Value, len(values))
		for i, value := range values {
			v[i] = s.ToValue(value)
		}
		return f(Undefined(), v...)
	}
	return nil, fmt.Errorf("%s not a function", fn)
}
func (s *Engine) Callable(fn Value) Callable {
	if f, ok := AssertFunction(fn); !ok {
		panic(fmt.Errorf("%s not a function", fn))
	} else {
		return f
	}
}

// CallMethod invoke a method (with this)
func (s *Engine) CallMethod(fn Value, self Value, values ...any) (v Value, err error) {
	s.TryStartEventLoop()
	defer func() {
		if r := recover(); r != nil {
			err = s.parse(r)
		}
	}()
	if f, ok := AssertFunction(fn); ok {
		v := make([]Value, len(values))
		for i, value := range values {
			v[i] = s.ToValue(value)
		}
		return f(self, v...)
	}
	return nil, fmt.Errorf("%s not a function", fn)
}
func (s *Engine) CallConstruct(fn Value, values ...any) (*Object, error) {
	s.TryStartEventLoop()
	if f, ok := AssertConstructor(fn); ok {
		v := make([]Value, len(values))
		for i, value := range values {
			v[i] = s.ToValue(value)
		}
		return f(nil, v...)
	}
	return nil, fmt.Errorf("%s not a Constructor", fn)
}
func (s *Engine) ToValues(args ...any) []Value {
	n := make([]Value, len(args))
	for i, arg := range args {
		n[i] = s.ToValue(arg)
	}
	return n
}

//endregion

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
	return s.Runtime.RunString(CompileTs(src))
}

// RunJs execute javascript code. Should manual control the execution, for a automatic timeout control see  RunJsTimeout.
func (s *Engine) RunJs(src string) (v Value, err error) {
	s.TryStartEventLoop()
	defer func() {
		if r := recover(); r != nil {
			err = s.parse(r)
		}
	}()
	return s.Runtime.RunString(CompileJs(src))
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
	return s.awaiting(CompileSource(src, false).Program, warm, ctx)
}

// RunTsContext run Ts source
// with context. If context closed early, the value will be HaltJobs, the error will be ErrTimeout.
func (s *Engine) RunTsContext(src string, warm time.Duration, ctx cx.Context) (v Value, err error) {
	return s.awaiting(CompileSource(src, true).Program, warm, ctx)
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

//region RegisterMod helper

// Set create global value
func (s *Engine) Set(name string, value any) {
	fn.Panic(s.Runtime.Set(name, value))
}

// RegisterFunction create global function
func (s *Engine) RegisterFunction(name string, ctor func(c FunctionCall) Value) {
	fn.Panic(s.Runtime.Set(name, ctor))
}

// RegisterType create global simple type
func (s *Engine) RegisterType(name string, ctor func(v []Value) any) {
	s.Set(name, s.ToConstructor(ctor))
}

// ToInstance create instance of a value with simple prototype, use for constructor only.
func (s *Engine) ToInstance(v any, c ConstructorCall) *Object {
	o := s.ToValue(v).(*Object)
	fn.Panic(o.SetPrototype(c.This.Prototype()))
	return o
}
func (s *Engine) ToConstructor(ct func(v []Value) any) func(ConstructorCall) *Object {
	return func(c ConstructorCall) *Object {
		val := ct(c.Arguments)
		if val != nil {
			return s.ToInstance(val, c)
		}
		panic("can't construct type: " + c.This.ClassName())
	}
}

// ToSelfReferConstructor create a Constructor which require use itself. see [ Bytes ]
func (s *Engine) ToSelfReferConstructor(ct func(ctor Value, v []Value) any) Value {
	var ctor Value
	ctor = s.ToValue(func(c ConstructorCall) *Object {
		val := ct(ctor, c.Arguments)
		if val != nil {
			return s.ToInstance(val, c)
		}
		panic("can't construct type: " + c.This.ClassName())
	})
	return ctor
}

// ToSelfReferRawConstructor create a Constructor which require use itself. see [ Bytes ]
func (s *Engine) ToSelfReferRawConstructor(ct func(ctor Value, call ConstructorCall) *Object) Value {
	var ctor Value
	ctor = s.ToValue(func(c ConstructorCall) *Object {
		return ct(ctor, c)
	})
	return ctor
}

//endregion

//region Const and helper

// IsNullish check if value is null or undefined
func (s *Engine) IsNullish(v Value) bool {
	return IsNullish(v)
}
func (s *Engine) Compile(src string, ts bool) *Code {
	return CompileSource(src, ts)
}
func (s *Engine) Undefined() Value {
	return Undefined()
}
func (s *Engine) Null() Value {
	return Null()
}
func (s *Engine) NaN() Value {
	return NaN()
}
func (s *Engine) PosInf() Value {
	return PositiveInf()
}
func (s *Engine) NegInf() Value {
	return NegativeInf()
}

// NewPromise create new Promise, must use StopEventBusAwait
func (s *Engine) NewPromise() (promise *Promise, resolve func(any), reject func(any)) {
	p, resolveFunc, rejectFunc := s.Runtime.NewPromise()
	callback := s.EventLoop.registerCallback()
	resolve = func(result any) {
		callback(func() {
			resolveFunc(result)
		})
	}
	reject = func(reason any) {
		callback(func() {
			rejectFunc(reason)
		})
	}
	return p, resolve, reject
}

//endregion

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
