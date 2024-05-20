// Code generated by define_gene; DO NOT EDIT.
package gif

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/image"
	_ "github.com/ZenLiuCN/engine/modules/golang/image/draw"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	"image/gif"
)

var (
	//go:embed image_gif.d.ts
	ImageGifDefine   []byte
	ImageGifDeclared = map[string]any{
		"encodeAll":          gif.EncodeAll,
		"decode":             gif.Decode,
		"decodeAll":          gif.DecodeAll,
		"decodeConfig":       gif.DecodeConfig,
		"DisposalBackground": gif.DisposalBackground,
		"DisposalNone":       gif.DisposalNone,
		"DisposalPrevious":   gif.DisposalPrevious,
		"encode":             gif.Encode,

		"emptyGIF":        engine.Empty[gif.GIF],
		"emptyRefGIF":     engine.EmptyRefer[gif.GIF],
		"refOfGIF":        engine.ReferOf[gif.GIF],
		"unRefGIF":        engine.UnRefer[gif.GIF],
		"emptyOptions":    engine.Empty[gif.Options],
		"emptyRefOptions": engine.EmptyRefer[gif.Options],
		"refOfOptions":    engine.ReferOf[gif.Options],
		"unRefOptions":    engine.UnRefer[gif.Options]}
)

func init() {
	engine.RegisterModule(ImageGifModule{})
}

type ImageGifModule struct{}

func (S ImageGifModule) Identity() string {
	return "golang/image/gif"
}
func (S ImageGifModule) TypeDefine() []byte {
	return ImageGifDefine
}
func (S ImageGifModule) Exports() map[string]any {
	return ImageGifDeclared
}