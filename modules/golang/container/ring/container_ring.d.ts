// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/container/ring'{

	// @ts-ignore
	import type {int,Ref,Struct} from 'go'
	export function New(n:int):Ref<Ring>

	export interface Ring extends Struct<Ring>{

			value:any
			next():Ref<Ring>
			prev():Ref<Ring>
			move(n:int):Ref<Ring>
			link(s:Ref<Ring>):Ref<Ring>
			unlink(n:int):Ref<Ring>
			len():int
			Do(v1:(v1:any)=>void):void
	}
	export function emptyRing():Ring
	export function emptyRefRing():Ref<Ring>
	export function refOfRing(x:Ring,v:Ref<Ring>)
	export function unRefRing(v:Ref<Ring>):Ring
}