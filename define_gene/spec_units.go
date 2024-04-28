package main

import (
	"fmt"
	"github.com/ZenLiuCN/fn"
	"go/ast"
	"go/token"
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
