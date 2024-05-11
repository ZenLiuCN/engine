// Code generated by define_gene; DO NOT EDIT.
package fcgi

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/golang/net"
	_ "github.com/ZenLiuCN/engine/golang/net/http"
	"net/http/fcgi"
)

var (
	//go:embed net_http_fcgi.d.ts
	NetHttpFcgiDefine   []byte
	NetHttpFcgiDeclared = map[string]any{
		"ErrConnClosed":     fcgi.ErrConnClosed,
		"ErrRequestAborted": fcgi.ErrRequestAborted,
		"processEnv":        fcgi.ProcessEnv,
		"serve":             fcgi.Serve,
	}
)

func init() {
	engine.RegisterModule(NetHttpFcgiModule{})
}

type NetHttpFcgiModule struct{}

func (S NetHttpFcgiModule) Identity() string {
	return "golang/net/http/fcgi"
}
func (S NetHttpFcgiModule) TypeDefine() []byte {
	return NetHttpFcgiDefine
}
func (S NetHttpFcgiModule) Exports() map[string]any {
	return NetHttpFcgiDeclared
}
