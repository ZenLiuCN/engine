declare module "golang/math/rand" {
	// @ts-ignore
	import type {int,error,uint64,Proto,float32,uint32,float64,int32,Struct,int64,Ref} from 'go'


	export function int():int
	export function float64():float64
	export function perm(n:int):int[]
	export function read(p:Uint8Array):[int,error]

	export interface Source64 extends Proto<Source64>,Source{
		uint64():uint64

	}
	export function uint64():uint64
	export function int31():int32
	export function shuffle(n:int,swap:(i,j:int)=>void):void
	export interface Zipf extends Struct<Zipf>{
		uint64():uint64

	}

	export interface Source extends Proto<Source>{
		int63():int64

		seed(seed:int64):void

	}
	export function seed(seed:int64):void
	export function int63():int64
	export function int63n(n:int64):int64
	export function int31n(n:int32):int32
	export function intn(n:int):int
	export interface Rand extends Struct<Rand>{
		uint32():uint32
		int():int
		int31n(n:int32):int32
		perm(n:int):int[]
		read(p:Uint8Array):[int,error]
		normFloat64():float64
		float32():float32
		shuffle(n:int,swap:(i,j:int)=>void):void
		expFloat64():float64
		seed(seed:int64):void
		int31():int32
		int63n(n:int64):int64
		int63():int64
		uint64():uint64
		intn(n:int):int
		float64():float64

	}
	export function New(src:Source):Ref<Rand>
	export function float32():float32
	export function normFloat64():float64
	export function expFloat64():float64
	export function newZipf(r:Ref<Rand>,s:float64,v:float64,imax:uint64):Ref<Zipf>
	export function newSource(seed:int64):Source
	export function uint32():uint32

}