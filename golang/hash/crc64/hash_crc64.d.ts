// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/hash/crc64'{
	// @ts-ignore
	import * as hash from 'golang/hash'
	// @ts-ignore
	import type {Ref,uint64,int} from 'go'
	export function checksum(data:Uint8Array,tab:Ref<Table>):uint64
	//14514072000185962306
	export const ECMA:int
	//15564440312192434176
	export const ISO:int
	export function makeTable(poly:uint64):Ref<Table>
	export function New(tab:Ref<Table>):hash.Hash64
	//8
	export const Size:int
	export interface Table extends Array<uint64>/*256*/{
	}
	export function update(crc:uint64,tab:Ref<Table>,p:Uint8Array):uint64
}
