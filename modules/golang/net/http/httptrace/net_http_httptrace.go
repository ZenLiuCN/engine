// Code generated by define_gene; DO NOT EDIT.
package httptrace

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/golang/crypto/tls"
	_ "github.com/ZenLiuCN/engine/modules/golang/net"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/textproto"
	_ "github.com/ZenLiuCN/engine/modules/golang/time"
	"net/http/httptrace"
)

var (
	//go:embed net_http_httptrace.d.ts
	NetHttpHttptraceDefine   []byte
	NetHttpHttptraceDeclared = map[string]any{
		"withClientTrace":    httptrace.WithClientTrace,
		"contextClientTrace": httptrace.ContextClientTrace,

		"emptyWroteRequestInfo":    engine.Empty[httptrace.WroteRequestInfo],
		"emptyRefWroteRequestInfo": engine.EmptyRefer[httptrace.WroteRequestInfo],
		"refOfWroteRequestInfo":    engine.ReferOf[httptrace.WroteRequestInfo],
		"unRefWroteRequestInfo":    engine.UnRefer[httptrace.WroteRequestInfo],
		"emptyClientTrace":         engine.Empty[httptrace.ClientTrace],
		"emptyRefClientTrace":      engine.EmptyRefer[httptrace.ClientTrace],
		"refOfClientTrace":         engine.ReferOf[httptrace.ClientTrace],
		"unRefClientTrace":         engine.UnRefer[httptrace.ClientTrace],
		"emptyDNSDoneInfo":         engine.Empty[httptrace.DNSDoneInfo],
		"emptyRefDNSDoneInfo":      engine.EmptyRefer[httptrace.DNSDoneInfo],
		"refOfDNSDoneInfo":         engine.ReferOf[httptrace.DNSDoneInfo],
		"unRefDNSDoneInfo":         engine.UnRefer[httptrace.DNSDoneInfo],
		"emptyDNSStartInfo":        engine.Empty[httptrace.DNSStartInfo],
		"emptyRefDNSStartInfo":     engine.EmptyRefer[httptrace.DNSStartInfo],
		"refOfDNSStartInfo":        engine.ReferOf[httptrace.DNSStartInfo],
		"unRefDNSStartInfo":        engine.UnRefer[httptrace.DNSStartInfo],
		"emptyGotConnInfo":         engine.Empty[httptrace.GotConnInfo],
		"emptyRefGotConnInfo":      engine.EmptyRefer[httptrace.GotConnInfo],
		"refOfGotConnInfo":         engine.ReferOf[httptrace.GotConnInfo],
		"unRefGotConnInfo":         engine.UnRefer[httptrace.GotConnInfo]}
)

func init() {
	engine.RegisterModule(NetHttpHttptraceModule{})
}

type NetHttpHttptraceModule struct{}

func (S NetHttpHttptraceModule) Identity() string {
	return "golang/net/http/httptrace"
}
func (S NetHttpHttptraceModule) TypeDefine() []byte {
	return NetHttpHttptraceDefine
}
func (S NetHttpHttptraceModule) Exports() map[string]any {
	return NetHttpHttptraceDeclared
}
