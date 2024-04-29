package main

import (
	"fmt"
	"github.com/ZenLiuCN/fn"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
)

/*

^// Code generated .* DO NOT EDIT\.$

*/

func IsDir(name string) bool {
	if info, err := os.Stat(name); err != nil {
		log.Fatal(err)
	} else {
		return info.IsDir()
	}
	return false
}

// ParseTypeInfo files to packages with tags
func ParseTypeInfo(flags, files []string, log func(format string, args ...any)) []*packages.Package {
	return fn.Panic1(packages.Load(&packages.Config{
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		BuildFlags: flags,
		Tests:      false,
		Logf:       log,
	}, files...))
}

type DeclCases interface {
	FuncDecl(d *ast.FuncDecl)
	MethodDecl(d *ast.FuncDecl, r *ast.Field) bool //return ture to walk for receivers
	IdentTypeSpec(s *ast.TypeSpec, t *ast.Ident)
	StructTypeSpec(s *ast.TypeSpec, t *ast.StructType)
	InterfaceTypeSpec(s *ast.TypeSpec, t *ast.InterfaceType)
	MapTypeSpec(s *ast.TypeSpec, t *ast.MapType)
	ArrayTypeSpec(s *ast.TypeSpec, t *ast.ArrayType)
	FuncTypeSpec(s *ast.TypeSpec, t *ast.FuncType)
	MethodDeclStarRecv(d *ast.FuncDecl, r *ast.Field, t *ast.StarExpr)
	MethodDeclIdentRecv(d *ast.FuncDecl, r *ast.Field, t *ast.Ident)
}

//region ExportedDeclCases

type ExportedDeclCases struct {
	Inner DeclCases
}

func (e ExportedDeclCases) MethodDeclStarRecv(d *ast.FuncDecl, r *ast.Field, t *ast.StarExpr) {
	if t.X.(*ast.Ident).IsExported() {
		e.Inner.MethodDeclStarRecv(d, r, t)
	}
}

func (e ExportedDeclCases) MethodDeclIdentRecv(d *ast.FuncDecl, r *ast.Field, t *ast.Ident) {
	if t.IsExported() {
		e.Inner.MethodDeclIdentRecv(d, r, t)
	}

}

func (e ExportedDeclCases) FuncDecl(d *ast.FuncDecl) {
	if !ast.IsExported(d.Name.Name) {
		return
	}
	e.Inner.FuncDecl(d)
}

func (e ExportedDeclCases) MethodDecl(d *ast.FuncDecl, r *ast.Field) bool {
	if !ast.IsExported(d.Name.Name) {
		return false
	}
	return e.Inner.MethodDecl(d, r)
}

func (e ExportedDeclCases) IdentTypeSpec(s *ast.TypeSpec, t *ast.Ident) {
	if ast.IsExported(s.Name.Name) {
		e.Inner.IdentTypeSpec(s, t)
	}
}

func (e ExportedDeclCases) StructTypeSpec(s *ast.TypeSpec, t *ast.StructType) {
	if !ast.IsExported(s.Name.Name) {
		return
	}
	e.Inner.StructTypeSpec(s, t)
}

func (e ExportedDeclCases) InterfaceTypeSpec(s *ast.TypeSpec, t *ast.InterfaceType) {
	if !ast.IsExported(s.Name.Name) {
		return
	}
	e.Inner.InterfaceTypeSpec(s, t)
}

func (e ExportedDeclCases) MapTypeSpec(s *ast.TypeSpec, t *ast.MapType) {
	if !ast.IsExported(s.Name.Name) {
		return
	}
	e.Inner.MapTypeSpec(s, t)
}

func (e ExportedDeclCases) ArrayTypeSpec(s *ast.TypeSpec, t *ast.ArrayType) {
	if !ast.IsExported(s.Name.Name) {
		return
	}
	e.Inner.ArrayTypeSpec(s, t)
}

func (e ExportedDeclCases) FuncTypeSpec(s *ast.TypeSpec, t *ast.FuncType) {
	if !ast.IsExported(s.Name.Name) {
		return
	}
	e.Inner.FuncTypeSpec(s, t)
}

//endregion

//region BaseDeclCases

type BaseDeclCases struct {
}

func (b BaseDeclCases) FuncDecl(decl *ast.FuncDecl) {

}

func (b BaseDeclCases) MethodDecl(d *ast.FuncDecl, field *ast.Field) {

}

func (b BaseDeclCases) IdentTypeSpec(s *ast.TypeSpec, t *ast.Ident) {

}

func (b BaseDeclCases) StructTypeSpec(s *ast.TypeSpec, t *ast.StructType) {

}

func (b BaseDeclCases) InterfaceTypeSpec(s *ast.TypeSpec, t *ast.InterfaceType) {

}

func (b BaseDeclCases) MapTypeSpec(s *ast.TypeSpec, t *ast.MapType) {

}

func (b BaseDeclCases) ArrayTypeSpec(s *ast.TypeSpec, t *ast.ArrayType) {

}

func (b BaseDeclCases) FuncTypeSpec(s *ast.TypeSpec, t *ast.FuncType) {

}

//endregion

func CaseDecl(file *ast.File, w DeclCases) {
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Recv == nil {
				w.FuncDecl(d)
			} else if d.Recv.NumFields() != 1 {
				panic(fmt.Errorf("more than one reciver: %#+v", d))
			} else {
				r := d.Recv.List[0]
				if w.MethodDecl(d, r) {
					switch t := r.Type.(type) {
					case *ast.StarExpr:
						w.MethodDeclStarRecv(d, r, t)
					case *ast.Ident:
						w.MethodDeclIdentRecv(d, r, t)
					default:
						log.Printf("miss method reciver type: %#+v \n", t)
					}
				}

			}
		case *ast.GenDecl:
			switch d.Tok {
			case token.TYPE:
				for _, spec := range d.Specs {
					if s, ok := spec.(*ast.TypeSpec); ok {
						switch t := s.Type.(type) {
						case *ast.Ident:
							w.IdentTypeSpec(s, t)
						case *ast.StructType:
							w.StructTypeSpec(s, t)
						case *ast.InterfaceType:
							w.InterfaceTypeSpec(s, t)
						case *ast.MapType:
							w.MapTypeSpec(s, t)
						case *ast.FuncType:
							w.FuncTypeSpec(s, t)
						case *ast.ArrayType:
							w.ArrayTypeSpec(s, t)
						default:
							fmt.Printf("miss %#+v\n", t)
						}
					}
				}
			default:
				continue
			}
		}
	}
}

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
	AstIdent
	AstSelectorExpr
	AstStarExpr
	AstFuncType
	AstArrayType
	AstMapType
	AstStructType
	AstChanType
	AstEllipsis
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
		Ellipsis       *ast.Ellipsis
	}
	DepthTypeCases struct {
		Stack []Layer
	}
)

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
	default:
		fmt.Printf("\nmissing type %#+v\n", t)
	}
}
