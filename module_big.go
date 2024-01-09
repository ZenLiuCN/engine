package engine

import (
	_ "embed"
	"fmt"
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
	ic := eng.ToValue(eng.ToConstructor(func(v []goja.Value) (any, error) {
		switch t := v[0].Export().(type) {
		case int64:
			return big.NewInt(t), nil
		case float64:
			return big.NewInt(int64(t)), nil
		case string:
			return big.NewInt(fn.Panic1(strconv.ParseInt(t, 0, 64))), nil
		default:
			return nil, fmt.Errorf("only number accepted")
		}
	}))
	rc := eng.ToValue(eng.ToConstructor(func(v []goja.Value) (any, error) {
		switch len(v) {
		case 1:
			switch t := v[0].Export().(type) {
			case int64:
				return big.NewRat(t, 1), nil
			case float64:
				return big.NewRat(0, 1).SetFloat64(t), nil
			case string:
				r, _ := big.NewRat(0, 1).SetString(t)
				return r, nil
			default:
				return nil, fmt.Errorf("only number accepted")
			}
		case 2:
			switch a := v[0].Export().(type) {
			case int64:
				switch b := v[1].Export().(type) {
				case int64:
					return big.NewRat(a, b), nil
				default:
					return nil, fmt.Errorf("only integer with integer accepted")
				}
			case float64:
				return big.NewRat(0, 1).SetFloat64(a), nil
			default:
				return nil, fmt.Errorf("only number accepted")
			}
		default:
			return nil, fmt.Errorf("bad arguments")
		}

	}))
	return map[string]any{
		"zeroInt": func() goja.Value {
			return fn.Panic1(eng.CallConstruct(ic, 0))
		},
		"oneRat": func() goja.Value {
			return fn.Panic1(eng.CallConstruct(rc, 1))
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
