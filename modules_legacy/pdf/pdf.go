package pdf

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/dop251/goja"
)

var (
	//go:embed pdf.d.ts
	pdfDefine []byte
)

type PDF struct {
}

func (P PDF) TypeDefine() []byte {
	return pdfDefine
}

func (P PDF) Name() string {
	return "Pdf"
}

func (P PDF) Register(engine *engine.Engine) {
	engine.ToConstructor(func(v []goja.Value) (any, error) {
		return nil, nil
	})
}
