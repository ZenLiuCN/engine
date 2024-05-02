package spec

import (
	"fmt"
	"go/ast"
	"go/token"
)

type VarDeclCases interface {
	IdentVarSpec(d *ast.GenDecl, s *ast.ValueSpec, t *ast.Ident)
	NilVarSpec(d *ast.GenDecl, s *ast.ValueSpec)
	FuncVarSpec(d *ast.GenDecl, s *ast.ValueSpec, t *ast.FuncType)
	SelectorVarSpec(d *ast.GenDecl, s *ast.ValueSpec, t *ast.SelectorExpr)
	MapVarSpec(d *ast.GenDecl, s *ast.ValueSpec, t *ast.MapType)
	StarVarSpec(d *ast.GenDecl, s *ast.ValueSpec, t *ast.StarExpr)
}

func CaseVarDecl(file *ast.File, w VarDeclCases) {
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			switch d.Tok {
			case token.VAR:
				for _, spec := range d.Specs {
					if s, ok := spec.(*ast.ValueSpec); ok {
						switch t := s.Type.(type) {
						case *ast.Ident:
							w.IdentVarSpec(d, s, t)
						case *ast.FuncType:
							w.FuncVarSpec(d, s, t)
						case *ast.SelectorExpr:
							w.SelectorVarSpec(d, s, t)
						case *ast.MapType:
							w.MapVarSpec(d, s, t)
						case *ast.StarExpr:
							w.StarVarSpec(d, s, t)
						case nil:
							w.NilVarSpec(d, s)
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
