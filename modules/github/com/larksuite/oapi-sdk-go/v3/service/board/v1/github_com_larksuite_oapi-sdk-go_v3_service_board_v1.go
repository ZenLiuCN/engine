// Code generated by define_gene; DO NOT EDIT.
package larkboard

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	"github.com/larksuite/oapi-sdk-go/v3/service/board/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_board_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceBoard1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceBoard1Declared = map[string]any{
		"newConnectorCaptionBuilder":        larkboard.NewConnectorCaptionBuilder,
		"newMindMapBuilder":                 larkboard.NewMindMapBuilder,
		"newSectionBuilder":                 larkboard.NewSectionBuilder,
		"newTableCellBuilder":               larkboard.NewTableCellBuilder,
		"newTableCellMergeInfoBuilder":      larkboard.NewTableCellMergeInfoBuilder,
		"newTextBuilder":                    larkboard.NewTextBuilder,
		"newConnectorAttachedObjectBuilder": larkboard.NewConnectorAttachedObjectBuilder,
		"newListWhiteboardNodeReqBuilder":   larkboard.NewListWhiteboardNodeReqBuilder,
		"newStyleBuilder":                   larkboard.NewStyleBuilder,
		"newTableMetaBuilder":               larkboard.NewTableMetaBuilder,
		"New":                               larkboard.New,
		"newConnectorBuilder":               larkboard.NewConnectorBuilder,
		"newDepartmentIdBuilder":            larkboard.NewDepartmentIdBuilder,
		"newImageBuilder":                   larkboard.NewImageBuilder,
		"newTableBuilder":                   larkboard.NewTableBuilder,
		"newWhiteboardNodeBuilder":          larkboard.NewWhiteboardNodeBuilder,
		"newCompositeShapeBuilder":          larkboard.NewCompositeShapeBuilder,

		"emptyMindMap":                       engine.Empty[larkboard.MindMap],
		"emptyRefMindMap":                    engine.EmptyRefer[larkboard.MindMap],
		"refOfMindMap":                       engine.ReferOf[larkboard.MindMap],
		"unRefMindMap":                       engine.UnRefer[larkboard.MindMap],
		"emptyStyle":                         engine.Empty[larkboard.Style],
		"emptyRefStyle":                      engine.EmptyRefer[larkboard.Style],
		"refOfStyle":                         engine.ReferOf[larkboard.Style],
		"unRefStyle":                         engine.UnRefer[larkboard.Style],
		"emptyV1":                            engine.Empty[larkboard.V1],
		"emptyRefV1":                         engine.EmptyRefer[larkboard.V1],
		"refOfV1":                            engine.ReferOf[larkboard.V1],
		"unRefV1":                            engine.UnRefer[larkboard.V1],
		"emptyWhiteboardNode":                engine.Empty[larkboard.WhiteboardNode],
		"emptyRefWhiteboardNode":             engine.EmptyRefer[larkboard.WhiteboardNode],
		"refOfWhiteboardNode":                engine.ReferOf[larkboard.WhiteboardNode],
		"unRefWhiteboardNode":                engine.UnRefer[larkboard.WhiteboardNode],
		"emptyConnectorCaption":              engine.Empty[larkboard.ConnectorCaption],
		"emptyRefConnectorCaption":           engine.EmptyRefer[larkboard.ConnectorCaption],
		"refOfConnectorCaption":              engine.ReferOf[larkboard.ConnectorCaption],
		"unRefConnectorCaption":              engine.UnRefer[larkboard.ConnectorCaption],
		"emptyImage":                         engine.Empty[larkboard.Image],
		"emptyRefImage":                      engine.EmptyRefer[larkboard.Image],
		"refOfImage":                         engine.ReferOf[larkboard.Image],
		"unRefImage":                         engine.UnRefer[larkboard.Image],
		"emptyListWhiteboardNodeResp":        engine.Empty[larkboard.ListWhiteboardNodeResp],
		"emptyRefListWhiteboardNodeResp":     engine.EmptyRefer[larkboard.ListWhiteboardNodeResp],
		"refOfListWhiteboardNodeResp":        engine.ReferOf[larkboard.ListWhiteboardNodeResp],
		"unRefListWhiteboardNodeResp":        engine.UnRefer[larkboard.ListWhiteboardNodeResp],
		"emptyTableCellMergeInfo":            engine.Empty[larkboard.TableCellMergeInfo],
		"emptyRefTableCellMergeInfo":         engine.EmptyRefer[larkboard.TableCellMergeInfo],
		"refOfTableCellMergeInfo":            engine.ReferOf[larkboard.TableCellMergeInfo],
		"unRefTableCellMergeInfo":            engine.UnRefer[larkboard.TableCellMergeInfo],
		"emptyConnectorAttachedObject":       engine.Empty[larkboard.ConnectorAttachedObject],
		"emptyRefConnectorAttachedObject":    engine.EmptyRefer[larkboard.ConnectorAttachedObject],
		"refOfConnectorAttachedObject":       engine.ReferOf[larkboard.ConnectorAttachedObject],
		"unRefConnectorAttachedObject":       engine.UnRefer[larkboard.ConnectorAttachedObject],
		"emptyText":                          engine.Empty[larkboard.Text],
		"emptyRefText":                       engine.EmptyRefer[larkboard.Text],
		"refOfText":                          engine.ReferOf[larkboard.Text],
		"unRefText":                          engine.UnRefer[larkboard.Text],
		"emptySection":                       engine.Empty[larkboard.Section],
		"emptyRefSection":                    engine.EmptyRefer[larkboard.Section],
		"refOfSection":                       engine.ReferOf[larkboard.Section],
		"unRefSection":                       engine.UnRefer[larkboard.Section],
		"emptyTable":                         engine.Empty[larkboard.Table],
		"emptyRefTable":                      engine.EmptyRefer[larkboard.Table],
		"refOfTable":                         engine.ReferOf[larkboard.Table],
		"unRefTable":                         engine.UnRefer[larkboard.Table],
		"emptyConnector":                     engine.Empty[larkboard.Connector],
		"emptyRefConnector":                  engine.EmptyRefer[larkboard.Connector],
		"refOfConnector":                     engine.ReferOf[larkboard.Connector],
		"unRefConnector":                     engine.UnRefer[larkboard.Connector],
		"emptyListWhiteboardNodeRespData":    engine.Empty[larkboard.ListWhiteboardNodeRespData],
		"emptyRefListWhiteboardNodeRespData": engine.EmptyRefer[larkboard.ListWhiteboardNodeRespData],
		"refOfListWhiteboardNodeRespData":    engine.ReferOf[larkboard.ListWhiteboardNodeRespData],
		"unRefListWhiteboardNodeRespData":    engine.UnRefer[larkboard.ListWhiteboardNodeRespData],
		"emptyTableCell":                     engine.Empty[larkboard.TableCell],
		"emptyRefTableCell":                  engine.EmptyRefer[larkboard.TableCell],
		"refOfTableCell":                     engine.ReferOf[larkboard.TableCell],
		"unRefTableCell":                     engine.UnRefer[larkboard.TableCell],
		"emptyCompositeShape":                engine.Empty[larkboard.CompositeShape],
		"emptyRefCompositeShape":             engine.EmptyRefer[larkboard.CompositeShape],
		"refOfCompositeShape":                engine.ReferOf[larkboard.CompositeShape],
		"unRefCompositeShape":                engine.UnRefer[larkboard.CompositeShape],
		"emptyListWhiteboardNodeReq":         engine.Empty[larkboard.ListWhiteboardNodeReq],
		"emptyRefListWhiteboardNodeReq":      engine.EmptyRefer[larkboard.ListWhiteboardNodeReq],
		"refOfListWhiteboardNodeReq":         engine.ReferOf[larkboard.ListWhiteboardNodeReq],
		"unRefListWhiteboardNodeReq":         engine.UnRefer[larkboard.ListWhiteboardNodeReq],
		"emptyDepartmentId":                  engine.Empty[larkboard.DepartmentId],
		"emptyRefDepartmentId":               engine.EmptyRefer[larkboard.DepartmentId],
		"refOfDepartmentId":                  engine.ReferOf[larkboard.DepartmentId],
		"unRefDepartmentId":                  engine.UnRefer[larkboard.DepartmentId],
		"emptyTableMeta":                     engine.Empty[larkboard.TableMeta],
		"emptyRefTableMeta":                  engine.EmptyRefer[larkboard.TableMeta],
		"refOfTableMeta":                     engine.ReferOf[larkboard.TableMeta],
		"unRefTableMeta":                     engine.UnRefer[larkboard.TableMeta]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceBoard1Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceBoard1Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceBoard1Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/board/v1"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceBoard1Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceBoard1Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceBoard1Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceBoard1Declared
}