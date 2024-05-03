package pdf_reader

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/dslipak/pdf"
)

var (
	//go:embed pdf_reader.d.ts
	define []byte

	declare = map[string]any{
		"open":               pdf.Open,
		"newReader":          pdf.NewReader,
		"newReaderEncrypted": pdf.NewReaderEncrypted,
	}
)

func init() {
	engine.RegisterModule(Module{})
}

type Module struct {
}

func (m Module) TypeDefine() []byte {
	return define
}

func (m Module) Identity() string {
	return "go/pdf/reader"
}

func (m Module) Exports() map[string]any {
	return declare
}
