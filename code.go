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

func CompileFile(path string) *Code {
	if len(path) > 3 && (strings.EqualFold(path[len(path)-3:], ".js") || strings.EqualFold(path[len(path)-3:], ".cjs") || strings.EqualFold(path[len(path)-3:], ".mjs")) {
		data := string(fn.Panic1(os.ReadFile(path)))
		return &Code{fn.Panic1(goja.Compile(path, CompileJs(data), true))}
	}
	if len(path) > 3 && strings.EqualFold(path[len(path)-3:], ".ts") {
		data := string(fn.Panic1(os.ReadFile(path)))
		return &Code{fn.Panic1(goja.Compile(path, CompileTs(data), true))}
	}
	panic(fmt.Errorf(`unsupported file %s`, path))
}

func CompileSource(source string, ts bool) *Code {
	if !ts {
		return &Code{fn.Panic1(goja.Compile("", CompileJs(source), false))}
	} else {
		return &Code{fn.Panic1(goja.Compile("", CompileTs(source), false))}
	}
}
