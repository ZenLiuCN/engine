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
	import type {Ref,Struct,error,GoError,int,bool} from 'go'
	export interface CRTValue extends Struct<CRTValue>,crypto.PrivateKey,crypto.DecrypterOpts,crypto.PublicKey{
		exp:Ref<big.Int>
		coeff:Ref<big.Int>
		R:Ref<big.Int>
	}
	export function decryptOAEP(hash:hash.Hash,random:io.Reader,priv:Ref<PrivateKey>,ciphertext:Uint8Array,label:Uint8Array):[Uint8Array,error]
	export function decryptPKCS1v15(random:io.Reader,priv:Ref<PrivateKey>,ciphertext:Uint8Array):[Uint8Array,error]
	export function decryptPKCS1v15SessionKey(random:io.Reader,priv:Ref<PrivateKey>,ciphertext:Uint8Array,key:Uint8Array):error
	export function encryptOAEP(hash:hash.Hash,random:io.Reader,pub:Ref<PublicKey>,msg:Uint8Array,label:Uint8Array):[Uint8Array,error]
	export function encryptPKCS1v15(random:io.Reader,pub:Ref<PublicKey>,msg:Uint8Array):[Uint8Array,error]
	export const ErrDecryption:GoError
	export const ErrMessageTooLong:GoError
	export const ErrVerification:GoError
	export function generateKey(random:io.Reader,bits:int):[Ref<PrivateKey>,error]
	export function generateMultiPrimeKey(random:io.Reader,nprimes:int,bits:int):[Ref<PrivateKey>,error]
	export interface OAEPOptions extends Struct<OAEPOptions>,crypto.PublicKey,crypto.PrivateKey,crypto.DecrypterOpts{
		hash:crypto.Hash
		mgfHash:crypto.Hash
		label:Uint8Array
	}
	export interface PKCS1v15DecryptOptions extends crypto.PublicKey,Struct<PKCS1v15DecryptOptions>,crypto.PrivateKey,crypto.DecrypterOpts{
		sessionKeyLen:int
	}
	export interface PSSOptions extends Struct<PSSOptions>,crypto.PublicKey,crypto.SignerOpts,crypto.PrivateKey,crypto.DecrypterOpts{
		saltLength:int
		hash:crypto.Hash
		hashFunc():crypto.Hash
	}
	//0
	export const PSSSaltLengthAuto:int
	//-1
	export const PSSSaltLengthEqualsHash:int
	export interface PrecomputedValues extends crypto.PrivateKey,Struct<PrecomputedValues>,crypto.DecrypterOpts,crypto.PublicKey{
		dp:Ref<big.Int>
		dq:Ref<big.Int>
		qinv:Ref<big.Int>
		crtValues:CRTValue[]
	}
	export interface PrivateKey extends crypto.PublicKey,crypto.Signer,crypto.PrivateKey,Struct<PrivateKey>,crypto.DecrypterOpts,crypto.Decrypter{
		publicKey:PublicKey
		D:Ref<big.Int>
		primes:Ref<big.Int>[]
		precomputed:PrecomputedValues
		public():crypto.PublicKey
		equal(x:crypto.PrivateKey):bool
		sign(rand:io.Reader,digest:Uint8Array,opts:crypto.SignerOpts):[Uint8Array,error]
		decrypt(rand:io.Reader,ciphertext:Uint8Array,opts:crypto.DecrypterOpts):[Uint8Array,error]
		validate():error
		precompute():void
	}
	export interface PublicKey extends Struct<PublicKey>,crypto.DecrypterOpts,crypto.PublicKey,crypto.PrivateKey{
		N:Ref<big.Int>
		E:int
		size():int
		equal(x:crypto.PublicKey):bool
	}
	export function signPKCS1v15(random:io.Reader,priv:Ref<PrivateKey>,hash:crypto.Hash,hashed:Uint8Array):[Uint8Array,error]
	export function signPSS(rand:io.Reader,priv:Ref<PrivateKey>,hash:crypto.Hash,digest:Uint8Array,opts:Ref<PSSOptions>):[Uint8Array,error]
	export function verifyPKCS1v15(pub:Ref<PublicKey>,hash:crypto.Hash,hashed:Uint8Array,sig:Uint8Array):error
	export function verifyPSS(pub:Ref<PublicKey>,hash:crypto.Hash,digest:Uint8Array,sig:Uint8Array,opts:Ref<PSSOptions>):error

export function emptyCRTValue():CRTValue
export function refCRTValue():Ref<CRTValue>
export function refOfCRTValue(x:CRTValue):Ref<CRTValue>
export function emptyOAEPOptions():OAEPOptions
export function refOAEPOptions():Ref<OAEPOptions>
export function refOfOAEPOptions(x:OAEPOptions):Ref<OAEPOptions>
export function emptyPKCS1v15DecryptOptions():PKCS1v15DecryptOptions
export function refPKCS1v15DecryptOptions():Ref<PKCS1v15DecryptOptions>
export function refOfPKCS1v15DecryptOptions(x:PKCS1v15DecryptOptions):Ref<PKCS1v15DecryptOptions>
export function emptyPSSOptions():PSSOptions
export function refPSSOptions():Ref<PSSOptions>
export function refOfPSSOptions(x:PSSOptions):Ref<PSSOptions>
export function emptyPrecomputedValues():PrecomputedValues
export function refPrecomputedValues():Ref<PrecomputedValues>
export function refOfPrecomputedValues(x:PrecomputedValues):Ref<PrecomputedValues>
export function emptyPrivateKey():PrivateKey
export function refPrivateKey():Ref<PrivateKey>
export function refOfPrivateKey(x:PrivateKey):Ref<PrivateKey>
export function emptyPublicKey():PublicKey
export function refPublicKey():Ref<PublicKey>
export function refOfPublicKey(x:PublicKey):Ref<PublicKey>}
