// Code generated by define_gene; DO NOT EDIT.
package base64

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"encoding/base64"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
)

var (
	//go:embed encoding_base64.d.ts
	EncodingBase64Define   []byte
	EncodingBase64Declared = map[string]any{
		"newDecoder":     base64.NewDecoder,
		"newEncoder":     base64.NewEncoder,
		"newEncoding":    base64.NewEncoding,
		"NoPadding":      base64.NoPadding,
		"StdEncoding":    base64.StdEncoding,
		"StdPadding":     base64.StdPadding,
		"URLEncoding":    base64.URLEncoding,
		"RawStdEncoding": base64.RawStdEncoding,
		"RawURLEncoding": base64.RawURLEncoding,
	}
)

func init() {
	engine.RegisterModule(EncodingBase64Module{})
}

type EncodingBase64Module struct{}

func (S EncodingBase64Module) Identity() string {
	return "golang/encoding/base64"
}
func (S EncodingBase64Module) TypeDefine() []byte {
	return EncodingBase64Define
}
func (S EncodingBase64Module) Exports() map[string]any {
	return EncodingBase64Declared
}