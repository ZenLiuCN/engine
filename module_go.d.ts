// noinspection JSUnusedGlobalSymbols

/**
 * the module contains basic and built-in function and types for go. except generic types
 */
declare module "go" {
    export type int = number
    export type int8 = number
    export type int16 = number
    export type int32 = number
    export type int64 = number
    export type uint = number
    export type byte = uint8
    export type uint8 = number
    export type uint16 = number
    export type uint32 = number
    export type uint64 = number
    export type rune = int32
    export type float64 = number
    export type float32 = number
    export type error = Error | undefined

    export interface uintptr {
    }


    export type bool = boolean
    //A go pointer of type T
    export type Ref<T> = T | undefined
    //A go refined type
    export type Alias<T> = T
    //A go Interface of type T
    export type Proto<T> = T
    //A go struct of type T
    export type Struct<T> = T
    //A go empty anonymous struct
    export type Nothing = Struct<{}>

    export interface TypeInfo<T> {
        toString(): string
    }

    export interface TypeUsage<T> {
        slice?: (len?: int, cap?: int) => T
    }

    export function typeOf<T>(c: T): TypeInfo<T>

    export function usage<T>(t: TypeInfo<T>): TypeUsage<T> | undefined

    export function imag32(c: complex64): float32

    export function real32(c: complex64): float32

    export function imag64(c: complex64): float32

    export function real64(c: complex64): float32

    export function complex32(r, i: float32): complex64

    export function complex64(r, i: float64): complex64

    export interface complex128 {

    }

    export interface complex64 {

    }

    export interface Stringer {
        string(): string
    }

    export interface GoStringer {
        goString(): string
    }

    export class Text {
        constructor(t?: string | Uint8Array | rune[])

        bytes(): Uint8Array

        runes(): rune[]

        string(): string

        toString(): string
    }

    export function bytesFromString(t: string): Uint8Array

    export function runesFromString(t: string): rune[]

    export function stringFromRunes(t: rune[]): string

    export function toInt8(t: number): int8

    export function toInt16(t: number): int16

    export function toInt32(t: number): int32

    export function toInt64(t: number): int64

    export function toUint8(t: number): uint8

    export function toUint16(t: number): uint16

    export function toUint32(t: number): uint32

    export function toUint64(t: number): uint64

    export interface Chan<T> extends ChanRecv<T>, ChanSend<T> {


    }

    export interface ChanRecv<T> {

        recv(handle: (t: T) => void): Promise<void>

        closed(): bool

        close(): error
    }

    export interface ChanSend<T> {

        send(v: T)

        closed(): boolean

        close(): error
    }

    //Helper type
    export interface Maybe<V> {
        Value?: V
        Error?: Error

        result(): [V, Error]
    }

    /**
     * convert int64 inside object to string, fail if any property is not an int64
     * @param v object
     * @param keys the keys of big integer
     */
    export function intToString(v: Record<string, any>, keys: string[]): Record<string, any>

    /**
     * convert string inside object to int64, fail if any property is not an int64
     * @param v object
     * @param keys the keys of big integer
     */
    export function intFromString(v: Record<string, any>, keys: string[]): Record<string, any>

    /**
     * convert bigint in objects to string, fail if any property is not an int64
     * @param v object array
     * @param keys the keys of big integer properties
     */
    export function intToStringArray(v: Record<string, any>[], keys: string[]): Record<string, any>[]

    /**
     * convert bigint in objects to string, fail if any property is not an int64
     * @param v object array
     * @param keys the keys of big integer properties
     */
    export function intFromStringArray(v: Record<string, any>[], keys: string[]): Record<string, any>[]
}
