declare module "golang/math/big" {
	// @ts-ignore
	import * as fmt from 'golang/fmt'
	// @ts-ignore
	import * as rand from 'golang/math/rand'
	// @ts-ignore
	import type {int8,uint64,int,Alias,rune,uint,bool,Struct,float64,float32,error,Ref,int64,byte} from 'go'


	export interface Rat extends Struct<Rat>{
		marshalText():[Uint8Array,error]
		inv(x:Ref<Rat>):Ref<Rat>
		add(x,y:Ref<Rat>):Ref<Rat>
		quo(x,y:Ref<Rat>):Ref<Rat>
		string():string
		setInt64(x:int64):Ref<Rat>
		abs(x:Ref<Rat>):Ref<Rat>
		num():Ref<Int>
		floatString(prec:int):string
		float32():[float32,bool]
		mul(x,y:Ref<Rat>):Ref<Rat>
		scan(s:fmt.ScanState,ch:rune):error
		ratString():string
		denom():Ref<Int>
		sub(x,y:Ref<Rat>):Ref<Rat>
		unmarshalText(text:Uint8Array):error
		setFloat64(f:float64):Ref<Rat>
		set(x:Ref<Rat>):Ref<Rat>
		gobEncode():[Uint8Array,error]
		setFrac(a,b:Ref<Int>):Ref<Rat>
		isInt():bool
		gobDecode(buf:Uint8Array):error
		setUint64(x:uint64):Ref<Rat>
		cmp(y:Ref<Rat>):int
		setString(s:string):[Ref<Rat>,bool]
		sign():int
		float64():[float64,bool]
		setFrac64(a,b:int64):Ref<Rat>
		setInt(x:Ref<Int>):Ref<Rat>
		neg(x:Ref<Rat>):Ref<Rat>

	}
	export function newRat(a,b:int64):Ref<Rat>
	export const MaxExp=0
	export const MinExp=1
	export const MaxPrec=2
	export function newFloat(x:float64):Ref<Float>
	export function parseFloat(s:string,base:int,prec:uint,mode:RoundingMode):[Ref<Float>,int,error]
	export interface Int extends Struct<Int>{
		quo(x,y:Ref<Int>):Ref<Int>
		text(base:int):string
		cmpAbs(y:Ref<Int>):int
		modInverse(g,n:Ref<Int>):Ref<Int>
		marshalJson():[Uint8Array,error]
		setUint64(x:uint64):Ref<Int>
		sub(x,y:Ref<Int>):Ref<Int>
		quoRem(x,y,r:Ref<Int>):[Ref<Int>,Ref<Int>]
		bit(i:int):uint
		divMod(x,y,m:Ref<Int>):[Ref<Int>,Ref<Int>]
		scan(s:fmt.ScanState,ch:rune):error
		set(x:Ref<Int>):Ref<Int>
		isInt64():bool
		setString(s:string,base:int):[Ref<Int>,bool]
		rand(rnd:rand.Rand,n:Ref<Int>):Ref<Int>
		format(s:fmt.State,ch:rune):void
		unmarshalText(text:Uint8Array):error
		sign():int
		bytes():Uint8Array
		bitLen():int
		trailingZeroBits():uint
		exp(x,y,m:Ref<Int>):Ref<Int>
		and(x,y:Ref<Int>):Ref<Int>
		isUint64():bool
		setBytes(buf:Uint8Array):Ref<Int>
		gcd(x,y,a,b:Ref<Int>):Ref<Int>
		rsh(x:Ref<Int>,n:uint):Ref<Int>
		setInt64(x:int64):Ref<Int>
		abs(x:Ref<Int>):Ref<Int>
		rem(x,y:Ref<Int>):Ref<Int>
		xor(x,y:Ref<Int>):Ref<Int>
		mulRange(a,b:int64):Ref<Int>
		div(x,y:Ref<Int>):Ref<Int>
		int64():int64
		or(x,y:Ref<Int>):Ref<Int>
		cmp(y:Ref<Int>):int
		not(x:Ref<Int>):Ref<Int>
		uint64():uint64
		modSqrt(x,p:Ref<Int>):Ref<Int>
		setBit(x:Ref<Int>,i:int,b:uint):Ref<Int>
		gobEncode():[Uint8Array,error]
		unmarshalJson(text:Uint8Array):error
		neg(x:Ref<Int>):Ref<Int>
		mul(x,y:Ref<Int>):Ref<Int>
		lsh(x:Ref<Int>,n:uint):Ref<Int>
		append(buf:Uint8Array,base:int):Uint8Array
		setBits(abs:Word[]):Ref<Int>
		mod(x,y:Ref<Int>):Ref<Int>
		float64():[float64,Accuracy]
		sqrt(x:Ref<Int>):Ref<Int>
		string():string
		marshalText():[Uint8Array,error]
		probablyPrime(n:int):bool
		bits():Word[]
		add(x,y:Ref<Int>):Ref<Int>
		binomial(n,k:int64):Ref<Int>
		fillBytes(buf:Uint8Array):Uint8Array
		andNot(x,y:Ref<Int>):Ref<Int>
		gobDecode(buf:Uint8Array):error

	}
	export function newInt(x:int64):Ref<Int>
	export function jacobi(x,y:Ref<Int>):int
	export const ToNearestAway:RoundingMode
	export const ToNearestEven:RoundingMode
	export const ToZero:RoundingMode
	export const AwayFromZero:RoundingMode
	export const ToNegativeInf:RoundingMode
	export const ToPositiveInf:RoundingMode
	export const MaxBase=0
	export interface ErrNaN extends Struct<ErrNaN>{
		error():string

	}
	export interface RoundingMode extends Alias<byte>{
		string():string

	}
	export interface Float extends Struct<Float>{
		scan(s:fmt.ScanState,ch:rune):error
		text(format:byte,prec:int):string
		mul(x,y:Ref<Float>):Ref<Float>
		float64():[float64,Accuracy]
		sub(x,y:Ref<Float>):Ref<Float>
		isInf():bool
		set(x:Ref<Float>):Ref<Float>
		rat(z:Ref<Rat>):[Ref<Rat>,Accuracy]
		gobDecode(buf:Uint8Array):error
		prec():uint
		setUint64(x:uint64):Ref<Float>
		int(z:Ref<Int>):[Ref<Int>,Accuracy]
		neg(x:Ref<Float>):Ref<Float>
		setString(s:string):[Ref<Float>,bool]
		sqrt(x:Ref<Float>):Ref<Float>
		setMode(mode:RoundingMode):Ref<Float>
		marshalText():[Uint8Array,error]
		unmarshalText(text:Uint8Array):error
		format(s:fmt.State,format:rune):void
		acc():Accuracy
		isInt():bool
		setInt64(x:int64):Ref<Float>
		setInf(signbit:bool):Ref<Float>
		string():string
		mode():RoundingMode
		abs(x:Ref<Float>):Ref<Float>
		quo(x,y:Ref<Float>):Ref<Float>
		signbit():bool
		setFloat64(x:float64):Ref<Float>
		uint64():[uint64,Accuracy]
		int64():[int64,Accuracy]
		minPrec():uint
		sign():int
		copy(x:Ref<Float>):Ref<Float>
		gobEncode():[Uint8Array,error]
		append(buf:Uint8Array,fmt:byte,prec:int):Uint8Array
		setMantExp(mant:Ref<Float>,exp:int):Ref<Float>
		setInt(x:Ref<Int>):Ref<Float>
		setRat(x:Ref<Rat>):Ref<Float>
		float32():[float32,Accuracy]
		cmp(y:Ref<Float>):int
		parse(s:string,base:int):[Ref<Float>,int,error]
		setPrec(prec:uint):Ref<Float>
		mantExp(mant:Ref<Float>):int
		add(x,y:Ref<Float>):Ref<Float>

	}
	export interface Accuracy extends Alias<int8>{
		string():string

	}
	export interface Word extends Alias<uint>{
	}

}