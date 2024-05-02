package spec

import "go/ast"

type (
	FnFuncDecl            = func(d *ast.FuncDecl)
	FnMethodDecl          = func(d *ast.FuncDecl, r *ast.Field) bool //return ture to walk for receivers
	FnIdentTypeSpec       = func(s *ast.TypeSpec, t *ast.Ident)
	FnStructTypeSpec      = func(s *ast.TypeSpec, t *ast.StructType)
	FnInterfaceTypeSpec   = func(s *ast.TypeSpec, t *ast.InterfaceType)
	FnMapTypeSpec         = func(s *ast.TypeSpec, t *ast.MapType)
	FnArrayTypeSpec       = func(s *ast.TypeSpec, t *ast.ArrayType)
	FnFuncTypeSpec        = func(s *ast.TypeSpec, t *ast.FuncType)
	FnMethodDeclStarRecv  = func(d *ast.FuncDecl, r *ast.Field, t *ast.StarExpr)
	FnMethodDeclIdentRecv = func(d *ast.FuncDecl, r *ast.Field, t *ast.Ident)
	FnChanTypeSpec        = func(s *ast.TypeSpec, t *ast.ChanType)
	FnSelectorTypeSpec    = func(s *ast.TypeSpec, t *ast.SelectorExpr)
	FnStarExprSpec        = func(s *ast.TypeSpec, t *ast.StarExpr)
)

func ExportedFnStarExprSpec(fn FnStarExprSpec) FnStarExprSpec {
	return func(s *ast.TypeSpec, t *ast.StarExpr) {
		if s.Name.IsExported() {
			fn(s, t)
		}
	}
}
func ExportedFnFuncDecl(fn FnFuncDecl) FnFuncDecl {
	return func(s *ast.FuncDecl) {
		if s.Name.IsExported() {
			fn(s)
		}
	}
}
func ExportedFnMethodDecl(fn FnMethodDecl) FnMethodDecl {
	return func(s *ast.FuncDecl, t *ast.Field) bool {
		if s.Name.IsExported() {
			return fn(s, t)
		}
		return false
	}
}
func ExportedFnIdentTypeSpec(fn FnIdentTypeSpec) FnIdentTypeSpec {
	return func(s *ast.TypeSpec, t *ast.Ident) {
		if s.Name.IsExported() {
			fn(s, t)
		}
	}
}
func ExportedFnStructTypeSpec(fn FnStructTypeSpec) FnStructTypeSpec {
	return func(s *ast.TypeSpec, t *ast.StructType) {
		if s.Name.IsExported() {
			fn(s, t)
		}
	}
}
func ExportedFnInterfaceTypeSpec(fn FnInterfaceTypeSpec) FnInterfaceTypeSpec {
	return func(s *ast.TypeSpec, t *ast.InterfaceType) {
		if s.Name.IsExported() {
			fn(s, t)
		}
	}
}
func ExportedFnMapTypeSpec(fn FnMapTypeSpec) FnMapTypeSpec {
	return func(s *ast.TypeSpec, t *ast.MapType) {
		if s.Name.IsExported() {
			fn(s, t)
		}
	}
}
func ExportedFnArrayTypeSpec(fn FnArrayTypeSpec) FnArrayTypeSpec {
	return func(s *ast.TypeSpec, t *ast.ArrayType) {
		if s.Name.IsExported() {
			fn(s, t)
		}
	}
}
func ExportedFnFuncTypeSpec(fn FnFuncTypeSpec) FnFuncTypeSpec {
	return func(s *ast.TypeSpec, t *ast.FuncType) {
		if s.Name.IsExported() {
			fn(s, t)
		}
	}
}
func ExportedFnMethodDeclStarRecv(fn FnMethodDeclStarRecv) FnMethodDeclStarRecv {
	return func(s *ast.FuncDecl, r *ast.Field, t *ast.StarExpr) {
		if s.Name.IsExported() {
			fn(s, r, t)
		}
	}
}
func ExportedFnMethodDeclIdentRecv(fn FnMethodDeclIdentRecv) FnMethodDeclIdentRecv {
	return func(s *ast.FuncDecl, r *ast.Field, t *ast.Ident) {
		if s.Name.IsExported() {
			fn(s, r, t)
		}
	}
}
func ExportedFnChanTypeSpec(fn FnChanTypeSpec) FnChanTypeSpec {
	return func(s *ast.TypeSpec, t *ast.ChanType) {
		if s.Name.IsExported() {
			fn(s, t)
		}
	}
}
func ExportedFnSelectorTypeSpec(fn FnSelectorTypeSpec) FnSelectorTypeSpec {
	return func(s *ast.TypeSpec, t *ast.SelectorExpr) {
		if s.Name.IsExported() {
			fn(s, t)
		}
	}
}
