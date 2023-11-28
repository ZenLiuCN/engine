package os

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	engine.RegisterModule(&Os{})
}

var (
	//go:embed os.d.ts
	osDefine []byte
)

type Os struct {
}

func (o *Os) Identity() string {
	return "go/os"
}

func (o *Os) Exports() map[string]any {
	return nil
}

func (o *Os) ExportsWithEngine(e *engine.Engine) map[string]any {
	m := map[string]any{}
	m["root"] = engine.GetRoot()
	m["simpleName"] = engine.GetName()
	m["ext"] = engine.GetExt()
	m["executable"] = engine.GetExecutable()
	m["pathSeparator"] = engine.GetPathSeparator()
	m["pathListSeparator"] = engine.GetPathListSeparator()

	m["expand"] = func(path string) string {
		return expand(path)
	}
	m["pre"] = func(key string, value ...string) {
		src := os.Getenv(key)
		tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
		if len(src) != 0 {
			tar = tar + engine.GetPathListSeparator() + src
		}
		fn.Panic(os.Setenv(key, tar))
	}
	m["prep"] = func(key string, value ...string) {
		src := os.Getenv(key)
		tar := fn.SliceJoinRune(value, os.PathListSeparator, expand)
		if len(src) != 0 {
			tar = tar + engine.GetPathListSeparator() + src
		}
		fn.Panic(os.Setenv(key, tar))
	}
	m["ap"] = func(key string, value ...string) {
		src := os.Getenv(key)
		tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
		if len(src) != 0 {
			tar = src + engine.GetPathListSeparator() + tar
		}
		fn.Panic(os.Setenv(key, tar))
	}
	m["app"] = func(key string, value ...string) {
		src := os.Getenv(key)
		tar := fn.SliceJoinRune(value, os.PathListSeparator, expand)
		if len(src) != 0 {
			tar = src + engine.GetPathListSeparator() + tar
		}
		fn.Panic(os.Setenv(key, tar))
	}
	m["set"] = func(key string, value ...string) {
		tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
		fn.Panic(os.Setenv(key, tar))
	}
	m["setp"] = func(key string, value ...string) {
		tar := fn.SliceJoinRune(value, os.PathListSeparator, expand)
		fn.Panic(os.Setenv(key, tar))
	}
	m["put"] = func(key string, value ...string) {
		src := os.Getenv(key)
		if len(src) != 0 {
			return
		}
		fn.Panic(os.Setenv(key, fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])))
	}
	m["putp"] = func(key string, value ...string) {
		src := os.Getenv(key)
		if len(src) != 0 {
			return
		}
		fn.Panic(os.Setenv(key, fn.SliceJoinRune(value, os.PathListSeparator, expand)))
	}
	m["variable"] = func(key string) string {
		return os.Getenv(key)
	}
	m["evalFile"] = func(path string) any {
		return fn.Panic1(e.Execute(engine.CompileFile(expand(path)))).Export()
	}
	m["evalFiles"] = func(paths ...string) (r []any) {
		for _, s := range paths {
			r = append(r, fn.Panic1(e.Execute(engine.CompileFile(expand(s)))).Export())
		}
		return
	}
	m["exec"] = func(option *ExecOption) {
		fn.Panic(Execute(option))
	}
	m["proc"] = func(option *ProcOption) SubProc {
		return SubProc{s: OpenProc(option)}
	}
	m["mkdir"] = func(path string) {
		pt := expand(path)
		if exists(pt) {
			return
		}
		fn.Panic(os.Mkdir(pt, os.ModePerm))
	}
	m["mkdirAll"] = func(path string) {
		pt := expand(path)
		if exists(pt) {
			return
		}
		fn.Panic(os.MkdirAll(pt, os.ModePerm))
	}
	m["exists"] = exists
	m["write"] = func(path string, data goja.ArrayBuffer) {
		fn.Panic(os.WriteFile(expand(path), data.Bytes(), os.ModePerm))
	}
	m["writeText"] = func(path string, data string) {
		fn.Panic(os.WriteFile(expand(path), []byte(data), os.ModePerm))
	}
	m["read"] = func(path string) []byte {
		return fn.Panic1(os.ReadFile(expand(path)))
	}
	m["readText"] = func(path string) string {
		return string(fn.Panic1(os.ReadFile(expand(path))))
	}

	m["chdir"] = func(path string) {
		fn.Panic(os.Chdir(expand(path)))
	}
	m["pwd"] = pwd
	m["ls"] = func(path string) (r []map[string]any) {
		if path == "" {
			path = pwd()
		}
		f := fn.Panic1(os.Open(expand(path)))
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
	}

	return m
}

func (o *Os) TypeDefine() []byte {
	return osDefine
}
func pwd() string {
	return fn.Panic1(os.Getwd())
}
func exists(path string) bool {
	_, err := os.Stat(expand(path))
	return err == nil || !os.IsNotExist(err)
}
func expand(path string) string {
	if strings.HasPrefix(path, "@") {
		return fixPath(os.ExpandEnv(filepath.Join(engine.GetRoot(), strings.TrimPrefix(path, "@"))))
	}
	p, e := filepath.Abs(path)
	if e != nil {
		return fixPath(os.ExpandEnv(path))
	}
	return fixPath(os.ExpandEnv(p))
}

type SubProc struct {
	s   *SubProcess
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
	fn.Panic(FromNativeConsole(buf))
	return buf.Bytes()
}
func (s SubProc) ToConsole(data []byte) []byte {
	buf := engine.GetBytesBuffer()
	defer engine.PutBytesBuffer(buf)
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
