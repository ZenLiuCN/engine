declare interface Jose {
    generateRsaKey(bits: number): PrivateKey
}

declare interface PublicKey {

}

declare interface Hash {
}

declare interface SignerOpts {
    hashFunc(): Hash
}

declare interface PrivateKey {
    public(): PublicKey

    equal(x: PrivateKey): boolean

    sign(digest: ArrayBuffer, opts: SignerOpts): ArrayBuffer
    decrypt( msg :ArrayBuffer, opts:any): ArrayBuffer
}

declare const jose: Jose
