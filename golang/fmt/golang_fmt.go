package fmt

import (
	_ "embed"
	_ "github.com/ZenLiuCN/engine/golang/io"

	"fmt"
	"github.com/ZenLiuCN/engine"
)

var (
	//go:embed golang_fmt.d.ts
	FmtDefine   []byte
	FmtDeclared = map[string]any{
		"fscanf":       fmt.Fscanf,
		"fprint":       fmt.Fprint,
		"appendln":     fmt.Appendln,
		"scanf":        fmt.Scanf,
		"println":      fmt.Println,
		"scanln":       fmt.Scanln,
		"sscanf":       fmt.Sscanf,
		"fprintln":     fmt.Fprintln,
		"scan":         fmt.Scan,
		"printf":       fmt.Printf,
		"fscanln":      fmt.Fscanln,
		"formatString": fmt.FormatString,
		"fprintf":      fmt.Fprintf,
		"append":       fmt.Append,
		"errorf":       fmt.Errorf,
		"sprintf":      fmt.Sprintf,
		"sprint":       fmt.Sprint,
		"sprintln":     fmt.Sprintln,
		"sscanln":      fmt.Sscanln,
		"appendf":      fmt.Appendf,
		"print":        fmt.Print,
		"sscan":        fmt.Sscan,
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
