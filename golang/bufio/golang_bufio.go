package bufio

import (
	_ "embed"

	"bufio"
	"github.com/ZenLiuCN/engine"
	_ "github.com/ZenLiuCN/engine/golang/io"
)

var (
	//go:embed golang_bufio.d.ts
	BufioDefine   []byte
	BufioDeclared = map[string]any{
		"scanRunes":     bufio.ScanRunes,
		"scanLines":     bufio.ScanLines,
		"scanWords":     bufio.ScanWords,
		"newScanner":    bufio.NewScanner,
		"newReader":     bufio.NewReader,
		"scanBytes":     bufio.ScanBytes,
		"newReaderSize": bufio.NewReaderSize,
		"newWriterSize": bufio.NewWriterSize,
		"newWriter":     bufio.NewWriter,
		"newReadWriter": bufio.NewReadWriter,
	}
)

func init() {
	engine.RegisterModule(BufioModule{})
}

type BufioModule struct{}

func (S BufioModule) Identity() string {
	return "golang/bufio"
}
func (S BufioModule) TypeDefine() []byte {
	return BufioDefine
}
func (S BufioModule) Exports() map[string]any {
	return BufioDeclared
}
