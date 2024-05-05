package internal

import (
	"github.com/ZenLiuCN/fn"
	"go/types"
	"testing"
)

func TestName(t *testing.T) {
	i := new(Inspector)
	i.Visitor = FnTypeVisitor{
		FnVisitConst: func(dir Dir, name string, e *types.Const) {
			t.Logf("%s export const %s :%s= %s", dir, name, e.Type(), e.Val().String())
		},
		FnVisitFunc: func(dir Dir, name string, e *types.Func) {
			t.Logf("%s export function %s :%s", dir, name, e.String())
		},
		FnVisitTypeName: func(dir Dir, name string, e *types.TypeName) {
			t.Logf("%s type %s: %s", dir, name, e.String())
		},
		FnVisitVar: func(dir Dir, name string, e *types.Var) {
			t.Logf("%s export const Var%s:%s", dir, name, e.Type())
		},
		FnVisitTypeVar: func(dir Dir, o types.Object, x *types.Var, mods Mods, seen ...types.Type) {
			t.Logf("%s visit: %s ==> %s", dir, mods, x.Name())
		},
		FnVisitTypeFunc: func(dir Dir, o types.Object, x *types.Func, mods Mods, seen ...types.Type) {
			t.Logf("%s visit method: %s ==> %s", dir, mods, x.Name())
		},
		FnVisitTypeBasic:     nil,
		FnVisitTypeMap:       nil,
		FnVisitTypeArray:     nil,
		FnVisitTypeStruct:    nil,
		FnVisitTypeTuple:     nil,
		FnVisitTypeUnion:     nil,
		FnVisitTypeSignature: nil,
		FnVisitTypeParam:     nil,
		FnVisitTypePointer:   nil,
		FnVisitTypeSlice:     nil,
		FnVisitTypeInterface: nil,
		FnVisitTypeChan:      nil,
		FnVisitTypeNamed:     nil,
	}

	fn.Panic(i.Inspect(nil, nil, "D:\\Dev\\env\\golang\\go.1.21x64\\src\\os"))

}
