package engine

import (
	_ "embed"
	"github.com/dop251/goja"
	"slices"
)

//go:embed engine.d.ts
var engineDefine []byte

type EngineModule struct {
}

func (e EngineModule) Name() string {
	return "engine"
}

func (e EngineModule) Register(engine *Engine) {
	engine.RegisterType("Engine", func(v []goja.Value) any {
		var r *Engine
		if len(v) == 0 {
			r = Get()
		} else {
			ex := v[0].ToBoolean()
			ms := v[1].Export().([]string)
			if ex {
				m := Modules()
				var mm []Module
				for _, module := range m {
					if !slices.Contains(ms, module.Name()) {
						mm = append(mm, module)
					}
				}
				r = NewRawEngine(mm...)
			} else {
				m := Modules()
				var mm []Module
				for _, module := range m {
					if slices.Contains(ms, module.Name()) {
						mm = append(mm, module)
					}
				}
				r = NewRawEngine(mm...)
			}
		}
		return r
	})
}

func (e EngineModule) TypeDefine() []byte {
	return engineDefine
}
