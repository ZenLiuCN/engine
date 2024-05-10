// Code generated by define_gene; DO NOT EDIT.
package cipher

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"crypto/cipher"
	_ "github.com/ZenLiuCN/engine/golang/io"
)

var (
	//go:embed crypto_cipher.d.ts
	CryptoCipherDefine   []byte
	CryptoCipherDeclared = map[string]any{
		"newGCMWithTagSize":   cipher.NewGCMWithTagSize,
		"newCBCDecrypter":     cipher.NewCBCDecrypter,
		"newCBCEncrypter":     cipher.NewCBCEncrypter,
		"newCFBEncrypter":     cipher.NewCFBEncrypter,
		"newCTR":              cipher.NewCTR,
		"newGCM":              cipher.NewGCM,
		"newGCMWithNonceSize": cipher.NewGCMWithNonceSize,
		"newCFBDecrypter":     cipher.NewCFBDecrypter,
		"newOFB":              cipher.NewOFB,

		"emptyStreamReader": func() (v cipher.StreamReader) {
			return v
		},
		"refStreamReader": func() *cipher.StreamReader {
			var x cipher.StreamReader
			return &x
		},
		"refOfStreamReader": func(x cipher.StreamReader) *cipher.StreamReader {
			return &x
		},
		"emptyStreamWriter": func() (v cipher.StreamWriter) {
			return v
		},
		"refStreamWriter": func() *cipher.StreamWriter {
			var x cipher.StreamWriter
			return &x
		},
		"refOfStreamWriter": func(x cipher.StreamWriter) *cipher.StreamWriter {
			return &x
		}}
)

func init() {
	engine.RegisterModule(CryptoCipherModule{})
}

type CryptoCipherModule struct{}

func (S CryptoCipherModule) Identity() string {
	return "golang/crypto/cipher"
}
func (S CryptoCipherModule) TypeDefine() []byte {
	return CryptoCipherDefine
}
func (S CryptoCipherModule) Exports() map[string]any {
	return CryptoCipherDeclared
}
