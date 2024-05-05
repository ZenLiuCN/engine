package internal

import "go/types"



type Types []types.Type

func (t Types) IsEmpty() bool {
	return len(t) == 0
}
func (t Types) Nth(n int) types.Type {
	if len(t) > n {
		return t[n]
	}
	return nil
}
func (t Types) LNth(n int) types.Type {
	if len(t) >= n {
		return t[len(t)-n]
	}
	return nil
}

func (t Types) NthBasic(n int) (v *types.Basic) {
	x := t.Nth(n)
	v, _ = x.(*types.Basic)
	return
}

func (t Types) NthMap(n int) (v *types.Map) {
	x := t.Nth(n)
	v, _ = x.(*types.Map)
	return
}

func (t Types) NthArray(n int) (v *types.Array) {
	x := t.Nth(n)
	v, _ = x.(*types.Array)
	return
}

func (t Types) NthStruct(n int) (v *types.Struct) {
	x := t.Nth(n)
	v, _ = x.(*types.Struct)
	return
}

func (t Types) NthTuple(n int) (v *types.Tuple) {
	x := t.Nth(n)
	v, _ = x.(*types.Tuple)
	return
}

func (t Types) NthUnion(n int) (v *types.Union) {
	x := t.Nth(n)
	v, _ = x.(*types.Union)
	return
}

func (t Types) NthSignature(n int) (v *types.Signature) {
	x := t.Nth(n)
	v, _ = x.(*types.Signature)
	return
}

func (t Types) NthTypeParam(n int) (v *types.TypeParam) {
	x := t.Nth(n)
	v, _ = x.(*types.TypeParam)
	return
}

func (t Types) NthPointer(n int) (v *types.Pointer) {
	x := t.Nth(n)
	v, _ = x.(*types.Pointer)
	return
}

func (t Types) NthSlice(n int) (v *types.Slice) {
	x := t.Nth(n)
	v, _ = x.(*types.Slice)
	return
}

func (t Types) NthInterface(n int) (v *types.Interface) {
	x := t.Nth(n)
	v, _ = x.(*types.Interface)
	return
}

func (t Types) NthChan(n int) (v *types.Chan) {
	x := t.Nth(n)
	v, _ = x.(*types.Chan)
	return
}

func (t Types) NthNamed(n int) (v *types.Named) {
	x := t.Nth(n)
	v, _ = x.(*types.Named)
	return
}

func (t Types) LNthBasic(n int) (v *types.Basic) {
	v, _ = t.LNth(n).(*types.Basic)
	return
}
func (t Types) LNthMap(n int) (v *types.Map) {
	v, _ = t.LNth(n).(*types.Map)
	return
}
func (t Types) LNthArray(n int) (v *types.Array) {
	v, _ = t.LNth(n).(*types.Array)
	return
}
func (t Types) LNthStruct(n int) (v *types.Struct) {
	v, _ = t.LNth(n).(*types.Struct)
	return
}
func (t Types) LNthTuple(n int) (v *types.Tuple) {
	v, _ = t.LNth(n).(*types.Tuple)
	return
}
func (t Types) LNthUnion(n int) (v *types.Union) {
	v, _ = t.LNth(n).(*types.Union)
	return
}
func (t Types) LNthSignature(n int) (v *types.Signature) {
	v, _ = t.LNth(n).(*types.Signature)
	return
}
func (t Types) LNthTypeParam(n int) (v *types.TypeParam) {
	v, _ = t.LNth(n).(*types.TypeParam)
	return
}
func (t Types) LNthPointer(n int) (v *types.Pointer) {
	v, _ = t.LNth(n).(*types.Pointer)
	return
}
func (t Types) LNthSlice(n int) (v *types.Slice) {
	v, _ = t.LNth(n).(*types.Slice)
	return
}
func (t Types) LNthInterface(n int) (v *types.Interface) {
	v, _ = t.LNth(n).(*types.Interface)
	return
}
func (t Types) LNthChan(n int) (v *types.Chan) {
	v, _ = t.LNth(n).(*types.Chan)
	return
}
func (t Types) LNthNamed(n int) (v *types.Named) {
	v, _ = t.LNth(n).(*types.Named)
	return
}

type (
	TypeRootVisitor[I Inspector, T types.Object]      func(i I, d Dir, name string, e T) bool
	TypeElementTypeVisitor[I Inspector, T types.Type] func(i I, d Dir, o types.Object, x T, mods Mods, seen Types) bool
	TypeElementVisitor[I Inspector, T types.Object]   func(i I, d Dir, o types.Object, x T, mods Mods, seen Types) bool
	TypeVisitor[I Inspector]                          interface {
		VisitConst(i I, d Dir, name string, e *types.Const) bool
		VisitFunc(i I, d Dir, name string, e *types.Func) bool
		VisitTypeName(i I, d Dir, name string, e *types.TypeName) bool
		VisitVar(i I, d Dir, name string, e *types.Var) bool

		VisitTypeVar(i I, d Dir, o types.Object, x *types.Var, mods Mods, seen Types) bool
		VisitTypeFunc(i I, d Dir, o types.Object, x *types.Func, mods Mods, seen Types) bool

		VisitTypeBasic(i I, d Dir, o types.Object, x *types.Basic, mods Mods, seen Types) bool
		VisitTypeMap(i I, d Dir, o types.Object, x *types.Map, mods Mods, seen Types) bool
		VisitTypeArray(i I, d Dir, o types.Object, x *types.Array, mods Mods, seen Types) bool
		VisitTypeStruct(i I, d Dir, o types.Object, x *types.Struct, mods Mods, seen Types) bool
		VisitTypeTuple(i I, d Dir, o types.Object, x *types.Tuple, mods Mods, seen Types) bool
		VisitTypeUnion(i I, d Dir, o types.Object, x *types.Union, mods Mods, seen Types) bool
		VisitTypeSignature(i I, d Dir, o types.Object, x *types.Signature, mods Mods, seen Types) bool
		VisitTypeParam(i I, d Dir, o types.Object, x *types.TypeParam, mods Mods, seen Types) bool
		VisitTypePointer(i I, d Dir, o types.Object, x *types.Pointer, mods Mods, seen Types) bool
		VisitTypeSlice(i I, d Dir, o types.Object, x *types.Slice, mods Mods, seen Types) bool
		VisitTypeInterface(i I, d Dir, o types.Object, x *types.Interface, mods Mods, seen Types) bool
		VisitTypeChan(i I, d Dir, o types.Object, x *types.Chan, mods Mods, seen Types) bool
		VisitTypeNamed(i I, d Dir, o types.Object, x *types.Named, mods Mods, seen Types) bool
	}
	FnTypeVisitor[I Inspector] struct {
		FnVisitConst    TypeRootVisitor[I, *types.Const]
		FnVisitFunc     TypeRootVisitor[I, *types.Func]
		FnVisitTypeName TypeRootVisitor[I, *types.TypeName]
		FnVisitVar      TypeRootVisitor[I, *types.Var]

		FnVisitTypeVar  TypeElementVisitor[I, *types.Var]
		FnVisitTypeFunc TypeElementVisitor[I, *types.Func]

		FnVisitTypeBasic     TypeElementTypeVisitor[I, *types.Basic]
		FnVisitTypeMap       TypeElementTypeVisitor[I, *types.Map]
		FnVisitTypeArray     TypeElementTypeVisitor[I, *types.Array]
		FnVisitTypeStruct    TypeElementTypeVisitor[I, *types.Struct]
		FnVisitTypeTuple     TypeElementTypeVisitor[I, *types.Tuple]
		FnVisitTypeUnion     TypeElementTypeVisitor[I, *types.Union]
		FnVisitTypeSignature TypeElementTypeVisitor[I, *types.Signature]
		FnVisitTypeParam     TypeElementTypeVisitor[I, *types.TypeParam]
		FnVisitTypePointer   TypeElementTypeVisitor[I, *types.Pointer]
		FnVisitTypeSlice     TypeElementTypeVisitor[I, *types.Slice]
		FnVisitTypeInterface TypeElementTypeVisitor[I, *types.Interface]
		FnVisitTypeChan      TypeElementTypeVisitor[I, *types.Chan]
		FnVisitTypeNamed     TypeElementTypeVisitor[I, *types.Named]
	}
)

func (t FnTypeVisitor[I]) VisitConst(i I, d Dir, name string, e *types.Const) bool {
	if t.FnVisitConst != nil {
		return t.FnVisitConst(i, d, name, e)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitFunc(i I, d Dir, name string, e *types.Func) bool {
	if t.FnVisitFunc != nil {
		return t.FnVisitFunc(i, d, name, e)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeName(i I, d Dir, name string, e *types.TypeName) bool {
	if t.FnVisitTypeName != nil {
		return t.FnVisitTypeName(i, d, name, e)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitVar(i I, d Dir, name string, e *types.Var) bool {
	if t.FnVisitVar != nil {
		return t.FnVisitVar(i, d, name, e)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeVar(i I, d Dir, o types.Object, x *types.Var, mods Mods, seen Types) bool {
	if t.FnVisitTypeVar != nil {
		return t.FnVisitTypeVar(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeFunc(i I, d Dir, o types.Object, x *types.Func, mods Mods, seen Types) bool {
	if t.FnVisitTypeFunc != nil {
		return t.FnVisitTypeFunc(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeBasic(i I, d Dir, o types.Object, x *types.Basic, mods Mods, seen Types) bool {
	if t.FnVisitTypeBasic != nil {
		return t.FnVisitTypeBasic(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeMap(i I, d Dir, o types.Object, x *types.Map, mods Mods, seen Types) bool {
	if t.FnVisitTypeMap != nil {
		return t.FnVisitTypeMap(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeArray(i I, d Dir, o types.Object, x *types.Array, mods Mods, seen Types) bool {
	if t.FnVisitTypeArray != nil {
		return t.FnVisitTypeArray(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeStruct(i I, d Dir, o types.Object, x *types.Struct, mods Mods, seen Types) bool {
	if t.FnVisitTypeStruct != nil {
		return t.FnVisitTypeStruct(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeTuple(i I, d Dir, o types.Object, x *types.Tuple, mods Mods, seen Types) bool {
	if t.FnVisitTypeTuple != nil {
		return t.FnVisitTypeTuple(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeUnion(i I, d Dir, o types.Object, x *types.Union, mods Mods, seen Types) bool {
	if t.FnVisitTypeUnion != nil {
		return t.FnVisitTypeUnion(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeSignature(i I, d Dir, o types.Object, x *types.Signature, mods Mods, seen Types) bool {
	if t.FnVisitTypeSignature != nil {
		return t.FnVisitTypeSignature(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeParam(i I, d Dir, o types.Object, x *types.TypeParam, mods Mods, seen Types) bool {
	if t.FnVisitTypeParam != nil {
		return t.FnVisitTypeParam(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypePointer(i I, d Dir, o types.Object, x *types.Pointer, mods Mods, seen Types) bool {
	if t.FnVisitTypePointer != nil {
		return t.FnVisitTypePointer(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeSlice(i I, d Dir, o types.Object, x *types.Slice, mods Mods, seen Types) bool {
	if t.FnVisitTypeSlice != nil {
		return t.FnVisitTypeSlice(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeInterface(i I, d Dir, o types.Object, x *types.Interface, mods Mods, seen Types) bool {
	if t.FnVisitTypeInterface != nil {
		return t.FnVisitTypeInterface(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeChan(i I, d Dir, o types.Object, x *types.Chan, mods Mods, seen Types) bool {
	if t.FnVisitTypeChan != nil {
		return t.FnVisitTypeChan(i, d, o, x, mods, seen)
	}
	return false
}

func (t FnTypeVisitor[I]) VisitTypeNamed(i I, d Dir, o types.Object, x *types.Named, mods Mods, seen Types) bool {
	if t.FnVisitTypeNamed != nil {
		return t.FnVisitTypeNamed(i, d, o, x, mods, seen)
	}
	return false
}
