package engine

import (
	_ "embed"
	"io"
)

var (
	//go:embed io.d.ts
	ioDefine []byte
)

type IoModule struct {
	m map[string]any
}

func (i *IoModule) Identity() string {
	return "go/io"
}

func (i *IoModule) TypeDefine() []byte {
	return ioDefine
}

func (i *IoModule) Exports() map[string]any {
	if i.m == nil {
		i.m = map[string]any{
			"copy":        io.Copy,
			"copyN":       io.CopyN,
			"copyBuffer":  io.CopyBuffer,
			"limitReader": io.LimitReader,
			"readAll":     io.ReadAll,
		}
	}
	return i.m
}
