// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/fmt'{

	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {bool,error,rune,int,Ref} from 'go'
	export function append(b:Uint8Array,...a:any[]):Uint8Array

	export function appendf(b:Uint8Array,format:string,...a:any[]):Uint8Array

	export function appendln(b:Uint8Array,...a:any[]):Uint8Array

	export function errorf(format:string,...a:any[])/*error*/

	export function formatString(state:State,verb:rune):string

	export interface Formatter{

			format(f:State,verb:rune):void
	}
	export function fprint(w:io.Writer,...a:any[]):int

	export function fprintf(w:io.Writer,format:string,...a:any[]):int

	export function fprintln(w:io.Writer,...a:any[]):int

	export function fscan(r:io.Reader,...a:any[]):int

	export function fscanf(r:io.Reader,format:string,...a:any[]):int

	export function fscanln(r:io.Reader,...a:any[]):int

	export interface GoStringer{

			goString():string
	}
	export function print(...a:any[]):int

	export function printf(format:string,...a:any[]):int

	export function println(...a:any[]):int

	export function scan(...a:any[]):int

	export interface ScanState{

			read(buf:Uint8Array):int
			readRune():[rune,int]
			skipSpace():void
			token(skipSpace:bool,v1:(v1:rune)=>bool):Uint8Array
			unreadRune()/*error*/
			width():[int,bool]
	}
	export function scanf(format:string,...a:any[]):int

	export function scanln(...a:any[]):int

	export interface Scanner{

			scan(state:ScanState,verb:rune)/*error*/
	}
	export function sprint(...a:any[]):string

	export function sprintf(format:string,...a:any[]):string

	export function sprintln(...a:any[]):string

	export function sscan(str:string,...a:any[]):int

	export function sscanf(str:string,format:string,...a:any[]):int

	export function sscanln(str:string,...a:any[]):int

	export interface State{

			flag(c:int):bool
			precision():[int,bool]
			width():[int,bool]
			write(b:Uint8Array):int
	}
	export interface Stringer{

			string():string
	}
}