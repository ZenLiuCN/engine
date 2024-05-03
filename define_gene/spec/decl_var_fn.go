package spec

import "go/ast"

type (
	FnIdentVarSpec     = func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.Ident)
	FnNilVarSpec       = func(d *ast.GenDecl, s *ast.ValueSpec)
	FnFuncVarSpec      = func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.FuncType)
	FnSelectorVarSpec  = func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.SelectorExpr)
	FnMapVarSpec       = func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.MapType)
	FnArrayVarSpec     = func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.ArrayType)
	FnInterfaceVarSpec = func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.InterfaceType)
	FnStarVarSpec      = func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.StarExpr)
)

func ExportedFnIdentVarSpec(fn FnIdentVarSpec) FnIdentVarSpec {
	return func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.Ident) {
		for _, name := range s.Names {
			if name.IsExported() {
				fn(d, s, t)
			}
		}
	}
}
func ExportedFnNilVarSpec(fn FnNilVarSpec) FnNilVarSpec {
	return func(d *ast.GenDecl, s *ast.ValueSpec) {
		for _, name := range s.Names {
			if name.IsExported() {
				fn(d, s)
			}
		}
	}
}
func ExportedFnFuncVarSpec(fn FnFuncVarSpec) FnFuncVarSpec {
	return func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.FuncType) {
		for _, name := range s.Names {
			if name.IsExported() {
				fn(d, s, t)
			}
		}
	}
}
func ExportedFnSelectorVarSpec(fn FnSelectorVarSpec) FnSelectorVarSpec {
	return func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.SelectorExpr) {
		for _, name := range s.Names {
			if name.IsExported() {
				fn(d, s, t)
			}
		}
	}
}
func ExportedFnMapVarSpec(fn FnMapVarSpec) FnMapVarSpec {
	return func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.MapType) {
		for _, name := range s.Names {
			if name.IsExported() {
				fn(d, s, t)
			}
		}
	}
}
func ExportedFnStarVarSpec(fn FnStarVarSpec) FnStarVarSpec {
	return func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.StarExpr) {
		for _, name := range s.Names {
			if name.IsExported() {
				fn(d, s, t)
			}
		}
	}
}
