package engine

import "github.com/dop251/goja"

type TextEncoders struct {
}

func (t TextEncoders) Name() string {
	return "TextEncoders"
}

func (t TextEncoders) Register(e *Engine) {
	e.RegisterTypeRecover("TextEncoder", func(v []goja.Value) any {
		return TextEncoder
	})
	e.RegisterTypeRecover("TextDecoder", func(v []goja.Value) any {
		return TextDecoder
	})
}

var (
	TextEncoder = &EsTextEncoder{"utf-8"}
	TextDecoder = &EsTextDecoder{"utf-8"}
)

type EsTextEncoder struct {
	Encoding string
}

func (t EsTextEncoder) Encode(v string) []byte {
	return []byte(v)
}
func (t EsTextEncoder) EncodeInto(v string, arr []byte) {
	b := []byte(v)
	for i := range arr {
		if i < len(b) {
			arr[i] = b[i]
		} else {
			break
		}
	}
}

type EsTextDecoder struct {
	Encoding string
}

// Decode typedArray are impl as bytes in goja
func (t EsTextDecoder) Decode(v []byte) string {
	return string(v)
}
func (t EsTextDecoder) EncodeInto(v string, arr []byte) {
	b := []byte(v)
	for i := range arr {
		if i < len(b) {
			arr[i] = b[i]
		} else {
			break
		}
	}
}
