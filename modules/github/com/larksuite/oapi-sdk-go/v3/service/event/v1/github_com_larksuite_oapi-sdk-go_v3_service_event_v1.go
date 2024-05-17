// Code generated by define_gene; DO NOT EDIT.
package larkevent

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/event/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_event_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceEvent1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceEvent1Declared = map[string]any{
		"New":                         larkevent.New,
		"newDepartmentIdBuilder":      larkevent.NewDepartmentIdBuilder,
		"newListOutboundIpReqBuilder": larkevent.NewListOutboundIpReqBuilder,

		"emptyListOutboundIpIterator":      engine.Empty[larkevent.ListOutboundIpIterator],
		"emptyRefListOutboundIpIterator":   engine.EmptyRefer[larkevent.ListOutboundIpIterator],
		"refOfListOutboundIpIterator":      engine.ReferOf[larkevent.ListOutboundIpIterator],
		"unRefListOutboundIpIterator":      engine.UnRefer[larkevent.ListOutboundIpIterator],
		"emptyListOutboundIpReq":           engine.Empty[larkevent.ListOutboundIpReq],
		"emptyRefListOutboundIpReq":        engine.EmptyRefer[larkevent.ListOutboundIpReq],
		"refOfListOutboundIpReq":           engine.ReferOf[larkevent.ListOutboundIpReq],
		"unRefListOutboundIpReq":           engine.UnRefer[larkevent.ListOutboundIpReq],
		"emptyListOutboundIpReqBuilder":    engine.Empty[larkevent.ListOutboundIpReqBuilder],
		"emptyRefListOutboundIpReqBuilder": engine.EmptyRefer[larkevent.ListOutboundIpReqBuilder],
		"refOfListOutboundIpReqBuilder":    engine.ReferOf[larkevent.ListOutboundIpReqBuilder],
		"unRefListOutboundIpReqBuilder":    engine.UnRefer[larkevent.ListOutboundIpReqBuilder],
		"emptyListOutboundIpResp":          engine.Empty[larkevent.ListOutboundIpResp],
		"emptyRefListOutboundIpResp":       engine.EmptyRefer[larkevent.ListOutboundIpResp],
		"refOfListOutboundIpResp":          engine.ReferOf[larkevent.ListOutboundIpResp],
		"unRefListOutboundIpResp":          engine.UnRefer[larkevent.ListOutboundIpResp],
		"emptyListOutboundIpRespData":      engine.Empty[larkevent.ListOutboundIpRespData],
		"emptyRefListOutboundIpRespData":   engine.EmptyRefer[larkevent.ListOutboundIpRespData],
		"refOfListOutboundIpRespData":      engine.ReferOf[larkevent.ListOutboundIpRespData],
		"unRefListOutboundIpRespData":      engine.UnRefer[larkevent.ListOutboundIpRespData],
		"emptyV1":                          engine.Empty[larkevent.V1],
		"emptyRefV1":                       engine.EmptyRefer[larkevent.V1],
		"refOfV1":                          engine.ReferOf[larkevent.V1],
		"unRefV1":                          engine.UnRefer[larkevent.V1],
		"emptyDepartmentId":                engine.Empty[larkevent.DepartmentId],
		"emptyRefDepartmentId":             engine.EmptyRefer[larkevent.DepartmentId],
		"refOfDepartmentId":                engine.ReferOf[larkevent.DepartmentId],
		"unRefDepartmentId":                engine.UnRefer[larkevent.DepartmentId],
		"emptyDepartmentIdBuilder":         engine.Empty[larkevent.DepartmentIdBuilder],
		"emptyRefDepartmentIdBuilder":      engine.EmptyRefer[larkevent.DepartmentIdBuilder],
		"refOfDepartmentIdBuilder":         engine.ReferOf[larkevent.DepartmentIdBuilder],
		"unRefDepartmentIdBuilder":         engine.UnRefer[larkevent.DepartmentIdBuilder]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceEvent1Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceEvent1Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceEvent1Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/event/v1"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceEvent1Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceEvent1Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceEvent1Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceEvent1Declared
}
