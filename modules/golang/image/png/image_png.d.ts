// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/image/png'{

	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as image from 'golang/image'
	// @ts-ignore
	import type {error,Struct,Ref,GoError,int} from 'go'
	export const BestCompression:CompressionLevel
	export const BestSpeed:CompressionLevel
	export interface CompressionLevel extends int{

	}
	export function decode(r:io.Reader):image.Image

	export function decodeConfig(r:io.Reader):image.Config

	export const DefaultCompression:CompressionLevel
	export function encode(w:io.Writer,m:image.Image)/*error*/

	export interface Encoder extends Struct<Encoder>{

			compressionLevel:CompressionLevel
			bufferPool:EncoderBufferPool
			encode(w:io.Writer,m:image.Image)/*error*/
	}
	export interface EncoderBuffer extends Struct<EncoderBuffer>{

	}
	export interface EncoderBufferPool{

			get():Ref<EncoderBuffer>
			put(v1:Ref<EncoderBuffer>):void
	}
	export interface FormatError extends string,GoError{

	error():string
	}
	export const NoCompression:CompressionLevel
	export interface UnsupportedError extends string,GoError{

	error():string
	}
	export function emptyEncoder():Encoder
	export function emptyRefEncoder():Ref<Encoder>
	export function refOfEncoder(x:Encoder,v:Ref<Encoder>)
	export function unRefEncoder(v:Ref<Encoder>):Encoder
	export function emptyEncoderBuffer():EncoderBuffer
	export function emptyRefEncoderBuffer():Ref<EncoderBuffer>
	export function refOfEncoderBuffer(x:EncoderBuffer,v:Ref<EncoderBuffer>)
	export function unRefEncoderBuffer(v:Ref<EncoderBuffer>):EncoderBuffer
}