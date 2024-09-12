// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/service/auth/v3'{

	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import type {error,bool,Alias,Nothing,Struct,Ref} from 'go'
	export interface AppAccessToken extends Alias<Nothing>{

	}
	export interface AppTicket extends Alias<Nothing>{

	}
	export interface CreateAppAccessTokenPathReqBodyBuilder extends Struct<CreateAppAccessTokenPathReqBodyBuilder>{

			appId(appId:string):Ref<CreateAppAccessTokenPathReqBodyBuilder>
			appSecret(appSecret:string):Ref<CreateAppAccessTokenPathReqBodyBuilder>
			appTicket(appTicket:string):Ref<CreateAppAccessTokenPathReqBodyBuilder>
			build():Ref<CreateAppAccessTokenReqBody>
	}
	export interface CreateAppAccessTokenReq extends Struct<CreateAppAccessTokenReq>{

			body:Ref<CreateAppAccessTokenReqBody>
	}
	export interface CreateAppAccessTokenReqBody extends Struct<CreateAppAccessTokenReqBody>{

			appId:Ref<string>
			appSecret:Ref<string>
			appTicket:Ref<string>
	}
	export interface CreateAppAccessTokenReqBodyBuilder extends Struct<CreateAppAccessTokenReqBodyBuilder>{

			appId(appId:string):Ref<CreateAppAccessTokenReqBodyBuilder>
			appSecret(appSecret:string):Ref<CreateAppAccessTokenReqBodyBuilder>
			appTicket(appTicket:string):Ref<CreateAppAccessTokenReqBodyBuilder>
			build():Ref<CreateAppAccessTokenReqBody>
	}
	export interface CreateAppAccessTokenReqBuilder extends Struct<CreateAppAccessTokenReqBuilder>{

			body(body:Ref<CreateAppAccessTokenReqBody>):Ref<CreateAppAccessTokenReqBuilder>
			build():Ref<CreateAppAccessTokenReq>
	}
	export interface CreateAppAccessTokenResp extends Struct<CreateAppAccessTokenResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			success():bool
	}
	export interface CreateTenantAccessTokenPathReqBodyBuilder extends Struct<CreateTenantAccessTokenPathReqBodyBuilder>{

			appAccessToken(appAccessToken:string):Ref<CreateTenantAccessTokenPathReqBodyBuilder>
			tenantKey(tenantKey:string):Ref<CreateTenantAccessTokenPathReqBodyBuilder>
			build():Ref<CreateTenantAccessTokenReqBody>
	}
	export interface CreateTenantAccessTokenReq extends Struct<CreateTenantAccessTokenReq>{

			body:Ref<CreateTenantAccessTokenReqBody>
	}
	export interface CreateTenantAccessTokenReqBody extends Struct<CreateTenantAccessTokenReqBody>{

			appAccessToken:Ref<string>
			tenantKey:Ref<string>
	}
	export interface CreateTenantAccessTokenReqBodyBuilder extends Struct<CreateTenantAccessTokenReqBodyBuilder>{

			appAccessToken(appAccessToken:string):Ref<CreateTenantAccessTokenReqBodyBuilder>
			tenantKey(tenantKey:string):Ref<CreateTenantAccessTokenReqBodyBuilder>
			build():Ref<CreateTenantAccessTokenReqBody>
	}
	export interface CreateTenantAccessTokenReqBuilder extends Struct<CreateTenantAccessTokenReqBuilder>{

			body(body:Ref<CreateTenantAccessTokenReqBody>):Ref<CreateTenantAccessTokenReqBuilder>
			build():Ref<CreateTenantAccessTokenReq>
	}
	export interface CreateTenantAccessTokenResp extends Struct<CreateTenantAccessTokenResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			success():bool
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
	export interface InternalAppAccessTokenPathReqBodyBuilder extends Struct<InternalAppAccessTokenPathReqBodyBuilder>{

			appId(appId:string):Ref<InternalAppAccessTokenPathReqBodyBuilder>
			appSecret(appSecret:string):Ref<InternalAppAccessTokenPathReqBodyBuilder>
			build():Ref<InternalAppAccessTokenReqBody>
	}
	export interface InternalAppAccessTokenReq extends Struct<InternalAppAccessTokenReq>{

			body:Ref<InternalAppAccessTokenReqBody>
	}
	export interface InternalAppAccessTokenReqBody extends Struct<InternalAppAccessTokenReqBody>{

			appId:Ref<string>
			appSecret:Ref<string>
	}
	export interface InternalAppAccessTokenReqBodyBuilder extends Struct<InternalAppAccessTokenReqBodyBuilder>{

			appId(appId:string):Ref<InternalAppAccessTokenReqBodyBuilder>
			appSecret(appSecret:string):Ref<InternalAppAccessTokenReqBodyBuilder>
			build():Ref<InternalAppAccessTokenReqBody>
	}
	export interface InternalAppAccessTokenReqBuilder extends Struct<InternalAppAccessTokenReqBuilder>{

			body(body:Ref<InternalAppAccessTokenReqBody>):Ref<InternalAppAccessTokenReqBuilder>
			build():Ref<InternalAppAccessTokenReq>
	}
	export interface InternalAppAccessTokenResp extends Struct<InternalAppAccessTokenResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			success():bool
	}
	export interface InternalTenantAccessTokenPathReqBodyBuilder extends Struct<InternalTenantAccessTokenPathReqBodyBuilder>{

			appId(appId:string):Ref<InternalTenantAccessTokenPathReqBodyBuilder>
			appSecret(appSecret:string):Ref<InternalTenantAccessTokenPathReqBodyBuilder>
			build():Ref<InternalTenantAccessTokenReqBody>
	}
	export interface InternalTenantAccessTokenReq extends Struct<InternalTenantAccessTokenReq>{

			body:Ref<InternalTenantAccessTokenReqBody>
	}
	export interface InternalTenantAccessTokenReqBody extends Struct<InternalTenantAccessTokenReqBody>{

			appId:Ref<string>
			appSecret:Ref<string>
	}
	export interface InternalTenantAccessTokenReqBodyBuilder extends Struct<InternalTenantAccessTokenReqBodyBuilder>{

			appId(appId:string):Ref<InternalTenantAccessTokenReqBodyBuilder>
			appSecret(appSecret:string):Ref<InternalTenantAccessTokenReqBodyBuilder>
			build():Ref<InternalTenantAccessTokenReqBody>
	}
	export interface InternalTenantAccessTokenReqBuilder extends Struct<InternalTenantAccessTokenReqBuilder>{

			body(body:Ref<InternalTenantAccessTokenReqBody>):Ref<InternalTenantAccessTokenReqBuilder>
			build():Ref<InternalTenantAccessTokenReq>
	}
	export interface InternalTenantAccessTokenResp extends Struct<InternalTenantAccessTokenResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			success():bool
	}
	export function New(config:Ref<larkcore.Config>):Ref<V3>

	export function newCreateAppAccessTokenPathReqBodyBuilder():Ref<CreateAppAccessTokenPathReqBodyBuilder>

	export function newCreateAppAccessTokenReqBodyBuilder():Ref<CreateAppAccessTokenReqBodyBuilder>

	export function newCreateAppAccessTokenReqBuilder():Ref<CreateAppAccessTokenReqBuilder>

	export function newCreateTenantAccessTokenPathReqBodyBuilder():Ref<CreateTenantAccessTokenPathReqBodyBuilder>

	export function newCreateTenantAccessTokenReqBodyBuilder():Ref<CreateTenantAccessTokenReqBodyBuilder>

	export function newCreateTenantAccessTokenReqBuilder():Ref<CreateTenantAccessTokenReqBuilder>

	export function newDepartmentIdBuilder():Ref<DepartmentIdBuilder>

	export function newInternalAppAccessTokenPathReqBodyBuilder():Ref<InternalAppAccessTokenPathReqBodyBuilder>

	export function newInternalAppAccessTokenReqBodyBuilder():Ref<InternalAppAccessTokenReqBodyBuilder>

	export function newInternalAppAccessTokenReqBuilder():Ref<InternalAppAccessTokenReqBuilder>

	export function newInternalTenantAccessTokenPathReqBodyBuilder():Ref<InternalTenantAccessTokenPathReqBodyBuilder>

	export function newInternalTenantAccessTokenReqBodyBuilder():Ref<InternalTenantAccessTokenReqBodyBuilder>

	export function newInternalTenantAccessTokenReqBuilder():Ref<InternalTenantAccessTokenReqBuilder>

	export function newResendAppTicketPathReqBodyBuilder():Ref<ResendAppTicketPathReqBodyBuilder>

	export function newResendAppTicketReqBodyBuilder():Ref<ResendAppTicketReqBodyBuilder>

	export function newResendAppTicketReqBuilder():Ref<ResendAppTicketReqBuilder>

	export function newRevokeTokenEventBuilder():Ref<RevokeTokenEventBuilder>

	export interface ResendAppTicketPathReqBodyBuilder extends Struct<ResendAppTicketPathReqBodyBuilder>{

			appId(appId:string):Ref<ResendAppTicketPathReqBodyBuilder>
			appSecret(appSecret:string):Ref<ResendAppTicketPathReqBodyBuilder>
			build():Ref<ResendAppTicketReqBody>
	}
	export interface ResendAppTicketReq extends Struct<ResendAppTicketReq>{

			body:Ref<ResendAppTicketReqBody>
	}
	export interface ResendAppTicketReqBody extends Struct<ResendAppTicketReqBody>{

			appId:Ref<string>
			appSecret:Ref<string>
	}
	export interface ResendAppTicketReqBodyBuilder extends Struct<ResendAppTicketReqBodyBuilder>{

			appId(appId:string):Ref<ResendAppTicketReqBodyBuilder>
			appSecret(appSecret:string):Ref<ResendAppTicketReqBodyBuilder>
			build():Ref<ResendAppTicketReqBody>
	}
	export interface ResendAppTicketReqBuilder extends Struct<ResendAppTicketReqBuilder>{

			body(body:Ref<ResendAppTicketReqBody>):Ref<ResendAppTicketReqBuilder>
			build():Ref<ResendAppTicketReq>
	}
	export interface ResendAppTicketResp extends Struct<ResendAppTicketResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			success():bool
	}
	export interface RevokeTokenEvent extends Struct<RevokeTokenEvent>{

			revokeTokenType:Ref<string>
			revokeReason:Ref<string>
			openId:Ref<string>
			unionId:Ref<string>
			userId:Ref<string>
	}
	export interface RevokeTokenEventBuilder extends Struct<RevokeTokenEventBuilder>{

			revokeTokenType(revokeTokenType:string):Ref<RevokeTokenEventBuilder>
			revokeReason(revokeReason:string):Ref<RevokeTokenEventBuilder>
			openId(openId:string):Ref<RevokeTokenEventBuilder>
			unionId(unionId:string):Ref<RevokeTokenEventBuilder>
			userId(userId:string):Ref<RevokeTokenEventBuilder>
			build():Ref<RevokeTokenEvent>
	}
	export interface TenantAccessToken extends Alias<Nothing>{

	}
	export interface V3 extends Struct<V3>{

			appAccessToken:Ref<{
			
				create(ctx:context.Context,req:Ref<CreateAppAccessTokenReq>,...options:larkcore.RequestOptionFunc[]):Ref<CreateAppAccessTokenResp>
				internal(ctx:context.Context,req:Ref<InternalAppAccessTokenReq>,...options:larkcore.RequestOptionFunc[]):Ref<InternalAppAccessTokenResp>
			}>
			appTicket:Ref<{
			
				resend(ctx:context.Context,req:Ref<ResendAppTicketReq>,...options:larkcore.RequestOptionFunc[]):Ref<ResendAppTicketResp>
			}>
			tenantAccessToken:Ref<{
			
				create(ctx:context.Context,req:Ref<CreateTenantAccessTokenReq>,...options:larkcore.RequestOptionFunc[]):Ref<CreateTenantAccessTokenResp>
				internal(ctx:context.Context,req:Ref<InternalTenantAccessTokenReq>,...options:larkcore.RequestOptionFunc[]):Ref<InternalTenantAccessTokenResp>
			}>
	}
	export function emptyDepartmentId():DepartmentId
	export function emptyRefDepartmentId():Ref<DepartmentId>
	export function refOfDepartmentId(x:DepartmentId,v:Ref<DepartmentId>)
	export function unRefDepartmentId(v:Ref<DepartmentId>):DepartmentId
	export function emptyCreateAppAccessTokenReq():CreateAppAccessTokenReq
	export function emptyRefCreateAppAccessTokenReq():Ref<CreateAppAccessTokenReq>
	export function refOfCreateAppAccessTokenReq(x:CreateAppAccessTokenReq,v:Ref<CreateAppAccessTokenReq>)
	export function unRefCreateAppAccessTokenReq(v:Ref<CreateAppAccessTokenReq>):CreateAppAccessTokenReq
	export function emptyCreateTenantAccessTokenReq():CreateTenantAccessTokenReq
	export function emptyRefCreateTenantAccessTokenReq():Ref<CreateTenantAccessTokenReq>
	export function refOfCreateTenantAccessTokenReq(x:CreateTenantAccessTokenReq,v:Ref<CreateTenantAccessTokenReq>)
	export function unRefCreateTenantAccessTokenReq(v:Ref<CreateTenantAccessTokenReq>):CreateTenantAccessTokenReq
	export function emptyCreateTenantAccessTokenReqBody():CreateTenantAccessTokenReqBody
	export function emptyRefCreateTenantAccessTokenReqBody():Ref<CreateTenantAccessTokenReqBody>
	export function refOfCreateTenantAccessTokenReqBody(x:CreateTenantAccessTokenReqBody,v:Ref<CreateTenantAccessTokenReqBody>)
	export function unRefCreateTenantAccessTokenReqBody(v:Ref<CreateTenantAccessTokenReqBody>):CreateTenantAccessTokenReqBody
	export function emptyResendAppTicketReqBody():ResendAppTicketReqBody
	export function emptyRefResendAppTicketReqBody():Ref<ResendAppTicketReqBody>
	export function refOfResendAppTicketReqBody(x:ResendAppTicketReqBody,v:Ref<ResendAppTicketReqBody>)
	export function unRefResendAppTicketReqBody(v:Ref<ResendAppTicketReqBody>):ResendAppTicketReqBody
	export function emptyInternalAppAccessTokenReqBody():InternalAppAccessTokenReqBody
	export function emptyRefInternalAppAccessTokenReqBody():Ref<InternalAppAccessTokenReqBody>
	export function refOfInternalAppAccessTokenReqBody(x:InternalAppAccessTokenReqBody,v:Ref<InternalAppAccessTokenReqBody>)
	export function unRefInternalAppAccessTokenReqBody(v:Ref<InternalAppAccessTokenReqBody>):InternalAppAccessTokenReqBody
	export function emptyResendAppTicketResp():ResendAppTicketResp
	export function emptyRefResendAppTicketResp():Ref<ResendAppTicketResp>
	export function refOfResendAppTicketResp(x:ResendAppTicketResp,v:Ref<ResendAppTicketResp>)
	export function unRefResendAppTicketResp(v:Ref<ResendAppTicketResp>):ResendAppTicketResp
	export function emptyInternalTenantAccessTokenResp():InternalTenantAccessTokenResp
	export function emptyRefInternalTenantAccessTokenResp():Ref<InternalTenantAccessTokenResp>
	export function refOfInternalTenantAccessTokenResp(x:InternalTenantAccessTokenResp,v:Ref<InternalTenantAccessTokenResp>)
	export function unRefInternalTenantAccessTokenResp(v:Ref<InternalTenantAccessTokenResp>):InternalTenantAccessTokenResp
	export function emptyInternalTenantAccessTokenReq():InternalTenantAccessTokenReq
	export function emptyRefInternalTenantAccessTokenReq():Ref<InternalTenantAccessTokenReq>
	export function refOfInternalTenantAccessTokenReq(x:InternalTenantAccessTokenReq,v:Ref<InternalTenantAccessTokenReq>)
	export function unRefInternalTenantAccessTokenReq(v:Ref<InternalTenantAccessTokenReq>):InternalTenantAccessTokenReq
	export function emptyCreateAppAccessTokenResp():CreateAppAccessTokenResp
	export function emptyRefCreateAppAccessTokenResp():Ref<CreateAppAccessTokenResp>
	export function refOfCreateAppAccessTokenResp(x:CreateAppAccessTokenResp,v:Ref<CreateAppAccessTokenResp>)
	export function unRefCreateAppAccessTokenResp(v:Ref<CreateAppAccessTokenResp>):CreateAppAccessTokenResp
	export function emptyInternalTenantAccessTokenReqBody():InternalTenantAccessTokenReqBody
	export function emptyRefInternalTenantAccessTokenReqBody():Ref<InternalTenantAccessTokenReqBody>
	export function refOfInternalTenantAccessTokenReqBody(x:InternalTenantAccessTokenReqBody,v:Ref<InternalTenantAccessTokenReqBody>)
	export function unRefInternalTenantAccessTokenReqBody(v:Ref<InternalTenantAccessTokenReqBody>):InternalTenantAccessTokenReqBody
	export function emptyV3():V3
	export function emptyRefV3():Ref<V3>
	export function refOfV3(x:V3,v:Ref<V3>)
	export function unRefV3(v:Ref<V3>):V3
	export function emptyCreateAppAccessTokenReqBody():CreateAppAccessTokenReqBody
	export function emptyRefCreateAppAccessTokenReqBody():Ref<CreateAppAccessTokenReqBody>
	export function refOfCreateAppAccessTokenReqBody(x:CreateAppAccessTokenReqBody,v:Ref<CreateAppAccessTokenReqBody>)
	export function unRefCreateAppAccessTokenReqBody(v:Ref<CreateAppAccessTokenReqBody>):CreateAppAccessTokenReqBody
	export function emptyCreateTenantAccessTokenResp():CreateTenantAccessTokenResp
	export function emptyRefCreateTenantAccessTokenResp():Ref<CreateTenantAccessTokenResp>
	export function refOfCreateTenantAccessTokenResp(x:CreateTenantAccessTokenResp,v:Ref<CreateTenantAccessTokenResp>)
	export function unRefCreateTenantAccessTokenResp(v:Ref<CreateTenantAccessTokenResp>):CreateTenantAccessTokenResp
	export function emptyRevokeTokenEvent():RevokeTokenEvent
	export function emptyRefRevokeTokenEvent():Ref<RevokeTokenEvent>
	export function refOfRevokeTokenEvent(x:RevokeTokenEvent,v:Ref<RevokeTokenEvent>)
	export function unRefRevokeTokenEvent(v:Ref<RevokeTokenEvent>):RevokeTokenEvent
	export function emptyInternalAppAccessTokenReq():InternalAppAccessTokenReq
	export function emptyRefInternalAppAccessTokenReq():Ref<InternalAppAccessTokenReq>
	export function refOfInternalAppAccessTokenReq(x:InternalAppAccessTokenReq,v:Ref<InternalAppAccessTokenReq>)
	export function unRefInternalAppAccessTokenReq(v:Ref<InternalAppAccessTokenReq>):InternalAppAccessTokenReq
	export function emptyResendAppTicketReq():ResendAppTicketReq
	export function emptyRefResendAppTicketReq():Ref<ResendAppTicketReq>
	export function refOfResendAppTicketReq(x:ResendAppTicketReq,v:Ref<ResendAppTicketReq>)
	export function unRefResendAppTicketReq(v:Ref<ResendAppTicketReq>):ResendAppTicketReq
	export function emptyInternalAppAccessTokenResp():InternalAppAccessTokenResp
	export function emptyRefInternalAppAccessTokenResp():Ref<InternalAppAccessTokenResp>
	export function refOfInternalAppAccessTokenResp(x:InternalAppAccessTokenResp,v:Ref<InternalAppAccessTokenResp>)
	export function unRefInternalAppAccessTokenResp(v:Ref<InternalAppAccessTokenResp>):InternalAppAccessTokenResp
}