package engine

import (
	_ "embed"
	"github.com/dop251/goja"
	"slices"
)

//go:embed module_engine.d.ts
var engineDefine []byte

type EngineModule struct {
}

func (e EngineModule) ExportsWithEngine(engine *Engine) map[string]any {
	m := make(map[string]any)
	m["Engine"] = engine.ToConstructor(func(v []goja.Value) any {
		var r *Engine
		if len(v) == 0 {
			r = Get()
		} else {
			ex := v[0].ToBoolean()
			ms := v[1].Export().([]string)
			if ex {
				m := Mods()
				var mm []Mod
				for _, module := range m {
					if !slices.Contains(ms, module.Name()) {
						mm = append(mm, module)
					}
				}
				r = NewRawEngine(mm...)
			} else {
				m := Mods()
				var mm []Mod
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
	return m
}

func (e EngineModule) Identity() string {
	return "go/engine"
}

func (e EngineModule) Exports() map[string]any {
	return nil
}

func (e EngineModule) TypeDefine() []byte {
	return engineDefine
}
