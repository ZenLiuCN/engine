// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/encoding/csv'{
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {rune,bool,int64,GoError,Ref,int,Struct,error} from 'go'
	export const ErrBareQuote:GoError
	export const ErrFieldCount:GoError
	export const ErrQuote:GoError
	export const ErrTrailingComma:GoError
	export function newReader(r:io.Reader):Ref<Reader>
	export function newWriter(w:io.Writer):Ref<Writer>
	export interface ParseError extends Struct<ParseError>,Error,GoError{
		startLine:int
		line:int
		column:int
		err:GoError
		error():string
		unwrap():error
	}
	export interface Reader extends Struct<Reader>{
		comma:rune
		comment:rune
		fieldsPerRecord:int
		lazyQuotes:bool
		trimLeadingSpace:bool
		reuseRecord:bool
		trailingComma:bool
		read():[string[],error]
		fieldPos(field:int):[int,int]
		inputOffset():int64
		readAll():[Array<string[]>,error]
	}
	export interface Writer extends Struct<Writer>{
		comma:rune
		useCRLF:bool
		write(record:string[]):error
		flush():void
		error():error
		writeAll(records:Array<string[]>):error
	}

export function emptyReader():Reader
export function refReader():Ref<Reader>
export function refOfReader(x:Reader):Ref<Reader>
export function emptyWriter():Writer
export function refWriter():Ref<Writer>
export function refOfWriter(x:Writer):Ref<Writer>}
