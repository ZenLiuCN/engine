// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/net/http/pprof'{
	// @ts-ignore
	import * as http from 'golang/net/http'
	// @ts-ignore
	import type {Ref} from 'go'
	export function cmdline(w:http.ResponseWriter,r:Ref<http.Request>):void
	export function handler(name:string):http.Handler
	export function index(w:http.ResponseWriter,r:Ref<http.Request>):void
	export function profile(w:http.ResponseWriter,r:Ref<http.Request>):void
	export function symbol(w:http.ResponseWriter,r:Ref<http.Request>):void
	export function trace(w:http.ResponseWriter,r:Ref<http.Request>):void

}
