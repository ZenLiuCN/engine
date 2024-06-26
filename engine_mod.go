package engine

import (
	"bytes"
	_ "embed"
	"github.com/ZenLiuCN/fn"
)

// Mod is a top level objects that not required to import
type Mod interface {
	Name() string
}
type InitializeMod interface {
	Mod
	Initialize(e *Engine) Mod // create copy of module with new Engine
}
type TopMod interface {
	Mod
	Register(e *Engine)
}

var (
	//go:embed global.d.ts
	globalDefine []byte
	registry     = map[string]Mod{}
)

// RegisterMod a default global module , returns false if already exists
func RegisterMod(module Mod) bool {
	if _, ok := registry[module.Name()]; ok {
		return false
	}
	registry[module.Name()] = module
	return true
}

// Mods of global registered
func Mods() []Mod {
	return fn.MapValues(registry)
}

// RemoveMod preloaded module
func RemoveMod(mod string) {
	delete(registry, mod)
}

// TypeDefined the element with typescript defines
type TypeDefined interface {
	TypeDefine() []byte
}

// ModDefines dump all possible type define (d.ts format) in registry
func ModDefines() []byte {
	var b bytes.Buffer
	b.WriteString("declare global{\n")
	b.Write(globalDefine)
	for _, module := range registry {
		if d, ok := module.(TypeDefined); ok {
			b.WriteRune('\n')
			b.Write(bytes.ReplaceAll(d.TypeDefine(), []byte("declare "), []byte("export ")))
		}
	}
	b.WriteString("}\n")
	return b.Bytes()
}
