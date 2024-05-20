// Code generated by define_gene; DO NOT EDIT.
package workplace

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/workplace/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/workplace"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_workplace.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceWorkplaceDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceWorkplaceDeclared = map[string]any{
		"newService": workplace.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceWorkplaceModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceWorkplaceModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceWorkplaceModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/workplace"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceWorkplaceModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceWorkplaceDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceWorkplaceModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceWorkplaceDeclared
}
