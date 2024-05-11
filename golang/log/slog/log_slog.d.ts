// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/log/slog'{
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import * as time from 'golang/time'
	// @ts-ignore
	import * as encoding from 'golang/encoding'
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as log from 'golang/log'
	// @ts-ignore
	import * as json from 'golang/encoding/json'
	// @ts-ignore
	import * as fmt from 'golang/fmt'
	// @ts-ignore
	import type {Struct,Ref,float64,int,int64,uintptr,uint64,bool,error} from 'go'
	export function any(key:string,value:any):Attr
	export function anyValue(v:any):Value
	export interface Attr extends Struct<Attr>,json.Token,fmt.Stringer{
		key:string
		value:Value
		equal(b:Attr):bool
		string():string
	}
	export function bool(key:string,v:bool):Attr
	export function boolValue(v:bool):Value
	export function debug(msg:string,...args:any[]):void
	export function debugContext(ctx:context.Context,msg:string,...args:any[]):void
	export function Default():Ref<Logger>
	export function duration(key:string,v:time.Duration):Attr
	export function durationValue(v:time.Duration):Value
	export function error(msg:string,...args:any[]):void
	export function errorContext(ctx:context.Context,msg:string,...args:any[]):void
	export function float64(key:string,v:float64):Attr
	export function float64Value(v:float64):Value
	export function group(key:string,...args:any[]):Attr
	export function groupValue(...as:Attr[]):Value
	export interface Handler{
		enabled(v2:context.Context,v1:Level):bool
		handle(v2:context.Context,v1:Record)/*error*/
		withAttrs(attrs:Attr[]):Handler
		withGroup(name:string):Handler
	}
	export interface HandlerOptions extends Struct<HandlerOptions>,json.Token{
		addSource:bool
		level:Leveler
		replaceAttr:(groups:string[],a:Attr)=>Attr
	}
	export function info(msg:string,...args:any[]):void
	export function infoContext(ctx:context.Context,msg:string,...args:any[]):void
	export function int(key:string,value:int):Attr
	export function int64(key:string,value:int64):Attr
	export function int64Value(v:int64):Value
	export function intValue(v:int):Value
	export interface JSONHandler extends Struct<JSONHandler>,Handler,json.Token{
		enabled(_:context.Context,level:Level):bool
		withAttrs(attrs:Attr[]):Handler
		withGroup(name:string):Handler
		handle(_:context.Context,r:Record)/*error*/
	}
	export interface Kind extends int{
		string():string
	}
	export const KindAny:Kind
	export const KindBool:Kind
	export const KindDuration:Kind
	export const KindFloat64:Kind
	export const KindGroup:Kind
	export const KindInt64:Kind
	export const KindLogValuer:Kind
	export const KindString:Kind
	export const KindTime:Kind
	export const KindUint64:Kind
	export interface Level extends int{
		string():string
		marshalJSON():Uint8Array
		unmarshalJSON(data:Uint8Array)/*error*/
		marshalText():Uint8Array
		unmarshalText(data:Uint8Array)/*error*/
		level():Level
	}
	export const LevelDebug:Level
	export const LevelError:Level
	export const LevelInfo:Level
	//"level"
	export const LevelKey:string
	export interface LevelVar extends Struct<LevelVar>,Leveler,fmt.Stringer,encoding.TextUnmarshaler,json.Token,encoding.TextMarshaler{
		level():Level
		set(l:Level):void
		string():string
		marshalText():Uint8Array
		unmarshalText(data:Uint8Array)/*error*/
	}
	export const LevelWarn:Level
	export interface Leveler{
		level():Level
	}
	export function log(ctx:context.Context,level:Level,msg:string,...args:any[]):void
	export function logAttrs(ctx:context.Context,level:Level,msg:string,...attrs:Attr[]):void
	export interface LogValuer{
		logValue():Value
	}
	export interface Logger extends Struct<Logger>,json.Token{
		handler():Handler
		With(...args:any[]):Ref<Logger>
		withGroup(name:string):Ref<Logger>
		enabled(ctx:context.Context,level:Level):bool
		log(ctx:context.Context,level:Level,msg:string,...args:any[]):void
		logAttrs(ctx:context.Context,level:Level,msg:string,...attrs:Attr[]):void
		debug(msg:string,...args:any[]):void
		debugContext(ctx:context.Context,msg:string,...args:any[]):void
		info(msg:string,...args:any[]):void
		infoContext(ctx:context.Context,msg:string,...args:any[]):void
		warn(msg:string,...args:any[]):void
		warnContext(ctx:context.Context,msg:string,...args:any[]):void
		error(msg:string,...args:any[]):void
		errorContext(ctx:context.Context,msg:string,...args:any[]):void
	}
	//"msg"
	export const MessageKey:string
	export function New(h:Handler):Ref<Logger>
	export function newJSONHandler(w:io.Writer,opts:Ref<HandlerOptions>):Ref<JSONHandler>
	export function newLogLogger(h:Handler,level:Level):Ref<log.Logger>
	export function newRecord(t:time.Time,level:Level,msg:string,pc:uintptr):Record
	export function newTextHandler(w:io.Writer,opts:Ref<HandlerOptions>):Ref<TextHandler>
	export interface Record extends Struct<Record>,json.Token{
		time:time.Time
		message:string
		level:Level
		pc:uintptr
		clone():Record
		numAttrs():int
		attrs(f:(v1:Attr)=>bool):void
		addAttrs(...attrs:Attr[]):void
		add(...args:any[]):void
	}
	export function setDefault(l:Ref<Logger>):void
	export interface Source extends Struct<Source>,json.Token{
		function:string
		file:string
		line:int
	}
	//"source"
	export const SourceKey:string
	export function string(key:string,value:string):Attr
	export function stringValue(value:string):Value
	export interface TextHandler extends Struct<TextHandler>,Handler,json.Token{
		enabled(_:context.Context,level:Level):bool
		withAttrs(attrs:Attr[]):Handler
		withGroup(name:string):Handler
		handle(_:context.Context,r:Record)/*error*/
	}
	export function time(key:string,v:time.Time):Attr
	//"time"
	export const TimeKey:string
	export function timeValue(v:time.Time):Value
	export function uint64(key:string,v:uint64):Attr
	export function uint64Value(v:uint64):Value
	export interface Value extends Struct<Value>,fmt.Stringer,json.Token{
		kind():Kind
		any():any
		string():string
		int64():int64
		uint64():uint64
		bool():bool
		duration():time.Duration
		float64():float64
		time():time.Time
		logValuer():LogValuer
		group():Attr[]
		equal(w:Value):bool
		resolve():Value
	}
	export function warn(msg:string,...args:any[]):void
	export function warnContext(ctx:context.Context,msg:string,...args:any[]):void
	export function With(...args:any[]):Ref<Logger>

export function emptyJSONHandler():JSONHandler
export function refJSONHandler():Ref<JSONHandler>
export function refOfJSONHandler(x:JSONHandler):Ref<JSONHandler>
export function emptyLogger():Logger
export function refLogger():Ref<Logger>
export function refOfLogger(x:Logger):Ref<Logger>
export function emptySource():Source
export function refSource():Ref<Source>
export function refOfSource(x:Source):Ref<Source>
export function emptyValue():Value
export function refValue():Ref<Value>
export function refOfValue(x:Value):Ref<Value>
export function emptyAttr():Attr
export function refAttr():Ref<Attr>
export function refOfAttr(x:Attr):Ref<Attr>
export function emptyHandlerOptions():HandlerOptions
export function refHandlerOptions():Ref<HandlerOptions>
export function refOfHandlerOptions(x:HandlerOptions):Ref<HandlerOptions>
export function emptyLevelVar():LevelVar
export function refLevelVar():Ref<LevelVar>
export function refOfLevelVar(x:LevelVar):Ref<LevelVar>
export function emptyRecord():Record
export function refRecord():Ref<Record>
export function refOfRecord(x:Record):Ref<Record>
export function emptyTextHandler():TextHandler
export function refTextHandler():Ref<TextHandler>
export function refOfTextHandler(x:TextHandler):Ref<TextHandler>
}
