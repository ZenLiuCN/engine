declare module "go/io" {
    export interface Reader {
    }

    export interface Writer {
    }

    export interface Seeker {
        seek(offset: number, whence: number): number
    }

    export interface Closer {
        close()
    }

    export interface WriteCloser extends Writer, Closer {

    }

    export interface ReadCloser extends Reader, Closer {

    }

    export function copy(w: Writer, r: Reader): number

    export function copyN(w: Writer, r: Reader, n: number): number

    export function copyBuffer(w: Writer, r: Reader, buffer: Uint8Array): number

    export function limitReader(r: Reader, n: number): Reader

    export function readAll(r: Reader): Uint8Array

    export function WriteString(w: Writer, text: string): number
}