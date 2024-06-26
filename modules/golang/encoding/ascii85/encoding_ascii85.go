// Code generated by define_gene; DO NOT EDIT.
package ascii85

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"encoding/ascii85"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
)

var (
	//go:embed encoding_ascii85.d.ts
	EncodingAscii85Define   []byte
	EncodingAscii85Declared = map[string]any{
		"decode":        ascii85.Decode,
		"encode":        ascii85.Encode,
		"maxEncodedLen": ascii85.MaxEncodedLen,
		"newDecoder":    ascii85.NewDecoder,
		"newEncoder":    ascii85.NewEncoder,
	}
)

func init() {
	engine.RegisterModule(EncodingAscii85Module{})
}

type EncodingAscii85Module struct{}

func (S EncodingAscii85Module) Identity() string {
	return "golang/encoding/ascii85"
}
func (S EncodingAscii85Module) TypeDefine() []byte {
	return EncodingAscii85Define
}
func (S EncodingAscii85Module) Exports() map[string]any {
	return EncodingAscii85Declared
}
