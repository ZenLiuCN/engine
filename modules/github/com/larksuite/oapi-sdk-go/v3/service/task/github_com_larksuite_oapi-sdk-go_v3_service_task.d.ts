// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/service/task'{

	// @ts-ignore
	import * as larktask from 'github.com/larksuite/oapi-sdk-go/v3/service/task/v1'
	// @ts-ignore
	import * as larktask from 'github.com/larksuite/oapi-sdk-go/v3/service/task/v2'
	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import type {Struct,Ref} from 'go'
	export function newService(config:Ref<larkcore.Config>):Ref<Service>

	export interface Service extends Struct<Service>{

			V1:Ref<larktask.V1>
			V2:Ref<larktask.V2>
	}
}