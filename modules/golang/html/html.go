// Code generated by define_gene; DO NOT EDIT.
package html

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"html"
)

var (
	//go:embed html.d.ts
	HtmlDefine   []byte
	HtmlDeclared = map[string]any{
		"escapeString":   html.EscapeString,
		"unescapeString": html.UnescapeString,
	}
)

func init() {
	engine.RegisterModule(HtmlModule{})
}

type HtmlModule struct{}

func (S HtmlModule) Identity() string {
	return "golang/html"
}
func (S HtmlModule) TypeDefine() []byte {
	return HtmlDefine
}
func (S HtmlModule) Exports() map[string]any {
	return HtmlDeclared
}
