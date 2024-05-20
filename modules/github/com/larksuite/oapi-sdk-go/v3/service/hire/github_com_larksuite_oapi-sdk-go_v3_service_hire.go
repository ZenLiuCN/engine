// Code generated by define_gene; DO NOT EDIT.
package hire

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/hire/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/hire"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_hire.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceHireDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceHireDeclared = map[string]any{
		"newService": hire.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceHireModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceHireModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceHireModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/hire"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceHireModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceHireDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceHireModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceHireDeclared
}