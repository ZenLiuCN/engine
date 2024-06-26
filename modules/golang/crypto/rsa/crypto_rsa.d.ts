// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/crypto/rsa'{

	// @ts-ignore
	import * as big from 'golang/math/big'
	// @ts-ignore
	import * as crypto from 'golang/crypto'
	// @ts-ignore
	import * as hash from 'golang/hash'
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import type {int,bool,Ref,Struct,error,GoError} from 'go'
	export interface CRTValue extends crypto.PublicKey,crypto.DecrypterOpts,Struct<CRTValue>,crypto.PrivateKey{

			exp:Ref<big.Int>
			coeff:Ref<big.Int>
			R:Ref<big.Int>
	}
	export function decryptOAEP(hash:hash.Hash,random:io.Reader,priv:Ref<PrivateKey>,ciphertext:Uint8Array,label:Uint8Array):Uint8Array

	export function decryptPKCS1v15(random:io.Reader,priv:Ref<PrivateKey>,ciphertext:Uint8Array):Uint8Array

	export function decryptPKCS1v15SessionKey(random:io.Reader,priv:Ref<PrivateKey>,ciphertext:Uint8Array,key:Uint8Array)/*error*/

	export function encryptOAEP(hash:hash.Hash,random:io.Reader,pub:Ref<PublicKey>,msg:Uint8Array,label:Uint8Array):Uint8Array

	export function encryptPKCS1v15(random:io.Reader,pub:Ref<PublicKey>,msg:Uint8Array):Uint8Array

	export const ErrDecryption:GoError
	export const ErrMessageTooLong:GoError
	export const ErrVerification:GoError
	export function generateKey(random:io.Reader,bits:int):Ref<PrivateKey>

	export function generateMultiPrimeKey(random:io.Reader,nprimes:int,bits:int):Ref<PrivateKey>

	export interface OAEPOptions extends crypto.PrivateKey,crypto.PublicKey,Struct<OAEPOptions>,crypto.DecrypterOpts{

			hash:crypto.Hash
			mgfHash:crypto.Hash
			label:Uint8Array
	}
	export interface PKCS1v15DecryptOptions extends Struct<PKCS1v15DecryptOptions>,crypto.PrivateKey,crypto.PublicKey,crypto.DecrypterOpts{

			sessionKeyLen:int
	}
	export interface PSSOptions extends crypto.PrivateKey,crypto.SignerOpts,crypto.PublicKey,crypto.DecrypterOpts,Struct<PSSOptions>{

			saltLength:int
			hash:crypto.Hash
			hashFunc():crypto.Hash
	}
	//0
	export const PSSSaltLengthAuto:int
	//-1
	export const PSSSaltLengthEqualsHash:int
	export interface PrecomputedValues extends crypto.PublicKey,crypto.DecrypterOpts,Struct<PrecomputedValues>,crypto.PrivateKey{

			dp:Ref<big.Int>
			dq:Ref<big.Int>
			qinv:Ref<big.Int>
			crtValues:CRTValue[]
	}
	export interface PrivateKey extends crypto.Decrypter,crypto.DecrypterOpts,Struct<PrivateKey>,crypto.PrivateKey,crypto.PublicKey,crypto.Signer{

			publicKey:PublicKey
			D:Ref<big.Int>
			primes:Ref<big.Int>[]
			precomputed:PrecomputedValues
			public():crypto.PublicKey
			equal(x:crypto.PrivateKey):bool
			sign(rand:io.Reader,digest:Uint8Array,opts:crypto.SignerOpts):Uint8Array
			decrypt(rand:io.Reader,ciphertext:Uint8Array,opts:crypto.DecrypterOpts):Uint8Array
			validate()/*error*/
			precompute():void
	}
	export interface PublicKey extends Struct<PublicKey>,crypto.PublicKey,crypto.DecrypterOpts,crypto.PrivateKey{

			N:Ref<big.Int>
			E:int
			size():int
			equal(x:crypto.PublicKey):bool
	}
	export function signPKCS1v15(random:io.Reader,priv:Ref<PrivateKey>,hash:crypto.Hash,hashed:Uint8Array):Uint8Array

	export function signPSS(rand:io.Reader,priv:Ref<PrivateKey>,hash:crypto.Hash,digest:Uint8Array,opts:Ref<PSSOptions>):Uint8Array

	export function verifyPKCS1v15(pub:Ref<PublicKey>,hash:crypto.Hash,hashed:Uint8Array,sig:Uint8Array)/*error*/

	export function verifyPSS(pub:Ref<PublicKey>,hash:crypto.Hash,digest:Uint8Array,sig:Uint8Array,opts:Ref<PSSOptions>)/*error*/

	export function emptyPSSOptions():PSSOptions
	export function emptyRefPSSOptions():Ref<PSSOptions>
	export function refOfPSSOptions(x:PSSOptions,v:Ref<PSSOptions>)
	export function unRefPSSOptions(v:Ref<PSSOptions>):PSSOptions
	export function emptyPrecomputedValues():PrecomputedValues
	export function emptyRefPrecomputedValues():Ref<PrecomputedValues>
	export function refOfPrecomputedValues(x:PrecomputedValues,v:Ref<PrecomputedValues>)
	export function unRefPrecomputedValues(v:Ref<PrecomputedValues>):PrecomputedValues
	export function emptyPrivateKey():PrivateKey
	export function emptyRefPrivateKey():Ref<PrivateKey>
	export function refOfPrivateKey(x:PrivateKey,v:Ref<PrivateKey>)
	export function unRefPrivateKey(v:Ref<PrivateKey>):PrivateKey
	export function emptyPublicKey():PublicKey
	export function emptyRefPublicKey():Ref<PublicKey>
	export function refOfPublicKey(x:PublicKey,v:Ref<PublicKey>)
	export function unRefPublicKey(v:Ref<PublicKey>):PublicKey
	export function emptyCRTValue():CRTValue
	export function emptyRefCRTValue():Ref<CRTValue>
	export function refOfCRTValue(x:CRTValue,v:Ref<CRTValue>)
	export function unRefCRTValue(v:Ref<CRTValue>):CRTValue
	export function emptyOAEPOptions():OAEPOptions
	export function emptyRefOAEPOptions():Ref<OAEPOptions>
	export function refOfOAEPOptions(x:OAEPOptions,v:Ref<OAEPOptions>)
	export function unRefOAEPOptions(v:Ref<OAEPOptions>):OAEPOptions
	export function emptyPKCS1v15DecryptOptions():PKCS1v15DecryptOptions
	export function emptyRefPKCS1v15DecryptOptions():Ref<PKCS1v15DecryptOptions>
	export function refOfPKCS1v15DecryptOptions(x:PKCS1v15DecryptOptions,v:Ref<PKCS1v15DecryptOptions>)
	export function unRefPKCS1v15DecryptOptions(v:Ref<PKCS1v15DecryptOptions>):PKCS1v15DecryptOptions
}