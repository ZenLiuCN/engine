package engine

import (
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"os"
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

func CompileFileWithMapping(path string, entry bool) (*Code, SourceMapping) {
	if len(path) > 3 && (strings.EqualFold(path[len(path)-3:], ".js") || strings.EqualFold(path[len(path)-3:], ".cjs") || strings.EqualFold(path[len(path)-3:], ".mjs")) {
		data := string(fn.Panic1(os.ReadFile(path)))
		s, m := CompileJsWithMapping(data, entry)
		return &Code{Path: path, Program: fn.Panic1(goja.Compile(path, s, true))}, m
	}
	if len(path) > 3 && strings.EqualFold(path[len(path)-3:], ".ts") {
		data := string(fn.Panic1(os.ReadFile(path)))
		s, m := CompileTsWithMapping(data, entry)
		return &Code{Path: path, Program: fn.Panic1(goja.Compile(path, s, true))}, m
	}
	panic(fmt.Errorf(`unsupported file %s`, path))
}

func CompileSourceWithMapping(source string, ts, entry bool) (*Code, SourceMapping) {
	if !ts {
		s, m := CompileJsWithMapping(source, entry)
		return &Code{Path: "", Program: fn.Panic1(goja.Compile("", s, false))}, m
	} else {
		s, m := CompileTsWithMapping(source, entry)
		return &Code{Path: "", Program: fn.Panic1(goja.Compile("", s, false))}, m
	}
}
func CompileFileSourceWithMapping(path, source string, ts, entry bool) (*Code, SourceMapping) {
	if !ts {
		s, m := CompileJsWithMapping(source, entry)
		return &Code{Path: path, Program: fn.Panic1(goja.Compile(path, s, false))}, m
	} else {
		s, m := CompileTsWithMapping(source, entry)
		return &Code{Path: path, Program: fn.Panic1(goja.Compile(path, s, false))}, m
	}
}
