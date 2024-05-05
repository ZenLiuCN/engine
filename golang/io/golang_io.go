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
		"readAll":          io.ReadAll,
		"teeReader":        io.TeeReader,
		"limitReader":      io.LimitReader,
		"newOffsetWriter":  io.NewOffsetWriter,
		"writeString":      io.WriteString,
		"multiWriter":      io.MultiWriter,
		"copyN":            io.CopyN,
		"copyBuffer":       io.CopyBuffer,
		"nopCloser":        io.NopCloser,
		"pipe":             io.Pipe,
		"newSectionReader": io.NewSectionReader,
		"readAtLeast":      io.ReadAtLeast,
		"copy":             io.Copy,
		"multiReader":      io.MultiReader,
		"readFull":         io.ReadFull,
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
