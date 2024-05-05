declare module "golang/math/rand" {
	// @ts-ignore
	import type {uint32,Struct,int,Proto,float64,Ref,float32,uint64,int64,int32,error} from 'go'


	export interface Zipf extends Struct<Zipf>{
		uint64():uint64

	}
	export function newSource(seed:int64):Source
	export function int31():int32
	export function int():int
	export function intn(n:int):int
	export function perm(n:int):int[]
	export function shuffle(n:int,swap:(i,j:int)=>void):void

	export interface Source extends Proto<Source>{
		int63():int64

		seed(seed:int64):void

	}
	export function uint64():uint64
	export function int31n(n:int32):int32
	export function int63n(n:int64):int64
	export function normFloat64():float64

	export interface Source64 extends Proto<Source64>,Source{
		uint64():uint64

	}
	export function New(src:Source):Ref<Rand>
	export function float64():float64
	export function float32():float32
	export function expFloat64():float64
	export function newZipf(r:Ref<Rand>,s:float64,v:float64,imax:uint64):Ref<Zipf>
	export interface Rand extends Struct<Rand>{
		uint32():uint32
		int31():int32
		intn(n:int):int
		float32():float32
		shuffle(n:int,swap:(i,j:int)=>void):void
		read(p:Uint8Array):[int,error]
		normFloat64():float64
		int63():int64
		uint64():uint64
		seed(seed:int64):void
		float64():float64
		perm(n:int):int[]
		int():int
		int63n(n:int64):int64
		int31n(n:int32):int32
		expFloat64():float64

	}
	export function int63():int64
	export function uint32():uint32

}