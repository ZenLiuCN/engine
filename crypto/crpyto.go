package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"github.com/google/uuid"
)

var (
	//go:embed crypto.d.ts
	d []byte
)

type Crypto struct {
	*engine.Engine
}

func (c *Crypto) TypeDefine() []byte {
	return d
}
func (c *Crypto) Name() string {
	return "crypto"
}

func (c *Crypto) Initialize(engine *engine.Engine) engine.Module {
	cx := &Crypto{Engine: engine}
	return cx
}

func (c *Crypto) RandomUUID() string {
	return uuid.New().String()
}
func (c *Crypto) GetRandomValues(buf goja.ArrayBuffer) goja.Value {
	fn.Panic1(rand.Read(buf.Bytes()))
	return c.Engine.ToValue(buf)
}

func (c *Crypto) GenerateRsaKey(bits int) *PrivateKey {
	return &PrivateKey{rsa: fn.Panic1(rsa.GenerateKey(rand.Reader, bits))}
}

func (c *Crypto) Cipher(bits int) crypto.Hash {
	return crypto.MD5
}

type PrivateKey struct {
	rsa *rsa.PrivateKey
}
type Cipher struct {
}
