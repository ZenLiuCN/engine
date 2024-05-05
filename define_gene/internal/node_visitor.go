package internal

import "go/ast"

type (
	Nodes []ast.Node

	NodeDeclVisit[I Inspector, T ast.Node] func(i I, d Dir, name string, o *ast.Object, x T, mods Mods, seen Nodes) bool
	NodeVisit[I Inspector, T ast.Node]     func(i I, d Dir, o *ast.Object, x T, mods Mods, seen Nodes) bool
	NodeExprVisit[I Inspector, T ast.Expr] func(i I, d Dir, o *ast.Object, x T, mods Mods, seen Nodes) bool
	NodeVisitor[I Inspector]               interface {
		VisitFieldDecl(i I, d Dir, name string, o *ast.Object, x *ast.Field, mods Mods, seen Nodes) bool
		VisitFuncDecl(i I, d Dir, name string, o *ast.Object, x *ast.FuncDecl, mods Mods, nodes Nodes) bool
		VisitTypeDecl(i I, d Dir, name string, o *ast.Object, x *ast.TypeSpec, mods Mods, nodes Nodes) bool
		VisitValueDecl(i I, d Dir, name string, o *ast.Object, x *ast.ValueSpec, mods Mods, nodes Nodes) bool

		VisitFieldList(i I, d Dir, o *ast.Object, x *ast.FieldList, mods Mods, nodes Nodes) bool
		VisitField(i I, d Dir, o *ast.Object, x *ast.Field, mods Mods, seen Nodes) bool

		VisitIdent(i I, d Dir, o *ast.Object, x *ast.Ident, mods Mods, nodes Nodes) bool
		VisitFuncType(i I, d Dir, o *ast.Object, x *ast.FuncType, mods Mods, nodes Nodes) bool
		VisitStructType(i I, d Dir, o *ast.Object, x *ast.StructType, mods Mods, nodes Nodes) bool
		VisitInterfaceType(i I, d Dir, o *ast.Object, x *ast.InterfaceType, mods Mods, nodes Nodes) bool
		VisitArrayType(i I, d Dir, o *ast.Object, x *ast.ArrayType, mods Mods, nodes Nodes) bool
		VisitMapType(i I, d Dir, o *ast.Object, x *ast.MapType, mods Mods, nodes Nodes) bool
		VisitChanType(i I, d Dir, o *ast.Object, x *ast.ChanType, mods Mods, nodes Nodes) bool
		VisitStarExpr(i I, d Dir, o *ast.Object, x *ast.StarExpr, mods Mods, nodes Nodes) bool
		VisitSelectorExpr(i I, d Dir, o *ast.Object, x *ast.SelectorExpr, mods Mods, nodes Nodes) bool
		VisitBasicLit(i I, d Dir, o *ast.Object, x *ast.BasicLit, mods Mods, nodes Nodes) bool
		VisitEllipsis(i I, d Dir, o *ast.Object, x *ast.Ellipsis, mods Mods, nodes Nodes) bool
		VisitUnaryExpr(i I, d Dir, o *ast.Object, x *ast.UnaryExpr, mods Mods, nodes Nodes) bool
		VisitBinaryExpr(i I, d Dir, o *ast.Object, x *ast.BinaryExpr, mods Mods, nodes Nodes) bool
	}
	FnNodeVisitor[I Inspector] struct {
		FnVisitFieldDecl NodeDeclVisit[I, *ast.Field]
		FnVisitFuncDecl  NodeDeclVisit[I, *ast.FuncDecl]
		FnVisitTypeDecl  NodeDeclVisit[I, *ast.TypeSpec]
		FnVisitValueDecl NodeDeclVisit[I, *ast.ValueSpec]

		FnVisitFieldList NodeVisit[I, *ast.FieldList]
		FnVisitField     NodeVisit[I, *ast.Field]

		FnVisitIdent         NodeExprVisit[I, *ast.Ident]
		FnVisitFuncType      NodeExprVisit[I, *ast.FuncType]
		FnVisitStructType    NodeExprVisit[I, *ast.StructType]
		FnVisitInterfaceType NodeExprVisit[I, *ast.InterfaceType]
		FnVisitArrayType     NodeExprVisit[I, *ast.ArrayType]
		FnVisitMapType       NodeExprVisit[I, *ast.MapType]
		FnVisitChanType      NodeExprVisit[I, *ast.ChanType]
		FnVisitStarExpr      NodeExprVisit[I, *ast.StarExpr]
		FnVisitSelectorExpr  NodeExprVisit[I, *ast.SelectorExpr]
		FnVisitBasicLit      NodeExprVisit[I, *ast.BasicLit]
		FnVisitEllipsis      NodeExprVisit[I, *ast.Ellipsis]
		FnVisitUnaryExpr     NodeExprVisit[I, *ast.UnaryExpr]
		FnVisitBinaryExpr    NodeExprVisit[I, *ast.BinaryExpr]
	}
)

func (n FnNodeVisitor[I]) VisitFieldDecl(i I, d Dir, name string, o *ast.Object, x *ast.Field, mods Mods, nodes Nodes) bool {
	if n.FnVisitFieldDecl != nil {
		return n.FnVisitFieldDecl(i, d, name, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitFuncDecl(i I, d Dir, name string, o *ast.Object, x *ast.FuncDecl, mods Mods, nodes Nodes) bool {
	if n.FnVisitFuncDecl != nil {
		return n.FnVisitFuncDecl(i, d, name, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitTypeDecl(i I, d Dir, name string, o *ast.Object, x *ast.TypeSpec, mods Mods, nodes Nodes) bool {
	if n.FnVisitTypeDecl != nil {
		return n.FnVisitTypeDecl(i, d, name, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitValueDecl(i I, d Dir, name string, o *ast.Object, x *ast.ValueSpec, mods Mods, nodes Nodes) bool {
	if n.FnVisitValueDecl != nil {
		return n.FnVisitValueDecl(i, d, name, o, x, mods, nodes)
	}
	return true
}

func (n FnNodeVisitor[I]) VisitFieldList(i I, d Dir, o *ast.Object, x *ast.FieldList, mods Mods, nodes Nodes) bool {
	if n.FnVisitFieldList != nil {
		return n.FnVisitFieldList(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitField(i I, d Dir, o *ast.Object, x *ast.Field, mods Mods, nodes Nodes) bool {
	if n.FnVisitField != nil {
		return n.FnVisitField(i, d, o, x, mods, nodes)
	}
	return true
}

func (n FnNodeVisitor[I]) VisitIdent(i I, d Dir, o *ast.Object, x *ast.Ident, mods Mods, nodes Nodes) bool {
	if n.FnVisitIdent != nil {
		return n.FnVisitIdent(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitFuncType(i I, d Dir, o *ast.Object, x *ast.FuncType, mods Mods, nodes Nodes) bool {
	if n.FnVisitFuncType != nil {
		return n.FnVisitFuncType(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitStructType(i I, d Dir, o *ast.Object, x *ast.StructType, mods Mods, nodes Nodes) bool {
	if n.FnVisitStructType != nil {
		return n.FnVisitStructType(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitInterfaceType(i I, d Dir, o *ast.Object, x *ast.InterfaceType, mods Mods, nodes Nodes) bool {
	if n.FnVisitInterfaceType != nil {
		return n.FnVisitInterfaceType(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitArrayType(i I, d Dir, o *ast.Object, x *ast.ArrayType, mods Mods, nodes Nodes) bool {
	if n.FnVisitArrayType != nil {
		return n.FnVisitArrayType(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitMapType(i I, d Dir, o *ast.Object, x *ast.MapType, mods Mods, nodes Nodes) bool {
	if n.FnVisitMapType != nil {
		return n.FnVisitMapType(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitChanType(i I, d Dir, o *ast.Object, x *ast.ChanType, mods Mods, nodes Nodes) bool {
	if n.FnVisitChanType != nil {
		return n.FnVisitChanType(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitStarExpr(i I, d Dir, o *ast.Object, x *ast.StarExpr, mods Mods, nodes Nodes) bool {
	if n.FnVisitStarExpr != nil {
		return n.FnVisitStarExpr(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitSelectorExpr(i I, d Dir, o *ast.Object, x *ast.SelectorExpr, mods Mods, nodes Nodes) bool {
	if n.FnVisitSelectorExpr != nil {
		return n.FnVisitSelectorExpr(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitBasicLit(i I, d Dir, o *ast.Object, x *ast.BasicLit, mods Mods, nodes Nodes) bool {
	if n.FnVisitBasicLit != nil {
		return n.FnVisitBasicLit(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitEllipsis(i I, d Dir, o *ast.Object, x *ast.Ellipsis, mods Mods, nodes Nodes) bool {
	if n.FnVisitEllipsis != nil {
		return n.FnVisitEllipsis(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitUnaryExpr(i I, d Dir, o *ast.Object, x *ast.UnaryExpr, mods Mods, nodes Nodes) bool {
	if n.FnVisitUnaryExpr != nil {
		return n.FnVisitUnaryExpr(i, d, o, x, mods, nodes)
	}
	return true
}
func (n FnNodeVisitor[I]) VisitBinaryExpr(i I, d Dir, o *ast.Object, x *ast.BinaryExpr, mods Mods, nodes Nodes) bool {
	if n.FnVisitBinaryExpr != nil {
		return n.FnVisitBinaryExpr(i, d, o, x, mods, nodes)
	}
	return true
}
