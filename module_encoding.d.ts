declare module 'go/encoding' {
    export interface Encoding {
        /**
         * encode utf8 text array to encoding
         */
        encode(u: Uint8Array): Uint8Array
        encodeText(u: string): Uint8Array

        /**
         * encode encoding text array to utf8
         */
        decode(u: Uint8Array): Uint8Array

        decodeText(u: string): Uint8Array

        toText(u:Uint8Array):string
        fromText(u:string):Uint8Array
    }

    export const GBK: Encoding
    export const GB2312: Encoding
    export const GB18030: Encoding
    export const BIG5: Encoding
    export const EucJP: Encoding
    export const EucKR: Encoding
    export const ShiftJIS: Encoding
    export const ISO2022JP: Encoding
    export const ISO8859_6: Encoding
    export const ISO8859_6E: Encoding
    export const ISO8859_6I: Encoding
    export const ISO8859_8E: Encoding
    export const ISO8859_8I: Encoding
    export const UTF8_BOM: Encoding
    export const UTF32_BOM_BE: Encoding
    export const UTF32_BOM_LE: Encoding
    export const UTF32_BE: Encoding
    export const UTF32_LE: Encoding
}
