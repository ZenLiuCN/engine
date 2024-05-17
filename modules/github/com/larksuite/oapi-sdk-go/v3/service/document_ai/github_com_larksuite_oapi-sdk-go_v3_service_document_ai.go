// Code generated by define_gene; DO NOT EDIT.
package document_ai

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/document_ai/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/document_ai"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_document_ai.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceDocument_aiDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceDocument_aiDeclared = map[string]any{
		"newService": document_ai.NewService,

		"emptyService":    engine.Empty[document_ai.Service],
		"emptyRefService": engine.EmptyRefer[document_ai.Service],
		"refOfService":    engine.ReferOf[document_ai.Service],
		"unRefService":    engine.UnRefer[document_ai.Service]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceDocument_aiModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceDocument_aiModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceDocument_aiModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/document_ai"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceDocument_aiModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceDocument_aiDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceDocument_aiModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceDocument_aiDeclared
}
