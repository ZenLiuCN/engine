// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/compress/zlib'{
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {error,Ref,Struct,int,GoError} from 'go'
	//9
	export const BestCompression:int
	//1
	export const BestSpeed:int
	//-1
	export const DefaultCompression:int
	export const ErrChecksum:GoError
	export const ErrDictionary:GoError
	export const ErrHeader:GoError
	//-2
	export const HuffmanOnly:int
	export function newReader(r:io.Reader):io.ReadCloser
	export function newReaderDict(r:io.Reader,dict:Uint8Array):[io.ReadCloser,error]
	export function newWriter(w:io.Writer):Ref<Writer>
	export function newWriterLevel(w:io.Writer,level:int):[Ref<Writer>,error]
	export function newWriterLevelDict(w:io.Writer,level:int,dict:Uint8Array):[Ref<Writer>,error]
	//0
	export const NoCompression:int
	export interface Resetter{
		reset(r:io.Reader,dict:Uint8Array):error
	}
	export interface Writer extends Struct<Writer>,io.WriteCloser,io.Closer{
		reset(w:io.Writer):void
		write(p:Uint8Array):[int,error]
		flush():error
		close():error
	}

export function emptyWriter():Writer
export function refWriter():Ref<Writer>
export function refOfWriter(x:Writer):Ref<Writer>}
