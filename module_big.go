package engine

import (
	_ "embed"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"math/big"
	"strconv"
)

var (
	//go:embed module_big.d.ts
	bigDefine []byte
)

type BigModule struct {
}

func (c BigModule) Identity() string {
	return "go/big"
}

func (c BigModule) TypeDefine() []byte {
	return bigDefine
}

func (c BigModule) Exports() map[string]any {
	return nil
}

func (c BigModule) ExportsWithEngine(eng *Engine) map[string]any {
	ic := eng.ToValue(eng.ToConstructorRecover(func(v []goja.Value) any {
		switch t := v[0].Export().(type) {
		case int64:
			return big.NewInt(t)
		case float64:
			return big.NewInt(int64(t))
		case string:
			return big.NewInt(fn.Panic1(strconv.ParseInt(t, 0, 64)))
		default:
			panic("only number accepted")
		}
	}))
	rc := eng.ToValue(eng.ToConstructorRecover(func(v []goja.Value) any {
		switch len(v) {
		case 1:
			switch t := v[0].Export().(type) {
			case int64:
				return big.NewRat(t, 1)
			case float64:
				return big.NewRat(0, 1).SetFloat64(t)
			case string:
				r, _ := big.NewRat(0, 1).SetString(t)
				return r
			default:
				panic("only number accepted")
			}
		case 2:
			switch a := v[0].Export().(type) {
			case int64:
				switch b := v[1].Export().(type) {
				case int64:
					return big.NewRat(a, b)
				default:
					panic("only integer with integer accepted")
				}
			case float64:
				return big.NewRat(0, 1).SetFloat64(a)
			default:
				panic("only number accepted")
			}
		default:
			panic("bad arguments")
		}

	}))
	return map[string]any{
		"zeroInt": func() (goja.Value, error) {
			return eng.CallConstruct(ic, 0)
		},
		"oneRat": func() (goja.Value, error) {
			return eng.CallConstruct(rc, 1)
		},
		"equals": func(a, b *big.Int) bool {
			return a.Cmp(b) == 0
		},
		"bigger": func(a, b *big.Int) bool {
			return a.Cmp(b) == 1
		},
		"smaller": func(a, b *big.Int) bool {
			return a.Cmp(b) == -1
		},
		"Int": ic,
		"Rat": rc,
	}
}
