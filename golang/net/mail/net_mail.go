// Code generated by define_gene; DO NOT EDIT.
package mail

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/golang/fmt"
	_ "github.com/ZenLiuCN/engine/golang/io"
	_ "github.com/ZenLiuCN/engine/golang/mime"
	_ "github.com/ZenLiuCN/engine/golang/time"
	"net/mail"
)

var (
	//go:embed net_mail.d.ts
	NetMailDefine   []byte
	NetMailDeclared = map[string]any{
		"ErrHeaderNotPresent": mail.ErrHeaderNotPresent,
		"parseAddress":        mail.ParseAddress,
		"parseAddressList":    mail.ParseAddressList,
		"parseDate":           mail.ParseDate,
		"readMessage":         mail.ReadMessage,

		"emptyAddress": func() (v mail.Address) {
			return v
		},
		"refAddress": func() *mail.Address {
			var x mail.Address
			return &x
		},
		"refOfAddress": func(x mail.Address) *mail.Address {
			return &x
		},
		"emptyAddressParser": func() (v mail.AddressParser) {
			return v
		},
		"refAddressParser": func() *mail.AddressParser {
			var x mail.AddressParser
			return &x
		},
		"refOfAddressParser": func(x mail.AddressParser) *mail.AddressParser {
			return &x
		},
		"emptyMessage": func() (v mail.Message) {
			return v
		},
		"refMessage": func() *mail.Message {
			var x mail.Message
			return &x
		},
		"refOfMessage": func(x mail.Message) *mail.Message {
			return &x
		}}
)

func init() {
	engine.RegisterModule(NetMailModule{})
}

type NetMailModule struct{}

func (S NetMailModule) Identity() string {
	return "golang/net/mail"
}
func (S NetMailModule) TypeDefine() []byte {
	return NetMailDefine
}
func (S NetMailModule) Exports() map[string]any {
	return NetMailDeclared
}
