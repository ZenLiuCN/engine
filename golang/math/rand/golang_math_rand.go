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
		"int":         rand.Int,
		"float64":     rand.Float64,
		"perm":        rand.Perm,
		"read":        rand.Read,
		"uint64":      rand.Uint64,
		"int31":       rand.Int31,
		"shuffle":     rand.Shuffle,
		"seed":        rand.Seed,
		"int63":       rand.Int63,
		"int63n":      rand.Int63n,
		"int31n":      rand.Int31n,
		"intn":        rand.Intn,
		"New":         rand.New,
		"float32":     rand.Float32,
		"normFloat64": rand.NormFloat64,
		"expFloat64":  rand.ExpFloat64,
		"newZipf":     rand.NewZipf,
		"newSource":   rand.NewSource,
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
