// Code generated by define_gene; DO NOT EDIT.
package bitable

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/bitable"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_bitable.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceBitableDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceBitableDeclared = map[string]any{
		"newService": bitable.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceBitableModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceBitableModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceBitableModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/bitable"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceBitableModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceBitableDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceBitableModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceBitableDeclared
}