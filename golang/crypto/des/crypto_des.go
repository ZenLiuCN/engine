// Code generated by define_gene; DO NOT EDIT.
package des

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"crypto/des"
	_ "github.com/ZenLiuCN/engine/golang/crypto/cipher"
)

var (
	//go:embed crypto_des.d.ts
	CryptoDesDefine   []byte
	CryptoDesDeclared = map[string]any{
		"BlockSize":          des.BlockSize,
		"newCipher":          des.NewCipher,
		"newTripleDESCipher": des.NewTripleDESCipher,
	}
)

func init() {
	engine.RegisterModule(CryptoDesModule{})
}

type CryptoDesModule struct{}

func (S CryptoDesModule) Identity() string {
	return "golang/crypto/des"
}
func (S CryptoDesModule) TypeDefine() []byte {
	return CryptoDesDefine
}
func (S CryptoDesModule) Exports() map[string]any {
	return CryptoDesDeclared
}
