// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/math/rand'{
	// @ts-ignore
	import type {int64,Ref,uint64,uint32,error,float64,float32,int,int32,Struct} from 'go'
	export function expFloat64():float64
	export function float32():float32
	export function float64():float64
	export function int():int
	export function int31():int32
	export function int31n(n:int32):int32
	export function int63():int64
	export function int63n(n:int64):int64
	export function intn(n:int):int
	export function New(src:Source):Ref<Rand>
	export function newSource(seed:int64):Source
	export function newZipf(r:Ref<Rand>,s:float64,v:float64,imax:uint64):Ref<Zipf>
	export function normFloat64():float64
	export function perm(n:int):int[]
	export interface Rand extends Source64,Struct<Rand>{
		expFloat64():float64
		normFloat64():float64
		seed(seed:int64):void
		int63():int64
		uint32():uint32
		uint64():uint64
		int31():int32
		int():int
		int63n(n:int64):int64
		int31n(n:int32):int32
		intn(n:int):int
		float64():float64
		float32():float32
		perm(n:int):int[]
		shuffle(n:int,swap:(i:int,j:int)=>void):void
		read(p:Uint8Array):int
	}
	export function read(p:Uint8Array):int
	export function seed(seed:int64):void
	export function shuffle(n:int,swap:(i:int,j:int)=>void):void
	export interface Source{
		int63():int64
		seed(seed:int64):void
	}
	export interface Source64 extends Source{
		uint64():uint64
	}
	export function uint32():uint32
	export function uint64():uint64
	export interface Zipf extends Struct<Zipf>{
		uint64():uint64
	}

export function emptyRand():Rand
export function refRand():Ref<Rand>
export function refOfRand(x:Rand):Ref<Rand>
export function emptyZipf():Zipf
export function refZipf():Ref<Zipf>
export function refOfZipf(x:Zipf):Ref<Zipf>
}
