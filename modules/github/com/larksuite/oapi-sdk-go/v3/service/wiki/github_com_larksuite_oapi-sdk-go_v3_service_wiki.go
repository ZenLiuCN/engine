// Code generated by define_gene; DO NOT EDIT.
package wiki

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
	"github.com/larksuite/oapi-sdk-go/v3/service/wiki"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_wiki.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceWikiDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceWikiDeclared = map[string]any{
		"newService": wiki.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceWikiModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceWikiModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceWikiModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/wiki"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceWikiModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceWikiDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceWikiModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceWikiDeclared
}
