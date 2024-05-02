package spec

import (
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

type (
	FnIdentType     = func(t *ast.Ident)
	FnSelectorType  = func(t *ast.SelectorExpr, target types.Object)
	FnStarType      = func(t *ast.StarExpr)
	FnFuncType      = func(t *ast.FuncType)
	FnArrayType     = func(t *ast.ArrayType)
	FnMapType       = func(t *ast.MapType)
	FnStructType    = func(t *ast.StructType)
	FnChanType      = func(t *ast.ChanType)
	FnEllipsis      = func(t *ast.Ellipsis)
	FnInterfaceType = func(t *ast.InterfaceType)
)
type TypeCases interface {
	IdentType(t *ast.Ident)
	SelectorType(t *ast.SelectorExpr, target types.Object)
	StarType(t *ast.StarExpr)
	FuncType(t *ast.FuncType)
	ArrayType(t *ast.ArrayType)
	MapType(t *ast.MapType)
	StructType(t *ast.StructType)
	ChanType(t *ast.ChanType)
	Ellipsis(t *ast.Ellipsis)
	InterfaceType(t *ast.InterfaceType)
}

//region BaseTypeCases

type BaseTypeCases struct {
}

func (b BaseTypeCases) IdentType(t *ast.Ident) {

}

func (b BaseTypeCases) SelectorType(t *ast.SelectorExpr, target types.Object) {

}

func (b BaseTypeCases) StarType(t *ast.StarExpr) {

}

func (b BaseTypeCases) FuncType(t *ast.FuncType) {

}

func (b BaseTypeCases) ArrayType(t *ast.ArrayType) {

}

func (b BaseTypeCases) MapType(t *ast.MapType) {

}

func (b BaseTypeCases) StructType(t *ast.StructType) {

}

func (b BaseTypeCases) ChanType(t *ast.ChanType) {

}

// endregion
// region DepthTypeCases
//
//go:generate stringer -type=AstNode
const (
	AstNothing AstNode = iota
	AstMethodDecl
	AstMethodStarDecl
	AstMethodIdentDecl
	AstFuncDecl
	AstIdent
	AstSelectorExpr
	AstStarExpr
	AstFuncType
	AstArrayType
	AstMapType
	AstStructType
	AstInterfaceType
	AstChanType
	AstEllipsis
	AstNodeBuiltInMax
)

type (
	AstNode int
	Layer   struct {
		Node           AstNode
		Ident          *ast.Ident
		SelectorTarget types.Object
		Selector       *ast.SelectorExpr
		Star           *ast.StarExpr
		FuncType       *ast.FuncType
		ArrayType      *ast.ArrayType
		MapType        *ast.MapType
		StructType     *ast.StructType
		ChanType       *ast.ChanType
		InterfaceType  *ast.InterfaceType
		Ellipsis       *ast.Ellipsis
	}
	DepthTypeCases struct {
		Stack []Layer
	}
)

func (d *DepthTypeCases) Append(x *DepthTypeCases) {
	d.Stack = append(d.Stack, x.Stack...)
}
func (d *DepthTypeCases) ReplaceLast(l0 Layer) (l Layer, ok bool) {
	l, ok = d.Last()
	if ok {
		d.Stack[len(d.Stack)-1] = l0
	}
	return
}
func (d *DepthTypeCases) LastType() (l Layer, ok bool) {
	n := len(d.Stack)
	if n == 0 {
		return Layer{}, false
	}
	for i := n - 1; i >= 0; i-- {
		l = d.Stack[i]
		switch l.Node {
		case AstStructType:
			return l, true
		}
	}

	return
}
func (d *DepthTypeCases) Pop() (l Layer, ok bool) {
	l, ok = d.Last()
	if ok {
		d.Stack = d.Stack[:len(d.Stack)-1]
	}
	return
}
func (d *DepthTypeCases) Depth() int {
	return len(d.Stack)
}
func (d *DepthTypeCases) Last() (Layer, bool) {
	if len(d.Stack) == 0 {
		return Layer{}, false
	}
	return d.Stack[len(d.Stack)-1], true
}
func (d *DepthTypeCases) IdentType(t *ast.Ident) {
	d.Stack = append(d.Stack, Layer{Node: AstIdent, Ident: t})
}
func (d *DepthTypeCases) InterfaceType(t *ast.InterfaceType) {
	d.Stack = append(d.Stack, Layer{Node: AstInterfaceType, InterfaceType: t})
}

func (d *DepthTypeCases) SelectorType(t *ast.SelectorExpr, target types.Object) {
	d.Stack = append(d.Stack, Layer{
		Node:           AstSelectorExpr,
		Selector:       t,
		SelectorTarget: target,
	})
}

func (d *DepthTypeCases) StarType(t *ast.StarExpr) {
	d.Stack = append(d.Stack, Layer{
		Node: AstStarExpr,
		Star: t,
	})
}

func (d *DepthTypeCases) FuncType(t *ast.FuncType) {
	d.Stack = append(d.Stack, Layer{
		Node:     AstFuncType,
		FuncType: t,
	})
}

func (d *DepthTypeCases) ArrayType(t *ast.ArrayType) {
	d.Stack = append(d.Stack, Layer{
		Node:      AstArrayType,
		ArrayType: t,
	})
}

func (d *DepthTypeCases) MapType(t *ast.MapType) {
	d.Stack = append(d.Stack, Layer{
		Node:    AstMapType,
		MapType: t,
	})
}

func (d *DepthTypeCases) StructType(t *ast.StructType) {
	d.Stack = append(d.Stack, Layer{
		Node:       AstStructType,
		StructType: t,
	})
}

func (d *DepthTypeCases) ChanType(t *ast.ChanType) {
	d.Stack = append(d.Stack, Layer{
		Node:     AstChanType,
		ChanType: t,
	})
}
func (d *DepthTypeCases) Ellipsis(t *ast.Ellipsis) {
	d.Stack = append(d.Stack, Layer{
		Node:     AstEllipsis,
		Ellipsis: t,
	})
}

//endregion
//region UnionTypeCases

type UnionTypeCases []TypeCases

func (u UnionTypeCases) IdentType(t *ast.Ident) {
	for _, walker := range u {
		walker.IdentType(t)
	}
}
func (u UnionTypeCases) InterfaceType(t *ast.InterfaceType) {
	for _, walker := range u {
		walker.InterfaceType(t)
	}
}

func (u UnionTypeCases) SelectorType(t *ast.SelectorExpr, target types.Object) {
	for _, walker := range u {
		walker.SelectorType(t, target)
	}
}

func (u UnionTypeCases) StarType(t *ast.StarExpr) {
	for _, walker := range u {
		walker.StarType(t)
	}
}

func (u UnionTypeCases) FuncType(t *ast.FuncType) {
	for _, walker := range u {
		walker.FuncType(t)
	}
}

func (u UnionTypeCases) ArrayType(t *ast.ArrayType) {
	for _, walker := range u {
		walker.ArrayType(t)
	}
}

func (u UnionTypeCases) MapType(t *ast.MapType) {
	for _, walker := range u {
		walker.MapType(t)
	}
}

func (u UnionTypeCases) StructType(t *ast.StructType) {
	for _, walker := range u {
		walker.StructType(t)
	}
}

func (u UnionTypeCases) ChanType(t *ast.ChanType) {
	for _, walker := range u {
		walker.ChanType(t)
	}
}

func (u UnionTypeCases) Ellipsis(t *ast.Ellipsis) {
	for _, walker := range u {
		walker.Ellipsis(t)
	}
}

//endregion
//region UnionStopTypeCases

// UnionStopTypeCases can interrupt by panic of last cases
type UnionStopTypeCases []TypeCases

func (u UnionStopTypeCases) IdentType(t *ast.Ident) {
	defer func() {
		_ = recover()
	}()
	for _, walker := range u {
		walker.IdentType(t)
	}
}
func (u UnionStopTypeCases) InterfaceType(t *ast.InterfaceType) {
	defer func() {
		_ = recover()
	}()
	for _, walker := range u {
		walker.InterfaceType(t)
	}
}

func (u UnionStopTypeCases) SelectorType(t *ast.SelectorExpr, target types.Object) {
	defer func() {
		_ = recover()
	}()
	for _, walker := range u {
		walker.SelectorType(t, target)
	}
}

func (u UnionStopTypeCases) StarType(t *ast.StarExpr) {
	defer func() {
		_ = recover()
	}()
	for _, walker := range u {
		walker.StarType(t)
	}
}

func (u UnionStopTypeCases) FuncType(t *ast.FuncType) {
	defer func() {
		_ = recover()
	}()
	for _, walker := range u {
		walker.FuncType(t)
	}
}

func (u UnionStopTypeCases) ArrayType(t *ast.ArrayType) {
	defer func() {
		_ = recover()
	}()
	for _, walker := range u {
		walker.ArrayType(t)
	}
}

func (u UnionStopTypeCases) MapType(t *ast.MapType) {
	defer func() {
		_ = recover()
	}()
	for _, walker := range u {
		walker.MapType(t)
	}
}

func (u UnionStopTypeCases) StructType(t *ast.StructType) {
	defer func() {
		_ = recover()
	}()
	for _, walker := range u {
		walker.StructType(t)
	}
}

func (u UnionStopTypeCases) ChanType(t *ast.ChanType) {
	defer func() {
		_ = recover()
	}()
	for _, walker := range u {
		walker.ChanType(t)
	}
}

func (u UnionStopTypeCases) Ellipsis(t *ast.Ellipsis) {
	defer func() {
		_ = recover()
	}()
	for _, walker := range u {
		walker.Ellipsis(t)
	}
}

//endregion

func CaseType(pkg *packages.Package, typ ast.Expr, w TypeCases) {
	switch t := typ.(type) {
	case *ast.Ident:
		w.IdentType(t)
	case *ast.SelectorExpr:
		w.SelectorType(t, pkg.TypesInfo.Uses[t.Sel])
	case *ast.StarExpr:
		w.StarType(t)
	case *ast.FuncType:
		w.FuncType(t)
	case *ast.ArrayType:
		w.ArrayType(t)
	case *ast.MapType:
		w.MapType(t)
	case *ast.StructType:
		w.StructType(t)
	case *ast.ChanType:
		w.ChanType(t)
	case *ast.Ellipsis:
		w.Ellipsis(t)
	case *ast.InterfaceType:
		w.InterfaceType(t)
	default:
		fmt.Printf("\nmissing type %#+v\n", t)
	}
}
