// Code generated by define_gene; DO NOT EDIT.
package httptest

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/golang/bytes"
	_ "github.com/ZenLiuCN/engine/golang/crypto/tls"
	_ "github.com/ZenLiuCN/engine/golang/crypto/x509"
	_ "github.com/ZenLiuCN/engine/golang/io"
	_ "github.com/ZenLiuCN/engine/golang/net"
	_ "github.com/ZenLiuCN/engine/golang/net/http"
	"net/http/httptest"
)

var (
	//go:embed net_http_httptest.d.ts
	NetHttpHttptestDefine   []byte
	NetHttpHttptestDeclared = map[string]any{
		"DefaultRemoteAddr":  httptest.DefaultRemoteAddr,
		"newRecorder":        httptest.NewRecorder,
		"newRequest":         httptest.NewRequest,
		"newServer":          httptest.NewServer,
		"newTLSServer":       httptest.NewTLSServer,
		"newUnstartedServer": httptest.NewUnstartedServer,

		"emptyResponseRecorder": func() (v httptest.ResponseRecorder) {
			return v
		},
		"refResponseRecorder": func() *httptest.ResponseRecorder {
			var x httptest.ResponseRecorder
			return &x
		},
		"refOfResponseRecorder": func(x httptest.ResponseRecorder) *httptest.ResponseRecorder {
			return &x
		},
		"emptyServer": func() (v httptest.Server) {
			return v
		},
		"refServer": func() *httptest.Server {
			var x httptest.Server
			return &x
		},
		"refOfServer": func(x httptest.Server) *httptest.Server {
			return &x
		}}
)

func init() {
	engine.RegisterModule(NetHttpHttptestModule{})
}

type NetHttpHttptestModule struct{}

func (S NetHttpHttptestModule) Identity() string {
	return "golang/net/http/httptest"
}
func (S NetHttpHttptestModule) TypeDefine() []byte {
	return NetHttpHttptestDefine
}
func (S NetHttpHttptestModule) Exports() map[string]any {
	return NetHttpHttptestDeclared
}
