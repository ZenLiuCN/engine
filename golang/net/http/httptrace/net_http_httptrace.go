// Code generated by define_gene; DO NOT EDIT.
package httptrace

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/golang/context"
	_ "github.com/ZenLiuCN/engine/golang/crypto/tls"
	_ "github.com/ZenLiuCN/engine/golang/net"
	_ "github.com/ZenLiuCN/engine/golang/net/textproto"
	_ "github.com/ZenLiuCN/engine/golang/time"
	"net/http/httptrace"
)

var (
	//go:embed net_http_httptrace.d.ts
	NetHttpHttptraceDefine   []byte
	NetHttpHttptraceDeclared = map[string]any{
		"contextClientTrace": httptrace.ContextClientTrace,
		"withClientTrace":    httptrace.WithClientTrace,

		"emptyWroteRequestInfo": func() (v httptrace.WroteRequestInfo) {
			return v
		},
		"refWroteRequestInfo": func() *httptrace.WroteRequestInfo {
			var x httptrace.WroteRequestInfo
			return &x
		},
		"refOfWroteRequestInfo": func(x httptrace.WroteRequestInfo) *httptrace.WroteRequestInfo {
			return &x
		},
		"emptyClientTrace": func() (v httptrace.ClientTrace) {
			return v
		},
		"refClientTrace": func() *httptrace.ClientTrace {
			var x httptrace.ClientTrace
			return &x
		},
		"refOfClientTrace": func(x httptrace.ClientTrace) *httptrace.ClientTrace {
			return &x
		},
		"emptyDNSDoneInfo": func() (v httptrace.DNSDoneInfo) {
			return v
		},
		"refDNSDoneInfo": func() *httptrace.DNSDoneInfo {
			var x httptrace.DNSDoneInfo
			return &x
		},
		"refOfDNSDoneInfo": func(x httptrace.DNSDoneInfo) *httptrace.DNSDoneInfo {
			return &x
		},
		"emptyDNSStartInfo": func() (v httptrace.DNSStartInfo) {
			return v
		},
		"refDNSStartInfo": func() *httptrace.DNSStartInfo {
			var x httptrace.DNSStartInfo
			return &x
		},
		"refOfDNSStartInfo": func(x httptrace.DNSStartInfo) *httptrace.DNSStartInfo {
			return &x
		},
		"emptyGotConnInfo": func() (v httptrace.GotConnInfo) {
			return v
		},
		"refGotConnInfo": func() *httptrace.GotConnInfo {
			var x httptrace.GotConnInfo
			return &x
		},
		"refOfGotConnInfo": func(x httptrace.GotConnInfo) *httptrace.GotConnInfo {
			return &x
		}}
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
