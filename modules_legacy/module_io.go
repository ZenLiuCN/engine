package legacy

import (
	_ "embed"
	"io"
)

var (
	//go:embed module_io.d.ts
	ioDefine []byte
	ioMap    = map[string]any{
		"copy":             io.Copy,
		"copyN":            io.CopyN,
		"copyBuffer":       io.CopyBuffer,
		"readAll":          io.ReadAll,
		"writeString":      io.WriteString,
		"readAtLeast":      io.ReadAtLeast,
		"readFull":         io.ReadFull,
		"limitReader":      io.LimitReader,
		"newSectionReader": io.NewSectionReader,
		"newOffsetWriter":  io.NewOffsetWriter,
		"teeReader":        io.TeeReader,
		"nopCloser":        io.NopCloser,
		"multiReader":      io.MultiReader,
		"multiWriter":      io.MultiWriter,
		"pipe": func() struct {
			Reader *io.PipeReader
			Writer *io.PipeWriter
		} {
			r, w := io.Pipe()
			return struct {
				Reader *io.PipeReader
				Writer *io.PipeWriter
			}{
				Reader: r,
				Writer: w,
			}
		},
	}
)

type IoModule struct {
}

func (i IoModule) Identity() string {
	return "go/io"
}

func (i IoModule) TypeDefine() []byte {
	return ioDefine
}

func (i IoModule) Exports() map[string]any {
	return ioMap
}
