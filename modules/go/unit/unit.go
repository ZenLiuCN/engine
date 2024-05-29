package unit

import (
	_ "embed"
	"fmt"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"io"
	"maps"
	"os"
	"strconv"
)

var (
	//go:embed module_unit.d.ts
	UnitDefine []byte
	UnitMap    = map[string]any{
		"root":              engine.GetRoot(),
		"simpleName":        engine.GetName(),
		"ext":               engine.GetExt(),
		"executable":        engine.GetExecutable(),
		"pathSeparator":     engine.GetPathSeparator(),
		"pathListSeparator": engine.GetPathListSeparator(),
		"expand":            engine.EnvExpand,
		"pre":               engine.EnvPrepend,
		"prep":              engine.EnvPrependPath,
		"ap":                engine.EnvAppend,
		"app":               engine.EnvAppendPath,
		"set":               engine.EnvSet,
		"setp":              engine.EnvSetPath,
		"put":               engine.EnvPut,
		"putp":              engine.EnvPutPath,
		"variable":          engine.EnvVar,
		"mkdir":             engine.Mkdir,
		"mkdirAll":          engine.MkdirAll,
		"exists":            engine.FileExists,
		"write":             engine.WriteBinaryFile,
		"writeText":         engine.WriteTextFile,
		"read":              engine.ReadBinaryFile,
		"readText":          engine.ReadTextFile,
		"chdir":             engine.Chdir,
		"pwd":               engine.Pwd,
		"getGID":            os.Geteuid,
		"getUID":            os.Geteuid,
		"getPagesize":       os.Getpagesize,
		"getPid":            os.Getpid,
		"userCacheDir":      os.UserCacheDir,
		"userConfigDir":     os.UserConfigDir,
		"userHomeDir":       os.UserHomeDir,
		"hostname":          os.Hostname,
		"tempDir":           os.TempDir,
		"cp": func(from, to string) (err error) {
			var f, f1 *os.File
			f, err = os.OpenFile(engine.EnvExpand(from), os.O_RDONLY, os.ModePerm)
			if err != nil {
				return
			}
			defer f.Close()
			f1, err = os.OpenFile(engine.EnvExpand(to), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
			if err != nil {
				return
			}
			defer f1.Close()
			_, err = io.Copy(f1, f)
			return
		},
		"mkdirALL": func(path string, perm string) error {
			return os.MkdirAll(engine.EnvExpand(path), os.FileMode(fn.Panic1(strconv.ParseInt(perm, 8, 32))))
		},
		"rename": func(path, new string) error {
			return os.Rename(engine.EnvExpand(path), engine.EnvExpand(new))
		},
		"remove": func(path string) error {
			return os.Remove(engine.EnvExpand(path))
		},
		"removeAll": func(path string) error {
			return os.RemoveAll(engine.EnvExpand(path))
		},
		"chown": func(path string, uid, gid int) error {
			return os.Chown(engine.EnvExpand(path), uid, gid)
		},
		"chmod": func(path string, mod string) error {
			return os.Chmod(engine.EnvExpand(path), os.FileMode(fn.Panic1(strconv.ParseInt(mod, 8, 32))))
		},
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
				path, _ = engine.Pwd()
			}
			f := fn.Panic1(os.Open(engine.EnvExpand(path)))
			defer fn.IgnoreClose(f)
			dir := fn.Panic1(f.ReadDir(0))
			for _, entry := range dir {
				info := fn.Panic1(entry.Info())
				r = append(r, map[string]any{
					//"dir":      entry.IsDir(),
					"name":     entry.Name(),
					"mode":     info.Mode().String(),
					"size":     info.Size(),
					"modified": info.ModTime().Format("2006-01-02 15:04:05.000Z07"),
				})
			}
			return
		},
		"stat": engine.Stat,
		"exec": engine.Execute,
		"proc": func(option *engine.ProcOption) (r *SubProc, err error) {
			s, err := engine.OpenProc(option)
			if err != nil {
				return nil, err
			}
			return &SubProc{s: s}, nil
		},
	}
)

func init() {
	engine.RegisterModule(Unit{})
}

type Unit struct {
}

func (o Unit) Identity() string {
	return "go/unit"
}
func (o Unit) TypeDefine() []byte {
	return UnitDefine
}
func (o Unit) Exports() map[string]any {
	return UnitMap
}
func (o Unit) ExportsWithEngine(e *engine.Engine) (m map[string]any) {
	m = maps.Clone(UnitMap)
	m["evalFile"] = func(p string) any {
		return engine.EvalFile(e, p)
	}
	m["evalFiles"] = func(paths ...string) (r []any) {
		return engine.EvalFiles(e, paths...)
	}
	return
}

type SubProc struct {
	s   *engine.SubProcess
	out io.ReadCloser
	err io.ReadCloser
	in  io.WriteCloser
	buf []byte
}

func (s SubProc) FromConsole(data []byte) []byte {
	buf := engine.GetBytesBuffer()
	defer engine.PutBytesBuffer(buf)
	defer buf.Reset()
	buf.Write(data)
	fn.Panic(engine.FromNativeConsole(buf))
	return buf.Bytes()
}
func (s SubProc) ToConsole(data []byte) []byte {
	buf := engine.GetBytesBuffer()
	defer engine.PutBytesBuffer(buf)
	defer buf.Reset()
	buf.Write(data)
	fn.Panic(engine.ToNativeConsole(buf))
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
