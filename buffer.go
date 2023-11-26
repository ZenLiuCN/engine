package engine

import (
	"bytes"
	_ "embed"
	"errors"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"io"
	"os"
	"slices"
	"sync"
)

var (
	//go:embed buffer.d.ts
	bufDefine []byte
)

type BufferModule struct {
}

func (b BufferModule) TypeDefine() []byte {
	return bufDefine
}

func (b BufferModule) Register(engine *Engine) {
	engine.Constructor("Buffer", func(c goja.ConstructorCall) *goja.Object {
		var v *Buffer
		if len(c.Arguments) == 0 {
			v = GetBuffer()
		} else if len(c.Arguments) == 1 {
			vars := c.Arguments[0]
			switch t := vars.Export().(type) {
			case string:
				v = &Buffer{
					Buffer: bytes.NewBufferString(t),
				}
			case []byte:
				v = &Buffer{
					Buffer: bytes.NewBuffer(t),
				}
			default:
				panic("bad parameter type")
			}
		} else {
			panic("bad parameters ")
		}
		v.e = engine
		return engine.ToInstance(v, c)
	})
	engine.Function("binFrom", func(c goja.FunctionCall) goja.Value {
		vars := c.Arguments
		var d []byte
		if len(vars) == 1 {
			switch x := vars[0].Export().(type) {
			case []byte:
				d = x
			case string:
				d = []byte(x)
			case []rune:
				d = []byte(string(x))
			case io.Reader:
				d = fn.Panic1(io.ReadAll(x))
			case rune:
				d = []byte(string(x))
			default:
				fn.Panic(engine.ExportTo(vars[0], &d))
			}
		} else {
			for _, value := range vars {
				d = append(d, byte(value.ToNumber().Export().(int64)))
			}
		}
		return engine.ToValue(d)
	})
	engine.Function("binLen", func(c goja.FunctionCall) goja.Value {
		return engine.ToValue(len(c.Arguments[0].Export().([]byte)))
	})
	engine.Function("binToString", func(c goja.FunctionCall) goja.Value {
		return engine.ToValue(string(c.Arguments[0].Export().([]byte)))
	})
	engine.Function("binGet", func(c goja.FunctionCall) goja.Value {
		return engine.ToValue(c.Arguments[0].Export().([]byte)[c.Arguments[1].Export().(int64)])
	})
	engine.Function("binSet", func(c goja.FunctionCall) goja.Value {
		c.Arguments[0].Export().([]byte)[c.Arguments[1].Export().(int64)] = byte(c.Arguments[0].Export().(int64))
		return engine.Undefined()
	})
	engine.Function("binSlice", func(c goja.FunctionCall) goja.Value {
		return engine.ToValue(c.Arguments[0].Export().([]byte)[c.Arguments[1].Export().(int64):c.Arguments[2].Export().(int64)])
	})
	engine.Function("binMap", func(c goja.FunctionCall) goja.Value {
		v := c.Argument(0).Export().([]byte)
		f := engine.Callable(c.Argument(1))
		var r []any
		for i, b2 := range v {
			r = append(r, fn.Panic1(f(goja.Undefined(), engine.ToValue(b2), engine.ToValue(i))))
		}
		return engine.ToValue(r)
	})
	engine.Function("binEquals", func(c goja.FunctionCall) goja.Value {
		return engine.ToValue(slices.Equal(c.Argument(0).Export().([]byte), c.Argument(1).Export().([]byte)))
	})
	engine.Function("binAppend", func(c goja.FunctionCall) goja.Value {
		return engine.ToValue(append(c.Argument(0).Export().([]byte), byte(c.Argument(1).Export().(int64))))
	})
	engine.Function("binAppends", func(c goja.FunctionCall) goja.Value {
		return engine.ToValue(append(c.Argument(0).Export().([]byte), c.Argument(1).Export().([]byte)...))
	})
	engine.Function("binToArrayBuffer", func(c goja.FunctionCall) goja.Value {
		v := c.Argument(0).Export().([]byte)
		a := engine.NewArrayBuffer(v)
		return engine.ToValue(a)
	})
	engine.Function("binToReader", func(c goja.FunctionCall) goja.Value {
		v := c.Argument(0).Export().([]byte)
		a := bytes.NewReader(v)
		return engine.ToValue(a)
	})
}

func (b BufferModule) Name() string {
	return "buffer"
}

type Buffer struct {
	e *Engine
	*bytes.Buffer
	Detached bool
}

func (b *Buffer) Chars() (r []string) {
	for {
		if u, _, err := b.ReadRune(); err == nil {
			r = append(r, string(u))
		} else if errors.Is(err, io.EOF) {
			break
		} else {
			panic(err)
		}
	}
	return
}
func (b *Buffer) Binary() (r []byte) {

	return b.Bytes()
}
func (b *Buffer) MapU8(v goja.Value) *goja.Object {
	if f, ok := goja.AssertFunction(v); ok {
		var ar []any
		for {
			if u, err := b.ReadByte(); err == nil {
				ar = append(ar, fn.Panic1(f(goja.Undefined(), b.e.ToValue(u))))
			} else if errors.Is(err, io.EOF) {
				break
			} else {
				panic(err)
			}
		}
		return b.e.NewArray(ar...)
	}
	return nil
}
func (b *Buffer) MapChar(v goja.Value) *goja.Object {
	if f, ok := goja.AssertFunction(v); ok {
		var ar []any
		for {
			if u, _, err := b.ReadRune(); err == nil {
				ar = append(ar, fn.Panic1(f(goja.Undefined(), b.e.ToValue(string(u)))))
			} else if errors.Is(err, io.EOF) {
				break
			} else {
				panic(err)
			}
		}
		return b.e.NewArray(ar...)
	}
	return nil
}
func (b *Buffer) EachU8(v goja.Value) {
	if f, ok := goja.AssertFunction(v); ok {
		for {
			if u, err := b.ReadByte(); err == nil {
				if !fn.Panic1(f(goja.Undefined(), b.e.ToValue(u))).ToBoolean() {
					break
				}
			} else if errors.Is(err, io.EOF) {
				break
			} else {
				panic(err)
			}
		}
	}
}
func (b *Buffer) EachChar(v goja.Value) {
	if f, ok := goja.AssertFunction(v); ok {
		for {
			if u, _, err := b.ReadRune(); err == nil {
				if !fn.Panic1(f(goja.Undefined(), b.e.ToValue(string(u)))).ToBoolean() {
					break
				}
			} else if errors.Is(err, io.EOF) {
				break
			} else {
				panic(err)
			}
		}
	}
}
func (b *Buffer) Slice(from, to int) *Buffer {
	if from < 0 || from >= to || to > b.Len() {
		panic("index overflow")
	}
	buf := GetBuffer()
	buf.e = b.e
	buf.Buffer.Write(b.Bytes()[from:to])
	return buf
}
func (b *Buffer) ReadChar() string {
	r, _ := fn.Panic2(b.ReadRune())
	return string(r)
}
func (b *Buffer) SaveTo(path string) {
	f := fn.Panic1(os.OpenFile(path, os.O_CREATE|os.O_TRUNC, os.ModePerm))
	defer fn.IgnoreClose(f)
	fn.Panic1(io.Copy(f, b.Buffer))
}
func (b *Buffer) ReadU8() uint8 {
	return fn.Panic1(b.Buffer.ReadByte())
}

func (b *Buffer) WriteU8(c uint8) {
	fn.Panic(b.Buffer.WriteByte(c))
}
func (b *Buffer) WriteText(c string) int {
	return fn.Panic1(b.Buffer.WriteString(c))
}
func (b *Buffer) LoadFile(p string) {
	b.Buffer.Reset()
	b.Buffer.Write(fn.Panic1(os.ReadFile(p)))
}
func (b *Buffer) MergeFile(p string) {
	b.Buffer.Write(fn.Panic1(os.ReadFile(p)))
}

// ArrayBuffer only for using with Engine
func (b *Buffer) ArrayBuffer() goja.ArrayBuffer {
	return b.e.NewArrayBuffer(b.Bytes())
}
func (b *Buffer) WriteBuffer(buf goja.ArrayBuffer) int {
	return fn.Panic1(b.Write(buf.Bytes()))
}
func (b *Buffer) ToString() string {
	return b.Buffer.String()
}
func (b *Buffer) Free() {
	if b.Buffer != nil {
		b.Buffer.Reset()
		buffers.Put(b.Buffer)
		b.Detached = true
	}
}
func (b *Buffer) ToWriter() io.Writer {
	return b.Buffer
}
func GetBytesBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}
func PutBytesBuffer(buf *bytes.Buffer) {
	buffers.Put(buf)
}
func GetBuffer() *Buffer {
	return &Buffer{
		Buffer: buffers.Get().(*bytes.Buffer),
	}
}

var (
	buffers = sync.Pool{New: func() any {
		buf := &bytes.Buffer{}
		return buf
	}}
)
