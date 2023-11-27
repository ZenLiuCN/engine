/**
 * Binary Buffer, act as Buffer and Slice
 */
declare class Buffer {
    constructor()
    constructor(text: string)
    constructor(bin: Uint8Array)
    constructor(bin: Bytes)
    constructor(reader: IoReader)

    readonly detached: boolean

    /**
     * free the buffer (detach)
     */
    free()

    available(): number

    length(): number

    cap(): number

    truncate(n: number)

    grow(n: number)

    /**
     * clean the buffer
     */
    reset()

    slice(from, to: number): Buffer

    runes(): Array<String>

    bytes(): Uint8Array

    binary(): Bytes

    eachByte(act: (u: number) => boolean)

    mapByte<T>(map: (u: number) => T): Array<T>

    eachRune(act: (u: string) => boolean)

    mapRune<T>(map: (u: string) => T): Array<T>

    toString(): string

    /**
     * read string until find delimiter
     * @param delimiter one ascii character string
     * @returns string end with delimiter, or empty string when reach EOF
     */
    readString(delimiter: string): string

    writeString(v: string): number

    readByte(): number

    writeByte(v: number)

    readRune(): string

    /**
     *
     * @param r a string contains one rune
     */
    writeRune(r: string)

    arrayBuffer(): ArrayBuffer

    writeBuffer(buf: ArrayBuffer): number

    saveTo(path: string)

    /**
     * clean and load file into buffer
     * @param path
     */
    loadFile(path: string)

    /**
     * load file into buffer, keep already exists data
     * @param path
     */
    mergeFile(path: string)

    toWriter(): IoWriter

    toReader(): IoReader
}

/**
 * the golang []byte
 */
// @ts-ignore
declare class Bytes extends Array<number> {
    constructor()
    constructor(reader: IoReader)
    constructor(text: string)
    /**
     * clone
     */
    constructor(bin: Bytes)
    constructor(bin: ArrayBuffer)
    constructor(bin: Uint8Array)
    constructor(bin: Buffer)
    /**
     * @param values array of uint8
     */
    constructor(...values: number[])

    bytes(): Uint8Array

    append(v: Uint8Array): Bytes
    append(v: ArrayBuffer): Bytes
    append(v: string): Bytes
    append(v: Bytes): Bytes
    append(v: Buffer): Bytes

    slice(from, to: number): Bytes
    toText():string
    clone(): Bytes
    toReader():IoReader
}

declare interface IoReader {
}

declare interface IoWriter {
}
