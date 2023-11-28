declare module "go/codec" {
    export function base64UrlEncode(b: Uint8Array): string

    export function base64UrlDecode(b: string): Uint8Array


    export function base64StdEncode(b: Uint8Array): string

    export function base64StdDecode(b: string): Uint8Array

    export function base64RawStdEncode(b: Uint8Array): string

    export function base64RawStdDecode(b: string): Uint8Array

    export function base64RawUrlEncode(b: Uint8Array): string

    export function base64RawUrlDecode(b: string): Uint8Array

    export function hexEncode(b: Uint8Array): string

    export function hexDecode(b: string): Uint8Array

    export function base32StdEncode(b: Uint8Array): string

    export function base32StdDecode(b: string): Uint8Array

    export function base32HexEncode(b: Uint8Array): string

    export function base32HexDecode(b: string): Uint8Array
}

