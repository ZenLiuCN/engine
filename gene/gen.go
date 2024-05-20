package main

import (
	"fmt"
	. "github.com/ZenLiuCN/go-inspect"
	"golang.org/x/tools/go/packages"
	"os"
	"strings"
)

type InspectorX = *TypeInspector[*context]
type Generator struct {
	dir      string
	tags     []string
	files    []string
	out      string
	suffix   string
	log      func(format string, a ...any)
	over     bool
	overTest bool
	print    bool
	trace    bool
	env      []string
	errors   bool
	ignores  []string
}

func (g *Generator) generate() (err error) {
	debug = g.log != nil
	if g.trace {
		Debug = true
		Trace = true
	}
	var flags []string
	if len(g.tags) > 0 {
		flags = append(flags, fmt.Sprintf("-tags=%s", strings.Join(g.tags, " ")))
	}
	c := new(context)
	c.ignores = g.ignores
	c.errors = g.errors
	c.init()
	ip := NewTypeInspector[*context](false)
	ip.Inspector.PopEach = true
	ip.Inspector.Visitor = c
	ip.Inspector.X = c
	err = ip.Inspect(
		func(config *packages.Config) {
			config.BuildFlags = flags
			if len(g.env) > 0 {
				config.Env = append(os.Environ(), g.env...)
			}

			//config.Mode |= packages.NeedImports
		},
		func(i []*packages.Package) error {
			if len(i) != 1 {
				return fmt.Errorf("only one package each run")
			}
			c.pkg = i[0]
			return nil
		}, g.files...)
	if err != nil {
		return err
	}

	return c.flush(g)
}
