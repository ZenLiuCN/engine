package engine

import (
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"os"
	"strings"
)

type Code struct {
	*goja.Program
}

// CompileFile entry for whether source is as script or library
func CompileFile(path string, entry bool) *Code {
	if len(path) > 3 && (strings.EqualFold(path[len(path)-3:], ".js") || strings.EqualFold(path[len(path)-3:], ".cjs") || strings.EqualFold(path[len(path)-3:], ".mjs")) {
		data := string(fn.Panic1(os.ReadFile(path)))
		return &Code{fn.Panic1(goja.Compile(path, CompileJs(data, entry), true))}
	}
	if len(path) > 3 && strings.EqualFold(path[len(path)-3:], ".ts") {
		data := string(fn.Panic1(os.ReadFile(path)))
		return &Code{fn.Panic1(goja.Compile(path, CompileTs(data, entry), true))}
	}
	panic(fmt.Errorf(`unsupported file %s`, path))
}

// CompileSource compile script source, ts for source is typescript or javascript. entry for whether source is as script or library
func CompileSource(source string, ts, entry bool) *Code {
	if !ts {
		return &Code{fn.Panic1(goja.Compile("", CompileJs(source, entry), false))}
	} else {
		return &Code{fn.Panic1(goja.Compile("", CompileTs(source, entry), false))}
	}
}
