// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/unicode/utf16'{
	// @ts-ignore
	import type {rune,bool} from 'go'
	export function appendRune(a:Uint16Array,r:rune):Uint16Array
	export function decode(s:Uint16Array):rune[]
	export function decodeRune(r1:rune,r2:rune):rune
	export function encode(s:rune[]):Uint16Array
	export function encodeRune(r:rune):[rune,rune]
	export function isSurrogate(r:rune):bool
}
