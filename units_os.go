package engine

import (
	"errors"
	"github.com/ZenLiuCN/fn"
	"os"
	"path/filepath"
	"strings"
)

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
