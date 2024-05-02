package spec

import (
	"fmt"
	"go/ast"
	"go/token"
)

type ConstDeclCases interface {
	IdentConstSpec(d *ast.GenDecl, s *ast.ValueSpec, t *ast.Ident)
	NilConstSpec(d *ast.GenDecl, s *ast.ValueSpec)
}

//region ExportedConstDeclCases

type ExportedConstDeclCases struct {
	Inner ConstDeclCases
}

func (e ExportedConstDeclCases) IdentConstSpec(d *ast.GenDecl, s *ast.ValueSpec, t *ast.Ident) {
	for _, name := range s.Names {
		if ast.IsExported(name.Name) {
			e.Inner.IdentConstSpec(d, s, t)
			break
		}
	}
}

func (e ExportedConstDeclCases) NilConstSpec(d *ast.GenDecl, s *ast.ValueSpec) {
	for _, name := range s.Names {
		if ast.IsExported(name.Name) {
			e.Inner.NilConstSpec(d, s)
			break
		}
	}
}

//endregion

func CaseConstDecl(file *ast.File, w ConstDeclCases) {
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			switch d.Tok {
			case token.CONST:
				for _, spec := range d.Specs {
					if s, ok := spec.(*ast.ValueSpec); ok {
						switch t := s.Type.(type) {
						case *ast.Ident:
							w.IdentConstSpec(d, s, t)
						case nil:
							w.NilConstSpec(d, s)
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
