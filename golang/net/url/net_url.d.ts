// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/net/url'{
	// @ts-ignore
	import * as fmt from 'golang/fmt'
	// @ts-ignore
	import type {GoError,Struct,error,bool,Ref} from 'go'
	export interface Error extends GoError,Struct<Error>,Error{
		op:string
		url:string
		err:GoError
		unwrap()/*error*/
		error():string
		timeout():bool
		temporary():bool
	}
	export interface EscapeError extends string,GoError{
		error():string
	}
	export interface InvalidHostError extends string,GoError{
		error():string
	}
	export function joinPath(base:string,...elem:string[]):string
	export function parse(rawURL:string):Ref<URL>
	export function parseQuery(query:string):Values
	export function parseRequestURI(rawURL:string):Ref<URL>
	export function pathEscape(s:string):string
	export function pathUnescape(s:string):string
	export function queryEscape(s:string):string
	export function queryUnescape(s:string):string
	export interface URL extends Struct<URL>,fmt.Stringer{
		scheme:string
		opaque:string
		user:Ref<Userinfo>
		host:string
		path:string
		rawPath:string
		omitHost:bool
		forceQuery:bool
		rawQuery:string
		fragment:string
		rawFragment:string
		escapedPath():string
		escapedFragment():string
		string():string
		redacted():string
		isAbs():bool
		parse(ref:string):Ref<URL>
		resolveReference(ref:Ref<URL>):Ref<URL>
		query():Values
		requestURI():string
		hostname():string
		port():string
		marshalBinary():Uint8Array
		unmarshalBinary(text:Uint8Array)/*error*/
		joinPath(...elem:string[]):Ref<URL>
	}
	export function user(username:string):Ref<Userinfo>
	export function userPassword(username:string,password:string):Ref<Userinfo>
	export interface Userinfo extends Struct<Userinfo>,fmt.Stringer{
		username():string
		password():[string,bool]
		string():string
	}
	export interface Values extends Record<string,string[]>{
		get(key:string):string
		set(key:string,value:string):void
		add(key:string,value:string):void
		del(key:string):void
		has(key:string):bool
		encode():string
	}

export function emptyURL():URL
export function refURL():Ref<URL>
export function refOfURL(x:URL):Ref<URL>
export function emptyUserinfo():Userinfo
export function refUserinfo():Ref<Userinfo>
export function refOfUserinfo(x:Userinfo):Ref<Userinfo>
}
