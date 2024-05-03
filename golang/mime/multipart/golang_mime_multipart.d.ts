declare module "golang/mime/multipart" {
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as textproto from 'golang/net/textproto'
	// @ts-ignore
	import type {int64,error,Proto,int,Ref,Struct} from 'go'


	export function newWriter(w:io.Writer):Ref<Writer>
	export interface Reader extends Struct<Reader>{
		nextRawPart():[Ref<Part>,error]
		readForm(maxMemory:int64):[Ref<Form>,error]
		nextPart():[Ref<Part>,error]

	}
	export interface Form extends Struct<Form>{
		value:Record<string,string[]>

		file:Record<string,Ref<FileHeader>>
		removeAll():error

	}
	export interface FileHeader extends Struct<FileHeader>{
		filename:string

		header:textproto.MIMEHeader

		size:int64




		open():[File,error]

	}

	export interface File extends Proto<File>,io.Reader,io.ReaderAt,io.Seeker,io.Closer{

	}
	export interface Part extends Struct<Part>,io.Closer{
		header:textproto.MIMEHeader








		read(d:Uint8Array):[int,error]
		close():error
		formName():string
		fileName():string
io
	}
	export function newReader(r:io.Reader,boundary:string):Ref<Reader>
	export interface Writer extends Struct<Writer>,io.Closer{
		formDataContentType():string
		setBoundary(boundary:string):error
		createFormFile(fieldname,filename:string):[io.Writer,error]
		createFormField(fieldname:string):[io.Writer,error]
		writeField(fieldname,value:string):error
		close():error
		boundary():string
		createPart(header:textproto.MIMEHeader):[io.Writer,error]
io
	}

}