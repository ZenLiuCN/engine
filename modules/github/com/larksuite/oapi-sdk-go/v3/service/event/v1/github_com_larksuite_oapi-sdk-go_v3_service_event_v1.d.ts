// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/service/event/v1'{

	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import type {Ref,Struct,bool,error,int} from 'go'
	export interface DepartmentId extends Struct<DepartmentId>{

			departmentId:Ref<string>
			openDepartmentId:Ref<string>
	}
	export interface DepartmentIdBuilder extends Struct<DepartmentIdBuilder>{

			departmentId(departmentId:string):Ref<DepartmentIdBuilder>
			openDepartmentId(openDepartmentId:string):Ref<DepartmentIdBuilder>
			build():Ref<DepartmentId>
	}
	export interface ListOutboundIpIterator extends Struct<ListOutboundIpIterator>{

			next():[bool,string]
			nextPageToken():Ref<string>
	}
	export interface ListOutboundIpReq extends Struct<ListOutboundIpReq>{

			limit:int
	}
	export interface ListOutboundIpReqBuilder extends Struct<ListOutboundIpReqBuilder>{

			limit(limit:int):Ref<ListOutboundIpReqBuilder>
			pageSize(pageSize:int):Ref<ListOutboundIpReqBuilder>
			pageToken(pageToken:string):Ref<ListOutboundIpReqBuilder>
			build():Ref<ListOutboundIpReq>
	}
	export interface ListOutboundIpResp extends Struct<ListOutboundIpResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<ListOutboundIpRespData>
			success():bool
	}
	export interface ListOutboundIpRespData extends Struct<ListOutboundIpRespData>{

			ipList:string[]
			pageToken:Ref<string>
			hasMore:Ref<bool>
	}
	export function New(config:Ref<larkcore.Config>):Ref<V1>

	export function newDepartmentIdBuilder():Ref<DepartmentIdBuilder>

	export function newListOutboundIpReqBuilder():Ref<ListOutboundIpReqBuilder>

	export interface V1 extends Struct<V1>{

			outboundIp:Ref<{
			
				list(ctx:context.Context,req:Ref<ListOutboundIpReq>,...options:larkcore.RequestOptionFunc[]):Ref<ListOutboundIpResp>
				listByIterator(ctx:context.Context,req:Ref<ListOutboundIpReq>,...options:larkcore.RequestOptionFunc[]):Ref<ListOutboundIpIterator>
			}>
	}
	export function emptyListOutboundIpResp():ListOutboundIpResp
	export function emptyRefListOutboundIpResp():Ref<ListOutboundIpResp>
	export function refOfListOutboundIpResp(x:ListOutboundIpResp,v:Ref<ListOutboundIpResp>)
	export function unRefListOutboundIpResp(v:Ref<ListOutboundIpResp>):ListOutboundIpResp
	export function emptyListOutboundIpRespData():ListOutboundIpRespData
	export function emptyRefListOutboundIpRespData():Ref<ListOutboundIpRespData>
	export function refOfListOutboundIpRespData(x:ListOutboundIpRespData,v:Ref<ListOutboundIpRespData>)
	export function unRefListOutboundIpRespData(v:Ref<ListOutboundIpRespData>):ListOutboundIpRespData
	export function emptyV1():V1
	export function emptyRefV1():Ref<V1>
	export function refOfV1(x:V1,v:Ref<V1>)
	export function unRefV1(v:Ref<V1>):V1
	export function emptyDepartmentId():DepartmentId
	export function emptyRefDepartmentId():Ref<DepartmentId>
	export function refOfDepartmentId(x:DepartmentId,v:Ref<DepartmentId>)
	export function unRefDepartmentId(v:Ref<DepartmentId>):DepartmentId
	export function emptyListOutboundIpIterator():ListOutboundIpIterator
	export function emptyRefListOutboundIpIterator():Ref<ListOutboundIpIterator>
	export function refOfListOutboundIpIterator(x:ListOutboundIpIterator,v:Ref<ListOutboundIpIterator>)
	export function unRefListOutboundIpIterator(v:Ref<ListOutboundIpIterator>):ListOutboundIpIterator
	export function emptyListOutboundIpReq():ListOutboundIpReq
	export function emptyRefListOutboundIpReq():Ref<ListOutboundIpReq>
	export function refOfListOutboundIpReq(x:ListOutboundIpReq,v:Ref<ListOutboundIpReq>)
	export function unRefListOutboundIpReq(v:Ref<ListOutboundIpReq>):ListOutboundIpReq
}