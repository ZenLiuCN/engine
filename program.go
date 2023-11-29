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
		return &Code{fn.Panic1(goja.CompileAST(fn.Panic1(goja.Parse(path, CompileJs(data))), false))}
	}
	if len(path) > 3 && strings.EqualFold(path[len(path)-3:], ".ts") {
		data := string(fn.Panic1(os.ReadFile(path)))
		return &Code{fn.Panic1(goja.CompileAST(fn.Panic1(goja.Parse(path, CompileTs(data))), false))}
	}
	panic(fmt.Errorf(`unsupported file %s`, path))
}

func CompileSource(source string, ts bool) *Code {
	if !ts {
		return &Code{fn.Panic1(goja.CompileAST(fn.Panic1(goja.Parse("", CompileJs(source))), false))}
	} else {
		return &Code{fn.Panic1(goja.CompileAST(fn.Panic1(goja.Parse("", CompileTs(source))), false))}
	}
}
