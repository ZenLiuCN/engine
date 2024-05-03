package spec

import "go/ast"

type (
	FnFuncDeclSpec            = func(d *ast.FuncDecl)
	FnMethodDeclSpec          = func(d *ast.FuncDecl, r *ast.Field) bool //return ture to walk for receivers
	FnIdentTypeSpec           = func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.Ident)
	FnStructTypeSpec          = func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.StructType)
	FnInterfaceTypeSpec       = func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.InterfaceType)
	FnMapTypeSpec             = func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.MapType)
	FnArrayTypeSpec           = func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.ArrayType)
	FnFuncTypeSpec            = func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.FuncType)
	FnMethodDeclStarRecvSpec  = func(d *ast.FuncDecl, r *ast.Field, t *ast.StarExpr)
	FnMethodDeclIdentRecvSpec = func(d *ast.FuncDecl, r *ast.Field, t *ast.Ident)
	FnChanTypeSpec            = func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.ChanType)
	FnSelectorTypeSpec        = func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.SelectorExpr)
	FnStarExprSpec            = func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.StarExpr)
)

func ExportedFnStarExprSpec(fn FnStarExprSpec) FnStarExprSpec {
	return func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.StarExpr) {
		if s.Name.IsExported() {
			fn(d, s, t)
		}
	}
}
func ExportedFnFuncDecl(fn FnFuncDeclSpec) FnFuncDeclSpec {
	return func(s *ast.FuncDecl) {
		if s.Name.IsExported() {
			fn(s)
		}
	}
}
func ExportedFnMethodDecl(fn FnMethodDeclSpec) FnMethodDeclSpec {
	return func(s *ast.FuncDecl, t *ast.Field) bool {
		if s.Name.IsExported() {
			return fn(s, t)
		}
		return false
	}
}
func ExportedFnIdentTypeSpec(fn FnIdentTypeSpec) FnIdentTypeSpec {
	return func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.Ident) {
		if s.Name.IsExported() {
			fn(d, s, t)
		}
	}
}
func ExportedFnStructTypeSpec(fn FnStructTypeSpec) FnStructTypeSpec {
	return func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.StructType) {
		if s.Name.IsExported() {
			fn(d, s, t)
		}
	}
}
func ExportedFnInterfaceTypeSpec(fn FnInterfaceTypeSpec) FnInterfaceTypeSpec {
	return func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.InterfaceType) {
		if s.Name.IsExported() {
			fn(d, s, t)
		}
	}
}
func ExportedFnMapTypeSpec(fn FnMapTypeSpec) FnMapTypeSpec {
	return func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.MapType) {
		if s.Name.IsExported() {
			fn(d, s, t)
		}
	}
}
func ExportedFnArrayTypeSpec(fn FnArrayTypeSpec) FnArrayTypeSpec {
	return func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.ArrayType) {
		if s.Name.IsExported() {
			fn(d, s, t)
		}
	}
}
func ExportedFnFuncTypeSpec(fn FnFuncTypeSpec) FnFuncTypeSpec {
	return func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.FuncType) {
		if s.Name.IsExported() {
			fn(d, s, t)
		}
	}
}
func ExportedFnMethodDeclStarRecv(fn FnMethodDeclStarRecvSpec) FnMethodDeclStarRecvSpec {
	return func(s *ast.FuncDecl, r *ast.Field, t *ast.StarExpr) {
		if s.Name.IsExported() {
			fn(s, r, t)
		}
	}
}
func ExportedFnMethodDeclIdentRecv(fn FnMethodDeclIdentRecvSpec) FnMethodDeclIdentRecvSpec {
	return func(s *ast.FuncDecl, r *ast.Field, t *ast.Ident) {
		if s.Name.IsExported() {
			fn(s, r, t)
		}
	}
}
func ExportedFnChanTypeSpec(fn FnChanTypeSpec) FnChanTypeSpec {
	return func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.ChanType) {
		if s.Name.IsExported() {
			fn(d, s, t)
		}
	}
}
func ExportedFnSelectorTypeSpec(fn FnSelectorTypeSpec) FnSelectorTypeSpec {
	return func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.SelectorExpr) {
		if s.Name.IsExported() {
			fn(d, s, t)
		}
	}
}
