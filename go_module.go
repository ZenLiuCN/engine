package engine

import (
	"fmt"
	"github.com/dop251/goja"
	"os"
	"path/filepath"
	"strings"
)

// GoModule support for import by require, not a global instance as Module
type (
	GoModule interface {
		Identity() string // Identity the full module identity
		TypeDefine() []byte
		Exports() map[string]any
	}
	GoInitializeModule interface {
		GoModule
		ExportsWithEngine(engine *Engine) map[string]any
	}
)

var (
	goRegistry = map[string]GoModule{}
)

func RegisterModule(module GoModule) bool {
	if _, ok := goRegistry[module.Identity()]; ok {
		return false
	}
	goRegistry[module.Identity()] = module
	return true
}
func RemoveModule(module string) {
	delete(goRegistry, module)
}

func GoModuleDefines() map[string][]byte {
	m := make(map[string][]byte)
	for s, module := range goRegistry {
		m[strings.ReplaceAll(s, "/", "_")] = module.TypeDefine()
	}
	return m
}
func resolveGoModule(specifier string) GoModule {
	if m, ok := goRegistry[specifier]; ok {
		return m
	}
	return nil
}
func instanceGoModule(engine *Engine, module GoModule) (*goja.Object, error) {
	exports := engine.NewObject()
	var elements map[string]any
	if igm, ok := module.(GoInitializeModule); ok {
		elements = igm.ExportsWithEngine(engine)
	} else {
		elements = module.Exports()
	}
	for name, value := range elements {
		err := exports.Set(name, value)
		if err != nil {
			return nil, fmt.Errorf("error prepare import GoModule, couldn't set exports: %w", err)
		}
	}
	return exports, nil
}

// DumpDefines to path, global.d.ts contains top level types , pkg_name.d.ts contains go modules
func DumpDefines(path string) {
	_ = os.WriteFile(filepath.Join(path, "global.d.ts"), TypeDefines(), os.ModePerm)
	for name, bytes := range GoModuleDefines() {
		_ = os.WriteFile(filepath.Join(path, name+".d.ts"), bytes, os.ModePerm)
	}
}
