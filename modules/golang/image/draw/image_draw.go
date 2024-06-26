// Code generated by define_gene; DO NOT EDIT.
package draw

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/image"
	_ "github.com/ZenLiuCN/engine/modules/golang/image/color"
	"image/draw"
)

var (
	//go:embed image_draw.d.ts
	ImageDrawDefine   []byte
	ImageDrawDeclared = map[string]any{
		"Src":            draw.Src,
		"draw":           draw.Draw,
		"drawMask":       draw.DrawMask,
		"FloydSteinberg": draw.FloydSteinberg,
		"Over":           draw.Over,
	}
)

func init() {
	engine.RegisterModule(ImageDrawModule{})
}

type ImageDrawModule struct{}

func (S ImageDrawModule) Identity() string {
	return "golang/image/draw"
}
func (S ImageDrawModule) TypeDefine() []byte {
	return ImageDrawDefine
}
func (S ImageDrawModule) Exports() map[string]any {
	return ImageDrawDeclared
}
