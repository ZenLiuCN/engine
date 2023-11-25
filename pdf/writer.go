package pdf

import (
	"encoding/hex"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"io"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type (
	Writer interface {
		RawBinary(v []byte) Writer
		RawRune(x rune) Writer
		Raw(x string) Writer

		Use(func(Writer)) Writer
		Name(x string) Writer
		Text(x string) Writer
		Date(x time.Time) Writer
		Boolean(x bool) Writer
		Int(v int) Writer
		Real(x float64) Writer
		Ref(id, gen int) Writer
		Binary(x []byte) Writer
		Space() Writer
		LineFeed() Writer
		Dict(func(w Writer)) Writer
		Entry(name string, val Value) Writer
		Array(func(w Writer)) Writer
		Object(id, gen int, f func(w Writer)) Writer
		Stream(stream io.Reader) Writer
		Null() Writer
		True() Writer
		False() Writer
		Size() int
	}
	writer struct {
		io.Writer
		n int
	}
)

func NewWriter(w io.Writer) Writer {
	return &writer{Writer: w}
}

func (w *writer) RawRune(v rune) Writer {
	b := make([]byte, 4)
	w.n += fn.Panic1(w.Write(b[:utf8.EncodeRune(b, v)]))
	return w
}
func (w *writer) Use(f func(Writer)) Writer {
	f(w)
	return w
}
func (w *writer) Entry(name string, val Value) Writer {
	return w.Name(name).Space().Use(val.Write)
}
func (w *writer) RawBinary(v []byte) Writer {
	w.n += fn.Panic1(w.Write(v))
	return w
}
func (w *writer) Raw(v string) Writer {
	w.n += fn.Panic1(w.Write([]byte(v)))
	return w
}
func (w *writer) Name(x string) Writer {
	return w.Raw("/" + strings.ReplaceAll(x, " ", "#20"))
}

func (w *writer) Text(x string) Writer {
	w.RawBinary(textBegin)
	for _, value := range []rune(x) {
		switch value {
		case '(', '\\', ')':
			w.RawBinary(escape)
		case '\n':
			w.RawBinary(escLineFeed)
		case '\r':
			w.RawBinary(escReturns)
		case '\t':
			w.RawBinary(escTab)
		case '\b':
			w.RawBinary(escBackspace)
		case '\f':
			w.RawBinary(escFormFeed)
		default:
			w.RawRune(value)
		}
	}
	return w.RawBinary(textEnd)
}
func (w *writer) Date(x time.Time) Writer {
	return w.Raw("(D:").Raw(x.Format("20060102150405-07")).Raw("'00')")
}
func (w *writer) Boolean(x bool) Writer {
	if x {
		return w.RawBinary(trueValue)
	} else {
		return w.RawBinary(falseValue)
	}
}
func (w *writer) Real(x float64) Writer {
	return w.Raw(strconv.FormatFloat(x, 'f', -1, 10))
}
func (w *writer) Ref(id, gen int) Writer {
	return w.Raw(fmt.Sprintf(`%d %d R`, id, gen))
}
func (w *writer) Binary(x []byte) Writer {
	return w.RawRune('<').Raw(hex.EncodeToString(x)).RawRune('>')
}

func (w *writer) LineFeed() Writer {
	return w.RawBinary(lineFeed)
}
func (w *writer) Space() Writer {
	return w.RawBinary(space)
}

func (w *writer) Null() Writer {
	return w.RawBinary(null)
}
func (w *writer) True() Writer {
	return w.RawBinary(trueValue)
}
func (w *writer) Int(v int) Writer {
	return w.Raw(strconv.Itoa(v))
}
func (w *writer) False() Writer {
	return w.RawBinary(falseValue)
}
func (w *writer) Stream(stream io.Reader) Writer {
	w.RawBinary(streamBegin)
	w.n += int(fn.Panic1(io.Copy(w.Writer, stream)))
	w.RawBinary(streamEnd)
	return w
}

func (w *writer) Dict(f func(w Writer)) Writer {
	w.RawBinary(dictBegin)
	f(w)
	return w.RawBinary(dictEnd)
}
func (w *writer) Array(f func(w Writer)) Writer {
	w.RawBinary(arrayBegin)
	f(w)
	return w.RawBinary(arrayEnd)
}

func (w *writer) Object(id, gen int, f func(w Writer)) Writer {
	w.Raw(fmt.Sprintf("%d %d ", id, gen))
	w.RawBinary(objectBegin)
	f(w)
	return w.RawBinary(objectEnd)
}

func (w *writer) Size() int {
	return w.n
}
