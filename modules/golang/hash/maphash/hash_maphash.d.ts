// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/hash/maphash'{

	// @ts-ignore
	import type {Struct,byte,error,int,uint64,Ref} from 'go'
	export function bytes(seed:Seed,b:Uint8Array):uint64

	export interface Hash extends Struct<Hash>{

			writeByte(b:byte)/*error*/
			write(b:Uint8Array):int
			writeString(s:string):int
			seed():Seed
			setSeed(seed:Seed):void
			reset():void
			sum64():uint64
			sum(b:Uint8Array):Uint8Array
			size():int
			blockSize():int
	}
	export function makeSeed():Seed

	export interface Seed extends Struct<Seed>{

	}
	export function string(seed:Seed,s:string):uint64

	export function emptySeed():Seed
	export function emptyRefSeed():Ref<Seed>
	export function refOfSeed(x:Seed,v:Ref<Seed>)
	export function unRefSeed(v:Ref<Seed>):Seed
	export function emptyHash():Hash
	export function emptyRefHash():Ref<Hash>
	export function refOfHash(x:Hash,v:Ref<Hash>)
	export function unRefHash(v:Ref<Hash>):Hash
}