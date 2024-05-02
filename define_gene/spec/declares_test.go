package spec

import (
	"fmt"
	"go/ast"
	"log"
	"path/filepath"
	"testing"
)

func TestVar(t *testing.T) {
	gr := "D:\\Dev\\env\\golang\\go.1.21x64" // runtime.GOROOT()
	root := filepath.ToSlash(gr)
	src := filepath.Join(root, "src")
	pkg := "net/http"
	source := filepath.Join(src, pkg)
	p := ParseTypeInfo(nil, []string{source}, IdentityLoadMode, log.Printf)
	pc := p[0]
	for _, syntax := range pc.Syntax {
		CaseVarDecl(syntax, varPrinter)
	}
}

type VarPrinter struct {
}

func (v VarPrinter) SelectorVarSpec(s *ast.ValueSpec, t *ast.SelectorExpr) {
	fmt.Println(s, t)
}

func (v VarPrinter) MapVarSpec(s *ast.ValueSpec, t *ast.MapType) {
	fmt.Println(s, t)
}

func (v VarPrinter) StarVarSpec(s *ast.ValueSpec, t *ast.StarExpr) {
	fmt.Println(s, t)
}

func (v VarPrinter) NilVarSpec(s *ast.ValueSpec) {
	fmt.Println(s)
}

func (v VarPrinter) FuncVarSpec(s *ast.ValueSpec, t *ast.FuncType) {
	fmt.Println(s, t)
}

func (v VarPrinter) IdentVarSpec(s *ast.ValueSpec, t *ast.Ident) {
	fmt.Println(s, t)
}

var varPrinter = VarPrinter{}

func TestConst(t *testing.T) {
	gr := "D:\\Dev\\env\\golang\\go.1.21x64" // runtime.GOROOT()
	root := filepath.ToSlash(gr)
	src := filepath.Join(root, "src")
	pkg := "net/http"
	source := filepath.Join(src, pkg)
	p := ParseTypeInfo(nil, []string{source}, IdentityLoadMode, log.Printf)
	pc := p[0]
	for _, syntax := range pc.Syntax {
		CaseConstDecl(syntax, constPrinter)
	}
}

type ConstPrinter struct {
}

func (c ConstPrinter) NilConstSpec(s *ast.ValueSpec) {
	fmt.Println(s)
}

func (c ConstPrinter) IdentConstSpec(s *ast.ValueSpec, t *ast.Ident) {
	fmt.Println(s, t)
}

var constPrinter = ConstPrinter{}
