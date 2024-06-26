// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/crypto/dsa'{

	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as big from 'golang/math/big'
	// @ts-ignore
	import type {GoError,Ref,error,int,Struct,bool} from 'go'
	export const ErrInvalidPublicKey:GoError
	export function generateKey(priv:Ref<PrivateKey>,rand:io.Reader)/*error*/

	export function generateParameters(params:Ref<Parameters>,rand:io.Reader,sizes:ParameterSizes)/*error*/

	export const L1024N160:ParameterSizes
	export const L2048N224:ParameterSizes
	export const L2048N256:ParameterSizes
	export const L3072N256:ParameterSizes
	export interface ParameterSizes extends int{

	}
	export interface Parameters extends Struct<Parameters>{

			P:Ref<big.Int>
			Q:Ref<big.Int>
			G:Ref<big.Int>
	}
	export interface PrivateKey extends Struct<PrivateKey>{

			publicKey:PublicKey
			X:Ref<big.Int>
	}
	export interface PublicKey extends Struct<PublicKey>{

			parameters:Parameters
			Y:Ref<big.Int>
	}
	export function sign(rand:io.Reader,priv:Ref<PrivateKey>,hash:Uint8Array):[Ref<big.Int>,Ref<big.Int>]

	export function verify(pub:Ref<PublicKey>,hash:Uint8Array,r:Ref<big.Int>,s:Ref<big.Int>):bool

	export function emptyParameters():Parameters
	export function emptyRefParameters():Ref<Parameters>
	export function refOfParameters(x:Parameters,v:Ref<Parameters>)
	export function unRefParameters(v:Ref<Parameters>):Parameters
	export function emptyPrivateKey():PrivateKey
	export function emptyRefPrivateKey():Ref<PrivateKey>
	export function refOfPrivateKey(x:PrivateKey,v:Ref<PrivateKey>)
	export function unRefPrivateKey(v:Ref<PrivateKey>):PrivateKey
	export function emptyPublicKey():PublicKey
	export function emptyRefPublicKey():Ref<PublicKey>
	export function refOfPublicKey(x:PublicKey,v:Ref<PublicKey>)
	export function unRefPublicKey(v:Ref<PublicKey>):PublicKey
}