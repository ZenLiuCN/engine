// Code generated by define_gene; DO NOT EDIT.
package ring

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"container/ring"
)

var (
	//go:embed container_ring.d.ts
	ContainerRingDefine   []byte
	ContainerRingDeclared = map[string]any{
		"New": ring.New,

		"emptyRing":    engine.Empty[ring.Ring],
		"emptyRefRing": engine.EmptyRefer[ring.Ring],
		"refOfRing":    engine.ReferOf[ring.Ring],
		"unRefRing":    engine.UnRefer[ring.Ring]}
)

func init() {
	engine.RegisterModule(ContainerRingModule{})
}

type ContainerRingModule struct{}

func (S ContainerRingModule) Identity() string {
	return "golang/container/ring"
}
func (S ContainerRingModule) TypeDefine() []byte {
	return ContainerRingDefine
}
func (S ContainerRingModule) Exports() map[string]any {
	return ContainerRingDeclared
}
