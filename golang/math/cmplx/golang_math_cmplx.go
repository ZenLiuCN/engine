package cmplx

import (
	_ "embed"

	"github.com/ZenLiuCN/engine"
	"math/cmplx"
)

var (
	//go:embed golang_math_cmplx.d.ts
	MathCmplxDefine   []byte
	MathCmplxDeclared = map[string]any{
		"tanh":  cmplx.Tanh,
		"atanh": cmplx.Atanh,
		"isNaN": cmplx.IsNaN,
		"log":   cmplx.Log,
		"sinh":  cmplx.Sinh,
		"sqrt":  cmplx.Sqrt,
		"pow":   cmplx.Pow,
		"cos":   cmplx.Cos,
		"abs":   cmplx.Abs,
		"asinh": cmplx.Asinh,
		"log10": cmplx.Log10,
		"polar": cmplx.Polar,
		"tan":   cmplx.Tan,
		"rect":  cmplx.Rect,
		"cot":   cmplx.Cot,
		"isInf": cmplx.IsInf,
		"sin":   cmplx.Sin,
		"asin":  cmplx.Asin,
		"exp":   cmplx.Exp,
		"phase": cmplx.Phase,
		"acos":  cmplx.Acos,
		"acosh": cmplx.Acosh,
		"conj":  cmplx.Conj,
		"inf":   cmplx.Inf,
		"atan":  cmplx.Atan,
		"naN":   cmplx.NaN,
		"cosh":  cmplx.Cosh,
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
