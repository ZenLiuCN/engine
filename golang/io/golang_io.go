package io

import (
	_ "embed"

	"github.com/ZenLiuCN/engine"
	"io"
)

var (
	//go:embed golang_io.d.ts
	IoDefine   []byte
	IoDeclared = map[string]any{
		"newOffsetWriter":  io.NewOffsetWriter,
		"multiWriter":      io.MultiWriter,
		"copy":             io.Copy,
		"limitReader":      io.LimitReader,
		"nopCloser":        io.NopCloser,
		"writeString":      io.WriteString,
		"readFull":         io.ReadFull,
		"teeReader":        io.TeeReader,
		"pipe":             io.Pipe,
		"copyBuffer":       io.CopyBuffer,
		"newSectionReader": io.NewSectionReader,
		"multiReader":      io.MultiReader,
		"readAtLeast":      io.ReadAtLeast,
		"readAll":          io.ReadAll,
		"copyN":            io.CopyN,
	}
)

func init() {
	engine.RegisterModule(IoModule{})
}

type IoModule struct{}

func (S IoModule) Identity() string {
	return "golang/io"
}
func (S IoModule) TypeDefine() []byte {
	return IoDefine
}
func (S IoModule) Exports() map[string]any {
	return IoDeclared
}
