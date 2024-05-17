// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/service/report/v1'{

	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import type {Ref,Struct,int,bool,error,Alias,Nothing} from 'go'
	export interface DepartmentId extends Struct<DepartmentId>{

			departmentId:Ref<string>
			openDepartmentId:Ref<string>
	}
	export interface DepartmentIdBuilder extends Struct<DepartmentIdBuilder>{

			departmentId(departmentId:string):Ref<DepartmentIdBuilder>
			openDepartmentId(openDepartmentId:string):Ref<DepartmentIdBuilder>
			build():Ref<DepartmentId>
	}
	export interface FormContent extends Struct<FormContent>{

			fieldId:Ref<string>
			fieldName:Ref<string>
			fieldValue:Ref<string>
	}
	export interface FormContentBuilder extends Struct<FormContentBuilder>{

			fieldId(fieldId:string):Ref<FormContentBuilder>
			fieldName(fieldName:string):Ref<FormContentBuilder>
			fieldValue(fieldValue:string):Ref<FormContentBuilder>
			build():Ref<FormContent>
	}
	export interface FormField extends Struct<FormField>{

			name:Ref<string>
			type:Ref<string>
	}
	export interface FormFieldBuilder extends Struct<FormFieldBuilder>{

			name(name:string):Ref<FormFieldBuilder>
			type(type_:string):Ref<FormFieldBuilder>
			build():Ref<FormField>
	}
	//0
	export const IncludeDeletedExclude:int
	//1
	export const IncludeDeletedInclude:int
	export function New(config:Ref<larkcore.Config>):Ref<V1>

	export function newDepartmentIdBuilder():Ref<DepartmentIdBuilder>

	export function newFormContentBuilder():Ref<FormContentBuilder>

	export function newFormFieldBuilder():Ref<FormFieldBuilder>

	export function newQueryRuleReqBuilder():Ref<QueryRuleReqBuilder>

	export function newQueryTaskPathReqBodyBuilder():Ref<QueryTaskPathReqBodyBuilder>

	export function newQueryTaskReqBodyBuilder():Ref<QueryTaskReqBodyBuilder>

	export function newQueryTaskReqBuilder():Ref<QueryTaskReqBuilder>

	export function newRemoveRuleViewPathReqBodyBuilder():Ref<RemoveRuleViewPathReqBodyBuilder>

	export function newRemoveRuleViewReqBodyBuilder():Ref<RemoveRuleViewReqBodyBuilder>

	export function newRemoveRuleViewReqBuilder():Ref<RemoveRuleViewReqBuilder>

	export function newRuleBuilder():Ref<RuleBuilder>

	export function newTaskBuilder():Ref<TaskBuilder>

	export interface QueryRuleReq extends Struct<QueryRuleReq>{

	}
	export interface QueryRuleReqBuilder extends Struct<QueryRuleReqBuilder>{

			ruleName(ruleName:string):Ref<QueryRuleReqBuilder>
			includeDeleted(includeDeleted:int):Ref<QueryRuleReqBuilder>
			userIdType(userIdType:string):Ref<QueryRuleReqBuilder>
			build():Ref<QueryRuleReq>
	}
	export interface QueryRuleResp extends Struct<QueryRuleResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<QueryRuleRespData>
			success():bool
	}
	export interface QueryRuleRespData extends Struct<QueryRuleRespData>{

			rules:Ref<Rule>[]
	}
	export interface QueryTaskPathReqBodyBuilder extends Struct<QueryTaskPathReqBodyBuilder>{

			commitStartTime(commitStartTime:int):Ref<QueryTaskPathReqBodyBuilder>
			commitEndTime(commitEndTime:int):Ref<QueryTaskPathReqBodyBuilder>
			ruleId(ruleId:string):Ref<QueryTaskPathReqBodyBuilder>
			userId(userId:string):Ref<QueryTaskPathReqBodyBuilder>
			pageToken(pageToken:string):Ref<QueryTaskPathReqBodyBuilder>
			pageSize(pageSize:int):Ref<QueryTaskPathReqBodyBuilder>
			build():Ref<QueryTaskReqBody>
	}
	export interface QueryTaskReq extends Struct<QueryTaskReq>{

			body:Ref<QueryTaskReqBody>
	}
	export interface QueryTaskReqBody extends Struct<QueryTaskReqBody>{

			commitStartTime:Ref<int>
			commitEndTime:Ref<int>
			ruleId:Ref<string>
			userId:Ref<string>
			pageToken:Ref<string>
			pageSize:Ref<int>
	}
	export interface QueryTaskReqBodyBuilder extends Struct<QueryTaskReqBodyBuilder>{

			commitStartTime(commitStartTime:int):Ref<QueryTaskReqBodyBuilder>
			commitEndTime(commitEndTime:int):Ref<QueryTaskReqBodyBuilder>
			ruleId(ruleId:string):Ref<QueryTaskReqBodyBuilder>
			userId(userId:string):Ref<QueryTaskReqBodyBuilder>
			pageToken(pageToken:string):Ref<QueryTaskReqBodyBuilder>
			pageSize(pageSize:int):Ref<QueryTaskReqBodyBuilder>
			build():Ref<QueryTaskReqBody>
	}
	export interface QueryTaskReqBuilder extends Struct<QueryTaskReqBuilder>{

			userIdType(userIdType:string):Ref<QueryTaskReqBuilder>
			body(body:Ref<QueryTaskReqBody>):Ref<QueryTaskReqBuilder>
			build():Ref<QueryTaskReq>
	}
	export interface QueryTaskResp extends Struct<QueryTaskResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<QueryTaskRespData>
			success():bool
	}
	export interface QueryTaskRespData extends Struct<QueryTaskRespData>{

			items:Ref<Task>[]
			hasMore:Ref<bool>
			pageToken:Ref<string>
	}
	export interface RemoveRuleViewPathReqBodyBuilder extends Struct<RemoveRuleViewPathReqBodyBuilder>{

			userIds(userIds:string[]):Ref<RemoveRuleViewPathReqBodyBuilder>
			build():Ref<RemoveRuleViewReqBody>
	}
	export interface RemoveRuleViewReq extends Struct<RemoveRuleViewReq>{

			body:Ref<RemoveRuleViewReqBody>
	}
	export interface RemoveRuleViewReqBody extends Struct<RemoveRuleViewReqBody>{

			userIds:string[]
	}
	export interface RemoveRuleViewReqBodyBuilder extends Struct<RemoveRuleViewReqBodyBuilder>{

			userIds(userIds:string[]):Ref<RemoveRuleViewReqBodyBuilder>
			build():Ref<RemoveRuleViewReqBody>
	}
	export interface RemoveRuleViewReqBuilder extends Struct<RemoveRuleViewReqBuilder>{

			ruleId(ruleId:string):Ref<RemoveRuleViewReqBuilder>
			userIdType(userIdType:string):Ref<RemoveRuleViewReqBuilder>
			body(body:Ref<RemoveRuleViewReqBody>):Ref<RemoveRuleViewReqBuilder>
			build():Ref<RemoveRuleViewReq>
	}
	export interface RemoveRuleViewResp extends Struct<RemoveRuleViewResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			success():bool
	}
	export interface Rule extends Struct<Rule>{

			ruleId:Ref<string>
			name:Ref<string>
			iconName:Ref<string>
			createdAt:Ref<int>
			creatorUserId:Ref<string>
			creatorUserName:Ref<string>
			ownerUserId:Ref<string>
			ownerUserName:Ref<string>
			formSchema:Ref<FormField>[]
			isDeleted:Ref<int>
			needReportUserIds:string[]
			needReportDepartmentIds:string[]
			needReportChatIds:string[]
			ccUserIds:string[]
			ccDepartmentIds:string[]
			toUserIds:string[]
			toChatIds:string[]
			toLeaders:int[]
			toDepartmentOwners:int[]
			managerUserIds:string[]
			ccChatIds:string[]
	}
	export interface RuleBuilder extends Struct<RuleBuilder>{

			ruleId(ruleId:string):Ref<RuleBuilder>
			name(name:string):Ref<RuleBuilder>
			iconName(iconName:string):Ref<RuleBuilder>
			createdAt(createdAt:int):Ref<RuleBuilder>
			creatorUserId(creatorUserId:string):Ref<RuleBuilder>
			creatorUserName(creatorUserName:string):Ref<RuleBuilder>
			ownerUserId(ownerUserId:string):Ref<RuleBuilder>
			ownerUserName(ownerUserName:string):Ref<RuleBuilder>
			formSchema(formSchema:Ref<FormField>[]):Ref<RuleBuilder>
			isDeleted(isDeleted:int):Ref<RuleBuilder>
			needReportUserIds(needReportUserIds:string[]):Ref<RuleBuilder>
			needReportDepartmentIds(needReportDepartmentIds:string[]):Ref<RuleBuilder>
			needReportChatIds(needReportChatIds:string[]):Ref<RuleBuilder>
			ccUserIds(ccUserIds:string[]):Ref<RuleBuilder>
			ccDepartmentIds(ccDepartmentIds:string[]):Ref<RuleBuilder>
			toUserIds(toUserIds:string[]):Ref<RuleBuilder>
			toChatIds(toChatIds:string[]):Ref<RuleBuilder>
			toLeaders(toLeaders:int[]):Ref<RuleBuilder>
			toDepartmentOwners(toDepartmentOwners:int[]):Ref<RuleBuilder>
			managerUserIds(managerUserIds:string[]):Ref<RuleBuilder>
			ccChatIds(ccChatIds:string[]):Ref<RuleBuilder>
			build():Ref<Rule>
	}
	export interface Task extends Struct<Task>{

			taskId:Ref<string>
			ruleName:Ref<string>
			fromUserId:Ref<string>
			fromUserName:Ref<string>
			departmentName:Ref<string>
			commitTime:Ref<int>
			formContents:Ref<FormContent>[]
			ruleId:Ref<string>
	}
	export interface TaskBuilder extends Struct<TaskBuilder>{

			taskId(taskId:string):Ref<TaskBuilder>
			ruleName(ruleName:string):Ref<TaskBuilder>
			fromUserId(fromUserId:string):Ref<TaskBuilder>
			fromUserName(fromUserName:string):Ref<TaskBuilder>
			departmentName(departmentName:string):Ref<TaskBuilder>
			commitTime(commitTime:int):Ref<TaskBuilder>
			formContents(formContents:Ref<FormContent>[]):Ref<TaskBuilder>
			ruleId(ruleId:string):Ref<TaskBuilder>
			build():Ref<Task>
	}
	//"open_id"
	export const UserIdTypeOpenId:string
	//"open_id"
	export const UserIdTypeQueryTaskOpenId:string
	//"union_id"
	export const UserIdTypeQueryTaskUnionId:string
	//"user_id"
	export const UserIdTypeQueryTaskUserId:string
	//"open_id"
	export const UserIdTypeRemoveRuleViewOpenId:string
	//"union_id"
	export const UserIdTypeRemoveRuleViewUnionId:string
	//"user_id"
	export const UserIdTypeRemoveRuleViewUserId:string
	//"union_id"
	export const UserIdTypeUnionId:string
	//"user_id"
	export const UserIdTypeUserId:string
	export interface V1 extends Struct<V1>{

			rule:Ref<{
			
				query(ctx:context.Context,req:Ref<QueryRuleReq>,...options:larkcore.RequestOptionFunc[]):Ref<QueryRuleResp>
			}>
			ruleView:Ref<{
			
				remove(ctx:context.Context,req:Ref<RemoveRuleViewReq>,...options:larkcore.RequestOptionFunc[]):Ref<RemoveRuleViewResp>
			}>
			task:Ref<{
			
				query(ctx:context.Context,req:Ref<QueryTaskReq>,...options:larkcore.RequestOptionFunc[]):Ref<QueryTaskResp>
			}>
	}
	export interface View extends Alias<Nothing>{

	}
	export function emptyQueryRuleReq():QueryRuleReq
	export function emptyRefQueryRuleReq():Ref<QueryRuleReq>
	export function refOfQueryRuleReq(x:QueryRuleReq,v:Ref<QueryRuleReq>)
	export function unRefQueryRuleReq(v:Ref<QueryRuleReq>):QueryRuleReq
	export function emptyRemoveRuleViewReqBody():RemoveRuleViewReqBody
	export function emptyRefRemoveRuleViewReqBody():Ref<RemoveRuleViewReqBody>
	export function refOfRemoveRuleViewReqBody(x:RemoveRuleViewReqBody,v:Ref<RemoveRuleViewReqBody>)
	export function unRefRemoveRuleViewReqBody(v:Ref<RemoveRuleViewReqBody>):RemoveRuleViewReqBody
	export function emptyQueryTaskReq():QueryTaskReq
	export function emptyRefQueryTaskReq():Ref<QueryTaskReq>
	export function refOfQueryTaskReq(x:QueryTaskReq,v:Ref<QueryTaskReq>)
	export function unRefQueryTaskReq(v:Ref<QueryTaskReq>):QueryTaskReq
	export function emptyQueryTaskResp():QueryTaskResp
	export function emptyRefQueryTaskResp():Ref<QueryTaskResp>
	export function refOfQueryTaskResp(x:QueryTaskResp,v:Ref<QueryTaskResp>)
	export function unRefQueryTaskResp(v:Ref<QueryTaskResp>):QueryTaskResp
	export function emptyRule():Rule
	export function emptyRefRule():Ref<Rule>
	export function refOfRule(x:Rule,v:Ref<Rule>)
	export function unRefRule(v:Ref<Rule>):Rule
	export function emptyV1():V1
	export function emptyRefV1():Ref<V1>
	export function refOfV1(x:V1,v:Ref<V1>)
	export function unRefV1(v:Ref<V1>):V1
	export function emptyQueryTaskReqBodyBuilder():QueryTaskReqBodyBuilder
	export function emptyRefQueryTaskReqBodyBuilder():Ref<QueryTaskReqBodyBuilder>
	export function refOfQueryTaskReqBodyBuilder(x:QueryTaskReqBodyBuilder,v:Ref<QueryTaskReqBodyBuilder>)
	export function unRefQueryTaskReqBodyBuilder(v:Ref<QueryTaskReqBodyBuilder>):QueryTaskReqBodyBuilder
	export function emptyQueryTaskReqBuilder():QueryTaskReqBuilder
	export function emptyRefQueryTaskReqBuilder():Ref<QueryTaskReqBuilder>
	export function refOfQueryTaskReqBuilder(x:QueryTaskReqBuilder,v:Ref<QueryTaskReqBuilder>)
	export function unRefQueryTaskReqBuilder(v:Ref<QueryTaskReqBuilder>):QueryTaskReqBuilder
	export function emptyTask():Task
	export function emptyRefTask():Ref<Task>
	export function refOfTask(x:Task,v:Ref<Task>)
	export function unRefTask(v:Ref<Task>):Task
	export function emptyTaskBuilder():TaskBuilder
	export function emptyRefTaskBuilder():Ref<TaskBuilder>
	export function refOfTaskBuilder(x:TaskBuilder,v:Ref<TaskBuilder>)
	export function unRefTaskBuilder(v:Ref<TaskBuilder>):TaskBuilder
	export function emptyFormContent():FormContent
	export function emptyRefFormContent():Ref<FormContent>
	export function refOfFormContent(x:FormContent,v:Ref<FormContent>)
	export function unRefFormContent(v:Ref<FormContent>):FormContent
	export function emptyQueryRuleReqBuilder():QueryRuleReqBuilder
	export function emptyRefQueryRuleReqBuilder():Ref<QueryRuleReqBuilder>
	export function refOfQueryRuleReqBuilder(x:QueryRuleReqBuilder,v:Ref<QueryRuleReqBuilder>)
	export function unRefQueryRuleReqBuilder(v:Ref<QueryRuleReqBuilder>):QueryRuleReqBuilder
	export function emptyQueryRuleRespData():QueryRuleRespData
	export function emptyRefQueryRuleRespData():Ref<QueryRuleRespData>
	export function refOfQueryRuleRespData(x:QueryRuleRespData,v:Ref<QueryRuleRespData>)
	export function unRefQueryRuleRespData(v:Ref<QueryRuleRespData>):QueryRuleRespData
	export function emptyRemoveRuleViewPathReqBodyBuilder():RemoveRuleViewPathReqBodyBuilder
	export function emptyRefRemoveRuleViewPathReqBodyBuilder():Ref<RemoveRuleViewPathReqBodyBuilder>
	export function refOfRemoveRuleViewPathReqBodyBuilder(x:RemoveRuleViewPathReqBodyBuilder,v:Ref<RemoveRuleViewPathReqBodyBuilder>)
	export function unRefRemoveRuleViewPathReqBodyBuilder(v:Ref<RemoveRuleViewPathReqBodyBuilder>):RemoveRuleViewPathReqBodyBuilder
	export function emptyRuleBuilder():RuleBuilder
	export function emptyRefRuleBuilder():Ref<RuleBuilder>
	export function refOfRuleBuilder(x:RuleBuilder,v:Ref<RuleBuilder>)
	export function unRefRuleBuilder(v:Ref<RuleBuilder>):RuleBuilder
	export function emptyDepartmentId():DepartmentId
	export function emptyRefDepartmentId():Ref<DepartmentId>
	export function refOfDepartmentId(x:DepartmentId,v:Ref<DepartmentId>)
	export function unRefDepartmentId(v:Ref<DepartmentId>):DepartmentId
	export function emptyFormContentBuilder():FormContentBuilder
	export function emptyRefFormContentBuilder():Ref<FormContentBuilder>
	export function refOfFormContentBuilder(x:FormContentBuilder,v:Ref<FormContentBuilder>)
	export function unRefFormContentBuilder(v:Ref<FormContentBuilder>):FormContentBuilder
	export function emptyFormField():FormField
	export function emptyRefFormField():Ref<FormField>
	export function refOfFormField(x:FormField,v:Ref<FormField>)
	export function unRefFormField(v:Ref<FormField>):FormField
	export function emptyQueryRuleResp():QueryRuleResp
	export function emptyRefQueryRuleResp():Ref<QueryRuleResp>
	export function refOfQueryRuleResp(x:QueryRuleResp,v:Ref<QueryRuleResp>)
	export function unRefQueryRuleResp(v:Ref<QueryRuleResp>):QueryRuleResp
	export function emptyQueryTaskPathReqBodyBuilder():QueryTaskPathReqBodyBuilder
	export function emptyRefQueryTaskPathReqBodyBuilder():Ref<QueryTaskPathReqBodyBuilder>
	export function refOfQueryTaskPathReqBodyBuilder(x:QueryTaskPathReqBodyBuilder,v:Ref<QueryTaskPathReqBodyBuilder>)
	export function unRefQueryTaskPathReqBodyBuilder(v:Ref<QueryTaskPathReqBodyBuilder>):QueryTaskPathReqBodyBuilder
	export function emptyQueryTaskRespData():QueryTaskRespData
	export function emptyRefQueryTaskRespData():Ref<QueryTaskRespData>
	export function refOfQueryTaskRespData(x:QueryTaskRespData,v:Ref<QueryTaskRespData>)
	export function unRefQueryTaskRespData(v:Ref<QueryTaskRespData>):QueryTaskRespData
	export function emptyRemoveRuleViewReqBodyBuilder():RemoveRuleViewReqBodyBuilder
	export function emptyRefRemoveRuleViewReqBodyBuilder():Ref<RemoveRuleViewReqBodyBuilder>
	export function refOfRemoveRuleViewReqBodyBuilder(x:RemoveRuleViewReqBodyBuilder,v:Ref<RemoveRuleViewReqBodyBuilder>)
	export function unRefRemoveRuleViewReqBodyBuilder(v:Ref<RemoveRuleViewReqBodyBuilder>):RemoveRuleViewReqBodyBuilder
	export function emptyRemoveRuleViewResp():RemoveRuleViewResp
	export function emptyRefRemoveRuleViewResp():Ref<RemoveRuleViewResp>
	export function refOfRemoveRuleViewResp(x:RemoveRuleViewResp,v:Ref<RemoveRuleViewResp>)
	export function unRefRemoveRuleViewResp(v:Ref<RemoveRuleViewResp>):RemoveRuleViewResp
	export function emptyDepartmentIdBuilder():DepartmentIdBuilder
	export function emptyRefDepartmentIdBuilder():Ref<DepartmentIdBuilder>
	export function refOfDepartmentIdBuilder(x:DepartmentIdBuilder,v:Ref<DepartmentIdBuilder>)
	export function unRefDepartmentIdBuilder(v:Ref<DepartmentIdBuilder>):DepartmentIdBuilder
	export function emptyFormFieldBuilder():FormFieldBuilder
	export function emptyRefFormFieldBuilder():Ref<FormFieldBuilder>
	export function refOfFormFieldBuilder(x:FormFieldBuilder,v:Ref<FormFieldBuilder>)
	export function unRefFormFieldBuilder(v:Ref<FormFieldBuilder>):FormFieldBuilder
	export function emptyQueryTaskReqBody():QueryTaskReqBody
	export function emptyRefQueryTaskReqBody():Ref<QueryTaskReqBody>
	export function refOfQueryTaskReqBody(x:QueryTaskReqBody,v:Ref<QueryTaskReqBody>)
	export function unRefQueryTaskReqBody(v:Ref<QueryTaskReqBody>):QueryTaskReqBody
	export function emptyRemoveRuleViewReq():RemoveRuleViewReq
	export function emptyRefRemoveRuleViewReq():Ref<RemoveRuleViewReq>
	export function refOfRemoveRuleViewReq(x:RemoveRuleViewReq,v:Ref<RemoveRuleViewReq>)
	export function unRefRemoveRuleViewReq(v:Ref<RemoveRuleViewReq>):RemoveRuleViewReq
	export function emptyRemoveRuleViewReqBuilder():RemoveRuleViewReqBuilder
	export function emptyRefRemoveRuleViewReqBuilder():Ref<RemoveRuleViewReqBuilder>
	export function refOfRemoveRuleViewReqBuilder(x:RemoveRuleViewReqBuilder,v:Ref<RemoveRuleViewReqBuilder>)
	export function unRefRemoveRuleViewReqBuilder(v:Ref<RemoveRuleViewReqBuilder>):RemoveRuleViewReqBuilder
}