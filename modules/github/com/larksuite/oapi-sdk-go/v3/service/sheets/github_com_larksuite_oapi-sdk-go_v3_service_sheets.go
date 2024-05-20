// Code generated by define_gene; DO NOT EDIT.
package sheets

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/sheets/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/sheets"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_sheets.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceSheetsDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceSheetsDeclared = map[string]any{
		"newService": sheets.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceSheetsModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceSheetsModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceSheetsModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/sheets"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceSheetsModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceSheetsDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceSheetsModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceSheetsDeclared
}