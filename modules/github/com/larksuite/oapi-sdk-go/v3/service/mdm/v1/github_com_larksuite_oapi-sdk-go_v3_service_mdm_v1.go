// Code generated by define_gene; DO NOT EDIT.
package larkmdm

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/mdm/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_mdm_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceMdm1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceMdm1Declared = map[string]any{
		"newAppendixBuilder":                          larkmdm.NewAppendixBuilder,
		"newDepartmentCostCenterRelationshipBuilder":  larkmdm.NewDepartmentCostCenterRelationshipBuilder,
		"newVendorAccountBuilder":                     larkmdm.NewVendorAccountBuilder,
		"UserIdTypeUnionId":                           larkmdm.UserIdTypeUnionId,
		"newUserAuthDataRelationBuilder":              larkmdm.NewUserAuthDataRelationBuilder,
		"newVendorAddressBuilder":                     larkmdm.NewVendorAddressBuilder,
		"New":                                         larkmdm.New,
		"newConfigBuilder":                            larkmdm.NewConfigBuilder,
		"newDepartmentIdBuilder":                      larkmdm.NewDepartmentIdBuilder,
		"newGlAccountCompanyRelationshipBuilder":      larkmdm.NewGlAccountCompanyRelationshipBuilder,
		"newMultiLanguageBuilder":                     larkmdm.NewMultiLanguageBuilder,
		"newProjectCompanyDeptMappingBuilder":         larkmdm.NewProjectCompanyDeptMappingBuilder,
		"newVendorBuilder":                            larkmdm.NewVendorBuilder,
		"newVendorContactBuilder":                     larkmdm.NewVendorContactBuilder,
		"UserIdTypeUnbindUserAuthDataRelationOpenId":  larkmdm.UserIdTypeUnbindUserAuthDataRelationOpenId,
		"newCostCenterBuilder":                        larkmdm.NewCostCenterBuilder,
		"newGlAccountBuilder":                         larkmdm.NewGlAccountBuilder,
		"newVendorCompanyViewBuilder":                 larkmdm.NewVendorCompanyViewBuilder,
		"UserIdTypeUserId":                            larkmdm.UserIdTypeUserId,
		"newBindUserAuthDataRelationReqBuilder":       larkmdm.NewBindUserAuthDataRelationReqBuilder,
		"newCompanyCompanyBankAccountBuilder":         larkmdm.NewCompanyCompanyBankAccountBuilder,
		"newCompanyAssetBuilder":                      larkmdm.NewCompanyAssetBuilder,
		"newExtendFieldBuilder":                       larkmdm.NewExtendFieldBuilder,
		"newLegalEntityBuilder":                       larkmdm.NewLegalEntityBuilder,
		"UserIdTypeUnbindUserAuthDataRelationUserId":  larkmdm.UserIdTypeUnbindUserAuthDataRelationUserId,
		"newCompanyBuilder":                           larkmdm.NewCompanyBuilder,
		"newI18nStructBuilder":                        larkmdm.NewI18nStructBuilder,
		"newOpenApiUpdateVendorBuilder":               larkmdm.NewOpenApiUpdateVendorBuilder,
		"newUnbindUserAuthDataRelationReqBuilder":     larkmdm.NewUnbindUserAuthDataRelationReqBuilder,
		"UserIdTypeUnbindUserAuthDataRelationUnionId": larkmdm.UserIdTypeUnbindUserAuthDataRelationUnionId,
		"newFixedExchangeRateBuilder":                 larkmdm.NewFixedExchangeRateBuilder,
		"newInternalOrderBuilder":                     larkmdm.NewInternalOrderBuilder,
		"newLegalEntityBankBuilder":                   larkmdm.NewLegalEntityBankBuilder,
		"newProjectBuilder":                           larkmdm.NewProjectBuilder,
		"UserIdTypeOpenId":                            larkmdm.UserIdTypeOpenId,

		"emptyProject":    engine.Empty[larkmdm.Project],
		"emptyRefProject": engine.EmptyRefer[larkmdm.Project],
		"refOfProject":    engine.ReferOf[larkmdm.Project],
		"unRefProject":    engine.UnRefer[larkmdm.Project],
		"emptyGlAccountCompanyRelationshipBuilder":        engine.Empty[larkmdm.GlAccountCompanyRelationshipBuilder],
		"emptyRefGlAccountCompanyRelationshipBuilder":     engine.EmptyRefer[larkmdm.GlAccountCompanyRelationshipBuilder],
		"refOfGlAccountCompanyRelationshipBuilder":        engine.ReferOf[larkmdm.GlAccountCompanyRelationshipBuilder],
		"unRefGlAccountCompanyRelationshipBuilder":        engine.UnRefer[larkmdm.GlAccountCompanyRelationshipBuilder],
		"emptyVendorAccount":                              engine.Empty[larkmdm.VendorAccount],
		"emptyRefVendorAccount":                           engine.EmptyRefer[larkmdm.VendorAccount],
		"refOfVendorAccount":                              engine.ReferOf[larkmdm.VendorAccount],
		"unRefVendorAccount":                              engine.UnRefer[larkmdm.VendorAccount],
		"emptyGlAccountBuilder":                           engine.Empty[larkmdm.GlAccountBuilder],
		"emptyRefGlAccountBuilder":                        engine.EmptyRefer[larkmdm.GlAccountBuilder],
		"refOfGlAccountBuilder":                           engine.ReferOf[larkmdm.GlAccountBuilder],
		"unRefGlAccountBuilder":                           engine.UnRefer[larkmdm.GlAccountBuilder],
		"emptyLegalEntityBuilder":                         engine.Empty[larkmdm.LegalEntityBuilder],
		"emptyRefLegalEntityBuilder":                      engine.EmptyRefer[larkmdm.LegalEntityBuilder],
		"refOfLegalEntityBuilder":                         engine.ReferOf[larkmdm.LegalEntityBuilder],
		"unRefLegalEntityBuilder":                         engine.UnRefer[larkmdm.LegalEntityBuilder],
		"emptyUserAuthDataRelation":                       engine.Empty[larkmdm.UserAuthDataRelation],
		"emptyRefUserAuthDataRelation":                    engine.EmptyRefer[larkmdm.UserAuthDataRelation],
		"refOfUserAuthDataRelation":                       engine.ReferOf[larkmdm.UserAuthDataRelation],
		"unRefUserAuthDataRelation":                       engine.UnRefer[larkmdm.UserAuthDataRelation],
		"emptyVendorCompanyView":                          engine.Empty[larkmdm.VendorCompanyView],
		"emptyRefVendorCompanyView":                       engine.EmptyRefer[larkmdm.VendorCompanyView],
		"refOfVendorCompanyView":                          engine.ReferOf[larkmdm.VendorCompanyView],
		"unRefVendorCompanyView":                          engine.UnRefer[larkmdm.VendorCompanyView],
		"emptyBindUserAuthDataRelationResp":               engine.Empty[larkmdm.BindUserAuthDataRelationResp],
		"emptyRefBindUserAuthDataRelationResp":            engine.EmptyRefer[larkmdm.BindUserAuthDataRelationResp],
		"refOfBindUserAuthDataRelationResp":               engine.ReferOf[larkmdm.BindUserAuthDataRelationResp],
		"unRefBindUserAuthDataRelationResp":               engine.UnRefer[larkmdm.BindUserAuthDataRelationResp],
		"emptyConfig":                                     engine.Empty[larkmdm.Config],
		"emptyRefConfig":                                  engine.EmptyRefer[larkmdm.Config],
		"refOfConfig":                                     engine.ReferOf[larkmdm.Config],
		"unRefConfig":                                     engine.UnRefer[larkmdm.Config],
		"emptyI18nStruct":                                 engine.Empty[larkmdm.I18nStruct],
		"emptyRefI18nStruct":                              engine.EmptyRefer[larkmdm.I18nStruct],
		"refOfI18nStruct":                                 engine.ReferOf[larkmdm.I18nStruct],
		"unRefI18nStruct":                                 engine.UnRefer[larkmdm.I18nStruct],
		"emptyConfigBuilder":                              engine.Empty[larkmdm.ConfigBuilder],
		"emptyRefConfigBuilder":                           engine.EmptyRefer[larkmdm.ConfigBuilder],
		"refOfConfigBuilder":                              engine.ReferOf[larkmdm.ConfigBuilder],
		"unRefConfigBuilder":                              engine.UnRefer[larkmdm.ConfigBuilder],
		"emptyDepartmentCostCenterRelationshipBuilder":    engine.Empty[larkmdm.DepartmentCostCenterRelationshipBuilder],
		"emptyRefDepartmentCostCenterRelationshipBuilder": engine.EmptyRefer[larkmdm.DepartmentCostCenterRelationshipBuilder],
		"refOfDepartmentCostCenterRelationshipBuilder":    engine.ReferOf[larkmdm.DepartmentCostCenterRelationshipBuilder],
		"unRefDepartmentCostCenterRelationshipBuilder":    engine.UnRefer[larkmdm.DepartmentCostCenterRelationshipBuilder],
		"emptyDepartmentCostCenterRelationship":           engine.Empty[larkmdm.DepartmentCostCenterRelationship],
		"emptyRefDepartmentCostCenterRelationship":        engine.EmptyRefer[larkmdm.DepartmentCostCenterRelationship],
		"refOfDepartmentCostCenterRelationship":           engine.ReferOf[larkmdm.DepartmentCostCenterRelationship],
		"unRefDepartmentCostCenterRelationship":           engine.UnRefer[larkmdm.DepartmentCostCenterRelationship],
		"emptyInternalOrderBuilder":                       engine.Empty[larkmdm.InternalOrderBuilder],
		"emptyRefInternalOrderBuilder":                    engine.EmptyRefer[larkmdm.InternalOrderBuilder],
		"refOfInternalOrderBuilder":                       engine.ReferOf[larkmdm.InternalOrderBuilder],
		"unRefInternalOrderBuilder":                       engine.UnRefer[larkmdm.InternalOrderBuilder],
		"emptyLegalEntity":                                engine.Empty[larkmdm.LegalEntity],
		"emptyRefLegalEntity":                             engine.EmptyRefer[larkmdm.LegalEntity],
		"refOfLegalEntity":                                engine.ReferOf[larkmdm.LegalEntity],
		"unRefLegalEntity":                                engine.UnRefer[larkmdm.LegalEntity],
		"emptyProjectBuilder":                             engine.Empty[larkmdm.ProjectBuilder],
		"emptyRefProjectBuilder":                          engine.EmptyRefer[larkmdm.ProjectBuilder],
		"refOfProjectBuilder":                             engine.ReferOf[larkmdm.ProjectBuilder],
		"unRefProjectBuilder":                             engine.UnRefer[larkmdm.ProjectBuilder],
		"emptyVendorAccountBuilder":                       engine.Empty[larkmdm.VendorAccountBuilder],
		"emptyRefVendorAccountBuilder":                    engine.EmptyRefer[larkmdm.VendorAccountBuilder],
		"refOfVendorAccountBuilder":                       engine.ReferOf[larkmdm.VendorAccountBuilder],
		"unRefVendorAccountBuilder":                       engine.UnRefer[larkmdm.VendorAccountBuilder],
		"emptyCompanyAssetBuilder":                        engine.Empty[larkmdm.CompanyAssetBuilder],
		"emptyRefCompanyAssetBuilder":                     engine.EmptyRefer[larkmdm.CompanyAssetBuilder],
		"refOfCompanyAssetBuilder":                        engine.ReferOf[larkmdm.CompanyAssetBuilder],
		"unRefCompanyAssetBuilder":                        engine.UnRefer[larkmdm.CompanyAssetBuilder],
		"emptyCompanyCompanyBankAccount":                  engine.Empty[larkmdm.CompanyCompanyBankAccount],
		"emptyRefCompanyCompanyBankAccount":               engine.EmptyRefer[larkmdm.CompanyCompanyBankAccount],
		"refOfCompanyCompanyBankAccount":                  engine.ReferOf[larkmdm.CompanyCompanyBankAccount],
		"unRefCompanyCompanyBankAccount":                  engine.UnRefer[larkmdm.CompanyCompanyBankAccount],
		"emptyUnbindUserAuthDataRelationResp":             engine.Empty[larkmdm.UnbindUserAuthDataRelationResp],
		"emptyRefUnbindUserAuthDataRelationResp":          engine.EmptyRefer[larkmdm.UnbindUserAuthDataRelationResp],
		"refOfUnbindUserAuthDataRelationResp":             engine.ReferOf[larkmdm.UnbindUserAuthDataRelationResp],
		"unRefUnbindUserAuthDataRelationResp":             engine.UnRefer[larkmdm.UnbindUserAuthDataRelationResp],
		"emptyDepartmentId":                               engine.Empty[larkmdm.DepartmentId],
		"emptyRefDepartmentId":                            engine.EmptyRefer[larkmdm.DepartmentId],
		"refOfDepartmentId":                               engine.ReferOf[larkmdm.DepartmentId],
		"unRefDepartmentId":                               engine.UnRefer[larkmdm.DepartmentId],
		"emptyFixedExchangeRateBuilder":                   engine.Empty[larkmdm.FixedExchangeRateBuilder],
		"emptyRefFixedExchangeRateBuilder":                engine.EmptyRefer[larkmdm.FixedExchangeRateBuilder],
		"refOfFixedExchangeRateBuilder":                   engine.ReferOf[larkmdm.FixedExchangeRateBuilder],
		"unRefFixedExchangeRateBuilder":                   engine.UnRefer[larkmdm.FixedExchangeRateBuilder],
		"emptyVendorAddressBuilder":                       engine.Empty[larkmdm.VendorAddressBuilder],
		"emptyRefVendorAddressBuilder":                    engine.EmptyRefer[larkmdm.VendorAddressBuilder],
		"refOfVendorAddressBuilder":                       engine.ReferOf[larkmdm.VendorAddressBuilder],
		"unRefVendorAddressBuilder":                       engine.UnRefer[larkmdm.VendorAddressBuilder],
		"emptyVendorContact":                              engine.Empty[larkmdm.VendorContact],
		"emptyRefVendorContact":                           engine.EmptyRefer[larkmdm.VendorContact],
		"refOfVendorContact":                              engine.ReferOf[larkmdm.VendorContact],
		"unRefVendorContact":                              engine.UnRefer[larkmdm.VendorContact],
		"emptyCostCenterBuilder":                          engine.Empty[larkmdm.CostCenterBuilder],
		"emptyRefCostCenterBuilder":                       engine.EmptyRefer[larkmdm.CostCenterBuilder],
		"refOfCostCenterBuilder":                          engine.ReferOf[larkmdm.CostCenterBuilder],
		"unRefCostCenterBuilder":                          engine.UnRefer[larkmdm.CostCenterBuilder],
		"emptyLegalEntityBank":                            engine.Empty[larkmdm.LegalEntityBank],
		"emptyRefLegalEntityBank":                         engine.EmptyRefer[larkmdm.LegalEntityBank],
		"refOfLegalEntityBank":                            engine.ReferOf[larkmdm.LegalEntityBank],
		"unRefLegalEntityBank":                            engine.UnRefer[larkmdm.LegalEntityBank],
		"emptyDepartmentIdBuilder":                        engine.Empty[larkmdm.DepartmentIdBuilder],
		"emptyRefDepartmentIdBuilder":                     engine.EmptyRefer[larkmdm.DepartmentIdBuilder],
		"refOfDepartmentIdBuilder":                        engine.ReferOf[larkmdm.DepartmentIdBuilder],
		"unRefDepartmentIdBuilder":                        engine.UnRefer[larkmdm.DepartmentIdBuilder],
		"emptyGlAccountCompanyRelationship":               engine.Empty[larkmdm.GlAccountCompanyRelationship],
		"emptyRefGlAccountCompanyRelationship":            engine.EmptyRefer[larkmdm.GlAccountCompanyRelationship],
		"refOfGlAccountCompanyRelationship":               engine.ReferOf[larkmdm.GlAccountCompanyRelationship],
		"unRefGlAccountCompanyRelationship":               engine.UnRefer[larkmdm.GlAccountCompanyRelationship],
		"emptyVendorAddress":                              engine.Empty[larkmdm.VendorAddress],
		"emptyRefVendorAddress":                           engine.EmptyRefer[larkmdm.VendorAddress],
		"refOfVendorAddress":                              engine.ReferOf[larkmdm.VendorAddress],
		"unRefVendorAddress":                              engine.UnRefer[larkmdm.VendorAddress],
		"emptyCompanyAsset":                               engine.Empty[larkmdm.CompanyAsset],
		"emptyRefCompanyAsset":                            engine.EmptyRefer[larkmdm.CompanyAsset],
		"refOfCompanyAsset":                               engine.ReferOf[larkmdm.CompanyAsset],
		"unRefCompanyAsset":                               engine.UnRefer[larkmdm.CompanyAsset],
		"emptyCompanyCompanyBankAccountBuilder":           engine.Empty[larkmdm.CompanyCompanyBankAccountBuilder],
		"emptyRefCompanyCompanyBankAccountBuilder":        engine.EmptyRefer[larkmdm.CompanyCompanyBankAccountBuilder],
		"refOfCompanyCompanyBankAccountBuilder":           engine.ReferOf[larkmdm.CompanyCompanyBankAccountBuilder],
		"unRefCompanyCompanyBankAccountBuilder":           engine.UnRefer[larkmdm.CompanyCompanyBankAccountBuilder],
		"emptyCostCenter":                                 engine.Empty[larkmdm.CostCenter],
		"emptyRefCostCenter":                              engine.EmptyRefer[larkmdm.CostCenter],
		"refOfCostCenter":                                 engine.ReferOf[larkmdm.CostCenter],
		"unRefCostCenter":                                 engine.UnRefer[larkmdm.CostCenter],
		"emptyV1":                                         engine.Empty[larkmdm.V1],
		"emptyRefV1":                                      engine.EmptyRefer[larkmdm.V1],
		"refOfV1":                                         engine.ReferOf[larkmdm.V1],
		"unRefV1":                                         engine.UnRefer[larkmdm.V1],
		"emptyI18nStructBuilder":                          engine.Empty[larkmdm.I18nStructBuilder],
		"emptyRefI18nStructBuilder":                       engine.EmptyRefer[larkmdm.I18nStructBuilder],
		"refOfI18nStructBuilder":                          engine.ReferOf[larkmdm.I18nStructBuilder],
		"unRefI18nStructBuilder":                          engine.UnRefer[larkmdm.I18nStructBuilder],
		"emptyInternalOrder":                              engine.Empty[larkmdm.InternalOrder],
		"emptyRefInternalOrder":                           engine.EmptyRefer[larkmdm.InternalOrder],
		"refOfInternalOrder":                              engine.ReferOf[larkmdm.InternalOrder],
		"unRefInternalOrder":                              engine.UnRefer[larkmdm.InternalOrder],
		"emptyCompany":                                    engine.Empty[larkmdm.Company],
		"emptyRefCompany":                                 engine.EmptyRefer[larkmdm.Company],
		"refOfCompany":                                    engine.ReferOf[larkmdm.Company],
		"unRefCompany":                                    engine.UnRefer[larkmdm.Company],
		"emptyCompanyBuilder":                             engine.Empty[larkmdm.CompanyBuilder],
		"emptyRefCompanyBuilder":                          engine.EmptyRefer[larkmdm.CompanyBuilder],
		"refOfCompanyBuilder":                             engine.ReferOf[larkmdm.CompanyBuilder],
		"unRefCompanyBuilder":                             engine.UnRefer[larkmdm.CompanyBuilder],
		"emptyMultiLanguageBuilder":                       engine.Empty[larkmdm.MultiLanguageBuilder],
		"emptyRefMultiLanguageBuilder":                    engine.EmptyRefer[larkmdm.MultiLanguageBuilder],
		"refOfMultiLanguageBuilder":                       engine.ReferOf[larkmdm.MultiLanguageBuilder],
		"unRefMultiLanguageBuilder":                       engine.UnRefer[larkmdm.MultiLanguageBuilder],
		"emptyVendorBuilder":                              engine.Empty[larkmdm.VendorBuilder],
		"emptyRefVendorBuilder":                           engine.EmptyRefer[larkmdm.VendorBuilder],
		"refOfVendorBuilder":                              engine.ReferOf[larkmdm.VendorBuilder],
		"unRefVendorBuilder":                              engine.UnRefer[larkmdm.VendorBuilder],
		"emptyBindUserAuthDataRelationReq":                engine.Empty[larkmdm.BindUserAuthDataRelationReq],
		"emptyRefBindUserAuthDataRelationReq":             engine.EmptyRefer[larkmdm.BindUserAuthDataRelationReq],
		"refOfBindUserAuthDataRelationReq":                engine.ReferOf[larkmdm.BindUserAuthDataRelationReq],
		"unRefBindUserAuthDataRelationReq":                engine.UnRefer[larkmdm.BindUserAuthDataRelationReq],
		"emptyBindUserAuthDataRelationReqBuilder":         engine.Empty[larkmdm.BindUserAuthDataRelationReqBuilder],
		"emptyRefBindUserAuthDataRelationReqBuilder":      engine.EmptyRefer[larkmdm.BindUserAuthDataRelationReqBuilder],
		"refOfBindUserAuthDataRelationReqBuilder":         engine.ReferOf[larkmdm.BindUserAuthDataRelationReqBuilder],
		"unRefBindUserAuthDataRelationReqBuilder":         engine.UnRefer[larkmdm.BindUserAuthDataRelationReqBuilder],
		"emptyVendorContactBuilder":                       engine.Empty[larkmdm.VendorContactBuilder],
		"emptyRefVendorContactBuilder":                    engine.EmptyRefer[larkmdm.VendorContactBuilder],
		"refOfVendorContactBuilder":                       engine.ReferOf[larkmdm.VendorContactBuilder],
		"unRefVendorContactBuilder":                       engine.UnRefer[larkmdm.VendorContactBuilder],
		"emptyOpenApiUpdateVendorBuilder":                 engine.Empty[larkmdm.OpenApiUpdateVendorBuilder],
		"emptyRefOpenApiUpdateVendorBuilder":              engine.EmptyRefer[larkmdm.OpenApiUpdateVendorBuilder],
		"refOfOpenApiUpdateVendorBuilder":                 engine.ReferOf[larkmdm.OpenApiUpdateVendorBuilder],
		"unRefOpenApiUpdateVendorBuilder":                 engine.UnRefer[larkmdm.OpenApiUpdateVendorBuilder],
		"emptyUserAuthDataRelationBuilder":                engine.Empty[larkmdm.UserAuthDataRelationBuilder],
		"emptyRefUserAuthDataRelationBuilder":             engine.EmptyRefer[larkmdm.UserAuthDataRelationBuilder],
		"refOfUserAuthDataRelationBuilder":                engine.ReferOf[larkmdm.UserAuthDataRelationBuilder],
		"unRefUserAuthDataRelationBuilder":                engine.UnRefer[larkmdm.UserAuthDataRelationBuilder],
		"emptyLegalEntityBankBuilder":                     engine.Empty[larkmdm.LegalEntityBankBuilder],
		"emptyRefLegalEntityBankBuilder":                  engine.EmptyRefer[larkmdm.LegalEntityBankBuilder],
		"refOfLegalEntityBankBuilder":                     engine.ReferOf[larkmdm.LegalEntityBankBuilder],
		"unRefLegalEntityBankBuilder":                     engine.UnRefer[larkmdm.LegalEntityBankBuilder],
		"emptyMultiLanguage":                              engine.Empty[larkmdm.MultiLanguage],
		"emptyRefMultiLanguage":                           engine.EmptyRefer[larkmdm.MultiLanguage],
		"refOfMultiLanguage":                              engine.ReferOf[larkmdm.MultiLanguage],
		"unRefMultiLanguage":                              engine.UnRefer[larkmdm.MultiLanguage],
		"emptyOpenApiUpdateVendor":                        engine.Empty[larkmdm.OpenApiUpdateVendor],
		"emptyRefOpenApiUpdateVendor":                     engine.EmptyRefer[larkmdm.OpenApiUpdateVendor],
		"refOfOpenApiUpdateVendor":                        engine.ReferOf[larkmdm.OpenApiUpdateVendor],
		"unRefOpenApiUpdateVendor":                        engine.UnRefer[larkmdm.OpenApiUpdateVendor],
		"emptyUnbindUserAuthDataRelationReqBuilder":       engine.Empty[larkmdm.UnbindUserAuthDataRelationReqBuilder],
		"emptyRefUnbindUserAuthDataRelationReqBuilder":    engine.EmptyRefer[larkmdm.UnbindUserAuthDataRelationReqBuilder],
		"refOfUnbindUserAuthDataRelationReqBuilder":       engine.ReferOf[larkmdm.UnbindUserAuthDataRelationReqBuilder],
		"unRefUnbindUserAuthDataRelationReqBuilder":       engine.UnRefer[larkmdm.UnbindUserAuthDataRelationReqBuilder],
		"emptyVendor":                                     engine.Empty[larkmdm.Vendor],
		"emptyRefVendor":                                  engine.EmptyRefer[larkmdm.Vendor],
		"refOfVendor":                                     engine.ReferOf[larkmdm.Vendor],
		"unRefVendor":                                     engine.UnRefer[larkmdm.Vendor],
		"emptyFixedExchangeRate":                          engine.Empty[larkmdm.FixedExchangeRate],
		"emptyRefFixedExchangeRate":                       engine.EmptyRefer[larkmdm.FixedExchangeRate],
		"refOfFixedExchangeRate":                          engine.ReferOf[larkmdm.FixedExchangeRate],
		"unRefFixedExchangeRate":                          engine.UnRefer[larkmdm.FixedExchangeRate],
		"emptyGlAccount":                                  engine.Empty[larkmdm.GlAccount],
		"emptyRefGlAccount":                               engine.EmptyRefer[larkmdm.GlAccount],
		"refOfGlAccount":                                  engine.ReferOf[larkmdm.GlAccount],
		"unRefGlAccount":                                  engine.UnRefer[larkmdm.GlAccount],
		"emptyExtendField":                                engine.Empty[larkmdm.ExtendField],
		"emptyRefExtendField":                             engine.EmptyRefer[larkmdm.ExtendField],
		"refOfExtendField":                                engine.ReferOf[larkmdm.ExtendField],
		"unRefExtendField":                                engine.UnRefer[larkmdm.ExtendField],
		"emptyAppendix":                                   engine.Empty[larkmdm.Appendix],
		"emptyRefAppendix":                                engine.EmptyRefer[larkmdm.Appendix],
		"refOfAppendix":                                   engine.ReferOf[larkmdm.Appendix],
		"unRefAppendix":                                   engine.UnRefer[larkmdm.Appendix],
		"emptyAppendixBuilder":                            engine.Empty[larkmdm.AppendixBuilder],
		"emptyRefAppendixBuilder":                         engine.EmptyRefer[larkmdm.AppendixBuilder],
		"refOfAppendixBuilder":                            engine.ReferOf[larkmdm.AppendixBuilder],
		"unRefAppendixBuilder":                            engine.UnRefer[larkmdm.AppendixBuilder],
		"emptyUnbindUserAuthDataRelationReq":              engine.Empty[larkmdm.UnbindUserAuthDataRelationReq],
		"emptyRefUnbindUserAuthDataRelationReq":           engine.EmptyRefer[larkmdm.UnbindUserAuthDataRelationReq],
		"refOfUnbindUserAuthDataRelationReq":              engine.ReferOf[larkmdm.UnbindUserAuthDataRelationReq],
		"unRefUnbindUserAuthDataRelationReq":              engine.UnRefer[larkmdm.UnbindUserAuthDataRelationReq],
		"emptyProjectCompanyDeptMapping":                  engine.Empty[larkmdm.ProjectCompanyDeptMapping],
		"emptyRefProjectCompanyDeptMapping":               engine.EmptyRefer[larkmdm.ProjectCompanyDeptMapping],
		"refOfProjectCompanyDeptMapping":                  engine.ReferOf[larkmdm.ProjectCompanyDeptMapping],
		"unRefProjectCompanyDeptMapping":                  engine.UnRefer[larkmdm.ProjectCompanyDeptMapping],
		"emptyProjectCompanyDeptMappingBuilder":           engine.Empty[larkmdm.ProjectCompanyDeptMappingBuilder],
		"emptyRefProjectCompanyDeptMappingBuilder":        engine.EmptyRefer[larkmdm.ProjectCompanyDeptMappingBuilder],
		"refOfProjectCompanyDeptMappingBuilder":           engine.ReferOf[larkmdm.ProjectCompanyDeptMappingBuilder],
		"unRefProjectCompanyDeptMappingBuilder":           engine.UnRefer[larkmdm.ProjectCompanyDeptMappingBuilder],
		"emptyExtendFieldBuilder":                         engine.Empty[larkmdm.ExtendFieldBuilder],
		"emptyRefExtendFieldBuilder":                      engine.EmptyRefer[larkmdm.ExtendFieldBuilder],
		"refOfExtendFieldBuilder":                         engine.ReferOf[larkmdm.ExtendFieldBuilder],
		"unRefExtendFieldBuilder":                         engine.UnRefer[larkmdm.ExtendFieldBuilder],
		"emptyVendorCompanyViewBuilder":                   engine.Empty[larkmdm.VendorCompanyViewBuilder],
		"emptyRefVendorCompanyViewBuilder":                engine.EmptyRefer[larkmdm.VendorCompanyViewBuilder],
		"refOfVendorCompanyViewBuilder":                   engine.ReferOf[larkmdm.VendorCompanyViewBuilder],
		"unRefVendorCompanyViewBuilder":                   engine.UnRefer[larkmdm.VendorCompanyViewBuilder]}
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
