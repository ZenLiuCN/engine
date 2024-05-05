package internal

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type AstInspector struct {
	WithUnexported bool
	Visitor        NodeVisitor[*AstInspector]
	Pkg            *packages.Package
	File           *ast.File
	Object         *ast.Object
}

func NewAstInspector(withUnexported bool) *BaseInspector[*AstInspector] {
	b := new(BaseInspector[*AstInspector])
	b.Inspector = &AstInspector{WithUnexported: withUnexported}
	return b
}

func (a *AstInspector) inspect(p *packages.Package) {
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

func (a *AstInspector) initialize(conf *packages.Config) {
	if a.Visitor == nil {
		a.Visitor = new(FnNodeVisitor[*AstInspector])
	}
	conf.Mode |= packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo
}
func (a *AstInspector) isExported(s string) bool {
	if a.WithUnexported {
		return true
	}
	return exported(s)
}

func (a *AstInspector) visitObject(name string, o *ast.Object) {
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

func (a *AstInspector) visitObjects(name string, o *ast.Object) {
	mods := Mods{ModeNamedElt}
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

		if a.Visitor.VisitFieldDecl(a, ENT, name, o, x, mods, Nodes{}) {
			a.visitTypeExpr(x.Type, o, mods, Nodes{x})
		}
		a.Visitor.VisitFieldDecl(a, EXT, name, o, x, mods, Nodes{x})
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
		if a.Visitor.VisitValueDecl(a, ENT, name, o, x, mods, Nodes{}) {
			a.visitTypeExpr(x.Type, o, mods, Nodes{x})
		}
		a.Visitor.VisitValueDecl(a, EXT, name, o, x, mods, Nodes{x})
	case *ast.TypeSpec:
		if !a.isExported(x.Name.Name) {
			return
		}
		if a.Visitor.VisitTypeDecl(a, ENT, name, o, x, mods, Nodes{}) {
			a.visitTypeExpr(x.Type, o, mods, Nodes{x})
		}
		a.Visitor.VisitTypeDecl(a, EXT, name, o, x, mods, Nodes{x})
	case *ast.FuncDecl:
		if !a.isExported(x.Name.Name) {
			return
		}
		if x.Recv.NumFields() > 0 {
			mods = append(mods, ModMethod)
		} else {
			mods = append(mods, ModFunction)
		}
		if a.Visitor.VisitFuncDecl(a, ENT, name, o, x, mods, Nodes{}) {
			a.visitTypeExpr(x.Type, o, mods, Nodes{x})
		}
		a.Visitor.VisitFuncDecl(a, EXT, name, o, x, mods, Nodes{x})
	case *ast.LabeledStmt, *ast.AssignStmt, *ast.Scope, nil:
		tracef(0, "ignore %s %T", o.Name, o.Type)
	default:
		debugf("undefined %s %T", o.Name, o.Type)
	}
}

func (a *AstInspector) visitTypeExpr(expr ast.Expr, o *ast.Object, mods Mods, nodes Nodes) {
	switch x := expr.(type) {
	case *ast.Ident:
		if a.Visitor.VisitIdent(a, ENT, o, x, mods, nodes) {
		}
		a.Visitor.VisitIdent(a, EXT, o, x, mods, append(nodes, x))
	case *ast.FuncType:
		if a.Visitor.VisitFuncType(a, ENT, o, x, mods, nodes) {
			if a.Visitor.VisitFieldList(a, ENT, o, x.TypeParams, append(mods, ModTypeParam), append(nodes, x)) {
				a.visitFieldList(x.TypeParams, o, mods, append(nodes, x, x.TypeParams), ModTypeParam)
				a.Visitor.VisitFieldList(a, EXT, o, x.TypeParams, append(mods, ModTypeParam), append(nodes, x))
			}
			if a.Visitor.VisitFieldList(a, ENT, o, x.Params, append(mods, ModParam), append(nodes, x)) {
				a.visitFieldList(x.Params, o, mods, append(nodes, x, x.Params), ModParam)
				a.Visitor.VisitFieldList(a, ENT, o, x.Params, append(mods, ModParam), append(nodes, x))
			}
			if a.Visitor.VisitFieldList(a, EXT, o, x.Results, append(mods, ModResult), append(nodes, x)) {
				a.visitFieldList(x.Results, o, mods, append(nodes, x, x.Results), ModResult)
				a.Visitor.VisitFieldList(a, ENT, o, x.Results, append(mods, ModResult), append(nodes, x))
			}
		}
		a.Visitor.VisitFuncType(a, EXT, o, x, mods, append(nodes, x))
	case *ast.StructType:
		if a.Visitor.VisitStructType(a, ENT, o, x, mods, nodes) {
			a.visitFieldList(x.Fields, o, mods, append(nodes, x), ModField)
		}
		a.Visitor.VisitStructType(a, EXT, o, x, mods, nodes)
	case *ast.InterfaceType:
		if a.Visitor.VisitInterfaceType(a, ENT, o, x, mods, nodes) {
			a.visitFieldList(x.Methods, o, mods, append(nodes, x), ModMethod)
		}
		a.Visitor.VisitInterfaceType(a, EXT, o, x, mods, nodes)
	case *ast.ArrayType:
		if a.Visitor.VisitArrayType(a, ENT, o, x, mods, nodes) {
			if x.Len == nil {
				a.visitTypeExpr(x.Elt, o, append(mods, ModSliceElt), append(nodes, x))
			} else {
				a.visitTypeExpr(x.Elt, o, append(mods, ModArrayElt), append(nodes, x))
			}
		}
		a.Visitor.VisitArrayType(a, EXT, o, x, mods, nodes)
	case *ast.MapType:
		if a.Visitor.VisitMapType(a, ENT, o, x, mods, nodes) {
			a.visitTypeExpr(x.Key, o, append(mods, ModMapKey), append(nodes, x))
			a.visitTypeExpr(x.Value, o, append(mods, ModMapValue), append(nodes, x))
		}
		a.Visitor.VisitMapType(a, EXT, o, x, mods, nodes)
	case *ast.ChanType:
		if a.Visitor.VisitChanType(a, ENT, o, x, mods, nodes) {
			a.visitTypeExpr(x.Value, o, append(mods, ModChanElt), append(nodes, x))
		}
		a.Visitor.VisitChanType(a, EXT, o, x, mods, nodes)
	case *ast.StarExpr:
		if a.Visitor.VisitStarExpr(a, ENT, o, x, mods, nodes) {
			a.visitTypeExpr(x.X, o, append(mods, ModPointerElt), append(nodes, x))
		}
		a.Visitor.VisitStarExpr(a, EXT, o, x, mods, nodes)
	case *ast.SelectorExpr:
		if a.Visitor.VisitSelectorExpr(a, ENT, o, x, mods, nodes) {
			a.visitTypeExpr(x.X, o, append(mods, ModSelect), append(nodes, x))
		}
		a.Visitor.VisitSelectorExpr(a, EXT, o, x, mods, nodes)

	case *ast.BasicLit:
		if a.Visitor.VisitBasicLit(a, ENT, o, x, mods, nodes) {

		}
		a.Visitor.VisitBasicLit(a, EXT, o, x, mods, nodes)
	case *ast.Ellipsis:
		if a.Visitor.VisitEllipsis(a, ENT, o, x, mods, nodes) {
			a.visitTypeExpr(x.Elt, o, append(mods, ModEllipsis), append(nodes, x))
		}
		a.Visitor.VisitEllipsis(a, EXT, o, x, mods, nodes)
	case *ast.UnaryExpr:
		if a.Visitor.VisitUnaryExpr(a, ENT, o, x, mods, nodes) {
			a.visitTypeExpr(x.X, o, append(mods, ModOperate), append(nodes, x))
		}
		a.Visitor.VisitUnaryExpr(a, EXT, o, x, mods, nodes)
	case *ast.BinaryExpr:
		if a.Visitor.VisitBinaryExpr(a, ENT, o, x, mods, nodes) {
			a.visitTypeExpr(x.X, o, append(mods, ModOperate, ModOpLeft), append(nodes, x))
			a.visitTypeExpr(x.Y, o, append(mods, ModOperate, ModOpRight), append(nodes, x))
		}
		a.Visitor.VisitBinaryExpr(a, EXT, o, x, mods, nodes)
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

func (a *AstInspector) visitFieldList(params *ast.FieldList, o *ast.Object, mods Mods, nodes Nodes, mod Mod) {
	if params.NumFields() == 0 {
		return
	}
	lst := params.List
	n := len(lst)
	mods = append(mods, ModTuple)
	for i := 0; i < n; i++ {
		x := lst[n]
		if mod == ModField {
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
		if a.Visitor.VisitField(a, ENT, o, x, mods, nodes) {
			a.visitTypeExpr(x.Type, o, append(mods, mod), append(nodes, x))
		}
		a.Visitor.VisitField(a, EXT, o, x, mods, nodes)
	}
}
