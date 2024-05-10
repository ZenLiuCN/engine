// Code generated by define_gene; DO NOT EDIT.
package png

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/golang/image"
	_ "github.com/ZenLiuCN/engine/golang/io"
	"image/png"
)

var (
	//go:embed image_png.d.ts
	ImagePngDefine   []byte
	ImagePngDeclared = map[string]any{
		"encode":             png.Encode,
		"NoCompression":      png.NoCompression,
		"BestCompression":    png.BestCompression,
		"BestSpeed":          png.BestSpeed,
		"decode":             png.Decode,
		"decodeConfig":       png.DecodeConfig,
		"DefaultCompression": png.DefaultCompression,

		"emptyEncoder": func() (v png.Encoder) {
			return v
		},
		"refEncoder": func() *png.Encoder {
			var x png.Encoder
			return &x
		},
		"refOfEncoder": func(x png.Encoder) *png.Encoder {
			return &x
		},
		"emptyEncoderBuffer": func() (v png.EncoderBuffer) {
			return v
		},
		"refEncoderBuffer": func() *png.EncoderBuffer {
			var x png.EncoderBuffer
			return &x
		},
		"refOfEncoderBuffer": func(x png.EncoderBuffer) *png.EncoderBuffer {
			return &x
		}}
)

func init() {
	engine.RegisterModule(ImagePngModule{})
}

type ImagePngModule struct{}

func (S ImagePngModule) Identity() string {
	return "golang/image/png"
}
func (S ImagePngModule) TypeDefine() []byte {
	return ImagePngDefine
}
func (S ImagePngModule) Exports() map[string]any {
	return ImagePngDeclared
}
