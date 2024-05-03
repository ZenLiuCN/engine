declare module "golang/mime" {
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {Struct,error,Alias} from 'go'


	export function extensionsByType(typ:string):[string[],error]
	export function addExtensionType(ext,typ:string):error
	export const BEncoding=98
	export const QEncoding=113
	export interface WordEncoder extends Alias<WordEncoder>{
		encode(charset,s:string):string

	}
	export interface WordDecoder extends Struct<WordDecoder>{
		charsetReader:(charset:string,input:io.Reader)=>[io.Reader,error]
		decode(word:string):[string,error]
		decodeHeader(header:string):[string,error]

	}
	export function formatMediaType(t:string,param:Record<string,string>):string
	export function parseMediaType(v:string):[string,Record<string,string>,error]
	export function typeByExtension(ext:string):string

}