// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/archive/zip'{

	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as time from 'golang/time'
	// @ts-ignore
	import * as fs from 'golang/io/fs'
	// @ts-ignore
	import type {uint16,GoError,Ref,error,Alias,bool,uint32,uint64,Struct,int64} from 'go'
	export interface Compressor extends Alias<(w:io.Writer)=>[io.WriteCloser,error]>{

	}
	export interface Decompressor extends Alias<(r:io.Reader)=>io.ReadCloser>{

	}
	//8
	export const Deflate:uint16
	export const ErrAlgorithm:GoError
	export const ErrChecksum:GoError
	export const ErrFormat:GoError
	export const ErrInsecurePath:GoError
	export interface File extends Struct<File>{

			fileHeader:FileHeader
			dataOffset():int64
			open():io.ReadCloser
			openRaw():io.Reader
	}
	export interface FileHeader extends Struct<FileHeader>{

			name:string
			comment:string
			nonUTF8:bool
			creatorVersion:uint16
			readerVersion:uint16
			flags:uint16
			method:uint16
			modified:time.Time
			modifiedTime:uint16
			modifiedDate:uint16
			crC32:uint32
			compressedSize:uint32
			uncompressedSize:uint32
			compressedSize64:uint64
			uncompressedSize64:uint64
			extra:Uint8Array
			externalAttrs:uint32
			fileInfo():fs.FileInfo
			modTime():time.Time
			setModTime(t:time.Time):void
			mode():fs.FileMode
			setMode(mode:fs.FileMode):void
	}
	export function fileInfoHeader(fi:fs.FileInfo):Ref<FileHeader>

	export function newReader(r:io.ReaderAt,size:int64):Ref<Reader>

	export function newWriter(w:io.Writer):Ref<Writer>

	export function openReader(name:string):Ref<ReadCloser>

	export interface ReadCloser extends Struct<ReadCloser>,fs.FS,io.Closer{

			reader:Reader
			close():error
	}
	export interface Reader extends Struct<Reader>,fs.FS{

			file:Ref<File>[]
			comment:string
			registerDecompressor(method:uint16,dcomp:Decompressor):void
			open(name:string):fs.File
	}
	export function registerCompressor(method:uint16,comp:Compressor):void

	export function registerDecompressor(method:uint16,dcomp:Decompressor):void

	//0
	export const Store:uint16
	export interface Writer extends Struct<Writer>,io.Closer{

			setOffset(n:int64):void
			flush()/*error*/
			setComment(comment:string)/*error*/
			close():error
			create(name:string):io.Writer
			createHeader(fh:Ref<FileHeader>):io.Writer
			createRaw(fh:Ref<FileHeader>):io.Writer
			copy(f:Ref<File>)/*error*/
			registerCompressor(method:uint16,comp:Compressor):void
	}
	export function emptyFile():File
	export function emptyRefFile():Ref<File>
	export function refOfFile(x:File,v:Ref<File>)
	export function unRefFile(v:Ref<File>):File
	export function emptyFileHeader():FileHeader
	export function emptyRefFileHeader():Ref<FileHeader>
	export function refOfFileHeader(x:FileHeader,v:Ref<FileHeader>)
	export function unRefFileHeader(v:Ref<FileHeader>):FileHeader
	export function emptyReadCloser():ReadCloser
	export function emptyRefReadCloser():Ref<ReadCloser>
	export function refOfReadCloser(x:ReadCloser,v:Ref<ReadCloser>)
	export function unRefReadCloser(v:Ref<ReadCloser>):ReadCloser
}