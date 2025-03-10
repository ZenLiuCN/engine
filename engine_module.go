package engine

import (
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"os"
	"path/filepath"
)

// Module support for import by require, not a global instance as Module
type (
	Module interface {
		TypeDefined
		Identity() string // Identity the full module identity
		Exports() map[string]any
	}
	InitializeModule interface {
		Module
		ExportsWithEngine(eng *Engine) map[string]any
	}
)
type BaseInitializeModule struct {
}

func (s *BaseInitializeModule) Exports() map[string]any {
	panic("not a static Module")
}

var (
	goRegistry = map[string]Module{}
)

// RegisterModule as import able module
func RegisterModule(module Module) bool {
	if _, ok := goRegistry[module.Identity()]; ok {
		return false
	}
	goRegistry[module.Identity()] = module
	return true
}
func RemoveModule(module string) {
	delete(goRegistry, module)
}
func ModuleNames() []string {
	return fn.MapKeys(goRegistry)
}

// ModuleDefines exports Module define as moduleName=>ModuleTypeDefine
func ModuleDefines() map[string][]byte {
	m := make(map[string][]byte)
	for s, module := range goRegistry {
		m[s] = module.TypeDefine()
	}
	return m
}
func resolveModule(specifier string) Module {
	if m, ok := goRegistry[specifier]; ok {
		return m
	}
	return nil
}
func instanceModule(engine *Engine, module Module) (*goja.Object, error) {
	exports := engine.NewObject()
	var elements map[string]any
	if igm, ok := module.(InitializeModule); ok {
		elements = igm.ExportsWithEngine(engine)
	} else {
		elements = module.Exports()
	}
	for name, value := range elements {
		err := exports.Set(name, value)
		if err != nil {
			return nil, fmt.Errorf("error prepare import Module, couldn't set exports: %w", err)
		}
	}
	return exports, nil
}

// DumpDefines to path, global.d.ts contains top level types , pkg/name.d.ts contains go modules
func DumpDefines(path string) {
	_ = os.WriteFile(filepath.Join(path, "globals.d.ts"), ModDefines(), os.ModePerm)
	for name, bytes := range ModuleDefines() {
		if name == "go" {
			name = "go/index"
		}
		p := filepath.Join(path, name+".d.ts")
		_ = os.MkdirAll(filepath.Dir(p), os.ModePerm)
		_ = os.WriteFile(p, bytes, os.ModePerm)
	}
}
