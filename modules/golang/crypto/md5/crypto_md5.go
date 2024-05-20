// Code generated by define_gene; DO NOT EDIT.
package md5

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"crypto/md5"
	_ "github.com/ZenLiuCN/engine/modules/golang/hash"
)

var (
	//go:embed crypto_md5.d.ts
	CryptoMd5Define   []byte
	CryptoMd5Declared = map[string]any{
		"New":       md5.New,
		"Size":      md5.Size,
		"sum":       md5.Sum,
		"BlockSize": md5.BlockSize,
	}
)

func init() {
	engine.RegisterModule(CryptoMd5Module{})
}

type CryptoMd5Module struct{}

func (S CryptoMd5Module) Identity() string {
	return "golang/crypto/md5"
}
func (S CryptoMd5Module) TypeDefine() []byte {
	return CryptoMd5Define
}
func (S CryptoMd5Module) Exports() map[string]any {
	return CryptoMd5Declared
}