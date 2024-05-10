// Code generated by define_gene; DO NOT EDIT.
package heap

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"container/heap"
	_ "github.com/ZenLiuCN/engine/golang/sort"
)

var (
	//go:embed container_heap.d.ts
	ContainerHeapDefine   []byte
	ContainerHeapDeclared = map[string]any{
		"remove": heap.Remove,
		"fix":    heap.Fix,
		"init":   heap.Init,
		"pop":    heap.Pop,
		"push":   heap.Push,
	}
)

func init() {
	engine.RegisterModule(ContainerHeapModule{})
}

type ContainerHeapModule struct{}

func (S ContainerHeapModule) Identity() string {
	return "golang/container/heap"
}
func (S ContainerHeapModule) TypeDefine() []byte {
	return ContainerHeapDefine
}
func (S ContainerHeapModule) Exports() map[string]any {
	return ContainerHeapDeclared
}
