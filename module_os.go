package engine

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	//go:embed module_os.d.ts
	osDefine []byte
)

type Os struct {
}

func (o Os) Identity() string {
	return "go/os"
}

func (o Os) Exports() map[string]any {
	return nil
}
func (o Os) TypeDefine() []byte {
	return osDefine
}
func (o Os) ExportsWithEngine(e *Engine) map[string]any {
	return map[string]any{
		"root":              root,
		"simpleName":        name,
		"ext":               ext,
		"executable":        executable,
		"pathSeparator":     pathSeparator,
		"pathListSeparator": pathListSeparator,
		"expand":            EnvExpand,
		"pre":               EnvPrepend,
		"prep":              EnvPrependPath,
		"ap":                EnvAppend,
		"app":               EnvAppendPath,
		"set":               EnvSet,
		"setp":              EnvSetPath,
		"put":               EnvPut,
		"putp":              EnvPutPath,
		"variable":          EnvVar,
		"evalFile": func(p string) any {
			return EvalFile(e, p)
		},
		"evalFiles": func(paths ...string) (r []any) {
			return EvalFiles(e, paths...)
		},
		"exec": func(option *ExecOption) error {
			return Execute(option)
		},
		"proc": func(option *ProcOption) (r *SubProc, err error) {
			s, err := OpenProc(option)
			if err != nil {
				return nil, err
			}
			return &SubProc{s: s}, nil
		},
		"mkdir":     Mkdir,
		"mkdirAll":  MkdirAll,
		"exists":    FileExists,
		"write":     WriteBinaryFile,
		"writeText": WriteTextFile,
		"read":      ReadBinaryFile,
		"readText":  ReadTextFile,
		"chdir":     Chdir,
		"pwd":       Pwd,
		"ls": func(path string) (r []map[string]any, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err = v
					default:
						err = fmt.Errorf("%s", v)
					}
				}
			}()
			if path == "" {
				path, _ = Pwd()
			}
			f := fn.Panic1(os.Open(EnvExpand(path)))
			defer fn.IgnoreClose(f)
			dir := fn.Panic1(f.ReadDir(0))
			for _, entry := range dir {
				info := fn.Panic1(entry.Info())
				r = append(r, map[string]any{
					"dir":      entry.IsDir(),
					"name":     entry.Name(),
					"mode":     entry.Type().String(),
					"size":     info.Size(),
					"modified": info.ModTime().Format("2006-01-02 15:04:05.000"),
				})
			}
			return
		},
		"stat": Stat,
		//std
		"open": func(path string) (f *os.File, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err = v
					default:
						err = fmt.Errorf("%s", v)
					}
				}
			}()
			return RegisterResource(e, fn.Panic1(os.Open(EnvExpand(path)))), nil
		},
		"O_CREATE": os.O_CREATE,
		"O_TRUNC":  os.O_TRUNC,
		"O_APPEND": os.O_APPEND,
		"O_RDONLY": os.O_RDONLY,
		"O_WRONLY": os.O_WRONLY,
		"O_SYNC":   os.O_SYNC,
		"O_EXCL":   os.O_EXCL,
		"flags": func(v ...int) (r int) {
			for _, i := range v {
				r |= i
			}
			return
		},
		"openFile": func(path string, flag int) (f *os.File, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err = v
					default:
						err = fmt.Errorf("%s", v)
					}
				}
			}()
			return RegisterResource(e, fn.Panic1(os.OpenFile(EnvExpand(path), flag, os.ModePerm))), nil
		},
		"chown": func(path string, uid, gid int) error {
			return os.Chown(EnvExpand(path), uid, gid)
		},
		"chmod": func(path string, mod os.FileMode) error {
			return os.Chmod(EnvExpand(path), mod)
		},
		"getGID":      os.Geteuid,
		"getUID":      os.Geteuid,
		"getPagesize": os.Getpagesize,
		"getPid":      os.Getpid,
		"mkdirALL": func(path string, perm os.FileMode) error {
			return os.MkdirAll(EnvExpand(path), perm)
		},
		"rename": func(path, new string) error {
			return os.Rename(EnvExpand(path), EnvExpand(new))
		},
		"userCacheDir":  os.UserCacheDir,
		"userConfigDir": os.UserConfigDir,
		"userHomeDir":   os.UserHomeDir,
		"writeFile": func(path string, data []byte, perm os.FileMode) error {
			return os.WriteFile(EnvExpand(path), data, perm)
		},
		"remove": func(path string) error {
			return os.Remove(EnvExpand(path))
		},
		"removeAll": func(path string) error {
			return os.RemoveAll(EnvExpand(path))
		},
		"hostname": os.Hostname,
		"tempDir":  os.TempDir,
	}

}

func EnvPutPath(key string, value ...string) {
	src := os.Getenv(key)
	if len(src) != 0 {
		return
	}
	fn.Panic(os.Setenv(key, fn.SliceJoinRune(value, os.PathListSeparator, EnvExpand)))
}
func EnvVar(key string) string {
	return os.Getenv(key)
}
func EvalFile(e *Engine, path string) any {
	return fn.Panic1(e.RunCode(CompileFile(EnvExpand(path), true))).Export()
}
func EvalFiles(e *Engine, paths ...string) (r []any) {
	for _, s := range paths {
		r = append(r, fn.Panic1(e.RunCode(CompileFile(EnvExpand(s), true))).Export())
	}
	return
}
func Mkdir(path string) {
	pt := EnvExpand(path)
	if FileExists(pt) {
		return
	}
	fn.Panic(os.Mkdir(pt, os.ModePerm))
}
func MkdirAll(path string) {
	pt := EnvExpand(path)
	if FileExists(pt) {
		return
	}
	fn.Panic(os.MkdirAll(pt, os.ModePerm))
}
func WriteBinaryFile(path string, data []byte) {
	fn.Panic(os.WriteFile(EnvExpand(path), data, os.ModePerm))
}
func WriteTextFile(path string, data string) {
	fn.Panic(os.WriteFile(EnvExpand(path), []byte(data), os.ModePerm))
}
func ReadBinaryFile(path string) []byte {
	return fn.Panic1(os.ReadFile(EnvExpand(path)))
}
func ReadTextFile(path string) string {
	return string(fn.Panic1(os.ReadFile(EnvExpand(path))))
}
func Chdir(path string) error {
	return os.Chdir(EnvExpand(path))
}
func EnvPut(key string, value ...string) error {
	src := os.Getenv(key)
	if len(src) != 0 {
		return nil
	}
	return os.Setenv(key, fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string]))
}
func EnvSetPath(key string, value ...string) error {
	tar := fn.SliceJoinRune(value, os.PathListSeparator, EnvExpand)
	return os.Setenv(key, tar)
}
func EnvSet(key string, value ...string) error {
	tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
	return os.Setenv(key, tar)
}
func EnvPrepend(key string, value ...string) error {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
	if len(src) != 0 {
		tar = tar + pathListSeparator + src
	}
	return os.Setenv(key, tar)
}
func EnvAppendPath(key string, value ...string) error {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, EnvExpand)
	if len(src) != 0 {
		tar = src + pathListSeparator + tar
	}
	return os.Setenv(key, tar)
}
func EnvAppend(key string, value ...string) error {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
	if len(src) != 0 {
		tar = src + pathListSeparator + tar
	}
	return os.Setenv(key, tar)
}
func EnvPrependPath(key string, value ...string) error {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, EnvExpand)
	if len(src) != 0 {
		tar = tar + pathListSeparator + src
	}
	return os.Setenv(key, tar)
}

func Pwd() (string, error) {
	return os.Getwd()
}
func FileExists(path string) bool {
	_, err := os.Stat(EnvExpand(path))
	return err == nil || !os.IsNotExist(err)
}
func EnvExpand(path string) string {
	if strings.HasPrefix(path, "@") {
		return fixPath(os.ExpandEnv(filepath.Join(root, strings.TrimPrefix(path, "@"))))
	}
	p, e := filepath.Abs(path)
	if e != nil {
		return fixPath(os.ExpandEnv(path))
	}
	return fixPath(os.ExpandEnv(p))
}
func Stat(path string) os.FileInfo {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return info
}

type SubProc struct {
	s   *SubProcess
	out io.ReadCloser
	err io.ReadCloser
	in  io.WriteCloser
	buf []byte
}

func (s SubProc) FromConsole(data []byte) []byte {
	buf := GetBytesBuffer()
	defer PutBytesBuffer(buf)
	defer buf.Reset()
	buf.Write(data)
	fn.Panic(FromNativeConsole(buf))
	return buf.Bytes()
}
func (s SubProc) ToConsole(data []byte) []byte {
	buf := GetBytesBuffer()
	defer PutBytesBuffer(buf)
	defer buf.Reset()
	buf.Write(data)
	fn.Panic(ToNativeConsole(buf))
	return buf.Bytes()
}
func (s SubProc) Run() string {
	err := s.s.Run()
	if err != nil {
		return err.Error()
	}
	return ""
}
func (s SubProc) Wait() string {
	err := s.s.Wait()
	if err != nil {
		return err.Error()
	}
	return ""
}
func (s SubProc) ReadStdout() BinaryError {
	var e error
	if s.out == nil {
		s.out, e = s.s.StderrPipe()
		if e != nil {
			return BinaryError{
				Data:  nil,
				Error: e.Error(),
			}
		}
	}
	if s.buf == nil {
		s.buf = make([]byte, 512)
	}
	n, e := s.out.Read(s.buf)
	if e != nil {
		return BinaryError{
			Data:  nil,
			Error: e.Error(),
		}
	}
	return BinaryError{
		Data:  s.buf[:n],
		Error: "",
	}
}
func (s SubProc) ReadStderr() BinaryError {
	var e error
	if s.err == nil {
		s.err, e = s.s.StderrPipe()
		if e != nil {
			return BinaryError{
				Data:  nil,
				Error: e.Error(),
			}
		}
	}
	if s.buf == nil {
		s.buf = make([]byte, 512)
	}
	n, e := s.err.Read(s.buf)
	if e != nil {
		return BinaryError{
			Data:  nil,
			Error: e.Error(),
		}
	}
	return BinaryError{
		Data:  s.buf[:n],
		Error: "",
	}
}
func (s SubProc) WriteStdin(data []byte) WriteError {
	var e error
	if s.in == nil {
		s.in, e = s.s.StdinPipe()
		if e != nil {
			return WriteError{
				Write: 0,
				Error: e.Error(),
			}
		}
	}
	if data == nil {
		return WriteError{
			Write: 0,
			Error: "",
		}
	}
	if s.buf == nil {
		s.buf = make([]byte, 512)
	}
	n, e := s.in.Write(data)
	if e != nil {
		return WriteError{
			Write: 0,
			Error: e.Error(),
		}
	}
	return WriteError{
		Write: n,
		Error: "",
	}
}
func (s SubProc) Start() string {
	err := s.s.Start()
	if err != nil {
		return err.Error()
	}
	return ""
}
func (s SubProc) Exited() bool {
	if s.s.ProcessState == nil {
		return false
	}
	return s.s.ProcessState.Exited()
}
func (s SubProc) Success() bool {
	if s.s.ProcessState == nil {
		return false
	}
	return s.s.ProcessState.Exited()
}
func (s SubProc) SysTime() int64 {
	if s.s.ProcessState == nil {
		return -1
	}
	return s.s.ProcessState.SystemTime().Milliseconds()
}
func (s SubProc) UserTime() int64 {
	if s.s.ProcessState == nil {
		return -1
	}
	return s.s.ProcessState.UserTime().Milliseconds()
}
func (s SubProc) SysNanoTime() int64 {
	if s.s.ProcessState == nil {
		return -1
	}
	return s.s.ProcessState.SystemTime().Nanoseconds()
}
func (s SubProc) UserNanoTime() int64 {
	if s.s.ProcessState == nil {
		return -1
	}
	return s.s.ProcessState.UserTime().Nanoseconds()
}
func (s SubProc) Kill() string {
	err := s.s.Process.Kill()
	if err != nil {
		return err.Error()
	}
	return ""
}
func (s SubProc) Release() string {
	s.Free()
	err := s.s.Process.Release()
	if err != nil {
		return err.Error()
	}
	return ""
}
func (s SubProc) Output() BinaryError {
	b, e := s.s.Output()
	if e != nil {
		return BinaryError{
			Data:  nil,
			Error: e.Error(),
		}
	}
	return BinaryError{
		Data:  b,
		Error: "",
	}
}
func (s SubProc) CombinedOutput() BinaryError {
	b, e := s.s.CombinedOutput()
	if e != nil {
		return BinaryError{
			Data:  nil,
			Error: e.Error(),
		}
	}
	return BinaryError{
		Data:  b,
		Error: "",
	}
}
func (s SubProc) Free() {
	if s.out != nil {
		_ = s.out.Close()
		s.out = nil
	}
	if s.in != nil {
		_ = s.in.Close()
		s.in = nil
	}
	if s.err != nil {
		_ = s.err.Close()
		s.err = nil
	}
	s.buf = nil
}

type BinaryError struct {
	Data  []byte
	Error string
}
type WriteError struct {
	Write int
	Error string
}
