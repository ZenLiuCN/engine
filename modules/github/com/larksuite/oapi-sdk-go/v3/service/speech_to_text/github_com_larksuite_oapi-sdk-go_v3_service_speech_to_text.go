// Code generated by define_gene; DO NOT EDIT.
package speech_to_text

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/service/speech_to_text/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/speech_to_text"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_speech_to_text.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceSpeech_to_textDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceSpeech_to_textDeclared = map[string]any{
		"newService": speech_to_text.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceSpeech_to_textModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceSpeech_to_textModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceSpeech_to_textModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/speech_to_text"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceSpeech_to_textModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceSpeech_to_textDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceSpeech_to_textModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceSpeech_to_textDeclared
}
