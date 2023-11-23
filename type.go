package engine

import (
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
)

type JsType[T any] struct {
	TypeName string
	Ctor     func(*Engine, []goja.Value) T
}

func (i JsType[T]) Name() string {
	return i.TypeName
}

func (i JsType[T]) Register(engine *Engine) {
	fn.Panic(engine.Set(i.TypeName, func(call goja.ConstructorCall) *goja.Object {
		is := i.Ctor(engine, call.Arguments)
		v, ok := engine.ToValue(is).(*goja.Object)
		if ok {
			fn.Panic(v.SetPrototype(call.This.Prototype()))
			return v
		}
		return nil

	}))
}

type JsTypeDefined[T any] struct {
	TypeName string
	Ctor     func(*Engine, []goja.Value) T
	Defined  []byte
}

func (j JsTypeDefined[T]) TypeDefine() []byte {
	return j.Defined
}
func (i JsTypeDefined[T]) Name() string {
	return i.TypeName
}

func (i JsTypeDefined[T]) Register(engine *Engine) {
	fn.Panic(engine.Set(i.TypeName, func(call goja.ConstructorCall) *goja.Object {
		is := i.Ctor(engine, call.Arguments)
		v, ok := engine.ToValue(is).(*goja.Object)
		if ok {
			fn.Panic(v.SetPrototype(call.This.Prototype()))
			return v
		}
		return nil

	}))
}
