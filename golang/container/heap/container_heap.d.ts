// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/container/heap'{
	// @ts-ignore
	import * as sort from 'golang/sort'
	// @ts-ignore
	import type {int} from 'go'
	export function fix(h:Interface,i:int):void
	export function init(h:Interface):void
	export interface Interface extends sort.Interface{
		pop():any
		push(x:any):void
	}
	export function pop(h:Interface):any
	export function push(h:Interface,x:any):void
	export function remove(h:Interface,i:int):any

}
