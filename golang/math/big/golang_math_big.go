package big

import (
	_ "embed"

	"github.com/ZenLiuCN/engine"
	_ "github.com/ZenLiuCN/engine/golang/fmt"
	_ "github.com/ZenLiuCN/engine/golang/math/rand"
	"math/big"
)

var (
	//go:embed golang_math_big.d.ts
	MathBigDefine   []byte
	MathBigDeclared = map[string]any{
		"jacobi":        big.Jacobi,
		"newInt":        big.NewInt,
		"parseFloat":    big.ParseFloat,
		"ToNearestAway": big.ToNearestAway,
		"ToNearestEven": big.ToNearestEven,
		"ToZero":        big.ToZero,
		"AwayFromZero":  big.AwayFromZero,
		"ToNegativeInf": big.ToNegativeInf,
		"ToPositiveInf": big.ToPositiveInf,
		"newFloat":      big.NewFloat,
		"newRat":        big.NewRat,
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
