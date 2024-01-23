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
		"Base64Std":    base64.StdEncoding,
		"Base64Url":    base64.URLEncoding,
		"Base64RawStd": base64.RawStdEncoding,
		"Base64RawUrl": base64.RawURLEncoding,
		"Hex":          hexEncoding(true),
		"Utf8":         utf8Encoding(true),
		"Base32Std":    base32.StdEncoding,
		"Base32Hex":    base32.HexEncoding,
	}
)

type utf8Encoding bool

func (utf8Encoding) EncodeToString(b []byte) string {
	return string(b)
}
func (utf8Encoding) DecodeString(b string) ([]byte, error) {
	return []byte(b), nil
}

type hexEncoding bool

func (hexEncoding) EncodeToString(b []byte) string {
	return hex.EncodeToString(b)
}
func (hexEncoding) DecodeString(b string) ([]byte, error) {
	return hex.DecodeString(b)
}
func (hexEncoding) Encode(dst, src []byte) int {
	return hex.Encode(dst, src)
}
func (hexEncoding) EncodedLen(n int) int {
	return hex.EncodedLen(n)
}
func (hexEncoding) DecodedLen(n int) int {
	return hex.DecodedLen(n)
}
func (hexEncoding) Decode(dst, src []byte) (int, error) {
	return hex.Decode(dst, src)
}

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
