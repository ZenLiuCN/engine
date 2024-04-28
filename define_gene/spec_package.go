package main

import "golang.org/x/tools/go/packages"

type SpecPackage struct {
	*SpecContext
	Name      string
	Package   *packages.Package
	SpecFiles []*SpecFile
}

func (p *SpecPackage) ResolveFiles() {
	p.SpecFiles = make([]*SpecFile, len(p.Package.Syntax))
	for i, f := range p.Package.Syntax {
		p.SpecFiles[i] = &SpecFile{
			SpecPkg: p,
			AstFile: f,
		}
	}
}
