declare module "go/io" {
    export interface Reader {
        read(buf: Uint8Array): number
    }

    export interface Writer {
        write(buf: Uint8Array): number
    }

    export interface Seeker {
        seek(offset: number, whence: number): number
    }

    export interface Closer {
        close()
    }

    export interface ReadWriter extends Reader, Writer {

    }

    export interface WriteSeeker extends Writer, Seeker {

    }

    export interface ReadSeeker extends Reader, Seeker {

    }

    export interface ReadWriteSeeker extends Writer, Reader, Seeker {

    }

    export interface ReaderAt {
        readAt(p: Uint8Array, off: number): number
    }

    export interface WriterAt {
        writeAt(p: Uint8Array, off: number): number
    }

    export interface ByteReader {
        readByte(): number
    }

    export interface ByteScanner extends ByteReader {

        unreadByte()
    }

    export interface ByteWriter {
        writeByte(c: number)
    }

    export interface StringWriter {
        writeString(s: string): number
    }

    export interface ReadFrom {
        readFrom(r: Reader): number
    }

    export interface WriteTo {
        writeTo(w: Writer): number
    }

    export interface ReadSeekCloser extends ReadSeeker, Closer {

    }

    export interface ReadWriteCloser extends Writer, ReadCloser {

    }

    export interface WriteCloser extends Writer, Closer {

    }

    export interface ReadCloser extends Reader, Closer {

    }

    export function copy(w: Writer, r: Reader): number

    export function copyN(w: Writer, r: Reader, n: number): number

    export function copyBuffer(w: Writer, r: Reader, buffer: Uint8Array): number

    export function readAll(r: Reader): Uint8Array

    export function writeString(w: Writer, text: string): number

    export function readAtLeast(r: Reader, buf: Uint8Array, min: number): number

    export function readFull(r: Reader, buf: Uint8Array): number

    export function limitReader(r: Reader, n: number): Reader
    export interface SectionReader extends ReadSeeker,ReaderAt{
        size():number
    }
    export interface OffsetWriter extends WriteSeeker,WriterAt{

    }
    export interface PipeWriter extends WriteCloser{

    }
    export interface PipeReader extends ReadCloser{

    }
    export function newSectionReader(r: ReaderAt, off, n: number): SectionReader
    export function newOffsetWriter(r: WriterAt, off, n: number): OffsetWriter
    export function teeReader(r:Reader,w: Writer): Reader
    export function nopCloser(r:Reader): ReadCloser
    export function multiReader(... r:Reader[]): Reader
    export function multiWriter(... r:Writer[]): Writer
    export function pipe():{reader:PipeReader,writer:PipeWriter}
}
