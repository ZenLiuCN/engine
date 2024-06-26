// Code generated by define_gene; DO NOT EDIT.
package signal

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/golang/os"
	"os/signal"
)

var (
	//go:embed os_signal.d.ts
	OsSignalDefine   []byte
	OsSignalDeclared = map[string]any{
		"notify":        signal.Notify,
		"notifyContext": signal.NotifyContext,
		"reset":         signal.Reset,
		"stop":          signal.Stop,
		"ignore":        signal.Ignore,
		"ignored":       signal.Ignored,
	}
)

func init() {
	engine.RegisterModule(OsSignalModule{})
}

type OsSignalModule struct{}

func (S OsSignalModule) Identity() string {
	return "golang/os/signal"
}
func (S OsSignalModule) TypeDefine() []byte {
	return OsSignalDefine
}
func (S OsSignalModule) Exports() map[string]any {
	return OsSignalDeclared
}
