declare module "golang/mime" {
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {error,Alias,byte,Struct} from 'go'


	export interface WordEncoder extends Alias<byte>{
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
	export function extensionsByType(typ:string):[string[],error]
	export function addExtensionType(ext,typ:string):error
	export const BEncoding=0
	export const QEncoding=1

}