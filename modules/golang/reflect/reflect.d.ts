// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/reflect'{

	// @ts-ignore
	import type {uint,float64,Pointer,complex128,error,int64,uint64,bool,uintptr,int,Struct,Ref,GoError} from 'go'
	export function append(s:Value,...x:Value[]):Value

	export function appendSlice(s:Value,t:Value):Value

	export const Array:Kind
	export function arrayOf(length:int,elem:Type):Type

	export const Bool:Kind
	export const BothDir:ChanDir
	export const Chan:Kind
	export interface ChanDir extends int{

	string():string
	}
	export function chanOf(dir:ChanDir,t:Type):Type

	export const Complex128:Kind
	export const Complex64:Kind
	export function copy(dst:Value,src:Value):int

	export function deepEqual(x:any,y:any):bool

	export const Float32:Kind
	export const Float64:Kind
	export const Func:Kind
	export function funcOf(In:Type[],out:Type[],variadic:bool):Type

	export function indirect(v:Value):Value

	export const Int:Kind
	export const Int16:Kind
	export const Int32:Kind
	export const Int64:Kind
	export const Int8:Kind
	export const Interface:Kind
	export const Invalid:Kind
	export interface Kind extends uint{

	string():string
	}
	export function makeChan(typ:Type,buffer:int):Value

	export function makeFunc(typ:Type,v1:(args:Value[])=>Value[]):Value

	export function makeMap(typ:Type):Value

	export function makeMapWithSize(typ:Type,n:int):Value

	export function makeSlice(typ:Type,len:int,cap:int):Value

	export const Map:Kind
	export interface MapIter extends Struct<MapIter>{

			key():Value
			value():Value
			next():bool
			reset(v:Value):void
	}
	export function mapOf(key:Type,elem:Type):Type

	export interface Method extends Struct<Method>{

			name:string
			pkgPath:string
			type:Type
			func:Value
			index:int
			isExported():bool
	}
	export function New(typ:Type):Value

	export function newAt(typ:Type,p:Pointer):Value

	export const Pointer:Kind
	export function pointerTo(t:Type):Type

	export const Ptr:Kind
	export function ptrTo(t:Type):Type

	export const RecvDir:ChanDir
	export function select(cases:SelectCase[]):[int,Value,bool]

	export interface SelectCase extends Struct<SelectCase>{

			dir:SelectDir
			chan:Value
			send:Value
	}
	export const SelectDefault:SelectDir
	export interface SelectDir extends int{

	}
	export const SelectRecv:SelectDir
	export const SelectSend:SelectDir
	export const SendDir:ChanDir
	export const Slice:Kind
	export interface SliceHeader extends Struct<SliceHeader>{

			data:uintptr
			len:int
			cap:int
	}
	export function sliceOf(t:Type):Type

	export const String:Kind
	export interface StringHeader extends Struct<StringHeader>{

			data:uintptr
			len:int
	}
	export const Struct:Kind
	export interface StructField extends Struct<StructField>{

			name:string
			pkgPath:string
			type:Type
			tag:StructTag
			offset:uintptr
			index:int[]
			anonymous:bool
			isExported():bool
	}
	export function structOf(fields:StructField[]):Type

	export interface StructTag extends string{

	get(key:string):string
	lookup(key:string):[string,bool]
	}
	export function swapper(slice:any):(i:int,j:int)=>void

	export interface Type{

			align():int
			assignableTo(u:Type):bool
			bits():int
			chanDir():ChanDir
			comparable():bool
			convertibleTo(u:Type):bool
			elem():Type
			field(i:int):StructField
			fieldAlign():int
			fieldByIndex(index:int[]):StructField
			fieldByName(name:string):[StructField,bool]
			fieldByNameFunc(v1:(v1:string)=>bool):[StructField,bool]
			implements(u:Type):bool
			In(i:int):Type
			isVariadic():bool
			key():Type
			kind():Kind
			len():int
			method(v1:int):Method
			methodByName(v1:string):[Method,bool]
			name():string
			numField():int
			numIn():int
			numMethod():int
			numOut():int
			out(i:int):Type
			pkgPath():string
			size():uintptr
			string():string
	}
	export function typeOf(i:any):Type

	export const Uint:Kind
	export const Uint16:Kind
	export const Uint32:Kind
	export const Uint64:Kind
	export const Uint8:Kind
	export const Uintptr:Kind
	export const UnsafePointer:Kind
	export interface Value extends Struct<Value>{

			addr():Value
			bool():bool
			bytes():Uint8Array
			canAddr():bool
			canSet():bool
			call(In:Value[]):Value[]
			callSlice(In:Value[]):Value[]
			cap():int
			close():void
			canComplex():bool
			complex():complex128
			elem():Value
			field(i:int):Value
			fieldByIndex(index:int[]):Value
			fieldByIndexErr(index:int[]):Value
			fieldByName(name:string):Value
			fieldByNameFunc(v1:(v1:string)=>bool):Value
			canFloat():bool
			float():float64
			index(i:int):Value
			canInt():bool
			int():int64
			canInterface():bool
			interface():any
			interfaceData():uintptr/*2*/
			isNil():bool
			isValid():bool
			isZero():bool
			setZero():void
			kind():Kind
			len():int
			mapIndex(key:Value):Value
			mapKeys():Value[]
			setIterKey(iter:Ref<MapIter>):void
			setIterValue(iter:Ref<MapIter>):void
			mapRange():Ref<MapIter>
			method(i:int):Value
			numMethod():int
			methodByName(name:string):Value
			numField():int
			overflowComplex(x:complex128):bool
			overflowFloat(x:float64):bool
			overflowInt(x:int64):bool
			overflowUint(x:uint64):bool
			pointer():uintptr
			recv():[Value,bool]
			send(x:Value):void
			set(x:Value):void
			setBool(x:bool):void
			setBytes(x:Uint8Array):void
			setComplex(x:complex128):void
			setFloat(x:float64):void
			setInt(x:int64):void
			setLen(n:int):void
			setCap(n:int):void
			setMapIndex(key:Value,elem:Value):void
			setUint(x:uint64):void
			setPointer(x:Pointer):void
			setString(x:string):void
			slice(i:int,j:int):Value
			slice3(i:int,j:int,k:int):Value
			string():string
			tryRecv():[Value,bool]
			trySend(x:Value):bool
			type():Type
			canUint():bool
			uint():uint64
			unsafeAddr():uintptr
			unsafePointer():Pointer
			grow(n:int):void
			clear():void
			convert(t:Type):Value
			canConvert(t:Type):bool
			comparable():bool
			equal(u:Value):bool
	}
	export interface ValueError extends Struct<ValueError>,Error,GoError{

			method:string
			kind:Kind
			error():string
	}
	export function valueOf(i:any):Value

	export function visibleFields(t:Type):StructField[]

	export function zero(typ:Type):Value

	export function emptySelectCase():SelectCase
	export function emptyRefSelectCase():Ref<SelectCase>
	export function refOfSelectCase(x:SelectCase,v:Ref<SelectCase>)
	export function unRefSelectCase(v:Ref<SelectCase>):SelectCase
	export function emptySliceHeader():SliceHeader
	export function emptyRefSliceHeader():Ref<SliceHeader>
	export function refOfSliceHeader(x:SliceHeader,v:Ref<SliceHeader>)
	export function unRefSliceHeader(v:Ref<SliceHeader>):SliceHeader
	export function emptyStringHeader():StringHeader
	export function emptyRefStringHeader():Ref<StringHeader>
	export function refOfStringHeader(x:StringHeader,v:Ref<StringHeader>)
	export function unRefStringHeader(v:Ref<StringHeader>):StringHeader
	export function emptyStructField():StructField
	export function emptyRefStructField():Ref<StructField>
	export function refOfStructField(x:StructField,v:Ref<StructField>)
	export function unRefStructField(v:Ref<StructField>):StructField
	export function emptyValue():Value
	export function emptyRefValue():Ref<Value>
	export function refOfValue(x:Value,v:Ref<Value>)
	export function unRefValue(v:Ref<Value>):Value
	export function emptyMapIter():MapIter
	export function emptyRefMapIter():Ref<MapIter>
	export function refOfMapIter(x:MapIter,v:Ref<MapIter>)
	export function unRefMapIter(v:Ref<MapIter>):MapIter
	export function emptyMethod():Method
	export function emptyRefMethod():Ref<Method>
	export function refOfMethod(x:Method,v:Ref<Method>)
	export function unRefMethod(v:Ref<Method>):Method
}