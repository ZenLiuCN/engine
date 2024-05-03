package main

//https://cs.opensource.google/go/x/tools/+/master:cmd/stringer/stringer.go;drc=daf94608b5e2caf763ba634b84e7a5ba7970e155;l=382
import (
	"fmt"
	"github.com/ZenLiuCN/engine/define_gene/spec"
	"strings"
)

type Generator struct {
	dir   string
	tags  []string
	files []string
	types []string
	out   string
	log   func(format string, a ...any)
	over  bool
}

//go:generate stringer -type=Kind
type Kind int

const (
	StructType Kind = iota
	MapType
	InterfaceType
	ArrayType
	ChanType
	FuncType
	FuncDecl
	IdentType
	StarMethodType
	IdentMethodType
	SelectorType
	StarType
	ConstValue
)

func (g *Generator) generate() {
	debug = g.log != nil
	var flags []string
	if len(g.tags) > 0 {
		flags = append(flags, fmt.Sprintf("-tags=%s", strings.Join(g.tags, " ")))
	}
	p := spec.ParseTypeInfo(flags, g.files, spec.IdentityLoadMode, g.log)
	if len(p) > 1 {
		panic("only one package each time")
	}
	pkg := p[0]
	if pkg.Name == "" { //!! skip without any package
		return
	}
	fw := new(FileWriter)
	fw.debug = g.log != nil
	fw.out = g.out
	fw.override = g.over
	fw.With(pkg, "")
	fw.Process()
	fw.Flush()
	fw.Free()
}
