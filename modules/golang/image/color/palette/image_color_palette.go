// Code generated by define_gene; DO NOT EDIT.
package palette

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/image/color"
	"image/color/palette"
)

var (
	//go:embed image_color_palette.d.ts
	ImageColorPaletteDefine   []byte
	ImageColorPaletteDeclared = map[string]any{
		"Plan9":   palette.Plan9,
		"WebSafe": palette.WebSafe,
	}
)

func init() {
	engine.RegisterModule(ImageColorPaletteModule{})
}

type ImageColorPaletteModule struct{}

func (S ImageColorPaletteModule) Identity() string {
	return "golang/image/color/palette"
}
func (S ImageColorPaletteModule) TypeDefine() []byte {
	return ImageColorPaletteDefine
}
func (S ImageColorPaletteModule) Exports() map[string]any {
	return ImageColorPaletteDeclared
}
