package engine

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"io"
	"os"
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
	engine.RegisterType("Buffer", func(v []goja.Value) any {
		var buf *Buffer
		if len(v) == 0 {
			buf = GetBuffer()
		} else if len(v) == 1 {
			switch t := v[0].Export().(type) {
			case string:
				buf = &Buffer{
					e: engine,
					b: bytes.NewBufferString(t),
				}
			case []byte:
				buf = &Buffer{
					e: engine,
					b: bytes.NewBuffer(t),
				}
			case *bytes.Buffer:
				buf = &Buffer{
					e: engine,
					b: t,
				}
			case *Bytes:
				buf = &Buffer{
					e: engine,
					b: bytes.NewBuffer(t.b),
				}
			case io.Reader:
				buf = GetBuffer()
				fn.Panic1(io.Copy(buf.b, t))
				buf.e = engine
			default:
				panic("bad parameter type")
			}
		} else {
			panic("bad parameters ")
		}
		return buf
	})
	engine.Constructor("Bytes", func(c goja.ConstructorCall) *goja.Object {
		v := c.Arguments
		var bin []byte
		if len(v) == 0 {

		} else if len(v) == 1 {
			switch t := v[0].Export().(type) {
			case string:
				bin = []byte(t)
			case []byte:
				bin = bytes.Clone(t)
			case goja.ArrayBuffer:
				bin = bytes.Clone(t.Bytes())
			case *Buffer:
				bin = bytes.Clone(t.Bytes())
			case *Bytes:
				bin = bytes.Clone(t.b)
			case io.Reader:
				bin = fn.Panic1(io.ReadAll(t))
			default:
				panic("bad parameter type")
			}
		} else {
			for _, value := range v {
				bin = append(bin, byte(value.ToNumber().Export().(int64)))
			}
		}
		o := &Bytes{
			Engine: engine,
			b:      bin,
		}
		da := engine.NewDynamicArray(o)
		do := engine.ToValue(o).(*goja.Object)
		_ = do.SetPrototype(c.This.Prototype())
		_ = da.Prototype().SetPrototype(do)

		return da
	})
}

func (b BufferModule) Name() string {
	return "buffer"
}

type Buffer struct {
	e        *Engine
	b        *bytes.Buffer
	Detached bool
}

func (b *Buffer) Free() {
	if b.b != nil {
		b.b.Reset()
		buffers.Put(b.b)
		b.Detached = true
	}
}
func (b *Buffer) Available() int {
	return b.b.Available()
}
func (b *Buffer) Length() int {
	return b.b.Len()
}
func (b *Buffer) Cap() int {
	return b.b.Cap()
}
func (b *Buffer) Truncate(n int) {
	b.b.Truncate(n)
}
func (b *Buffer) Grow(n int) {
	b.b.Grow(n)
}
func (b *Buffer) Reset() {
	b.b.Reset()
}
func (b *Buffer) Slice(from, to int) *goja.Object {
	if from < 0 || from >= to || to > b.b.Len() {
		panic("index overflow")
	}
	return fn.Panic1(b.e.New(b.e.Get("Buffer"), b.e.ToValue(b.b.Bytes()[from:to])))
}
func (b *Buffer) Runes() (r []string) {
	for {
		if u, _, err := b.b.ReadRune(); err == nil {
			r = append(r, string(u))
		} else if errors.Is(err, io.EOF) {
			break
		} else {
			panic(err)
		}
	}
	return
}
func (b *Buffer) Bytes() []byte {
	return b.b.Bytes()
}
func (b *Buffer) Binary() (r *goja.Object) {
	return fn.Panic1(b.e.Runtime.New(b.e.Get("Bytes"), b.e.ToValue(b.b.Bytes())))
}
func (b *Buffer) EachByte(v goja.Value) {
	if f, ok := goja.AssertFunction(v); ok {
		for {
			if u, err := b.b.ReadByte(); err == nil {
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
func (b *Buffer) MapByte(v goja.Value) *goja.Object {
	if f, ok := goja.AssertFunction(v); ok {
		var ar []any
		for {
			if u, err := b.b.ReadByte(); err == nil {
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
func (b *Buffer) EachRune(v goja.Value) {
	if f, ok := goja.AssertFunction(v); ok {
		for {
			if u, _, err := b.b.ReadRune(); err == nil {
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
func (b *Buffer) MapRune(v goja.Value) *goja.Object {
	if f, ok := goja.AssertFunction(v); ok {
		var ar []any
		for {
			if u, _, err := b.b.ReadRune(); err == nil {
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
func (b *Buffer) ToString() string {
	return b.b.String()
}
func (b *Buffer) ReadString(delimiter string) string {
	s, err := b.b.ReadString([]byte(delimiter)[0])
	if err != nil && errors.Is(err, io.EOF) {
		return ""
	} else if err != nil {
		panic(err)
	}
	return s
}
func (b *Buffer) WriteString(c string) int {
	return fn.Panic1(b.b.WriteString(c))
}
func (b *Buffer) ReadByte() byte {
	return fn.Panic1(b.b.ReadByte())
}
func (b *Buffer) WriteByte(c uint8) {
	fn.Panic(b.b.WriteByte(c))
}
func (b *Buffer) ReadRune() string {
	r, _ := fn.Panic2(b.b.ReadRune())
	return string(r)
}
func (b *Buffer) WriteRune(r string) {
	b.b.WriteRune([]rune(r)[0])
}

// ArrayBuffer only for using with Engine
func (b *Buffer) ArrayBuffer() goja.ArrayBuffer {
	return b.e.NewArrayBuffer(b.b.Bytes())
}
func (b *Buffer) WriteBuffer(buf goja.ArrayBuffer) int {
	return fn.Panic1(b.b.Write(buf.Bytes()))
}

func (b *Buffer) SaveTo(path string) {
	f := fn.Panic1(os.OpenFile(path, os.O_CREATE|os.O_TRUNC, os.ModePerm))
	defer fn.IgnoreClose(f)
	fn.Panic1(io.Copy(f, b.b))
}

func (b *Buffer) LoadFile(p string) {
	b.b.Reset()
	b.b.Write(fn.Panic1(os.ReadFile(p)))
}
func (b *Buffer) MergeFile(p string) {
	b.b.Write(fn.Panic1(os.ReadFile(p)))
}

func (b *Buffer) ToWriter() io.Writer {
	return b.b
}
func (b *Buffer) ToReader() io.Reader {
	return b.b
}

func GetBytesBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}
func PutBytesBuffer(buf *bytes.Buffer) {
	buffers.Put(buf)
}
func GetBuffer() *Buffer {
	return &Buffer{
		b: buffers.Get().(*bytes.Buffer),
	}
}

type Bytes struct {
	*Engine
	b      []byte
	Length int
}

func (b *Bytes) Len() int {
	return len(b.b)
}

func (b *Bytes) Get(idx int) goja.Value {
	return b.Engine.ToValue(b.b[idx])
}

func (b *Bytes) Set(idx int, val goja.Value) bool {
	if idx >= 0 && idx < len(b.b) {
		b.b[idx] = byte(val.ToInteger())
	}
	return false
}

func (b *Bytes) SetLen(i int) bool {
	if i != len(b.b) {
		return false
	}
	b.Length = i
	return true
}

func (b *Bytes) Slice(from, to int) *goja.Object {
	return b.Engine.Construct("Bytes", b.b[from:to])
}

func (b *Bytes) Clone() *goja.Object {
	return b.Engine.Construct("Bytes", bytes.Clone(b.b))
}
func (b *Bytes) Bytes() []byte {
	return b.b
}
func (b *Bytes) Append(v goja.Value) *goja.Object {
	var bin []byte
	switch t := v.Export().(type) {
	case []byte:
		bin = t
	case string:
		bin = []byte(t)
	case *Bytes:
		bin = b.b
	case *Buffer:
		bin = t.b.Bytes()
	case goja.ArrayBuffer:
		bin = t.Bytes()
	default:
		panic(fmt.Errorf("bad argument type of: %#v", v))
	}
	return b.Engine.Construct("Bytes", append(b.b, bin...))
}
func (b *Bytes) ToText() string {
	return string(b.b)
}
func (b *Bytes) ToReader() io.Reader {
	return bytes.NewReader(b.b)
}

var (
	buffers = sync.Pool{New: func() any {
		buf := &bytes.Buffer{}
		return buf
	}}
)
