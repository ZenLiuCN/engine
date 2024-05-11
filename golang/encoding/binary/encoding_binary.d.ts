// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/encoding/binary'{
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {uint16,uint32,uint64,int64,int,error} from 'go'
	export interface AppendByteOrder{
		appendUint16(v2:Uint8Array,v1:uint16):Uint8Array
		appendUint32(v2:Uint8Array,v1:uint32):Uint8Array
		appendUint64(v2:Uint8Array,v1:uint64):Uint8Array
		string():string
	}
	export function appendUvarint(buf:Uint8Array,x:uint64):Uint8Array
	export function appendVarint(buf:Uint8Array,x:int64):Uint8Array
	export const BigEndian:AppendByteOrder&ByteOrder
	export interface ByteOrder{
		putUint16(v2:Uint8Array,v1:uint16):void
		putUint32(v2:Uint8Array,v1:uint32):void
		putUint64(v2:Uint8Array,v1:uint64):void
		string():string
		uint16(v1:Uint8Array):uint16
		uint32(v1:Uint8Array):uint32
		uint64(v1:Uint8Array):uint64
	}
	export const LittleEndian:AppendByteOrder&ByteOrder
	//3
	export const MaxVarintLen16:int
	//5
	export const MaxVarintLen32:int
	//10
	export const MaxVarintLen64:int
	export const NativeEndian:AppendByteOrder&ByteOrder
	export function putUvarint(buf:Uint8Array,x:uint64):int
	export function putVarint(buf:Uint8Array,x:int64):int
	export function read(r:io.Reader,order:ByteOrder,data:any)/*error*/
	export function readUvarint(r:io.ByteReader):uint64
	export function readVarint(r:io.ByteReader):int64
	export function size(v:any):int
	export function uvarint(buf:Uint8Array):[uint64,int]
	export function varint(buf:Uint8Array):[int64,int]
	export function write(w:io.Writer,order:ByteOrder,data:any)/*error*/

}
