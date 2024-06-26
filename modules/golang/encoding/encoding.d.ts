// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/encoding'{

	// @ts-ignore
	import type {error,Ref} from 'go'
	export interface BinaryMarshaler{

			marshalBinary():Uint8Array
	}
	export interface BinaryUnmarshaler{

			unmarshalBinary(data:Uint8Array)/*error*/
	}
	export interface TextMarshaler{

			marshalText():Uint8Array
	}
	export interface TextUnmarshaler{

			unmarshalText(text:Uint8Array)/*error*/
	}
}