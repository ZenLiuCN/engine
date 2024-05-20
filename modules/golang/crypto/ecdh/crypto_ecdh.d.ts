// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/crypto/ecdh'{

	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as crypto from 'golang/crypto'
	// @ts-ignore
	import type {Ref,error,Struct,bool} from 'go'
	export interface Curve{

			generateKey(rand:io.Reader):Ref<PrivateKey>
			newPrivateKey(key:Uint8Array):Ref<PrivateKey>
			newPublicKey(key:Uint8Array):Ref<PublicKey>
	}
	export function P256():Curve

	export function P384():Curve

	export function P521():Curve

	export interface PrivateKey extends Struct<PrivateKey>,crypto.PrivateKey,crypto.PublicKey,crypto.DecrypterOpts{

			ecdh(remote:Ref<PublicKey>):Uint8Array
			bytes():Uint8Array
			equal(x:crypto.PrivateKey):bool
			curve():Curve
			publicKey():Ref<PublicKey>
			public():crypto.PublicKey
	}
	export interface PublicKey extends Struct<PublicKey>,crypto.DecrypterOpts,crypto.PrivateKey,crypto.PublicKey{

			bytes():Uint8Array
			equal(x:crypto.PublicKey):bool
			curve():Curve
	}
	export function X25519():Curve

	export function emptyPrivateKey():PrivateKey
	export function emptyRefPrivateKey():Ref<PrivateKey>
	export function refOfPrivateKey(x:PrivateKey,v:Ref<PrivateKey>)
	export function unRefPrivateKey(v:Ref<PrivateKey>):PrivateKey
	export function emptyPublicKey():PublicKey
	export function emptyRefPublicKey():Ref<PublicKey>
	export function refOfPublicKey(x:PublicKey,v:Ref<PublicKey>)
	export function unRefPublicKey(v:Ref<PublicKey>):PublicKey
}