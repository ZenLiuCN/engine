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
)

type HashW crypto.Hash

func (h HashW) Get() hash.Hash {
	return crypto.Hash(h).New()
}

type HashModule struct {
	m map[string]any
}

func (h *HashModule) Identity() string {
	return "go/hash"
}

func (h *HashModule) TypeDefine() []byte {
	return hashDefine
}

func (h *HashModule) Exports() map[string]any {
	if len(h.m) == 0 {
		h.init()
	}
	return h.m
}

func (h *HashModule) init() {
	h.m = make(map[string]any)
	h.m["MD4"] = HashW(crypto.MD4)
	h.m["MD5"] = HashW(crypto.MD5)
	h.m["SHA1"] = HashW(crypto.SHA1)
	h.m["SHA224"] = HashW(crypto.SHA224)
	h.m["SHA256"] = HashW(crypto.SHA256)
	h.m["SHA384"] = HashW(crypto.SHA384)
	h.m["SHA512"] = HashW(crypto.SHA512)
	h.m["MD5SHA1"] = HashW(crypto.MD5SHA1)
	h.m["RIPEMD160"] = HashW(crypto.RIPEMD160)
	h.m["SHA3_224"] = HashW(crypto.SHA3_224)
	h.m["SHA3_256"] = HashW(crypto.SHA3_256)
	h.m["SHA3_384"] = HashW(crypto.SHA3_384)
	h.m["SHA3_512"] = HashW(crypto.SHA3_512)
	h.m["SHA512_224"] = HashW(crypto.SHA512_224)
	h.m["SHA512_256"] = HashW(crypto.SHA512_256)
	h.m["BLAKE2s_256"] = HashW(crypto.BLAKE2s_256)
	h.m["BLAKE2b_256"] = HashW(crypto.BLAKE2b_256)
	h.m["BLAKE2b_384"] = HashW(crypto.BLAKE2b_384)
	h.m["BLAKE2b_512"] = HashW(crypto.BLAKE2b_512)
}
