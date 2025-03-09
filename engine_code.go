package engine

import (
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"os"
	"path/filepath"
	"strings"
)

type Code struct {
	Path string
	*goja.Program
}

// CompileFile entry for whether source is as script or library
func CompileFile(path string, entry bool) *Code {
	if len(path) > 3 && (strings.EqualFold(path[len(path)-3:], ".js") || strings.EqualFold(path[len(path)-3:], ".cjs") || strings.EqualFold(path[len(path)-3:], ".mjs")) {
		data := string(fn.Panic1(os.ReadFile(path)))
		return &Code{Path: path, Program: fn.Panic1(goja.Compile(path, CompileJs(data, entry), true))}
	}
	if len(path) > 3 && strings.EqualFold(path[len(path)-3:], ".ts") {
		data := string(fn.Panic1(os.ReadFile(path)))
		return &Code{Path: path, Program: fn.Panic1(goja.Compile(path, CompileTs(data, entry), true))}
	}
	panic(fmt.Errorf(`unsupported file %s`, path))
}

// CompileSource compile script source, ts for source is typescript or javascript. entry for whether source is as script or library
func CompileSource(source string, ts, entry bool) *Code {
	if !ts {
		return &Code{Path: "", Program: fn.Panic1(goja.Compile("", CompileJs(source, entry), false))}
	} else {
		return &Code{Path: "", Program: fn.Panic1(goja.Compile("", CompileTs(source, entry), false))}
	}
}
func CompileFileSource(path, source string, ts, entry bool) *Code {
	if !ts {
		return &Code{Path: path, Program: fn.Panic1(goja.Compile("", CompileJs(source, entry), false))}
	} else {
		return &Code{Path: path, Program: fn.Panic1(goja.Compile("", CompileTs(source, entry), false))}
	}
}

func CompileFileWithMapping(path string, entry bool) (*Code, *SourceMap, string) {
	if len(path) > 3 && (strings.EqualFold(path[len(path)-3:], ".js") || strings.EqualFold(path[len(path)-3:], ".cjs") || strings.EqualFold(path[len(path)-3:], ".mjs")) {
		data := string(fn.Panic1(os.ReadFile(path)))
		base := filepath.Base(path)
		s, m, b := CompileJsWithMapping(base, data, entry)
		return &Code{Path: path, Program: compile(base, s, b)}, m, base
	}
	if len(path) > 3 && strings.EqualFold(path[len(path)-3:], ".ts") {
		data := string(fn.Panic1(os.ReadFile(path)))
		base := filepath.Base(path)
		s, m, b := CompileTsWithMapping(base, data, entry)
		return &Code{Path: path, Program: compile(base, s, b)}, m, base
	}
	panic(fmt.Errorf(`unsupported file %s`, path))
}

func CompileSourceWithMapping(file, source string, ts, entry bool) (*Code, *SourceMap) {
	if !ts {
		s, m, b := CompileJsWithMapping(file, source, entry)
		return &Code{Path: "", Program: compile(file, s, b)}, m
	} else {
		s, m, b := CompileTsWithMapping(file, source, entry)
		return &Code{Path: "", Program: compile(file, s, b)}, m
	}
}
func CompileFileSourceWithMapping(path, source string, ts, entry bool) (*Code, *SourceMap) {
	name := filepath.Base(path)
	if !ts {
		s, m, b := CompileJsWithMapping(name, source, entry)
		return &Code{Path: path, Program: compile(name, s, b)}, m
	} else {
		s, m, b := CompileTsWithMapping(name, source, entry)
		return &Code{Path: path, Program: compile(name, s, b)}, m
	}
}
func compile(name, source string, m []byte) *goja.Program {
	p := fn.Panic1(goja.Parse(name, source, parser.WithSourceMapLoader(func(path string) ([]byte, error) {
		return m, nil
	})))
	return fn.Panic1(goja.CompileAST(p, false))
}
