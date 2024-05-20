// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/encoding/xml'{

	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as reflect from 'golang/reflect'
	// @ts-ignore
	import type {GoError,Struct,bool,map,error,Ref,int64,int} from 'go'
	export interface Attr extends Struct<Attr>,Token{

			name:Name
			value:string
	}
	export interface CharData extends Uint8Array{

			copy():CharData
	}
	export interface Comment extends Uint8Array{

			copy():Comment
	}
	export function copyToken(t:Token):Token

	export interface Decoder extends TokenReader,Token,Struct<Decoder>{

			strict:bool
			autoClose:string[]
			entity:map<string,string>
			charsetReader:(charset:string,input:io.Reader)=>io.Reader
			defaultSpace:string
			decode(v:any)/*error*/
			decodeElement(v:any,start:Ref<StartElement>)/*error*/
			skip()/*error*/
			token():Token
			rawToken():Token
			inputOffset():int64
			inputPos():[int,int]
	}
	export interface Directive extends Uint8Array{

			copy():Directive
	}
	export interface Encoder extends Struct<Encoder>,io.Closer,Token{

			indent(prefix:string,indent:string):void
			encode(v:any)/*error*/
			encodeElement(v:any,start:StartElement)/*error*/
			encodeToken(t:Token)/*error*/
			flush()/*error*/
			close():error
	}
	export interface EndElement extends Struct<EndElement>,Token{

			name:Name
	}
	export function escape(w:io.Writer,s:Uint8Array):void

	export function escapeText(w:io.Writer,s:Uint8Array)/*error*/

	export const HTMLAutoClose:string[]
	export const HTMLEntity:map<string,string>
	//"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	export const Header:string
	export function marshal(v:any):Uint8Array

	export function marshalIndent(v:any,prefix:string,indent:string):Uint8Array

	export interface Marshaler{

			marshalXML(e:Ref<Encoder>,start:StartElement)/*error*/
	}
	export interface MarshalerAttr{

			marshalXMLAttr(name:Name):Attr
	}
	export interface Name extends Struct<Name>,Token{

			space:string
			local:string
	}
	export function newDecoder(r:io.Reader):Ref<Decoder>

	export function newEncoder(w:io.Writer):Ref<Encoder>

	export function newTokenDecoder(t:TokenReader):Ref<Decoder>

	export interface ProcInst extends Struct<ProcInst>,Token{

			target:string
			inst:Uint8Array
			copy():ProcInst
	}
	export interface StartElement extends Struct<StartElement>,Token{

			name:Name
			attr:Attr[]
			copy():StartElement
			end():EndElement
	}
	export interface SyntaxError extends Struct<SyntaxError>,Error,GoError{

			msg:string
			line:int
			error():string
	}
	export interface TagPathError extends Struct<TagPathError>,Error,GoError{

			struct:reflect.Type
			field1:string
			tag1:string
			field2:string
			tag2:string
			error():string
	}
	export interface Token{

	}
	export interface TokenReader{

			token():Token
	}
	export function unmarshal(data:Uint8Array,v:any)/*error*/

	export interface UnmarshalError extends GoError,string{

	error():string
	}
	export interface Unmarshaler{

			unmarshalXML(d:Ref<Decoder>,start:StartElement)/*error*/
	}
	export interface UnmarshalerAttr{

			unmarshalXMLAttr(attr:Attr)/*error*/
	}
	export interface UnsupportedTypeError extends Struct<UnsupportedTypeError>,Error,GoError{

			type:reflect.Type
			error():string
	}
	export function emptyAttr():Attr
	export function emptyRefAttr():Ref<Attr>
	export function refOfAttr(x:Attr,v:Ref<Attr>)
	export function unRefAttr(v:Ref<Attr>):Attr
	export function emptyEndElement():EndElement
	export function emptyRefEndElement():Ref<EndElement>
	export function refOfEndElement(x:EndElement,v:Ref<EndElement>)
	export function unRefEndElement(v:Ref<EndElement>):EndElement
	export function emptyProcInst():ProcInst
	export function emptyRefProcInst():Ref<ProcInst>
	export function refOfProcInst(x:ProcInst,v:Ref<ProcInst>)
	export function unRefProcInst(v:Ref<ProcInst>):ProcInst
	export function emptyStartElement():StartElement
	export function emptyRefStartElement():Ref<StartElement>
	export function refOfStartElement(x:StartElement,v:Ref<StartElement>)
	export function unRefStartElement(v:Ref<StartElement>):StartElement
	export function emptyName():Name
	export function emptyRefName():Ref<Name>
	export function refOfName(x:Name,v:Ref<Name>)
	export function unRefName(v:Ref<Name>):Name
}