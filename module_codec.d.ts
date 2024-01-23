declare module "go/codec" {
    interface Codec{
        encodeToString(src:Uint8Array):string
        decodeString(src:string):Uint8Array

    }
    interface BinCodec {
        decode(dst:Uint8Array,src:Uint8Array):number
        encode(dst:Uint8Array,src:Uint8Array):number
        encodedLen(n:number):number
        decodedLen(n:number):number
    }
    export const Base64Std:Codec&BinCodec
    export const Base64Url:Codec&BinCodec
    export const Base64RawStd:Codec&BinCodec
    export const Base64RawUrl:Codec&BinCodec
    export const Base32Hex:Codec&BinCodec
    export const Base32Std:Codec&BinCodec
    export const Hex:Codec&BinCodec
    export const Utf8:Codec

}

