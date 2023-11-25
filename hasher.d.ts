declare interface Hasher {
    readonly MD4: Hash
    readonly MD5: Hash
    readonly SHA1: Hash
    readonly SHA224: Hash
    readonly SHA256: Hash
    readonly SHA384: Hash
    readonly SHA512: Hash
    readonly MD5SHA1: Hash
    readonly RIPEMD160: Hash
    readonly SHA3_224: Hash
    readonly SHA3_256: Hash
    readonly SHA3_384: Hash
    readonly SHA3_512: Hash
    readonly SHA512_224: Hash
    readonly SHA512_256: Hash
    readonly BLAKE2s_256: Hash
    readonly BLAKE2b_256: Hash
    readonly BLAKE2b_384: Hash
    readonly BLAKE2b_512: Hash
}

declare interface HashFunc {
    sum(bin: BinaryData): BinaryData

    reset()

    size(): number

    blockSize(): number
}

declare interface Hash {
    available(): boolean

    get(): HashFunc
}

declare interface Codec {
    base64UrlEncode(b: BinaryData): string

    base64UrlDecode(b: string): BinaryData


    base64StdEncode(b: BinaryData): string

    base64StdDecode(b: string): BinaryData

    base64RawStdEncode(b: BinaryData): string

    base64RawStdDecode(b: string): BinaryData

    base64RawUrlEncode(b: BinaryData): string

    base64RawUrlDecode(b: string): BinaryData

    hexEncode(b: BinaryData): string

    hexDecode(b: string): BinaryData

    base32StdEncode(b: BinaryData): string

    base32StdDecode(b: string): BinaryData

    base32HexEncode(b: BinaryData): string

    base32HexDecode(b: string): BinaryData
}

declare const hasher: Hasher
declare const codec: Codec