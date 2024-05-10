// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/compress/lzw'{
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {int,Struct,error} from 'go'
	export const LSB:Order
	export const MSB:Order
	export function newReader(r:io.Reader,order:Order,litWidth:int):io.ReadCloser
	export function newWriter(w:io.Writer,order:Order,litWidth:int):io.WriteCloser
	export interface Order extends int{
	}
	export interface Reader extends Struct<Reader>,io.ReadCloser,io.Closer{
		read(b:Uint8Array):[int,error]
		close():error
		reset(src:io.Reader,order:Order,litWidth:int):void
	}
	export interface Writer extends Struct<Writer>,io.WriteCloser,io.Closer{
		write(p:Uint8Array):[int,error]
		close():error
		reset(dst:io.Writer,order:Order,litWidth:int):void
	}

export function emptyWriter():Writer
export function refWriter():Ref<Writer>
export function refOfWriter(x:Writer):Ref<Writer>
export function emptyReader():Reader
export function refReader():Ref<Reader>
export function refOfReader(x:Reader):Ref<Reader>}
