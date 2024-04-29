package main

import (
	"fmt"
	"github.com/ZenLiuCN/fn"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"strings"
)

func IsDir(name string) bool {
	if info, err := os.Stat(name); err != nil {
		log.Fatal(err)
	} else {
		return info.IsDir()
	}
	return false
}
func Parse(tags, files []string, f func(format string, args ...any)) []*packages.Package {
	return fn.Panic1(packages.Load(&packages.Config{
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
		Tests:      false,
		Logf:       f,
	}, files...))
}

type (
	DeclaredWalk interface {
		FuncDecl(decl *ast.FuncDecl)
		MethodDecl(decl *ast.FuncDecl)
		IdentTypeSpec(spec *ast.TypeSpec)
		StructTypeSpec(spec *ast.TypeSpec)
		InterfaceTypeSpec(spec *ast.TypeSpec)
		MapTypeSpec(spec *ast.TypeSpec)
		ArrayTypeSpec(spec *ast.TypeSpec)
		FuncTypeSpec(spec *ast.TypeSpec)
	}
)

func DeclsWalk(file *ast.File, w DeclaredWalk) {
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Recv == nil {
				w.FuncDecl(d)
			} else if d.Recv.NumFields() != 1 {
				panic(fmt.Errorf("more than one reciver: %#+v", d))
			} else {
				w.MethodDecl(d)
			}
		case *ast.GenDecl:
			switch d.Tok {
			case token.TYPE:
				for _, spec := range d.Specs {
					if s, ok := spec.(*ast.TypeSpec); ok {
						switch t := s.Type.(type) {
						case *ast.Ident:
							w.IdentTypeSpec(s)
						case *ast.StructType:
							w.StructTypeSpec(s)
						case *ast.InterfaceType:
							w.InterfaceTypeSpec(s)
						case *ast.MapType:
							w.MapTypeSpec(s)
						case *ast.FuncType:
							w.FuncTypeSpec(s)
						case *ast.ArrayType:
							w.ArrayTypeSpec(s)
						default:
							fmt.Printf("miss %#+v\n", t)
						}
					}
				}
			default:
				continue
			}
		}
	}
}

//go:generate go install golang.org/x/tools/cmd/stringer
//go:generate stringer -type=TypeKind,FieldKind
type TypeKind int

const (
	TypeKindStruct TypeKind = iota + 1
	TypeKindArray
	TypeKindInterface
	TypeKindMap
	TypeKindFunc
	TypeKindIdent
	TypeKindSelector
	TypeKindStar
	TypeKindChan
	TypeKindChanSend
	TypeKindChanRecv
)

func LookupType(pkg *packages.Package, file *ast.File, typ ast.Expr) (kind []TypeKind, ty types.Object, pk *types.Package) {
	switch t := typ.(type) {
	case *ast.Ident:
		kind = append(kind, TypeKindIdent)
		return
	case *ast.SelectorExpr:
		var ok bool
		ty, ok = pkg.TypesInfo.Uses[t.Sel]
		if ok {
			pk = ty.Pkg()
			kind = append(kind, TypeKindSelector)
		} else {
			fmt.Printf("\nmissing selector %#+v\n", t)
		}
		return
	case *ast.StarExpr:
		return LookupTypeInner(pkg, file, t.X, append(kind, TypeKindStar))
	case *ast.FuncType:
		kind = append(kind, TypeKindFunc)
		return
	case *ast.ArrayType:
		kind = append(kind, TypeKindArray)
		kind, ty, pk = LookupTypeInner(pkg, file, t.Elt, kind)
		return
	case *ast.MapType:
		kind = append(kind, TypeKindMap)
		kind, ty, pk = LookupTypeInner(pkg, file, t.Key, kind)
		kind, ty, pk = LookupTypeInner(pkg, file, t.Value, kind)
		return
	case *ast.StructType:
		kind = append(kind, TypeKindStruct)
	case *ast.ChanType:
		switch {
		case ast.RECV == t.Dir:
			return LookupTypeInner(pkg, file, t.Value, append(kind, TypeKindChanRecv))
		case ast.SEND == t.Dir:
			return LookupTypeInner(pkg, file, t.Value, append(kind, TypeKindChanSend))
		default:
			return LookupTypeInner(pkg, file, t.Value, append(kind, TypeKindChan))
		}
	default:
		fmt.Printf("\nmissing type %#+v\n", t)
	}
	return
}
func LookupTypeInner(pkg *packages.Package, file *ast.File, typ ast.Expr, k []TypeKind) (kind []TypeKind, ty types.Object, pk *types.Package) {
	switch t := typ.(type) {
	case *ast.Ident:
		kind = append(k, TypeKindIdent)
		return
	case *ast.SelectorExpr:
		var ok bool
		ty, ok = pkg.TypesInfo.Uses[t.Sel]
		if ok {
			pk = ty.Pkg()
			kind = append(k, TypeKindSelector)
		} else {
			fmt.Printf("\nmissing selector %#+v\n", t)
		}
		return
	case *ast.StarExpr:
		return LookupTypeInner(pkg, file, t.X, append(k, TypeKindStar))
	case *ast.FuncType:
		kind = append(k, TypeKindFunc)
		return
	case *ast.ArrayType:
		kind = append(k, TypeKindArray)
		kind, ty, pk = LookupTypeInner(pkg, file, t.Elt, kind)
		return
	case *ast.MapType:
		kind = append(k, TypeKindMap)
		kind, ty, pk = LookupTypeInner(pkg, file, t.Key, kind)
		kind, ty, pk = LookupTypeInner(pkg, file, t.Value, kind)
		return
	case *ast.StructType:
		kind = append(k, TypeKindStruct)
	case *ast.ChanType:
		switch {
		case ast.RECV == t.Dir:
			return LookupTypeInner(pkg, file, t.Value, append(kind, TypeKindChanRecv))
		case ast.SEND == t.Dir:
			return LookupTypeInner(pkg, file, t.Value, append(kind, TypeKindChanSend))
		default:
			return LookupTypeInner(pkg, file, t.Value, append(kind, TypeKindChan))
		}
	default:
		fmt.Printf("\nmissing type %#+v\n", t)
	}
	return
}
