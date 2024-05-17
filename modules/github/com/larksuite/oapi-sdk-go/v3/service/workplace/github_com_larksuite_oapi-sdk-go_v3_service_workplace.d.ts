// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/service/workplace'{

	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as larkworkplace from 'github.com/larksuite/oapi-sdk-go/v3/service/workplace/v1'
	// @ts-ignore
	import type {Ref,Struct} from 'go'
	export function newService(config:Ref<larkcore.Config>):Ref<Service>

	export interface Service extends Struct<Service>{

			V1:Ref<larkworkplace.V1>
	}
	export function emptyService():Service
	export function emptyRefService():Ref<Service>
	export function refOfService(x:Service,v:Ref<Service>)
	export function unRefService(v:Ref<Service>):Service
}