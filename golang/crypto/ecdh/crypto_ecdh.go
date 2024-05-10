// Code generated by define_gene; DO NOT EDIT.
package ecdh

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"crypto/ecdh"
	_ "github.com/ZenLiuCN/engine/golang/crypto"
	_ "github.com/ZenLiuCN/engine/golang/io"
)

var (
	//go:embed crypto_ecdh.d.ts
	CryptoEcdhDefine   []byte
	CryptoEcdhDeclared = map[string]any{
		"p256":   ecdh.P256,
		"p384":   ecdh.P384,
		"p521":   ecdh.P521,
		"x25519": ecdh.X25519,

		"emptyPrivateKey": func() (v ecdh.PrivateKey) {
			return v
		},
		"refPrivateKey": func() *ecdh.PrivateKey {
			var x ecdh.PrivateKey
			return &x
		},
		"refOfPrivateKey": func(x ecdh.PrivateKey) *ecdh.PrivateKey {
			return &x
		},
		"emptyPublicKey": func() (v ecdh.PublicKey) {
			return v
		},
		"refPublicKey": func() *ecdh.PublicKey {
			var x ecdh.PublicKey
			return &x
		},
		"refOfPublicKey": func(x ecdh.PublicKey) *ecdh.PublicKey {
			return &x
		}}
)

func init() {
	engine.RegisterModule(CryptoEcdhModule{})
}

type CryptoEcdhModule struct{}

func (S CryptoEcdhModule) Identity() string {
	return "golang/crypto/ecdh"
}
func (S CryptoEcdhModule) TypeDefine() []byte {
	return CryptoEcdhDefine
}
func (S CryptoEcdhModule) Exports() map[string]any {
	return CryptoEcdhDeclared
}
