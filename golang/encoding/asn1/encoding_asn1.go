// Code generated by define_gene; DO NOT EDIT.
package asn1

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"encoding/asn1"
)

var (
	//go:embed encoding_asn1.d.ts
	EncodingAsn1Define   []byte
	EncodingAsn1Declared = map[string]any{
		"TagBMPString":         asn1.TagBMPString,
		"TagBitString":         asn1.TagBitString,
		"TagNull":              asn1.TagNull,
		"TagSequence":          asn1.TagSequence,
		"ClassContextSpecific": asn1.ClassContextSpecific,
		"ClassPrivate":         asn1.ClassPrivate,
		"marshal":              asn1.Marshal,
		"TagBoolean":           asn1.TagBoolean,
		"TagEnum":              asn1.TagEnum,
		"TagUTF8String":        asn1.TagUTF8String,
		"ClassApplication":     asn1.ClassApplication,
		"ClassUniversal":       asn1.ClassUniversal,
		"NullBytes":            asn1.NullBytes,
		"TagGeneralString":     asn1.TagGeneralString,
		"TagNumericString":     asn1.TagNumericString,
		"TagSet":               asn1.TagSet,
		"TagInteger":           asn1.TagInteger,
		"TagOctetString":       asn1.TagOctetString,
		"TagUTCTime":           asn1.TagUTCTime,
		"unmarshal":            asn1.Unmarshal,
		"TagGeneralizedTime":   asn1.TagGeneralizedTime,
		"TagOID":               asn1.TagOID,
		"TagT61String":         asn1.TagT61String,
		"marshalWithParams":    asn1.MarshalWithParams,
		"NullRawValue":         asn1.NullRawValue,
		"TagIA5String":         asn1.TagIA5String,
		"TagPrintableString":   asn1.TagPrintableString,
		"unmarshalWithParams":  asn1.UnmarshalWithParams,

		"emptyBitString": func() (v asn1.BitString) {
			return v
		},
		"refBitString": func() *asn1.BitString {
			var x asn1.BitString
			return &x
		},
		"refOfBitString": func(x asn1.BitString) *asn1.BitString {
			return &x
		},
		"emptyRawValue": func() (v asn1.RawValue) {
			return v
		},
		"refRawValue": func() *asn1.RawValue {
			var x asn1.RawValue
			return &x
		},
		"refOfRawValue": func(x asn1.RawValue) *asn1.RawValue {
			return &x
		}}
)

func init() {
	engine.RegisterModule(EncodingAsn1Module{})
}

type EncodingAsn1Module struct{}

func (S EncodingAsn1Module) Identity() string {
	return "golang/encoding/asn1"
}
func (S EncodingAsn1Module) TypeDefine() []byte {
	return EncodingAsn1Define
}
func (S EncodingAsn1Module) Exports() map[string]any {
	return EncodingAsn1Declared
}
