package bufio

import (
	_ "embed"
	_ "github.com/ZenLiuCN/engine/golang/io"

	"bufio"
	"github.com/ZenLiuCN/engine"
)

var (
	//go:embed golang_bufio.d.ts
	BufioDefine   []byte
	BufioDeclared = map[string]any{
		"newScanner":    bufio.NewScanner,
		"scanWords":     bufio.ScanWords,
		"newReadWriter": bufio.NewReadWriter,
		"scanBytes":     bufio.ScanBytes,
		"newReaderSize": bufio.NewReaderSize,
		"newReader":     bufio.NewReader,
		"newWriter":     bufio.NewWriter,
		"scanRunes":     bufio.ScanRunes,
		"newWriterSize": bufio.NewWriterSize,
		"scanLines":     bufio.ScanLines,
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
