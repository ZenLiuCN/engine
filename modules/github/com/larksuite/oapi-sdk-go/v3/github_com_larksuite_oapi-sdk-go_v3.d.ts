// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3'{

	// @ts-ignore
	import * as security_and_compliance from 'github.com/larksuite/oapi-sdk-go/v3/service/security_and_compliance'
	// @ts-ignore
	import * as admin from 'github.com/larksuite/oapi-sdk-go/v3/service/admin'
	// @ts-ignore
	import * as bitable from 'github.com/larksuite/oapi-sdk-go/v3/service/bitable'
	// @ts-ignore
	import * as helpdesk from 'github.com/larksuite/oapi-sdk-go/v3/service/helpdesk'
	// @ts-ignore
	import * as calendar from 'github.com/larksuite/oapi-sdk-go/v3/service/calendar'
	// @ts-ignore
	import * as authen from 'github.com/larksuite/oapi-sdk-go/v3/service/authen'
	// @ts-ignore
	import * as auth from 'github.com/larksuite/oapi-sdk-go/v3/service/auth'
	// @ts-ignore
	import * as ehr from 'github.com/larksuite/oapi-sdk-go/v3/service/ehr'
	// @ts-ignore
	import * as meeting_room from 'github.com/larksuite/oapi-sdk-go/v3/service/meeting_room'
	// @ts-ignore
	import * as docx from 'github.com/larksuite/oapi-sdk-go/v3/service/docx'
	// @ts-ignore
	import * as optical_char_recognition from 'github.com/larksuite/oapi-sdk-go/v3/service/optical_char_recognition'
	// @ts-ignore
	import * as http from 'golang/net/http'
	// @ts-ignore
	import * as human_authentication from 'github.com/larksuite/oapi-sdk-go/v3/service/human_authentication'
	// @ts-ignore
	import * as event from 'github.com/larksuite/oapi-sdk-go/v3/service/event'
	// @ts-ignore
	import * as lingo from 'github.com/larksuite/oapi-sdk-go/v3/service/lingo'
	// @ts-ignore
	import * as block from 'github.com/larksuite/oapi-sdk-go/v3/service/block'
	// @ts-ignore
	import * as mail from 'github.com/larksuite/oapi-sdk-go/v3/service/mail'
	// @ts-ignore
	import * as wiki from 'github.com/larksuite/oapi-sdk-go/v3/service/wiki'
	// @ts-ignore
	import * as baike from 'github.com/larksuite/oapi-sdk-go/v3/service/baike'
	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as corehr from 'github.com/larksuite/oapi-sdk-go/v3/service/corehr'
	// @ts-ignore
	import * as search from 'github.com/larksuite/oapi-sdk-go/v3/service/search'
	// @ts-ignore
	import * as document_ai from 'github.com/larksuite/oapi-sdk-go/v3/service/document_ai'
	// @ts-ignore
	import * as passport from 'github.com/larksuite/oapi-sdk-go/v3/service/passport'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import * as im from 'github.com/larksuite/oapi-sdk-go/v3/service/im'
	// @ts-ignore
	import * as sheets from 'github.com/larksuite/oapi-sdk-go/v3/service/sheets'
	// @ts-ignore
	import * as larkext from 'github.com/larksuite/oapi-sdk-go/v3/service/ext'
	// @ts-ignore
	import * as time from 'golang/time'
	// @ts-ignore
	import * as tenant from 'github.com/larksuite/oapi-sdk-go/v3/service/tenant'
	// @ts-ignore
	import * as workplace from 'github.com/larksuite/oapi-sdk-go/v3/service/workplace'
	// @ts-ignore
	import * as vc from 'github.com/larksuite/oapi-sdk-go/v3/service/vc'
	// @ts-ignore
	import * as speech_to_text from 'github.com/larksuite/oapi-sdk-go/v3/service/speech_to_text'
	// @ts-ignore
	import * as task from 'github.com/larksuite/oapi-sdk-go/v3/service/task'
	// @ts-ignore
	import * as gray_test_open_sg from 'github.com/larksuite/oapi-sdk-go/v3/service/gray_test_open_sg'
	// @ts-ignore
	import * as translation from 'github.com/larksuite/oapi-sdk-go/v3/service/translation'
	// @ts-ignore
	import * as contact from 'github.com/larksuite/oapi-sdk-go/v3/service/contact'
	// @ts-ignore
	import * as drive from 'github.com/larksuite/oapi-sdk-go/v3/service/drive'
	// @ts-ignore
	import * as personal_settings from 'github.com/larksuite/oapi-sdk-go/v3/service/personal_settings'
	// @ts-ignore
	import * as verification from 'github.com/larksuite/oapi-sdk-go/v3/service/verification'
	// @ts-ignore
	import * as json from 'golang/encoding/json'
	// @ts-ignore
	import * as board from 'github.com/larksuite/oapi-sdk-go/v3/service/board'
	// @ts-ignore
	import * as okr from 'github.com/larksuite/oapi-sdk-go/v3/service/okr'
	// @ts-ignore
	import * as aily from 'github.com/larksuite/oapi-sdk-go/v3/service/aily'
	// @ts-ignore
	import * as attendance from 'github.com/larksuite/oapi-sdk-go/v3/service/attendance'
	// @ts-ignore
	import * as application from 'github.com/larksuite/oapi-sdk-go/v3/service/application'
	// @ts-ignore
	import * as approval from 'github.com/larksuite/oapi-sdk-go/v3/service/approval'
	// @ts-ignore
	import * as report from 'github.com/larksuite/oapi-sdk-go/v3/service/report'
	// @ts-ignore
	import * as hire from 'github.com/larksuite/oapi-sdk-go/v3/service/hire'
	// @ts-ignore
	import * as mdm from 'github.com/larksuite/oapi-sdk-go/v3/service/mdm'
	// @ts-ignore
	import * as acs from 'github.com/larksuite/oapi-sdk-go/v3/service/acs'
	// @ts-ignore
	import type {Ref,Struct,error,Alias,bool} from 'go'
	export interface Client extends Struct<Client>,json.Token{

			humanAuthentication:Ref<human_authentication.Service>
			task:Ref<task.Service>
			calendar:Ref<calendar.Service>
			corehr:Ref<corehr.Service>
			im:Ref<im.Service>
			search:Ref<search.Service>
			approval:Ref<approval.Service>
			authen:Ref<authen.Service>
			event:Ref<event.Service>
			personalSettings:Ref<personal_settings.Service>
			report:Ref<report.Service>
			tenant:Ref<tenant.Service>
			auth:Ref<auth.Service>
			board:Ref<board.Service>
			ehr:Ref<ehr.Service>
			grayTestOpenSg:Ref<gray_test_open_sg.Service>
			hire:Ref<hire.Service>
			mdm:Ref<mdm.Service>
			meetingRoom:Ref<meeting_room.Service>
			securityAndCompliance:Ref<security_and_compliance.Service>
			acs:Ref<acs.Service>
			documentAi:Ref<document_ai.Service>
			wiki:Ref<wiki.Service>
			workplace:Ref<workplace.Service>
			translation:Ref<translation.Service>
			vc:Ref<vc.Service>
			lingo:Ref<lingo.Service>
			okr:Ref<okr.Service>
			admin:Ref<admin.Service>
			docx:Ref<docx.Service>
			block:Ref<block.Service>
			mail:Ref<mail.Service>
			baike:Ref<baike.Service>
			contact:Ref<contact.Service>
			helpdesk:Ref<helpdesk.Service>
			opticalCharRecognition:Ref<optical_char_recognition.Service>
			sheets:Ref<sheets.Service>
			speechToText:Ref<speech_to_text.Service>
			aily:Ref<aily.Service>
			attendance:Ref<attendance.Service>
			verification:Ref<verification.Service>
			drive:Ref<drive.Service>
			passport:Ref<passport.Service>
			application:Ref<application.Service>
			bitable:Ref<bitable.Service>
			ext:Ref<larkext.ExtService>
			post(ctx:context.Context,httpPath:string,body:any,accessTokeType:larkcore.AccessTokenType,...options:larkcore.RequestOptionFunc[]):Ref<larkcore.ApiResp>
			Do(ctx:context.Context,apiReq:Ref<larkcore.ApiReq>,...options:larkcore.RequestOptionFunc[]):Ref<larkcore.ApiResp>
			get(ctx:context.Context,httpPath:string,body:any,accessTokeType:larkcore.AccessTokenType,...options:larkcore.RequestOptionFunc[]):Ref<larkcore.ApiResp>
			delete(ctx:context.Context,httpPath:string,body:any,accessTokeType:larkcore.AccessTokenType,...options:larkcore.RequestOptionFunc[]):Ref<larkcore.ApiResp>
			put(ctx:context.Context,httpPath:string,body:any,accessTokeType:larkcore.AccessTokenType,...options:larkcore.RequestOptionFunc[]):Ref<larkcore.ApiResp>
			patch(ctx:context.Context,httpPath:string,body:any,accessTokeType:larkcore.AccessTokenType,...options:larkcore.RequestOptionFunc[]):Ref<larkcore.ApiResp>
			getAppAccessTokenBySelfBuiltApp(ctx:context.Context,req:Ref<larkcore.SelfBuiltAppAccessTokenReq>):Ref<larkcore.AppAccessTokenResp>
			getAppAccessTokenByMarketplaceApp(ctx:context.Context,req:Ref<larkcore.MarketplaceAppAccessTokenReq>):Ref<larkcore.AppAccessTokenResp>
			getTenantAccessTokenBySelfBuiltApp(ctx:context.Context,req:Ref<larkcore.SelfBuiltTenantAccessTokenReq>):Ref<larkcore.TenantAccessTokenResp>
			getTenantAccessTokenByMarketplaceApp(ctx:context.Context,req:Ref<larkcore.MarketplaceTenantAccessTokenReq>):Ref<larkcore.TenantAccessTokenResp>
			resendAppTicket(ctx:context.Context,req:Ref<larkcore.ResendAppTicketReq>):Ref<larkcore.ResendAppTicketResp>
	}
	export interface ClientOptionFunc extends Alias<(config:Ref<larkcore.Config>)=>void>{

	}
	export const FeishuBaseUrl:string
	export const LarkBaseUrl:string
	export function newClient(appId:string,appSecret:string,...options:ClientOptionFunc[]):Ref<Client>

	export function withAppType(appType:larkcore.AppType):ClientOptionFunc

	export function withEnableTokenCache(enableTokenCache:bool):ClientOptionFunc

	export function withHeaders(header:http.Header):ClientOptionFunc

	export function withHelpdeskCredential(helpdeskID:string,helpdeskToken:string):ClientOptionFunc

	export function withHttpClient(httpClient:larkcore.HttpClient):ClientOptionFunc

	export function withLogLevel(logLevel:larkcore.LogLevel):ClientOptionFunc

	export function withLogReqAtDebug(printReqRespLog:bool):ClientOptionFunc

	export function withLogger(logger:larkcore.Logger):ClientOptionFunc

	export function withMarketplaceApp():ClientOptionFunc

	export function withOpenBaseUrl(baseUrl:string):ClientOptionFunc

	export function withReqTimeout(reqTimeout:time.Duration):ClientOptionFunc

	export function withSerialization(serializable:larkcore.Serializable):ClientOptionFunc

	export function withTokenCache(cache:larkcore.Cache):ClientOptionFunc

}