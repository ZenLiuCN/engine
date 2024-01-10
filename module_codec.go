package engine

import (
	_ "embed"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
)

var (
	//go:embed module_codec.d.ts
	codecDefine []byte
	codecMap    = map[string]any{
		"base64StdEncode": func(b []byte) string {
			return base64.StdEncoding.EncodeToString(b)
		},
		"base64StdDecode": func(b string) ([]byte, error) {
			return base64.StdEncoding.DecodeString(b)
		},
		"base64UrlEncode": func(b []byte) string {
			return base64.URLEncoding.EncodeToString(b)
		},
		"base64UrlDecode": func(b string) ([]byte, error) {
			return base64.URLEncoding.DecodeString(b)
		},
		"base64RawStdEncode": func(b []byte) string {
			return base64.RawStdEncoding.EncodeToString(b)
		},
		"base64RawStdDecode": func(b string) ([]byte, error) {
			return base64.RawStdEncoding.DecodeString(b)
		},
		"base64RawUrlEncode": func(b []byte) string {
			return base64.RawURLEncoding.EncodeToString(b)
		},
		"base64RawUrlDecode": func(b string) ([]byte, error) {
			return base64.RawURLEncoding.DecodeString(b)
		},
		"hexEncode": func(b []byte) string {
			return hex.EncodeToString(b)
		},
		"hexDecode": func(b string) ([]byte, error) {
			return hex.DecodeString(b)
		},
		"base32StdEncode": func(b []byte) string {
			return base32.StdEncoding.EncodeToString(b)
		},
		"base32StdDecode": func(b string) ([]byte, error) {
			return base32.StdEncoding.DecodeString(b)
		},
		"base32HexEncode": func(b []byte) string {
			return base32.HexEncoding.EncodeToString(b)
		},
		"base32HexDecode": func(b string) ([]byte, error) {
			return base32.HexEncoding.DecodeString(b)
		},
	}
)

type CodecModule struct {
}

func (h CodecModule) Identity() string {
	return "go/codec"
}

func (h CodecModule) TypeDefine() []byte {
	return codecDefine
}

func (h CodecModule) Exports() map[string]any {

	return codecMap
}
