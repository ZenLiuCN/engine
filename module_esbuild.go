package engine

import (
	_ "embed"
	"github.com/evanw/esbuild/pkg/api"
)

var (
	//go:embed module_esbuild.d.ts
	esbuildDefine []byte
	esbuildMap    = map[string]any{
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
)

type EsBuild struct {
}

func (e EsBuild) Identity() string {
	return "go/esbuild"
}

func (e EsBuild) TypeDefine() []byte {
	return esbuildDefine
}

func (e EsBuild) Exports() map[string]any {

	return esbuildMap
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

func (c *Context) Watch(options api.WatchOptions) error {
	return c.c.Watch(options)
}

func (c *Context) Serve(options api.ServeOptions) (api.ServeResult, error) {
	return c.c.Serve(options)
}

func (c *Context) Cancel() {
	c.c.Cancel()
}

func (c *Context) Dispose() {
	c.c.Dispose()
}
