// Code generated by define_gene; DO NOT EDIT.
package larkoptical_char_recognition

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/optical_char_recognition/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_optical_char_recognition_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognition1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognition1Declared = map[string]any{
		"New": larkoptical_char_recognition.New,
		"newBasicRecognizeImagePathReqBodyBuilder": larkoptical_char_recognition.NewBasicRecognizeImagePathReqBodyBuilder,
		"newBasicRecognizeImageReqBodyBuilder":     larkoptical_char_recognition.NewBasicRecognizeImageReqBodyBuilder,
		"newBasicRecognizeImageReqBuilder":         larkoptical_char_recognition.NewBasicRecognizeImageReqBuilder,
		"newDepartmentIdBuilder":                   larkoptical_char_recognition.NewDepartmentIdBuilder,

		"emptyBasicRecognizeImageReqBodyBuilder":        engine.Empty[larkoptical_char_recognition.BasicRecognizeImageReqBodyBuilder],
		"emptyRefBasicRecognizeImageReqBodyBuilder":     engine.EmptyRefer[larkoptical_char_recognition.BasicRecognizeImageReqBodyBuilder],
		"refOfBasicRecognizeImageReqBodyBuilder":        engine.ReferOf[larkoptical_char_recognition.BasicRecognizeImageReqBodyBuilder],
		"unRefBasicRecognizeImageReqBodyBuilder":        engine.UnRefer[larkoptical_char_recognition.BasicRecognizeImageReqBodyBuilder],
		"emptyBasicRecognizeImageReqBuilder":            engine.Empty[larkoptical_char_recognition.BasicRecognizeImageReqBuilder],
		"emptyRefBasicRecognizeImageReqBuilder":         engine.EmptyRefer[larkoptical_char_recognition.BasicRecognizeImageReqBuilder],
		"refOfBasicRecognizeImageReqBuilder":            engine.ReferOf[larkoptical_char_recognition.BasicRecognizeImageReqBuilder],
		"unRefBasicRecognizeImageReqBuilder":            engine.UnRefer[larkoptical_char_recognition.BasicRecognizeImageReqBuilder],
		"emptyBasicRecognizeImageResp":                  engine.Empty[larkoptical_char_recognition.BasicRecognizeImageResp],
		"emptyRefBasicRecognizeImageResp":               engine.EmptyRefer[larkoptical_char_recognition.BasicRecognizeImageResp],
		"refOfBasicRecognizeImageResp":                  engine.ReferOf[larkoptical_char_recognition.BasicRecognizeImageResp],
		"unRefBasicRecognizeImageResp":                  engine.UnRefer[larkoptical_char_recognition.BasicRecognizeImageResp],
		"emptyBasicRecognizeImageRespData":              engine.Empty[larkoptical_char_recognition.BasicRecognizeImageRespData],
		"emptyRefBasicRecognizeImageRespData":           engine.EmptyRefer[larkoptical_char_recognition.BasicRecognizeImageRespData],
		"refOfBasicRecognizeImageRespData":              engine.ReferOf[larkoptical_char_recognition.BasicRecognizeImageRespData],
		"unRefBasicRecognizeImageRespData":              engine.UnRefer[larkoptical_char_recognition.BasicRecognizeImageRespData],
		"emptyDepartmentId":                             engine.Empty[larkoptical_char_recognition.DepartmentId],
		"emptyRefDepartmentId":                          engine.EmptyRefer[larkoptical_char_recognition.DepartmentId],
		"refOfDepartmentId":                             engine.ReferOf[larkoptical_char_recognition.DepartmentId],
		"unRefDepartmentId":                             engine.UnRefer[larkoptical_char_recognition.DepartmentId],
		"emptyDepartmentIdBuilder":                      engine.Empty[larkoptical_char_recognition.DepartmentIdBuilder],
		"emptyRefDepartmentIdBuilder":                   engine.EmptyRefer[larkoptical_char_recognition.DepartmentIdBuilder],
		"refOfDepartmentIdBuilder":                      engine.ReferOf[larkoptical_char_recognition.DepartmentIdBuilder],
		"unRefDepartmentIdBuilder":                      engine.UnRefer[larkoptical_char_recognition.DepartmentIdBuilder],
		"emptyV1":                                       engine.Empty[larkoptical_char_recognition.V1],
		"emptyRefV1":                                    engine.EmptyRefer[larkoptical_char_recognition.V1],
		"refOfV1":                                       engine.ReferOf[larkoptical_char_recognition.V1],
		"unRefV1":                                       engine.UnRefer[larkoptical_char_recognition.V1],
		"emptyBasicRecognizeImagePathReqBodyBuilder":    engine.Empty[larkoptical_char_recognition.BasicRecognizeImagePathReqBodyBuilder],
		"emptyRefBasicRecognizeImagePathReqBodyBuilder": engine.EmptyRefer[larkoptical_char_recognition.BasicRecognizeImagePathReqBodyBuilder],
		"refOfBasicRecognizeImagePathReqBodyBuilder":    engine.ReferOf[larkoptical_char_recognition.BasicRecognizeImagePathReqBodyBuilder],
		"unRefBasicRecognizeImagePathReqBodyBuilder":    engine.UnRefer[larkoptical_char_recognition.BasicRecognizeImagePathReqBodyBuilder],
		"emptyBasicRecognizeImageReq":                   engine.Empty[larkoptical_char_recognition.BasicRecognizeImageReq],
		"emptyRefBasicRecognizeImageReq":                engine.EmptyRefer[larkoptical_char_recognition.BasicRecognizeImageReq],
		"refOfBasicRecognizeImageReq":                   engine.ReferOf[larkoptical_char_recognition.BasicRecognizeImageReq],
		"unRefBasicRecognizeImageReq":                   engine.UnRefer[larkoptical_char_recognition.BasicRecognizeImageReq],
		"emptyBasicRecognizeImageReqBody":               engine.Empty[larkoptical_char_recognition.BasicRecognizeImageReqBody],
		"emptyRefBasicRecognizeImageReqBody":            engine.EmptyRefer[larkoptical_char_recognition.BasicRecognizeImageReqBody],
		"refOfBasicRecognizeImageReqBody":               engine.ReferOf[larkoptical_char_recognition.BasicRecognizeImageReqBody],
		"unRefBasicRecognizeImageReqBody":               engine.UnRefer[larkoptical_char_recognition.BasicRecognizeImageReqBody]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognition1Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognition1Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognition1Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/optical_char_recognition/v1"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognition1Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognition1Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognition1Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognition1Declared
}
