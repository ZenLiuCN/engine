// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/encoding/gob'{
	// @ts-ignore
	import * as reflect from 'golang/reflect'
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {Struct,error,Ref} from 'go'
	export interface CommonType extends Struct<CommonType>{
		name:string
	}
	export interface Decoder extends Struct<Decoder>{
		decode(e:any)/*error*/
		decodeValue(v:reflect.Value)/*error*/
	}
	export interface Encoder extends Struct<Encoder>{
		encode(e:any)/*error*/
		encodeValue(value:reflect.Value)/*error*/
	}
	export interface GobDecoder{
		gobDecode(v1:Uint8Array)/*error*/
	}
	export interface GobEncoder{
		gobEncode():Uint8Array
	}
	export function newDecoder(r:io.Reader):Ref<Decoder>
	export function newEncoder(w:io.Writer):Ref<Encoder>
	export function register(value:any):void
	export function registerName(name:string,value:any):void

export function emptyCommonType():CommonType
export function refCommonType():Ref<CommonType>
export function refOfCommonType(x:CommonType):Ref<CommonType>
export function emptyDecoder():Decoder
export function refDecoder():Ref<Decoder>
export function refOfDecoder(x:Decoder):Ref<Decoder>
export function emptyEncoder():Encoder
export function refEncoder():Ref<Encoder>
export function refOfEncoder(x:Encoder):Ref<Encoder>
}
