// Code generated by define_gene; DO NOT EDIT.
package zip

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"archive/zip"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	_ "github.com/ZenLiuCN/engine/modules/golang/io/fs"
	_ "github.com/ZenLiuCN/engine/modules/golang/time"
)

var (
	//go:embed archive_zip.d.ts
	ArchiveZipDefine   []byte
	ArchiveZipDeclared = map[string]any{
		"ErrChecksum":          zip.ErrChecksum,
		"ErrInsecurePath":      zip.ErrInsecurePath,
		"newReader":            zip.NewReader,
		"openReader":           zip.OpenReader,
		"registerCompressor":   zip.RegisterCompressor,
		"Deflate":              zip.Deflate,
		"ErrAlgorithm":         zip.ErrAlgorithm,
		"ErrFormat":            zip.ErrFormat,
		"fileInfoHeader":       zip.FileInfoHeader,
		"newWriter":            zip.NewWriter,
		"registerDecompressor": zip.RegisterDecompressor,
		"Store":                zip.Store,

		"emptyReadCloser":    engine.Empty[zip.ReadCloser],
		"emptyRefReadCloser": engine.EmptyRefer[zip.ReadCloser],
		"refOfReadCloser":    engine.ReferOf[zip.ReadCloser],
		"unRefReadCloser":    engine.UnRefer[zip.ReadCloser],
		"emptyFile":          engine.Empty[zip.File],
		"emptyRefFile":       engine.EmptyRefer[zip.File],
		"refOfFile":          engine.ReferOf[zip.File],
		"unRefFile":          engine.UnRefer[zip.File],
		"emptyFileHeader":    engine.Empty[zip.FileHeader],
		"emptyRefFileHeader": engine.EmptyRefer[zip.FileHeader],
		"refOfFileHeader":    engine.ReferOf[zip.FileHeader],
		"unRefFileHeader":    engine.UnRefer[zip.FileHeader]}
)

func init() {
	engine.RegisterModule(ArchiveZipModule{})
}

type ArchiveZipModule struct{}

func (S ArchiveZipModule) Identity() string {
	return "golang/archive/zip"
}
func (S ArchiveZipModule) TypeDefine() []byte {
	return ArchiveZipDefine
}
func (S ArchiveZipModule) Exports() map[string]any {
	return ArchiveZipDeclared
}
