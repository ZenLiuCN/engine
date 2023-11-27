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
    base64UrlEncode(b: Uint8Array): string

    base64UrlDecode(b: string): Uint8Array


    base64StdEncode(b: Uint8Array): string

    base64StdDecode(b: string): Uint8Array

    base64RawStdEncode(b: Uint8Array): string

    base64RawStdDecode(b: string): Uint8Array

    base64RawUrlEncode(b: Uint8Array): string

    base64RawUrlDecode(b: string): Uint8Array

    hexEncode(b: Uint8Array): string

    hexDecode(b: string): Uint8Array

    base32StdEncode(b: Uint8Array): string

    base32StdDecode(b: string): Uint8Array

    base32HexEncode(b: Uint8Array): string

    base32HexDecode(b: string): Uint8Array
}

declare const hasher: Hasher
declare const codec: Codec
