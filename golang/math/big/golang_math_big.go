package big

import (
	_ "embed"
	_ "github.com/ZenLiuCN/engine/golang/fmt"
	_ "github.com/ZenLiuCN/engine/golang/math/rand"

	"github.com/ZenLiuCN/engine"
	"math/big"
)

var (
	//go:embed golang_math_big.d.ts
	MathBigDefine   []byte
	MathBigDeclared = map[string]any{
		"newRat":        big.NewRat,
		"newFloat":      big.NewFloat,
		"parseFloat":    big.ParseFloat,
		"newInt":        big.NewInt,
		"jacobi":        big.Jacobi,
		"ToNearestAway": big.ToNearestAway,
		"ToNearestEven": big.ToNearestEven,
		"ToZero":        big.ToZero,
		"AwayFromZero":  big.AwayFromZero,
		"ToNegativeInf": big.ToNegativeInf,
		"ToPositiveInf": big.ToPositiveInf,
	}
)

func init() {
	engine.RegisterModule(MathBigModule{})
}

type MathBigModule struct{}

func (S MathBigModule) Identity() string {
	return "golang/math/big"
}
func (S MathBigModule) TypeDefine() []byte {
	return MathBigDefine
}
func (S MathBigModule) Exports() map[string]any {
	return MathBigDeclared
}
