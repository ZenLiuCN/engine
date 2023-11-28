package engine

import (
	_ "embed"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"github.com/ZenLiuCN/fn"
)

var (
	//go:embed codec.d.ts
	codecDefine []byte
)

type CodecModule struct {
	m map[string]any
}

func (h *CodecModule) Identity() string {
	return "go/codec"
}

func (h *CodecModule) TypeDefine() []byte {
	return codecDefine
}

func (h *CodecModule) Exports() map[string]any {
	if len(h.m) == 0 {
		h.m = make(map[string]any)
		h.m["base64StdEncode"] = func(b []byte) string {
			return base64.StdEncoding.EncodeToString(b)
		}
		h.m["base64StdDecode"] = func(b string) []byte {
			return fn.Panic1(base64.StdEncoding.DecodeString(b))
		}
		h.m["base64UrlEncode"] = func(b []byte) string {
			return base64.URLEncoding.EncodeToString(b)
		}
		h.m["base64UrlDecode"] = func(b string) []byte {
			return fn.Panic1(base64.URLEncoding.DecodeString(b))
		}
		h.m["base64RawStdEncode"] = func(b []byte) string {
			return base64.RawStdEncoding.EncodeToString(b)
		}
		h.m["base64RawStdDecode"] = func(b string) []byte {
			return fn.Panic1(base64.RawStdEncoding.DecodeString(b))
		}
		h.m["base64RawUrlEncode"] = func(b []byte) string {
			return base64.RawURLEncoding.EncodeToString(b)
		}
		h.m["base64RawUrlDecode"] = func(b string) []byte {
			return fn.Panic1(base64.RawURLEncoding.DecodeString(b))
		}
		h.m["hexEncode"] = func(b []byte) string {
			return hex.EncodeToString(b)
		}
		h.m["hexDecode"] = func(b string) []byte {
			return fn.Panic1(hex.DecodeString(b))
		}
		h.m["base32StdEncode"] = func(b []byte) string {
			return base32.StdEncoding.EncodeToString(b)
		}
		h.m["base32StdDecode"] = func(b string) []byte {
			return fn.Panic1(base32.StdEncoding.DecodeString(b))
		}
		h.m["base32HexEncode"] = func(b []byte) string {
			return base32.HexEncoding.EncodeToString(b)
		}
		h.m["base32HexDecode"] = func(b string) []byte {
			return fn.Panic1(base32.HexEncoding.DecodeString(b))
		}
	}
	return h.m
}
