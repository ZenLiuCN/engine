// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/hash'{

	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {int,uint32,uint64,Ref} from 'go'
	export interface Hash extends io.Writer{

			blockSize():int
			reset():void
			size():int
			sum(b:Uint8Array):Uint8Array
	}
	export interface Hash32 extends Hash{

			sum32():uint32
	}
	export interface Hash64 extends Hash{

			sum64():uint64
	}
}