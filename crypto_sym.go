package engine

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"github.com/ZenLiuCN/fn"
)

type Symmetric struct {
}

func (Symmetric) Xaes() CipherFunc {
	return func(bytes []byte) cipher.Block {
		return fn.Panic1(aes.NewCipher(bytes))
	}
}
func (Symmetric) Xdes() CipherFunc {
	return func(bytes []byte) cipher.Block {
		return fn.Panic1(des.NewCipher(bytes))
	}
}
func (Symmetric) Xcbc(iv []byte, encrypt bool) BlockModeFunc {
	if encrypt {
		return func(block cipher.Block) cipher.BlockMode {
			return cipher.NewCBCEncrypter(block, iv)
		}
	} else {
		return func(block cipher.Block) cipher.BlockMode {
			return cipher.NewCBCDecrypter(block, iv)
		}
	}

}
func (Symmetric) Xecb(encrypt bool) BlockModeFunc {
	if encrypt {
		return func(block cipher.Block) cipher.BlockMode {
			return NewECBEncrypter(block)
		}
	} else {
		return func(block cipher.Block) cipher.BlockMode {
			return NewECBDecrypter(block)
		}
	}

}
func (Symmetric) Xcfb(iv []byte, encrypt bool) StreamFunc {
	if encrypt {
		return func(block cipher.Block) cipher.Stream {
			return cipher.NewCFBEncrypter(block, iv)
		}
	} else {
		return func(block cipher.Block) cipher.Stream {
			return cipher.NewCFBDecrypter(block, iv)
		}
	}

}
func (Symmetric) Xctr(iv []byte) StreamFunc {
	return func(block cipher.Block) cipher.Stream {
		return cipher.NewCTR(block, iv)
	}
}
func (Symmetric) Xofb(iv []byte) StreamFunc {
	return func(block cipher.Block) cipher.Stream {
		return cipher.NewOFB(block, iv)
	}
}
func (Symmetric) Xgcm() AEADFunc {
	return func(block cipher.Block) cipher.AEAD {
		return fn.Panic1(cipher.NewGCM(block))
	}
}
func (Symmetric) Xpkcs7() BlockPadding {
	return PKCS7(0)
}
func (Symmetric) Xpkcs5() BlockPadding {
	return PKCS5(0)
}
func (Symmetric) Cipher(cipher *Cipher) *Cipher {
	return cipher
}

type CipherFunc func([]byte) cipher.Block
type BlockModeFunc func(cipher.Block) cipher.BlockMode
type StreamFunc func(cipher.Block) cipher.Stream
type AEADFunc func(cipher.Block) cipher.AEAD
type BlockPadding interface {
	Padding(data []byte, blockSize int) []byte
	UnPadding(data []byte, blockSize int) []byte
}
type Cipher struct {
	Padding BlockPadding
	Cipher  CipherFunc
	Block   BlockModeFunc
	Stream  StreamFunc
	AEAD    AEADFunc `js:"aead"`
	Nonce   []byte
	Label   []byte
	Encrypt bool
}

func (c *Cipher) build() func(key, data []byte) []byte {
	if c.Cipher == nil {
		panic("cipher required")
	}
	if c.Encrypt {
		if c.Block != nil {
			return func(key, data []byte) []byte {
				b := c.Block(c.Cipher(key))
				if c.Padding != nil {
					data = c.Padding.Padding(data, b.BlockSize())
				}
				out := make([]byte, len(data))
				b.CryptBlocks(out, data)
				return out
			}
		}
		if c.Stream != nil {
			return func(key, data []byte) []byte {
				b := c.Stream(c.Cipher(key))
				out := make([]byte, len(data))
				b.XORKeyStream(out, data)
				return out
			}
		}
		if c.AEAD != nil {
			return func(key, data []byte) []byte {
				out := make([]byte, len(data))
				c.AEAD(c.Cipher(key)).Seal(out, c.Nonce, data, c.Label)
				return out
			}
		}
		return func(key, data []byte) []byte {
			b := c.Cipher(key)
			if c.Padding != nil {
				data = c.Padding.Padding(data, b.BlockSize())
			}
			out := make([]byte, len(data))
			b.Encrypt(out, c.Label)
			return out
		}
	} else {
		if c.Block != nil {
			return func(key, data []byte) []byte {
				out := make([]byte, len(data))
				b := c.Block(c.Cipher(key))
				b.CryptBlocks(out, data)
				if c.Padding != nil {
					out = c.Padding.UnPadding(out, b.BlockSize())
				}
				return out
			}
		}
		if c.Stream != nil {
			return func(key, data []byte) []byte {
				out := make([]byte, len(data))
				c.Stream(c.Cipher(key)).XORKeyStream(out, data)
				return out
			}
		}
		if c.AEAD != nil {
			return func(key, data []byte) []byte {
				out := make([]byte, len(data))
				c.AEAD(c.Cipher(key)).Seal(out, c.Nonce, data, c.Label)
				return out
			}
		}
		return func(key, data []byte) []byte {
			out := make([]byte, len(data))
			b := c.Cipher(key)
			b.Decrypt(out, c.Label)
			if c.Padding != nil {
				out = c.Padding.UnPadding(out, b.BlockSize())
			}
			return out
		}
	}
}

func (c *Cipher) Crypto(key, data []byte) []byte {
	return c.build()(key, data)
}
