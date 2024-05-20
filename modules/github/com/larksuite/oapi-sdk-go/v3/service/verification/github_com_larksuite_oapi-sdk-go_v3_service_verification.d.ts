// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/service/verification'{

	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as larkverification from 'github.com/larksuite/oapi-sdk-go/v3/service/verification/v1'
	// @ts-ignore
	import type {Ref,Struct} from 'go'
	export function newService(config:Ref<larkcore.Config>):Ref<Service>

	export interface Service extends Struct<Service>{

			V1:Ref<larkverification.V1>
	}
}