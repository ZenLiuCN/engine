package legacy

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	. "github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"io"
	"os"
)

var (
	//go:embed module_buffer.d.ts
	bufDefine []byte
)

type BufferModule struct {
}

func (b BufferModule) Exports() map[string]any {
	return nil
}

func (b BufferModule) ExportsWithEngine(engine *Engine) map[string]any {

	return map[string]any{
		"Buffer": engine.ToConstructor(func(v []goja.Value) (any, error) {
			var buf *Buffer
			if len(v) == 0 {
				buf = GetBuffer()
			} else if len(v) == 1 {
				switch t := v[0].Export().(type) {
				case string:
					buf = &Buffer{
						e:      engine,
						Buffer: bytes.NewBufferString(t),
					}
				case []byte:
					buf = &Buffer{
						e:      engine,
						Buffer: bytes.NewBuffer(t),
					}
				case *bytes.Buffer:
					buf = &Buffer{
						e:      engine,
						Buffer: t,
					}
				case *Bytes:
					buf = &Buffer{
						e:      engine,
						Buffer: bytes.NewBuffer(t.b),
					}
				case io.Reader:
					buf = GetBuffer()
					fn.Panic1(io.Copy(buf.Buffer, t))
					buf.e = engine
				default:
					return nil, fmt.Errorf("bad parameter type")
				}
			} else {
				return nil, fmt.Errorf("bad parameters ")
			}
			return buf, nil
		}),
		"Bytes": engine.ToSelfReferRawConstructor(func(ctor goja.Value, c goja.ConstructorCall) *goja.Object {
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
				ctor:   ctor,
				Engine: engine,
				b:      bin,
			}
			da := engine.NewDynamicArray(o)
			do := engine.ToValue(o).(*goja.Object)
			_ = do.SetPrototype(c.This.Prototype())
			_ = da.Prototype().SetPrototype(do)
			return da
		}),
	}
}

func (b BufferModule) TypeDefine() []byte {
	return bufDefine
}

func (b BufferModule) Identity() string {
	return "go/buffer"
}

type Buffer struct {
	e        *Engine
	Buffer   *bytes.Buffer `js:"-"`
	Detached bool
}

func (b *Buffer) Free() {
	if b.Buffer != nil {
		b.Buffer.Reset()
		PutBytesBuffer(b.Buffer)
		b.Detached = true
	}
}
func (b *Buffer) Available() int {
	return b.Buffer.Available()
}
func (b *Buffer) Length() int {
	return b.Buffer.Len()
}
func (b *Buffer) Cap() int {
	return b.Buffer.Cap()
}
func (b *Buffer) Truncate(n int) {
	b.Buffer.Truncate(n)
}
func (b *Buffer) Grow(n int) {
	b.Buffer.Grow(n)
}
func (b *Buffer) Reset() {
	b.Buffer.Reset()
}
func (b *Buffer) Slice(from, to int) (*goja.Object, error) {
	if from < 0 || from >= to || to > b.Buffer.Len() {
		panic("index overflow")
	}
	return b.e.New(b.e.Get("Buffer"), b.e.ToValue(b.Buffer.Bytes()[from:to]))
}
func (b *Buffer) Runes() (r []string, err error) {
	var u rune
	for {
		if u, _, err = b.Buffer.ReadRune(); err == nil {
			r = append(r, string(u))
		} else if errors.Is(err, io.EOF) {
			break
		} else {
			return
		}
	}
	return
}
func (b *Buffer) Bytes() []byte {
	return b.Buffer.Bytes()
}
func (b *Buffer) EachByte(v goja.Value, err error) {
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
	if f, ok := goja.AssertFunction(v); ok {
		for {
			if u, err := b.Buffer.ReadByte(); err == nil {
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
func (b *Buffer) MapByte(v goja.Value) (r *goja.Object, err error) {
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
	if f, ok := goja.AssertFunction(v); ok {
		var ar []any
		for {
			if u, err := b.Buffer.ReadByte(); err == nil {
				ar = append(ar, fn.Panic1(f(goja.Undefined(), b.e.ToValue(u))))
			} else if errors.Is(err, io.EOF) {
				break
			} else {
				panic(err)
			}
		}
		return b.e.NewArray(ar...), nil
	}
	return nil, nil
}
func (b *Buffer) EachRune(v goja.Value, err error) {
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
	if f, ok := goja.AssertFunction(v); ok {
		for {
			if u, _, err := b.Buffer.ReadRune(); err == nil {
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
func (b *Buffer) MapRune(v goja.Value) (r *goja.Object, err error) {
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
	if f, ok := goja.AssertFunction(v); ok {
		var ar []any
		for {
			if u, _, err := b.Buffer.ReadRune(); err == nil {
				ar = append(ar, fn.Panic1(f(goja.Undefined(), b.e.ToValue(string(u)))))
			} else if errors.Is(err, io.EOF) {
				break
			} else {
				panic(err)
			}
		}
		return b.e.NewArray(ar...), nil
	}
	return nil, nil
}
func (b *Buffer) ToString() string {
	return b.Buffer.String()
}
func (b *Buffer) ReadString(delimiter string) (s string, err error) {
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

	s, err = b.Buffer.ReadString([]byte(delimiter)[0])
	if err != nil && errors.Is(err, io.EOF) {
		return "", nil
	} else if err != nil {
		panic(err)
	}
	return s, nil
}
func (b *Buffer) WriteString(c string) (int, error) {
	return b.Buffer.WriteString(c)
}
func (b *Buffer) ReadByte() (byte, error) {
	return b.Buffer.ReadByte()
}
func (b *Buffer) WriteByte(c uint8) error {
	return b.Buffer.WriteByte(c)
}
func (b *Buffer) ReadRune() (string, error) {
	r, _, err := b.Buffer.ReadRune()
	if err != nil {
		return "", err
	}
	return string(r), nil
}
func (b *Buffer) WriteRune(r string) {
	b.Buffer.WriteRune([]rune(r)[0])
}

// ArrayBuffer only for using with Engine
func (b *Buffer) ArrayBuffer() goja.ArrayBuffer {
	return b.e.NewArrayBuffer(b.Buffer.Bytes())
}
func (b *Buffer) WriteBuffer(buf goja.ArrayBuffer) int {
	return fn.Panic1(b.Buffer.Write(buf.Bytes()))
}

func (b *Buffer) SaveTo(path string) (err error) {
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
	f := fn.Panic1(os.OpenFile(path, os.O_CREATE|os.O_TRUNC, os.ModePerm))
	defer fn.IgnoreClose(f)
	fn.Panic1(io.Copy(f, b.Buffer))
	return
}

func (b *Buffer) LoadFile(p string) (err error) {
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
	b.Buffer.Reset()
	b.Buffer.Write(fn.Panic1(os.ReadFile(p)))
	return
}
func (b *Buffer) MergeFile(p string) (err error) {
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
	b.Buffer.Write(fn.Panic1(os.ReadFile(p)))
	return
}

func (b *Buffer) ToWriter() io.Writer {
	return b.Buffer
}
func (b *Buffer) ToReader() io.Reader {
	return b.Buffer
}

func GetBuffer() *Buffer {
	return &Buffer{
		Buffer: GetBytesBuffer(),
	}
}

type Bytes struct {
	*Engine
	ctor   goja.Value
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

func (b *Bytes) Slice(from, to int) (r *goja.Object, err error) {
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
	return fn.Panic1(b.Engine.CallConstruct(b.ctor, b.b[from:to])), nil
}
func (b *Bytes) Equals(v goja.Value) bool {
	switch t := v.Export().(type) {
	case string:
		return bytes.Equal(b.b, []byte(t))
	case []byte:
		return bytes.Equal(b.b, t)
	case *Bytes:
		return bytes.Equal(b.b, t.b)
	default:
		return false
	}
}
func (b *Bytes) Clone() (*goja.Object, error) {
	return b.Engine.CallConstruct(b.ctor, bytes.Clone(b.b))
}
func (b *Bytes) Bytes() []byte {
	return b.b
}
func (b *Bytes) Append(v goja.Value) (r *goja.Object, err error) {
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
	var bin []byte
	switch t := v.Export().(type) {
	case []byte:
		bin = t
	case string:
		bin = []byte(t)
	case *Bytes:
		bin = b.b
	case *Buffer:
		bin = t.Buffer.Bytes()
	case goja.ArrayBuffer:
		bin = t.Bytes()
	default:
		panic(fmt.Errorf("bad argument type of: %#v", v))
	}
	return fn.Panic1(b.Engine.CallConstruct(b.ctor, append(b.b, bin...))), nil
}
func (b *Bytes) ToText() string {
	return string(b.b)
}
func (b *Bytes) ToReader() io.Reader {
	return bytes.NewReader(b.b)
}
