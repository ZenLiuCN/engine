declare module "golang/math/big" {
	// @ts-ignore
	import * as fmt from 'golang/fmt'
	// @ts-ignore
	import * as rand from 'golang/math/rand'
	// @ts-ignore
	import type {uint,rune,error,int,int64,bool,Ref,float64,byte,uint64,Alias,Struct,float32} from 'go'


	export interface Word extends Alias<uint>{
	}
	export interface Float extends Struct<Float>{
		add(x,y:Ref<Float>):Ref<Float>
		setMode(mode:RoundingMode):Ref<Float>
		setString(s:string):[Ref<Float>,bool]
		append(buf:Uint8Array,fmt:byte,prec:int):Uint8Array
		setMantExp(mant:Ref<Float>,exp:int):Ref<Float>
		quo(x,y:Ref<Float>):Ref<Float>
		setFloat64(x:float64):Ref<Float>
		setInt64(x:int64):Ref<Float>
		set(x:Ref<Float>):Ref<Float>
		acc():Accuracy
		setPrec(prec:uint):Ref<Float>
		setInt(x:Ref<Int>):Ref<Float>
		setRat(x:Ref<Rat>):Ref<Float>
		gobDecode(buf:Uint8Array):error
		isInt():bool
		copy(x:Ref<Float>):Ref<Float>
		int(z:Ref<Int>):[Ref<Int>,Accuracy]
		scan(s:fmt.ScanState,ch:rune):error
		sign():int
		int64():[int64,Accuracy]
		abs(x:Ref<Float>):Ref<Float>
		gobEncode():[Uint8Array,error]
		setUint64(x:uint64):Ref<Float>
		sub(x,y:Ref<Float>):Ref<Float>
		parse(s:string,base:int):[Ref<Float>,int,error]
		unmarshalText(text:Uint8Array):error
		mode():RoundingMode
		neg(x:Ref<Float>):Ref<Float>
		mul(x,y:Ref<Float>):Ref<Float>
		text(format:byte,prec:int):string
		setInf(signbit:bool):Ref<Float>
		isInf():bool
		float32():[float32,Accuracy]
		cmp(y:Ref<Float>):int
		mantExp(mant:Ref<Float>):int
		sqrt(x:Ref<Float>):Ref<Float>
		prec():uint
		signbit():bool
		marshalText():[Uint8Array,error]
		string():string
		format(s:fmt.State,format:rune):void
		minPrec():uint
		float64():[float64,Accuracy]
		rat(z:Ref<Rat>):[Ref<Rat>,Accuracy]
		uint64():[uint64,Accuracy]

	}
	export function jacobi(x,y:Ref<Int>):int
	export function newInt(x:int64):Ref<Int>
	export function parseFloat(s:string,base:int,prec:uint,mode:RoundingMode):[Ref<Float>,int,error]
	export interface Int extends Struct<Int>{
		setString(s:string,base:int):[Ref<Int>,bool]
		bit(i:int):uint
		cmp(y:Ref<Int>):int
		rand(rnd:rand.Rand,n:Ref<Int>):Ref<Int>
		isInt64():bool
		setBytes(buf:Uint8Array):Ref<Int>
		exp(x,y,m:Ref<Int>):Ref<Int>
		setBits(abs:Word[]):Ref<Int>
		binomial(n,k:int64):Ref<Int>
		int64():int64
		setBit(x:Ref<Int>,i:int,b:uint):Ref<Int>
		andNot(x,y:Ref<Int>):Ref<Int>
		setUint64(x:uint64):Ref<Int>
		add(x,y:Ref<Int>):Ref<Int>
		lsh(x:Ref<Int>,n:uint):Ref<Int>
		rsh(x:Ref<Int>,n:uint):Ref<Int>
		xor(x,y:Ref<Int>):Ref<Int>
		scan(s:fmt.ScanState,ch:rune):error
		gobDecode(buf:Uint8Array):error
		unmarshalJson(text:Uint8Array):error
		divMod(x,y,m:Ref<Int>):[Ref<Int>,Ref<Int>]
		cmpAbs(y:Ref<Int>):int
		sign():int
		gobEncode():[Uint8Array,error]
		isUint64():bool
		string():string
		marshalText():[Uint8Array,error]
		rem(x,y:Ref<Int>):Ref<Int>
		quoRem(x,y,r:Ref<Int>):[Ref<Int>,Ref<Int>]
		modSqrt(x,p:Ref<Int>):Ref<Int>
		quo(x,y:Ref<Int>):Ref<Int>
		bytes():Uint8Array
		setInt64(x:int64):Ref<Int>
		abs(x:Ref<Int>):Ref<Int>
		marshalJson():[Uint8Array,error]
		text(base:int):string
		and(x,y:Ref<Int>):Ref<Int>
		sqrt(x:Ref<Int>):Ref<Int>
		format(s:fmt.State,ch:rune):void
		bitLen():int
		trailingZeroBits():uint
		modInverse(g,n:Ref<Int>):Ref<Int>
		unmarshalText(text:Uint8Array):error
		neg(x:Ref<Int>):Ref<Int>
		set(x:Ref<Int>):Ref<Int>
		mulRange(a,b:int64):Ref<Int>
		div(x,y:Ref<Int>):Ref<Int>
		uint64():uint64
		gcd(x,y,a,b:Ref<Int>):Ref<Int>
		sub(x,y:Ref<Int>):Ref<Int>
		mul(x,y:Ref<Int>):Ref<Int>
		probablyPrime(n:int):bool
		or(x,y:Ref<Int>):Ref<Int>
		not(x:Ref<Int>):Ref<Int>
		float64():[float64,Accuracy]
		fillBytes(buf:Uint8Array):Uint8Array
		append(buf:Uint8Array,base:int):Uint8Array
		bits():Word[]
		mod(x,y:Ref<Int>):Ref<Int>

	}
	export interface Rat extends Struct<Rat>{
		floatString(prec:int):string
		set(x:Ref<Rat>):Ref<Rat>
		sign():int
		cmp(y:Ref<Rat>):int
		denom():Ref<Int>
		setFloat64(f:float64):Ref<Rat>
		float32():[float32,bool]
		setInt64(x:int64):Ref<Rat>
		ratString():string
		marshalText():[Uint8Array,error]
		setInt(x:Ref<Int>):Ref<Rat>
		mul(x,y:Ref<Rat>):Ref<Rat>
		setString(s:string):[Ref<Rat>,bool]
		neg(x:Ref<Rat>):Ref<Rat>
		quo(x,y:Ref<Rat>):Ref<Rat>
		scan(s:fmt.ScanState,ch:rune):error
		gobEncode():[Uint8Array,error]
		setFrac(a,b:Ref<Int>):Ref<Rat>
		isInt():bool
		num():Ref<Int>
		sub(x,y:Ref<Rat>):Ref<Rat>
		string():string
		unmarshalText(text:Uint8Array):error
		setUint64(x:uint64):Ref<Rat>
		inv(x:Ref<Rat>):Ref<Rat>
		add(x,y:Ref<Rat>):Ref<Rat>
		gobDecode(buf:Uint8Array):error
		float64():[float64,bool]
		setFrac64(a,b:int64):Ref<Rat>
		abs(x:Ref<Rat>):Ref<Rat>

	}
	export const MaxExp=2147483647
	export const MinExp=-2147483648
	export const MaxPrec=4294967295
	export const ToNearestAway:RoundingMode
	export const ToNearestEven:RoundingMode
	export const ToZero:RoundingMode
	export const AwayFromZero:RoundingMode
	export const ToNegativeInf:RoundingMode
	export const ToPositiveInf:RoundingMode
	export interface Accuracy extends Alias<Accuracy>{
		string():string

	}
	export interface ErrNaN extends Struct<ErrNaN>{
		error():string

	}
	export function newFloat(x:float64):Ref<Float>
	export interface RoundingMode extends Alias<RoundingMode>{
		string():string

	}
	export function newRat(a,b:int64):Ref<Rat>
	export const MaxBase=62

}