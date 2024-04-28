package main

import (
	"fmt"
	"go/ast"
	"go/token"
)

//go:generate stringer -type TypeKind,FieldKind
type TypeKind int

const (
	KindStruct TypeKind = iota + 1
	KindArray
	KindInterface
	KindMap
	KindFunc
	KindIdent
)

type SpecType struct {
	Kind        TypeKind
	SpecFile    *SpecFile
	AstType     ast.Expr
	AstTypeSpec *ast.TypeSpec
}

func (s *SpecType) Resolve() *SpecResolveType {
	switch s.Kind {
	case KindStruct:
		return resolveStructType(s, s.AstType.(*ast.StructType))
	case KindArray:
		return resolveArrayType(s, s.AstType.(*ast.ArrayType))
	case KindInterface:
		return resolveInterfaceType(s, s.AstType.(*ast.InterfaceType))
	case KindMap:
		return resolveMapType(s, s.AstType.(*ast.MapType))
	case KindFunc:
		return resolveFuncType(s, s.AstType.(*ast.FuncType))
	case KindIdent:
		return resolveIdentType(s, s.AstType.(*ast.Ident))
	default:
		panic(fmt.Errorf("kind %s of %#+v unknown", s.Kind, s))
	}
}

func resolveIdentType(s *SpecType, t *ast.Ident) *SpecResolveType {
	return &SpecResolveType{
		SpecType: s,
		Name:     t.Name,
	}
}

func resolveFuncType(s *SpecType, t *ast.FuncType) *SpecResolveType {
	return nil
}

func resolveMapType(s *SpecType, t *ast.MapType) *SpecResolveType {
	return nil
}

func resolveInterfaceType(s *SpecType, t *ast.InterfaceType) *SpecResolveType {
	return nil
}

func resolveArrayType(s *SpecType, t *ast.ArrayType) *SpecResolveType {

	return nil
}

func resolveStructType(s *SpecType, t *ast.StructType) (r *SpecResolveType) {
	r = new(SpecResolveType)
	r.SpecType = s
	r.Name = s.AstTypeSpec.Name.Name
	r.Fields = make([]*SpecField, t.Fields.NumFields())
	for i, field := range t.Fields.List {
		r.Fields[i] = resolveField(s, t, field)
	}
	return r
}

func FindTypeSpec(f *SpecFile, export bool, on func(*SpecType)) func(node ast.Node) bool {
	return func(node ast.Node) bool {
		dec, ok := node.(*ast.GenDecl)
		if !ok || dec.Tok != token.TYPE {
			return true
		}
		for _, spec := range dec.Specs {
			if ts, ok := spec.(*ast.TypeSpec); ok {
				if export && !ast.IsExported(ts.Name.Name) {
					continue
				}
				switch t := ts.Type.(type) {
				case *ast.StructType:
					on(&SpecType{
						Kind:        KindStruct,
						SpecFile:    f,
						AstType:     t,
						AstTypeSpec: ts,
					})
				case *ast.ArrayType:
					on(&SpecType{
						Kind:        KindArray,
						SpecFile:    f,
						AstType:     t,
						AstTypeSpec: ts,
					})
				case *ast.InterfaceType:
					on(&SpecType{
						Kind:        KindInterface,
						SpecFile:    f,
						AstType:     t,
						AstTypeSpec: ts,
					})
				case *ast.MapType:
					on(&SpecType{
						Kind:        KindMap,
						SpecFile:    f,
						AstType:     t,
						AstTypeSpec: ts,
					})
				case *ast.FuncType:
					on(&SpecType{
						Kind:        KindFunc,
						SpecFile:    f,
						AstType:     t,
						AstTypeSpec: ts,
					})
				case *ast.Ident:
					on(&SpecType{
						Kind:        KindIdent,
						SpecFile:    f,
						AstType:     t,
						AstTypeSpec: ts,
					})
				default:
					fmt.Printf("%[1]T %+[1]v\n", t)
				}
			}
		}
		return true
	}
}

type SpecResolveType struct {
	SpecType *SpecType
	Name     string

	Fields []*SpecField
}
type FieldKind int

const (
	FieldKindIdent FieldKind = iota + 1
	FieldKindStar
)

type SpecField struct {
	Embedded bool
	Kind     FieldKind
	AstType  ast.Expr
	Refer    *SpecReferType
	TypeName string
}

func resolveField(s *SpecType, t *ast.StructType, field *ast.Field) (r *SpecField) {
	r = new(SpecField)
	r.Embedded = len(field.Names) == 0
	r.AstType = t
	switch x := field.Type.(type) {
	case *ast.StarExpr:
		r.Kind = FieldKindStar
		r.Refer = ResolveType(s.SpecFile, x.X)
	case *ast.Ident:
		r.Kind = FieldKindIdent
		r.TypeName = x.Name
	default:
		fmt.Printf("field type %#+v\n", x)
	}
	return
}

type SpecReferType struct {
	SpecFile *SpecFile
	Select   ast.Expr
	Name     string
}

func ResolveType(f *SpecFile, t ast.Expr) (r *SpecReferType) {
	r = new(SpecReferType)
	r.SpecFile = f
	switch x := t.(type) {
	case *ast.SelectorExpr:
		r.Select = x.X
		r.Name = x.Sel.Name
	default:
		fmt.Printf("refer type %#+v\n", x)
	}
	return
}
