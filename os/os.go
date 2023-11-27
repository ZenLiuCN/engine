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

	engine.Register(&Os{})
}

var (
	//go:embed os.d.ts
	osDefine []byte
)

type Os struct {
	*engine.Engine
	Root              string
	SimpleName        string `js:"name"`
	Ext               string
	Executable        string
	PathSeparator     string
	PathListSeparator string
}

func (o *Os) TypeDefine() []byte {
	return osDefine
}

func (o *Os) Name() string {
	return "os"
}

func (o *Os) Initialize(e *engine.Engine) engine.Module {
	return &Os{
		Engine:            e,
		Root:              engine.GetRoot(),
		SimpleName:        engine.GetName(),
		Ext:               engine.GetExt(),
		Executable:        engine.GetExecutable(),
		PathSeparator:     engine.GetPathSeparator(),
		PathListSeparator: engine.GetPathListSeparator(),
	}
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
func (o *Os) Expand(path string) string {
	return expand(path)
}
func (o *Os) Pre(key string, value ...string) {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
	if len(src) != 0 {
		tar = tar + o.PathListSeparator + src
	}
	fn.Panic(os.Setenv(key, tar))
}
func (o *Os) Prep(key string, value ...string) {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, o.Expand)
	if len(src) != 0 {
		tar = tar + o.PathListSeparator + src
	}
	fn.Panic(os.Setenv(key, tar))
}
func (o *Os) Ap(key string, value ...string) {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
	if len(src) != 0 {
		tar = src + o.PathListSeparator + tar
	}
	fn.Panic(os.Setenv(key, tar))
}
func (o *Os) App(key string, value ...string) {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, o.Expand)
	if len(src) != 0 {
		tar = src + o.PathListSeparator + tar
	}
	fn.Panic(os.Setenv(key, tar))
}
func (o *Os) Set(key string, value ...string) {
	tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
	fn.Panic(os.Setenv(key, tar))
}
func (o *Os) Setp(key string, value ...string) {
	tar := fn.SliceJoinRune(value, os.PathListSeparator, o.Expand)
	fn.Panic(os.Setenv(key, tar))
}
func (o *Os) Put(key string, value ...string) {
	src := os.Getenv(key)
	if len(src) != 0 {
		return
	}
	fn.Panic(os.Setenv(key, fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])))
}
func (o *Os) Putp(key string, value ...string) {
	src := os.Getenv(key)
	if len(src) != 0 {
		return
	}
	fn.Panic(os.Setenv(key, fn.SliceJoinRune(value, os.PathListSeparator, o.Expand)))
}
func (o *Os) Var(key string) string {
	return os.Getenv(key)
}
func (o *Os) EvalFile(path string) goja.Value {
	return fn.Panic1(o.Engine.Execute(engine.CompileFile(o.Expand(path))))
}
func (o *Os) EvalFiles(paths ...string) (r []goja.Value) {
	for _, s := range paths {
		r = append(r, fn.Panic1(o.Engine.Execute(engine.CompileFile(o.Expand(s)))))
	}
	return
}
func (o *Os) Exec(option *ExecOption) {
	fn.Panic(Execute(option))
}
func (o *Os) Proc(option *ProcOption) SubProc {
	return SubProc{s: OpenProc(option)}
}
func (o *Os) Mkdir(path string) {
	pt := o.Expand(path)
	if o.Exists(pt) {
		return
	}
	fn.Panic(os.Mkdir(pt, os.ModePerm))
}
func (o *Os) MkdirAll(path string) {
	pt := o.Expand(path)
	if o.Exists(pt) {
		return
	}
	fn.Panic(os.MkdirAll(pt, os.ModePerm))
}
func (o *Os) Exists(path string) bool {
	_, err := os.Stat(o.Expand(path))
	return err == nil || !os.IsNotExist(err)
}
func (o *Os) Write(path string, data goja.ArrayBuffer) {
	fn.Panic(os.WriteFile(o.Expand(path), data.Bytes(), os.ModePerm))
}
func (o *Os) WriteText(path string, data string) {
	fn.Panic(os.WriteFile(o.Expand(path), []byte(data), os.ModePerm))
}
func (o *Os) Read(path string) []byte {
	return fn.Panic1(os.ReadFile(o.Expand(path)))
}
func (o *Os) ReadText(path string) string {
	return string(fn.Panic1(os.ReadFile(o.Expand(path))))
}

func (o *Os) Chdir(path string) {
	fn.Panic(os.Chdir(o.Expand(path)))
}
func (o *Os) Pwd() string {
	return fn.Panic1(os.Getwd())
}
func (o *Os) Ls(path string) (r []map[string]any) {
	if path == "" {
		path = o.Pwd()
	}
	f := fn.Panic1(os.Open(o.Expand(path)))
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
