package engine

import (
	"fmt"
	"github.com/ZenLiuCN/fn"
	. "github.com/dop251/goja"
)

type Engine struct {
	*Runtime
	*EventLoop
}

// Register register modules
func (s *Engine) Register(modules ...Module) {
	for _, module := range modules {
		if im, ok := module.(InitializeModule); ok {
			m := im.Initialize(s)
			fn.Panic(s.Runtime.Set(m.Name(), m))
		} else if tm, ok := module.(TopModule); ok {
			tm.Register(s)
		} else {
			fn.Panic(s.Runtime.Set(module.Name(), module))
		}

	}
}

// Free recycle this engine
func (s *Engine) Free() {
	s.EventLoop.StopEventLoopNoWait()
}

//region Value helper

// CallFunction invoke a function (without this)
func (s *Engine) CallFunction(fn Value, values ...any) (Value, error) {
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
func (s *Engine) CallMethod(fn Value, self Value, values ...any) (Value, error) {
	if f, ok := AssertFunction(fn); ok {
		v := make([]Value, len(values))
		for i, value := range values {
			v[i] = s.ToValue(value)
		}
		return f(self, v...)
	}
	return nil, fmt.Errorf("%s not a function", fn)
}
func (s *Engine) Construct(name string, args ...any) *Object {
	return fn.Panic1(s.New(s.Get(name), s.ToValues(args...)...))
}
func (s *Engine) ToValues(args ...any) []Value {
	n := make([]Value, len(args))
	for i, arg := range args {
		n[i] = s.ToValue(arg)
	}
	return n
}

//endregion

// RunTypeScript execute typescript code
func (s *Engine) RunTypeScript(src string) (Value, error) {
	s.StartEventLoopWhenNotStarted()
	return s.Runtime.RunString(CompileTs(src))
}

// RunJavaScript execute javascript code
func (s *Engine) RunJavaScript(src string) (Value, error) {
	s.StartEventLoopWhenNotStarted()
	return s.Runtime.RunString(CompileJs(src))
}

// RunScript execute javascript (es5|es6 without import) code
func (s *Engine) RunScript(src string) (Value, error) {
	s.StartEventLoopWhenNotStarted()
	return s.Runtime.RunString(src)
}

// Execute execute compiled code
func (s *Engine) Execute(code *Code) (Value, error) {
	s.StartEventLoopWhenNotStarted()
	return s.Runtime.RunProgram(code.Program)
}

//region Register helper

func (s *Engine) Set(name string, value any) {
	fn.Panic(s.Runtime.Set(name, value))
}

func (s *Engine) Function(name string, ctor func(c FunctionCall) Value) {
	fn.Panic(s.Runtime.Set(name, ctor))
}
func (s *Engine) RegisterType(name string, ctor func(v []Value) any) {
	s.Set(name, s.ToConstructor(ctor))
}
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

//endregion

//region Const and helper

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

func NewEngine(modules ...Module) (r *Engine) {
	r = &Engine{Runtime: New()}
	r.Runtime.SetFieldNameMapper(EngineFieldMapper{})
	r.EventLoop = NewEventLoop(r)
	r.Register(Modules()...)
	r.Register(modules...)
	return
}
func NewRawEngine(modules ...Module) (r *Engine) {
	r = &Engine{Runtime: New()}
	r.Runtime.SetFieldNameMapper(EngineFieldMapper{})
	r.EventLoop = NewEventLoop(r)
	r.Register(modules...)
	return
}
