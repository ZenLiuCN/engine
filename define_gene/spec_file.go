package main

import "go/ast"

type SpecFile struct {
	SpecPkg *SpecPackage
	AstFile *ast.File
}
