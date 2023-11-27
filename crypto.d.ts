
declare interface CipherAlg{}
declare interface BlockMode{}
declare interface StreamMode{}
declare interface AEADMode{}
declare interface PaddingMode{}
declare interface PrivateKey {
    /**
     * 0: empty
     *
     * 1: RSA
     *
     * 2: ECDH
     *
     * 3: ECDSA
     *
     * 4: ED25519
     */
    readonly Alg: 0 | 1 | 2 | 3 | 4

    /**
     * serialize to pem bytes
     */
    bytes(): Uint8Array

    /**
     * deserialize from pem bytes
     */
    load(bin: Uint8Array)

    public(): PublicKey

    equal(pk: PrivateKey): boolean
}
declare interface PublicKey {
    /**
     * 0: empty
     *
     * 1: RSA
     *
     * 2: ECDH
     *
     * 3: ECDSA
     *
     * 4: ED25519
     */
    readonly Alg: 0 | 1 | 2 | 3 | 4

    /**
     * serialize to pem bytes
     */
    bytes(): Uint8Array

    /**
     * deserialize from pem bytes
     */
    load(bin: Uint8Array)

    equal(pk: PublicKey): boolean
}

declare interface Cipher {
    crypto(key,src: Uint8Array): Uint8Array
}

declare const Symmetric: {
    aes(): CipherAlg
    des(): CipherAlg
    cbc(iv:Uint8Array,encrypt:boolean): BlockMode
    ecb(encrypt:boolean): BlockMode
    cfb(iv:Uint8Array,encrypt:boolean): StreamMode
    ctr(iv:Uint8Array): StreamMode
    ofb(iv:Uint8Array): StreamMode
    gcm(): AEADMode
    pkcs7(): PaddingMode
    pkcs5(): PaddingMode
    cipher(conf:{
        cipher:  CipherAlg
        encrypt: boolean
        block?:   BlockMode
        stream?:  StreamMode
        padding?: PaddingMode
        aead?   : AEADMode
        nonce?  : Uint8Array
        label?  : Uint8Array
    }):Cipher
}
declare const Asymmetric: {

    /**
     * @param mode 1:RSA 2: ECDH 3: ECDSA 4: ED25519
     * @param opt
     */
    generateKey(mode: 1 | 2 | 3 | 4, opt: {
        //for rsa
        bits?: number
        //for ECC, X25519 only for ecdh , P224 only for ECDSA
        curve?: 'P256' | 'P384' | 'P521' | 'X25519' | 'P224'
    }): PrivateKey
    parsePrivateKey(pem: Uint8Array): PrivateKey
    parsePublicKey(pem: Uint8Array): PublicKey
    /**
     * signature for data
     * @param key the private key
     * @param data the original data
     * @param hash the hash of hasher, require for RSA
     * @param opt use PSS alg when opt presents, otherwise use PKCS1v15, optional for RSA
     * @returns binary signature
     */
    sign(key: PrivateKey, data: Uint8Array, hash?: Hash, opt?: {
        //-1: length equal hash  0: auto, default  positive: salt length
        saltLength?: number
        hash: Hash
    }): Uint8Array
    /**
     * verify a signature
     * @param key the public key
     * @param data the original data
     * @param sign the signature
     * @param hash the hash of hasher, require for RSA
     * @param opt use PSS alg when opt presents, otherwise use PKCS1v15, optional for RSA
     * @returns valid or not
     */
    verify(key: PublicKey, data, sign: Uint8Array, hash?: Hash, opt?: {
        //-1: length equal hash  0: auto, default  positive: salt length
        saltLength?: number
        hash: Hash
    }): boolean
    /**
     *
     * @param key support ECDSA and RSA
     * @param data the plain data
     * @param hash the HashFunc for RSA OAEP encrypt
     */
    encrypt(key: PublicKey, data: Uint8Array, hash?: HashFunc): Uint8Array
    /**
     *
     * @param key support ECDSA and RSA
     * @param secret the encrypted data
     * @param hash the HashFunc for RSA OAEP decrypt
     */
    decrypt(key: PrivateKey, secret: Uint8Array, hash?: HashFunc): Uint8Array
}
