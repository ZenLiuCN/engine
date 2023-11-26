package engine

import (
	_ "embed"
)

//go:embed crypto.d.ts
var cryptoDefine []byte

type CryptoModule struct {
	Symmetric  Symmetric  `js:"Symmetric"`
	Asymmetric Asymmetric `js:"Asymmetric"`
}

func (s CryptoModule) Register(engine *Engine) {
	engine.Set("Symmetric", s.Symmetric)
	engine.Set("Asymmetric", s.Asymmetric)
}

func (s CryptoModule) Name() string {
	return "crypto"
}

func (s CryptoModule) TypeDefine() []byte {
	return cryptoDefine
}
