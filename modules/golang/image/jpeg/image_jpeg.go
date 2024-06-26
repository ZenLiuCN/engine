// Code generated by define_gene; DO NOT EDIT.
package jpeg

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/image"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	"image/jpeg"
)

var (
	//go:embed image_jpeg.d.ts
	ImageJpegDefine   []byte
	ImageJpegDeclared = map[string]any{
		"decode":         jpeg.Decode,
		"decodeConfig":   jpeg.DecodeConfig,
		"DefaultQuality": jpeg.DefaultQuality,
		"encode":         jpeg.Encode,

		"emptyOptions":    engine.Empty[jpeg.Options],
		"emptyRefOptions": engine.EmptyRefer[jpeg.Options],
		"refOfOptions":    engine.ReferOf[jpeg.Options],
		"unRefOptions":    engine.UnRefer[jpeg.Options]}
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
