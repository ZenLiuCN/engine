// Code generated by define_gene; DO NOT EDIT.
package lark

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/base-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/base-sdk-go/v3/service/base/v1"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/base-sdk-go/v3/service/drive/v1"
	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/http"
	_ "github.com/ZenLiuCN/engine/modules/golang/time"
	"github.com/larksuite/base-sdk-go/v3"
)

var (
	//go:embed github_com_larksuite_base-sdk-go_v3.d.ts
	GithubComLarksuiteBaseSdkGo3Define   []byte
	GithubComLarksuiteBaseSdkGo3Declared = map[string]any{
		"withLogLevel":      lark.WithLogLevel,
		"withLogger":        lark.WithLogger,
		"withOpenBaseUrl":   lark.WithOpenBaseUrl,
		"withSerialization": lark.WithSerialization,
		"FeishuBaseUrl":     lark.FeishuBaseUrl,
		"newClient":         lark.NewClient,
		"withHeaders":       lark.WithHeaders,
		"withHttpClient":    lark.WithHttpClient,
		"LarkBaseUrl":       lark.LarkBaseUrl,
		"withLogReqAtDebug": lark.WithLogReqAtDebug,
		"withReqTimeout":    lark.WithReqTimeout,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteBaseSdkGo3Module{})
}

type GithubComLarksuiteBaseSdkGo3Module struct{}

func (S GithubComLarksuiteBaseSdkGo3Module) Identity() string {
	return "github.com/larksuite/base-sdk-go/v3"
}
func (S GithubComLarksuiteBaseSdkGo3Module) TypeDefine() []byte {
	return GithubComLarksuiteBaseSdkGo3Define
}
func (S GithubComLarksuiteBaseSdkGo3Module) Exports() map[string]any {
	return GithubComLarksuiteBaseSdkGo3Declared
}
