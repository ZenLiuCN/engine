// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/core'{

	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import * as time from 'golang/time'
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as json from 'golang/encoding/json'
	// @ts-ignore
	import * as http from 'golang/net/http'
	// @ts-ignore
	import * as fmt from 'golang/fmt'
	// @ts-ignore
	import type {Struct,float32,map,int,Ref,error,GoError,bool,Nothing,int16,int32,Alias,float64,int64,int8} from 'go'
	export interface AccessTokenType extends string{

	}
	export const AccessTokenTypeApp:AccessTokenType
	export const AccessTokenTypeNone:AccessTokenType
	export const AccessTokenTypeTenant:AccessTokenType
	export const AccessTokenTypeUser:AccessTokenType
	export interface ApiReq extends Struct<ApiReq>,json.Token{

			httpMethod:string
			apiPath:string
			body:any
			queryParams:QueryParams
			pathParams:PathParams
			supportedAccessTokenTypes:AccessTokenType[]
	}
	export interface ApiResp extends Struct<ApiResp>,fmt.Stringer,json.Token{

			statusCode:int
			header:http.Header
			rawBody:Uint8Array
			write(writer:http.ResponseWriter):void
			jsonUnmarshalBody(val:any,config:Ref<Config>)/*error*/
			requestId():string
			string():string
	}
	//"/open-apis/auth/v3/app_access_token/internal"
	export const AppAccessTokenInternalUrlPath:string
	export interface AppAccessTokenResp extends Struct<AppAccessTokenResp>,json.Token{

			apiResp:Ref<ApiResp>
			codeError:CodeError
			expire:int
			appAccessToken:string
			success():bool
	}
	//"/open-apis/auth/v3/app_access_token"
	export const AppAccessTokenUrlPath:string
	export interface AppTicketManager extends Struct<AppTicketManager>,json.Token{

			get(ctx:context.Context,config:Ref<Config>):string
			set(ctx:context.Context,appId:string,value:string,ttl:time.Duration)/*error*/
	}
	export interface AppType extends string{

	}
	export const AppTypeMarketplace:AppType
	export const AppTypeSelfBuilt:AppType
	//"/open-apis/auth/v3/app_ticket/resend"
	export const ApplyAppTicketPath:string
	export function boolPtr(v:bool):Ref<bool>

	export function boolValue(v:Ref<bool>):bool

	export interface Cache{

			get(ctx:context.Context,key:string):string
			set(ctx:context.Context,key:string,value:string,expireTime:time.Duration)/*error*/
	}
	export interface ClientTimeoutError extends Error,GoError,Struct<ClientTimeoutError>{

			error():string
	}
	export interface CodeError extends Struct<CodeError>,Error,GoError{

			code:int
			msg:string
			err:Ref<Struct<{
			
				details:Ref<CodeErrorDetail>[]
				permissionViolations:Ref<CodeErrorPermissionViolation>[]
				fieldViolations:Ref<CodeErrorFieldViolation>[]
			}>>
			error():string
			string():string
	}
	export interface CodeErrorDetail extends Struct<CodeErrorDetail>,Error{

			key:string
			value:string
	}
	export interface CodeErrorFieldViolation extends Struct<CodeErrorFieldViolation>,Error{

			field:string
			value:string
			description:string
	}
	export interface CodeErrorPermissionViolation extends Struct<CodeErrorPermissionViolation>,Error{

			type:string
			subject:string
			description:string
	}
	export interface Config extends Struct<Config>,json.Token{

			baseUrl:string
			appId:string
			appSecret:string
			helpDeskId:string
			helpDeskToken:string
			helpdeskAuthToken:string
			reqTimeout:time.Duration
			logLevel:LogLevel
			httpClient:HttpClient
			logger:Logger
			appType:AppType
			enableTokenCache:bool
			tokenCache:Cache
			logReqAtDebug:bool
			header:http.Header
			serializable:Serializable
			skipSignVerify:bool
	}
	export interface DecryptErr extends Struct<DecryptErr>,Error,GoError{

			message:string
			error():string
	}
	export interface DefaultSerialization extends Alias<Nothing>{

			serialize(v:any):Uint8Array
			deserialize(data:Uint8Array,v:any)/*error*/
	}
	export interface DialFailedError extends Error,GoError,Struct<DialFailedError>{

			error():string
	}
	export function downloadFile(ctx:context.Context,url:string):Uint8Array

	export function encryptedEventMsg(ctx:context.Context,data:any,encryptKey:string):string

	export const ErrAppTicketIsEmpty:GoError
	export function file2Bytes(fileName:string):Uint8Array

	export function fileNameByHeader(header:http.Header):string

	export function float32Ptr(v:float32):Ref<float32>

	export function float32Value(v:Ref<float32>):float32

	export function float64Ptr(v:float64):Ref<float64>

	export function float64Value(v:Ref<float64>):float64

	export interface Formdata extends Struct<Formdata>,json.Token{

			addField(field:string,val:any):Ref<Formdata>
			addFile(field:string,r:io.Reader):Ref<Formdata>
	}
	export function getAppTicketManager():Ref<AppTicketManager>

	export interface HttpClient{

			Do(v1:Ref<http.Request>):Ref<http.Response>
	}
	//"X-Tt-Logid"
	export const HttpHeaderKeyLogId:string
	//"X-Request-Id"
	export const HttpHeaderKeyRequestId:string
	export interface IllegalParamError extends Struct<IllegalParamError>,Error,GoError{

			error():string
	}
	export function int16Ptr(v:int16):Ref<int16>

	export function int16Value(v:Ref<int16>):int16

	export function int32Ptr(v:int32):Ref<int32>

	export function int32Value(v:Ref<int32>):int32

	export function int64Ptr(v:int64):Ref<int64>

	export function int64Value(v:Ref<int64>):int64

	export function int8Ptr(v:int8):Ref<int8>

	export function int8Value(v:Ref<int8>):int8

	export function intPtr(v:int):Ref<int>

	export function intValue(v:Ref<int>):int

	export interface LogLevel extends int{

	}
	export const LogLevelDebug:LogLevel
	export const LogLevelError:LogLevel
	export const LogLevelInfo:LogLevel
	export const LogLevelWarn:LogLevel
	export interface Logger{

			debug(v2:context.Context,...v1:any[]):void
			error(v2:context.Context,...v1:any[]):void
			info(v2:context.Context,...v1:any[]):void
			warn(v2:context.Context,...v1:any[]):void
	}
	export interface MarketplaceAppAccessTokenReq extends json.Token,Struct<MarketplaceAppAccessTokenReq>{

			appID:string
			appSecret:string
			appTicket:string
	}
	export interface MarketplaceTenantAccessTokenReq extends Struct<MarketplaceTenantAccessTokenReq>,json.Token{

			appAccessToken:string
			tenantKey:string
	}
	export function newCache(config:Ref<Config>):void

	export function newEventLogger():Logger

	export function newFormdata():Ref<Formdata>

	export function newHttpClient(config:Ref<Config>):void

	export function newLogger(config:Ref<Config>):void

	export function newSerialization(config:Ref<Config>):void

	export interface PathParams extends map<string,string>{

			get(key:string):string
			set(key:string,value:string):void
	}
	export function prettify(i:any):string

	export interface QueryParams extends map<string,Array<string>>{

			get(key:string):string
			set(key:string,value:string):void
			encode():string
			add(key:string,value:string):void
	}
	export interface ReqTranslator extends Alias<Nothing>{

	}
	export function request(ctx:context.Context,req:Ref<ApiReq>,config:Ref<Config>,...options:RequestOptionFunc[]):Ref<ApiResp>

	export interface RequestOption extends Struct<RequestOption>,json.Token{

			tenantKey:string
			userAccessToken:string
			appAccessToken:string
			tenantAccessToken:string
			needHelpDeskAuth:bool
			requestId:string
			appTicket:string
			fileUpload:bool
			fileDownload:bool
			header:http.Header
	}
	export interface RequestOptionFunc extends Alias<(option:Ref<RequestOption>)=>void>{

	}
	export interface ResendAppTicketReq extends json.Token,Struct<ResendAppTicketReq>{

			appID:string
			appSecret:string
	}
	export interface ResendAppTicketResp extends Struct<ResendAppTicketResp>,json.Token{

			apiResp:Ref<ApiResp>
			codeError:CodeError
			success():bool
	}
	export interface SelfBuiltAppAccessTokenReq extends Struct<SelfBuiltAppAccessTokenReq>,json.Token{

			appID:string
			appSecret:string
	}
	export interface SelfBuiltTenantAccessTokenReq extends Struct<SelfBuiltTenantAccessTokenReq>,json.Token{

			appID:string
			appSecret:string
	}
	export interface Serializable{

			deserialize(data:Uint8Array,v:any)/*error*/
			serialize(v:any):Uint8Array
	}
	export interface ServerTimeoutError extends Struct<ServerTimeoutError>,Error,GoError{

			error():string
	}
	export function stringPtr(v:string):Ref<string>

	export function stringValue(v:Ref<string>):string

	export function structToMap(val:any):[map<string,any>]

	//"/open-apis/auth/v3/tenant_access_token/internal"
	export const TenantAccessTokenInternalUrlPath:string
	export interface TenantAccessTokenResp extends Struct<TenantAccessTokenResp>,json.Token{

			apiResp:Ref<ApiResp>
			codeError:CodeError
			expire:int
			tenantAccessToken:string
			success():bool
	}
	//"/open-apis/auth/v3/tenant_access_token"
	export const TenantAccessTokenUrlPath:string
	export function timePtr(v:time.Time):Ref<time.Time>

	export function timeValue(v:Ref<time.Time>):time.Time

	export interface TokenManager extends Struct<TokenManager>,json.Token{

	}
	export interface Value extends Struct<Value>,json.Token{

	}
	export function withAppTicket(appTicket:string):RequestOptionFunc

	export function withFileDownload():RequestOptionFunc

	export function withFileUpload():RequestOptionFunc

	export function withHeaders(header:http.Header):RequestOptionFunc

	export function withNeedHelpDeskAuth():RequestOptionFunc

	export function withRequestId(requestId:string):RequestOptionFunc

	export function withTenantAccessToken(tenantAccessToken:string):RequestOptionFunc

	export function withTenantKey(tenantKey:string):RequestOptionFunc

	export function withUserAccessToken(userAccessToken:string):RequestOptionFunc

	export function emptyCodeErrorDetail():CodeErrorDetail
	export function emptyRefCodeErrorDetail():Ref<CodeErrorDetail>
	export function refOfCodeErrorDetail(x:CodeErrorDetail,v:Ref<CodeErrorDetail>)
	export function unRefCodeErrorDetail(v:Ref<CodeErrorDetail>):CodeErrorDetail
	export function emptyAppAccessTokenResp():AppAccessTokenResp
	export function emptyRefAppAccessTokenResp():Ref<AppAccessTokenResp>
	export function refOfAppAccessTokenResp(x:AppAccessTokenResp,v:Ref<AppAccessTokenResp>)
	export function unRefAppAccessTokenResp(v:Ref<AppAccessTokenResp>):AppAccessTokenResp
	export function emptyAppTicketManager():AppTicketManager
	export function emptyRefAppTicketManager():Ref<AppTicketManager>
	export function refOfAppTicketManager(x:AppTicketManager,v:Ref<AppTicketManager>)
	export function unRefAppTicketManager(v:Ref<AppTicketManager>):AppTicketManager
	export function emptyCodeErrorFieldViolation():CodeErrorFieldViolation
	export function emptyRefCodeErrorFieldViolation():Ref<CodeErrorFieldViolation>
	export function refOfCodeErrorFieldViolation(x:CodeErrorFieldViolation,v:Ref<CodeErrorFieldViolation>)
	export function unRefCodeErrorFieldViolation(v:Ref<CodeErrorFieldViolation>):CodeErrorFieldViolation
	export function emptyCodeErrorPermissionViolation():CodeErrorPermissionViolation
	export function emptyRefCodeErrorPermissionViolation():Ref<CodeErrorPermissionViolation>
	export function refOfCodeErrorPermissionViolation(x:CodeErrorPermissionViolation,v:Ref<CodeErrorPermissionViolation>)
	export function unRefCodeErrorPermissionViolation(v:Ref<CodeErrorPermissionViolation>):CodeErrorPermissionViolation
	export function emptyMarketplaceTenantAccessTokenReq():MarketplaceTenantAccessTokenReq
	export function emptyRefMarketplaceTenantAccessTokenReq():Ref<MarketplaceTenantAccessTokenReq>
	export function refOfMarketplaceTenantAccessTokenReq(x:MarketplaceTenantAccessTokenReq,v:Ref<MarketplaceTenantAccessTokenReq>)
	export function unRefMarketplaceTenantAccessTokenReq(v:Ref<MarketplaceTenantAccessTokenReq>):MarketplaceTenantAccessTokenReq
	export function emptyRequestOption():RequestOption
	export function emptyRefRequestOption():Ref<RequestOption>
	export function refOfRequestOption(x:RequestOption,v:Ref<RequestOption>)
	export function unRefRequestOption(v:Ref<RequestOption>):RequestOption
	export function emptySelfBuiltAppAccessTokenReq():SelfBuiltAppAccessTokenReq
	export function emptyRefSelfBuiltAppAccessTokenReq():Ref<SelfBuiltAppAccessTokenReq>
	export function refOfSelfBuiltAppAccessTokenReq(x:SelfBuiltAppAccessTokenReq,v:Ref<SelfBuiltAppAccessTokenReq>)
	export function unRefSelfBuiltAppAccessTokenReq(v:Ref<SelfBuiltAppAccessTokenReq>):SelfBuiltAppAccessTokenReq
	export function emptySelfBuiltTenantAccessTokenReq():SelfBuiltTenantAccessTokenReq
	export function emptyRefSelfBuiltTenantAccessTokenReq():Ref<SelfBuiltTenantAccessTokenReq>
	export function refOfSelfBuiltTenantAccessTokenReq(x:SelfBuiltTenantAccessTokenReq,v:Ref<SelfBuiltTenantAccessTokenReq>)
	export function unRefSelfBuiltTenantAccessTokenReq(v:Ref<SelfBuiltTenantAccessTokenReq>):SelfBuiltTenantAccessTokenReq
	export function emptyApiReq():ApiReq
	export function emptyRefApiReq():Ref<ApiReq>
	export function refOfApiReq(x:ApiReq,v:Ref<ApiReq>)
	export function unRefApiReq(v:Ref<ApiReq>):ApiReq
	export function emptyConfig():Config
	export function emptyRefConfig():Ref<Config>
	export function refOfConfig(x:Config,v:Ref<Config>)
	export function unRefConfig(v:Ref<Config>):Config
	export function emptyResendAppTicketReq():ResendAppTicketReq
	export function emptyRefResendAppTicketReq():Ref<ResendAppTicketReq>
	export function refOfResendAppTicketReq(x:ResendAppTicketReq,v:Ref<ResendAppTicketReq>)
	export function unRefResendAppTicketReq(v:Ref<ResendAppTicketReq>):ResendAppTicketReq
	export function emptyTenantAccessTokenResp():TenantAccessTokenResp
	export function emptyRefTenantAccessTokenResp():Ref<TenantAccessTokenResp>
	export function refOfTenantAccessTokenResp(x:TenantAccessTokenResp,v:Ref<TenantAccessTokenResp>)
	export function unRefTenantAccessTokenResp(v:Ref<TenantAccessTokenResp>):TenantAccessTokenResp
	export function emptyTokenManager():TokenManager
	export function emptyRefTokenManager():Ref<TokenManager>
	export function refOfTokenManager(x:TokenManager,v:Ref<TokenManager>)
	export function unRefTokenManager(v:Ref<TokenManager>):TokenManager
	export function emptyValue():Value
	export function emptyRefValue():Ref<Value>
	export function refOfValue(x:Value,v:Ref<Value>)
	export function unRefValue(v:Ref<Value>):Value
	export function emptyResendAppTicketResp():ResendAppTicketResp
	export function emptyRefResendAppTicketResp():Ref<ResendAppTicketResp>
	export function refOfResendAppTicketResp(x:ResendAppTicketResp,v:Ref<ResendAppTicketResp>)
	export function unRefResendAppTicketResp(v:Ref<ResendAppTicketResp>):ResendAppTicketResp
	export function emptyApiResp():ApiResp
	export function emptyRefApiResp():Ref<ApiResp>
	export function refOfApiResp(x:ApiResp,v:Ref<ApiResp>)
	export function unRefApiResp(v:Ref<ApiResp>):ApiResp
	export function emptyFormdata():Formdata
	export function emptyRefFormdata():Ref<Formdata>
	export function refOfFormdata(x:Formdata,v:Ref<Formdata>)
	export function unRefFormdata(v:Ref<Formdata>):Formdata
	export function emptyMarketplaceAppAccessTokenReq():MarketplaceAppAccessTokenReq
	export function emptyRefMarketplaceAppAccessTokenReq():Ref<MarketplaceAppAccessTokenReq>
	export function refOfMarketplaceAppAccessTokenReq(x:MarketplaceAppAccessTokenReq,v:Ref<MarketplaceAppAccessTokenReq>)
	export function unRefMarketplaceAppAccessTokenReq(v:Ref<MarketplaceAppAccessTokenReq>):MarketplaceAppAccessTokenReq
}