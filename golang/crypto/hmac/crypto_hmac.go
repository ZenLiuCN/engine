// Code generated by define_gene; DO NOT EDIT.
package hmac

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"crypto/hmac"
	_ "github.com/ZenLiuCN/engine/golang/hash"
)

var (
	//go:embed crypto_hmac.d.ts
	CryptoHmacDefine   []byte
	CryptoHmacDeclared = map[string]any{
		"equal": hmac.Equal,
		"New":   hmac.New,
	}
)

func init() {
	engine.RegisterModule(CryptoHmacModule{})
}

type CryptoHmacModule struct{}

func (S CryptoHmacModule) Identity() string {
	return "golang/crypto/hmac"
}
func (S CryptoHmacModule) TypeDefine() []byte {
	return CryptoHmacDefine
}
func (S CryptoHmacModule) Exports() map[string]any {
	return CryptoHmacDeclared
}
