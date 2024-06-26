// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/compress/lzw'{

	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {int,Struct,error,Ref} from 'go'
	export const LSB:Order
	export const MSB:Order
	export function newReader(r:io.Reader,order:Order,litWidth:int):io.ReadCloser

	export function newWriter(w:io.Writer,order:Order,litWidth:int):io.WriteCloser

	export interface Order extends int{

	}
	export interface Reader extends Struct<Reader>,io.ReadCloser,io.Closer{

			read(b:Uint8Array):int
			close():error
			reset(src:io.Reader,order:Order,litWidth:int):void
	}
	export interface Writer extends io.Closer,Struct<Writer>,io.WriteCloser{

			write(p:Uint8Array):int
			close():error
			reset(dst:io.Writer,order:Order,litWidth:int):void
	}
}