package syscall

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	_ "github.com/ZenLiuCN/engine/modules/golang/sync"
)

var (
	//go:embed syscall.d.ts
	SyscallDefine   []byte
	SyscallDeclared = map[string]any{}
)

func init() {
	engine.RegisterModule(SyscallModule{})
}

type SyscallModule struct{}

func (S SyscallModule) Identity() string {
	return "golang/syscall"
}
func (S SyscallModule) TypeDefine() []byte {
	return SyscallDefine
}
func (S SyscallModule) Exports() map[string]any {
	return SyscallDeclared
}
