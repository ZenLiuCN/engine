/**
 * Binary Buffer, act as Buffer and Slice
 */
declare class Buffer {
    constructor()
    constructor(text: string)
    constructor(bin: BinaryData)

    readonly detached: boolean

    /**
     * free the buffer (detach)
     */
    free()

    available(): number

    len(): number

    cap(): number

    truncate(n: number)

    grow(n: number)

    slice(from, to: number): Buffer

    chars(): Array<String>

    binary(): Array<number>

    eachU8(act: (u: number) => boolean)

    mapU8<T>(map: (u: number) => T): Array<T>

    eachChar(act: (u: string) => boolean)

    mapChar<T>(map: (u: string) => T): Array<T>

    /**
     * clean the buffer
     */
    reset()

    toString(): string

    writeBuffer(buf: ArrayBuffer): number

    writeText(v: string): number

    /**
     * read string until find delimiter
     * @param delimiter
     * @returns string end with delimiter
     */
    readString(delimiter: string): string

    readChar(): string

    readU8(): number

    arrayBuffer(): ArrayBuffer

    readByte(): number

    writeU8(v: number)

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

    saveTo(path: string)

    bytes(): BinaryData

    toWriter(): IoWriter
}

declare function binFrom(num: Array<number>): BinaryData
declare function binFrom(text: string): BinaryData
declare function binFrom(...num: Array<number>): BinaryData

declare function binLen(bin: BinaryData): number

declare function binGet(bin: BinaryData, index: number): number

declare function binSet(bin: BinaryData, index, value: number)

declare function binSlice(bin: BinaryData, from, to: number): BinaryData

declare function binAppend(bin: BinaryData, v: number): BinaryData

declare function binAppends(bin: BinaryData, tail: BinaryData): BinaryData
declare function binToString(bin: BinaryData): string

declare function binEach(bin: BinaryData, act: (v, i: number) => boolean)

declare function binEquals(bin1, bin2: BinaryData): boolean

declare function binMap<T>(bin: BinaryData, act: (v, i: number) => T): Array<T>

declare function binToArrayBuffer(bin: BinaryData): ArrayBuffer

declare function binToReader(bin: BinaryData): IoReader

declare function binFrom(reader: IoReader): BinaryData
/**
 * the golang []byte
 */
// @ts-ignore
declare interface BinaryData {
}

declare interface IoReader {
}

declare interface IoWriter {
}
