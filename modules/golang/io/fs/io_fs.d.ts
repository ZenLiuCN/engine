// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/io/fs'{

	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as time from 'golang/time'
	// @ts-ignore
	import type {Alias,error,bool,GoError,int,int64,uint32,Struct,Ref} from 'go'
	export interface DirEntry{

			info():FileInfo
			isDir():bool
			name():string
			type():FileMode
	}
	export const ErrClosed:GoError
	export const ErrExist:GoError
	export const ErrInvalid:GoError
	export const ErrNotExist:GoError
	export const ErrPermission:GoError
	export interface FS{

			open(name:string):File
	}
	export interface File extends io.Closer{

			close():error
			read(v1:Uint8Array):int
			stat():FileInfo
	}
	export interface FileInfo{

			isDir():bool
			modTime():time.Time
			mode():FileMode
			name():string
			size():int64
			sys():any
	}
	export function fileInfoToDirEntry(info:FileInfo):DirEntry

	export interface FileMode extends uint32{

	string():string
	isDir():bool
	isRegular():bool
	perm():FileMode
	type():FileMode
	}
	export function toFileMode(n:uint32):FileMode
	export function formatDirEntry(dir:DirEntry):string

	export function formatFileInfo(info:FileInfo):string

	export function glob(fsys:FS,pattern:string):string[]

	export interface GlobFS extends FS{

			glob(pattern:string):string[]
	}
	export const ModeAppend:FileMode
	export const ModeCharDevice:FileMode
	export const ModeDevice:FileMode
	export const ModeDir:FileMode
	export const ModeExclusive:FileMode
	export const ModeIrregular:FileMode
	export const ModeNamedPipe:FileMode
	export const ModePerm:FileMode
	export const ModeSetgid:FileMode
	export const ModeSetuid:FileMode
	export const ModeSocket:FileMode
	export const ModeSticky:FileMode
	export const ModeSymlink:FileMode
	export const ModeTemporary:FileMode
	export const ModeType:FileMode
	export interface PathError extends Struct<PathError>,Error,GoError{

			op:string
			path:string
			err:GoError
			error():string
			unwrap()/*error*/
			timeout():bool
	}
	export function readDir(fsys:FS,name:string):DirEntry[]

	export interface ReadDirFS extends FS{

			readDir(name:string):DirEntry[]
	}
	export interface ReadDirFile extends File{

			readDir(n:int):DirEntry[]
	}
	export function readFile(fsys:FS,name:string):Uint8Array

	export interface ReadFileFS extends FS{

			readFile(name:string):Uint8Array
	}
	export const SkipAll:GoError
	export const SkipDir:GoError
	export function stat(fsys:FS,name:string):FileInfo

	export interface StatFS extends FS{

			stat(name:string):FileInfo
	}
	export function sub(fsys:FS,dir:string):FS

	export interface SubFS extends FS{

			sub(dir:string):FS
	}
	export function validPath(name:string):bool

	export function walkDir(fsys:FS,root:string,fn:WalkDirFunc)/*error*/

	export interface WalkDirFunc extends Alias<(path:string,d:DirEntry,err:error)=>error>{

	}
}