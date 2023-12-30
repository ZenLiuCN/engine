package engine

import (
	"bytes"
	"fmt"
	"github.com/ZenLiuCN/fn"
	. "github.com/dop251/goja"
	"io"
	"log/slog"
)

//region Register  helper

func (s *Engine) BufferConsole(value *bytes.Buffer) {
	if value == nil {
		fn.Panic(s.Runtime.Set("console", NewBufferConsole()))
		return
	}
	fn.Panic(s.Runtime.Set("console", NewBufferConsoleOf(value)))
}
func (s *Engine) LoggerConsole(value *slog.Logger) {
	fn.Panic(s.Runtime.Set("console", NewConsole(value)))
}

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
func (s *Engine) Compile(src string, ts, entry bool) *Code {
	return CompileSource(src, ts, entry)
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

func (s *Engine) RegisterResources(r io.Closer) {
	if r != nil {
		if _, ok := s.Resources[r]; !ok {
			s.Resources[r] = resHolder
		}
	}
	return
}
func (s *Engine) RemoveResources(r io.Closer) {
	if r != nil {
		if _, ok := s.Resources[r]; ok {
			delete(s.Resources, r)
		}
	}
	return
}
func RegisterResource[T io.Closer](e *Engine, v T) T {
	e.RegisterResources(v)
	return v
}
func RemoveResource[T io.Closer](e *Engine, v T) T {
	e.RemoveResources(v)
	return v
}

var (
	resHolder = struct{}{}
)
