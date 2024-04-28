package main

import (
	"fmt"
	"github.com/ZenLiuCN/fn"
	"golang.org/x/tools/go/packages"
	"strings"
)

type SpecContext struct {
	SpecPackages []*SpecPackage
	Logf         func(format string, args ...any)
}

func NewContext(logf func(format string, args ...any)) *SpecContext {
	return &SpecContext{Logf: logf}
}

func (c *SpecContext) Printf(format string, args ...any) {
	if c.Logf != nil {
		c.Logf(format, args...)
	}
}
func (c *SpecContext) Parse(tags []string, files []string) {
	pkg := fn.Panic1(packages.Load(&packages.Config{
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
		Tests:      false,
		Logf:       c.Logf,
	}, files...))
	c.SpecPackages = make([]*SpecPackage, len(pkg))
	for i, p := range pkg {
		c.SpecPackages[i] = &SpecPackage{
			SpecContext: c,
			Name:        p.Name,
			Package:     p,
		}
		c.SpecPackages[i].ResolveFiles()
	}
}
