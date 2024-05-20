// Code generated by define_gene; DO NOT EDIT.
package larksecurity_and_compliance

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	"github.com/larksuite/oapi-sdk-go/v3/service/security_and_compliance/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_security_and_compliance_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceSecurity_and_compliance1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceSecurity_and_compliance1Declared = map[string]any{
		"newUserBuilder":                          larksecurity_and_compliance.NewUserBuilder,
		"newUserIdBuilder":                        larksecurity_and_compliance.NewUserIdBuilder,
		"newAppDlpExecuteLogBuilder":              larksecurity_and_compliance.NewAppDlpExecuteLogBuilder,
		"newGwCommonBuilder":                      larksecurity_and_compliance.NewGwCommonBuilder,
		"newGwRequestBuilder":                     larksecurity_and_compliance.NewGwRequestBuilder,
		"newKeyPersonBuilder":                     larksecurity_and_compliance.NewKeyPersonBuilder,
		"newMigrationEntityBuilder":               larksecurity_and_compliance.NewMigrationEntityBuilder,
		"newNotificationBuilder":                  larksecurity_and_compliance.NewNotificationBuilder,
		"newVaultExportFileBuilder":               larksecurity_and_compliance.NewVaultExportFileBuilder,
		"newTenantBuilder":                        larksecurity_and_compliance.NewTenantBuilder,
		"newTimeRangeBuilder":                     larksecurity_and_compliance.NewTimeRangeBuilder,
		"newAdminLogBuilder":                      larksecurity_and_compliance.NewAdminLogBuilder,
		"newDepartmentIdBuilder":                  larksecurity_and_compliance.NewDepartmentIdBuilder,
		"newDeviceApplyRecordBuilder":             larksecurity_and_compliance.NewDeviceApplyRecordBuilder,
		"newDlpExecuteEvidenceBuilder":            larksecurity_and_compliance.NewDlpExecuteEvidenceBuilder,
		"newDlpExecuteLogBuilder":                 larksecurity_and_compliance.NewDlpExecuteLogBuilder,
		"newEmailBuilder":                         larksecurity_and_compliance.NewEmailBuilder,
		"newOpenapiLogBuilder":                    larksecurity_and_compliance.NewOpenapiLogBuilder,
		"newVaultTaskBuilder":                     larksecurity_and_compliance.NewVaultTaskBuilder,
		"newEmailFilterBuilder":                   larksecurity_and_compliance.NewEmailFilterBuilder,
		"newMessageBuilder":                       larksecurity_and_compliance.NewMessageBuilder,
		"New":                                     larksecurity_and_compliance.New,
		"newCreateMigrationItemsBuilder":          larksecurity_and_compliance.NewCreateMigrationItemsBuilder,
		"newDeviceRecordBuilder":                  larksecurity_and_compliance.NewDeviceRecordBuilder,
		"newDirectoryItemsBuilder":                larksecurity_and_compliance.NewDirectoryItemsBuilder,
		"newDocumentBuilder":                      larksecurity_and_compliance.NewDocumentBuilder,
		"newDownloadTokenBuilder":                 larksecurity_and_compliance.NewDownloadTokenBuilder,
		"newSimpleUserBuilder":                    larksecurity_and_compliance.NewSimpleUserBuilder,
		"newCreateMigrationEntityBuilder":         larksecurity_and_compliance.NewCreateMigrationEntityBuilder,
		"newDataArchivingUserBuilder":             larksecurity_and_compliance.NewDataArchivingUserBuilder,
		"newDlpEvidenceDetailBuilder":             larksecurity_and_compliance.NewDlpEvidenceDetailBuilder,
		"newDlpPolicyHitProofBuilder":             larksecurity_and_compliance.NewDlpPolicyHitProofBuilder,
		"newListDataOpenapiLogReqBuilder":         larksecurity_and_compliance.NewListDataOpenapiLogReqBuilder,
		"newMigrationItemsBuilder":                larksecurity_and_compliance.NewMigrationItemsBuilder,
		"newTenantThirdPartyEncryptionAppBuilder": larksecurity_and_compliance.NewTenantThirdPartyEncryptionAppBuilder,
		"newDataArchivingMessageStructBuilder":    larksecurity_and_compliance.NewDataArchivingMessageStructBuilder,
		"newDataArchivingUserStructBuilder":       larksecurity_and_compliance.NewDataArchivingUserStructBuilder,
		"newGwResponseBuilder":                    larksecurity_and_compliance.NewGwResponseBuilder,
		"newListOpenapiLogRequestBuilder":         larksecurity_and_compliance.NewListOpenapiLogRequestBuilder,
		"newParamBuilder":                         larksecurity_and_compliance.NewParamBuilder,
		"newSecurityLogErrorBuilder":              larksecurity_and_compliance.NewSecurityLogErrorBuilder,
		"newUserMigrationBuilder":                 larksecurity_and_compliance.NewUserMigrationBuilder,
		"newDataArchivingMessageBuilder":          larksecurity_and_compliance.NewDataArchivingMessageBuilder,
		"newDlpDetectModeProofContextBuilder":     larksecurity_and_compliance.NewDlpDetectModeProofContextBuilder,
		"newDlpHitPolicyBuilder":                  larksecurity_and_compliance.NewDlpHitPolicyBuilder,
		"newDlpProofContextBuilder":               larksecurity_and_compliance.NewDlpProofContextBuilder,
		"newOpenapiLogDetailBuilder":              larksecurity_and_compliance.NewOpenapiLogDetailBuilder,
		"newTaskStatusBuilder":                    larksecurity_and_compliance.NewTaskStatusBuilder,

		"emptyEmailFilter":                      engine.Empty[larksecurity_and_compliance.EmailFilter],
		"emptyRefEmailFilter":                   engine.EmptyRefer[larksecurity_and_compliance.EmailFilter],
		"refOfEmailFilter":                      engine.ReferOf[larksecurity_and_compliance.EmailFilter],
		"unRefEmailFilter":                      engine.UnRefer[larksecurity_and_compliance.EmailFilter],
		"emptyDocument":                         engine.Empty[larksecurity_and_compliance.Document],
		"emptyRefDocument":                      engine.EmptyRefer[larksecurity_and_compliance.Document],
		"refOfDocument":                         engine.ReferOf[larksecurity_and_compliance.Document],
		"unRefDocument":                         engine.UnRefer[larksecurity_and_compliance.Document],
		"emptyParam":                            engine.Empty[larksecurity_and_compliance.Param],
		"emptyRefParam":                         engine.EmptyRefer[larksecurity_and_compliance.Param],
		"refOfParam":                            engine.ReferOf[larksecurity_and_compliance.Param],
		"unRefParam":                            engine.UnRefer[larksecurity_and_compliance.Param],
		"emptyCreateMigrationEntity":            engine.Empty[larksecurity_and_compliance.CreateMigrationEntity],
		"emptyRefCreateMigrationEntity":         engine.EmptyRefer[larksecurity_and_compliance.CreateMigrationEntity],
		"refOfCreateMigrationEntity":            engine.ReferOf[larksecurity_and_compliance.CreateMigrationEntity],
		"unRefCreateMigrationEntity":            engine.UnRefer[larksecurity_and_compliance.CreateMigrationEntity],
		"emptyDirectoryItems":                   engine.Empty[larksecurity_and_compliance.DirectoryItems],
		"emptyRefDirectoryItems":                engine.EmptyRefer[larksecurity_and_compliance.DirectoryItems],
		"refOfDirectoryItems":                   engine.ReferOf[larksecurity_and_compliance.DirectoryItems],
		"unRefDirectoryItems":                   engine.UnRefer[larksecurity_and_compliance.DirectoryItems],
		"emptyVaultExportFile":                  engine.Empty[larksecurity_and_compliance.VaultExportFile],
		"emptyRefVaultExportFile":               engine.EmptyRefer[larksecurity_and_compliance.VaultExportFile],
		"refOfVaultExportFile":                  engine.ReferOf[larksecurity_and_compliance.VaultExportFile],
		"unRefVaultExportFile":                  engine.UnRefer[larksecurity_and_compliance.VaultExportFile],
		"emptyDataArchivingMessageStruct":       engine.Empty[larksecurity_and_compliance.DataArchivingMessageStruct],
		"emptyRefDataArchivingMessageStruct":    engine.EmptyRefer[larksecurity_and_compliance.DataArchivingMessageStruct],
		"refOfDataArchivingMessageStruct":       engine.ReferOf[larksecurity_and_compliance.DataArchivingMessageStruct],
		"unRefDataArchivingMessageStruct":       engine.UnRefer[larksecurity_and_compliance.DataArchivingMessageStruct],
		"emptyDlpPolicyHitProof":                engine.Empty[larksecurity_and_compliance.DlpPolicyHitProof],
		"emptyRefDlpPolicyHitProof":             engine.EmptyRefer[larksecurity_and_compliance.DlpPolicyHitProof],
		"refOfDlpPolicyHitProof":                engine.ReferOf[larksecurity_and_compliance.DlpPolicyHitProof],
		"unRefDlpPolicyHitProof":                engine.UnRefer[larksecurity_and_compliance.DlpPolicyHitProof],
		"emptyDownloadToken":                    engine.Empty[larksecurity_and_compliance.DownloadToken],
		"emptyRefDownloadToken":                 engine.EmptyRefer[larksecurity_and_compliance.DownloadToken],
		"refOfDownloadToken":                    engine.ReferOf[larksecurity_and_compliance.DownloadToken],
		"unRefDownloadToken":                    engine.UnRefer[larksecurity_and_compliance.DownloadToken],
		"emptyDeviceApplyRecord":                engine.Empty[larksecurity_and_compliance.DeviceApplyRecord],
		"emptyRefDeviceApplyRecord":             engine.EmptyRefer[larksecurity_and_compliance.DeviceApplyRecord],
		"refOfDeviceApplyRecord":                engine.ReferOf[larksecurity_and_compliance.DeviceApplyRecord],
		"unRefDeviceApplyRecord":                engine.UnRefer[larksecurity_and_compliance.DeviceApplyRecord],
		"emptyUserId":                           engine.Empty[larksecurity_and_compliance.UserId],
		"emptyRefUserId":                        engine.EmptyRefer[larksecurity_and_compliance.UserId],
		"refOfUserId":                           engine.ReferOf[larksecurity_and_compliance.UserId],
		"unRefUserId":                           engine.UnRefer[larksecurity_and_compliance.UserId],
		"emptyListDataOpenapiLogReq":            engine.Empty[larksecurity_and_compliance.ListDataOpenapiLogReq],
		"emptyRefListDataOpenapiLogReq":         engine.EmptyRefer[larksecurity_and_compliance.ListDataOpenapiLogReq],
		"refOfListDataOpenapiLogReq":            engine.ReferOf[larksecurity_and_compliance.ListDataOpenapiLogReq],
		"unRefListDataOpenapiLogReq":            engine.UnRefer[larksecurity_and_compliance.ListDataOpenapiLogReq],
		"emptySimpleUser":                       engine.Empty[larksecurity_and_compliance.SimpleUser],
		"emptyRefSimpleUser":                    engine.EmptyRefer[larksecurity_and_compliance.SimpleUser],
		"refOfSimpleUser":                       engine.ReferOf[larksecurity_and_compliance.SimpleUser],
		"unRefSimpleUser":                       engine.UnRefer[larksecurity_and_compliance.SimpleUser],
		"emptyTenant":                           engine.Empty[larksecurity_and_compliance.Tenant],
		"emptyRefTenant":                        engine.EmptyRefer[larksecurity_and_compliance.Tenant],
		"refOfTenant":                           engine.ReferOf[larksecurity_and_compliance.Tenant],
		"unRefTenant":                           engine.UnRefer[larksecurity_and_compliance.Tenant],
		"emptyAdminLog":                         engine.Empty[larksecurity_and_compliance.AdminLog],
		"emptyRefAdminLog":                      engine.EmptyRefer[larksecurity_and_compliance.AdminLog],
		"refOfAdminLog":                         engine.ReferOf[larksecurity_and_compliance.AdminLog],
		"unRefAdminLog":                         engine.UnRefer[larksecurity_and_compliance.AdminLog],
		"emptyDeviceRecord":                     engine.Empty[larksecurity_and_compliance.DeviceRecord],
		"emptyRefDeviceRecord":                  engine.EmptyRefer[larksecurity_and_compliance.DeviceRecord],
		"refOfDeviceRecord":                     engine.ReferOf[larksecurity_and_compliance.DeviceRecord],
		"unRefDeviceRecord":                     engine.UnRefer[larksecurity_and_compliance.DeviceRecord],
		"emptyDlpExecuteEvidence":               engine.Empty[larksecurity_and_compliance.DlpExecuteEvidence],
		"emptyRefDlpExecuteEvidence":            engine.EmptyRefer[larksecurity_and_compliance.DlpExecuteEvidence],
		"refOfDlpExecuteEvidence":               engine.ReferOf[larksecurity_and_compliance.DlpExecuteEvidence],
		"unRefDlpExecuteEvidence":               engine.UnRefer[larksecurity_and_compliance.DlpExecuteEvidence],
		"emptyV1":                               engine.Empty[larksecurity_and_compliance.V1],
		"emptyRefV1":                            engine.EmptyRefer[larksecurity_and_compliance.V1],
		"refOfV1":                               engine.ReferOf[larksecurity_and_compliance.V1],
		"unRefV1":                               engine.UnRefer[larksecurity_and_compliance.V1],
		"emptyDataArchivingMessage":             engine.Empty[larksecurity_and_compliance.DataArchivingMessage],
		"emptyRefDataArchivingMessage":          engine.EmptyRefer[larksecurity_and_compliance.DataArchivingMessage],
		"refOfDataArchivingMessage":             engine.ReferOf[larksecurity_and_compliance.DataArchivingMessage],
		"unRefDataArchivingMessage":             engine.UnRefer[larksecurity_and_compliance.DataArchivingMessage],
		"emptyDataArchivingUser":                engine.Empty[larksecurity_and_compliance.DataArchivingUser],
		"emptyRefDataArchivingUser":             engine.EmptyRefer[larksecurity_and_compliance.DataArchivingUser],
		"refOfDataArchivingUser":                engine.ReferOf[larksecurity_and_compliance.DataArchivingUser],
		"unRefDataArchivingUser":                engine.UnRefer[larksecurity_and_compliance.DataArchivingUser],
		"emptyDlpProofContext":                  engine.Empty[larksecurity_and_compliance.DlpProofContext],
		"emptyRefDlpProofContext":               engine.EmptyRefer[larksecurity_and_compliance.DlpProofContext],
		"refOfDlpProofContext":                  engine.ReferOf[larksecurity_and_compliance.DlpProofContext],
		"unRefDlpProofContext":                  engine.UnRefer[larksecurity_and_compliance.DlpProofContext],
		"emptyListDataOpenapiLogRespData":       engine.Empty[larksecurity_and_compliance.ListDataOpenapiLogRespData],
		"emptyRefListDataOpenapiLogRespData":    engine.EmptyRefer[larksecurity_and_compliance.ListDataOpenapiLogRespData],
		"refOfListDataOpenapiLogRespData":       engine.ReferOf[larksecurity_and_compliance.ListDataOpenapiLogRespData],
		"unRefListDataOpenapiLogRespData":       engine.UnRefer[larksecurity_and_compliance.ListDataOpenapiLogRespData],
		"emptyDepartmentId":                     engine.Empty[larksecurity_and_compliance.DepartmentId],
		"emptyRefDepartmentId":                  engine.EmptyRefer[larksecurity_and_compliance.DepartmentId],
		"refOfDepartmentId":                     engine.ReferOf[larksecurity_and_compliance.DepartmentId],
		"unRefDepartmentId":                     engine.UnRefer[larksecurity_and_compliance.DepartmentId],
		"emptyVaultTask":                        engine.Empty[larksecurity_and_compliance.VaultTask],
		"emptyRefVaultTask":                     engine.EmptyRefer[larksecurity_and_compliance.VaultTask],
		"refOfVaultTask":                        engine.ReferOf[larksecurity_and_compliance.VaultTask],
		"unRefVaultTask":                        engine.UnRefer[larksecurity_and_compliance.VaultTask],
		"emptyMigrationItems":                   engine.Empty[larksecurity_and_compliance.MigrationItems],
		"emptyRefMigrationItems":                engine.EmptyRefer[larksecurity_and_compliance.MigrationItems],
		"refOfMigrationItems":                   engine.ReferOf[larksecurity_and_compliance.MigrationItems],
		"unRefMigrationItems":                   engine.UnRefer[larksecurity_and_compliance.MigrationItems],
		"emptyDlpDetectModeProofContext":        engine.Empty[larksecurity_and_compliance.DlpDetectModeProofContext],
		"emptyRefDlpDetectModeProofContext":     engine.EmptyRefer[larksecurity_and_compliance.DlpDetectModeProofContext],
		"refOfDlpDetectModeProofContext":        engine.ReferOf[larksecurity_and_compliance.DlpDetectModeProofContext],
		"unRefDlpDetectModeProofContext":        engine.UnRefer[larksecurity_and_compliance.DlpDetectModeProofContext],
		"emptyOpenapiLog":                       engine.Empty[larksecurity_and_compliance.OpenapiLog],
		"emptyRefOpenapiLog":                    engine.EmptyRefer[larksecurity_and_compliance.OpenapiLog],
		"refOfOpenapiLog":                       engine.ReferOf[larksecurity_and_compliance.OpenapiLog],
		"unRefOpenapiLog":                       engine.UnRefer[larksecurity_and_compliance.OpenapiLog],
		"emptyDlpExecuteLog":                    engine.Empty[larksecurity_and_compliance.DlpExecuteLog],
		"emptyRefDlpExecuteLog":                 engine.EmptyRefer[larksecurity_and_compliance.DlpExecuteLog],
		"refOfDlpExecuteLog":                    engine.ReferOf[larksecurity_and_compliance.DlpExecuteLog],
		"unRefDlpExecuteLog":                    engine.UnRefer[larksecurity_and_compliance.DlpExecuteLog],
		"emptyGwResponse":                       engine.Empty[larksecurity_and_compliance.GwResponse],
		"emptyRefGwResponse":                    engine.EmptyRefer[larksecurity_and_compliance.GwResponse],
		"refOfGwResponse":                       engine.ReferOf[larksecurity_and_compliance.GwResponse],
		"unRefGwResponse":                       engine.UnRefer[larksecurity_and_compliance.GwResponse],
		"emptyMigrationEntity":                  engine.Empty[larksecurity_and_compliance.MigrationEntity],
		"emptyRefMigrationEntity":               engine.EmptyRefer[larksecurity_and_compliance.MigrationEntity],
		"refOfMigrationEntity":                  engine.ReferOf[larksecurity_and_compliance.MigrationEntity],
		"unRefMigrationEntity":                  engine.UnRefer[larksecurity_and_compliance.MigrationEntity],
		"emptyDlpEvidenceDetail":                engine.Empty[larksecurity_and_compliance.DlpEvidenceDetail],
		"emptyRefDlpEvidenceDetail":             engine.EmptyRefer[larksecurity_and_compliance.DlpEvidenceDetail],
		"refOfDlpEvidenceDetail":                engine.ReferOf[larksecurity_and_compliance.DlpEvidenceDetail],
		"unRefDlpEvidenceDetail":                engine.UnRefer[larksecurity_and_compliance.DlpEvidenceDetail],
		"emptyUser":                             engine.Empty[larksecurity_and_compliance.User],
		"emptyRefUser":                          engine.EmptyRefer[larksecurity_and_compliance.User],
		"refOfUser":                             engine.ReferOf[larksecurity_and_compliance.User],
		"unRefUser":                             engine.UnRefer[larksecurity_and_compliance.User],
		"emptyAppDlpExecuteLog":                 engine.Empty[larksecurity_and_compliance.AppDlpExecuteLog],
		"emptyRefAppDlpExecuteLog":              engine.EmptyRefer[larksecurity_and_compliance.AppDlpExecuteLog],
		"refOfAppDlpExecuteLog":                 engine.ReferOf[larksecurity_and_compliance.AppDlpExecuteLog],
		"unRefAppDlpExecuteLog":                 engine.UnRefer[larksecurity_and_compliance.AppDlpExecuteLog],
		"emptyEmail":                            engine.Empty[larksecurity_and_compliance.Email],
		"emptyRefEmail":                         engine.EmptyRefer[larksecurity_and_compliance.Email],
		"refOfEmail":                            engine.ReferOf[larksecurity_and_compliance.Email],
		"unRefEmail":                            engine.UnRefer[larksecurity_and_compliance.Email],
		"emptyGwRequest":                        engine.Empty[larksecurity_and_compliance.GwRequest],
		"emptyRefGwRequest":                     engine.EmptyRefer[larksecurity_and_compliance.GwRequest],
		"refOfGwRequest":                        engine.ReferOf[larksecurity_and_compliance.GwRequest],
		"unRefGwRequest":                        engine.UnRefer[larksecurity_and_compliance.GwRequest],
		"emptyKeyPerson":                        engine.Empty[larksecurity_and_compliance.KeyPerson],
		"emptyRefKeyPerson":                     engine.EmptyRefer[larksecurity_and_compliance.KeyPerson],
		"refOfKeyPerson":                        engine.ReferOf[larksecurity_and_compliance.KeyPerson],
		"unRefKeyPerson":                        engine.UnRefer[larksecurity_and_compliance.KeyPerson],
		"emptyTaskStatus":                       engine.Empty[larksecurity_and_compliance.TaskStatus],
		"emptyRefTaskStatus":                    engine.EmptyRefer[larksecurity_and_compliance.TaskStatus],
		"refOfTaskStatus":                       engine.ReferOf[larksecurity_and_compliance.TaskStatus],
		"unRefTaskStatus":                       engine.UnRefer[larksecurity_and_compliance.TaskStatus],
		"emptyUserMigration":                    engine.Empty[larksecurity_and_compliance.UserMigration],
		"emptyRefUserMigration":                 engine.EmptyRefer[larksecurity_and_compliance.UserMigration],
		"refOfUserMigration":                    engine.ReferOf[larksecurity_and_compliance.UserMigration],
		"unRefUserMigration":                    engine.UnRefer[larksecurity_and_compliance.UserMigration],
		"emptyMessage":                          engine.Empty[larksecurity_and_compliance.Message],
		"emptyRefMessage":                       engine.EmptyRefer[larksecurity_and_compliance.Message],
		"refOfMessage":                          engine.ReferOf[larksecurity_and_compliance.Message],
		"unRefMessage":                          engine.UnRefer[larksecurity_and_compliance.Message],
		"emptyGwCommon":                         engine.Empty[larksecurity_and_compliance.GwCommon],
		"emptyRefGwCommon":                      engine.EmptyRefer[larksecurity_and_compliance.GwCommon],
		"refOfGwCommon":                         engine.ReferOf[larksecurity_and_compliance.GwCommon],
		"unRefGwCommon":                         engine.UnRefer[larksecurity_and_compliance.GwCommon],
		"emptyListDataOpenapiLogResp":           engine.Empty[larksecurity_and_compliance.ListDataOpenapiLogResp],
		"emptyRefListDataOpenapiLogResp":        engine.EmptyRefer[larksecurity_and_compliance.ListDataOpenapiLogResp],
		"refOfListDataOpenapiLogResp":           engine.ReferOf[larksecurity_and_compliance.ListDataOpenapiLogResp],
		"unRefListDataOpenapiLogResp":           engine.UnRefer[larksecurity_and_compliance.ListDataOpenapiLogResp],
		"emptyNotification":                     engine.Empty[larksecurity_and_compliance.Notification],
		"emptyRefNotification":                  engine.EmptyRefer[larksecurity_and_compliance.Notification],
		"refOfNotification":                     engine.ReferOf[larksecurity_and_compliance.Notification],
		"unRefNotification":                     engine.UnRefer[larksecurity_and_compliance.Notification],
		"emptyTimeRange":                        engine.Empty[larksecurity_and_compliance.TimeRange],
		"emptyRefTimeRange":                     engine.EmptyRefer[larksecurity_and_compliance.TimeRange],
		"refOfTimeRange":                        engine.ReferOf[larksecurity_and_compliance.TimeRange],
		"unRefTimeRange":                        engine.UnRefer[larksecurity_and_compliance.TimeRange],
		"emptyTenantThirdPartyEncryptionApp":    engine.Empty[larksecurity_and_compliance.TenantThirdPartyEncryptionApp],
		"emptyRefTenantThirdPartyEncryptionApp": engine.EmptyRefer[larksecurity_and_compliance.TenantThirdPartyEncryptionApp],
		"refOfTenantThirdPartyEncryptionApp":    engine.ReferOf[larksecurity_and_compliance.TenantThirdPartyEncryptionApp],
		"unRefTenantThirdPartyEncryptionApp":    engine.UnRefer[larksecurity_and_compliance.TenantThirdPartyEncryptionApp],
		"emptyDataArchivingUserStruct":          engine.Empty[larksecurity_and_compliance.DataArchivingUserStruct],
		"emptyRefDataArchivingUserStruct":       engine.EmptyRefer[larksecurity_and_compliance.DataArchivingUserStruct],
		"refOfDataArchivingUserStruct":          engine.ReferOf[larksecurity_and_compliance.DataArchivingUserStruct],
		"unRefDataArchivingUserStruct":          engine.UnRefer[larksecurity_and_compliance.DataArchivingUserStruct],
		"emptyOpenapiLogDetail":                 engine.Empty[larksecurity_and_compliance.OpenapiLogDetail],
		"emptyRefOpenapiLogDetail":              engine.EmptyRefer[larksecurity_and_compliance.OpenapiLogDetail],
		"refOfOpenapiLogDetail":                 engine.ReferOf[larksecurity_and_compliance.OpenapiLogDetail],
		"unRefOpenapiLogDetail":                 engine.UnRefer[larksecurity_and_compliance.OpenapiLogDetail],
		"emptyDlpHitPolicy":                     engine.Empty[larksecurity_and_compliance.DlpHitPolicy],
		"emptyRefDlpHitPolicy":                  engine.EmptyRefer[larksecurity_and_compliance.DlpHitPolicy],
		"refOfDlpHitPolicy":                     engine.ReferOf[larksecurity_and_compliance.DlpHitPolicy],
		"unRefDlpHitPolicy":                     engine.UnRefer[larksecurity_and_compliance.DlpHitPolicy],
		"emptyListOpenapiLogRequest":            engine.Empty[larksecurity_and_compliance.ListOpenapiLogRequest],
		"emptyRefListOpenapiLogRequest":         engine.EmptyRefer[larksecurity_and_compliance.ListOpenapiLogRequest],
		"refOfListOpenapiLogRequest":            engine.ReferOf[larksecurity_and_compliance.ListOpenapiLogRequest],
		"unRefListOpenapiLogRequest":            engine.UnRefer[larksecurity_and_compliance.ListOpenapiLogRequest],
		"emptyCreateMigrationItems":             engine.Empty[larksecurity_and_compliance.CreateMigrationItems],
		"emptyRefCreateMigrationItems":          engine.EmptyRefer[larksecurity_and_compliance.CreateMigrationItems],
		"refOfCreateMigrationItems":             engine.ReferOf[larksecurity_and_compliance.CreateMigrationItems],
		"unRefCreateMigrationItems":             engine.UnRefer[larksecurity_and_compliance.CreateMigrationItems]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceSecurity_and_compliance1Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceSecurity_and_compliance1Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceSecurity_and_compliance1Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/security_and_compliance/v1"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceSecurity_and_compliance1Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceSecurity_and_compliance1Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceSecurity_and_compliance1Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceSecurity_and_compliance1Declared
}