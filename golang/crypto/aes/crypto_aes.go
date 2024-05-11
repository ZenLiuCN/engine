// Code generated by define_gene; DO NOT EDIT.
package aes

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"crypto/aes"
	_ "github.com/ZenLiuCN/engine/golang/crypto/cipher"
)

var (
	//go:embed crypto_aes.d.ts
	CryptoAesDefine   []byte
	CryptoAesDeclared = map[string]any{
		"newCipher": aes.NewCipher,
		"BlockSize": aes.BlockSize,
	}
)

func init() {
	engine.RegisterModule(CryptoAesModule{})
}

type CryptoAesModule struct{}

func (S CryptoAesModule) Identity() string {
	return "golang/crypto/aes"
}
func (S CryptoAesModule) TypeDefine() []byte {
	return CryptoAesDefine
}
func (S CryptoAesModule) Exports() map[string]any {
	return CryptoAesDeclared
}
