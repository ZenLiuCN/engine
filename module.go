package engine

import (
	"bytes"
	"github.com/ZenLiuCN/fn"
	"log/slog"
)

type Module interface {
	Name() string
}
type InitializeModule interface {
	Module
	Initialize(engine *Engine) Module // create copy of module with new Engine
}
type TopModule interface {
	Module
	Register(engine *Engine)
}

var (
	registry map[string]Module
)

func init() {
	registry = make(map[string]Module)
	Register(NewConsole(slog.Default()))
	Register(&Compiler{})
	Register(&Require{})
	Register(BufferModule{})
	Register(EngineModule{})
	Register(CryptoModule{})
	Register(HasherInstance)
}

// Register a module , returns false if already exists
func Register(module Module) bool {
	if _, ok := registry[module.Name()]; ok {
		return false
	}
	registry[module.Name()] = module
	return true
}

// ForceRegister a new Module, replace old one if exists.
func ForceRegister(module Module) {
	registry[module.Name()] = module
}

// Modules of global registered
func Modules() []Module {
	return fn.MapValues(registry)
}

// Registry the global module registry
func Registry() map[string]Module {
	return registry
}

// Remove preloaded module
func Remove(module string) {
	delete(registry, module)
}

type TypeDefined interface {
	TypeDefine() []byte
}

// TypeDefines dump all possible type define (d.ts format) in registry
func TypeDefines() []byte {
	var b bytes.Buffer
	for _, module := range registry {
		if d, ok := module.(TypeDefined); ok {
			b.WriteRune('\n')
			b.Write(d.TypeDefine())
		}
	}
	return b.Bytes()
}
