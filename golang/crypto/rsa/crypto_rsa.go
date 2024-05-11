// Code generated by define_gene; DO NOT EDIT.
package rsa

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"crypto/rsa"
	_ "github.com/ZenLiuCN/engine/golang/crypto"
	_ "github.com/ZenLiuCN/engine/golang/hash"
	_ "github.com/ZenLiuCN/engine/golang/io"
	_ "github.com/ZenLiuCN/engine/golang/math/big"
)

var (
	//go:embed crypto_rsa.d.ts
	CryptoRsaDefine   []byte
	CryptoRsaDeclared = map[string]any{
		"decryptPKCS1v15":           rsa.DecryptPKCS1v15,
		"ErrVerification":           rsa.ErrVerification,
		"PSSSaltLengthAuto":         rsa.PSSSaltLengthAuto,
		"signPSS":                   rsa.SignPSS,
		"verifyPSS":                 rsa.VerifyPSS,
		"decryptOAEP":               rsa.DecryptOAEP,
		"encryptPKCS1v15":           rsa.EncryptPKCS1v15,
		"ErrMessageTooLong":         rsa.ErrMessageTooLong,
		"generateMultiPrimeKey":     rsa.GenerateMultiPrimeKey,
		"PSSSaltLengthEqualsHash":   rsa.PSSSaltLengthEqualsHash,
		"signPKCS1v15":              rsa.SignPKCS1v15,
		"verifyPKCS1v15":            rsa.VerifyPKCS1v15,
		"encryptOAEP":               rsa.EncryptOAEP,
		"ErrDecryption":             rsa.ErrDecryption,
		"generateKey":               rsa.GenerateKey,
		"decryptPKCS1v15SessionKey": rsa.DecryptPKCS1v15SessionKey,

		"emptyPublicKey": func() (v rsa.PublicKey) {
			return v
		},
		"refPublicKey": func() *rsa.PublicKey {
			var x rsa.PublicKey
			return &x
		},
		"refOfPublicKey": func(x rsa.PublicKey) *rsa.PublicKey {
			return &x
		},
		"emptyCRTValue": func() (v rsa.CRTValue) {
			return v
		},
		"refCRTValue": func() *rsa.CRTValue {
			var x rsa.CRTValue
			return &x
		},
		"refOfCRTValue": func(x rsa.CRTValue) *rsa.CRTValue {
			return &x
		},
		"emptyOAEPOptions": func() (v rsa.OAEPOptions) {
			return v
		},
		"refOAEPOptions": func() *rsa.OAEPOptions {
			var x rsa.OAEPOptions
			return &x
		},
		"refOfOAEPOptions": func(x rsa.OAEPOptions) *rsa.OAEPOptions {
			return &x
		},
		"emptyPKCS1v15DecryptOptions": func() (v rsa.PKCS1v15DecryptOptions) {
			return v
		},
		"refPKCS1v15DecryptOptions": func() *rsa.PKCS1v15DecryptOptions {
			var x rsa.PKCS1v15DecryptOptions
			return &x
		},
		"refOfPKCS1v15DecryptOptions": func(x rsa.PKCS1v15DecryptOptions) *rsa.PKCS1v15DecryptOptions {
			return &x
		},
		"emptyPSSOptions": func() (v rsa.PSSOptions) {
			return v
		},
		"refPSSOptions": func() *rsa.PSSOptions {
			var x rsa.PSSOptions
			return &x
		},
		"refOfPSSOptions": func(x rsa.PSSOptions) *rsa.PSSOptions {
			return &x
		},
		"emptyPrecomputedValues": func() (v rsa.PrecomputedValues) {
			return v
		},
		"refPrecomputedValues": func() *rsa.PrecomputedValues {
			var x rsa.PrecomputedValues
			return &x
		},
		"refOfPrecomputedValues": func(x rsa.PrecomputedValues) *rsa.PrecomputedValues {
			return &x
		},
		"emptyPrivateKey": func() (v rsa.PrivateKey) {
			return v
		},
		"refPrivateKey": func() *rsa.PrivateKey {
			var x rsa.PrivateKey
			return &x
		},
		"refOfPrivateKey": func(x rsa.PrivateKey) *rsa.PrivateKey {
			return &x
		}}
)

func init() {
	engine.RegisterModule(CryptoRsaModule{})
}

type CryptoRsaModule struct{}

func (S CryptoRsaModule) Identity() string {
	return "golang/crypto/rsa"
}
func (S CryptoRsaModule) TypeDefine() []byte {
	return CryptoRsaDefine
}
func (S CryptoRsaModule) Exports() map[string]any {
	return CryptoRsaDeclared
}
