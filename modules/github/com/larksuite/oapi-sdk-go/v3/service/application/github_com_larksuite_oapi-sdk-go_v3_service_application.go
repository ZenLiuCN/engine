// Code generated by define_gene; DO NOT EDIT.
package application

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/application/v6"
	"github.com/larksuite/oapi-sdk-go/v3/service/application"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_application.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceApplicationDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceApplicationDeclared = map[string]any{
		"newService": application.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceApplicationModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceApplicationModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceApplicationModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/application"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceApplicationModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceApplicationDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceApplicationModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceApplicationDeclared
}