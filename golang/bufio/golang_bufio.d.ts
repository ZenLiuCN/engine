declare module "golang/bufio" {
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {error,bool,Ref,int64,Struct,byte,rune,int} from 'go'


	export interface Writer extends Struct<Writer>{
		size():int
		availableBuffer():Uint8Array
		writeByte(c:byte):error
		writeString(s:string):[int,error]
		writeRune(r:rune):[int,error]
		readFrom(r:io.Reader):[int64,error]
		reset(w:io.Writer):void
		flush():error
		available():int
		buffered():int
		write(p:Uint8Array):[int,error]

	}
	export function scanRunes(data:Uint8Array,atEOF:bool):[int,Uint8Array,error]
	export function scanLines(data:Uint8Array,atEOF:bool):[int,Uint8Array,error]
	export function scanWords(data:Uint8Array,atEOF:bool):[int,Uint8Array,error]
	export interface Reader extends Struct<Reader>{
		unreadByte():error
		buffered():int
		readLine():[Uint8Array,bool,error]
		readBytes(delim:byte):[Uint8Array,error]
		reset(r:io.Reader):void
		read(p:Uint8Array):[int,error]
		readString(delim:byte):[string,error]
		writeTo(w:io.Writer):[int64,error]
		peek(n:int):[Uint8Array,error]
		discard(n:int):[int,error]
		readRune():[rune,int,error]
		readSlice(delim:byte):[Uint8Array,error]
		size():int
		readByte():[byte,error]
		unreadRune():error

	}

	export type SplitFunc = (data:Uint8Array,atEOF:bool)=>[int,Uint8Array,error]
	export function newScanner(r:io.Reader):Ref<Scanner>
	export function newReader(rd:io.Reader):Ref<Reader>
	export interface ReadWriter extends Struct<ReadWriter>,Ref<Reader>,Ref<Writer>{
		reader:Ref<Reader>

		writer:Ref<Writer>

	}
	export interface Scanner extends Struct<Scanner>{
		bytes():Uint8Array
		text():string
		scan():bool
		buffer(buf:Uint8Array,max:int):void
		split(split:SplitFunc):void
		err():error

	}
	export function scanBytes(data:Uint8Array,atEOF:bool):[int,Uint8Array,error]
	export const MaxScanTokenSize=65536
	export const startBufSize=4096
	export function newReaderSize(rd:io.Reader,size:int):Ref<Reader>
	export function newWriterSize(w:io.Writer,size:int):Ref<Writer>
	export function newWriter(w:io.Writer):Ref<Writer>
	export function newReadWriter(r:Ref<Reader>,w:Ref<Writer>):Ref<ReadWriter>

}