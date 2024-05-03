package multipart

import (
	_ "embed"
	_ "github.com/ZenLiuCN/engine/golang/io"
	_ "github.com/ZenLiuCN/engine/golang/net/textproto"

	"github.com/ZenLiuCN/engine"
	"mime/multipart"
)

var (
	//go:embed golang_mime_multipart.d.ts
	MimeMultipartDefine   []byte
	MimeMultipartDeclared = map[string]any{
		"newWriter": multipart.NewWriter,
		"newReader": multipart.NewReader,
	}
)

func init() {
	engine.RegisterModule(MimeMultipartModule{})
}

type MimeMultipartModule struct{}

func (S MimeMultipartModule) Identity() string {
	return "golang/mime/multipart"
}
func (S MimeMultipartModule) TypeDefine() []byte {
	return MimeMultipartDefine
}
func (S MimeMultipartModule) Exports() map[string]any {
	return MimeMultipartDeclared
}
