// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/archive/tar'{
	// @ts-ignore
	import * as fs from 'golang/io/fs'
	// @ts-ignore
	import * as time from 'golang/time'
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {Ref,error,int,byte,int64,Struct,rune,GoError} from 'go'
	export const ErrFieldTooLong:GoError
	export const ErrHeader:GoError
	export const ErrInsecurePath:GoError
	export const ErrWriteAfterClose:GoError
	export const ErrWriteTooLong:GoError
	export function fileInfoHeader(fi:fs.FileInfo,link:string):Ref<Header>
	export interface Format extends int{
		string():string
	}
	export const FormatGNU:Format
	export const FormatPAX:Format
	export const FormatUSTAR:Format
	export const FormatUnknown:Format
	export interface Header extends Struct<Header>{
		typeflag:byte
		name:string
		linkname:string
		size:int64
		mode:int64
		uid:int
		gid:int
		uname:string
		gname:string
		modTime:time.Time
		accessTime:time.Time
		changeTime:time.Time
		devmajor:int64
		devminor:int64
		xattrs:Record<string,string>
		paxRecords:Record<string,string>
		format:Format
		fileInfo():fs.FileInfo
	}
	export function newReader(r:io.Reader):Ref<Reader>
	export function newWriter(w:io.Writer):Ref<Writer>
	export interface Reader extends Struct<Reader>,io.Reader{
		next():Ref<Header>
		read(b:Uint8Array):int
	}
	//52
	export const TypeBlock:rune
	//51
	export const TypeChar:rune
	//55
	export const TypeCont:rune
	//53
	export const TypeDir:rune
	//54
	export const TypeFifo:rune
	//75
	export const TypeGNULongLink:rune
	//76
	export const TypeGNULongName:rune
	//83
	export const TypeGNUSparse:rune
	//49
	export const TypeLink:rune
	//48
	export const TypeReg:rune
	//0
	export const TypeRegA:rune
	//50
	export const TypeSymlink:rune
	//103
	export const TypeXGlobalHeader:rune
	//120
	export const TypeXHeader:rune
	export interface Writer extends Struct<Writer>,io.WriteCloser,io.Closer{
		flush()/*error*/
		writeHeader(hdr:Ref<Header>)/*error*/
		write(b:Uint8Array):int
		close():error
	}

export function emptyHeader():Header
export function refHeader():Ref<Header>
export function refOfHeader(x:Header):Ref<Header>
export function emptyReader():Reader
export function refReader():Ref<Reader>
export function refOfReader(x:Reader):Ref<Reader>
export function emptyWriter():Writer
export function refWriter():Ref<Writer>
export function refOfWriter(x:Writer):Ref<Writer>
}
