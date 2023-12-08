package engine

import (
	"context"
	_ "embed"
	"time"
)

var (
	//go:embed module_context.d.ts
	contextDefine []byte
	contextMap    = map[string]any{
		"background": context.Background,
		"withCancel": func(p context.Context) ContextCancel {
			c, cc := context.WithCancel(p)
			return ContextCancel{Context: c, Cancel: cc}
		},
		"withTimeout": func(p context.Context, t time.Duration) ContextCancel {
			c, cc := context.WithTimeout(p, t)
			return ContextCancel{Context: c, Cancel: cc}
		},
	}
)

type ContextModule struct {
}

func (c ContextModule) TypeDefine() []byte {
	return contextDefine
}

func (c ContextModule) Identity() string {
	return "go/context"
}

func (c ContextModule) Exports() map[string]any {
	return contextMap
}

type ContextCancel struct {
	Context context.Context
	Cancel  context.CancelFunc
}
