// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/net/http/fcgi'{
	// @ts-ignore
	import * as http from 'golang/net/http'
	// @ts-ignore
	import * as net from 'golang/net'
	// @ts-ignore
	import type {Ref,error,GoError} from 'go'
	export const ErrConnClosed:GoError
	export const ErrRequestAborted:GoError
	export function processEnv(r:Ref<http.Request>):Record<string,string>
	export function serve(l:net.Listener,handler:http.Handler):error
}
