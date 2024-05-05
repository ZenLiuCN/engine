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
		"trailingZeros":   bits.TrailingZeros,
		"len64":           bits.Len64,
		"div":             bits.Div,
		"leadingZeros":    bits.LeadingZeros,
		"onesCount":       bits.OnesCount,
		"onesCount64":     bits.OnesCount64,
		"reverse":         bits.Reverse,
		"reverseBytes64":  bits.ReverseBytes64,
		"rem32":           bits.Rem32,
		"rem64":           bits.Rem64,
		"leadingZeros64":  bits.LeadingZeros64,
		"onesCount32":     bits.OnesCount32,
		"len16":           bits.Len16,
		"mul64":           bits.Mul64,
		"trailingZeros32": bits.TrailingZeros32,
		"rotateLeft64":    bits.RotateLeft64,
		"rem":             bits.Rem,
		"trailingZeros8":  bits.TrailingZeros8,
		"rotateLeft8":     bits.RotateLeft8,
		"sub64":           bits.Sub64,
		"reverse64":       bits.Reverse64,
		"reverseBytes16":  bits.ReverseBytes16,
		"sub32":           bits.Sub32,
		"leadingZeros8":   bits.LeadingZeros8,
		"trailingZeros64": bits.TrailingZeros64,
		"reverse8":        bits.Reverse8,
		"reverseBytes32":  bits.ReverseBytes32,
		"sub":             bits.Sub,
		"leadingZeros16":  bits.LeadingZeros16,
		"trailingZeros16": bits.TrailingZeros16,
		"rotateLeft16":    bits.RotateLeft16,
		"reverse16":       bits.Reverse16,
		"reverse32":       bits.Reverse32,
		"len32":           bits.Len32,
		"div64":           bits.Div64,
		"add":             bits.Add,
		"onesCount16":     bits.OnesCount16,
		"len":             bits.Len,
		"add32":           bits.Add32,
		"mul32":           bits.Mul32,
		"onesCount8":      bits.OnesCount8,
		"rotateLeft":      bits.RotateLeft,
		"add64":           bits.Add64,
		"div32":           bits.Div32,
		"leadingZeros32":  bits.LeadingZeros32,
		"reverseBytes":    bits.ReverseBytes,
		"len8":            bits.Len8,
		"mul":             bits.Mul,
		"rotateLeft32":    bits.RotateLeft32,
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
