declare module "golang/net/textproto" {
	// @ts-ignore
	import * as bufio from 'golang/bufio'
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {error,Ref,Alias,Struct,uint,int} from 'go'


	export interface MIMEHeader extends Struct<MIMEHeader>,Record<string,string[]>{
		//@ts-ignore
		set(key,value:string):void
		//@ts-ignore
		get(key:string):string
		//@ts-ignore
		values(key:string):string[]
		//@ts-ignore
		del(key:string):void
		//@ts-ignore
		add(key,value:string):void

	}
	export function canonicalMimeHeaderKey(s:string):string
	export function trimBytes(b:Uint8Array):Uint8Array
	export interface Pipeline extends Struct<Pipeline>{
		endRequest(id:uint):void
		startResponse(id:uint):void
		endResponse(id:uint):void
		next():uint
		startRequest(id:uint):void

	}
	export interface Reader extends Struct<Reader>{
		r:Ref<bufio.Reader>


		readDotBytes():[Uint8Array,error]
		dotReader():io.Reader
		readLineBytes():[Uint8Array,error]
		readDotLines():[string[],error]
		readResponse(expectCode:int):[int,string,error]
		readLine():[string,error]
		readContinuedLine():[string,error]
		readContinuedLineBytes():[Uint8Array,error]
		readCodeLine(expectCode:int):[int,string,error]
		readMimeHeader():[MIMEHeader,error]

	}
	export interface Error extends Struct<Error>{
		code:int

		msg:string
		error():string

	}
	export function newConn(conn:io.ReadWriteCloser):Ref<Conn>
	export function dial(network,addr:string):[Ref<Conn>,error]
	export function newReader(r:bufio.Reader):Ref<Reader>
	export interface ProtocolError extends Alias<string>{
		error():string

	}
	export interface Conn extends Struct<Conn>,Ref<Reader>,Ref<Writer>,Ref<Pipeline>,io.Closer{
		reader:Ref<Reader>

		writer:Ref<Writer>

		pipeline:Ref<Pipeline>

		close():error
		cmd(format:string,... args:any[]):[uint,error]

	}
	export function trimString(s:string):string
	export interface Writer extends Struct<Writer>{
		w:Ref<bufio.Writer>

		printfLine(format:string,... args:any[]):error
		dotWriter():io.WriteCloser

	}
	export function newWriter(w:bufio.Writer):Ref<Writer>

}