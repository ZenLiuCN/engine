package engine

import (
	_ "embed"
	"github.com/evanw/esbuild/pkg/api"
)

var (
	//go:embed module_esbuild.d.ts
	esbuildDefine []byte
)

type EsBuild struct {
	m map[string]any
}

func (e *EsBuild) Identity() string {
	return "go/esbuild"
}

func (e *EsBuild) TypeDefine() []byte {
	return esbuildDefine
}

func (e *EsBuild) Exports() map[string]any {
	if e.m == nil {
		e.m = map[string]any{
			"analyzeMetafile": api.AnalyzeMetafile,
			"formatMessages":  api.FormatMessages,
			"transform":       api.Transform,
			"build":           api.Build,
			"context": Transform12(api.Context, func(bc api.BuildContext, be *api.ContextError) *ContextResult {
				return &ContextResult{
					Context: &Context{bc},
					Err:     be,
				}
			}),
		}
	}
	return e.m
}

type ContextResult struct {
	Context *Context
	Err     *api.ContextError
}
type Context struct {
	c api.BuildContext
}

func (c *Context) Rebuild() api.BuildResult {
	return c.c.Rebuild()
}

func (c *Context) Watch(options api.WatchOptions) *GoError {
	err := c.c.Watch(options)
	if err != nil {
		return &GoError{Err: err}
	}
	return nil
}

type ServeOut struct {
	Result api.ServeResult
	Err    *GoError
}

func (c *Context) Serve(options api.ServeOptions) *ServeOut {
	r, err := c.c.Serve(options)
	if err != nil {
		return &ServeOut{
			Result: r,
			Err:    &GoError{Err: err},
		}
	}
	return &ServeOut{
		Result: r,
		Err:    nil,
	}
}

func (c *Context) Cancel() {
	c.c.Cancel()
}

func (c *Context) Dispose() {
	c.c.Dispose()
}
