// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/container/list'{

	// @ts-ignore
	import type {Struct,Ref,int} from 'go'
	export interface Element extends Struct<Element>{

			value:any
			next():Ref<Element>
			prev():Ref<Element>
	}
	export interface List extends Struct<List>{

			init():Ref<List>
			len():int
			front():Ref<Element>
			back():Ref<Element>
			remove(e:Ref<Element>):any
			pushFront(v:any):Ref<Element>
			pushBack(v:any):Ref<Element>
			insertBefore(v:any,mark:Ref<Element>):Ref<Element>
			insertAfter(v:any,mark:Ref<Element>):Ref<Element>
			moveToFront(e:Ref<Element>):void
			moveToBack(e:Ref<Element>):void
			moveBefore(e:Ref<Element>,mark:Ref<Element>):void
			moveAfter(e:Ref<Element>,mark:Ref<Element>):void
			pushBackList(other:Ref<List>):void
			pushFrontList(other:Ref<List>):void
	}
	export function New():Ref<List>

	export function emptyElement():Element
	export function emptyRefElement():Ref<Element>
	export function refOfElement(x:Element,v:Ref<Element>)
	export function unRefElement(v:Ref<Element>):Element
	export function emptyList():List
	export function emptyRefList():Ref<List>
	export function refOfList(x:List,v:Ref<List>)
	export function unRefList(v:Ref<List>):List
}