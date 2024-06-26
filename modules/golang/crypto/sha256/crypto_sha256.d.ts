// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/crypto/sha256'{

	// @ts-ignore
	import * as hash from 'golang/hash'
	// @ts-ignore
	import type {int,byte,Ref} from 'go'
	//64
	export const BlockSize:int
	export function New():hash.Hash

	export function new224():hash.Hash

	//32
	export const Size:int
	//28
	export const Size224:int
	export function sum224(data:Uint8Array):byte/*28*/

	export function sum256(data:Uint8Array):byte/*32*/

}