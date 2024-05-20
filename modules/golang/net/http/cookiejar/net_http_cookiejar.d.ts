// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/net/http/cookiejar'{

	// @ts-ignore
	import * as http from 'golang/net/http'
	// @ts-ignore
	import * as url from 'golang/net/url'
	// @ts-ignore
	import type {error,Struct,Ref} from 'go'
	export interface Jar extends Struct<Jar>,http.CookieJar{

			cookies(u:Ref<url.URL>):Ref<http.Cookie>[]
			setCookies(u:Ref<url.URL>,cookies:Ref<http.Cookie>[]):void
	}
	export function New(o:Ref<Options>):Ref<Jar>

	export interface Options extends Struct<Options>{

			publicSuffixList:PublicSuffixList
	}
	export interface PublicSuffixList{

			publicSuffix(domain:string):string
			string():string
	}
	export function emptyJar():Jar
	export function emptyRefJar():Ref<Jar>
	export function refOfJar(x:Jar,v:Ref<Jar>)
	export function unRefJar(v:Ref<Jar>):Jar
	export function emptyOptions():Options
	export function emptyRefOptions():Ref<Options>
	export function refOfOptions(x:Options,v:Ref<Options>)
	export function unRefOptions(v:Ref<Options>):Options
}