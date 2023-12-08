package pug

import (
	_ "embed"
	"github.com/Joker/jade"
	"github.com/ZenLiuCN/engine"
	"html/template"
)

var (
	//go:embed pug.d.ts
	pugDefine []byte
	pugMap    = map[string]any{
		"parse": jade.Parse,
		"parseText": func(name, code string) (string, error) {
			return jade.Parse(name, []byte(code))
		},
		"parseFile": jade.ParseFile,
		"config":    jade.Config,
		"template": func(name, code string) (*template.Template, error) {
			return template.New(name).Parse(code)
		},
	}
)

func init() {
	engine.RegisterModule(&Pug{})
}

type Pug struct {
	m map[string]any
}

func (p *Pug) Identity() string {
	return "go/pug"
}

func (p *Pug) TypeDefine() []byte {
	return pugDefine
}

func (p *Pug) Exports() map[string]any {

	return pugMap
}
