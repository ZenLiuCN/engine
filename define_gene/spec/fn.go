package spec

import (
	"fmt"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"log"
)

type FnProvider = func(node AstNode) any
type Providers map[AstNode]any

func (p Providers) Provide(node AstNode) any {
	if fn, ok := p[node]; ok {
		return fn
	}
	return nil
}

func CaseConstDeclFunc(file *ast.File, w FnProvider) {
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			switch d.Tok {
			case token.CONST:
				for _, spec := range d.Specs {
					if s, ok := spec.(*ast.ValueSpec); ok {
						switch t := s.Type.(type) {
						case *ast.Ident:
							if fn, ok := w(AstIdent).(FnIdentConstSpec); ok {
								fn(d, s, t, ok)
							}
						case nil:
							if fn, ok := w(AstNothing).(FnNilConstSpec); ok {
								fn(d, s, ok)
							}
						default:
							fmt.Printf("miss const decl %#+v\n", t)
						}
					}
				}
			default:
				continue
			}
		}
	}
}

func CaseVarDeclFunc(file *ast.File, w FnProvider) {
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			switch d.Tok {
			case token.VAR:
				for _, spec := range d.Specs {
					if s, ok := spec.(*ast.ValueSpec); ok {
						switch t := s.Type.(type) {
						case *ast.Ident:
							if fn, ok := w(AstIdent).(FnIdentVarSpec); ok {
								fn(d, s, t)
							}
						case *ast.FuncType:
							if fn, ok := w(AstFuncType).(FnFuncVarSpec); ok {
								fn(d, s, t)
							}
						case *ast.SelectorExpr:
							if fn, ok := w(AstSelectorExpr).(FnSelectorVarSpec); ok {
								fn(d, s, t)
							}
						case *ast.MapType:
							if fn, ok := w(AstMapType).(FnMapVarSpec); ok {
								fn(d, s, t)
							}
						case *ast.StarExpr:
							if fn, ok := w(AstStarExpr).(FnStarVarSpec); ok {
								fn(d, s, t)
							}
						case nil:
							if fn, ok := w(AstNothing).(FnNilVarSpec); ok {
								fn(d, s)
							}
						default:
							fmt.Printf("miss var decl %#+v\n", t)
						}
					}
				}
			default:
				continue
			}
		}
	}
}
func CaseTypeDeclFunc(file *ast.File, w FnProvider) {
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Recv == nil {
				if fn, ok := w(AstFuncDecl).(FnFuncDecl); ok {
					fn(d)
				}
			} else if d.Recv.NumFields() != 1 {
				panic(fmt.Errorf("more than one reciver: %#+v", d))
			} else {
				r := d.Recv.List[0]
				if fn, ok := w(AstMethodDecl).(FnMethodDecl); ok {
					if fn(d, r) {
						switch t := r.Type.(type) {
						case *ast.StarExpr:
							if fn, ok := w(AstMethodStarDecl).(FnMethodDeclStarRecv); ok {
								fn(d, r, t)
							}
						case *ast.Ident:
							if fn, ok := w(AstMethodIdentDecl).(FnMethodDeclIdentRecv); ok {
								fn(d, r, t)
							}
						default:
							log.Printf("miss method reciver type: %#+v \n", t)
						}
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
							if fn, ok := w(AstIdent).(FnIdentTypeSpec); ok {
								fn(s, t)
							}
						case *ast.StructType:
							if fn, ok := w(AstStructType).(FnStructTypeSpec); ok {
								fn(s, t)
							}

						case *ast.InterfaceType:
							if fn, ok := w(AstInterfaceType).(FnInterfaceTypeSpec); ok {
								fn(s, t)
							}

						case *ast.MapType:
							if fn, ok := w(AstMapType).(FnMapTypeSpec); ok {
								fn(s, t)
							}

						case *ast.FuncType:
							if fn, ok := w(AstFuncType).(FnFuncTypeSpec); ok {
								fn(s, t)
							}

						case *ast.ArrayType:
							if fn, ok := w(AstArrayType).(FnArrayTypeSpec); ok {
								fn(s, t)
							}

						case *ast.ChanType:
							if fn, ok := w(AstChanType).(FnChanTypeSpec); ok {
								fn(s, t)
							}

						default:
							fmt.Printf("miss decl %#+v\n", t)
						}
					}
				}

			default:
				continue
			}
		}
	}
}
func CaseTypeFunc(pkg *packages.Package, typ ast.Expr, w FnProvider) {
	switch t := typ.(type) {
	case *ast.Ident:
		if fn, ok := w(AstIdent).(FnIdentType); ok {
			fn(t)
		}
	case *ast.SelectorExpr:
		if fn, ok := w(AstSelectorExpr).(FnSelectorType); ok {
			fn(t, pkg.TypesInfo.Uses[t.Sel])
		}
	case *ast.StarExpr:
		if fn, ok := w(AstStarExpr).(FnStarType); ok {
			fn(t)
		}
	case *ast.FuncType:
		if fn, ok := w(AstFuncType).(FnFuncType); ok {
			fn(t)
		}
	case *ast.ArrayType:
		if fn, ok := w(AstArrayType).(FnArrayType); ok {
			fn(t)
		}
	case *ast.MapType:
		if fn, ok := w(AstMapType).(FnMapType); ok {
			fn(t)
		}
	case *ast.StructType:
		if fn, ok := w(AstStructType).(FnStructType); ok {
			fn(t)
		}
	case *ast.ChanType:
		if fn, ok := w(AstChanType).(FnChanType); ok {
			fn(t)
		}
	case *ast.Ellipsis:
		if fn, ok := w(AstEllipsis).(FnEllipsis); ok {
			fn(t)
		}
	case *ast.InterfaceType:
		if fn, ok := w(AstInterfaceType).(FnInterfaceType); ok {
			fn(t)
		}
	default:
		fmt.Printf("\nmissing type %#+v\n", t)
	}
}
