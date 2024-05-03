declare module "golang/net/textproto" {
	// @ts-ignore
	import * as bufio from 'golang/bufio'
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {Alias,int,error,Struct,uint,Ref} from 'go'


	export interface Reader extends Struct<Reader>{
		r:Ref<bufio.Reader>


		readContinuedLineBytes():[Uint8Array,error]
		readLine():[string,error]
		readLineBytes():[Uint8Array,error]
		dotReader():io.Reader
		readCodeLine(expectCode:int):[int,string,error]
		readResponse(expectCode:int):[int,string,error]
		readDotBytes():[Uint8Array,error]
		readDotLines():[string[],error]
		readMimeHeader():[MIMEHeader,error]
		readContinuedLine():[string,error]

	}
	export function trimBytes(b:Uint8Array):Uint8Array
	export interface Pipeline extends Struct<Pipeline>{
		next():uint
		startRequest(id:uint):void
		endRequest(id:uint):void
		startResponse(id:uint):void
		endResponse(id:uint):void

	}
	export function newReader(r:bufio.Reader):Ref<Reader>
	export interface Error extends Struct<Error>{
		code:int

		msg:string
		error():string

	}
	export interface Writer extends Struct<Writer>{
		w:Ref<bufio.Writer>

		printfLine(format:string,... args:any[]):error
		dotWriter():io.WriteCloser

	}
	export function newWriter(w:bufio.Writer):Ref<Writer>
	export function trimString(s:string):string
	export interface MIMEHeader extends Struct<MIMEHeader>,Record<string,string[]>{
		//@ts-ignore
		add(key,value:string):void
		//@ts-ignore
		set(key,value:string):void
		//@ts-ignore
		get(key:string):string
		//@ts-ignore
		values(key:string):string[]
		//@ts-ignore
		del(key:string):void

	}
	export function canonicalMimeHeaderKey(s:string):string
	export interface ProtocolError extends Alias<ProtocolError>{
		error():string

	}
	export function newConn(conn:io.ReadWriteCloser):Ref<Conn>
	export function dial(network,addr:string):[Ref<Conn>,error]
	export interface Conn extends Struct<Conn>,Ref<Reader>,Ref<Writer>,Ref<Pipeline>,io.Closer{
		reader:Ref<Reader>

		writer:Ref<Writer>

		pipeline:Ref<Pipeline>

		close():error
		cmd(format:string,... args:any[]):[uint,error]
io
	}

}