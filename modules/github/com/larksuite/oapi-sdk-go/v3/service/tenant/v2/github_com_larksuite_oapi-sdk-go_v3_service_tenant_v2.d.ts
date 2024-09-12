// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/service/tenant/v2'{

	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import type {Ref,Struct,bool,int,error} from 'go'
	export interface Avatar extends Struct<Avatar>{

			avatarOrigin:Ref<string>
			avatar72:Ref<string>
			avatar240:Ref<string>
			avatar640:Ref<string>
	}
	export interface AvatarBuilder extends Struct<AvatarBuilder>{

			avatarOrigin(avatarOrigin:string):Ref<AvatarBuilder>
			avatar72(avatar72:string):Ref<AvatarBuilder>
			avatar240(avatar240:string):Ref<AvatarBuilder>
			avatar640(avatar640:string):Ref<AvatarBuilder>
			build():Ref<Avatar>
	}
	export interface DepartmentId extends Struct<DepartmentId>{

			departmentId:Ref<string>
			openDepartmentId:Ref<string>
	}
	export interface DepartmentIdBuilder extends Struct<DepartmentIdBuilder>{

			departmentId(departmentId:string):Ref<DepartmentIdBuilder>
			openDepartmentId(openDepartmentId:string):Ref<DepartmentIdBuilder>
			build():Ref<DepartmentId>
	}
	export function New(config:Ref<larkcore.Config>):Ref<V2>

	export function newAvatarBuilder():Ref<AvatarBuilder>

	export function newDepartmentIdBuilder():Ref<DepartmentIdBuilder>

	export function newProductI18nNameBuilder():Ref<ProductI18nNameBuilder>

	export function newTenantAssignInfoBuilder():Ref<TenantAssignInfoBuilder>

	export function newTenantBuilder():Ref<TenantBuilder>

	export interface ProductI18nName extends Struct<ProductI18nName>{

			zhCn:Ref<string>
			jaJp:Ref<string>
			enUs:Ref<string>
	}
	export interface ProductI18nNameBuilder extends Struct<ProductI18nNameBuilder>{

			zhCn(zhCn:string):Ref<ProductI18nNameBuilder>
			jaJp(jaJp:string):Ref<ProductI18nNameBuilder>
			enUs(enUs:string):Ref<ProductI18nNameBuilder>
			build():Ref<ProductI18nName>
	}
	export interface QueryTenantProductAssignInfoResp extends Struct<QueryTenantProductAssignInfoResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<QueryTenantProductAssignInfoRespData>
			success():bool
	}
	export interface QueryTenantProductAssignInfoRespData extends Struct<QueryTenantProductAssignInfoRespData>{

			assignInfoList:Ref<TenantAssignInfo>[]
	}
	export interface QueryTenantResp extends Struct<QueryTenantResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<QueryTenantRespData>
			success():bool
	}
	export interface QueryTenantRespData extends Struct<QueryTenantRespData>{

			tenant:Ref<Tenant>
	}
	export interface Tenant extends Struct<Tenant>{

			name:Ref<string>
			displayId:Ref<string>
			tenantTag:Ref<int>
			tenantKey:Ref<string>
			avatar:Ref<Avatar>
			domain:Ref<string>
	}
	export interface TenantAssignInfo extends Struct<TenantAssignInfo>{

			subscriptionId:Ref<string>
			licensePlanKey:Ref<string>
			productName:Ref<string>
			i18nName:Ref<ProductI18nName>
			totalSeats:Ref<string>
			assignedSeats:Ref<string>
			startTime:Ref<string>
			endTime:Ref<string>
	}
	export interface TenantAssignInfoBuilder extends Struct<TenantAssignInfoBuilder>{

			subscriptionId(subscriptionId:string):Ref<TenantAssignInfoBuilder>
			licensePlanKey(licensePlanKey:string):Ref<TenantAssignInfoBuilder>
			productName(productName:string):Ref<TenantAssignInfoBuilder>
			i18nName(i18nName:Ref<ProductI18nName>):Ref<TenantAssignInfoBuilder>
			totalSeats(totalSeats:string):Ref<TenantAssignInfoBuilder>
			assignedSeats(assignedSeats:string):Ref<TenantAssignInfoBuilder>
			startTime(startTime:string):Ref<TenantAssignInfoBuilder>
			endTime(endTime:string):Ref<TenantAssignInfoBuilder>
			build():Ref<TenantAssignInfo>
	}
	export interface TenantBuilder extends Struct<TenantBuilder>{

			name(name:string):Ref<TenantBuilder>
			displayId(displayId:string):Ref<TenantBuilder>
			tenantTag(tenantTag:int):Ref<TenantBuilder>
			tenantKey(tenantKey:string):Ref<TenantBuilder>
			avatar(avatar:Ref<Avatar>):Ref<TenantBuilder>
			domain(domain:string):Ref<TenantBuilder>
			build():Ref<Tenant>
	}
	export interface V2 extends Struct<V2>{

			tenant:Ref<{
			
				query(ctx:context.Context,...options:larkcore.RequestOptionFunc[]):Ref<QueryTenantResp>
			}>
			tenantProductAssignInfo:Ref<{
			
				query(ctx:context.Context,...options:larkcore.RequestOptionFunc[]):Ref<QueryTenantProductAssignInfoResp>
			}>
	}
	export function emptyDepartmentId():DepartmentId
	export function emptyRefDepartmentId():Ref<DepartmentId>
	export function refOfDepartmentId(x:DepartmentId,v:Ref<DepartmentId>)
	export function unRefDepartmentId(v:Ref<DepartmentId>):DepartmentId
	export function emptyQueryTenantProductAssignInfoResp():QueryTenantProductAssignInfoResp
	export function emptyRefQueryTenantProductAssignInfoResp():Ref<QueryTenantProductAssignInfoResp>
	export function refOfQueryTenantProductAssignInfoResp(x:QueryTenantProductAssignInfoResp,v:Ref<QueryTenantProductAssignInfoResp>)
	export function unRefQueryTenantProductAssignInfoResp(v:Ref<QueryTenantProductAssignInfoResp>):QueryTenantProductAssignInfoResp
	export function emptyQueryTenantProductAssignInfoRespData():QueryTenantProductAssignInfoRespData
	export function emptyRefQueryTenantProductAssignInfoRespData():Ref<QueryTenantProductAssignInfoRespData>
	export function refOfQueryTenantProductAssignInfoRespData(x:QueryTenantProductAssignInfoRespData,v:Ref<QueryTenantProductAssignInfoRespData>)
	export function unRefQueryTenantProductAssignInfoRespData(v:Ref<QueryTenantProductAssignInfoRespData>):QueryTenantProductAssignInfoRespData
	export function emptyQueryTenantResp():QueryTenantResp
	export function emptyRefQueryTenantResp():Ref<QueryTenantResp>
	export function refOfQueryTenantResp(x:QueryTenantResp,v:Ref<QueryTenantResp>)
	export function unRefQueryTenantResp(v:Ref<QueryTenantResp>):QueryTenantResp
	export function emptyQueryTenantRespData():QueryTenantRespData
	export function emptyRefQueryTenantRespData():Ref<QueryTenantRespData>
	export function refOfQueryTenantRespData(x:QueryTenantRespData,v:Ref<QueryTenantRespData>)
	export function unRefQueryTenantRespData(v:Ref<QueryTenantRespData>):QueryTenantRespData
	export function emptyTenant():Tenant
	export function emptyRefTenant():Ref<Tenant>
	export function refOfTenant(x:Tenant,v:Ref<Tenant>)
	export function unRefTenant(v:Ref<Tenant>):Tenant
	export function emptyTenantAssignInfo():TenantAssignInfo
	export function emptyRefTenantAssignInfo():Ref<TenantAssignInfo>
	export function refOfTenantAssignInfo(x:TenantAssignInfo,v:Ref<TenantAssignInfo>)
	export function unRefTenantAssignInfo(v:Ref<TenantAssignInfo>):TenantAssignInfo
	export function emptyV2():V2
	export function emptyRefV2():Ref<V2>
	export function refOfV2(x:V2,v:Ref<V2>)
	export function unRefV2(v:Ref<V2>):V2
	export function emptyProductI18nName():ProductI18nName
	export function emptyRefProductI18nName():Ref<ProductI18nName>
	export function refOfProductI18nName(x:ProductI18nName,v:Ref<ProductI18nName>)
	export function unRefProductI18nName(v:Ref<ProductI18nName>):ProductI18nName
	export function emptyAvatar():Avatar
	export function emptyRefAvatar():Ref<Avatar>
	export function refOfAvatar(x:Avatar,v:Ref<Avatar>)
	export function unRefAvatar(v:Ref<Avatar>):Avatar
}