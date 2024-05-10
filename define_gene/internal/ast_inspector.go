package internal

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type AstInspector[X any] struct {
	WithUnexported bool
	Visitor        NodeVisitor[*AstInspector[X], X]
	X              X
	Pkg            *packages.Package
	File           *ast.File
	Object         *ast.Object
}

func NewAstInspector[X any](withUnexported bool) *BaseInspector[*AstInspector[X]] {
	b := new(BaseInspector[*AstInspector[X]])
	b.Inspector = &AstInspector[X]{WithUnexported: withUnexported}
	return b
}

func (a *AstInspector[X]) inspect(p *packages.Package) {
	a.Pkg = p
	for _, file := range p.Syntax {
		a.File = file
		for name, object := range file.Scope.Objects {
			if a.isExported(name) {
				a.Object = object
				a.visitObject(name, object)
			}
		}
	}
	a.Pkg = nil
}

func (a *AstInspector[X]) initialize(conf *packages.Config) {
	if a.Visitor == nil {
		a.Visitor = new(FnNodeVisitor[*AstInspector[X], X])
	}
	conf.Mode |= packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo
}
func (a *AstInspector[X]) isExported(s string) bool {
	if a.WithUnexported {
		return true
	}
	return exported(s)
}

func (a *AstInspector[X]) visitObject(name string, o *ast.Object) {
	switch o.Kind {
	case ast.Con, ast.Typ, ast.Var, ast.Fun:
		a.visitObjects(name, o)
	case ast.Pkg,
		ast.Bad,
		ast.Lbl:
		tracef(0, "ignore %s", o.Kind)
	default:
		debugf("undefined %s", o.Kind)
	}
}

func (a *AstInspector[X]) visitObjects(name string, o *ast.Object) {
	switch x := o.Type.(type) {
	case *ast.Field:
		if !a.WithUnexported {
			n := 0
			for _, name := range x.Names {
				if a.isExported(name.Name) {
					n++
				}
			}
			if n != len(x.Names) {
				return
			}
		}

		if a.Visitor.VisitFieldDecl(a, ENT, name, o, x, Nodes{}, a.X) {
			a.visitTypeExpr(x.Type, o, Nodes{x})
		}
		a.Visitor.VisitFieldDecl(a, EXT, name, o, x, Nodes{x}, a.X)
	case *ast.ValueSpec:
		if !a.WithUnexported {
			n := 0
			for _, name := range x.Names {
				if a.isExported(name.Name) {
					n++
				}
			}
			if n != len(x.Names) {
				return
			}
		}
		if a.Visitor.VisitValueDecl(a, ENT, name, o, x, Nodes{}, a.X) {
			a.visitTypeExpr(x.Type, o, Nodes{x})
		}
		a.Visitor.VisitValueDecl(a, EXT, name, o, x, Nodes{x}, a.X)
	case *ast.TypeSpec:
		if !a.isExported(x.Name.Name) {
			return
		}
		if a.Visitor.VisitTypeDecl(a, ENT, name, o, x, Nodes{}, a.X) {
			a.visitTypeExpr(x.Type, o, Nodes{x})
		}
		a.Visitor.VisitTypeDecl(a, EXT, name, o, x, Nodes{x}, a.X)
	case *ast.FuncDecl:
		if !a.isExported(x.Name.Name) {
			return
		}
		if a.Visitor.VisitFuncDecl(a, ENT, name, o, x, Nodes{}, a.X) {
			a.visitTypeExpr(x.Type, o, Nodes{x})
		}
		a.Visitor.VisitFuncDecl(a, EXT, name, o, x, Nodes{x}, a.X)
	case *ast.LabeledStmt, *ast.AssignStmt, *ast.Scope, nil:
		tracef(0, "ignore %s %T", o.Name, o.Type)
	default:
		debugf("undefined %s %T", o.Name, o.Type)
	}
}

func (a *AstInspector[X]) visitTypeExpr(expr ast.Expr, o *ast.Object, nodes Nodes) {
	switch x := expr.(type) {
	case *ast.Ident:
		if a.Visitor.VisitIdent(a, ENT, o, x, nodes, a.X) {
		}
		a.Visitor.VisitIdent(a, EXT, o, x, append(nodes, x), a.X)
	case *ast.FuncType:
		if a.Visitor.VisitFuncType(a, ENT, o, x, nodes, a.X) {
			if a.Visitor.VisitFieldList(a, ENT, o, x.TypeParams, append(nodes, x), a.X) {
				a.visitFieldList(x.TypeParams, o, append(nodes, x, x.TypeParams))
				a.Visitor.VisitFieldList(a, EXT, o, x.TypeParams, append(nodes, x), a.X)
			}
			if a.Visitor.VisitFieldList(a, ENT, o, x.Params, append(nodes, x), a.X) {
				a.visitFieldList(x.Params, o, append(nodes, x, x.Params))
				a.Visitor.VisitFieldList(a, ENT, o, x.Params, append(nodes, x), a.X)
			}
			if a.Visitor.VisitFieldList(a, EXT, o, x.Results, append(nodes, x), a.X) {
				a.visitFieldList(x.Results, o, append(nodes, x, x.Results))
				a.Visitor.VisitFieldList(a, ENT, o, x.Results, append(nodes, x), a.X)
			}
		}
		a.Visitor.VisitFuncType(a, EXT, o, x, append(nodes, x), a.X)
	case *ast.StructType:
		if a.Visitor.VisitStructType(a, ENT, o, x, nodes, a.X) {
			a.visitFieldList(x.Fields, o, append(nodes, x))
		}
		a.Visitor.VisitStructType(a, EXT, o, x, nodes, a.X)
	case *ast.InterfaceType:
		if a.Visitor.VisitInterfaceType(a, ENT, o, x, nodes, a.X) {
			a.visitFieldList(x.Methods, o, append(nodes, x))
		}
		a.Visitor.VisitInterfaceType(a, EXT, o, x, nodes, a.X)
	case *ast.ArrayType:
		if a.Visitor.VisitArrayType(a, ENT, o, x, nodes, a.X) {
			if x.Len == nil {
				a.visitTypeExpr(x.Elt, o, append(nodes, x))
			} else {
				a.visitTypeExpr(x.Elt, o, append(nodes, x))
			}
		}
		a.Visitor.VisitArrayType(a, EXT, o, x, nodes, a.X)
	case *ast.MapType:
		if a.Visitor.VisitMapType(a, ENT, o, x, nodes, a.X) {
			a.visitTypeExpr(x.Key, o, append(nodes, x))
			a.visitTypeExpr(x.Value, o, append(nodes, x))
		}
		a.Visitor.VisitMapType(a, EXT, o, x, nodes, a.X)
	case *ast.ChanType:
		if a.Visitor.VisitChanType(a, ENT, o, x, nodes, a.X) {
			a.visitTypeExpr(x.Value, o, append(nodes, x))
		}
		a.Visitor.VisitChanType(a, EXT, o, x, nodes, a.X)
	case *ast.StarExpr:
		if a.Visitor.VisitStarExpr(a, ENT, o, x, nodes, a.X) {
			a.visitTypeExpr(x.X, o, append(nodes, x))
		}
		a.Visitor.VisitStarExpr(a, EXT, o, x, nodes, a.X)
	case *ast.SelectorExpr:
		if a.Visitor.VisitSelectorExpr(a, ENT, o, x, nodes, a.X) {
			a.visitTypeExpr(x.X, o, append(nodes, x))
		}
		a.Visitor.VisitSelectorExpr(a, EXT, o, x, nodes, a.X)

	case *ast.BasicLit:
		if a.Visitor.VisitBasicLit(a, ENT, o, x, nodes, a.X) {

		}
		a.Visitor.VisitBasicLit(a, EXT, o, x, nodes, a.X)
	case *ast.Ellipsis:
		if a.Visitor.VisitEllipsis(a, ENT, o, x, nodes, a.X) {
			a.visitTypeExpr(x.Elt, o, append(nodes, x))
		}
		a.Visitor.VisitEllipsis(a, EXT, o, x, nodes, a.X)
	case *ast.UnaryExpr:
		if a.Visitor.VisitUnaryExpr(a, ENT, o, x, nodes, a.X) {
			a.visitTypeExpr(x.X, o, append(nodes, x))
		}
		a.Visitor.VisitUnaryExpr(a, EXT, o, x, nodes, a.X)
	case *ast.BinaryExpr:
		if a.Visitor.VisitBinaryExpr(a, ENT, o, x, nodes, a.X) {
			a.visitTypeExpr(x.X, o, append(nodes, x))
			a.visitTypeExpr(x.Y, o, append(nodes, x))
		}
		a.Visitor.VisitBinaryExpr(a, EXT, o, x, nodes, a.X)
	case *ast.SliceExpr,
		*ast.IndexExpr,
		*ast.CallExpr,
		*ast.TypeAssertExpr,
		*ast.KeyValueExpr,
		*ast.BadExpr,
		*ast.CompositeLit,
		*ast.FuncLit,
		*ast.ParenExpr:
		tracef(0, "ignore expr $T", x)
	default:
		debugf("undefined %T", x)
	}
}

func (a *AstInspector[X]) visitFieldList(params *ast.FieldList, o *ast.Object, nodes Nodes) {
	if params.NumFields() == 0 {
		return
	}
	lst := params.List
	n := len(lst)
	for i := 0; i < n; i++ {
		x := lst[n]
		if len(nodes) > 0 && (IsNode[*ast.StructType](nodes[len(nodes)-1])) {

			if len(x.Names) > 0 {
				c := 0
				for _, name := range x.Names {
					if a.isExported(name.Name) {
						c++
					}
				}
				if c == 0 {
					return
				}
			}
		}
		if a.Visitor.VisitField(a, ENT, o, x, nodes, a.X) {
			a.visitTypeExpr(x.Type, o, append(nodes, x))
		}
		a.Visitor.VisitField(a, EXT, o, x, nodes, a.X)
	}
}
