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
		"acosh": cmplx.Acosh,
		"conj":  cmplx.Conj,
		"log10": cmplx.Log10,
		"phase": cmplx.Phase,
		"polar": cmplx.Polar,
		"cot":   cmplx.Cot,
		"asinh": cmplx.Asinh,
		"pow":   cmplx.Pow,
		"sqrt":  cmplx.Sqrt,
		"isInf": cmplx.IsInf,
		"sin":   cmplx.Sin,
		"cos":   cmplx.Cos,
		"tan":   cmplx.Tan,
		"acos":  cmplx.Acos,
		"isNaN": cmplx.IsNaN,
		"naN":   cmplx.NaN,
		"abs":   cmplx.Abs,
		"log":   cmplx.Log,
		"tanh":  cmplx.Tanh,
		"atanh": cmplx.Atanh,
		"sinh":  cmplx.Sinh,
		"asin":  cmplx.Asin,
		"cosh":  cmplx.Cosh,
		"atan":  cmplx.Atan,
		"exp":   cmplx.Exp,
		"inf":   cmplx.Inf,
		"rect":  cmplx.Rect,
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
