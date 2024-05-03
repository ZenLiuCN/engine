declare module "golang/io" {
	// @ts-ignore
	import type {int64,int,Ref,rune,byte,Proto,error,Struct} from 'go'



	export interface ByteScanner extends Proto<ByteScanner>,ByteReader{
		unreadByte():error

	}

	export interface SectionReader extends Struct<SectionReader>{
		read(p:Uint8Array):[int,error]
		seek(offset:int64,whence:int):[int64,error]
		readAt(p:Uint8Array,off:int64):[int,error]
		size():int64

	}
	export function newOffsetWriter(w:WriterAt,off:int64):Ref<OffsetWriter>
	export function multiWriter(... writers:Writer[]):Writer

	export interface ReadWriter extends Proto<ReadWriter>,Reader,Writer{

	}

	export interface ReadSeekCloser extends Proto<ReadSeekCloser>,Reader,Seeker,Closer{

	}

	export interface WriterAt extends Proto<WriterAt>{
		writeAt(p:Uint8Array,off:int64):[int,error]

	}
	export function copy(dst:Writer,src:Reader):[int64,error]
	export function limitReader(r:Reader,n:int64):Reader
	export function nopCloser(r:Reader):ReadCloser

	export interface PipeReader extends Struct<PipeReader>,Closer{
		read(data:Uint8Array):[int,error]
		close():error
		closeWithError(err:error):error

	}

	export interface ReaderFrom extends Proto<ReaderFrom>{
		readFrom(r:Reader):[int64,error]

	}

	export interface StringWriter extends Proto<StringWriter>{
		writeString(s:string):[int,error]

	}
	export function writeString(w:Writer,s:string):[int,error]
	export function readFull(r:Reader,buf:Uint8Array):[int,error]

	export interface PipeWriter extends Struct<PipeWriter>,Closer{
		write(data:Uint8Array):[int,error]
		close():error
		closeWithError(err:error):error

	}
	export function teeReader(r:Reader,w:Writer):Reader
	export function pipe():[Ref<PipeReader>,Ref<PipeWriter>]

	export interface WriteSeeker extends Proto<WriteSeeker>,Writer,Seeker{

	}

	export interface RuneReader extends Proto<RuneReader>{
		readRune():[rune,int,error]

	}

	export interface LimitedReader extends Struct<LimitedReader>{
		r:Reader

		n:int64
		read(p:Uint8Array):[int,error]

	}

	export interface OffsetWriter extends Struct<OffsetWriter>{
		write(p:Uint8Array):[int,error]
		writeAt(p:Uint8Array,off:int64):[int,error]
		seek(offset:int64,whence:int):[int64,error]

	}

	export interface ByteReader extends Proto<ByteReader>{
		readByte():[byte,error]

	}

	export interface RuneScanner extends Proto<RuneScanner>,RuneReader{
		unreadRune():error

	}
	export function copyBuffer(dst:Writer,src:Reader,buf:Uint8Array):[int64,error]

	export interface Closer extends Proto<Closer>{
		close():error

	}

	export interface ReadCloser extends Proto<ReadCloser>,Reader,Closer{

	}

	export interface ReadWriteCloser extends Proto<ReadWriteCloser>,Reader,Writer,Closer{

	}

	export interface ReadWriteSeeker extends Proto<ReadWriteSeeker>,Reader,Writer,Seeker{

	}

	export interface ReadSeeker extends Proto<ReadSeeker>,Reader,Seeker{

	}

	export interface WriterTo extends Proto<WriterTo>{
		writeTo(w:Writer):[int64,error]

	}
	export function newSectionReader(r:ReaderAt,off:int64,n:int64):Ref<SectionReader>
	export function multiReader(... readers:Reader[]):Reader

	export interface WriteCloser extends Proto<WriteCloser>,Writer,Closer{

	}

	export interface ByteWriter extends Proto<ByteWriter>{
		writeByte(c:byte):error

	}
	export function readAtLeast(r:Reader,buf:Uint8Array,min:int):[int,error]
	export function readAll(r:Reader):[Uint8Array,error]
	export function copyN(dst:Writer,src:Reader,n:int64):[int64,error]

	export interface Reader extends Proto<Reader>{
		read(p:Uint8Array):[int,error]

	}

	export interface Writer extends Proto<Writer>{
		write(p:Uint8Array):[int,error]

	}

	export interface Seeker extends Proto<Seeker>{
		seek(offset:int64,whence:int):[int64,error]

	}

	export interface ReaderAt extends Proto<ReaderAt>{
		readAt(p:Uint8Array,off:int64):[int,error]

	}

}