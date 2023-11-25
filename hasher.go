package engine

import (
	"crypto"
	_ "crypto/sha512"
	_ "embed"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"github.com/ZenLiuCN/fn"
	_ "golang.org/x/crypto/blake2s"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"
	"hash"
)

var (
	//go:embed hasher.d.ts
	hasherDefine []byte
)

type Hasher struct {
	MD4         HashW `js:"MD4"`
	MD5         HashW `js:"MD5"`
	SHA1        HashW `js:"SHA1"`
	SHA224      HashW `js:"SHA224"`
	SHA256      HashW `js:"SHA256"`
	SHA384      HashW `js:"SHA384"`
	SHA512      HashW `js:"SHA512"`
	MD5SHA1     HashW `js:"MD5SHA1"`
	RIPEMD160   HashW `js:"RIPEMD160"`
	SHA3_224    HashW `js:"SHA3_224"`
	SHA3_256    HashW `js:"SHA3_256"`
	SHA3_384    HashW `js:"SHA3_384"`
	SHA3_512    HashW `js:"SHA3_512"`
	SHA512_224  HashW `js:"SHA512_224"`
	SHA512_256  HashW `js:"SHA512_256"`
	BLAKE2s_256 HashW `js:"BLAKE2s_256"`
	BLAKE2b_256 HashW `js:"BLAKE2b_256"`
	BLAKE2b_384 HashW `js:"BLAKE2b_384"`
	BLAKE2b_512 HashW `js:"BLAKE2b_512"`
}

func (h Hasher) Name() string {
	return "hasher"
}

func (h Hasher) Register(engine *Engine) {
	engine.Set("hasher", h)
	engine.Set("codec", h)
}
func (h Hasher) Base64StdEncode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func (h Hasher) Base64StdDecode(b string) []byte {
	return fn.Panic1(base64.StdEncoding.DecodeString(b))
}
func (h Hasher) Base64UrlEncode(b []byte) string {
	return base64.URLEncoding.EncodeToString(b)
}
func (h Hasher) Base64UrlDecode(b string) []byte {
	return fn.Panic1(base64.URLEncoding.DecodeString(b))
}
func (h Hasher) Base64RawStdEncode(b []byte) string {
	return base64.RawStdEncoding.EncodeToString(b)
}
func (h Hasher) Base64RawStdDecode(b string) []byte {
	return fn.Panic1(base64.RawStdEncoding.DecodeString(b))
}

func (h Hasher) Base64RawUrlEncode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}
func (h Hasher) Base64RawUrlDecode(b string) []byte {
	return fn.Panic1(base64.RawURLEncoding.DecodeString(b))
}
func (h Hasher) HexEncode(b []byte) string {
	return hex.EncodeToString(b)
}
func (h Hasher) HexDecode(b string) []byte {
	return fn.Panic1(hex.DecodeString(b))
}
func (h Hasher) Base32StdEncode(b []byte) string {
	return base32.StdEncoding.EncodeToString(b)
}
func (h Hasher) Base32StdDecode(b string) []byte {
	return fn.Panic1(base32.StdEncoding.DecodeString(b))
}
func (h Hasher) Base32HexEncode(b []byte) string {
	return base32.HexEncoding.EncodeToString(b)
}
func (h Hasher) Base32HexDecode(b string) []byte {
	return fn.Panic1(base32.HexEncoding.DecodeString(b))
}
func (h Hasher) TypeDefine() []byte {
	return hasherDefine
}

type HashW crypto.Hash

func (h HashW) Get() hash.Hash {
	return crypto.Hash(h).New()
}

var HasherInstance = Hasher{
	MD4:         HashW(crypto.MD4),
	MD5:         HashW(crypto.MD5),
	SHA1:        HashW(crypto.SHA1),
	SHA224:      HashW(crypto.SHA224),
	SHA256:      HashW(crypto.SHA256),
	SHA384:      HashW(crypto.SHA384),
	SHA512:      HashW(crypto.SHA512),
	MD5SHA1:     HashW(crypto.MD5SHA1),
	RIPEMD160:   HashW(crypto.RIPEMD160),
	SHA3_224:    HashW(crypto.SHA3_224),
	SHA3_256:    HashW(crypto.SHA3_256),
	SHA3_384:    HashW(crypto.SHA3_384),
	SHA3_512:    HashW(crypto.SHA3_512),
	SHA512_224:  HashW(crypto.SHA512_224),
	SHA512_256:  HashW(crypto.SHA512_256),
	BLAKE2s_256: HashW(crypto.BLAKE2s_256),
	BLAKE2b_256: HashW(crypto.BLAKE2b_256),
	BLAKE2b_384: HashW(crypto.BLAKE2b_384),
	BLAKE2b_512: HashW(crypto.BLAKE2b_512),
}
