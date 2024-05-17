// Code generated by define_gene; DO NOT EDIT.
package larkgray_test_open_sg

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/gray_test_open_sg/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_gray_test_open_sg_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceGray_test_open_sg1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceGray_test_open_sg1Declared = map[string]any{
		"newLevelBuilder":         larkgray_test_open_sg.NewLevelBuilder,
		"newListMotoReqBuilder":   larkgray_test_open_sg.NewListMotoReqBuilder,
		"newMotoBuilder":          larkgray_test_open_sg.NewMotoBuilder,
		"New":                     larkgray_test_open_sg.New,
		"newCreateMotoReqBuilder": larkgray_test_open_sg.NewCreateMotoReqBuilder,
		"newDepartmentIdBuilder":  larkgray_test_open_sg.NewDepartmentIdBuilder,
		"newGetMotoReqBuilder":    larkgray_test_open_sg.NewGetMotoReqBuilder,

		"emptyCreateMotoResp":          engine.Empty[larkgray_test_open_sg.CreateMotoResp],
		"emptyRefCreateMotoResp":       engine.EmptyRefer[larkgray_test_open_sg.CreateMotoResp],
		"refOfCreateMotoResp":          engine.ReferOf[larkgray_test_open_sg.CreateMotoResp],
		"unRefCreateMotoResp":          engine.UnRefer[larkgray_test_open_sg.CreateMotoResp],
		"emptyGetMotoReqBuilder":       engine.Empty[larkgray_test_open_sg.GetMotoReqBuilder],
		"emptyRefGetMotoReqBuilder":    engine.EmptyRefer[larkgray_test_open_sg.GetMotoReqBuilder],
		"refOfGetMotoReqBuilder":       engine.ReferOf[larkgray_test_open_sg.GetMotoReqBuilder],
		"unRefGetMotoReqBuilder":       engine.UnRefer[larkgray_test_open_sg.GetMotoReqBuilder],
		"emptyGetMotoResp":             engine.Empty[larkgray_test_open_sg.GetMotoResp],
		"emptyRefGetMotoResp":          engine.EmptyRefer[larkgray_test_open_sg.GetMotoResp],
		"refOfGetMotoResp":             engine.ReferOf[larkgray_test_open_sg.GetMotoResp],
		"unRefGetMotoResp":             engine.UnRefer[larkgray_test_open_sg.GetMotoResp],
		"emptyGetMotoRespData":         engine.Empty[larkgray_test_open_sg.GetMotoRespData],
		"emptyRefGetMotoRespData":      engine.EmptyRefer[larkgray_test_open_sg.GetMotoRespData],
		"refOfGetMotoRespData":         engine.ReferOf[larkgray_test_open_sg.GetMotoRespData],
		"unRefGetMotoRespData":         engine.UnRefer[larkgray_test_open_sg.GetMotoRespData],
		"emptyListMotoIterator":        engine.Empty[larkgray_test_open_sg.ListMotoIterator],
		"emptyRefListMotoIterator":     engine.EmptyRefer[larkgray_test_open_sg.ListMotoIterator],
		"refOfListMotoIterator":        engine.ReferOf[larkgray_test_open_sg.ListMotoIterator],
		"unRefListMotoIterator":        engine.UnRefer[larkgray_test_open_sg.ListMotoIterator],
		"emptyListMotoReqBuilder":      engine.Empty[larkgray_test_open_sg.ListMotoReqBuilder],
		"emptyRefListMotoReqBuilder":   engine.EmptyRefer[larkgray_test_open_sg.ListMotoReqBuilder],
		"refOfListMotoReqBuilder":      engine.ReferOf[larkgray_test_open_sg.ListMotoReqBuilder],
		"unRefListMotoReqBuilder":      engine.UnRefer[larkgray_test_open_sg.ListMotoReqBuilder],
		"emptyListMotoRespData":        engine.Empty[larkgray_test_open_sg.ListMotoRespData],
		"emptyRefListMotoRespData":     engine.EmptyRefer[larkgray_test_open_sg.ListMotoRespData],
		"refOfListMotoRespData":        engine.ReferOf[larkgray_test_open_sg.ListMotoRespData],
		"unRefListMotoRespData":        engine.UnRefer[larkgray_test_open_sg.ListMotoRespData],
		"emptyMotoBuilder":             engine.Empty[larkgray_test_open_sg.MotoBuilder],
		"emptyRefMotoBuilder":          engine.EmptyRefer[larkgray_test_open_sg.MotoBuilder],
		"refOfMotoBuilder":             engine.ReferOf[larkgray_test_open_sg.MotoBuilder],
		"unRefMotoBuilder":             engine.UnRefer[larkgray_test_open_sg.MotoBuilder],
		"emptyCreateMotoReq":           engine.Empty[larkgray_test_open_sg.CreateMotoReq],
		"emptyRefCreateMotoReq":        engine.EmptyRefer[larkgray_test_open_sg.CreateMotoReq],
		"refOfCreateMotoReq":           engine.ReferOf[larkgray_test_open_sg.CreateMotoReq],
		"unRefCreateMotoReq":           engine.UnRefer[larkgray_test_open_sg.CreateMotoReq],
		"emptyCreateMotoRespData":      engine.Empty[larkgray_test_open_sg.CreateMotoRespData],
		"emptyRefCreateMotoRespData":   engine.EmptyRefer[larkgray_test_open_sg.CreateMotoRespData],
		"refOfCreateMotoRespData":      engine.ReferOf[larkgray_test_open_sg.CreateMotoRespData],
		"unRefCreateMotoRespData":      engine.UnRefer[larkgray_test_open_sg.CreateMotoRespData],
		"emptyGetMotoReq":              engine.Empty[larkgray_test_open_sg.GetMotoReq],
		"emptyRefGetMotoReq":           engine.EmptyRefer[larkgray_test_open_sg.GetMotoReq],
		"refOfGetMotoReq":              engine.ReferOf[larkgray_test_open_sg.GetMotoReq],
		"unRefGetMotoReq":              engine.UnRefer[larkgray_test_open_sg.GetMotoReq],
		"emptyListMotoResp":            engine.Empty[larkgray_test_open_sg.ListMotoResp],
		"emptyRefListMotoResp":         engine.EmptyRefer[larkgray_test_open_sg.ListMotoResp],
		"refOfListMotoResp":            engine.ReferOf[larkgray_test_open_sg.ListMotoResp],
		"unRefListMotoResp":            engine.UnRefer[larkgray_test_open_sg.ListMotoResp],
		"emptyDepartmentIdBuilder":     engine.Empty[larkgray_test_open_sg.DepartmentIdBuilder],
		"emptyRefDepartmentIdBuilder":  engine.EmptyRefer[larkgray_test_open_sg.DepartmentIdBuilder],
		"refOfDepartmentIdBuilder":     engine.ReferOf[larkgray_test_open_sg.DepartmentIdBuilder],
		"unRefDepartmentIdBuilder":     engine.UnRefer[larkgray_test_open_sg.DepartmentIdBuilder],
		"emptyLevel":                   engine.Empty[larkgray_test_open_sg.Level],
		"emptyRefLevel":                engine.EmptyRefer[larkgray_test_open_sg.Level],
		"refOfLevel":                   engine.ReferOf[larkgray_test_open_sg.Level],
		"unRefLevel":                   engine.UnRefer[larkgray_test_open_sg.Level],
		"emptyLevelBuilder":            engine.Empty[larkgray_test_open_sg.LevelBuilder],
		"emptyRefLevelBuilder":         engine.EmptyRefer[larkgray_test_open_sg.LevelBuilder],
		"refOfLevelBuilder":            engine.ReferOf[larkgray_test_open_sg.LevelBuilder],
		"unRefLevelBuilder":            engine.UnRefer[larkgray_test_open_sg.LevelBuilder],
		"emptyMoto":                    engine.Empty[larkgray_test_open_sg.Moto],
		"emptyRefMoto":                 engine.EmptyRefer[larkgray_test_open_sg.Moto],
		"refOfMoto":                    engine.ReferOf[larkgray_test_open_sg.Moto],
		"unRefMoto":                    engine.UnRefer[larkgray_test_open_sg.Moto],
		"emptyV1":                      engine.Empty[larkgray_test_open_sg.V1],
		"emptyRefV1":                   engine.EmptyRefer[larkgray_test_open_sg.V1],
		"refOfV1":                      engine.ReferOf[larkgray_test_open_sg.V1],
		"unRefV1":                      engine.UnRefer[larkgray_test_open_sg.V1],
		"emptyCreateMotoReqBuilder":    engine.Empty[larkgray_test_open_sg.CreateMotoReqBuilder],
		"emptyRefCreateMotoReqBuilder": engine.EmptyRefer[larkgray_test_open_sg.CreateMotoReqBuilder],
		"refOfCreateMotoReqBuilder":    engine.ReferOf[larkgray_test_open_sg.CreateMotoReqBuilder],
		"unRefCreateMotoReqBuilder":    engine.UnRefer[larkgray_test_open_sg.CreateMotoReqBuilder],
		"emptyDepartmentId":            engine.Empty[larkgray_test_open_sg.DepartmentId],
		"emptyRefDepartmentId":         engine.EmptyRefer[larkgray_test_open_sg.DepartmentId],
		"refOfDepartmentId":            engine.ReferOf[larkgray_test_open_sg.DepartmentId],
		"unRefDepartmentId":            engine.UnRefer[larkgray_test_open_sg.DepartmentId],
		"emptyListMotoReq":             engine.Empty[larkgray_test_open_sg.ListMotoReq],
		"emptyRefListMotoReq":          engine.EmptyRefer[larkgray_test_open_sg.ListMotoReq],
		"refOfListMotoReq":             engine.ReferOf[larkgray_test_open_sg.ListMotoReq],
		"unRefListMotoReq":             engine.UnRefer[larkgray_test_open_sg.ListMotoReq]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceGray_test_open_sg1Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceGray_test_open_sg1Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceGray_test_open_sg1Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/gray_test_open_sg/v1"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceGray_test_open_sg1Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceGray_test_open_sg1Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceGray_test_open_sg1Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceGray_test_open_sg1Declared
}
