package fmt

import (
	_ "embed"

	"fmt"
	"github.com/ZenLiuCN/engine"
	_ "github.com/ZenLiuCN/engine/golang/io"
)

var (
	//go:embed golang_fmt.d.ts
	FmtDefine   []byte
	FmtDeclared = map[string]any{
		"formatString": fmt.FormatString,
		"println":      fmt.Println,
		"scan":         fmt.Scan,
		"fscanln":      fmt.Fscanln,
		"scanf":        fmt.Scanf,
		"fscanf":       fmt.Fscanf,
		"printf":       fmt.Printf,
		"appendf":      fmt.Appendf,
		"sprint":       fmt.Sprint,
		"fprintln":     fmt.Fprintln,
		"sprintln":     fmt.Sprintln,
		"sscanln":      fmt.Sscanln,
		"errorf":       fmt.Errorf,
		"sprintf":      fmt.Sprintf,
		"append":       fmt.Append,
		"sscan":        fmt.Sscan,
		"sscanf":       fmt.Sscanf,
		"print":        fmt.Print,
		"appendln":     fmt.Appendln,
		"scanln":       fmt.Scanln,
		"fprintf":      fmt.Fprintf,
		"fprint":       fmt.Fprint,
		"fscan":        fmt.Fscan,
	}
)

func init() {
	engine.RegisterModule(FmtModule{})
}

type FmtModule struct{}

func (S FmtModule) Identity() string {
	return "golang/fmt"
}
func (S FmtModule) TypeDefine() []byte {
	return FmtDefine
}
func (S FmtModule) Exports() map[string]any {
	return FmtDeclared
}
