package spec

import "go/ast"

type (
	FnIdentConstSpec = func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.Ident, unused bool)
	FnNilConstSpec   = func(d *ast.GenDecl, s *ast.ValueSpec, unused bool)
)

func ExportedFnIdentConstSpec(fn FnIdentConstSpec) FnIdentConstSpec {
	return func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.Ident, unused bool) {
		for _, name := range s.Names {
			if name.IsExported() {
				fn(d, s, t, unused)
			}
		}
	}
}
func ExportedFnNilConstSpec(fn FnNilConstSpec) FnNilConstSpec {
	return func(d *ast.GenDecl, s *ast.ValueSpec, unused bool) {
		for _, name := range s.Names {
			if name.IsExported() {
				fn(d, s, unused)
			}
		}
	}
}
