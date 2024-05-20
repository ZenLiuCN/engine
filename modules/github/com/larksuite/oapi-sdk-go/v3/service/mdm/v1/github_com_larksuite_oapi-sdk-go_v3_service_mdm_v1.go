// Code generated by define_gene; DO NOT EDIT.
package larkmdm

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	"github.com/larksuite/oapi-sdk-go/v3/service/mdm/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_mdm_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceMdm1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceMdm1Declared = map[string]any{
		"newFixedExchangeRateBuilder":                 larkmdm.NewFixedExchangeRateBuilder,
		"newVendorBuilder":                            larkmdm.NewVendorBuilder,
		"newCompanyCompanyBankAccountBuilder":         larkmdm.NewCompanyCompanyBankAccountBuilder,
		"UserIdTypeUnbindUserAuthDataRelationOpenId":  larkmdm.UserIdTypeUnbindUserAuthDataRelationOpenId,
		"UserIdTypeUnionId":                           larkmdm.UserIdTypeUnionId,
		"newProjectCompanyDeptMappingBuilder":         larkmdm.NewProjectCompanyDeptMappingBuilder,
		"UserIdTypeUnbindUserAuthDataRelationUnionId": larkmdm.UserIdTypeUnbindUserAuthDataRelationUnionId,
		"New":                       larkmdm.New,
		"newAppendixBuilder":        larkmdm.NewAppendixBuilder,
		"newCompanyAssetBuilder":    larkmdm.NewCompanyAssetBuilder,
		"newLegalEntityBankBuilder": larkmdm.NewLegalEntityBankBuilder,
		"newVendorAddressBuilder":   larkmdm.NewVendorAddressBuilder,
		"UserIdTypeUnbindUserAuthDataRelationUserId": larkmdm.UserIdTypeUnbindUserAuthDataRelationUserId,
		"newCostCenterBuilder":                       larkmdm.NewCostCenterBuilder,
		"newDepartmentCostCenterRelationshipBuilder": larkmdm.NewDepartmentCostCenterRelationshipBuilder,
		"newInternalOrderBuilder":                    larkmdm.NewInternalOrderBuilder,
		"newProjectBuilder":                          larkmdm.NewProjectBuilder,
		"newVendorAccountBuilder":                    larkmdm.NewVendorAccountBuilder,
		"newVendorCompanyViewBuilder":                larkmdm.NewVendorCompanyViewBuilder,
		"newGlAccountCompanyRelationshipBuilder":     larkmdm.NewGlAccountCompanyRelationshipBuilder,
		"newI18nStructBuilder":                       larkmdm.NewI18nStructBuilder,
		"newLegalEntityBuilder":                      larkmdm.NewLegalEntityBuilder,
		"newVendorContactBuilder":                    larkmdm.NewVendorContactBuilder,
		"UserIdTypeOpenId":                           larkmdm.UserIdTypeOpenId,
		"UserIdTypeUserId":                           larkmdm.UserIdTypeUserId,
		"newDepartmentIdBuilder":                     larkmdm.NewDepartmentIdBuilder,
		"newGlAccountBuilder":                        larkmdm.NewGlAccountBuilder,
		"newMultiLanguageBuilder":                    larkmdm.NewMultiLanguageBuilder,
		"newUnbindUserAuthDataRelationReqBuilder":    larkmdm.NewUnbindUserAuthDataRelationReqBuilder,
		"newOpenApiUpdateVendorBuilder":              larkmdm.NewOpenApiUpdateVendorBuilder,
		"newUserAuthDataRelationBuilder":             larkmdm.NewUserAuthDataRelationBuilder,
		"newBindUserAuthDataRelationReqBuilder":      larkmdm.NewBindUserAuthDataRelationReqBuilder,
		"newCompanyBuilder":                          larkmdm.NewCompanyBuilder,
		"newConfigBuilder":                           larkmdm.NewConfigBuilder,
		"newExtendFieldBuilder":                      larkmdm.NewExtendFieldBuilder,

		"emptyI18nStruct":                          engine.Empty[larkmdm.I18nStruct],
		"emptyRefI18nStruct":                       engine.EmptyRefer[larkmdm.I18nStruct],
		"refOfI18nStruct":                          engine.ReferOf[larkmdm.I18nStruct],
		"unRefI18nStruct":                          engine.UnRefer[larkmdm.I18nStruct],
		"emptyOpenApiUpdateVendor":                 engine.Empty[larkmdm.OpenApiUpdateVendor],
		"emptyRefOpenApiUpdateVendor":              engine.EmptyRefer[larkmdm.OpenApiUpdateVendor],
		"refOfOpenApiUpdateVendor":                 engine.ReferOf[larkmdm.OpenApiUpdateVendor],
		"unRefOpenApiUpdateVendor":                 engine.UnRefer[larkmdm.OpenApiUpdateVendor],
		"emptyGlAccount":                           engine.Empty[larkmdm.GlAccount],
		"emptyRefGlAccount":                        engine.EmptyRefer[larkmdm.GlAccount],
		"refOfGlAccount":                           engine.ReferOf[larkmdm.GlAccount],
		"unRefGlAccount":                           engine.UnRefer[larkmdm.GlAccount],
		"emptyLegalEntity":                         engine.Empty[larkmdm.LegalEntity],
		"emptyRefLegalEntity":                      engine.EmptyRefer[larkmdm.LegalEntity],
		"refOfLegalEntity":                         engine.ReferOf[larkmdm.LegalEntity],
		"unRefLegalEntity":                         engine.UnRefer[larkmdm.LegalEntity],
		"emptyVendorContact":                       engine.Empty[larkmdm.VendorContact],
		"emptyRefVendorContact":                    engine.EmptyRefer[larkmdm.VendorContact],
		"refOfVendorContact":                       engine.ReferOf[larkmdm.VendorContact],
		"unRefVendorContact":                       engine.UnRefer[larkmdm.VendorContact],
		"emptyBindUserAuthDataRelationResp":        engine.Empty[larkmdm.BindUserAuthDataRelationResp],
		"emptyRefBindUserAuthDataRelationResp":     engine.EmptyRefer[larkmdm.BindUserAuthDataRelationResp],
		"refOfBindUserAuthDataRelationResp":        engine.ReferOf[larkmdm.BindUserAuthDataRelationResp],
		"unRefBindUserAuthDataRelationResp":        engine.UnRefer[larkmdm.BindUserAuthDataRelationResp],
		"emptyCompany":                             engine.Empty[larkmdm.Company],
		"emptyRefCompany":                          engine.EmptyRefer[larkmdm.Company],
		"refOfCompany":                             engine.ReferOf[larkmdm.Company],
		"unRefCompany":                             engine.UnRefer[larkmdm.Company],
		"emptyExtendField":                         engine.Empty[larkmdm.ExtendField],
		"emptyRefExtendField":                      engine.EmptyRefer[larkmdm.ExtendField],
		"refOfExtendField":                         engine.ReferOf[larkmdm.ExtendField],
		"unRefExtendField":                         engine.UnRefer[larkmdm.ExtendField],
		"emptyDepartmentCostCenterRelationship":    engine.Empty[larkmdm.DepartmentCostCenterRelationship],
		"emptyRefDepartmentCostCenterRelationship": engine.EmptyRefer[larkmdm.DepartmentCostCenterRelationship],
		"refOfDepartmentCostCenterRelationship":    engine.ReferOf[larkmdm.DepartmentCostCenterRelationship],
		"unRefDepartmentCostCenterRelationship":    engine.UnRefer[larkmdm.DepartmentCostCenterRelationship],
		"emptyCompanyCompanyBankAccount":           engine.Empty[larkmdm.CompanyCompanyBankAccount],
		"emptyRefCompanyCompanyBankAccount":        engine.EmptyRefer[larkmdm.CompanyCompanyBankAccount],
		"refOfCompanyCompanyBankAccount":           engine.ReferOf[larkmdm.CompanyCompanyBankAccount],
		"unRefCompanyCompanyBankAccount":           engine.UnRefer[larkmdm.CompanyCompanyBankAccount],
		"emptyLegalEntityBank":                     engine.Empty[larkmdm.LegalEntityBank],
		"emptyRefLegalEntityBank":                  engine.EmptyRefer[larkmdm.LegalEntityBank],
		"refOfLegalEntityBank":                     engine.ReferOf[larkmdm.LegalEntityBank],
		"unRefLegalEntityBank":                     engine.UnRefer[larkmdm.LegalEntityBank],
		"emptyInternalOrder":                       engine.Empty[larkmdm.InternalOrder],
		"emptyRefInternalOrder":                    engine.EmptyRefer[larkmdm.InternalOrder],
		"refOfInternalOrder":                       engine.ReferOf[larkmdm.InternalOrder],
		"unRefInternalOrder":                       engine.UnRefer[larkmdm.InternalOrder],
		"emptyAppendix":                            engine.Empty[larkmdm.Appendix],
		"emptyRefAppendix":                         engine.EmptyRefer[larkmdm.Appendix],
		"refOfAppendix":                            engine.ReferOf[larkmdm.Appendix],
		"unRefAppendix":                            engine.UnRefer[larkmdm.Appendix],
		"emptyCompanyAsset":                        engine.Empty[larkmdm.CompanyAsset],
		"emptyRefCompanyAsset":                     engine.EmptyRefer[larkmdm.CompanyAsset],
		"refOfCompanyAsset":                        engine.ReferOf[larkmdm.CompanyAsset],
		"unRefCompanyAsset":                        engine.UnRefer[larkmdm.CompanyAsset],
		"emptyMultiLanguage":                       engine.Empty[larkmdm.MultiLanguage],
		"emptyRefMultiLanguage":                    engine.EmptyRefer[larkmdm.MultiLanguage],
		"refOfMultiLanguage":                       engine.ReferOf[larkmdm.MultiLanguage],
		"unRefMultiLanguage":                       engine.UnRefer[larkmdm.MultiLanguage],
		"emptyUnbindUserAuthDataRelationResp":      engine.Empty[larkmdm.UnbindUserAuthDataRelationResp],
		"emptyRefUnbindUserAuthDataRelationResp":   engine.EmptyRefer[larkmdm.UnbindUserAuthDataRelationResp],
		"refOfUnbindUserAuthDataRelationResp":      engine.ReferOf[larkmdm.UnbindUserAuthDataRelationResp],
		"unRefUnbindUserAuthDataRelationResp":      engine.UnRefer[larkmdm.UnbindUserAuthDataRelationResp],
		"emptyUserAuthDataRelation":                engine.Empty[larkmdm.UserAuthDataRelation],
		"emptyRefUserAuthDataRelation":             engine.EmptyRefer[larkmdm.UserAuthDataRelation],
		"refOfUserAuthDataRelation":                engine.ReferOf[larkmdm.UserAuthDataRelation],
		"unRefUserAuthDataRelation":                engine.UnRefer[larkmdm.UserAuthDataRelation],
		"emptyCostCenter":                          engine.Empty[larkmdm.CostCenter],
		"emptyRefCostCenter":                       engine.EmptyRefer[larkmdm.CostCenter],
		"refOfCostCenter":                          engine.ReferOf[larkmdm.CostCenter],
		"unRefCostCenter":                          engine.UnRefer[larkmdm.CostCenter],
		"emptyProjectCompanyDeptMapping":           engine.Empty[larkmdm.ProjectCompanyDeptMapping],
		"emptyRefProjectCompanyDeptMapping":        engine.EmptyRefer[larkmdm.ProjectCompanyDeptMapping],
		"refOfProjectCompanyDeptMapping":           engine.ReferOf[larkmdm.ProjectCompanyDeptMapping],
		"unRefProjectCompanyDeptMapping":           engine.UnRefer[larkmdm.ProjectCompanyDeptMapping],
		"emptyDepartmentId":                        engine.Empty[larkmdm.DepartmentId],
		"emptyRefDepartmentId":                     engine.EmptyRefer[larkmdm.DepartmentId],
		"refOfDepartmentId":                        engine.ReferOf[larkmdm.DepartmentId],
		"unRefDepartmentId":                        engine.UnRefer[larkmdm.DepartmentId],
		"emptyVendor":                              engine.Empty[larkmdm.Vendor],
		"emptyRefVendor":                           engine.EmptyRefer[larkmdm.Vendor],
		"refOfVendor":                              engine.ReferOf[larkmdm.Vendor],
		"unRefVendor":                              engine.UnRefer[larkmdm.Vendor],
		"emptyProject":                             engine.Empty[larkmdm.Project],
		"emptyRefProject":                          engine.EmptyRefer[larkmdm.Project],
		"refOfProject":                             engine.ReferOf[larkmdm.Project],
		"unRefProject":                             engine.UnRefer[larkmdm.Project],
		"emptyVendorAddress":                       engine.Empty[larkmdm.VendorAddress],
		"emptyRefVendorAddress":                    engine.EmptyRefer[larkmdm.VendorAddress],
		"refOfVendorAddress":                       engine.ReferOf[larkmdm.VendorAddress],
		"unRefVendorAddress":                       engine.UnRefer[larkmdm.VendorAddress],
		"emptyConfig":                              engine.Empty[larkmdm.Config],
		"emptyRefConfig":                           engine.EmptyRefer[larkmdm.Config],
		"refOfConfig":                              engine.ReferOf[larkmdm.Config],
		"unRefConfig":                              engine.UnRefer[larkmdm.Config],
		"emptyVendorCompanyView":                   engine.Empty[larkmdm.VendorCompanyView],
		"emptyRefVendorCompanyView":                engine.EmptyRefer[larkmdm.VendorCompanyView],
		"refOfVendorCompanyView":                   engine.ReferOf[larkmdm.VendorCompanyView],
		"unRefVendorCompanyView":                   engine.UnRefer[larkmdm.VendorCompanyView],
		"emptyBindUserAuthDataRelationReq":         engine.Empty[larkmdm.BindUserAuthDataRelationReq],
		"emptyRefBindUserAuthDataRelationReq":      engine.EmptyRefer[larkmdm.BindUserAuthDataRelationReq],
		"refOfBindUserAuthDataRelationReq":         engine.ReferOf[larkmdm.BindUserAuthDataRelationReq],
		"unRefBindUserAuthDataRelationReq":         engine.UnRefer[larkmdm.BindUserAuthDataRelationReq],
		"emptyV1":                                  engine.Empty[larkmdm.V1],
		"emptyRefV1":                               engine.EmptyRefer[larkmdm.V1],
		"refOfV1":                                  engine.ReferOf[larkmdm.V1],
		"unRefV1":                                  engine.UnRefer[larkmdm.V1],
		"emptyUnbindUserAuthDataRelationReq":       engine.Empty[larkmdm.UnbindUserAuthDataRelationReq],
		"emptyRefUnbindUserAuthDataRelationReq":    engine.EmptyRefer[larkmdm.UnbindUserAuthDataRelationReq],
		"refOfUnbindUserAuthDataRelationReq":       engine.ReferOf[larkmdm.UnbindUserAuthDataRelationReq],
		"unRefUnbindUserAuthDataRelationReq":       engine.UnRefer[larkmdm.UnbindUserAuthDataRelationReq],
		"emptyVendorAccount":                       engine.Empty[larkmdm.VendorAccount],
		"emptyRefVendorAccount":                    engine.EmptyRefer[larkmdm.VendorAccount],
		"refOfVendorAccount":                       engine.ReferOf[larkmdm.VendorAccount],
		"unRefVendorAccount":                       engine.UnRefer[larkmdm.VendorAccount],
		"emptyFixedExchangeRate":                   engine.Empty[larkmdm.FixedExchangeRate],
		"emptyRefFixedExchangeRate":                engine.EmptyRefer[larkmdm.FixedExchangeRate],
		"refOfFixedExchangeRate":                   engine.ReferOf[larkmdm.FixedExchangeRate],
		"unRefFixedExchangeRate":                   engine.UnRefer[larkmdm.FixedExchangeRate],
		"emptyGlAccountCompanyRelationship":        engine.Empty[larkmdm.GlAccountCompanyRelationship],
		"emptyRefGlAccountCompanyRelationship":     engine.EmptyRefer[larkmdm.GlAccountCompanyRelationship],
		"refOfGlAccountCompanyRelationship":        engine.ReferOf[larkmdm.GlAccountCompanyRelationship],
		"unRefGlAccountCompanyRelationship":        engine.UnRefer[larkmdm.GlAccountCompanyRelationship]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceMdm1Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceMdm1Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceMdm1Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/mdm/v1"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceMdm1Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceMdm1Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceMdm1Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceMdm1Declared
}
