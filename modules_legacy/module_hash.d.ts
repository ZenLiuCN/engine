declare module "go/hash" {
    export const MD4: Hash
    export const MD5: Hash
    export const SHA1: Hash
    export const SHA224: Hash
    export const SHA256: Hash
    export const SHA384: Hash
    export const SHA512: Hash
    export const MD5SHA1: Hash
    export const RIPEMD160: Hash
    export const SHA3_224: Hash
    export const SHA3_256: Hash
    export const SHA3_384: Hash
    export const SHA3_512: Hash
    export const SHA512_224: Hash
    export const SHA512_256: Hash
    export const BLAKE2s_256: Hash
    export const BLAKE2b_256: Hash
    export const BLAKE2b_384: Hash
    export const BLAKE2b_512: Hash

    export interface HashFunc {
        sum(bin: Uint8Array): Uint8Array

        reset()

        size(): number

        blockSize(): number
    }

    export interface Hash {
        available(): boolean

        get(): HashFunc
    }
}


