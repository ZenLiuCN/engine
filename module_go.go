package engine

import (
	_ "embed"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"maps"
	"reflect"
	"strconv"
	"sync/atomic"
)

var (
	//go:embed module_go.d.ts
	goDefine []byte
	goMap    = map[string]any{
		"intToString": func(value goja.Value, v []string) map[string]any {
			m := value.Export().(map[string]any)
			for _, s := range v {
				m[s] = strconv.FormatInt(m[s].(int64), 10)
			}
			return m
		},
		"intFromString": func(value goja.Value, v []string) map[string]any {
			m := value.Export().(map[string]any)
			for _, s := range v {
				m[s] = fn.Panic1(strconv.ParseInt(m[s].(string), 0, 64))
			}
			return m
		},
		"intToStringArray": func(value goja.Value, v []string) []any {
			vx := value.Export().([]any)
			for _, a := range vx {
				m := a.(map[string]any)
				for _, s := range v {
					m[s] = strconv.FormatInt(m[s].(int64), 10)

				}
			}
			return vx
		},
		"intFromStringArray": func(value goja.Value, v []string) (r []any, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err = v
					default:
						err = fmt.Errorf("%s", v)
					}
				}
			}()
			vx := value.Export().([]any)
			for _, a := range vx {
				m := a.(map[string]any)
				for _, s := range v {
					m[s] = fn.Panic1(strconv.ParseInt(m[s].(string), 0, 64))
				}
			}
			return vx, nil
		},
		"complex32": func(r, i float32) complex64 {
			return complex(r, i)
		},
		"complex64": func(r, i float64) complex128 {
			return complex(r, i)
		},
		"real32": func(c complex64) float32 {
			return real(c)
		},
		"real364": func(c complex128) float64 {
			return real(c)
		},
		"imag32": func(c complex64) float32 {
			return real(c)
		},
		"imag64": func(c complex128) float64 {
			return real(c)
		},
		"bytesFromString":       func(v string) []byte { return []byte(v) },
		"runesFromString":       func(v string) []rune { return []rune(v) },
		"stringFromBytes":       func(v []byte) string { return string(v) },
		"stringFromRunes":       func(v []rune) string { return string(v) },
		"toInt8":                func(v float64) int8 { return int8(v) },
		"toInt16":               func(v float64) int16 { return int16(v) },
		"toInt32":               func(v float64) int32 { return int32(v) },
		"toInt64":               func(v float64) int64 { return int64(v) },
		"toUint8":               func(v float64) uint8 { return uint8(v) },
		"toUint16":              func(v float64) uint16 { return uint16(v) },
		"toUint32":              func(v float64) uint32 { return uint32(v) },
		"toUint64":              func(v float64) uint64 { return uint64(v) },
		"typeOf":                TypeOf,
		"usageOf":               TypeUsage,
		"elementOf":             ElementOf,
		"TypeKindInvalid":       reflect.Invalid,
		"TypeKindBool":          reflect.Bool,
		"TypeKindInt":           reflect.Int,
		"TypeKindInt8":          reflect.Int8,
		"TypeKindInt16":         reflect.Int16,
		"TypeKindInt32":         reflect.Int32,
		"TypeKindInt64":         reflect.Int64,
		"TypeKindUint":          reflect.Uint,
		"TypeKindUint8":         reflect.Uint8,
		"TypeKindUint16":        reflect.Uint16,
		"TypeKindUint32":        reflect.Uint32,
		"TypeKindUint64":        reflect.Uint64,
		"TypeKindUintptr":       reflect.Uintptr,
		"TypeKindFloat32":       reflect.Float32,
		"TypeKindFloat64":       reflect.Float64,
		"TypeKindComplex64":     reflect.Complex64,
		"TypeKindComplex128":    reflect.Complex128,
		"TypeKindArray":         reflect.Array,
		"TypeKindChan":          reflect.Chan,
		"TypeKindFunc":          reflect.Func,
		"TypeKindInterface":     reflect.Interface,
		"TypeKindMap":           reflect.Map,
		"TypeKindPointer":       reflect.Pointer,
		"TypeKindSlice":         reflect.Slice,
		"TypeKindString":        reflect.String,
		"TypeKindStruct":        reflect.Struct,
		"TypeKindUnsafePointer": reflect.UnsafePointer,
	}
)

type GoModule struct {
}

func (c GoModule) Identity() string {
	return "go"
}

func (c GoModule) TypeDefine() []byte {
	return goDefine
}

func (c GoModule) Exports() map[string]any {
	return goMap
}
func (c GoModule) ExportsWithEngine(e *Engine) map[string]any {
	m := maps.Clone(goMap)
	ctor := e.ToConstructorRecover(func(v []goja.Value) any {
		n := len(v)
		if n == 0 {
			return Text("")
		} else if n != 1 {
			panic("bad argument")
		}
		switch x := v[0].Export().(type) {
		case string:
			return Text(x)
		case []byte:
			return Text(x)
		case *goja.ArrayBuffer:
			return Text(x.Bytes())
		case []rune:
			return Text(x)
		default:
			panic("bad argument")
		}
	})
	m["Text"] = ctor
	return m
}

type Text string

func (t Text) Bytes() []byte {
	return []byte(t)
}
func (t Text) Runes() []rune {
	return []rune(t)
}
func (t Text) ToString() string {
	return string(t)
}
func (t Text) String() string {
	return string(t)
}

type (
	Chan[T any] interface {
		Recv(func(T)) *goja.Promise
		Send(v T)
		Closed() bool
		Close()
	}
	ReadChan[T any] interface {
		Recv(func(T)) *goja.Promise
		Closed() bool
		Close()
	}
	WriteChan[T any] interface {
		Send(v T)
		Closed() bool
		Close()
	}
)
type rwch[T any] struct {
	e      *Engine
	ch     chan T
	closer chan<- struct{}
	closed atomic.Bool
}

func NewChan[T any](ch chan T, e *Engine) Chan[T] {
	return &rwch[T]{e: e, ch: ch, closed: atomic.Bool{}}
}

func (c *rwch[T]) Recv(f func(T)) *goja.Promise {
	p, r, j := c.e.NewPromise()
	if c.closer != nil {
		j("already have receiver")
	}
	ch := make(chan struct{})
	c.closer = ch
	go func() {
		defer func() {
			r(nil)
		}()
		for {
			select {
			case <-ch:
				return
			case v, ok := <-c.ch:
				if ok {
					f(v)
				} else {
					c.closed.Store(true)
					return
				}
			}
		}
	}()
	return p
}

func (c *rwch[T]) Send(v T) {
	c.ch <- v
}
func (c *rwch[T]) Close() {
	if c.closed.Load() {
		return
	}
	if c.closer != nil {
		c.closer <- struct{}{}
	}
	c.closed.Store(true)
	close(c.ch)
	if c.closer != nil {
		close(c.closer)
		c.closer = nil
	}
}
func (c *rwch[T]) Closed() bool {
	return c.closed.Load()
}

type roch[T any] struct {
	e      *Engine
	ch     <-chan T
	closer chan<- struct{}
	closed atomic.Bool
}

func NewChanReadOnly[T any](ch <-chan T, e *Engine) ReadChan[T] {
	return &roch[T]{e: e, ch: ch, closed: atomic.Bool{}}
}
func (c *roch[T]) Recv(f func(T)) *goja.Promise {
	p, r, j := c.e.NewPromise()
	if c.closer != nil {
		j("already have receiver")
	}
	ch := make(chan struct{})
	c.closer = ch
	go func() {
		defer func() {
			r(nil)
		}()
		for {
			select {
			case <-ch:
				return
			case v, ok := <-c.ch:
				if ok {
					f(v)
				} else {
					c.closed.Store(true)
					return
				}
			}
		}
	}()
	return p
}
func (c *roch[T]) Close() {
	if c.closed.Load() {
		return
	}
	if c.closer != nil {
		c.closer <- struct{}{}
	}
	c.closed.Store(true)
	if c.closer != nil {
		close(c.closer)
		c.closer = nil
	}
}
func (c *roch[T]) Closed() bool {
	return c.closed.Load()
}

type wch[T any] struct {
	ch     chan<- T
	closed atomic.Bool
}

func NewChanWriteOnly[T any](ch chan<- T) WriteChan[T] {
	return &wch[T]{ch: ch, closed: atomic.Bool{}}
}
func (c *wch[T]) Send(v T) {
	c.ch <- v
}
func (c *wch[T]) Close() {
	if c.closed.Load() {
		return
	}
	c.closed.Store(true)
	close(c.ch)

}
func (c *wch[T]) Closed() bool {
	return c.closed.Load()
}

type Maybe[T any] struct {
	Value T
	Error error
}

func (m Maybe[T]) Result() (T, error) {
	return m.Value, m.Error
}
func MaybeOk[T any](v T) Maybe[T] {
	return Maybe[T]{Value: v}
}
func MaybeError[T any](e error) Maybe[T] {
	return Maybe[T]{Error: e}
}
func MaybeBoth[T any](v T, e error) Maybe[T] {
	if e == nil {
		return Maybe[T]{Value: v}
	}
	return Maybe[T]{Value: v, Error: e}
}
func MaybeMap[T, R any](v T, e error, f func(T) R) Maybe[R] {
	if e == nil {
		return Maybe[R]{Value: f(v)}
	}
	return Maybe[R]{Error: e}
}
