// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/net/http/httputil'{
	// @ts-ignore
	import * as http from 'golang/net/http'
	// @ts-ignore
	import * as url from 'golang/net/url'
	// @ts-ignore
	import * as time from 'golang/time'
	// @ts-ignore
	import * as log from 'golang/log'
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as net from 'golang/net'
	// @ts-ignore
	import * as bufio from 'golang/bufio'
	// @ts-ignore
	import type {Struct,Ref,error,int,bool,GoError} from 'go'
	export interface BufferPool{
		get():Uint8Array
		put(v1:Uint8Array):void
	}
	export interface ClientConn extends Struct<ClientConn>,io.Closer{
		hijack():[net.Conn,Ref<bufio.Reader>]
		close():error
		write(req:Ref<http.Request>):error
		pending():int
		read(req:Ref<http.Request>):[Ref<http.Response>,error]
		Do(req:Ref<http.Request>):[Ref<http.Response>,error]
	}
	export function dumpRequest(req:Ref<http.Request>,body:bool):[Uint8Array,error]
	export function dumpRequestOut(req:Ref<http.Request>,body:bool):[Uint8Array,error]
	export function dumpResponse(resp:Ref<http.Response>,body:bool):[Uint8Array,error]
	export const ErrClosed:Ref<http.ProtocolError>
	export const ErrLineTooLong:GoError
	export const ErrPersistEOF:Ref<http.ProtocolError>
	export const ErrPipeline:Ref<http.ProtocolError>
	export function newChunkedReader(r:io.Reader):io.Reader
	export function newChunkedWriter(w:io.Writer):io.WriteCloser
	export function newClientConn(c:net.Conn,r:Ref<bufio.Reader>):Ref<ClientConn>
	export function newProxyClientConn(c:net.Conn,r:Ref<bufio.Reader>):Ref<ClientConn>
	export function newServerConn(c:net.Conn,r:Ref<bufio.Reader>):Ref<ServerConn>
	export function newSingleHostReverseProxy(target:Ref<url.URL>):Ref<ReverseProxy>
	export interface ProxyRequest extends Struct<ProxyRequest>{
		In:Ref<http.Request>
		out:Ref<http.Request>
		setURL(target:Ref<url.URL>):void
		setXForwarded():void
	}
	export interface ReverseProxy extends Struct<ReverseProxy>,http.Handler{
		rewrite:(v1:Ref<ProxyRequest>)=>void
		director:(v1:Ref<http.Request>)=>void
		transport:http.RoundTripper
		flushInterval:time.Duration
		errorLog:Ref<log.Logger>
		bufferPool:BufferPool
		modifyResponse:(v1:Ref<http.Response>)=>error
		errorHandler:(v3:http.ResponseWriter,v2:Ref<http.Request>,v1:error)=>void
		serveHTTP(rw:http.ResponseWriter,req:Ref<http.Request>):void
	}
	export interface ServerConn extends Struct<ServerConn>,io.Closer{
		hijack():[net.Conn,Ref<bufio.Reader>]
		close():error
		read():[Ref<http.Request>,error]
		pending():int
		write(req:Ref<http.Request>,resp:Ref<http.Response>):error
	}

export function emptyClientConn():ClientConn
export function refClientConn():Ref<ClientConn>
export function refOfClientConn(x:ClientConn):Ref<ClientConn>
export function emptyProxyRequest():ProxyRequest
export function refProxyRequest():Ref<ProxyRequest>
export function refOfProxyRequest(x:ProxyRequest):Ref<ProxyRequest>
export function emptyReverseProxy():ReverseProxy
export function refReverseProxy():Ref<ReverseProxy>
export function refOfReverseProxy(x:ReverseProxy):Ref<ReverseProxy>
export function emptyServerConn():ServerConn
export function refServerConn():Ref<ServerConn>
export function refOfServerConn(x:ServerConn):Ref<ServerConn>}
