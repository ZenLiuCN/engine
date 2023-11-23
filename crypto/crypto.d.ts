//=========== Crypto
declare interface Crypto {


    getRandomValues(array: ArrayBuffer): ArrayBuffer;

    randomUUID(): string;
    generateRsaKey(bits: number): PrivateKey
}

declare interface PrivateKey {

}

// @ts-ignore
declare const crypto: Crypto
