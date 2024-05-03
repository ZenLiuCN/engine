package main

import (
	"bytes"
	"github.com/ZenLiuCN/engine/define_gene/spec"
	"github.com/ZenLiuCN/fn"
	"go/ast"
	"go/types"
)

type Type struct {
	File       *ast.File
	Type       Kind
	Spec       *ast.TypeSpec
	Interface  *ast.InterfaceType
	Ident      *ast.Ident
	Struct     *ast.StructType
	Map        *ast.MapType
	Array      *ast.ArrayType
	Func       *ast.FuncType
	FuncDecl   *ast.FuncDecl
	IdentRecv  *ast.Ident
	Recv       *ast.Field
	MethodDecl *ast.FuncDecl
	StarRecv   *ast.StarExpr
	ChanType   *ast.ChanType
	Select     *ast.SelectorExpr
	Star       *ast.StarExpr
	Decl       *ast.GenDecl
}
type Value struct {
	File  *ast.File
	Name  string
	Ident *ast.Ident
	Spec  *ast.ValueSpec
	Decl  *ast.GenDecl
}

type Define struct {
	File *ast.File
	Type Kind
	Spec *ast.TypeSpec
	Decl *ast.GenDecl
	Prim *Type
	Elt  fn.HashSet[*Type]
	Val  fn.HashSet[*Value]
}

func (d *Define) GoString() string {
	b := new(bytes.Buffer)
	b.WriteString("Define{")
	b.WriteString(d.Type.String())
	if d.Spec != nil {
		b.WriteByte(';')
		b.WriteString(d.Spec.Name.Name)
	}
	b.WriteByte('}')
	return b.String()
}

func (d *Define) AddElt(v *Type) *Define {
	if d.Elt == nil {
		d.Elt = fn.NewHashSet[*Type]()
	}
	d.Elt.Add(v)
	return d
}
func (d *Define) AddVal(v *Value) *Define {
	if d.Val == nil {
		d.Val = fn.NewHashSet[*Value]()
	}
	d.Val.Add(v)
	return d
}

type Types []TypeInfo

func (t Types) LastN(i int) TypeInfo {
	n := len(t)
	if n >= i {
		return t[n-i]
	}
	return TypeInfo{}
}

func (t Types) Empty() bool {
	return len(t) == 0
}
func (t Types) Len() int {
	return len(t)
}
func (t Types) Index(node spec.AstNode) int {
	for i, info := range t {
		if info.Node == node {
			return i
		}
	}
	return -1
}
func (t Types) LastIndex(node spec.AstNode) int {
	n := t.Len()
	for i := n - 1; i > -1; i-- {
		if t[i].Node == node {
			return n - i
		}
	}
	return -1
}

type TypeInfo struct {
	Node           spec.AstNode
	Ident          *ast.Ident
	SelectorTarget types.Object
	Selector       *ast.SelectorExpr
	Star           *ast.StarExpr
	Func           *ast.FuncType
	Array          *ast.ArrayType
	Map            *ast.MapType
	Struct         *ast.StructType
	Chan           *ast.ChanType
	Ellipsis       *ast.Ellipsis
	Interface      *ast.InterfaceType
}

func (t TypeInfo) GoString() string {
	b := new(bytes.Buffer)
	b.WriteString("Info{")
	b.WriteString(t.Node.String())
	if t.Ident != nil {
		b.WriteByte(';')
		b.WriteString(t.Ident.Name)
	}
	if t.Selector != nil {
		b.WriteByte(';')
		if t.SelectorTarget != nil {
			b.WriteString(t.SelectorTarget.Pkg().Name())
			b.WriteByte('.')
		}
		b.WriteString(t.Selector.Sel.Name)

	}

	b.WriteByte('}')
	return b.String()
}
