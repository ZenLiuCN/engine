package engine

import (
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
	engine.RegisterType(i.TypeName, func(v []goja.Value) any {
		return i.Ctor(engine, v)
	})
}

type JsTypeDefined[T any] struct {
	TypeName string
	Ctor     func(*Engine, []goja.Value) T
	Defined  []byte
}

func (i JsTypeDefined[T]) TypeDefine() []byte {
	return i.Defined
}
func (i JsTypeDefined[T]) Name() string {
	return i.TypeName
}

func (i JsTypeDefined[T]) Register(engine *Engine) {
	engine.RegisterType(i.TypeName, func(v []goja.Value) any {
		return i.Ctor(engine, v)
	})
}
