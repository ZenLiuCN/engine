// Code generated by define_gene; DO NOT EDIT.
package flate

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"compress/flate"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
)

var (
	//go:embed compress_flate.d.ts
	CompressFlateDefine   []byte
	CompressFlateDeclared = map[string]any{
		"DefaultCompression": flate.DefaultCompression,
		"HuffmanOnly":        flate.HuffmanOnly,
		"newReader":          flate.NewReader,
		"BestCompression":    flate.BestCompression,
		"BestSpeed":          flate.BestSpeed,
		"newWriterDict":      flate.NewWriterDict,
		"NoCompression":      flate.NoCompression,
		"newReaderDict":      flate.NewReaderDict,
		"newWriter":          flate.NewWriter,
	}
)

func init() {
	engine.RegisterModule(CompressFlateModule{})
}

type CompressFlateModule struct{}

func (S CompressFlateModule) Identity() string {
	return "golang/compress/flate"
}
func (S CompressFlateModule) TypeDefine() []byte {
	return CompressFlateDefine
}
func (S CompressFlateModule) Exports() map[string]any {
	return CompressFlateDeclared
}
