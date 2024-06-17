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
    export type float = number
    export type error = GoError | undefined
    export type reserved = void
    //@ts-ignore
    export type map<K, V> = Record<K, V>

    export interface GoError extends Error {
        error(): string
    }

    //Pointer is the unsafe.Pointer
    export interface Pointer {
    }

    export interface uintptr {
    }

    export interface Slice<T> extends Array<T> {

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

    // @ts-ignore
    export interface TypeKind extends Alias<int> {

    }

    const TypeKindInvalid: TypeKind
    const TypeKindBool: TypeKind
    const TypeKindInt: TypeKind
    const TypeKindInt8: TypeKind
    const TypeKindInt16: TypeKind
    const TypeKindInt32: TypeKind
    const TypeKindInt64: TypeKind
    const TypeKindUint: TypeKind
    const TypeKindUint8: TypeKind
    const TypeKindUint16: TypeKind
    const TypeKindUint32: TypeKind
    const TypeKindUint64: TypeKind
    const TypeKindUintptr: TypeKind
    const TypeKindFloat32: TypeKind
    const TypeKindFloat64: TypeKind
    const TypeKindComplex64: TypeKind
    const TypeKindComplex128: TypeKind
    const TypeKindArray: TypeKind
    const TypeKindChan: TypeKind
    const TypeKindFunc: TypeKind
    const TypeKindInterface: TypeKind
    const TypeKindMap: TypeKind
    const TypeKindPointer: TypeKind
    const TypeKindSlice: TypeKind
    const TypeKindString: TypeKind
    const TypeKindStruct: TypeKind
    const TypeKindUnsafePointer: TypeKind

    export interface TypeId<T> {
        valid(): boolean

        kind(): TypeKind

        identity(): string

        string(): string
    }

    export interface TypeUsage<T> {
        id(): TypeId<T>

        //get the slice creator
        slice(): (cap?: int, len?: int) => Slice<T>

        instance(): () => T

        channel(): (buf?: int) => Chan<T>
    }

    export function elementOf<T, V>(c: TypeId<T>): TypeId<V> | undefined

    export function typeOf<T>(c: T): TypeId<T>

    export function sliceOf<T>(c: TypeId<T>): TypeId<Slice<T>>

    export function chanOf<T>(c: TypeId<T>): TypeId<Chan<T>>

    // @ts-ignore
    export function mapOf<K, V>(c: TypeId<K>, v: TypeId<V>): TypeId<Map<K, V>>

    export const typeUnit: {
        usageOf<T>(c: TypeId<T>): TypeUsage<T> | undefined
    }

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

    export function stringFromBytes(t: Uint8Array): string

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


    }

    export interface ChanSend<T> {


    }

    export interface GoChan<T> {
        raw(): Chan<T>

        asSendOnly(): GoSendChan<T>

        asRecvOnly(): GoRecvChan<T>

        recv(handle: (t: T) => void): Promise<void>

        send(v: T)

        close(): error

        stop(): boolean
    }

    export interface GoRecvChan<T> {
        raw(): ChanRecv<T>

        recv(handle: (t: T) => void): Promise<void>

    }

    export interface GoSendChan<T> {
        raw(): ChanSend<T>

        send(v: T)

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
