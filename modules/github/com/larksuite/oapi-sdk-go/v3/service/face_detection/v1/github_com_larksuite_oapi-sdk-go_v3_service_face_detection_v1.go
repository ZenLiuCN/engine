// Code generated by define_gene; DO NOT EDIT.
package larkface_detection

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	"github.com/larksuite/oapi-sdk-go/v3/service/face_detection/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_face_detection_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceFace_detection1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceFace_detection1Declared = map[string]any{
		"newDetectFaceAttributesImageReqBuilder":         larkface_detection.NewDetectFaceAttributesImageReqBuilder,
		"newFacePositionBuilder":                         larkface_detection.NewFacePositionBuilder,
		"newDetectFaceAttributesImagePathReqBodyBuilder": larkface_detection.NewDetectFaceAttributesImagePathReqBodyBuilder,
		"newFaceOccludeBuilder":                          larkface_detection.NewFaceOccludeBuilder,
		"newService":                                     larkface_detection.NewService,
		"newFaceInfoBuilder":                             larkface_detection.NewFaceInfoBuilder,
		"newFaceQualityBuilder":                          larkface_detection.NewFaceQualityBuilder,
		"newImageBuilder":                                larkface_detection.NewImageBuilder,
		"newPointBuilder":                                larkface_detection.NewPointBuilder,
		"newAttributeItemBuilder":                        larkface_detection.NewAttributeItemBuilder,
		"newDetectFaceAttributesImageReqBodyBuilder":     larkface_detection.NewDetectFaceAttributesImageReqBodyBuilder,
		"newFaceAttributeBuilder":                        larkface_detection.NewFaceAttributeBuilder,
		"newFacePoseBuilder":                             larkface_detection.NewFacePoseBuilder,

		"emptyImage":                                engine.Empty[larkface_detection.Image],
		"emptyRefImage":                             engine.EmptyRefer[larkface_detection.Image],
		"refOfImage":                                engine.ReferOf[larkface_detection.Image],
		"unRefImage":                                engine.UnRefer[larkface_detection.Image],
		"emptyAttributeItem":                        engine.Empty[larkface_detection.AttributeItem],
		"emptyRefAttributeItem":                     engine.EmptyRefer[larkface_detection.AttributeItem],
		"refOfAttributeItem":                        engine.ReferOf[larkface_detection.AttributeItem],
		"unRefAttributeItem":                        engine.UnRefer[larkface_detection.AttributeItem],
		"emptyDetectFaceAttributesImageReq":         engine.Empty[larkface_detection.DetectFaceAttributesImageReq],
		"emptyRefDetectFaceAttributesImageReq":      engine.EmptyRefer[larkface_detection.DetectFaceAttributesImageReq],
		"refOfDetectFaceAttributesImageReq":         engine.ReferOf[larkface_detection.DetectFaceAttributesImageReq],
		"unRefDetectFaceAttributesImageReq":         engine.UnRefer[larkface_detection.DetectFaceAttributesImageReq],
		"emptyFaceAttribute":                        engine.Empty[larkface_detection.FaceAttribute],
		"emptyRefFaceAttribute":                     engine.EmptyRefer[larkface_detection.FaceAttribute],
		"refOfFaceAttribute":                        engine.ReferOf[larkface_detection.FaceAttribute],
		"unRefFaceAttribute":                        engine.UnRefer[larkface_detection.FaceAttribute],
		"emptyFaceOcclude":                          engine.Empty[larkface_detection.FaceOcclude],
		"emptyRefFaceOcclude":                       engine.EmptyRefer[larkface_detection.FaceOcclude],
		"refOfFaceOcclude":                          engine.ReferOf[larkface_detection.FaceOcclude],
		"unRefFaceOcclude":                          engine.UnRefer[larkface_detection.FaceOcclude],
		"emptyPoint":                                engine.Empty[larkface_detection.Point],
		"emptyRefPoint":                             engine.EmptyRefer[larkface_detection.Point],
		"refOfPoint":                                engine.ReferOf[larkface_detection.Point],
		"unRefPoint":                                engine.UnRefer[larkface_detection.Point],
		"emptyFaceInfo":                             engine.Empty[larkface_detection.FaceInfo],
		"emptyRefFaceInfo":                          engine.EmptyRefer[larkface_detection.FaceInfo],
		"refOfFaceInfo":                             engine.ReferOf[larkface_detection.FaceInfo],
		"unRefFaceInfo":                             engine.UnRefer[larkface_detection.FaceInfo],
		"emptyFaceQuality":                          engine.Empty[larkface_detection.FaceQuality],
		"emptyRefFaceQuality":                       engine.EmptyRefer[larkface_detection.FaceQuality],
		"refOfFaceQuality":                          engine.ReferOf[larkface_detection.FaceQuality],
		"unRefFaceQuality":                          engine.UnRefer[larkface_detection.FaceQuality],
		"emptyDetectFaceAttributesImageReqBody":     engine.Empty[larkface_detection.DetectFaceAttributesImageReqBody],
		"emptyRefDetectFaceAttributesImageReqBody":  engine.EmptyRefer[larkface_detection.DetectFaceAttributesImageReqBody],
		"refOfDetectFaceAttributesImageReqBody":     engine.ReferOf[larkface_detection.DetectFaceAttributesImageReqBody],
		"unRefDetectFaceAttributesImageReqBody":     engine.UnRefer[larkface_detection.DetectFaceAttributesImageReqBody],
		"emptyDetectFaceAttributesImageResp":        engine.Empty[larkface_detection.DetectFaceAttributesImageResp],
		"emptyRefDetectFaceAttributesImageResp":     engine.EmptyRefer[larkface_detection.DetectFaceAttributesImageResp],
		"refOfDetectFaceAttributesImageResp":        engine.ReferOf[larkface_detection.DetectFaceAttributesImageResp],
		"unRefDetectFaceAttributesImageResp":        engine.UnRefer[larkface_detection.DetectFaceAttributesImageResp],
		"emptyDetectFaceAttributesImageRespData":    engine.Empty[larkface_detection.DetectFaceAttributesImageRespData],
		"emptyRefDetectFaceAttributesImageRespData": engine.EmptyRefer[larkface_detection.DetectFaceAttributesImageRespData],
		"refOfDetectFaceAttributesImageRespData":    engine.ReferOf[larkface_detection.DetectFaceAttributesImageRespData],
		"unRefDetectFaceAttributesImageRespData":    engine.UnRefer[larkface_detection.DetectFaceAttributesImageRespData],
		"emptyFacePose":                             engine.Empty[larkface_detection.FacePose],
		"emptyRefFacePose":                          engine.EmptyRefer[larkface_detection.FacePose],
		"refOfFacePose":                             engine.ReferOf[larkface_detection.FacePose],
		"unRefFacePose":                             engine.UnRefer[larkface_detection.FacePose],
		"emptyFaceDetectionService":                 engine.Empty[larkface_detection.FaceDetectionService],
		"emptyRefFaceDetectionService":              engine.EmptyRefer[larkface_detection.FaceDetectionService],
		"refOfFaceDetectionService":                 engine.ReferOf[larkface_detection.FaceDetectionService],
		"unRefFaceDetectionService":                 engine.UnRefer[larkface_detection.FaceDetectionService],
		"emptyFacePosition":                         engine.Empty[larkface_detection.FacePosition],
		"emptyRefFacePosition":                      engine.EmptyRefer[larkface_detection.FacePosition],
		"refOfFacePosition":                         engine.ReferOf[larkface_detection.FacePosition],
		"unRefFacePosition":                         engine.UnRefer[larkface_detection.FacePosition]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceFace_detection1Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceFace_detection1Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceFace_detection1Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/face_detection/v1"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceFace_detection1Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceFace_detection1Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceFace_detection1Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceFace_detection1Declared
}