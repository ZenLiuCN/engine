package engine

import (
	"crypto"
	_ "crypto/sha512"
	_ "embed"
	_ "golang.org/x/crypto/blake2s"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"
	"hash"
)

var (
	//go:embed module_hash.d.ts
	hashDefine []byte
	hashMap    = map[string]any{
		"MD4":         HashW(crypto.MD4),
		"MD5":         HashW(crypto.MD5),
		"SHA1":        HashW(crypto.SHA1),
		"SHA224":      HashW(crypto.SHA224),
		"SHA256":      HashW(crypto.SHA256),
		"SHA384":      HashW(crypto.SHA384),
		"SHA512":      HashW(crypto.SHA512),
		"MD5SHA1":     HashW(crypto.MD5SHA1),
		"RIPEMD160":   HashW(crypto.RIPEMD160),
		"SHA3_224":    HashW(crypto.SHA3_224),
		"SHA3_256":    HashW(crypto.SHA3_256),
		"SHA3_384":    HashW(crypto.SHA3_384),
		"SHA3_512":    HashW(crypto.SHA3_512),
		"SHA512_224":  HashW(crypto.SHA512_224),
		"SHA512_256":  HashW(crypto.SHA512_256),
		"BLAKE2s_256": HashW(crypto.BLAKE2s_256),
		"BLAKE2b_256": HashW(crypto.BLAKE2b_256),
		"BLAKE2b_384": HashW(crypto.BLAKE2b_384),
		"BLAKE2b_512": HashW(crypto.BLAKE2b_512),
	}
)

type HashW crypto.Hash

func (h HashW) Get() hash.Hash {
	return crypto.Hash(h).New()
}

type HashModule struct {
	m map[string]any
}

func (h HashModule) Identity() string {
	return "go/hash"
}

func (h HashModule) TypeDefine() []byte {
	return hashDefine
}

func (h HashModule) Exports() map[string]any {

	return hashMap
}
