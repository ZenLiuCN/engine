// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/crypto/elliptic'{

	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as big from 'golang/math/big'
	// @ts-ignore
	import type {Ref,bool,int,Struct,error} from 'go'
	export interface Curve{

			add(x1:Ref<big.Int>,y1:Ref<big.Int>,x2:Ref<big.Int>,y2:Ref<big.Int>):[Ref<big.Int>,Ref<big.Int>]
			double(x1:Ref<big.Int>,y1:Ref<big.Int>):[Ref<big.Int>,Ref<big.Int>]
			isOnCurve(x:Ref<big.Int>,y:Ref<big.Int>):bool
			params():Ref<CurveParams>
			scalarBaseMult(k:Uint8Array):[Ref<big.Int>,Ref<big.Int>]
			scalarMult(x1:Ref<big.Int>,y1:Ref<big.Int>,k:Uint8Array):[Ref<big.Int>,Ref<big.Int>]
	}
	export interface CurveParams extends Struct<CurveParams>,Curve{

			P:Ref<big.Int>
			N:Ref<big.Int>
			B:Ref<big.Int>
			gx:Ref<big.Int>
			gy:Ref<big.Int>
			bitSize:int
			name:string
			params():Ref<CurveParams>
			isOnCurve(x:Ref<big.Int>,y:Ref<big.Int>):bool
			add(x1:Ref<big.Int>,y1:Ref<big.Int>,x2:Ref<big.Int>,y2:Ref<big.Int>):[Ref<big.Int>,Ref<big.Int>]
			double(x1:Ref<big.Int>,y1:Ref<big.Int>):[Ref<big.Int>,Ref<big.Int>]
			scalarMult(Bx:Ref<big.Int>,By:Ref<big.Int>,k:Uint8Array):[Ref<big.Int>,Ref<big.Int>]
			scalarBaseMult(k:Uint8Array):[Ref<big.Int>,Ref<big.Int>]
	}
	export function generateKey(curve:Curve,rand:io.Reader):[Uint8Array,Ref<big.Int>,Ref<big.Int>]

	export function marshal(curve:Curve,x:Ref<big.Int>,y:Ref<big.Int>):Uint8Array

	export function marshalCompressed(curve:Curve,x:Ref<big.Int>,y:Ref<big.Int>):Uint8Array

	export function P224():Curve

	export function P256():Curve

	export function P384():Curve

	export function P521():Curve

	export function unmarshal(curve:Curve,data:Uint8Array):[Ref<big.Int>,Ref<big.Int>]

	export function unmarshalCompressed(curve:Curve,data:Uint8Array):[Ref<big.Int>,Ref<big.Int>]

	export function emptyCurveParams():CurveParams
	export function emptyRefCurveParams():Ref<CurveParams>
	export function refOfCurveParams(x:CurveParams,v:Ref<CurveParams>)
	export function unRefCurveParams(v:Ref<CurveParams>):CurveParams
}