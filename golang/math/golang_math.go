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
		"atan":            math.Atan,
		"gamma":           math.Gamma,
		"modf":            math.Modf,
		"atan2":           math.Atan2,
		"y0":              math.Y0,
		"nextafter32":     math.Nextafter32,
		"log2":            math.Log2,
		"log1p":           math.Log1p,
		"float32bits":     math.Float32bits,
		"asin":            math.Asin,
		"erf":             math.Erf,
		"floor":           math.Floor,
		"trunc":           math.Trunc,
		"fma":             math.FMA,
		"inf":             math.Inf,
		"isInf":           math.IsInf,
		"erfinv":          math.Erfinv,
		"j0":              math.J0,
		"log10":           math.Log10,
		"expm1":           math.Expm1,
		"yn":              math.Yn,
		"nextafter":       math.Nextafter,
		"pow10":           math.Pow10,
		"copysign":        math.Copysign,
		"exp2":            math.Exp2,
		"frexp":           math.Frexp,
		"ldexp":           math.Ldexp,
		"remainder":       math.Remainder,
		"asinh":           math.Asinh,
		"dim":             math.Dim,
		"hypot":           math.Hypot,
		"pow":             math.Pow,
		"sinh":            math.Sinh,
		"acosh":           math.Acosh,
		"max":             math.Max,
		"mod":             math.Mod,
		"j1":              math.J1,
		"float32frombits": math.Float32frombits,
		"abs":             math.Abs,
		"acos":            math.Acos,
		"atanh":           math.Atanh,
		"exp":             math.Exp,
		"roundToEven":     math.RoundToEven,
		"naN":             math.NaN,
		"min":             math.Min,
		"ceil":            math.Ceil,
		"ilogb":           math.Ilogb,
		"float64bits":     math.Float64bits,
		"round":           math.Round,
		"sqrt":            math.Sqrt,
		"float64frombits": math.Float64frombits,
		"jn":              math.Jn,
		"logb":            math.Logb,
		"sin":             math.Sin,
		"sincos":          math.Sincos,
		"tan":             math.Tan,
		"erfc":            math.Erfc,
		"erfcinv":         math.Erfcinv,
		"lgamma":          math.Lgamma,
		"signbit":         math.Signbit,
		"cos":             math.Cos,
		"isNaN":           math.IsNaN,
		"cbrt":            math.Cbrt,
		"log":             math.Log,
		"y1":              math.Y1,
		"cosh":            math.Cosh,
		"tanh":            math.Tanh,
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
