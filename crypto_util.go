package engine

import (
	"bytes"
	"crypto/cipher"
)

type PKCS7 byte

func (p PKCS7) Padding(text []byte, blockSize int) []byte {
	padding := byte(blockSize - len(text)%blockSize)
	var paddingText []byte
	if padding == 0 {
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		paddingText = bytes.Repeat([]byte{padding}, int(padding))
	}
	return append(text, paddingText...)
}
func (p PKCS7) UnPadding(text []byte, blockSize int) []byte {
	if textLen := len(text); textLen != 0 {
		padLen := int(text[textLen-1])
		if padLen >= textLen || padLen > blockSize {
			return []byte{}
		}
		return text[:textLen-padLen]
	}
	return []byte{}
}

type PKCS5 byte

func (p PKCS5) Padding(text []byte, blockSize int) []byte {
	padding := 8 - len(text)%8
	var paddingText []byte
	if padding == 0 {
		paddingText = bytes.Repeat([]byte{byte(8)}, 8)
	} else {
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}
	return append(text, paddingText...)
}

func (p PKCS5) UnPadding(text []byte, blockSize int) []byte {
	if textLen := len(text); textLen != 0 {
		padLen := int(text[textLen-1])
		if padLen >= textLen || padLen > 8 {
			return []byte{}
		}
		return text[:textLen-padLen]
	}
	return []byte{}
}

// region ECB
type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ECBEncrypt ecb

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ECBEncrypt)(newECB(b))
}
func (x *ECBEncrypt) BlockSize() int { return x.blockSize }
func (x *ECBEncrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ECBDecrypt ecb

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ECBDecrypt)(newECB(b))
}
func (x *ECBDecrypt) BlockSize() int { return x.blockSize }
func (x *ECBDecrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

//endregion
