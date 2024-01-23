package engine

import (
	_ "embed"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
)

var (
	//go:embed module_encoding.d.ts
	encodingDefine []byte
)

type EncodingModule struct {
}

func (t EncodingModule) Identity() string {
	return "go/encoding"
}

func (t EncodingModule) TypeDefine() []byte {
	return encodingDefine
}

func (t EncodingModule) Exports() map[string]any {
	return encodings
}

var (
	encodings = map[string]any{
		"GBK":       &CharsetEncoding{simplifiedchinese.GBK},
		"GB2312":    &CharsetEncoding{simplifiedchinese.HZGB2312},
		"GB18030":   &CharsetEncoding{simplifiedchinese.GB18030},
		"BIG5":      &CharsetEncoding{traditionalchinese.Big5},
		"EucJP":     &CharsetEncoding{japanese.EUCJP},
		"ShiftJIS":  &CharsetEncoding{japanese.ShiftJIS},
		"ISO2022JP": &CharsetEncoding{japanese.ISO2022JP},
		"EucKR":     &CharsetEncoding{korean.EUCKR},

		"ISO8859_6":  &CharsetEncoding{charmap.ISO8859_6},
		"ISO8859_6E": &CharsetEncoding{charmap.ISO8859_6E},
		"ISO8859_6I": &CharsetEncoding{charmap.ISO8859_6I},
		"ISO8859_8E": &CharsetEncoding{charmap.ISO8859_8E},
		"ISO8859_8I": &CharsetEncoding{charmap.ISO8859_8I},

		"UTF8_BOM":     &CharsetEncoding{unicode.UTF8BOM},
		"UTF32_BOM_BE": &CharsetEncoding{utf32.UTF32(utf32.BigEndian, utf32.ExpectBOM)},
		"UTF32_BOM_LE": &CharsetEncoding{utf32.UTF32(utf32.LittleEndian, utf32.ExpectBOM)},
		"UTF32_BE":     &CharsetEncoding{utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM)},
		"UTF32_LE":     &CharsetEncoding{utf32.UTF32(utf32.LittleEndian, utf32.IgnoreBOM)},

		"toText": func(u string) []byte {
			return []byte(u)
		},
		"fromText": func(u []byte) string {
			return string(u)
		},
	}
)

type CharsetEncoding struct {
	enc encoding.Encoding
}

func (e *CharsetEncoding) Encode(u []byte) ([]byte, error) {
	return e.enc.NewEncoder().Bytes(u)
}
func (e *CharsetEncoding) EncodeText(u string) (string, error) {
	return e.enc.NewEncoder().String(u)
}
func (e *CharsetEncoding) Decode(u []byte) ([]byte, error) {
	return e.enc.NewDecoder().Bytes(u)
}
func (e *CharsetEncoding) DecodeText(u string) (string, error) {
	return e.enc.NewDecoder().String(u)
}
