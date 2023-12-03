package engine

import (
	"encoding/json"
	"fmt"
	"github.com/dop251/goja"
	"strings"
)

func Throw(rt *goja.Runtime, err error) {
	if e, ok := err.(*goja.Exception); ok {
		panic(e)
	}
	panic(rt.NewGoError(err))
}

func ToBytes(data any) ([]byte, error) {
	switch dt := data.(type) {
	case []byte:
		return dt, nil
	case string:
		return []byte(dt), nil
	case goja.ArrayBuffer:
		return dt.Bytes(), nil
	default:
		return nil, fmt.Errorf("invalid type %T, expected string, []byte or ArrayBuffer", data)
	}
}

func ToString(data any) (string, error) {
	switch dt := data.(type) {
	case []byte:
		return string(dt), nil
	case string:
		return dt, nil
	case goja.ArrayBuffer:
		return string(dt.Bytes()), nil
	default:
		return "", fmt.Errorf("invalid type %T, expected string, []byte or ArrayBuffer", data)
	}
}
func IsNullish(v goja.Value) bool {
	return v == nil || goja.IsUndefined(v) || goja.IsNull(v)
}

func IsAsyncFunction(rt *goja.Runtime, val goja.Value) bool {
	if IsNullish(val) {
		return false
	}
	return val.ToObject(rt).Get("constructor").ToObject(rt).Get("name").String() == "AsyncFunction"
}
func ValueString(v goja.Value) string {
	mv, ok := v.(json.Marshaler)
	if !ok {
		return v.String()
	}

	b, err := json.Marshal(mv)
	if err != nil {
		return v.String()
	}
	return string(b)
}
func ValuesString(args ...goja.Value) string {
	var msg strings.Builder
	for i := 0; i < len(args); i++ {
		if i > 0 {
			msg.WriteString(" ")
		}
		msg.WriteString(ValueString(args[i]))
	}
	return msg.String()
}

type Maybe[T any] struct {
	Value T
	Error error
}

func MaybeOk[T any](v T) Maybe[T] {
	return Maybe[T]{Value: v}
}
func MaybeError[T any](e error) Maybe[T] {
	return Maybe[T]{Error: e}
}
func MaybeBoth[T any](v T, e error) Maybe[T] {
	return Maybe[T]{Value: v, Error: e}
}
