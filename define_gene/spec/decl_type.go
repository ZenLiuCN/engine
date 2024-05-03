package spec

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
)

type TypeDeclCases interface {
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
	ChanTypeSpec(s *ast.TypeSpec, t *ast.ChanType)
	SelectorTypeSpec(s *ast.TypeSpec, t *ast.SelectorExpr)
	StarExprSpec(s *ast.TypeSpec, t *ast.StarExpr)
}

//region ExportedTypeDeclCases

type ExportedTypeDeclCases struct {
	Inner TypeDeclCases
}

func (e ExportedTypeDeclCases) StarExprSpec(s *ast.TypeSpec, t *ast.StarExpr) {
	if s.Name.IsExported() {
		e.Inner.StarExprSpec(s, t)
	}
}

func (e ExportedTypeDeclCases) SelectorTypeSpec(s *ast.TypeSpec, t *ast.SelectorExpr) {
	if ast.IsExported(s.Name.Name) {
		e.Inner.SelectorTypeSpec(s, t)
	}
}

func (e ExportedTypeDeclCases) MethodDeclStarRecv(d *ast.FuncDecl, r *ast.Field, t *ast.StarExpr) {
	if id, ok := t.X.(*ast.Ident); ok {
		if id.IsExported() {
			e.Inner.MethodDeclStarRecv(d, r, t)
		}
	} else {
		log.Printf("miss method decl of starExpr: %s %#+v \n", d.Name, t)
	}
}

func (e ExportedTypeDeclCases) MethodDeclIdentRecv(d *ast.FuncDecl, r *ast.Field, t *ast.Ident) {
	if t.IsExported() {
		e.Inner.MethodDeclIdentRecv(d, r, t)
	}

}

func (e ExportedTypeDeclCases) FuncDecl(d *ast.FuncDecl) {
	if !ast.IsExported(d.Name.Name) {
		return
	}
	e.Inner.FuncDecl(d)
}

func (e ExportedTypeDeclCases) MethodDecl(d *ast.FuncDecl, r *ast.Field) bool {
	if !ast.IsExported(d.Name.Name) {
		return false
	}
	return e.Inner.MethodDecl(d, r)
}

func (e ExportedTypeDeclCases) IdentTypeSpec(s *ast.TypeSpec, t *ast.Ident) {
	if ast.IsExported(s.Name.Name) {
		e.Inner.IdentTypeSpec(s, t)
	}
}

func (e ExportedTypeDeclCases) StructTypeSpec(s *ast.TypeSpec, t *ast.StructType) {
	if !ast.IsExported(s.Name.Name) {
		return
	}
	e.Inner.StructTypeSpec(s, t)
}

func (e ExportedTypeDeclCases) InterfaceTypeSpec(s *ast.TypeSpec, t *ast.InterfaceType) {
	if !ast.IsExported(s.Name.Name) {
		return
	}
	e.Inner.InterfaceTypeSpec(s, t)
}

func (e ExportedTypeDeclCases) MapTypeSpec(s *ast.TypeSpec, t *ast.MapType) {
	if !ast.IsExported(s.Name.Name) {
		return
	}
	e.Inner.MapTypeSpec(s, t)
}

func (e ExportedTypeDeclCases) ArrayTypeSpec(s *ast.TypeSpec, t *ast.ArrayType) {
	if !ast.IsExported(s.Name.Name) {
		return
	}
	e.Inner.ArrayTypeSpec(s, t)
}

func (e ExportedTypeDeclCases) FuncTypeSpec(s *ast.TypeSpec, t *ast.FuncType) {
	if !ast.IsExported(s.Name.Name) {
		return
	}
	e.Inner.FuncTypeSpec(s, t)
}
func (e ExportedTypeDeclCases) ChanTypeSpec(s *ast.TypeSpec, t *ast.ChanType) {
	if !ast.IsExported(s.Name.Name) {
		return
	}
	e.Inner.ChanTypeSpec(s, t)
}

//endregion

func CaseTypeDecl(file *ast.File, w TypeDeclCases) {
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
						case *ast.ChanType:
							w.ChanTypeSpec(s, t)
						case *ast.SelectorExpr:
							w.SelectorTypeSpec(s, t)
						case *ast.StarExpr:
							w.StarExprSpec(s, t)
						default:
							log.Printf("miss decl %#+v\n", t)
						}
					}
				}

			default:
				continue
			}
		}
	}
}
