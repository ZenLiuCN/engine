// Code generated by define_gene; DO NOT EDIT.
package httptest

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/bytes"
	_ "github.com/ZenLiuCN/engine/modules/golang/crypto/tls"
	_ "github.com/ZenLiuCN/engine/modules/golang/crypto/x509"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	_ "github.com/ZenLiuCN/engine/modules/golang/net"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/http"
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

		"emptyResponseRecorder":    engine.Empty[httptest.ResponseRecorder],
		"emptyRefResponseRecorder": engine.EmptyRefer[httptest.ResponseRecorder],
		"refOfResponseRecorder":    engine.ReferOf[httptest.ResponseRecorder],
		"unRefResponseRecorder":    engine.UnRefer[httptest.ResponseRecorder]}
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