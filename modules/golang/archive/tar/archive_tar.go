// Code generated by define_gene; DO NOT EDIT.
package tar

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"archive/tar"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	_ "github.com/ZenLiuCN/engine/modules/golang/io/fs"
	_ "github.com/ZenLiuCN/engine/modules/golang/time"
)

var (
	//go:embed archive_tar.d.ts
	ArchiveTarDefine   []byte
	ArchiveTarDeclared = map[string]any{
		"ErrHeader":          tar.ErrHeader,
		"newWriter":          tar.NewWriter,
		"TypeGNULongName":    tar.TypeGNULongName,
		"TypeReg":            tar.TypeReg,
		"ErrInsecurePath":    tar.ErrInsecurePath,
		"TypeBlock":          tar.TypeBlock,
		"TypeGNULongLink":    tar.TypeGNULongLink,
		"FormatUnknown":      tar.FormatUnknown,
		"TypeRegA":           tar.TypeRegA,
		"TypeXGlobalHeader":  tar.TypeXGlobalHeader,
		"TypeXHeader":        tar.TypeXHeader,
		"ErrFieldTooLong":    tar.ErrFieldTooLong,
		"TypeChar":           tar.TypeChar,
		"TypeDir":            tar.TypeDir,
		"TypeLink":           tar.TypeLink,
		"TypeSymlink":        tar.TypeSymlink,
		"FormatPAX":          tar.FormatPAX,
		"FormatUSTAR":        tar.FormatUSTAR,
		"ErrWriteAfterClose": tar.ErrWriteAfterClose,
		"FormatGNU":          tar.FormatGNU,
		"TypeCont":           tar.TypeCont,
		"TypeFifo":           tar.TypeFifo,
		"ErrWriteTooLong":    tar.ErrWriteTooLong,
		"fileInfoHeader":     tar.FileInfoHeader,
		"newReader":          tar.NewReader,
		"TypeGNUSparse":      tar.TypeGNUSparse,

		"emptyHeader":    engine.Empty[tar.Header],
		"emptyRefHeader": engine.EmptyRefer[tar.Header],
		"refOfHeader":    engine.ReferOf[tar.Header],
		"unRefHeader":    engine.UnRefer[tar.Header]}
)

func init() {
	engine.RegisterModule(ArchiveTarModule{})
}

type ArchiveTarModule struct{}

func (S ArchiveTarModule) Identity() string {
	return "golang/archive/tar"
}
func (S ArchiveTarModule) TypeDefine() []byte {
	return ArchiveTarDefine
}
func (S ArchiveTarModule) Exports() map[string]any {
	return ArchiveTarDeclared
}