// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/hash/crc32'{
	// @ts-ignore
	import * as hash from 'golang/hash'
	// @ts-ignore
	import type {int,Ref,uint32} from 'go'
	//2197175160
	export const Castagnoli:int
	export function checksum(data:Uint8Array,tab:Ref<Table>):uint32
	export function checksumIEEE(data:Uint8Array):uint32
	//3988292384
	export const IEEE:int
	export const IEEETable:Ref<Table>
	//3945912366
	export const Koopman:int
	export function makeTable(poly:uint32):Ref<Table>
	export function New(tab:Ref<Table>):hash.Hash32
	export function newIEEE():hash.Hash32
	//4
	export const Size:int
	export interface Table extends Uint32Array/*256*/{
	}
	export function update(crc:uint32,tab:Ref<Table>,p:Uint8Array):uint32
}
