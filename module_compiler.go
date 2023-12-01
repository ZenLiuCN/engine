package engine

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/evanw/esbuild/pkg/api"
	"strings"
)

var (
	//go:embed module_compiler.d.ts
	compilerDefine []byte
)

type Compiler struct {
	m map[string]any
}

func (s *Compiler) TypeDefine() []byte {
	return compilerDefine
}

func (s *Compiler) Identity() string {
	return "go/compiler"
}

func (s *Compiler) Exports() map[string]any {
	if s.m == nil {
		s.m = map[string]any{}
		s.m["compileJs"] = CompileJs
		s.m["compileTs"] = CompileTs
		s.m["compileTsCode"] = func(src string) *Code {
			return CompileSource(src, true)
		}
		s.m["compileJsCode"] = func(src string) *Code {
			return CompileSource(src, false)
		}
	}
	return s.m
}

func CompileJs(js string) string {
	format := api.FormatDefault
	if strings.Contains(js, "import ") {
		format = api.FormatCommonJS
	} else if strings.Contains(js, "export ") {
		format = api.FormatCommonJS
	} else if strings.Contains(js, "require ") {
		format = api.FormatCommonJS
	}
	res := api.Transform(js, api.TransformOptions{
		MinifyWhitespace:  true,
		MinifyIdentifiers: false,
		MinifySyntax:      true,

		KeepNames:   true,
		TreeShaking: api.TreeShakingTrue,

		Target:   api.ES2022,
		Platform: api.PlatformBrowser,
		Format:   format,

		Loader: api.LoaderJS,
	})
	if res.Errors != nil {
		panic(errors.New("Compile JS error\n" + fn.SliceJoinRune(res.Errors, '\n', func(m api.Message) string {
			return fmt.Sprintf(`%s: %s`, m.Location.LineText, m.Text)
		})))
	}
	return string(res.Code)
}
func CompileTs(ts string) string {
	format := api.FormatDefault
	if strings.Contains(ts, "import ") {
		format = api.FormatCommonJS
	} else if strings.Contains(ts, "export ") {
		format = api.FormatCommonJS
	} else if strings.Contains(ts, "require ") {
		format = api.FormatCommonJS
	}
	res := api.Transform(ts, api.TransformOptions{
		MinifyWhitespace:  true,
		MinifyIdentifiers: false,
		MinifySyntax:      true,

		KeepNames:   true,
		TreeShaking: api.TreeShakingTrue,

		Target:   api.ES2022,
		Platform: api.PlatformBrowser,
		Format:   format,

		Loader: api.LoaderTS,
	})
	if res.Errors != nil {
		panic(errors.New("Compile TS error\n" + fn.SliceJoinRune(res.Errors, '\n', func(m api.Message) string {
			return fmt.Sprintf(`%s: %s`, m.Location.LineText, m.Text)
		})))
	}
	return string(res.Code)
}