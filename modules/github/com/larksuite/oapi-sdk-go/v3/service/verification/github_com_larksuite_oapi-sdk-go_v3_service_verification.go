// Code generated by define_gene; DO NOT EDIT.
package verification

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/verification/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/verification"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_verification.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceVerificationDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceVerificationDeclared = map[string]any{
		"newService": verification.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceVerificationModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceVerificationModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceVerificationModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/verification"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceVerificationModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceVerificationDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceVerificationModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceVerificationDeclared
}