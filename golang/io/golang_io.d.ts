declare module "golang/io" {
	// @ts-ignore
	import type {int64,error,int,Struct,Ref,byte,rune,Proto} from 'go'



	export interface ReaderFrom extends Proto<ReaderFrom>{
		readFrom(r:Reader):[int64,error]

	}

	export interface WriterAt extends Proto<WriterAt>{
		writeAt(p:Uint8Array,off:int64):[int,error]

	}
	export function readAll(r:Reader):[Uint8Array,error]
	export const SeekEnd=2
	export const SeekStart=0
	export const SeekCurrent=1

	export interface ReadWriteCloser extends Proto<ReadWriteCloser>,Reader,Writer,Closer{

	}

	export interface WriteSeeker extends Proto<WriteSeeker>,Writer,Seeker{

	}
	export interface SectionReader extends Struct<SectionReader>{
		size():int64
		read(p:Uint8Array):[int,error]
		seek(offset:int64,whence:int):[int64,error]
		readAt(p:Uint8Array,off:int64):[int,error]

	}
	export interface OffsetWriter extends Struct<OffsetWriter>{
		write(p:Uint8Array):[int,error]
		writeAt(p:Uint8Array,off:int64):[int,error]
		seek(offset:int64,whence:int):[int64,error]

	}
	export function teeReader(r:Reader,w:Writer):Reader
	export interface PipeReader extends Struct<PipeReader>,Closer{
		closeWithError(err:error):error
		read(data:Uint8Array):[int,error]
		close():error

	}
	export function limitReader(r:Reader,n:int64):Reader
	export function newOffsetWriter(w:WriterAt,off:int64):Ref<OffsetWriter>

	export interface Reader extends Proto<Reader>{
		read(p:Uint8Array):[int,error]

	}

	export interface ReadSeekCloser extends Proto<ReadSeekCloser>,Reader,Seeker,Closer{

	}

	export interface ReadWriteSeeker extends Proto<ReadWriteSeeker>,Reader,Writer,Seeker{

	}

	export interface ByteWriter extends Proto<ByteWriter>{
		writeByte(c:byte):error

	}

	export interface RuneReader extends Proto<RuneReader>{
		readRune():[rune,int,error]

	}
	export function writeString(w:Writer,s:string):[int,error]
	export function multiWriter(... writers:Writer[]):Writer

	export interface Seeker extends Proto<Seeker>{
		seek(offset:int64,whence:int):[int64,error]

	}

	export interface ReadWriter extends Proto<ReadWriter>,Reader,Writer{

	}

	export interface WriterTo extends Proto<WriterTo>{
		writeTo(w:Writer):[int64,error]

	}
	export function copyN(dst:Writer,src:Reader,n:int64):[int64,error]
	export function copyBuffer(dst:Writer,src:Reader,buf:Uint8Array):[int64,error]
	export function nopCloser(r:Reader):ReadCloser
	export function pipe():[Ref<PipeReader>,Ref<PipeWriter>]

	export interface Writer extends Proto<Writer>{
		write(p:Uint8Array):[int,error]

	}

	export interface WriteCloser extends Proto<WriteCloser>,Writer,Closer{

	}

	export interface ByteReader extends Proto<ByteReader>{
		readByte():[byte,error]

	}

	export interface RuneScanner extends Proto<RuneScanner>,RuneReader{
		unreadRune():error

	}
	export function newSectionReader(r:ReaderAt,off:int64,n:int64):Ref<SectionReader>
	export interface PipeWriter extends Struct<PipeWriter>,Closer{
		write(data:Uint8Array):[int,error]
		close():error
		closeWithError(err:error):error

	}

	export interface ReadSeeker extends Proto<ReadSeeker>,Reader,Seeker{

	}
	export function readAtLeast(r:Reader,buf:Uint8Array,min:int):[int,error]
	export function copy(dst:Writer,src:Reader):[int64,error]
	export function multiReader(... readers:Reader[]):Reader
	export interface LimitedReader extends Struct<LimitedReader>{
		r:Ref<Reader>


		n:int64
		read(p:Uint8Array):[int,error]

	}

	export interface Closer extends Proto<Closer>{
		close():error

	}

	export interface ReadCloser extends Proto<ReadCloser>,Reader,Closer{

	}

	export interface ReaderAt extends Proto<ReaderAt>{
		readAt(p:Uint8Array,off:int64):[int,error]

	}

	export interface ByteScanner extends Proto<ByteScanner>,ByteReader{
		unreadByte():error

	}

	export interface StringWriter extends Proto<StringWriter>{
		writeString(s:string):[int,error]

	}
	export function readFull(r:Reader,buf:Uint8Array):[int,error]

}