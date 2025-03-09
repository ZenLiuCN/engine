package engine

import (
	"bytes"
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
	compilerMap    = map[string]any{
		"compileJs": func(js string, entry bool) (s string, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err = v
					default:
						err = fmt.Errorf("%s", v)
					}
				}
			}()
			return CompileJs(js, entry), nil
		},
		"compileTs": func(js string, entry bool) (s string, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err = v
					default:
						err = fmt.Errorf("%s", v)
					}
				}
			}()
			return CompileTs(js, entry), nil
		},
		"compileJsWithMapping": func(js string, entry bool) (s string, m SourceMapping, b []byte, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err = v
					default:
						err = fmt.Errorf("%s", v)
					}
				}
			}()
			s, m, b = CompileJsWithMapping("vm.js", js, entry)
			return
		},
		"compileTsWithMapping": func(js string, entry bool) (s string, m SourceMapping, b []byte, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err = v
					default:
						err = fmt.Errorf("%s", v)
					}
				}
			}()
			s, m, b = CompileTsWithMapping("vm.ts", js, entry)
			return
		},
		"compileTsCode": func(src string, entry bool) (c *Code, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err = v
					default:
						err = fmt.Errorf("%s", v)
					}
				}
			}()
			return CompileSource(src, true, entry), nil
		},
		"compileJsCode": func(src string, entry bool) (c *Code, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err = v
					default:
						err = fmt.Errorf("%s", v)
					}
				}
			}()
			return CompileSource(src, false, entry), nil
		},
		"compileTsCodeWithMapping": func(src string, entry bool) (c *Code, m SourceMapping, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err = v
					default:
						err = fmt.Errorf("%s", v)
					}
				}
			}()
			c, m = CompileSourceWithMapping("vm.ts", src, true, entry)
			return
		},
		"compileJsCodeWithMapping": func(src string, entry bool) (c *Code, m SourceMapping, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err = v
					default:
						err = fmt.Errorf("%s", v)
					}
				}
			}()
			c, m = CompileSourceWithMapping("vm.js", src, false, entry)
			return
		},
	}
)

type Compiler struct {
}

func (s Compiler) TypeDefine() []byte {
	return compilerDefine
}

func (s Compiler) Identity() string {
	return "go/compiler"
}

func (s Compiler) Exports() map[string]any {

	return compilerMap
}

func CompileJs(js string, entry bool) string {
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
	if entry && format == api.FormatCommonJS {
		idx := bytes.Index(res.Code, []byte("module.exports=__toCommonJS(stdin_exports);"))
		if idx >= 0 {
			return string(res.Code[idx+43:])
		}
		return string(res.Code)
	}
	return string(res.Code)
}
func CompileTs(ts string, entry bool) string {
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
	if entry && format == api.FormatCommonJS {
		idx := bytes.Index(res.Code, []byte("module.exports=__toCommonJS(stdin_exports);"))
		if idx >= 0 {
			return string(res.Code[idx+43:])
		}
		return string(res.Code)
	}
	return string(res.Code)
}

func CompileJsWithMapping(name, js string, entry bool) (string, SourceMapping, []byte) {
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

		Target:     api.ES2022,
		Platform:   api.PlatformBrowser,
		Format:     format,
		Sourcemap:  api.SourceMapExternal,
		Loader:     api.LoaderJS,
		Sourcefile: name,
	})
	if res.Errors != nil {
		panic(errors.New("Compile JS error\n" + fn.SliceJoinRune(res.Errors, '\n', func(m api.Message) string {
			return fmt.Sprintf(`%s: %s`, m.Location.LineText, m.Text)
		})))
	}
	sm := NewSourceMap(res.Map)
	if entry && format == api.FormatCommonJS {
		idx := bytes.Index(res.Code, []byte("module.exports=__toCommonJS(stdin_exports);"))
		if idx >= 0 {
			return string(res.Code[idx+43:]), sm, res.Map
		}
		return string(res.Code), sm, res.Map
	}
	return string(res.Code), sm, res.Map
}
func CompileTsWithMapping(name, ts string, entry bool) (string, SourceMapping, []byte) {
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

		Target:     api.ES2022,
		Platform:   api.PlatformBrowser,
		Format:     format,
		Sourcemap:  api.SourceMapExternal,
		Loader:     api.LoaderTS,
		Sourcefile: name,
	})
	if res.Errors != nil {
		panic(errors.New("Compile TS error\n" + fn.SliceJoinRune(res.Errors, '\n', func(m api.Message) string {
			return fmt.Sprintf(`%s: %s`, m.Location.LineText, m.Text)
		})))
	}
	sm := NewSourceMap(res.Map)
	if entry && format == api.FormatCommonJS {
		idx := bytes.Index(res.Code, []byte("module.exports=__toCommonJS(stdin_exports);"))
		if idx >= 0 {
			return string(res.Code[idx+43:]), sm, res.Map
		}
		return string(res.Code), sm, res.Map
	}
	return string(res.Code), sm, res.Map
}
