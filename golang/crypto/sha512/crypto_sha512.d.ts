// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/crypto/sha512'{
	// @ts-ignore
	import * as hash from 'golang/hash'
	// @ts-ignore
	import type {int} from 'go'
	//128
	export const BlockSize:int
	export function New():hash.Hash
	export function new384():hash.Hash
	export function new512_224():hash.Hash
	export function new512_256():hash.Hash
	//64
	export const Size:int
	//28
	export const Size224:int
	//32
	export const Size256:int
	//48
	export const Size384:int
	export function sum384(data:Uint8Array):Uint8Array/*48*/
	export function sum512(data:Uint8Array):Uint8Array/*64*/
	export function sum512_224(data:Uint8Array):Uint8Array/*28*/
	export function sum512_256(data:Uint8Array):Uint8Array/*32*/
}
