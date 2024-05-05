package rand

import (
	_ "embed"

	"github.com/ZenLiuCN/engine"
	"math/rand"
)

var (
	//go:embed golang_math_rand.d.ts
	MathRandDefine   []byte
	MathRandDeclared = map[string]any{
		"newSource":   rand.NewSource,
		"int31":       rand.Int31,
		"int":         rand.Int,
		"intn":        rand.Intn,
		"perm":        rand.Perm,
		"shuffle":     rand.Shuffle,
		"uint64":      rand.Uint64,
		"int31n":      rand.Int31n,
		"int63n":      rand.Int63n,
		"normFloat64": rand.NormFloat64,
		"New":         rand.New,
		"float64":     rand.Float64,
		"float32":     rand.Float32,
		"expFloat64":  rand.ExpFloat64,
		"newZipf":     rand.NewZipf,
		"int63":       rand.Int63,
		"uint32":      rand.Uint32,
	}
)

func init() {
	engine.RegisterModule(MathRandModule{})
}

type MathRandModule struct{}

func (S MathRandModule) Identity() string {
	return "golang/math/rand"
}
func (S MathRandModule) TypeDefine() []byte {
	return MathRandDefine
}
func (S MathRandModule) Exports() map[string]any {
	return MathRandDeclared
}
