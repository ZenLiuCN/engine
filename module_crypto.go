package engine

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"errors"
	"fmt"
	"github.com/ZenLiuCN/engine/eciesgo"
	"github.com/ZenLiuCN/fn"
	"hash"
)

//go:embed module_crypto.d.ts
var cryptoDefine []byte

type CryptoModule struct {
	m map[string]any
}

func (s *CryptoModule) Identity() string {
	return "go/crypto"
}

func (s *CryptoModule) Exports() map[string]any {
	if s.m == nil {
		s.m = map[string]any{}
		s.m["aes"] = func() CipherFunc {
			return func(bytes []byte) cipher.Block {
				return fn.Panic1(aes.NewCipher(bytes))
			}
		}
		s.m["des"] = func() CipherFunc {
			return func(bytes []byte) cipher.Block {
				return fn.Panic1(des.NewCipher(bytes))
			}
		}
		s.m["ecb"] = func(encrypt bool) BlockModeFunc {
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
		s.m["cbc"] = func(iv []byte, encrypt bool) BlockModeFunc {
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
		s.m["cfb"] = func(iv []byte, encrypt bool) StreamFunc {
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
		s.m["ctr"] = func(iv []byte) StreamFunc {
			return func(block cipher.Block) cipher.Stream {
				return cipher.NewCTR(block, iv)
			}
		}
		s.m["ofb"] = func(iv []byte) StreamFunc {
			return func(block cipher.Block) cipher.Stream {
				return cipher.NewOFB(block, iv)
			}
		}
		s.m["gcm"] = func() AEADFunc {
			return func(block cipher.Block) cipher.AEAD {
				return fn.Panic1(cipher.NewGCM(block))
			}
		}
		s.m["pkcs7"] = func() BlockPadding {
			return PKCS7(0)
		}
		s.m["pkcs5"] = func() BlockPadding {
			return PKCS5(0)
		}
		s.m["cipher"] = func(cipher *Cipher) *Cipher {
			return cipher
		}

		s.m["generateKey"] = func(algorithm Algorithm, option Option) *AsymmetricPrivateKey {
			switch algorithm {
			case algRSA:
				return &AsymmetricPrivateKey{
					Alg: algorithm,
					key: fn.Panic1(rsa.GenerateKey(rand.Reader, option.Bits)),
				}
			case algECDH:
				switch option.Curve {
				case "P256":
					return &AsymmetricPrivateKey{
						Alg: algorithm,
						key: fn.Panic1(ecdh.P256().GenerateKey(rand.Reader)),
					}
				case "P384":
					return &AsymmetricPrivateKey{
						Alg: algorithm,
						key: fn.Panic1(ecdh.P384().GenerateKey(rand.Reader)),
					}
				case "P521":
					return &AsymmetricPrivateKey{
						Alg: algorithm,
						key: fn.Panic1(ecdh.P521().GenerateKey(rand.Reader)),
					}
				case "X25519":
					return &AsymmetricPrivateKey{
						Alg: algorithm,
						key: fn.Panic1(ecdh.X25519().GenerateKey(rand.Reader)),
					}
				default:
					panic(fmt.Errorf("unknown ECDH curve type : %s", option.Curve))
				}
			case algECDSA:
				switch option.Curve {
				case "P256":
					return &AsymmetricPrivateKey{
						Alg: algorithm,
						key: fn.Panic1(ecdsa.GenerateKey(elliptic.P256(), rand.Reader)),
					}
				case "P384":
					return &AsymmetricPrivateKey{
						Alg: algorithm,
						key: fn.Panic1(ecdsa.GenerateKey(elliptic.P384(), rand.Reader)),
					}
				case "P521":
					return &AsymmetricPrivateKey{
						Alg: algorithm,
						key: fn.Panic1(ecdsa.GenerateKey(elliptic.P521(), rand.Reader)),
					}
				case "P224":
					return &AsymmetricPrivateKey{
						Alg: algorithm,
						key: fn.Panic1(ecdsa.GenerateKey(elliptic.P224(), rand.Reader)),
					}
				default:
					panic(fmt.Errorf("unknown ECDSA curve type : %s", option.Curve))
				}
			case algED25519:
				pk, puk := fn.Panic2(ed25519.GenerateKey(rand.Reader))
				return &AsymmetricPrivateKey{
					Alg:    algorithm,
					key:    pk,
					pubKey: puk,
				}
			default:
				panic(fmt.Errorf("unknown algorithm type : %d", algorithm))
			}
		}
		s.m["parsePrivateKey"] = func(pem []byte) *AsymmetricPrivateKey {
			p := new(AsymmetricPrivateKey)
			p.Load(pem)
			return p
		}
		s.m["parsePublicKey"] = func(pem []byte) *AsymmetricPublicKey {
			p := new(AsymmetricPublicKey)
			p.Load(pem)
			return p
		}
		s.m["sign"] = func(key *AsymmetricPrivateKey, data []byte, hash crypto.Hash, opt *rsa.PSSOptions) []byte {
			switch key.Alg {
			case algRSA:
				if opt == nil {
					return fn.Panic1(rsa.SignPKCS1v15(rand.Reader, key.mustRSA(), hash, data))
				} else {
					return fn.Panic1(rsa.SignPSS(rand.Reader, key.mustRSA(), hash, data, opt))
				}
			case algECDSA:
				return fn.Panic1(ecdsa.SignASN1(rand.Reader, key.mustECDSA(), data))
			case algED25519:
				return ed25519.Sign(key.mustED25519(), data)
			default:
				panic(fmt.Errorf("unsupported algorithm for sign: %d", key.Alg))
			}
		}
		s.m["verify"] = func(key *AsymmetricPublicKey, data, sign []byte, hash crypto.Hash, opt *rsa.PSSOptions) bool {
			switch key.Alg {
			case algRSA:
				if opt == nil {
					err := rsa.VerifyPKCS1v15(key.mustRSA(), hash, data, sign)
					if err != nil {
						if errors.Is(err, rsa.ErrVerification) {
							return false
						} else {
							panic(err)
						}
					} else {
						return true
					}
				} else {
					err := rsa.VerifyPSS(key.mustRSA(), hash, data, sign, opt)
					if err != nil {
						if errors.Is(err, rsa.ErrVerification) {
							return false
						} else {
							panic(err)
						}
					} else {
						return true
					}
				}
			case algECDSA:
				return ecdsa.VerifyASN1(key.mustECDSA(), data, sign)
			case algED25519:
				return ed25519.Verify(key.mustED25519(), data, sign)
			default:
				panic(fmt.Errorf("unsupported algorithm for sign: %d", key.Alg))
			}
		}
		s.m["encrypt"] = func(key *AsymmetricPublicKey, data []byte, hasher hash.Hash) []byte {
			switch key.Alg {
			case algRSA:
				return fn.Panic1(rsa.EncryptOAEP(hasher, rand.Reader, key.mustRSA(), data, nil))
			case algECDSA:
				return fn.Panic1(eciesgo.Encrypt(key.mustECISE(), data))
			default:
				panic(fmt.Errorf("unsupported algorithm for sign: %d", key.Alg))
			}
		}
		s.m["decrypt"] = func(key *AsymmetricPrivateKey, secret []byte, hasher hash.Hash) []byte {
			switch key.Alg {
			case algRSA:
				return fn.Panic1(rsa.DecryptOAEP(hasher, rand.Reader, key.mustRSA(), secret, nil))
			case algECDSA:
				return fn.Panic1(eciesgo.Decrypt(key.mustECISE(), secret))
			default:
				panic(fmt.Errorf("unsupported algorithm for sign: %d", key.Alg))
			}
		}

	}
	return s.m
}

func (s *CryptoModule) TypeDefine() []byte {
	return cryptoDefine
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

type Algorithm int

const (
	algUndefined Algorithm = iota
	algRSA
	algECDH
	algECDSA
	algED25519
)

type AsymmetricPrivateKey struct {
	Alg    Algorithm
	key    crypto.PrivateKey
	pubKey crypto.PublicKey // only for ED25519
}

func (s *AsymmetricPrivateKey) Bytes() []byte {
	return fn.Panic1(x509.MarshalPKCS8PrivateKey(s.key))
}
func (s *AsymmetricPrivateKey) Load(pem []byte) {
	k := fn.Panic1(x509.ParsePKCS8PrivateKey(pem))
	switch v := k.(type) {
	case *rsa.PrivateKey:
		{
			s.key = v
			s.Alg = algRSA
		}
	case *ecdsa.PrivateKey:
		{
			s.key = v
			s.Alg = algECDSA
		}
	case *ecdh.PrivateKey:
		{
			s.key = v
			s.Alg = algECDH
		}
	case ed25519.PrivateKey:
		{
			s.key = v
			s.Alg = algED25519
		}
	default:
		panic(fmt.Errorf("unknown key type : %T", v))
	}
}
func (s *AsymmetricPrivateKey) Equal(other *AsymmetricPrivateKey) bool {
	if s == nil || other == nil {
		return false
	}
	if s.Alg != other.Alg {
		return false
	}
	if s.key == nil || other.key == nil {
		return false
	}
	switch k := s.key.(type) {
	case *rsa.PrivateKey:
		return k.Equal(other.key)
	case *ecdsa.PrivateKey:
		return k.Equal(other.key)
	case *ecdh.PrivateKey:
		return k.Equal(other.key)
	case ed25519.PrivateKey:
		return k.Equal(other.key) && s.pubKey.(ed25519.PublicKey).Equal(other.pubKey)
	default:
		panic(fmt.Errorf("unknown key type : %T", s.key))
	}
}
func (s *AsymmetricPrivateKey) PublicKey() *AsymmetricPublicKey {
	switch k := s.key.(type) {
	case *rsa.PrivateKey:
		return &AsymmetricPublicKey{
			Alg: s.Alg,
			key: &k.PublicKey,
		}
	case *ecdsa.PrivateKey:
		return &AsymmetricPublicKey{
			Alg: s.Alg,
			key: &k.PublicKey,
		}
	case ed25519.PrivateKey:
		return &AsymmetricPublicKey{
			Alg: s.Alg,
			key: s.pubKey,
		}
	case *ecdh.PrivateKey:
		return &AsymmetricPublicKey{
			Alg: s.Alg,
			key: k.PublicKey(),
		}
	default:
		panic(fmt.Errorf("unknown key type : %T", s.key))
	}
}
func (s *AsymmetricPrivateKey) mustRSA() *rsa.PrivateKey {
	if s.Alg != algRSA {
		panic("require a RSA key")
	}
	if k, ok := s.key.(*rsa.PrivateKey); !ok {
		panic("require a RSA key")
	} else {
		return k
	}
}

func (s *AsymmetricPrivateKey) mustECDSA() *ecdsa.PrivateKey {
	if s.Alg != algECDSA {
		panic("require a ECDSA key")
	}
	if k, ok := s.key.(*ecdsa.PrivateKey); !ok {
		panic("require a ECDSA key")
	} else {
		return k
	}
}

func (s *AsymmetricPrivateKey) mustED25519() ed25519.PrivateKey {
	if s.Alg != algED25519 {
		panic("require a ED25519 key")
	}
	if k, ok := s.key.(ed25519.PrivateKey); !ok {
		panic("require a ED25519 key")
	} else {
		return k
	}
}
func (s *AsymmetricPrivateKey) mustECDH() *ecdh.PrivateKey {
	if s.Alg != algECDH {
		panic("require a ECDH key")
	}
	if k, ok := s.key.(*ecdh.PrivateKey); !ok {
		panic("require a ECDH key")
	} else {
		return k
	}
}

func (s *AsymmetricPrivateKey) mustECISE() *eciesgo.PrivateKey {
	k := s.mustECDSA()
	return &eciesgo.PrivateKey{
		PublicKey: &eciesgo.PublicKey{
			Curve: k.Curve,
			X:     k.X,
			Y:     k.Y,
		},
		D: k.D,
	}
}

type AsymmetricPublicKey struct {
	Alg Algorithm
	key crypto.PublicKey
}

func (s *AsymmetricPublicKey) Bytes() []byte {
	return fn.Panic1(x509.MarshalPKIXPublicKey(s.key))
}
func (s *AsymmetricPublicKey) Load(pem []byte) {
	k := fn.Panic1(x509.ParsePKIXPublicKey(pem))
	switch v := k.(type) {
	case *rsa.PublicKey:
		{
			s.key = v
			s.Alg = algRSA
		}
	case *ecdsa.PublicKey:
		{
			s.key = v
			s.Alg = algECDSA
		}
	case *ecdh.PublicKey:
		{
			s.key = v
			s.Alg = algECDH
		}
	case ed25519.PublicKey:
		{
			s.key = v
			s.Alg = algED25519
		}
	default:
		panic(fmt.Errorf("unknown key type : %T", v))
	}
}
func (s *AsymmetricPublicKey) Equal(other *AsymmetricPublicKey) bool {
	if s == nil || other == nil {
		return false
	}
	if s.Alg != other.Alg {
		return false
	}
	if s.key == nil || other.key == nil {
		return false
	}
	switch k := s.key.(type) {
	case *rsa.PublicKey:
		return k.Equal(other.key)
	case *ecdsa.PublicKey:
		return k.Equal(other.key)
	case *ecdh.PublicKey:
		return k.Equal(other.key)
	case ed25519.PublicKey:
		return k.Equal(other.key)
	default:
		panic(fmt.Errorf("unknown key type : %T", s.key))
	}
}
func (s *AsymmetricPublicKey) mustRSA() *rsa.PublicKey {
	if s.Alg != algRSA {
		panic("require a RSA  key")
	}
	if k, ok := s.key.(*rsa.PublicKey); !ok {
		panic("require a RSA  key")
	} else {
		return k
	}
}

func (s *AsymmetricPublicKey) mustECDSA() *ecdsa.PublicKey {
	if s.Alg != algECDSA {
		panic("require a ECDSA key")
	}
	if k, ok := s.key.(*ecdsa.PublicKey); !ok {
		panic("require a ECDSA key")
	} else {
		return k
	}
}
func (s *AsymmetricPublicKey) mustECISE() *eciesgo.PublicKey {
	k := s.mustECDSA()
	return &eciesgo.PublicKey{
		Curve: k.Curve,
		X:     k.X,
		Y:     k.Y,
	}
}
func (s *AsymmetricPublicKey) mustED25519() ed25519.PublicKey {
	if s.Alg != algED25519 {
		panic("require a ED25519 key")
	}
	if k, ok := s.key.(ed25519.PublicKey); !ok {
		panic("require a ED25519 key")
	} else {
		return k
	}
}

func (s *AsymmetricPublicKey) mustECDH() *ecdh.PublicKey {
	if s.Alg != algECDH {
		panic("require a ECDH key")
	}
	if k, ok := s.key.(*ecdh.PublicKey); !ok {
		panic("require a ECDH key")
	} else {
		return k
	}
}

type (
	Option struct {
		Bits  int
		Curve string
	}
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
