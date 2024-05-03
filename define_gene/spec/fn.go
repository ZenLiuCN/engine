package spec

import (
	"fmt"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"log"
)

type Func int

const (
	FuncFnNone Func = iota
	FuncTypeFnIdentType
	FuncTypeFnSelectorType
	FuncTypeFnStarType
	FuncTypeFnFuncType
	FuncTypeFnArrayType
	FuncTypeFnMapType
	FuncTypeFnStructType
	FuncTypeFnChanType
	FuncTypeFnEllipsis
	FuncTypeFnInterfaceType
	FuncSpecFnFuncDeclSpec
	FuncSpecFnMethodDeclSpec
	FuncSpecFnIdentTypeSpec
	FuncSpecFnStructTypeSpec
	FuncSpecFnInterfaceTypeSpec
	FuncSpecFnMapTypeSpec
	FuncSpecFnArrayTypeSpec
	FuncSpecFnFuncTypeSpec
	FuncSpecFnMethodDeclStarRecvSpec
	FuncSpecFnMethodDeclIdentRecvSpec
	FuncSpecFnChanTypeSpec
	FuncSpecFnSelectorTypeSpec
	FuncSpecFnStarExprSpec
	FuncConsFnIdentConstSpec
	FuncConsFnNilConstSpec
	FuncVarsFnIdentVarSpec
	FuncVarsFnNilVarSpec
	FuncVarsFnFuncVarSpec
	FuncVarsFnSelectorVarSpec
	FuncVarsFnMapVarSpec
	FuncVarsFnStarVarSpec
	FuncVarsFnArrayVarSpec
	FuncVarsFnInterfaceVarSpec

	FuncFnMax
)

type FnProvider = func(node Func) any
type Providers map[Func]any

func (p Providers) Provide(node Func) any {
	if fn, ok := p[node]; ok {
		return fn
	}
	return nil
}
func (p Providers) WithAllType(
	funcTypeFnIdentType FnIdentType,
	funcTypeFnSelectorType FnSelectorType,
	funcTypeFnStarType FnStarType,
	funcTypeFnFuncType FnFuncType,
	funcTypeFnArrayType FnArrayType,
	funcTypeFnMapType FnMapType,
	funcTypeFnStructType FnStructType,
	funcTypeFnChanType FnChanType,
	funcTypeFnEllipsis FnEllipsis,
	funcTypeFnInterfaceType FnInterfaceType,
) {
	p[FuncTypeFnIdentType] = funcTypeFnIdentType
	p[FuncTypeFnSelectorType] = funcTypeFnSelectorType
	p[FuncTypeFnStarType] = funcTypeFnStarType
	p[FuncTypeFnFuncType] = funcTypeFnFuncType
	p[FuncTypeFnArrayType] = funcTypeFnArrayType
	p[FuncTypeFnMapType] = funcTypeFnMapType
	p[FuncTypeFnStructType] = funcTypeFnStructType
	p[FuncTypeFnChanType] = funcTypeFnChanType
	p[FuncTypeFnEllipsis] = funcTypeFnEllipsis
	p[FuncTypeFnInterfaceType] = funcTypeFnInterfaceType
}
func (p Providers) WithAllSpec(
	funcSpecFnFuncDeclSpec FnFuncDeclSpec,
	funcSpecFnMethodDeclSpec FnMethodDeclSpec,
	funcSpecFnIdentTypeSpec FnIdentTypeSpec,
	funcSpecFnStructTypeSpec FnStructTypeSpec,
	funcSpecFnInterfaceTypeSpec FnInterfaceTypeSpec,
	funcSpecFnMapTypeSpec FnMapTypeSpec,
	funcSpecFnArrayTypeSpec FnArrayTypeSpec,
	funcSpecFnFuncTypeSpec FnFuncTypeSpec,
	funcSpecFnMethodDeclStarRecvSpec FnMethodDeclStarRecvSpec,
	funcSpecFnMethodDeclIdentRecvSpec FnMethodDeclIdentRecvSpec,
	funcSpecFnChanTypeSpec FnChanTypeSpec,
	funcSpecFnSelectorTypeSpec FnSelectorTypeSpec,
	funcSpecFnStarExprSpec FnStarExprSpec,
) {
	p[FuncSpecFnFuncDeclSpec] = funcSpecFnFuncDeclSpec
	p[FuncSpecFnMethodDeclSpec] = funcSpecFnMethodDeclSpec
	p[FuncSpecFnIdentTypeSpec] = funcSpecFnIdentTypeSpec
	p[FuncSpecFnStructTypeSpec] = funcSpecFnStructTypeSpec
	p[FuncSpecFnInterfaceTypeSpec] = funcSpecFnInterfaceTypeSpec
	p[FuncSpecFnMapTypeSpec] = funcSpecFnMapTypeSpec
	p[FuncSpecFnArrayTypeSpec] = funcSpecFnArrayTypeSpec
	p[FuncSpecFnFuncTypeSpec] = funcSpecFnFuncTypeSpec
	p[FuncSpecFnMethodDeclStarRecvSpec] = funcSpecFnMethodDeclStarRecvSpec
	p[FuncSpecFnMethodDeclIdentRecvSpec] = funcSpecFnMethodDeclIdentRecvSpec
	p[FuncSpecFnChanTypeSpec] = funcSpecFnChanTypeSpec
	p[FuncSpecFnSelectorTypeSpec] = funcSpecFnSelectorTypeSpec
	p[FuncSpecFnStarExprSpec] = funcSpecFnStarExprSpec
}
func (p Providers) WithAllConst(
	funcConstFnIdentConstSpec FnIdentConstSpec,
	funcConstFnNilConstSpec FnNilConstSpec,
) {
	p[FuncConsFnIdentConstSpec] = funcConstFnIdentConstSpec
	p[FuncConsFnNilConstSpec] = funcConstFnNilConstSpec

}
func (p Providers) WithAllVar(
	funcVarFnIdentVarSpec FnIdentVarSpec,
	funcVarFnNilVarSpec FnNilVarSpec,
	funcVarFnFuncVarSpec FnFuncVarSpec,
	funcVarFnSelectorVarSpec FnSelectorVarSpec,
	funcVarFnMapVarSpec FnMapVarSpec,
	funcVarFnStarVarSpec FnStarVarSpec,
	funcVarFnArrayVarSpec FnArrayVarSpec,
	funcVarFnInterfaceVarSpec FnInterfaceVarSpec,
) {
	p[FuncVarsFnIdentVarSpec] = funcVarFnIdentVarSpec
	p[FuncVarsFnNilVarSpec] = funcVarFnNilVarSpec
	p[FuncVarsFnFuncVarSpec] = funcVarFnFuncVarSpec
	p[FuncVarsFnSelectorVarSpec] = funcVarFnSelectorVarSpec
	p[FuncVarsFnMapVarSpec] = funcVarFnMapVarSpec
	p[FuncVarsFnStarVarSpec] = funcVarFnStarVarSpec
	p[FuncVarsFnArrayVarSpec] = funcVarFnArrayVarSpec
	p[FuncVarsFnInterfaceVarSpec] = funcVarFnInterfaceVarSpec

}
func (p Providers) WithAll(
	funcTypeFnIdentType FnIdentType,
	funcTypeFnSelectorType FnSelectorType,
	funcTypeFnStarType FnStarType,
	funcTypeFnFuncType FnFuncType,
	funcTypeFnArrayType FnArrayType,
	funcTypeFnMapType FnMapType,
	funcTypeFnStructType FnStructType,
	funcTypeFnChanType FnChanType,
	funcTypeFnEllipsis FnEllipsis,
	funcTypeFnInterfaceType FnInterfaceType,
	funcSpecFnFuncDeclSpec FnFuncDeclSpec,
	funcSpecFnMethodDeclSpec FnMethodDeclSpec,
	funcSpecFnIdentTypeSpec FnIdentTypeSpec,
	funcSpecFnStructTypeSpec FnStructTypeSpec,
	funcSpecFnInterfaceTypeSpec FnInterfaceTypeSpec,
	funcSpecFnMapTypeSpec FnMapTypeSpec,
	funcSpecFnArrayTypeSpec FnArrayTypeSpec,
	funcSpecFnFuncTypeSpec FnFuncTypeSpec,
	funcSpecFnMethodDeclStarRecvSpec FnMethodDeclStarRecvSpec,
	funcSpecFnMethodDeclIdentRecvSpec FnMethodDeclIdentRecvSpec,
	funcSpecFnChanTypeSpec FnChanTypeSpec,
	funcSpecFnSelectorTypeSpec FnSelectorTypeSpec,
	funcSpecFnStarExprSpec FnStarExprSpec,
	funcConstFnIdentConstSpec FnIdentConstSpec,
	funcConstFnNilConstSpec FnNilConstSpec,
	funcVarFnIdentVarSpec FnIdentVarSpec,
	funcVarFnNilVarSpec FnNilVarSpec,
	funcVarFnFuncVarSpec FnFuncVarSpec,
	funcVarFnSelectorVarSpec FnSelectorVarSpec,
	funcVarFnMapVarSpec FnMapVarSpec,
	funcVarFnStarVarSpec FnStarVarSpec,
	funcVarFnArrayVarSpec FnArrayVarSpec,
	funcVarFnInterfaceVarSpec FnInterfaceVarSpec,
) {
	p[FuncTypeFnIdentType] = funcTypeFnIdentType
	p[FuncTypeFnSelectorType] = funcTypeFnSelectorType
	p[FuncTypeFnStarType] = funcTypeFnStarType
	p[FuncTypeFnFuncType] = funcTypeFnFuncType
	p[FuncTypeFnArrayType] = funcTypeFnArrayType
	p[FuncTypeFnMapType] = funcTypeFnMapType
	p[FuncTypeFnStructType] = funcTypeFnStructType
	p[FuncTypeFnChanType] = funcTypeFnChanType
	p[FuncTypeFnEllipsis] = funcTypeFnEllipsis
	p[FuncTypeFnInterfaceType] = funcTypeFnInterfaceType
	p[FuncSpecFnFuncDeclSpec] = funcSpecFnFuncDeclSpec
	p[FuncSpecFnMethodDeclSpec] = funcSpecFnMethodDeclSpec
	p[FuncSpecFnIdentTypeSpec] = funcSpecFnIdentTypeSpec
	p[FuncSpecFnStructTypeSpec] = funcSpecFnStructTypeSpec
	p[FuncSpecFnInterfaceTypeSpec] = funcSpecFnInterfaceTypeSpec
	p[FuncSpecFnMapTypeSpec] = funcSpecFnMapTypeSpec
	p[FuncSpecFnArrayTypeSpec] = funcSpecFnArrayTypeSpec
	p[FuncSpecFnFuncTypeSpec] = funcSpecFnFuncTypeSpec
	p[FuncSpecFnMethodDeclStarRecvSpec] = funcSpecFnMethodDeclStarRecvSpec
	p[FuncSpecFnMethodDeclIdentRecvSpec] = funcSpecFnMethodDeclIdentRecvSpec
	p[FuncSpecFnChanTypeSpec] = funcSpecFnChanTypeSpec
	p[FuncSpecFnSelectorTypeSpec] = funcSpecFnSelectorTypeSpec
	p[FuncSpecFnStarExprSpec] = funcSpecFnStarExprSpec
	p[FuncConsFnIdentConstSpec] = funcConstFnIdentConstSpec
	p[FuncConsFnNilConstSpec] = funcConstFnNilConstSpec
	p[FuncVarsFnIdentVarSpec] = funcVarFnIdentVarSpec
	p[FuncVarsFnNilVarSpec] = funcVarFnNilVarSpec
	p[FuncVarsFnFuncVarSpec] = funcVarFnFuncVarSpec
	p[FuncVarsFnSelectorVarSpec] = funcVarFnSelectorVarSpec
	p[FuncVarsFnMapVarSpec] = funcVarFnMapVarSpec
	p[FuncVarsFnStarVarSpec] = funcVarFnStarVarSpec
	p[FuncVarsFnArrayVarSpec] = funcVarFnArrayVarSpec
	p[FuncVarsFnInterfaceVarSpec] = funcVarFnInterfaceVarSpec
}
func (p Providers) WithTypeFnIdentType(fn FnIdentType) FnIdentType {
	if f, ok := p[FuncTypeFnIdentType]; ok {
		p[FuncTypeFnIdentType] = fn
		return f.(FnIdentType)
	}
	p[FuncTypeFnIdentType] = fn
	return nil
}
func (p Providers) WithTypeFnSelectorType(fn FnSelectorType) FnSelectorType {
	if f, ok := p[FuncTypeFnSelectorType]; ok {
		p[FuncTypeFnSelectorType] = fn
		return f.(FnSelectorType)
	}
	p[FuncTypeFnSelectorType] = fn
	return nil
}
func (p Providers) WithTypeFnStarType(fn FnStarType) FnStarType {
	if f, ok := p[FuncTypeFnStarType]; ok {
		p[FuncTypeFnStarType] = fn
		return f.(FnStarType)
	}
	p[FuncTypeFnStarType] = fn
	return nil
}
func (p Providers) WithTypeFnFuncType(fn FnFuncType) FnFuncType {
	if f, ok := p[FuncTypeFnFuncType]; ok {
		p[FuncTypeFnFuncType] = fn
		return f.(FnFuncType)
	}
	p[FuncTypeFnFuncType] = fn
	return nil
}
func (p Providers) WithTypeFnArrayType(fn FnArrayType) FnArrayType {
	if f, ok := p[FuncTypeFnArrayType]; ok {
		p[FuncTypeFnArrayType] = fn
		return f.(FnArrayType)
	}
	p[FuncTypeFnArrayType] = fn
	return nil
}
func (p Providers) WithTypeFnMapType(fn FnMapType) FnMapType {
	if f, ok := p[FuncTypeFnMapType]; ok {
		p[FuncTypeFnMapType] = fn
		return f.(FnMapType)
	}
	p[FuncTypeFnMapType] = fn
	return nil
}
func (p Providers) WithTypeFnStructType(fn FnStructType) FnStructType {
	if f, ok := p[FuncTypeFnStructType]; ok {
		p[FuncTypeFnStructType] = fn
		return f.(FnStructType)
	}
	p[FuncTypeFnStructType] = fn
	return nil
}
func (p Providers) WithTypeFnChanType(fn FnChanType) FnChanType {
	if f, ok := p[FuncTypeFnChanType]; ok {
		p[FuncTypeFnChanType] = fn
		return f.(FnChanType)
	}
	p[FuncTypeFnChanType] = fn
	return nil
}
func (p Providers) WithTypeFnEllipsis(fn FnEllipsis) FnEllipsis {
	if f, ok := p[FuncTypeFnEllipsis]; ok {
		p[FuncTypeFnEllipsis] = fn
		return f.(FnEllipsis)
	}
	p[FuncTypeFnEllipsis] = fn
	return nil
}
func (p Providers) WithTypeFnInterfaceType(fn FnInterfaceType) FnInterfaceType {
	if f, ok := p[FuncTypeFnInterfaceType]; ok {
		p[FuncTypeFnInterfaceType] = fn
		return f.(FnInterfaceType)
	}
	p[FuncTypeFnInterfaceType] = fn
	return nil
}
func (p Providers) WithSpecFnFuncDeclSpec(fn FnFuncDeclSpec) FnFuncDeclSpec {
	if f, ok := p[FuncSpecFnFuncDeclSpec]; ok {
		p[FuncSpecFnFuncDeclSpec] = fn
		return f.(FnFuncDeclSpec)
	}
	p[FuncSpecFnFuncDeclSpec] = fn
	return nil
}
func (p Providers) WithSpecFnMethodDeclSpec(fn FnMethodDeclSpec) FnMethodDeclSpec {
	if f, ok := p[FuncSpecFnMethodDeclSpec]; ok {
		p[FuncSpecFnMethodDeclSpec] = fn
		return f.(FnMethodDeclSpec)
	}
	p[FuncSpecFnMethodDeclSpec] = fn
	return nil
}
func (p Providers) WithSpecFnIdentTypeSpec(fn FnIdentTypeSpec) FnIdentTypeSpec {
	if f, ok := p[FuncSpecFnIdentTypeSpec]; ok {
		p[FuncSpecFnIdentTypeSpec] = fn
		return f.(FnIdentTypeSpec)
	}
	p[FuncSpecFnIdentTypeSpec] = fn
	return nil
}
func (p Providers) WithSpecFnStructTypeSpec(fn FnStructTypeSpec) FnStructTypeSpec {
	if f, ok := p[FuncSpecFnStructTypeSpec]; ok {
		p[FuncSpecFnStructTypeSpec] = fn
		return f.(FnStructTypeSpec)
	}
	p[FuncSpecFnStructTypeSpec] = fn
	return nil
}
func (p Providers) WithSpecFnInterfaceTypeSpec(fn FnInterfaceTypeSpec) FnInterfaceTypeSpec {
	if f, ok := p[FuncSpecFnInterfaceTypeSpec]; ok {
		p[FuncSpecFnInterfaceTypeSpec] = fn
		return f.(FnInterfaceTypeSpec)
	}
	p[FuncSpecFnInterfaceTypeSpec] = fn
	return nil
}
func (p Providers) WithSpecFnMapTypeSpec(fn FnMapTypeSpec) FnMapTypeSpec {
	if f, ok := p[FuncSpecFnMapTypeSpec]; ok {
		p[FuncSpecFnMapTypeSpec] = fn
		return f.(FnMapTypeSpec)
	}
	p[FuncSpecFnMapTypeSpec] = fn
	return nil
}
func (p Providers) WithSpecFnArrayTypeSpec(fn FnArrayTypeSpec) FnArrayTypeSpec {
	if f, ok := p[FuncSpecFnArrayTypeSpec]; ok {
		p[FuncSpecFnArrayTypeSpec] = fn
		return f.(FnArrayTypeSpec)
	}
	p[FuncSpecFnArrayTypeSpec] = fn
	return nil
}
func (p Providers) WithSpecFnFuncTypeSpec(fn FnFuncTypeSpec) FnFuncTypeSpec {
	if f, ok := p[FuncSpecFnFuncTypeSpec]; ok {
		p[FuncSpecFnFuncTypeSpec] = fn
		return f.(FnFuncTypeSpec)
	}
	p[FuncSpecFnFuncTypeSpec] = fn
	return nil
}
func (p Providers) WithSpecFnMethodDeclStarRecvSpec(fn FnMethodDeclStarRecvSpec) FnMethodDeclStarRecvSpec {
	if f, ok := p[FuncSpecFnMethodDeclStarRecvSpec]; ok {
		p[FuncSpecFnMethodDeclStarRecvSpec] = fn
		return f.(FnMethodDeclStarRecvSpec)
	}
	p[FuncSpecFnMethodDeclStarRecvSpec] = fn
	return nil
}
func (p Providers) WithSpecFnMethodDeclIdentRecvSpec(fn FnMethodDeclIdentRecvSpec) FnMethodDeclIdentRecvSpec {
	if f, ok := p[FuncSpecFnMethodDeclIdentRecvSpec]; ok {
		p[FuncSpecFnMethodDeclIdentRecvSpec] = fn
		return f.(FnMethodDeclIdentRecvSpec)
	}
	p[FuncSpecFnMethodDeclIdentRecvSpec] = fn
	return nil
}
func (p Providers) WithSpecFnChanTypeSpec(fn FnChanTypeSpec) FnChanTypeSpec {
	if f, ok := p[FuncSpecFnChanTypeSpec]; ok {
		p[FuncSpecFnChanTypeSpec] = fn
		return f.(FnChanTypeSpec)
	}
	p[FuncSpecFnChanTypeSpec] = fn
	return nil
}
func (p Providers) WithSpecFnSelectorTypeSpec(fn FnSelectorTypeSpec) FnSelectorTypeSpec {
	if f, ok := p[FuncSpecFnSelectorTypeSpec]; ok {
		p[FuncSpecFnSelectorTypeSpec] = fn
		return f.(FnSelectorTypeSpec)
	}
	p[FuncSpecFnSelectorTypeSpec] = fn
	return nil
}
func (p Providers) WithSpecFnStarExprSpec(fn FnStarExprSpec) FnStarExprSpec {
	if f, ok := p[FuncSpecFnStarExprSpec]; ok {
		p[FuncSpecFnStarExprSpec] = fn
		return f.(FnStarExprSpec)
	}
	p[FuncSpecFnStarExprSpec] = fn
	return nil
}
func (p Providers) WithConstFnIdentConstSpec(fn FnIdentConstSpec) FnIdentConstSpec {
	if f, ok := p[FuncConsFnIdentConstSpec]; ok {
		p[FuncConsFnIdentConstSpec] = fn
		return f.(FnIdentConstSpec)
	}
	p[FuncConsFnIdentConstSpec] = fn
	return nil
}
func (p Providers) WithConstFnNilConstSpec(fn FnNilConstSpec) FnNilConstSpec {
	if f, ok := p[FuncConsFnNilConstSpec]; ok {
		p[FuncConsFnNilConstSpec] = fn
		return f.(FnNilConstSpec)
	}
	p[FuncConsFnNilConstSpec] = fn
	return nil
}
func (p Providers) WithVarFnIdentVarSpec(fn FnIdentVarSpec) FnIdentVarSpec {
	if f, ok := p[FuncVarsFnIdentVarSpec]; ok {
		p[FuncVarsFnIdentVarSpec] = fn
		return f.(FnIdentVarSpec)
	}
	p[FuncVarsFnIdentVarSpec] = fn
	return nil
}
func (p Providers) WithVarFnNilVarSpec(fn FnNilVarSpec) FnNilVarSpec {
	if f, ok := p[FuncVarsFnNilVarSpec]; ok {
		p[FuncVarsFnNilVarSpec] = fn
		return f.(FnNilVarSpec)
	}
	p[FuncVarsFnNilVarSpec] = fn
	return nil
}
func (p Providers) WithVarFnFuncVarSpec(fn FnFuncVarSpec) FnFuncVarSpec {
	if f, ok := p[FuncVarsFnFuncVarSpec]; ok {
		p[FuncVarsFnFuncVarSpec] = fn
		return f.(FnFuncVarSpec)
	}
	p[FuncVarsFnFuncVarSpec] = fn
	return nil
}
func (p Providers) WithVarFnSelectorVarSpec(fn FnSelectorVarSpec) FnSelectorVarSpec {
	if f, ok := p[FuncVarsFnSelectorVarSpec]; ok {
		p[FuncVarsFnSelectorVarSpec] = fn
		return f.(FnSelectorVarSpec)
	}
	p[FuncVarsFnSelectorVarSpec] = fn
	return nil
}
func (p Providers) WithVarFnMapVarSpec(fn FnMapVarSpec) FnMapVarSpec {
	if f, ok := p[FuncVarsFnMapVarSpec]; ok {
		p[FuncVarsFnMapVarSpec] = fn
		return f.(FnMapVarSpec)
	}
	p[FuncVarsFnMapVarSpec] = fn
	return nil
}
func (p Providers) WithVarFnStarVarSpec(fn FnStarVarSpec) FnStarVarSpec {
	if f, ok := p[FuncVarsFnStarVarSpec]; ok {
		p[FuncVarsFnStarVarSpec] = fn
		return f.(FnStarVarSpec)
	}
	p[FuncVarsFnStarVarSpec] = fn
	return nil
}
func (p Providers) WithVarFnArrayVarSpec(fn FnArrayVarSpec) FnArrayVarSpec {
	if f, ok := p[FuncVarsFnArrayVarSpec]; ok {
		p[FuncVarsFnArrayVarSpec] = fn
		return f.(FnArrayVarSpec)
	}
	p[FuncVarsFnArrayVarSpec] = fn
	return nil
}
func (p Providers) WithVarFnInterfaceVarSpec(fn FnInterfaceVarSpec) FnInterfaceVarSpec {
	if f, ok := p[FuncVarsFnInterfaceVarSpec]; ok {
		p[FuncVarsFnInterfaceVarSpec] = fn
		return f.(FnInterfaceVarSpec)
	}
	p[FuncVarsFnInterfaceVarSpec] = fn
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
							if fn, ok := w(FuncConsFnIdentConstSpec).(FnIdentConstSpec); ok {
								fn(d, s, t, ok)
							}
						case nil:
							if fn, ok := w(FuncConsFnNilConstSpec).(FnNilConstSpec); ok {
								fn(d, s, ok)
							}
						default:
							log.Printf("miss const decl %#+v\n", t)
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
							if fn, ok := w(FuncVarsFnIdentVarSpec).(FnIdentVarSpec); ok {
								fn(d, s, t)
							}
						case *ast.FuncType:
							if fn, ok := w(FuncVarsFnFuncVarSpec).(FnFuncVarSpec); ok {
								fn(d, s, t)
							}
						case *ast.SelectorExpr:
							if fn, ok := w(FuncVarsFnSelectorVarSpec).(FnSelectorVarSpec); ok {
								fn(d, s, t)
							}
						case *ast.MapType:
							if fn, ok := w(FuncVarsFnMapVarSpec).(FnMapVarSpec); ok {
								fn(d, s, t)
							}
						case *ast.ArrayType:
							if fn, ok := w(FuncVarsFnArrayVarSpec).(FnArrayVarSpec); ok {
								fn(d, s, t)
							}
						case *ast.InterfaceType:
							if fn, ok := w(FuncVarsFnInterfaceVarSpec).(FnInterfaceVarSpec); ok {
								fn(d, s, t)
							}
						case *ast.StarExpr:
							if fn, ok := w(FuncVarsFnStarVarSpec).(FnStarVarSpec); ok {
								fn(d, s, t)
							}
						case nil:
							if fn, ok := w(FuncVarsFnNilVarSpec).(FnNilVarSpec); ok {
								fn(d, s)
							}
						default:
							log.Printf("miss var decl %#+v\n", t)
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
				if fn, ok := w(FuncSpecFnFuncDeclSpec).(FnFuncDeclSpec); ok {
					fn(d)
				}
			} else if d.Recv.NumFields() != 1 {
				panic(fmt.Errorf("more than one reciver: %#+v", d))
			} else {
				r := d.Recv.List[0]
				if fn, ok := w(FuncSpecFnMethodDeclSpec).(FnMethodDeclSpec); ok {
					if fn(d, r) {
						switch t := r.Type.(type) {
						case *ast.StarExpr:
							if fn, ok := w(FuncSpecFnMethodDeclStarRecvSpec).(FnMethodDeclStarRecvSpec); ok {
								fn(d, r, t)
							}
						case *ast.Ident:
							if fn, ok := w(FuncSpecFnMethodDeclIdentRecvSpec).(FnMethodDeclIdentRecvSpec); ok {
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
							if fn, ok := w(FuncSpecFnIdentTypeSpec).(FnIdentTypeSpec); ok {
								fn(d, s, t)
							}
						case *ast.StructType:
							if fn, ok := w(FuncSpecFnStructTypeSpec).(FnStructTypeSpec); ok {
								fn(d, s, t)
							}
						case *ast.InterfaceType:
							if fn, ok := w(FuncSpecFnInterfaceTypeSpec).(FnInterfaceTypeSpec); ok {
								fn(d, s, t)
							}
						case *ast.MapType:
							if fn, ok := w(FuncSpecFnMapTypeSpec).(FnMapTypeSpec); ok {
								fn(d, s, t)
							}
						case *ast.FuncType:
							if fn, ok := w(FuncSpecFnFuncTypeSpec).(FnFuncTypeSpec); ok {
								fn(d, s, t)
							}
						case *ast.ArrayType:
							if fn, ok := w(FuncSpecFnArrayTypeSpec).(FnArrayTypeSpec); ok {
								fn(d, s, t)
							}
						case *ast.ChanType:
							if fn, ok := w(FuncSpecFnChanTypeSpec).(FnChanTypeSpec); ok {
								fn(d, s, t)
							}
						case *ast.SelectorExpr:
							if fn, ok := w(FuncSpecFnSelectorTypeSpec).(FnSelectorTypeSpec); ok {
								fn(d, s, t)
							}
						case *ast.StarExpr:
							if fn, ok := w(FuncSpecFnStarExprSpec).(FnStarExprSpec); ok {
								fn(d, s, t)
							}

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
func CaseTypeFunc(pkg *packages.Package, typ ast.Expr, w FnProvider) {
	switch t := typ.(type) {
	case *ast.Ident:
		if fn, ok := w(FuncTypeFnIdentType).(FnIdentType); ok {
			fn(t)
		}
	case *ast.SelectorExpr:
		if fn, ok := w(FuncTypeFnSelectorType).(FnSelectorType); ok {
			fn(t, pkg.TypesInfo.Uses[t.Sel])
		}
	case *ast.StarExpr:
		if fn, ok := w(FuncTypeFnStarType).(FnStarType); ok {
			fn(t)
		}
	case *ast.FuncType:
		if fn, ok := w(FuncTypeFnFuncType).(FnFuncType); ok {
			fn(t)
		}
	case *ast.ArrayType:
		if fn, ok := w(FuncTypeFnArrayType).(FnArrayType); ok {
			fn(t)
		}
	case *ast.MapType:
		if fn, ok := w(FuncTypeFnMapType).(FnMapType); ok {
			fn(t)
		}
	case *ast.StructType:
		if fn, ok := w(FuncTypeFnStructType).(FnStructType); ok {
			fn(t)
		}
	case *ast.ChanType:
		if fn, ok := w(FuncTypeFnChanType).(FnChanType); ok {
			fn(t)
		}
	case *ast.Ellipsis:
		if fn, ok := w(FuncTypeFnEllipsis).(FnEllipsis); ok {
			fn(t)
		}
	case *ast.InterfaceType:
		if fn, ok := w(FuncTypeFnInterfaceType).(FnInterfaceType); ok {
			fn(t)
		}
	default:
		log.Printf("\nmissing type %#+v\n", t)
	}
}

type Mode int

const (
	TypeDeclMode Mode = 1 << iota
	VariableDeclMode
	ConstantDeclMode
	AllDeclMode = TypeDeclMode | VariableDeclMode | ConstantDeclMode
)

func CaseDeclsFunc(file *ast.File, mode Mode, w FnProvider) {
	consts := mode&ConstantDeclMode > 0
	variables := mode&VariableDeclMode > 0
	types := mode&TypeDeclMode > 0
	if !consts && !variables && !types {
		return
	}
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if types {
				if d.Recv == nil {
					if fn, ok := w(FuncSpecFnFuncDeclSpec).(FnFuncDeclSpec); ok {
						fn(d)
					}
				} else if d.Recv.NumFields() != 1 {
					panic(fmt.Errorf("more than one reciver: %#+v", d))
				} else {
					r := d.Recv.List[0]
					if fn, ok := w(FuncSpecFnMethodDeclSpec).(FnMethodDeclSpec); ok {
						if fn(d, r) {
							switch t := r.Type.(type) {
							case *ast.StarExpr:
								if fn, ok := w(FuncSpecFnMethodDeclStarRecvSpec).(FnMethodDeclStarRecvSpec); ok {
									fn(d, r, t)
								}
							case *ast.Ident:
								if fn, ok := w(FuncSpecFnMethodDeclIdentRecvSpec).(FnMethodDeclIdentRecvSpec); ok {
									fn(d, r, t)
								}
							default:
								log.Printf("miss method reciver type: %#+v \n", t)
							}
						}
					}
				}
			}
		case *ast.GenDecl:
			switch d.Tok {
			case token.TYPE:
				if types {
					for _, spec := range d.Specs {
						if s, ok := spec.(*ast.TypeSpec); ok {
							switch t := s.Type.(type) {
							case *ast.Ident:
								if fn, ok := w(FuncSpecFnIdentTypeSpec).(FnIdentTypeSpec); ok {
									fn(d, s, t)
								}
							case *ast.StructType:
								if fn, ok := w(FuncSpecFnStructTypeSpec).(FnStructTypeSpec); ok {
									fn(d, s, t)
								}
							case *ast.InterfaceType:
								if fn, ok := w(FuncSpecFnInterfaceTypeSpec).(FnInterfaceTypeSpec); ok {
									fn(d, s, t)
								}
							case *ast.MapType:
								if fn, ok := w(FuncSpecFnMapTypeSpec).(FnMapTypeSpec); ok {
									fn(d, s, t)
								}
							case *ast.FuncType:
								if fn, ok := w(FuncSpecFnFuncTypeSpec).(FnFuncTypeSpec); ok {
									fn(d, s, t)
								}
							case *ast.ArrayType:
								if fn, ok := w(FuncSpecFnArrayTypeSpec).(FnArrayTypeSpec); ok {
									fn(d, s, t)
								}
							case *ast.ChanType:
								if fn, ok := w(FuncSpecFnChanTypeSpec).(FnChanTypeSpec); ok {
									fn(d, s, t)
								}
							case *ast.SelectorExpr:
								if fn, ok := w(FuncSpecFnSelectorTypeSpec).(FnSelectorTypeSpec); ok {
									fn(d, s, t)
								}
							case *ast.StarExpr:
								if fn, ok := w(FuncSpecFnStarExprSpec).(FnStarExprSpec); ok {
									fn(d, s, t)
								}

							default:
								log.Printf("miss decl %#+v\n", t)
							}
						}
					}
				}
			case token.VAR:
				if variables {
					for _, spec := range d.Specs {
						if s, ok := spec.(*ast.ValueSpec); ok {
							switch t := s.Type.(type) {
							case *ast.Ident:
								if fn, ok := w(FuncVarsFnIdentVarSpec).(FnIdentVarSpec); ok {
									fn(d, s, t)
								}
							case *ast.FuncType:
								if fn, ok := w(FuncVarsFnFuncVarSpec).(FnFuncVarSpec); ok {
									fn(d, s, t)
								}
							case *ast.SelectorExpr:
								if fn, ok := w(FuncVarsFnSelectorVarSpec).(FnSelectorVarSpec); ok {
									fn(d, s, t)
								}
							case *ast.MapType:
								if fn, ok := w(FuncVarsFnMapVarSpec).(FnMapVarSpec); ok {
									fn(d, s, t)
								}
							case *ast.ArrayType:
								if fn, ok := w(FuncVarsFnArrayVarSpec).(FnArrayVarSpec); ok {
									fn(d, s, t)
								}
							case *ast.InterfaceType:
								if fn, ok := w(FuncVarsFnInterfaceVarSpec).(FnInterfaceVarSpec); ok {
									fn(d, s, t)
								}
							case *ast.StarExpr:
								if fn, ok := w(FuncVarsFnStarVarSpec).(FnStarVarSpec); ok {
									fn(d, s, t)
								}
							case nil:
								if fn, ok := w(FuncVarsFnNilVarSpec).(FnNilVarSpec); ok {
									fn(d, s)
								}
							default:
								log.Printf("miss var decl %#+v\n", t)
							}
						}
					}
				}
			case token.CONST:
				if consts {
					for _, spec := range d.Specs {
						if s, ok := spec.(*ast.ValueSpec); ok {
							switch t := s.Type.(type) {
							case *ast.Ident:
								if fn, ok := w(FuncConsFnIdentConstSpec).(FnIdentConstSpec); ok {
									fn(d, s, t, ok)
								}
							case nil:
								if fn, ok := w(FuncConsFnNilConstSpec).(FnNilConstSpec); ok {
									fn(d, s, ok)
								}
							default:
								log.Printf("miss const decl %#+v\n", t)
							}
						}
					}
				}
			default:
				continue
			}
		}
	}
}
