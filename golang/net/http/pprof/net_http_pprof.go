// Code generated by define_gene; DO NOT EDIT.
package pprof

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/golang/net/http"
	"net/http/pprof"
)

var (
	//go:embed net_http_pprof.d.ts
	NetHttpPprofDefine   []byte
	NetHttpPprofDeclared = map[string]any{
		"cmdline": pprof.Cmdline,
		"handler": pprof.Handler,
		"index":   pprof.Index,
		"profile": pprof.Profile,
		"symbol":  pprof.Symbol,
		"trace":   pprof.Trace,
	}
)

func init() {
	engine.RegisterModule(NetHttpPprofModule{})
}

type NetHttpPprofModule struct{}

func (S NetHttpPprofModule) Identity() string {
	return "golang/net/http/pprof"
}
func (S NetHttpPprofModule) TypeDefine() []byte {
	return NetHttpPprofDefine
}
func (S NetHttpPprofModule) Exports() map[string]any {
	return NetHttpPprofDeclared
}
