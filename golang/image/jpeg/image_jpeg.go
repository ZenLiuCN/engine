// Code generated by define_gene; DO NOT EDIT.
package jpeg

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/golang/image"
	_ "github.com/ZenLiuCN/engine/golang/io"
	"image/jpeg"
)

var (
	//go:embed image_jpeg.d.ts
	ImageJpegDefine   []byte
	ImageJpegDeclared = map[string]any{
		"decodeConfig":   jpeg.DecodeConfig,
		"DefaultQuality": jpeg.DefaultQuality,
		"encode":         jpeg.Encode,
		"decode":         jpeg.Decode,

		"emptyOptions": func() (v jpeg.Options) {
			return v
		},
		"refOptions": func() *jpeg.Options {
			var x jpeg.Options
			return &x
		},
		"refOfOptions": func(x jpeg.Options) *jpeg.Options {
			return &x
		}}
)

func init() {
	engine.RegisterModule(ImageJpegModule{})
}

type ImageJpegModule struct{}

func (S ImageJpegModule) Identity() string {
	return "golang/image/jpeg"
}
func (S ImageJpegModule) TypeDefine() []byte {
	return ImageJpegDefine
}
func (S ImageJpegModule) Exports() map[string]any {
	return ImageJpegDeclared
}
