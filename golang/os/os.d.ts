// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/os'{
	// @ts-ignore
	import * as fs from 'golang/io/fs'
	// @ts-ignore
	import * as time from 'golang/time'
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as syscall from 'golang/syscall'
	// @ts-ignore
	import type {rune,error,int,Ref,int64,uintptr,bool,uint8,GoError,Struct} from 'go'
	export const Args:string[]
	export function chdir(dir:string):error
	export function chmod(name:string,mode:fs.FileMode):error
	export function chown(name:string,uid:int,gid:int):error
	export function chtimes(name:string,atime:time.Time,mtime:time.Time):error
	export function clearenv():void
	export function create(name:string):Ref<File>
	export function createTemp(dir:string,pattern:string):[Ref<File>,error]
	//"NUL"
	export const DevNull:string
	export type DirEntry=fs.DirEntry
	export function dirFS(dir:string):fs.FS
	export function environ():string[]
	export const ErrClosed:GoError
	export const ErrDeadlineExceeded:GoError
	export const ErrExist:GoError
	export const ErrInvalid:GoError
	export const ErrNoDeadline:GoError
	export const ErrNotExist:GoError
	export const ErrPermission:GoError
	export const ErrProcessDone:GoError
	export function executable():string
	export function exit(code:int):void
	export function expand(s:string,mapping:(v1:string)=>string):string
	export function expandEnv(s:string):string
	export interface File extends io.ReaderAt,io.ReaderFrom,io.WriterAt,io.Closer,Struct<File>,fs.ReadDirFile,io.StringWriter,io.ReadWriteSeeker,io.ReadWriter,io.WriteCloser,io.ReadSeeker,io.ReadCloser,syscall.Conn,io.WriteSeeker,io.ReadSeekCloser,io.ReadWriteCloser{
		readdir(n:int):[fs.FileInfo[],error]
		readdirnames(n:int):[string[],error]
		readDir(n:int):[fs.DirEntry[],error]
		name():string
		read(b:Uint8Array):[int,error]
		readAt(b:Uint8Array,off:int64):[int,error]
		readFrom(r:io.Reader):[int64,error]
		write(b:Uint8Array):[int,error]
		writeAt(b:Uint8Array,off:int64):[int,error]
		seek(offset:int64,whence:int):[int64,error]
		writeString(s:string):[int,error]
		chmod(mode:fs.FileMode):error
		setDeadline(t:time.Time):error
		setReadDeadline(t:time.Time):error
		setWriteDeadline(t:time.Time):error
		syscallConn():[syscall.RawConn,error]
		close():error
		chown(uid:int,gid:int):error
		truncate(size:int64):error
		sync():error
		chdir():error
		fd():uintptr
		stat():[fs.FileInfo,error]
	}
	export type FileInfo=fs.FileInfo
	export type FileMode=fs.FileMode
	export function findProcess(pid:int):Ref<Process>
	export function getegid():int
	export function getenv(key:string):string
	export function geteuid():int
	export function getgid():int
	export function getgroups():[int]
	export function getpagesize():int
	export function getpid():int
	export function getppid():int
	export function getuid():int
	export function getwd():string
	export function hostname():string
	export const Interrupt:Signal
	export function isExist(err:error):bool
	export function isNotExist(err:error):bool
	export function isPathSeparator(c:uint8):bool
	export function isPermission(err:error):bool
	export function isTimeout(err:error):bool
	export const Kill:Signal
	export function lchown(name:string,uid:int,gid:int):error
	export function link(oldname:string,newname:string):error
	export interface LinkError extends Struct<LinkError>,Error,GoError{
		op:string
		old:string
		New:string
		err:GoError
		error():string
		unwrap():error
	}
	export function lookupEnv(key:string):[string,bool]
	export function lstat(name:string):fs.FileInfo
	export function mkdir(name:string,perm:fs.FileMode):error
	export function mkdirAll(path:string,perm:fs.FileMode):error
	export function mkdirTemp(dir:string,pattern:string):[string,error]
	export const ModeAppend:fs.FileMode
	export const ModeCharDevice:fs.FileMode
	export const ModeDevice:fs.FileMode
	export const ModeDir:fs.FileMode
	export const ModeExclusive:fs.FileMode
	export const ModeIrregular:fs.FileMode
	export const ModeNamedPipe:fs.FileMode
	export const ModePerm:fs.FileMode
	export const ModeSetgid:fs.FileMode
	export const ModeSetuid:fs.FileMode
	export const ModeSocket:fs.FileMode
	export const ModeSticky:fs.FileMode
	export const ModeSymlink:fs.FileMode
	export const ModeTemporary:fs.FileMode
	export const ModeType:fs.FileMode
	export function newFile(fd:uintptr,name:string):Ref<File>
	export function newSyscallError(syscall:string,err:error):error
	//1024
	export const O_APPEND:int
	//64
	export const O_CREATE:int
	//128
	export const O_EXCL:int
	//0
	export const O_RDONLY:int
	//2
	export const O_RDWR:int
	//4096
	export const O_SYNC:int
	//512
	export const O_TRUNC:int
	//1
	export const O_WRONLY:int
	export function open(name:string):Ref<File>
	export function openFile(name:string,flag:int,perm:fs.FileMode):[Ref<File>,error]
	export type PathError=fs.PathError
	//59
	export const PathListSeparator:rune
	//92
	export const PathSeparator:rune
	export function pipe():[Ref<File>,Ref<File>,error]
	export interface ProcAttr extends Struct<ProcAttr>{
		dir:string
		env:string[]
		files:Ref<File>[]
		sys:Ref<syscall.SysProcAttr>
	}
	export interface Process extends Struct<Process>{
		pid:int
		release():error
		kill():error
		wait():[Ref<ProcessState>,error]
		signal(sig:Signal):error
	}
	export interface ProcessState extends Struct<ProcessState>{
		userTime():time.Duration
		systemTime():time.Duration
		exited():bool
		success():bool
		sys():any
		sysUsage():any
		pid():int
		string():string
		exitCode():int
	}
	export function readDir(name:string):[fs.DirEntry]
	export function readFile(name:string):Uint8Array
	export function readlink(name:string):string
	export function remove(name:string):error
	export function removeAll(path:string):error
	export function rename(oldpath:string,newpath:string):error
	//1
	export const SEEK_CUR:int
	//2
	export const SEEK_END:int
	//0
	export const SEEK_SET:int
	export function sameFile(fi1:fs.FileInfo,fi2:fs.FileInfo):bool
	export function setenv(key:string,value:string):error
	export interface Signal{
		signal():void
		string():string
	}
	export function startProcess(name:string,argv:string[],attr:Ref<ProcAttr>):[Ref<Process>,error]
	export function stat(name:string):fs.FileInfo
	export const Stderr:Ref<File>
	export const Stdin:Ref<File>
	export const Stdout:Ref<File>
	export function symlink(oldname:string,newname:string):error
	export interface SyscallError extends Error,GoError,Struct<SyscallError>{
		syscall:string
		err:GoError
		error():string
		unwrap():error
		timeout():bool
	}
	export function tempDir():string
	export function truncate(name:string,size:int64):error
	export function unsetenv(key:string):error
	export function userCacheDir():string
	export function userConfigDir():string
	export function userHomeDir():string
	export function writeFile(name:string,data:Uint8Array,perm:fs.FileMode):error

export function emptyFile():File
export function refFile():Ref<File>
export function refOfFile(x:File):Ref<File>
export function emptyProcAttr():ProcAttr
export function refProcAttr():Ref<ProcAttr>
export function refOfProcAttr(x:ProcAttr):Ref<ProcAttr>
export function emptyProcess():Process
export function refProcess():Ref<Process>
export function refOfProcess(x:Process):Ref<Process>
export function emptyProcessState():ProcessState
export function refProcessState():Ref<ProcessState>
export function refOfProcessState(x:ProcessState):Ref<ProcessState>}
