package bits

import (
	_ "embed"

	"github.com/ZenLiuCN/engine"
	"math/bits"
)

var (
	//go:embed golang_math_bits.d.ts
	MathBitsDefine   []byte
	MathBitsDeclared = map[string]any{
		"reverse":         bits.Reverse,
		"reverse16":       bits.Reverse16,
		"reverseBytes":    bits.ReverseBytes,
		"len8":            bits.Len8,
		"rotateLeft32":    bits.RotateLeft32,
		"reverseBytes32":  bits.ReverseBytes32,
		"rem32":           bits.Rem32,
		"onesCount16":     bits.OnesCount16,
		"rotateLeft16":    bits.RotateLeft16,
		"reverse64":       bits.Reverse64,
		"div":             bits.Div,
		"div64":           bits.Div64,
		"trailingZeros8":  bits.TrailingZeros8,
		"onesCount8":      bits.OnesCount8,
		"rotateLeft8":     bits.RotateLeft8,
		"reverseBytes16":  bits.ReverseBytes16,
		"len":             bits.Len,
		"len16":           bits.Len16,
		"len32":           bits.Len32,
		"sub":             bits.Sub,
		"leadingZeros32":  bits.LeadingZeros32,
		"reverseBytes64":  bits.ReverseBytes64,
		"rem":             bits.Rem,
		"rem64":           bits.Rem64,
		"len64":           bits.Len64,
		"onesCount":       bits.OnesCount,
		"onesCount64":     bits.OnesCount64,
		"rotateLeft":      bits.RotateLeft,
		"add64":           bits.Add64,
		"mul32":           bits.Mul32,
		"trailingZeros32": bits.TrailingZeros32,
		"rotateLeft64":    bits.RotateLeft64,
		"trailingZeros16": bits.TrailingZeros16,
		"sub32":           bits.Sub32,
		"mul":             bits.Mul,
		"leadingZeros16":  bits.LeadingZeros16,
		"leadingZeros":    bits.LeadingZeros,
		"sub64":           bits.Sub64,
		"div32":           bits.Div32,
		"trailingZeros":   bits.TrailingZeros,
		"mul64":           bits.Mul64,
		"leadingZeros8":   bits.LeadingZeros8,
		"trailingZeros64": bits.TrailingZeros64,
		"reverse8":        bits.Reverse8,
		"leadingZeros64":  bits.LeadingZeros64,
		"reverse32":       bits.Reverse32,
		"add":             bits.Add,
		"add32":           bits.Add32,
		"onesCount32":     bits.OnesCount32,
	}
)

func init() {
	engine.RegisterModule(MathBitsModule{})
}

type MathBitsModule struct{}

func (S MathBitsModule) Identity() string {
	return "golang/math/bits"
}
func (S MathBitsModule) TypeDefine() []byte {
	return MathBitsDefine
}
func (S MathBitsModule) Exports() map[string]any {
	return MathBitsDeclared
}
