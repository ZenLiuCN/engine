declare module "golang/bufio" {
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {int64,Struct,byte,rune,Ref,bool,int,error} from 'go'


	export function newScanner(r:io.Reader):Ref<Scanner>
	export function scanWords(data:Uint8Array,atEOF:bool):[int,Uint8Array,error]
	export const MaxScanTokenSize=0
	export const startBufSize=1
	export interface Reader extends Struct<Reader>{
		reset(r:io.Reader):void
		unreadByte():error
		readLine():[Uint8Array,bool,error]
		readBytes(delim:byte):[Uint8Array,error]
		read(p:Uint8Array):[int,error]
		readByte():[byte,error]
		readRune():[rune,int,error]
		readString(delim:byte):[string,error]
		writeTo(w:io.Writer):[int64,error]
		size():int
		peek(n:int):[Uint8Array,error]
		unreadRune():error
		discard(n:int):[int,error]
		buffered():int
		readSlice(delim:byte):[Uint8Array,error]

	}

	export type SplitFunc = (data:Uint8Array,atEOF:bool)=>[int,Uint8Array,error]
	export interface ReadWriter extends Struct<ReadWriter>,Ref<Reader>,Ref<Writer>{
		reader:Ref<Reader>

		writer:Ref<Writer>

	}
	export function newReadWriter(r:Ref<Reader>,w:Ref<Writer>):Ref<ReadWriter>
	export function scanBytes(data:Uint8Array,atEOF:bool):[int,Uint8Array,error]
	export function newReaderSize(rd:io.Reader,size:int):Ref<Reader>
	export function newReader(rd:io.Reader):Ref<Reader>
	export function newWriter(w:io.Writer):Ref<Writer>
	export function scanRunes(data:Uint8Array,atEOF:bool):[int,Uint8Array,error]
	export interface Writer extends Struct<Writer>{
		available():int
		readFrom(r:io.Reader):[int64,error]
		writeRune(r:rune):[int,error]
		size():int
		reset(w:io.Writer):void
		flush():error
		availableBuffer():Uint8Array
		buffered():int
		write(p:Uint8Array):[int,error]
		writeByte(c:byte):error
		writeString(s:string):[int,error]

	}
	export function newWriterSize(w:io.Writer,size:int):Ref<Writer>
	export interface Scanner extends Struct<Scanner>{
		err():error
		bytes():Uint8Array
		text():string
		split(split:SplitFunc):void
		scan():bool
		buffer(buf:Uint8Array,max:int):void

	}
	export function scanLines(data:Uint8Array,atEOF:bool):[int,Uint8Array,error]

}