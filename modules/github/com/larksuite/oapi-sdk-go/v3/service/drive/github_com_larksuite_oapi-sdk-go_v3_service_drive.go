// Code generated by define_gene; DO NOT EDIT.
package drive

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/drive/v1"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/drive/v2"
	"github.com/larksuite/oapi-sdk-go/v3/service/drive"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_drive.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceDriveDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceDriveDeclared = map[string]any{
		"newService": drive.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceDriveModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceDriveModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceDriveModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/drive"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceDriveModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceDriveDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceDriveModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceDriveDeclared
}