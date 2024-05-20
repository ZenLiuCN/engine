// Code generated by define_gene; DO NOT EDIT.
package personal_settings

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/personal_settings/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/personal_settings"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_personal_settings.d.ts
	GithubComLarksuiteOapiSdkGo3ServicePersonal_settingsDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServicePersonal_settingsDeclared = map[string]any{
		"newService": personal_settings.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServicePersonal_settingsModule{})
}

type GithubComLarksuiteOapiSdkGo3ServicePersonal_settingsModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServicePersonal_settingsModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/personal_settings"
}
func (S GithubComLarksuiteOapiSdkGo3ServicePersonal_settingsModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServicePersonal_settingsDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServicePersonal_settingsModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServicePersonal_settingsDeclared
}
