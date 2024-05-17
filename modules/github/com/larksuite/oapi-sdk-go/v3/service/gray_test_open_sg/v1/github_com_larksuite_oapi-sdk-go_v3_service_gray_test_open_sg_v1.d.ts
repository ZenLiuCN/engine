// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/service/gray_test_open_sg/v1'{

	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import type {Ref,Struct,bool,error,int} from 'go'
	export interface CreateMotoReq extends Struct<CreateMotoReq>{

			level:Ref<Level>
	}
	export interface CreateMotoReqBuilder extends Struct<CreateMotoReqBuilder>{

			departmentIdType(departmentIdType:string):Ref<CreateMotoReqBuilder>
			userIdType(userIdType:string):Ref<CreateMotoReqBuilder>
			level(level:Ref<Level>):Ref<CreateMotoReqBuilder>
			build():Ref<CreateMotoReq>
	}
	export interface CreateMotoResp extends Struct<CreateMotoResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<CreateMotoRespData>
			success():bool
	}
	export interface CreateMotoRespData extends Struct<CreateMotoRespData>{

			moto:Ref<Moto>
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
	export interface GetMotoReq extends Struct<GetMotoReq>{

	}
	export interface GetMotoReqBuilder extends Struct<GetMotoReqBuilder>{

			motoId(motoId:string):Ref<GetMotoReqBuilder>
			bodyLevel(bodyLevel:string):Ref<GetMotoReqBuilder>
			build():Ref<GetMotoReq>
	}
	export interface GetMotoResp extends Struct<GetMotoResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<GetMotoRespData>
			success():bool
	}
	export interface GetMotoRespData extends Struct<GetMotoRespData>{

			moto:Ref<Moto>
	}
	export interface Level extends Struct<Level>{

			level:Ref<string>
			body:Ref<string>
			type:Ref<string>
	}
	export interface LevelBuilder extends Struct<LevelBuilder>{

			level(level:string):Ref<LevelBuilder>
			body(body:string):Ref<LevelBuilder>
			type(type_:string):Ref<LevelBuilder>
			build():Ref<Level>
	}
	export interface ListMotoIterator extends Struct<ListMotoIterator>{

			next():[bool,string]
			nextPageToken():Ref<string>
	}
	export interface ListMotoReq extends Struct<ListMotoReq>{

			limit:int
	}
	export interface ListMotoReqBuilder extends Struct<ListMotoReqBuilder>{

			limit(limit:int):Ref<ListMotoReqBuilder>
			pageSize(pageSize:int):Ref<ListMotoReqBuilder>
			pageToken(pageToken:string):Ref<ListMotoReqBuilder>
			level(level:int):Ref<ListMotoReqBuilder>
			build():Ref<ListMotoReq>
	}
	export interface ListMotoResp extends Struct<ListMotoResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<ListMotoRespData>
			success():bool
	}
	export interface ListMotoRespData extends Struct<ListMotoRespData>{

			items:string[]
			pageToken:Ref<string>
			hasMore:Ref<bool>
	}
	export interface Moto extends Struct<Moto>{

			motoId:Ref<string>
			id:Ref<string>
			userName:Ref<string>
			type:Ref<string>
	}
	export interface MotoBuilder extends Struct<MotoBuilder>{

			motoId(motoId:string):Ref<MotoBuilder>
			id(id:string):Ref<MotoBuilder>
			userName(userName:string):Ref<MotoBuilder>
			type(type_:string):Ref<MotoBuilder>
			build():Ref<Moto>
	}
	export function New(config:Ref<larkcore.Config>):Ref<V1>

	export function newCreateMotoReqBuilder():Ref<CreateMotoReqBuilder>

	export function newDepartmentIdBuilder():Ref<DepartmentIdBuilder>

	export function newGetMotoReqBuilder():Ref<GetMotoReqBuilder>

	export function newLevelBuilder():Ref<LevelBuilder>

	export function newListMotoReqBuilder():Ref<ListMotoReqBuilder>

	export function newMotoBuilder():Ref<MotoBuilder>

	export interface V1 extends Struct<V1>{

			moto:Ref<{
			
				create(ctx:context.Context,req:Ref<CreateMotoReq>,...options:larkcore.RequestOptionFunc[]):Ref<CreateMotoResp>
				get(ctx:context.Context,req:Ref<GetMotoReq>,...options:larkcore.RequestOptionFunc[]):Ref<GetMotoResp>
				list(ctx:context.Context,req:Ref<ListMotoReq>,...options:larkcore.RequestOptionFunc[]):Ref<ListMotoResp>
				listByIterator(ctx:context.Context,req:Ref<ListMotoReq>,...options:larkcore.RequestOptionFunc[]):Ref<ListMotoIterator>
			}>
	}
	export function emptyCreateMotoRespData():CreateMotoRespData
	export function emptyRefCreateMotoRespData():Ref<CreateMotoRespData>
	export function refOfCreateMotoRespData(x:CreateMotoRespData,v:Ref<CreateMotoRespData>)
	export function unRefCreateMotoRespData(v:Ref<CreateMotoRespData>):CreateMotoRespData
	export function emptyGetMotoReq():GetMotoReq
	export function emptyRefGetMotoReq():Ref<GetMotoReq>
	export function refOfGetMotoReq(x:GetMotoReq,v:Ref<GetMotoReq>)
	export function unRefGetMotoReq(v:Ref<GetMotoReq>):GetMotoReq
	export function emptyListMotoResp():ListMotoResp
	export function emptyRefListMotoResp():Ref<ListMotoResp>
	export function refOfListMotoResp(x:ListMotoResp,v:Ref<ListMotoResp>)
	export function unRefListMotoResp(v:Ref<ListMotoResp>):ListMotoResp
	export function emptyCreateMotoReq():CreateMotoReq
	export function emptyRefCreateMotoReq():Ref<CreateMotoReq>
	export function refOfCreateMotoReq(x:CreateMotoReq,v:Ref<CreateMotoReq>)
	export function unRefCreateMotoReq(v:Ref<CreateMotoReq>):CreateMotoReq
	export function emptyLevel():Level
	export function emptyRefLevel():Ref<Level>
	export function refOfLevel(x:Level,v:Ref<Level>)
	export function unRefLevel(v:Ref<Level>):Level
	export function emptyLevelBuilder():LevelBuilder
	export function emptyRefLevelBuilder():Ref<LevelBuilder>
	export function refOfLevelBuilder(x:LevelBuilder,v:Ref<LevelBuilder>)
	export function unRefLevelBuilder(v:Ref<LevelBuilder>):LevelBuilder
	export function emptyMoto():Moto
	export function emptyRefMoto():Ref<Moto>
	export function refOfMoto(x:Moto,v:Ref<Moto>)
	export function unRefMoto(v:Ref<Moto>):Moto
	export function emptyV1():V1
	export function emptyRefV1():Ref<V1>
	export function refOfV1(x:V1,v:Ref<V1>)
	export function unRefV1(v:Ref<V1>):V1
	export function emptyDepartmentIdBuilder():DepartmentIdBuilder
	export function emptyRefDepartmentIdBuilder():Ref<DepartmentIdBuilder>
	export function refOfDepartmentIdBuilder(x:DepartmentIdBuilder,v:Ref<DepartmentIdBuilder>)
	export function unRefDepartmentIdBuilder(v:Ref<DepartmentIdBuilder>):DepartmentIdBuilder
	export function emptyDepartmentId():DepartmentId
	export function emptyRefDepartmentId():Ref<DepartmentId>
	export function refOfDepartmentId(x:DepartmentId,v:Ref<DepartmentId>)
	export function unRefDepartmentId(v:Ref<DepartmentId>):DepartmentId
	export function emptyListMotoReq():ListMotoReq
	export function emptyRefListMotoReq():Ref<ListMotoReq>
	export function refOfListMotoReq(x:ListMotoReq,v:Ref<ListMotoReq>)
	export function unRefListMotoReq(v:Ref<ListMotoReq>):ListMotoReq
	export function emptyCreateMotoReqBuilder():CreateMotoReqBuilder
	export function emptyRefCreateMotoReqBuilder():Ref<CreateMotoReqBuilder>
	export function refOfCreateMotoReqBuilder(x:CreateMotoReqBuilder,v:Ref<CreateMotoReqBuilder>)
	export function unRefCreateMotoReqBuilder(v:Ref<CreateMotoReqBuilder>):CreateMotoReqBuilder
	export function emptyGetMotoReqBuilder():GetMotoReqBuilder
	export function emptyRefGetMotoReqBuilder():Ref<GetMotoReqBuilder>
	export function refOfGetMotoReqBuilder(x:GetMotoReqBuilder,v:Ref<GetMotoReqBuilder>)
	export function unRefGetMotoReqBuilder(v:Ref<GetMotoReqBuilder>):GetMotoReqBuilder
	export function emptyGetMotoResp():GetMotoResp
	export function emptyRefGetMotoResp():Ref<GetMotoResp>
	export function refOfGetMotoResp(x:GetMotoResp,v:Ref<GetMotoResp>)
	export function unRefGetMotoResp(v:Ref<GetMotoResp>):GetMotoResp
	export function emptyGetMotoRespData():GetMotoRespData
	export function emptyRefGetMotoRespData():Ref<GetMotoRespData>
	export function refOfGetMotoRespData(x:GetMotoRespData,v:Ref<GetMotoRespData>)
	export function unRefGetMotoRespData(v:Ref<GetMotoRespData>):GetMotoRespData
	export function emptyListMotoIterator():ListMotoIterator
	export function emptyRefListMotoIterator():Ref<ListMotoIterator>
	export function refOfListMotoIterator(x:ListMotoIterator,v:Ref<ListMotoIterator>)
	export function unRefListMotoIterator(v:Ref<ListMotoIterator>):ListMotoIterator
	export function emptyListMotoReqBuilder():ListMotoReqBuilder
	export function emptyRefListMotoReqBuilder():Ref<ListMotoReqBuilder>
	export function refOfListMotoReqBuilder(x:ListMotoReqBuilder,v:Ref<ListMotoReqBuilder>)
	export function unRefListMotoReqBuilder(v:Ref<ListMotoReqBuilder>):ListMotoReqBuilder
	export function emptyListMotoRespData():ListMotoRespData
	export function emptyRefListMotoRespData():Ref<ListMotoRespData>
	export function refOfListMotoRespData(x:ListMotoRespData,v:Ref<ListMotoRespData>)
	export function unRefListMotoRespData(v:Ref<ListMotoRespData>):ListMotoRespData
	export function emptyMotoBuilder():MotoBuilder
	export function emptyRefMotoBuilder():Ref<MotoBuilder>
	export function refOfMotoBuilder(x:MotoBuilder,v:Ref<MotoBuilder>)
	export function unRefMotoBuilder(v:Ref<MotoBuilder>):MotoBuilder
	export function emptyCreateMotoResp():CreateMotoResp
	export function emptyRefCreateMotoResp():Ref<CreateMotoResp>
	export function refOfCreateMotoResp(x:CreateMotoResp,v:Ref<CreateMotoResp>)
	export function unRefCreateMotoResp(v:Ref<CreateMotoResp>):CreateMotoResp
}