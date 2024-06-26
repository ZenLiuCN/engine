// Code generated by define_gene; DO NOT EDIT.
package hex

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"encoding/hex"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
)

var (
	//go:embed encoding_hex.d.ts
	EncodingHexDefine   []byte
	EncodingHexDeclared = map[string]any{
		"newDecoder":     hex.NewDecoder,
		"decode":         hex.Decode,
		"decodedLen":     hex.DecodedLen,
		"dumper":         hex.Dumper,
		"encode":         hex.Encode,
		"encodeToString": hex.EncodeToString,
		"encodedLen":     hex.EncodedLen,
		"ErrLength":      hex.ErrLength,
		"decodeString":   hex.DecodeString,
		"dump":           hex.Dump,
		"newEncoder":     hex.NewEncoder,
	}
)

func init() {
	engine.RegisterModule(EncodingHexModule{})
}

type EncodingHexModule struct{}

func (S EncodingHexModule) Identity() string {
	return "golang/encoding/hex"
}
func (S EncodingHexModule) TypeDefine() []byte {
	return EncodingHexDefine
}
func (S EncodingHexModule) Exports() map[string]any {
	return EncodingHexDeclared
}
