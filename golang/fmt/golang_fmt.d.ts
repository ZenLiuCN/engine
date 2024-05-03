declare module "golang/fmt" {
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {int,error,bool,Proto,rune} from 'go'



	export interface State extends Proto<State>{
		write(b:Uint8Array):[int,error]

		width():[int,bool]

		precision():[int,bool]

		flag(c:int):bool

	}
	export function formatString(state:State,verb:rune):string
	export function println(... a:any[]):[int,error]
	export function scan(... a:any[]):[int,error]
	export function fscanln(r:io.Reader,... a:any[]):[int,error]
	export function scanf(format:string,... a:any[]):[int,error]
	export function fscanf(r:io.Reader,format:string,... a:any[]):[int,error]
	export function printf(format:string,... a:any[]):[int,error]
	export function appendf(b:Uint8Array,format:string,... a:any[]):Uint8Array

	export interface GoStringer extends Proto<GoStringer>{
		goString():string

	}
	export function sprint(... a:any[]):string
	export function fprintln(w:io.Writer,... a:any[]):[int,error]
	export function sprintln(... a:any[]):string

	export interface ScanState extends Proto<ScanState>{
		readRune():[rune,int,error]

		unreadRune():error

		skipSpace():void

		token(skipSpace:bool,f:(v1:rune)=>bool):[Uint8Array,error]

		width():[int,bool]

		read(buf:Uint8Array):[int,error]

	}
	export function sscanln(str:string,... a:any[]):[int,error]
	export function errorf(format:string,... a:any[]):error

	export interface Formatter extends Proto<Formatter>{
		format(f:State,verb:rune):void

	}
	export function sprintf(format:string,... a:any[]):string
	export function append(b:Uint8Array,... a:any[]):Uint8Array
	export function sscan(str:string,... a:any[]):[int,error]
	export function sscanf(str:string,format:string,... a:any[]):[int,error]
	export function print(... a:any[]):[int,error]
	export function appendln(b:Uint8Array,... a:any[]):Uint8Array

	export interface Scanner extends Proto<Scanner>{
		scan(state:ScanState,verb:rune):error

	}
	export function scanln(... a:any[]):[int,error]

	export interface Stringer extends Proto<Stringer>{
		string():string

	}
	export function fprintf(w:io.Writer,format:string,... a:any[]):[int,error]
	export function fprint(w:io.Writer,... a:any[]):[int,error]
	export function fscan(r:io.Reader,... a:any[]):[int,error]

}