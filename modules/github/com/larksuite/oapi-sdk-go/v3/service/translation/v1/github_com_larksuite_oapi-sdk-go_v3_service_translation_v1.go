// Code generated by define_gene; DO NOT EDIT.
package larktranslation

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/translation/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_translation_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceTranslation1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceTranslation1Declared = map[string]any{
		"New":                                larktranslation.New,
		"newDetectTextReqBodyBuilder":        larktranslation.NewDetectTextReqBodyBuilder,
		"newTranslateTextReqBodyBuilder":     larktranslation.NewTranslateTextReqBodyBuilder,
		"newTranslateTextPathReqBodyBuilder": larktranslation.NewTranslateTextPathReqBodyBuilder,
		"newTranslateTextReqBuilder":         larktranslation.NewTranslateTextReqBuilder,
		"newDepartmentIdBuilder":             larktranslation.NewDepartmentIdBuilder,
		"newDetectTextPathReqBodyBuilder":    larktranslation.NewDetectTextPathReqBodyBuilder,
		"newDetectTextReqBuilder":            larktranslation.NewDetectTextReqBuilder,
		"newTermBuilder":                     larktranslation.NewTermBuilder,

		"emptyTranslateTextReq":                   engine.Empty[larktranslation.TranslateTextReq],
		"emptyRefTranslateTextReq":                engine.EmptyRefer[larktranslation.TranslateTextReq],
		"refOfTranslateTextReq":                   engine.ReferOf[larktranslation.TranslateTextReq],
		"unRefTranslateTextReq":                   engine.UnRefer[larktranslation.TranslateTextReq],
		"emptyTranslateTextReqBuilder":            engine.Empty[larktranslation.TranslateTextReqBuilder],
		"emptyRefTranslateTextReqBuilder":         engine.EmptyRefer[larktranslation.TranslateTextReqBuilder],
		"refOfTranslateTextReqBuilder":            engine.ReferOf[larktranslation.TranslateTextReqBuilder],
		"unRefTranslateTextReqBuilder":            engine.UnRefer[larktranslation.TranslateTextReqBuilder],
		"emptyTranslateTextRespData":              engine.Empty[larktranslation.TranslateTextRespData],
		"emptyRefTranslateTextRespData":           engine.EmptyRefer[larktranslation.TranslateTextRespData],
		"refOfTranslateTextRespData":              engine.ReferOf[larktranslation.TranslateTextRespData],
		"unRefTranslateTextRespData":              engine.UnRefer[larktranslation.TranslateTextRespData],
		"emptyDetectTextReqBuilder":               engine.Empty[larktranslation.DetectTextReqBuilder],
		"emptyRefDetectTextReqBuilder":            engine.EmptyRefer[larktranslation.DetectTextReqBuilder],
		"refOfDetectTextReqBuilder":               engine.ReferOf[larktranslation.DetectTextReqBuilder],
		"unRefDetectTextReqBuilder":               engine.UnRefer[larktranslation.DetectTextReqBuilder],
		"emptyDetectTextRespData":                 engine.Empty[larktranslation.DetectTextRespData],
		"emptyRefDetectTextRespData":              engine.EmptyRefer[larktranslation.DetectTextRespData],
		"refOfDetectTextRespData":                 engine.ReferOf[larktranslation.DetectTextRespData],
		"unRefDetectTextRespData":                 engine.UnRefer[larktranslation.DetectTextRespData],
		"emptyTerm":                               engine.Empty[larktranslation.Term],
		"emptyRefTerm":                            engine.EmptyRefer[larktranslation.Term],
		"refOfTerm":                               engine.ReferOf[larktranslation.Term],
		"unRefTerm":                               engine.UnRefer[larktranslation.Term],
		"emptyTermBuilder":                        engine.Empty[larktranslation.TermBuilder],
		"emptyRefTermBuilder":                     engine.EmptyRefer[larktranslation.TermBuilder],
		"refOfTermBuilder":                        engine.ReferOf[larktranslation.TermBuilder],
		"unRefTermBuilder":                        engine.UnRefer[larktranslation.TermBuilder],
		"emptyDetectTextReqBody":                  engine.Empty[larktranslation.DetectTextReqBody],
		"emptyRefDetectTextReqBody":               engine.EmptyRefer[larktranslation.DetectTextReqBody],
		"refOfDetectTextReqBody":                  engine.ReferOf[larktranslation.DetectTextReqBody],
		"unRefDetectTextReqBody":                  engine.UnRefer[larktranslation.DetectTextReqBody],
		"emptyDetectTextResp":                     engine.Empty[larktranslation.DetectTextResp],
		"emptyRefDetectTextResp":                  engine.EmptyRefer[larktranslation.DetectTextResp],
		"refOfDetectTextResp":                     engine.ReferOf[larktranslation.DetectTextResp],
		"unRefDetectTextResp":                     engine.UnRefer[larktranslation.DetectTextResp],
		"emptyTranslateTextReqBody":               engine.Empty[larktranslation.TranslateTextReqBody],
		"emptyRefTranslateTextReqBody":            engine.EmptyRefer[larktranslation.TranslateTextReqBody],
		"refOfTranslateTextReqBody":               engine.ReferOf[larktranslation.TranslateTextReqBody],
		"unRefTranslateTextReqBody":               engine.UnRefer[larktranslation.TranslateTextReqBody],
		"emptyTranslateTextPathReqBodyBuilder":    engine.Empty[larktranslation.TranslateTextPathReqBodyBuilder],
		"emptyRefTranslateTextPathReqBodyBuilder": engine.EmptyRefer[larktranslation.TranslateTextPathReqBodyBuilder],
		"refOfTranslateTextPathReqBodyBuilder":    engine.ReferOf[larktranslation.TranslateTextPathReqBodyBuilder],
		"unRefTranslateTextPathReqBodyBuilder":    engine.UnRefer[larktranslation.TranslateTextPathReqBodyBuilder],
		"emptyTranslateTextReqBodyBuilder":        engine.Empty[larktranslation.TranslateTextReqBodyBuilder],
		"emptyRefTranslateTextReqBodyBuilder":     engine.EmptyRefer[larktranslation.TranslateTextReqBodyBuilder],
		"refOfTranslateTextReqBodyBuilder":        engine.ReferOf[larktranslation.TranslateTextReqBodyBuilder],
		"unRefTranslateTextReqBodyBuilder":        engine.UnRefer[larktranslation.TranslateTextReqBodyBuilder],
		"emptyTranslateTextResp":                  engine.Empty[larktranslation.TranslateTextResp],
		"emptyRefTranslateTextResp":               engine.EmptyRefer[larktranslation.TranslateTextResp],
		"refOfTranslateTextResp":                  engine.ReferOf[larktranslation.TranslateTextResp],
		"unRefTranslateTextResp":                  engine.UnRefer[larktranslation.TranslateTextResp],
		"emptyV1":                                 engine.Empty[larktranslation.V1],
		"emptyRefV1":                              engine.EmptyRefer[larktranslation.V1],
		"refOfV1":                                 engine.ReferOf[larktranslation.V1],
		"unRefV1":                                 engine.UnRefer[larktranslation.V1],
		"emptyDepartmentId":                       engine.Empty[larktranslation.DepartmentId],
		"emptyRefDepartmentId":                    engine.EmptyRefer[larktranslation.DepartmentId],
		"refOfDepartmentId":                       engine.ReferOf[larktranslation.DepartmentId],
		"unRefDepartmentId":                       engine.UnRefer[larktranslation.DepartmentId],
		"emptyDepartmentIdBuilder":                engine.Empty[larktranslation.DepartmentIdBuilder],
		"emptyRefDepartmentIdBuilder":             engine.EmptyRefer[larktranslation.DepartmentIdBuilder],
		"refOfDepartmentIdBuilder":                engine.ReferOf[larktranslation.DepartmentIdBuilder],
		"unRefDepartmentIdBuilder":                engine.UnRefer[larktranslation.DepartmentIdBuilder],
		"emptyDetectTextReq":                      engine.Empty[larktranslation.DetectTextReq],
		"emptyRefDetectTextReq":                   engine.EmptyRefer[larktranslation.DetectTextReq],
		"refOfDetectTextReq":                      engine.ReferOf[larktranslation.DetectTextReq],
		"unRefDetectTextReq":                      engine.UnRefer[larktranslation.DetectTextReq],
		"emptyDetectTextReqBodyBuilder":           engine.Empty[larktranslation.DetectTextReqBodyBuilder],
		"emptyRefDetectTextReqBodyBuilder":        engine.EmptyRefer[larktranslation.DetectTextReqBodyBuilder],
		"refOfDetectTextReqBodyBuilder":           engine.ReferOf[larktranslation.DetectTextReqBodyBuilder],
		"unRefDetectTextReqBodyBuilder":           engine.UnRefer[larktranslation.DetectTextReqBodyBuilder],
		"emptyDetectTextPathReqBodyBuilder":       engine.Empty[larktranslation.DetectTextPathReqBodyBuilder],
		"emptyRefDetectTextPathReqBodyBuilder":    engine.EmptyRefer[larktranslation.DetectTextPathReqBodyBuilder],
		"refOfDetectTextPathReqBodyBuilder":       engine.ReferOf[larktranslation.DetectTextPathReqBodyBuilder],
		"unRefDetectTextPathReqBodyBuilder":       engine.UnRefer[larktranslation.DetectTextPathReqBodyBuilder]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceTranslation1Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceTranslation1Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceTranslation1Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/translation/v1"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceTranslation1Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceTranslation1Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceTranslation1Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceTranslation1Declared
}
