// Code generated by define_gene; DO NOT EDIT.
package baike

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/baike/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/baike"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_baike.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceBaikeDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceBaikeDeclared = map[string]any{
		"newService": baike.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceBaikeModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceBaikeModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceBaikeModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/baike"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceBaikeModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceBaikeDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceBaikeModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceBaikeDeclared
}