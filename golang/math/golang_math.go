package math

import (
	_ "embed"

	"github.com/ZenLiuCN/engine"
	"math"
)

var (
	//go:embed golang_math.d.ts
	MathDefine   []byte
	MathDeclared = map[string]any{
		"tan":             math.Tan,
		"tanh":            math.Tanh,
		"lgamma":          math.Lgamma,
		"dim":             math.Dim,
		"roundToEven":     math.RoundToEven,
		"jn":              math.Jn,
		"log2":            math.Log2,
		"atanh":           math.Atanh,
		"y0":              math.Y0,
		"naN":             math.NaN,
		"nextafter32":     math.Nextafter32,
		"cos":             math.Cos,
		"sincos":          math.Sincos,
		"sqrt":            math.Sqrt,
		"floor":           math.Floor,
		"cbrt":            math.Cbrt,
		"frexp":           math.Frexp,
		"isInf":           math.IsInf,
		"erfc":            math.Erfc,
		"round":           math.Round,
		"j0":              math.J0,
		"asin":            math.Asin,
		"log10":           math.Log10,
		"modf":            math.Modf,
		"pow":             math.Pow,
		"sin":             math.Sin,
		"cosh":            math.Cosh,
		"float32bits":     math.Float32bits,
		"acosh":           math.Acosh,
		"erf":             math.Erf,
		"mod":             math.Mod,
		"nextafter":       math.Nextafter,
		"float64bits":     math.Float64bits,
		"float64frombits": math.Float64frombits,
		"max":             math.Max,
		"isNaN":           math.IsNaN,
		"ceil":            math.Ceil,
		"yn":              math.Yn,
		"ilogb":           math.Ilogb,
		"float32frombits": math.Float32frombits,
		"acos":            math.Acos,
		"trunc":           math.Trunc,
		"j1":              math.J1,
		"log":             math.Log,
		"pow10":           math.Pow10,
		"signbit":         math.Signbit,
		"exp2":            math.Exp2,
		"y1":              math.Y1,
		"copysign":        math.Copysign,
		"min":             math.Min,
		"expm1":           math.Expm1,
		"ldexp":           math.Ldexp,
		"log1p":           math.Log1p,
		"sinh":            math.Sinh,
		"asinh":           math.Asinh,
		"atan":            math.Atan,
		"inf":             math.Inf,
		"erfinv":          math.Erfinv,
		"erfcinv":         math.Erfcinv,
		"exp":             math.Exp,
		"fma":             math.FMA,
		"remainder":       math.Remainder,
		"abs":             math.Abs,
		"atan2":           math.Atan2,
		"hypot":           math.Hypot,
		"logb":            math.Logb,
		"gamma":           math.Gamma,
	}
)

func init() {
	engine.RegisterModule(MathModule{})
}

type MathModule struct{}

func (S MathModule) Identity() string {
	return "golang/math"
}
func (S MathModule) TypeDefine() []byte {
	return MathDefine
}
func (S MathModule) Exports() map[string]any {
	return MathDeclared
}
