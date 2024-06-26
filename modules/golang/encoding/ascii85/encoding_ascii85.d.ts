// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/encoding/ascii85'{

	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {int64,GoError,bool,int,error,Ref} from 'go'
	export interface CorruptInputError extends int64,GoError{

	error():string
	}
	export function decode(dst:Uint8Array,src:Uint8Array,flush:bool):[int,int]

	export function encode(dst:Uint8Array,src:Uint8Array):int

	export function maxEncodedLen(n:int):int

	export function newDecoder(r:io.Reader):io.Reader

	export function newEncoder(w:io.Writer):io.WriteCloser

}