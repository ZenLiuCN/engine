package os

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"os"
	"path/filepath"
	"strings"
)

var (
	execPath   string
	root       string
	executable string
	ext        string
	name       string

	pathSeparator     = string(os.PathSeparator)
	pathListSeparator = string(os.PathListSeparator)
)

func init() {
	var err error
	execPath, err = os.Executable()
	if err != nil {
		execPath, _ = filepath.Abs(os.Args[0])
	}
	execPath, err = filepath.EvalSymlinks(execPath)
	if err != nil {
		execPath, _ = filepath.Abs(os.Args[0])
	}
	root = filepath.Dir(execPath)
	_, executable = filepath.Split(execPath)
	ext = filepath.Ext(execPath)
	name = strings.Replace(executable, ext, "", -1)

	engine.Register(&Os{})
}

var (
	//go:embed os.d.ts
	d []byte
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
	return d
}

func (o *Os) Name() string {
	return "os"
}

func (o *Os) Initialize(engine *engine.Engine) engine.Module {
	return &Os{
		Engine:            engine,
		Root:              root,
		SimpleName:        name,
		Ext:               ext,
		Executable:        executable,
		PathSeparator:     pathSeparator,
		PathListSeparator: pathListSeparator,
	}
}
func expand(path string) string {
	if strings.HasPrefix(path, "@") {
		return fixPath(os.ExpandEnv(filepath.Join(root, strings.TrimPrefix(path, "@"))))
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
		tar = tar + pathListSeparator + src
	}
	fn.Panic(os.Setenv(key, tar))
}
func (o *Os) Prep(key string, value ...string) {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, o.Expand)
	if len(src) != 0 {
		tar = tar + pathListSeparator + src
	}
	fn.Panic(os.Setenv(key, tar))
}
func (o *Os) Ap(key string, value ...string) {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, fn.Identity[string])
	if len(src) != 0 {
		tar = src + pathListSeparator + tar
	}
	fn.Panic(os.Setenv(key, tar))
}
func (o *Os) App(key string, value ...string) {
	src := os.Getenv(key)
	tar := fn.SliceJoinRune(value, os.PathListSeparator, o.Expand)
	if len(src) != 0 {
		tar = src + pathListSeparator + tar
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
func (o *Os) Read(path string) goja.ArrayBuffer {
	ar := o.Engine.NewArrayBuffer(fn.Panic1(os.ReadFile(o.Expand(path))))
	return ar
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
			"modified": info.ModTime(),
		})
	}
	return
}
