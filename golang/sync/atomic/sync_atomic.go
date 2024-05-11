// Code generated by define_gene; DO NOT EDIT.
package atomic

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"sync/atomic"
)

var (
	//go:embed sync_atomic.d.ts
	SyncAtomicDefine   []byte
	SyncAtomicDeclared = map[string]any{
		"addUint32":             atomic.AddUint32,
		"compareAndSwapInt32":   atomic.CompareAndSwapInt32,
		"compareAndSwapInt64":   atomic.CompareAndSwapInt64,
		"storeInt32":            atomic.StoreInt32,
		"swapInt64":             atomic.SwapInt64,
		"swapUintptr":           atomic.SwapUintptr,
		"compareAndSwapUintptr": atomic.CompareAndSwapUintptr,
		"loadUint64":            atomic.LoadUint64,
		"swapUint32":            atomic.SwapUint32,
		"addInt32":              atomic.AddInt32,
		"addInt64":              atomic.AddInt64,
		"compareAndSwapUint64":  atomic.CompareAndSwapUint64,
		"loadUint32":            atomic.LoadUint32,
		"storeUint64":           atomic.StoreUint64,
		"storeUintptr":          atomic.StoreUintptr,
		"swapUint64":            atomic.SwapUint64,
		"loadPointer":           atomic.LoadPointer,
		"storePointer":          atomic.StorePointer,
		"addUint64":             atomic.AddUint64,
		"addUintptr":            atomic.AddUintptr,
		"loadInt64":             atomic.LoadInt64,
		"storeInt64":            atomic.StoreInt64,
		"storeUint32":           atomic.StoreUint32,
		"compareAndSwapPointer": atomic.CompareAndSwapPointer,
		"compareAndSwapUint32":  atomic.CompareAndSwapUint32,
		"loadInt32":             atomic.LoadInt32,
		"loadUintptr":           atomic.LoadUintptr,
		"swapInt32":             atomic.SwapInt32,
		"swapPointer":           atomic.SwapPointer,

		"emptyUint32": func() (v atomic.Uint32) {
			return v
		},
		"refUint32": func() *atomic.Uint32 {
			var x atomic.Uint32
			return &x
		},
		"refOfUint32": func(x atomic.Uint32) *atomic.Uint32 {
			return &x
		},
		"emptyUint64": func() (v atomic.Uint64) {
			return v
		},
		"refUint64": func() *atomic.Uint64 {
			var x atomic.Uint64
			return &x
		},
		"refOfUint64": func(x atomic.Uint64) *atomic.Uint64 {
			return &x
		},
		"emptyUintptr": func() (v atomic.Uintptr) {
			return v
		},
		"refUintptr": func() *atomic.Uintptr {
			var x atomic.Uintptr
			return &x
		},
		"refOfUintptr": func(x atomic.Uintptr) *atomic.Uintptr {
			return &x
		},
		"emptyValue": func() (v atomic.Value) {
			return v
		},
		"refValue": func() *atomic.Value {
			var x atomic.Value
			return &x
		},
		"refOfValue": func(x atomic.Value) *atomic.Value {
			return &x
		},
		"emptyBool": func() (v atomic.Bool) {
			return v
		},
		"refBool": func() *atomic.Bool {
			var x atomic.Bool
			return &x
		},
		"refOfBool": func(x atomic.Bool) *atomic.Bool {
			return &x
		},
		"emptyInt32": func() (v atomic.Int32) {
			return v
		},
		"refInt32": func() *atomic.Int32 {
			var x atomic.Int32
			return &x
		},
		"refOfInt32": func(x atomic.Int32) *atomic.Int32 {
			return &x
		},
		"emptyInt64": func() (v atomic.Int64) {
			return v
		},
		"refInt64": func() *atomic.Int64 {
			var x atomic.Int64
			return &x
		},
		"refOfInt64": func(x atomic.Int64) *atomic.Int64 {
			return &x
		}}
)

func init() {
	engine.RegisterModule(SyncAtomicModule{})
}

type SyncAtomicModule struct{}

func (S SyncAtomicModule) Identity() string {
	return "golang/sync/atomic"
}
func (S SyncAtomicModule) TypeDefine() []byte {
	return SyncAtomicDefine
}
func (S SyncAtomicModule) Exports() map[string]any {
	return SyncAtomicDeclared
}
