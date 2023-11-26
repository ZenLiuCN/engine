
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
    bytes(): BinaryData

    /**
     * deserialize from pem bytes
     */
    load(bin: BinaryData)

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
    bytes(): BinaryData

    /**
     * deserialize from pem bytes
     */
    load(bin: BinaryData)

    equal(pk: PublicKey): boolean
}

declare interface Cipher {
    crypto(key,src: BinaryData): BinaryData
}

declare const Symmetric: {
    aes(): CipherAlg
    des(): CipherAlg
    cbc(iv:BinaryData,encrypt:boolean): BlockMode
    ecb(encrypt:boolean): BlockMode
    cfb(iv:BinaryData,encrypt:boolean): StreamMode
    ctr(iv:BinaryData): StreamMode
    ofb(iv:BinaryData): StreamMode
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
        nonce?  : BinaryData
        label?  : BinaryData
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
    parsePrivateKey(pem: BinaryData): PrivateKey
    parsePublicKey(pem: BinaryData): PublicKey
    /**
     * signature for data
     * @param key the private key
     * @param data the original data
     * @param hash the hash of hasher, require for RSA
     * @param opt use PSS alg when opt presents, otherwise use PKCS1v15, optional for RSA
     * @returns binary signature
     */
    sign(key: PrivateKey, data: BinaryData, hash?: Hash, opt?: {
        //-1: length equal hash  0: auto, default  positive: salt length
        saltLength?: number
        hash: Hash
    }): BinaryData
    /**
     * verify a signature
     * @param key the public key
     * @param data the original data
     * @param sign the signature
     * @param hash the hash of hasher, require for RSA
     * @param opt use PSS alg when opt presents, otherwise use PKCS1v15, optional for RSA
     * @returns valid or not
     */
    verify(key: PublicKey, data, sign: BinaryData, hash?: Hash, opt?: {
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
    encrypt(key: PublicKey, data: BinaryData, hash?: HashFunc): BinaryData
    /**
     *
     * @param key support ECDSA and RSA
     * @param secret the encrypted data
     * @param hash the HashFunc for RSA OAEP decrypt
     */
    decrypt(key: PrivateKey, secret: BinaryData, hash?: HashFunc): BinaryData
}