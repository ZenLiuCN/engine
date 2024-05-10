// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github/com/dslipak/pdf'{
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as fmt from 'golang/fmt'
	// @ts-ignore
	import type {error,bool,int64,Struct,Ref,GoError,int,float64} from 'go'
	export const Array:ValueKind
	export const Bool:ValueKind
	export interface Column extends Struct<Column>{
		position:int64
		content:TextVertical
	}
	export interface Columns extends Array<Ref<Column>>{
	}
	export interface Content extends Struct<Content>{
		text:Text[]
		rect:Rect[]
	}
	export const Dict:ValueKind
	export const ErrInvalidPassword:GoError
	export interface Font extends Struct<Font>{
		V:Value
		baseFont():string
		firstChar():int
		lastChar():int
		widths():float64[]
		width(code:int):float64
		encoder():TextEncoding
	}
	export const Integer:ValueKind
	export function interpret(strm:Value,Do:(stk:Ref<Stack>,op:string)=>void):void
	export const Name:ValueKind
	export function newReader(f:io.ReaderAt,size:int64):[Ref<Reader>,error]
	export function newReaderEncrypted(f:io.ReaderAt,size:int64,pw:()=>string):[Ref<Reader>,error]
	export const Null:ValueKind
	export function open(file:string):Ref<Reader>
	export interface Outline extends Struct<Outline>{
		title:string
		child:Outline[]
	}
	export interface Page extends Struct<Page>{
		V:Value
		resources():Value
		fonts():string[]
		font(name:string):Font
		getPlainText(fonts:Record<string,Ref<Font>>):[string,error]
		getTextByColumn():[Columns,error]
		getTextByRow():[Rows,error]
		content():Content
	}
	export interface Point extends Struct<Point>{
		X:float64
		Y:float64
	}
	export interface Reader extends Struct<Reader>{
		page(num:int):Page
		numPage():int
		getPlainText():[io.Reader,error]
		outline():Outline
		trailer():Value
	}
	export const Real:ValueKind
	export interface Rect extends Struct<Rect>{
		min:Point
		max:Point
	}
	export interface Row extends Struct<Row>{
		position:int64
		content:TextHorizontal
	}
	export interface Rows extends Array<Ref<Row>>{
	}
	export interface Stack extends Struct<Stack>{
		len():int
		push(v:Value):void
		pop():Value
	}
	export const Stream:ValueKind
	export const String:ValueKind
	export interface Text extends Struct<Text>{
		font:string
		fontSize:float64
		X:float64
		Y:float64
		W:float64
		S:string
	}
	export interface TextEncoding{
		decode(raw:string):string
	}
	export interface TextHorizontal extends Array<Text>{
		len():int
		swap(i:int,j:int):void
		less(i:int,j:int):bool
	}
	export interface TextVertical extends Array<Text>{
		len():int
		swap(i:int,j:int):void
		less(i:int,j:int):bool
	}
	export interface Value extends Struct<Value>,fmt.Stringer{
		isNull():bool
		kind():ValueKind
		string():string
		bool():bool
		int64():int64
		float64():float64
		rawString():string
		text():string
		textFromUTF16():string
		name():string
		key(key:string):Value
		keys():string[]
		index(i:int):Value
		len():int
		reader():io.ReadCloser
	}
	export interface ValueKind extends int{
	}

export function emptyContent():Content
export function refContent():Ref<Content>
export function refOfContent(x:Content):Ref<Content>
export function emptyPage():Page
export function refPage():Ref<Page>
export function refOfPage(x:Page):Ref<Page>
export function emptyReader():Reader
export function refReader():Ref<Reader>
export function refOfReader(x:Reader):Ref<Reader>
export function emptyRect():Rect
export function refRect():Ref<Rect>
export function refOfRect(x:Rect):Ref<Rect>
export function emptyStack():Stack
export function refStack():Ref<Stack>
export function refOfStack(x:Stack):Ref<Stack>
export function emptyValue():Value
export function refValue():Ref<Value>
export function refOfValue(x:Value):Ref<Value>
export function emptyColumn():Column
export function refColumn():Ref<Column>
export function refOfColumn(x:Column):Ref<Column>
export function emptyFont():Font
export function refFont():Ref<Font>
export function refOfFont(x:Font):Ref<Font>
export function emptyOutline():Outline
export function refOutline():Ref<Outline>
export function refOfOutline(x:Outline):Ref<Outline>
export function emptyPoint():Point
export function refPoint():Ref<Point>
export function refOfPoint(x:Point):Ref<Point>
export function emptyRow():Row
export function refRow():Ref<Row>
export function refOfRow(x:Row):Ref<Row>
export function emptyText():Text
export function refText():Ref<Text>
export function refOfText(x:Text):Ref<Text>}
