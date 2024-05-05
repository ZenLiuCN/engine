package textproto

import (
	_ "embed"
	_ "github.com/ZenLiuCN/engine/golang/bufio"
	_ "github.com/ZenLiuCN/engine/golang/io"

	"github.com/ZenLiuCN/engine"
	"net/textproto"
)

var (
	//go:embed golang_net_textproto.d.ts
	NetTextprotoDefine   []byte
	NetTextprotoDeclared = map[string]any{
		"canonicalMimeHeaderKey": textproto.CanonicalMIMEHeaderKey,
		"trimBytes":              textproto.TrimBytes,
		"newConn":                textproto.NewConn,
		"dial":                   textproto.Dial,
		"newReader":              textproto.NewReader,
		"trimString":             textproto.TrimString,
		"newWriter":              textproto.NewWriter,
	}
)

func init() {
	engine.RegisterModule(NetTextprotoModule{})
}

type NetTextprotoModule struct{}

func (S NetTextprotoModule) Identity() string {
	return "golang/net/textproto"
}
func (S NetTextprotoModule) TypeDefine() []byte {
	return NetTextprotoDefine
}
func (S NetTextprotoModule) Exports() map[string]any {
	return NetTextprotoDeclared
}
