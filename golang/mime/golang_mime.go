package mime

import (
	_ "embed"
	_ "github.com/ZenLiuCN/engine/golang/io"

	"github.com/ZenLiuCN/engine"
	"mime"
)

var (
	//go:embed golang_mime.d.ts
	MimeDefine   []byte
	MimeDeclared = map[string]any{
		"extensionsByType": mime.ExtensionsByType,
		"addExtensionType": mime.AddExtensionType,
		"formatMediaType":  mime.FormatMediaType,
		"parseMediaType":   mime.ParseMediaType,
		"typeByExtension":  mime.TypeByExtension,
	}
)

func init() {
	engine.RegisterModule(MimeModule{})
}

type MimeModule struct{}

func (S MimeModule) Identity() string {
	return "golang/mime"
}
func (S MimeModule) TypeDefine() []byte {
	return MimeDefine
}
func (S MimeModule) Exports() map[string]any {
	return MimeDeclared
}
