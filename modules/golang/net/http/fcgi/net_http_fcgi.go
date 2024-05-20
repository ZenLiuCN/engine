// Code generated by define_gene; DO NOT EDIT.
package fcgi

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/net"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/http"
	"net/http/fcgi"
)

var (
	//go:embed net_http_fcgi.d.ts
	NetHttpFcgiDefine   []byte
	NetHttpFcgiDeclared = map[string]any{
		"processEnv":        fcgi.ProcessEnv,
		"serve":             fcgi.Serve,
		"ErrConnClosed":     fcgi.ErrConnClosed,
		"ErrRequestAborted": fcgi.ErrRequestAborted,
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