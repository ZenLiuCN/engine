// Code generated by define_gene; DO NOT EDIT.
package bzip2

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"compress/bzip2"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
)

var (
	//go:embed compress_bzip2.d.ts
	CompressBzip2Define   []byte
	CompressBzip2Declared = map[string]any{
		"newReader": bzip2.NewReader,
	}
)

func init() {
	engine.RegisterModule(CompressBzip2Module{})
}

type CompressBzip2Module struct{}

func (S CompressBzip2Module) Identity() string {
	return "golang/compress/bzip2"
}
func (S CompressBzip2Module) TypeDefine() []byte {
	return CompressBzip2Define
}
func (S CompressBzip2Module) Exports() map[string]any {
	return CompressBzip2Declared
}
