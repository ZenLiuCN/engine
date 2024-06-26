// Code generated by define_gene; DO NOT EDIT.
package larkacs

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/event"
	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	"github.com/larksuite/oapi-sdk-go/v3/service/acs/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_acs_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceAcs1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceAcs1Declared = map[string]any{
		"newDeleteVisitorReqBuilder":                  larkacs.NewDeleteVisitorReqBuilder,
		"newUserIdBuilder":                            larkacs.NewUserIdBuilder,
		"UserIdTypeUnionId":                           larkacs.UserIdTypeUnionId,
		"newCreateVisitorReqBodyBuilder":              larkacs.NewCreateVisitorReqBodyBuilder,
		"UserIdTypeGetRuleExternalUnionId":            larkacs.UserIdTypeGetRuleExternalUnionId,
		"UserIdTypeUpdateUserFaceUnionId":             larkacs.UserIdTypeUpdateUserFaceUnionId,
		"newDeviceBindRuleExternalReqBodyBuilder":     larkacs.NewDeviceBindRuleExternalReqBodyBuilder,
		"newDeviceExternalBuilder":                    larkacs.NewDeviceExternalBuilder,
		"UserIdTypeDeleteVisitorOpenId":               larkacs.UserIdTypeDeleteVisitorOpenId,
		"UserIdTypeGetUserFaceUnionId":                larkacs.UserIdTypeGetUserFaceUnionId,
		"UserIdTypeGetUserOpenId":                     larkacs.UserIdTypeGetUserOpenId,
		"UserIdTypeUpdateUserFaceOpenId":              larkacs.UserIdTypeUpdateUserFaceOpenId,
		"newCreateRuleExternalPathReqBodyBuilder":     larkacs.NewCreateRuleExternalPathReqBodyBuilder,
		"newCreateRuleExternalReqBuilder":             larkacs.NewCreateRuleExternalReqBuilder,
		"newOpeningTimeValidDayExternalBuilder":       larkacs.NewOpeningTimeValidDayExternalBuilder,
		"UserIdTypeCreateVisitorUnionId":              larkacs.UserIdTypeCreateVisitorUnionId,
		"UserIdTypeListUserOpenId":                    larkacs.UserIdTypeListUserOpenId,
		"newAccessRecordBuilder":                      larkacs.NewAccessRecordBuilder,
		"newCreateRuleExternalReqBodyBuilder":         larkacs.NewCreateRuleExternalReqBodyBuilder,
		"newFileBuilder":                              larkacs.NewFileBuilder,
		"UserIdTypePatchUserUserId":                   larkacs.UserIdTypePatchUserUserId,
		"newP2UserUpdatedV1Handler":                   larkacs.NewP2UserUpdatedV1Handler,
		"UserIdTypeCreateRuleExternalUserId":          larkacs.UserIdTypeCreateRuleExternalUserId,
		"UserIdTypeDeleteVisitorUserId":               larkacs.UserIdTypeDeleteVisitorUserId,
		"UserIdTypeGetUserFaceOpenId":                 larkacs.UserIdTypeGetUserFaceOpenId,
		"UserIdTypeGetUserUserId":                     larkacs.UserIdTypeGetUserUserId,
		"newCreateVisitorPathReqBodyBuilder":          larkacs.NewCreateVisitorPathReqBodyBuilder,
		"newDeviceBindRuleExternalReqBuilder":         larkacs.NewDeviceBindRuleExternalReqBuilder,
		"newDeviceBuilder":                            larkacs.NewDeviceBuilder,
		"newOpeningTimeExternalBuilder":               larkacs.NewOpeningTimeExternalBuilder,
		"UserIdTypeDeleteVisitorUnionId":              larkacs.UserIdTypeDeleteVisitorUnionId,
		"UserIdTypePatchUserOpenId":                   larkacs.UserIdTypePatchUserOpenId,
		"newGetAccessRecordAccessPhotoReqBuilder":     larkacs.NewGetAccessRecordAccessPhotoReqBuilder,
		"newP2AccessRecordCreatedV1Handler":           larkacs.NewP2AccessRecordCreatedV1Handler,
		"UserIdTypePatchUserUnionId":                  larkacs.UserIdTypePatchUserUnionId,
		"newDepartmentIdBuilder":                      larkacs.NewDepartmentIdBuilder,
		"newPropertyBuilder":                          larkacs.NewPropertyBuilder,
		"UserIdTypeCreateVisitorUserId":               larkacs.UserIdTypeCreateVisitorUserId,
		"newDeviceBindRuleExternalPathReqBodyBuilder": larkacs.NewDeviceBindRuleExternalPathReqBodyBuilder,
		"newFeatureBuilder":                           larkacs.NewFeatureBuilder,
		"newGetUserReqBuilder":                        larkacs.NewGetUserReqBuilder,
		"newListAccessRecordReqBuilder":               larkacs.NewListAccessRecordReqBuilder,
		"newOpeningTimePeriodExternalBuilder":         larkacs.NewOpeningTimePeriodExternalBuilder,
		"newUserBuilder":                              larkacs.NewUserBuilder,
		"newUserExternalBuilder":                      larkacs.NewUserExternalBuilder,
		"UserIdTypeCreateRuleExternalOpenId":          larkacs.UserIdTypeCreateRuleExternalOpenId,
		"UserIdTypeListUserUnionId":                   larkacs.UserIdTypeListUserUnionId,
		"UserIdTypeOpenId":                            larkacs.UserIdTypeOpenId,
		"UserIdTypeUpdateUserFaceUserId":              larkacs.UserIdTypeUpdateUserFaceUserId,
		"newCreateVisitorReqBuilder":                  larkacs.NewCreateVisitorReqBuilder,
		"newDeleteRuleExternalReqBuilder":             larkacs.NewDeleteRuleExternalReqBuilder,
		"newGetUserFaceReqBuilder":                    larkacs.NewGetUserFaceReqBuilder,
		"newPatchUserReqBuilder":                      larkacs.NewPatchUserReqBuilder,
		"UserIdTypeGetRuleExternalUserId":             larkacs.UserIdTypeGetRuleExternalUserId,
		"UserIdTypeGetUserFaceUserId":                 larkacs.UserIdTypeGetUserFaceUserId,
		"newGetRuleExternalReqBuilder":                larkacs.NewGetRuleExternalReqBuilder,
		"newListUserReqBuilder":                       larkacs.NewListUserReqBuilder,
		"UserIdTypeCreateVisitorOpenId":               larkacs.UserIdTypeCreateVisitorOpenId,
		"newRuleBuilder":                              larkacs.NewRuleBuilder,
		"newUpdateUserFaceReqBuilder":                 larkacs.NewUpdateUserFaceReqBuilder,
		"UserIdTypeGetUserUnionId":                    larkacs.UserIdTypeGetUserUnionId,
		"UserIdTypeListUserUserId":                    larkacs.UserIdTypeListUserUserId,
		"New":                                         larkacs.New,
		"UserIdTypeCreateRuleExternalUnionId":         larkacs.UserIdTypeCreateRuleExternalUnionId,
		"UserIdTypeGetRuleExternalOpenId":             larkacs.UserIdTypeGetRuleExternalOpenId,
		"UserIdTypeUserId":                            larkacs.UserIdTypeUserId,

		"emptyDeviceExternal":                    engine.Empty[larkacs.DeviceExternal],
		"emptyRefDeviceExternal":                 engine.EmptyRefer[larkacs.DeviceExternal],
		"refOfDeviceExternal":                    engine.ReferOf[larkacs.DeviceExternal],
		"unRefDeviceExternal":                    engine.UnRefer[larkacs.DeviceExternal],
		"emptyListDeviceResp":                    engine.Empty[larkacs.ListDeviceResp],
		"emptyRefListDeviceResp":                 engine.EmptyRefer[larkacs.ListDeviceResp],
		"refOfListDeviceResp":                    engine.ReferOf[larkacs.ListDeviceResp],
		"unRefListDeviceResp":                    engine.UnRefer[larkacs.ListDeviceResp],
		"emptyUser":                              engine.Empty[larkacs.User],
		"emptyRefUser":                           engine.EmptyRefer[larkacs.User],
		"refOfUser":                              engine.ReferOf[larkacs.User],
		"unRefUser":                              engine.UnRefer[larkacs.User],
		"emptyCreateRuleExternalReqBody":         engine.Empty[larkacs.CreateRuleExternalReqBody],
		"emptyRefCreateRuleExternalReqBody":      engine.EmptyRefer[larkacs.CreateRuleExternalReqBody],
		"refOfCreateRuleExternalReqBody":         engine.ReferOf[larkacs.CreateRuleExternalReqBody],
		"unRefCreateRuleExternalReqBody":         engine.UnRefer[larkacs.CreateRuleExternalReqBody],
		"emptyCreateVisitorReqBody":              engine.Empty[larkacs.CreateVisitorReqBody],
		"emptyRefCreateVisitorReqBody":           engine.EmptyRefer[larkacs.CreateVisitorReqBody],
		"refOfCreateVisitorReqBody":              engine.ReferOf[larkacs.CreateVisitorReqBody],
		"unRefCreateVisitorReqBody":              engine.UnRefer[larkacs.CreateVisitorReqBody],
		"emptyGetAccessRecordAccessPhotoReq":     engine.Empty[larkacs.GetAccessRecordAccessPhotoReq],
		"emptyRefGetAccessRecordAccessPhotoReq":  engine.EmptyRefer[larkacs.GetAccessRecordAccessPhotoReq],
		"refOfGetAccessRecordAccessPhotoReq":     engine.ReferOf[larkacs.GetAccessRecordAccessPhotoReq],
		"unRefGetAccessRecordAccessPhotoReq":     engine.UnRefer[larkacs.GetAccessRecordAccessPhotoReq],
		"emptyGetUserFaceResp":                   engine.Empty[larkacs.GetUserFaceResp],
		"emptyRefGetUserFaceResp":                engine.EmptyRefer[larkacs.GetUserFaceResp],
		"refOfGetUserFaceResp":                   engine.ReferOf[larkacs.GetUserFaceResp],
		"unRefGetUserFaceResp":                   engine.UnRefer[larkacs.GetUserFaceResp],
		"emptyPatchUserResp":                     engine.Empty[larkacs.PatchUserResp],
		"emptyRefPatchUserResp":                  engine.EmptyRefer[larkacs.PatchUserResp],
		"refOfPatchUserResp":                     engine.ReferOf[larkacs.PatchUserResp],
		"unRefPatchUserResp":                     engine.UnRefer[larkacs.PatchUserResp],
		"emptyPatchUserReq":                      engine.Empty[larkacs.PatchUserReq],
		"emptyRefPatchUserReq":                   engine.EmptyRefer[larkacs.PatchUserReq],
		"refOfPatchUserReq":                      engine.ReferOf[larkacs.PatchUserReq],
		"unRefPatchUserReq":                      engine.UnRefer[larkacs.PatchUserReq],
		"emptyDeviceBindRuleExternalReqBody":     engine.Empty[larkacs.DeviceBindRuleExternalReqBody],
		"emptyRefDeviceBindRuleExternalReqBody":  engine.EmptyRefer[larkacs.DeviceBindRuleExternalReqBody],
		"refOfDeviceBindRuleExternalReqBody":     engine.ReferOf[larkacs.DeviceBindRuleExternalReqBody],
		"unRefDeviceBindRuleExternalReqBody":     engine.UnRefer[larkacs.DeviceBindRuleExternalReqBody],
		"emptyListUserRespData":                  engine.Empty[larkacs.ListUserRespData],
		"emptyRefListUserRespData":               engine.EmptyRefer[larkacs.ListUserRespData],
		"refOfListUserRespData":                  engine.ReferOf[larkacs.ListUserRespData],
		"unRefListUserRespData":                  engine.UnRefer[larkacs.ListUserRespData],
		"emptyCreateRuleExternalResp":            engine.Empty[larkacs.CreateRuleExternalResp],
		"emptyRefCreateRuleExternalResp":         engine.EmptyRefer[larkacs.CreateRuleExternalResp],
		"refOfCreateRuleExternalResp":            engine.ReferOf[larkacs.CreateRuleExternalResp],
		"unRefCreateRuleExternalResp":            engine.UnRefer[larkacs.CreateRuleExternalResp],
		"emptyGetAccessRecordAccessPhotoResp":    engine.Empty[larkacs.GetAccessRecordAccessPhotoResp],
		"emptyRefGetAccessRecordAccessPhotoResp": engine.EmptyRefer[larkacs.GetAccessRecordAccessPhotoResp],
		"refOfGetAccessRecordAccessPhotoResp":    engine.ReferOf[larkacs.GetAccessRecordAccessPhotoResp],
		"unRefGetAccessRecordAccessPhotoResp":    engine.UnRefer[larkacs.GetAccessRecordAccessPhotoResp],
		"emptyListAccessRecordResp":              engine.Empty[larkacs.ListAccessRecordResp],
		"emptyRefListAccessRecordResp":           engine.EmptyRefer[larkacs.ListAccessRecordResp],
		"refOfListAccessRecordResp":              engine.ReferOf[larkacs.ListAccessRecordResp],
		"unRefListAccessRecordResp":              engine.UnRefer[larkacs.ListAccessRecordResp],
		"emptyDeleteVisitorResp":                 engine.Empty[larkacs.DeleteVisitorResp],
		"emptyRefDeleteVisitorResp":              engine.EmptyRefer[larkacs.DeleteVisitorResp],
		"refOfDeleteVisitorResp":                 engine.ReferOf[larkacs.DeleteVisitorResp],
		"unRefDeleteVisitorResp":                 engine.UnRefer[larkacs.DeleteVisitorResp],
		"emptyDepartmentId":                      engine.Empty[larkacs.DepartmentId],
		"emptyRefDepartmentId":                   engine.EmptyRefer[larkacs.DepartmentId],
		"refOfDepartmentId":                      engine.ReferOf[larkacs.DepartmentId],
		"unRefDepartmentId":                      engine.UnRefer[larkacs.DepartmentId],
		"emptyListUserIterator":                  engine.Empty[larkacs.ListUserIterator],
		"emptyRefListUserIterator":               engine.EmptyRefer[larkacs.ListUserIterator],
		"refOfListUserIterator":                  engine.ReferOf[larkacs.ListUserIterator],
		"unRefListUserIterator":                  engine.UnRefer[larkacs.ListUserIterator],
		"emptyFile":                              engine.Empty[larkacs.File],
		"emptyRefFile":                           engine.EmptyRefer[larkacs.File],
		"refOfFile":                              engine.ReferOf[larkacs.File],
		"unRefFile":                              engine.UnRefer[larkacs.File],
		"emptyGetRuleExternalReq":                engine.Empty[larkacs.GetRuleExternalReq],
		"emptyRefGetRuleExternalReq":             engine.EmptyRefer[larkacs.GetRuleExternalReq],
		"refOfGetRuleExternalReq":                engine.ReferOf[larkacs.GetRuleExternalReq],
		"unRefGetRuleExternalReq":                engine.UnRefer[larkacs.GetRuleExternalReq],
		"emptyGetUserResp":                       engine.Empty[larkacs.GetUserResp],
		"emptyRefGetUserResp":                    engine.EmptyRefer[larkacs.GetUserResp],
		"refOfGetUserResp":                       engine.ReferOf[larkacs.GetUserResp],
		"unRefGetUserResp":                       engine.UnRefer[larkacs.GetUserResp],
		"emptyUpdateUserFaceResp":                engine.Empty[larkacs.UpdateUserFaceResp],
		"emptyRefUpdateUserFaceResp":             engine.EmptyRefer[larkacs.UpdateUserFaceResp],
		"refOfUpdateUserFaceResp":                engine.ReferOf[larkacs.UpdateUserFaceResp],
		"unRefUpdateUserFaceResp":                engine.UnRefer[larkacs.UpdateUserFaceResp],
		"emptyAccessRecord":                      engine.Empty[larkacs.AccessRecord],
		"emptyRefAccessRecord":                   engine.EmptyRefer[larkacs.AccessRecord],
		"refOfAccessRecord":                      engine.ReferOf[larkacs.AccessRecord],
		"unRefAccessRecord":                      engine.UnRefer[larkacs.AccessRecord],
		"emptyDevice":                            engine.Empty[larkacs.Device],
		"emptyRefDevice":                         engine.EmptyRefer[larkacs.Device],
		"refOfDevice":                            engine.ReferOf[larkacs.Device],
		"unRefDevice":                            engine.UnRefer[larkacs.Device],
		"emptyGetUserFaceReq":                    engine.Empty[larkacs.GetUserFaceReq],
		"emptyRefGetUserFaceReq":                 engine.EmptyRefer[larkacs.GetUserFaceReq],
		"refOfGetUserFaceReq":                    engine.ReferOf[larkacs.GetUserFaceReq],
		"unRefGetUserFaceReq":                    engine.UnRefer[larkacs.GetUserFaceReq],
		"emptyListAccessRecordReq":               engine.Empty[larkacs.ListAccessRecordReq],
		"emptyRefListAccessRecordReq":            engine.EmptyRefer[larkacs.ListAccessRecordReq],
		"refOfListAccessRecordReq":               engine.ReferOf[larkacs.ListAccessRecordReq],
		"unRefListAccessRecordReq":               engine.UnRefer[larkacs.ListAccessRecordReq],
		"emptyP2UserUpdatedV1":                   engine.Empty[larkacs.P2UserUpdatedV1],
		"emptyRefP2UserUpdatedV1":                engine.EmptyRefer[larkacs.P2UserUpdatedV1],
		"refOfP2UserUpdatedV1":                   engine.ReferOf[larkacs.P2UserUpdatedV1],
		"unRefP2UserUpdatedV1":                   engine.UnRefer[larkacs.P2UserUpdatedV1],
		"emptyGetUserRespData":                   engine.Empty[larkacs.GetUserRespData],
		"emptyRefGetUserRespData":                engine.EmptyRefer[larkacs.GetUserRespData],
		"refOfGetUserRespData":                   engine.ReferOf[larkacs.GetUserRespData],
		"unRefGetUserRespData":                   engine.UnRefer[larkacs.GetUserRespData],
		"emptyListAccessRecordIterator":          engine.Empty[larkacs.ListAccessRecordIterator],
		"emptyRefListAccessRecordIterator":       engine.EmptyRefer[larkacs.ListAccessRecordIterator],
		"refOfListAccessRecordIterator":          engine.ReferOf[larkacs.ListAccessRecordIterator],
		"unRefListAccessRecordIterator":          engine.UnRefer[larkacs.ListAccessRecordIterator],
		"emptyOpeningTimeValidDayExternal":       engine.Empty[larkacs.OpeningTimeValidDayExternal],
		"emptyRefOpeningTimeValidDayExternal":    engine.EmptyRefer[larkacs.OpeningTimeValidDayExternal],
		"refOfOpeningTimeValidDayExternal":       engine.ReferOf[larkacs.OpeningTimeValidDayExternal],
		"unRefOpeningTimeValidDayExternal":       engine.UnRefer[larkacs.OpeningTimeValidDayExternal],
		"emptyP2UserUpdatedV1Data":               engine.Empty[larkacs.P2UserUpdatedV1Data],
		"emptyRefP2UserUpdatedV1Data":            engine.EmptyRefer[larkacs.P2UserUpdatedV1Data],
		"refOfP2UserUpdatedV1Data":               engine.ReferOf[larkacs.P2UserUpdatedV1Data],
		"unRefP2UserUpdatedV1Data":               engine.UnRefer[larkacs.P2UserUpdatedV1Data],
		"emptyProperty":                          engine.Empty[larkacs.Property],
		"emptyRefProperty":                       engine.EmptyRefer[larkacs.Property],
		"refOfProperty":                          engine.ReferOf[larkacs.Property],
		"unRefProperty":                          engine.UnRefer[larkacs.Property],
		"emptyRule":                              engine.Empty[larkacs.Rule],
		"emptyRefRule":                           engine.EmptyRefer[larkacs.Rule],
		"refOfRule":                              engine.ReferOf[larkacs.Rule],
		"unRefRule":                              engine.UnRefer[larkacs.Rule],
		"emptyCreateVisitorReq":                  engine.Empty[larkacs.CreateVisitorReq],
		"emptyRefCreateVisitorReq":               engine.EmptyRefer[larkacs.CreateVisitorReq],
		"refOfCreateVisitorReq":                  engine.ReferOf[larkacs.CreateVisitorReq],
		"unRefCreateVisitorReq":                  engine.UnRefer[larkacs.CreateVisitorReq],
		"emptyCreateVisitorResp":                 engine.Empty[larkacs.CreateVisitorResp],
		"emptyRefCreateVisitorResp":              engine.EmptyRefer[larkacs.CreateVisitorResp],
		"refOfCreateVisitorResp":                 engine.ReferOf[larkacs.CreateVisitorResp],
		"unRefCreateVisitorResp":                 engine.UnRefer[larkacs.CreateVisitorResp],
		"emptyCreateVisitorRespData":             engine.Empty[larkacs.CreateVisitorRespData],
		"emptyRefCreateVisitorRespData":          engine.EmptyRefer[larkacs.CreateVisitorRespData],
		"refOfCreateVisitorRespData":             engine.ReferOf[larkacs.CreateVisitorRespData],
		"unRefCreateVisitorRespData":             engine.UnRefer[larkacs.CreateVisitorRespData],
		"emptyDeleteRuleExternalResp":            engine.Empty[larkacs.DeleteRuleExternalResp],
		"emptyRefDeleteRuleExternalResp":         engine.EmptyRefer[larkacs.DeleteRuleExternalResp],
		"refOfDeleteRuleExternalResp":            engine.ReferOf[larkacs.DeleteRuleExternalResp],
		"unRefDeleteRuleExternalResp":            engine.UnRefer[larkacs.DeleteRuleExternalResp],
		"emptyGetRuleExternalResp":               engine.Empty[larkacs.GetRuleExternalResp],
		"emptyRefGetRuleExternalResp":            engine.EmptyRefer[larkacs.GetRuleExternalResp],
		"refOfGetRuleExternalResp":               engine.ReferOf[larkacs.GetRuleExternalResp],
		"unRefGetRuleExternalResp":               engine.UnRefer[larkacs.GetRuleExternalResp],
		"emptyP2AccessRecordCreatedV1":           engine.Empty[larkacs.P2AccessRecordCreatedV1],
		"emptyRefP2AccessRecordCreatedV1":        engine.EmptyRefer[larkacs.P2AccessRecordCreatedV1],
		"refOfP2AccessRecordCreatedV1":           engine.ReferOf[larkacs.P2AccessRecordCreatedV1],
		"unRefP2AccessRecordCreatedV1":           engine.UnRefer[larkacs.P2AccessRecordCreatedV1],
		"emptyCreateRuleExternalRespData":        engine.Empty[larkacs.CreateRuleExternalRespData],
		"emptyRefCreateRuleExternalRespData":     engine.EmptyRefer[larkacs.CreateRuleExternalRespData],
		"refOfCreateRuleExternalRespData":        engine.ReferOf[larkacs.CreateRuleExternalRespData],
		"unRefCreateRuleExternalRespData":        engine.UnRefer[larkacs.CreateRuleExternalRespData],
		"emptyDeviceBindRuleExternalReq":         engine.Empty[larkacs.DeviceBindRuleExternalReq],
		"emptyRefDeviceBindRuleExternalReq":      engine.EmptyRefer[larkacs.DeviceBindRuleExternalReq],
		"refOfDeviceBindRuleExternalReq":         engine.ReferOf[larkacs.DeviceBindRuleExternalReq],
		"unRefDeviceBindRuleExternalReq":         engine.UnRefer[larkacs.DeviceBindRuleExternalReq],
		"emptyFeature":                           engine.Empty[larkacs.Feature],
		"emptyRefFeature":                        engine.EmptyRefer[larkacs.Feature],
		"refOfFeature":                           engine.ReferOf[larkacs.Feature],
		"unRefFeature":                           engine.UnRefer[larkacs.Feature],
		"emptyOpeningTimeExternal":               engine.Empty[larkacs.OpeningTimeExternal],
		"emptyRefOpeningTimeExternal":            engine.EmptyRefer[larkacs.OpeningTimeExternal],
		"refOfOpeningTimeExternal":               engine.ReferOf[larkacs.OpeningTimeExternal],
		"unRefOpeningTimeExternal":               engine.UnRefer[larkacs.OpeningTimeExternal],
		"emptyUserId":                            engine.Empty[larkacs.UserId],
		"emptyRefUserId":                         engine.EmptyRefer[larkacs.UserId],
		"refOfUserId":                            engine.ReferOf[larkacs.UserId],
		"unRefUserId":                            engine.UnRefer[larkacs.UserId],
		"emptyGetRuleExternalRespData":           engine.Empty[larkacs.GetRuleExternalRespData],
		"emptyRefGetRuleExternalRespData":        engine.EmptyRefer[larkacs.GetRuleExternalRespData],
		"refOfGetRuleExternalRespData":           engine.ReferOf[larkacs.GetRuleExternalRespData],
		"unRefGetRuleExternalRespData":           engine.UnRefer[larkacs.GetRuleExternalRespData],
		"emptyGetUserReq":                        engine.Empty[larkacs.GetUserReq],
		"emptyRefGetUserReq":                     engine.EmptyRefer[larkacs.GetUserReq],
		"refOfGetUserReq":                        engine.ReferOf[larkacs.GetUserReq],
		"unRefGetUserReq":                        engine.UnRefer[larkacs.GetUserReq],
		"emptyListAccessRecordRespData":          engine.Empty[larkacs.ListAccessRecordRespData],
		"emptyRefListAccessRecordRespData":       engine.EmptyRefer[larkacs.ListAccessRecordRespData],
		"refOfListAccessRecordRespData":          engine.ReferOf[larkacs.ListAccessRecordRespData],
		"unRefListAccessRecordRespData":          engine.UnRefer[larkacs.ListAccessRecordRespData],
		"emptyListUserReq":                       engine.Empty[larkacs.ListUserReq],
		"emptyRefListUserReq":                    engine.EmptyRefer[larkacs.ListUserReq],
		"refOfListUserReq":                       engine.ReferOf[larkacs.ListUserReq],
		"unRefListUserReq":                       engine.UnRefer[larkacs.ListUserReq],
		"emptyListUserResp":                      engine.Empty[larkacs.ListUserResp],
		"emptyRefListUserResp":                   engine.EmptyRefer[larkacs.ListUserResp],
		"refOfListUserResp":                      engine.ReferOf[larkacs.ListUserResp],
		"unRefListUserResp":                      engine.UnRefer[larkacs.ListUserResp],
		"emptyV1":                                engine.Empty[larkacs.V1],
		"emptyRefV1":                             engine.EmptyRefer[larkacs.V1],
		"refOfV1":                                engine.ReferOf[larkacs.V1],
		"unRefV1":                                engine.UnRefer[larkacs.V1],
		"emptyUpdateUserFaceReq":                 engine.Empty[larkacs.UpdateUserFaceReq],
		"emptyRefUpdateUserFaceReq":              engine.EmptyRefer[larkacs.UpdateUserFaceReq],
		"refOfUpdateUserFaceReq":                 engine.ReferOf[larkacs.UpdateUserFaceReq],
		"unRefUpdateUserFaceReq":                 engine.UnRefer[larkacs.UpdateUserFaceReq],
		"emptyCreateRuleExternalReq":             engine.Empty[larkacs.CreateRuleExternalReq],
		"emptyRefCreateRuleExternalReq":          engine.EmptyRefer[larkacs.CreateRuleExternalReq],
		"refOfCreateRuleExternalReq":             engine.ReferOf[larkacs.CreateRuleExternalReq],
		"unRefCreateRuleExternalReq":             engine.UnRefer[larkacs.CreateRuleExternalReq],
		"emptyDeleteRuleExternalReq":             engine.Empty[larkacs.DeleteRuleExternalReq],
		"emptyRefDeleteRuleExternalReq":          engine.EmptyRefer[larkacs.DeleteRuleExternalReq],
		"refOfDeleteRuleExternalReq":             engine.ReferOf[larkacs.DeleteRuleExternalReq],
		"unRefDeleteRuleExternalReq":             engine.UnRefer[larkacs.DeleteRuleExternalReq],
		"emptyP2AccessRecordCreatedV1Data":       engine.Empty[larkacs.P2AccessRecordCreatedV1Data],
		"emptyRefP2AccessRecordCreatedV1Data":    engine.EmptyRefer[larkacs.P2AccessRecordCreatedV1Data],
		"refOfP2AccessRecordCreatedV1Data":       engine.ReferOf[larkacs.P2AccessRecordCreatedV1Data],
		"unRefP2AccessRecordCreatedV1Data":       engine.UnRefer[larkacs.P2AccessRecordCreatedV1Data],
		"emptyUserExternal":                      engine.Empty[larkacs.UserExternal],
		"emptyRefUserExternal":                   engine.EmptyRefer[larkacs.UserExternal],
		"refOfUserExternal":                      engine.ReferOf[larkacs.UserExternal],
		"unRefUserExternal":                      engine.UnRefer[larkacs.UserExternal],
		"emptyDeleteVisitorReq":                  engine.Empty[larkacs.DeleteVisitorReq],
		"emptyRefDeleteVisitorReq":               engine.EmptyRefer[larkacs.DeleteVisitorReq],
		"refOfDeleteVisitorReq":                  engine.ReferOf[larkacs.DeleteVisitorReq],
		"unRefDeleteVisitorReq":                  engine.UnRefer[larkacs.DeleteVisitorReq],
		"emptyDeviceBindRuleExternalResp":        engine.Empty[larkacs.DeviceBindRuleExternalResp],
		"emptyRefDeviceBindRuleExternalResp":     engine.EmptyRefer[larkacs.DeviceBindRuleExternalResp],
		"refOfDeviceBindRuleExternalResp":        engine.ReferOf[larkacs.DeviceBindRuleExternalResp],
		"unRefDeviceBindRuleExternalResp":        engine.UnRefer[larkacs.DeviceBindRuleExternalResp],
		"emptyListDeviceRespData":                engine.Empty[larkacs.ListDeviceRespData],
		"emptyRefListDeviceRespData":             engine.EmptyRefer[larkacs.ListDeviceRespData],
		"refOfListDeviceRespData":                engine.ReferOf[larkacs.ListDeviceRespData],
		"unRefListDeviceRespData":                engine.UnRefer[larkacs.ListDeviceRespData],
		"emptyOpeningTimePeriodExternal":         engine.Empty[larkacs.OpeningTimePeriodExternal],
		"emptyRefOpeningTimePeriodExternal":      engine.EmptyRefer[larkacs.OpeningTimePeriodExternal],
		"refOfOpeningTimePeriodExternal":         engine.ReferOf[larkacs.OpeningTimePeriodExternal],
		"unRefOpeningTimePeriodExternal":         engine.UnRefer[larkacs.OpeningTimePeriodExternal]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceAcs1Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceAcs1Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceAcs1Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/acs/v1"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceAcs1Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceAcs1Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceAcs1Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceAcs1Declared
}
