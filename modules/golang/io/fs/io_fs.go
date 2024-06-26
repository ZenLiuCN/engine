// Code generated by define_gene; DO NOT EDIT.
package fs

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	_ "github.com/ZenLiuCN/engine/modules/golang/time"
	"io/fs"
)

var (
	//go:embed io_fs.d.ts
	IoFsDefine   []byte
	IoFsDeclared = map[string]any{
		"readDir":            fs.ReadDir,
		"readFile":           fs.ReadFile,
		"ErrInvalid":         fs.ErrInvalid,
		"fileInfoToDirEntry": fs.FileInfoToDirEntry,
		"ModeDevice":         fs.ModeDevice,
		"ModeExclusive":      fs.ModeExclusive,
		"ModeTemporary":      fs.ModeTemporary,
		"SkipAll":            fs.SkipAll,
		"ErrExist":           fs.ErrExist,
		"formatFileInfo":     fs.FormatFileInfo,
		"glob":               fs.Glob,
		"ModePerm":           fs.ModePerm,
		"ModeSetgid":         fs.ModeSetgid,
		"stat":               fs.Stat,
		"walkDir":            fs.WalkDir,
		"ErrNotExist":        fs.ErrNotExist,
		"ModeCharDevice":     fs.ModeCharDevice,
		"sub":                fs.Sub,
		"ErrPermission":      fs.ErrPermission,
		"ModeSetuid":         fs.ModeSetuid,
		"ModeSymlink":        fs.ModeSymlink,
		"SkipDir":            fs.SkipDir,
		"ModeDir":            fs.ModeDir,
		"ErrClosed":          fs.ErrClosed,
		"formatDirEntry":     fs.FormatDirEntry,
		"ModeAppend":         fs.ModeAppend,
		"ModeSticky":         fs.ModeSticky,
		"validPath":          fs.ValidPath,
		"ModeIrregular":      fs.ModeIrregular,
		"ModeNamedPipe":      fs.ModeNamedPipe,
		"ModeSocket":         fs.ModeSocket,
		"ModeType":           fs.ModeType,
		"toFileMode": func(n uint32) fs.FileMode {
			return fs.FileMode(n)
		},
	}
)

func init() {
	engine.RegisterModule(IoFsModule{})
}

type IoFsModule struct{}

func (S IoFsModule) Identity() string {
	return "golang/io/fs"
}
func (S IoFsModule) TypeDefine() []byte {
	return IoFsDefine
}
func (S IoFsModule) Exports() map[string]any {
	return IoFsDeclared
}
