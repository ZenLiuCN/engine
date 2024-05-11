// Code generated by define_gene; DO NOT EDIT.
package cmplx

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"math/cmplx"
)

var (
	//go:embed math_cmplx.d.ts
	MathCmplxDefine   []byte
	MathCmplxDeclared = map[string]any{
		"atanh": cmplx.Atanh,
		"isInf": cmplx.IsInf,
		"isNaN": cmplx.IsNaN,
		"pow":   cmplx.Pow,
		"asin":  cmplx.Asin,
		"asinh": cmplx.Asinh,
		"phase": cmplx.Phase,
		"polar": cmplx.Polar,
		"tanh":  cmplx.Tanh,
		"log10": cmplx.Log10,
		"sqrt":  cmplx.Sqrt,
		"atan":  cmplx.Atan,
		"inf":   cmplx.Inf,
		"log":   cmplx.Log,
		"tan":   cmplx.Tan,
		"acos":  cmplx.Acos,
		"cosh":  cmplx.Cosh,
		"naN":   cmplx.NaN,
		"sin":   cmplx.Sin,
		"conj":  cmplx.Conj,
		"sinh":  cmplx.Sinh,
		"cos":   cmplx.Cos,
		"rect":  cmplx.Rect,
		"abs":   cmplx.Abs,
		"acosh": cmplx.Acosh,
		"cot":   cmplx.Cot,
		"exp":   cmplx.Exp,
	}
)

func init() {
	engine.RegisterModule(MathCmplxModule{})
}

type MathCmplxModule struct{}

func (S MathCmplxModule) Identity() string {
	return "golang/math/cmplx"
}
func (S MathCmplxModule) TypeDefine() []byte {
	return MathCmplxDefine
}
func (S MathCmplxModule) Exports() map[string]any {
	return MathCmplxDeclared
}
