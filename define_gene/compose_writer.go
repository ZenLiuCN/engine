package main

import (
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ComposeWriter struct {
	Declare *Writer
	Impl    *Writer
}

func (w *ComposeWriter) GoString() string {
	return fmt.Sprintf("Declare:%vImpl:%v", w.Declare, w.Impl)
}
func (w *ComposeWriter) Init() {
	if w.Declare == nil {
		w.Declare = get()
	} else {
		w.Declare.Clear()
	}
	if w.Impl == nil {
		w.Impl = get()
	} else {
		w.Impl.Clear()
	}
}
func (w *ComposeWriter) Free() {
	if w.Declare != nil {
		w.Declare.Free()
	}
	if w.Impl != nil {
		w.Impl.Free()
	}
}
func (w *ComposeWriter) Merge() {
	if len(w.Impl.imp) > 0 {
		w.Declare.Imports(w.Impl.imp.Values()...)
		w.Impl.imp.Clear()
	}
	if len(w.Impl.use) > 0 {
		w.Declare.Uses(w.Impl.use.Values()...)
		w.Impl.use.Clear()
	}
}

func (w *ComposeWriter) Empty() bool {
	return w.Declare.Len() == 0 && w.Impl.Len() == 0
}

func (w *ComposeWriter) Append(i *ComposeWriter) {
	if i != nil && !i.Empty() {
		i.Merge()
		w.Declare.Append(i.Declare)
		w.Impl.Append(i.Impl)
	}
}
func (w *ComposeWriter) MergeWithoutContent(i *ComposeWriter) {
	if i != nil && !i.Empty() {
		i.Merge()
		w.Declare.MergeWithoutContent(i.Declare)
		w.Impl.MergeWithoutContent(i.Impl)
	}
}

type ModuleWriter struct {
	Golang     *ComposeWriter
	Typescript *ComposeWriter
}

func (w *ModuleWriter) GoString() string {
	return fmt.Sprintf("Golang:%vTypescript:%v", w.Golang, w.Typescript)
}
func (w *ModuleWriter) Free() {
	if w != nil {
		w.Typescript.Free()
		w.Golang.Free()
	}

}
func (w *ModuleWriter) Init() {
	if w.Golang == nil {
		w.Golang = new(ComposeWriter)
	}
	w.Golang.Init()
	if w.Typescript == nil {
		w.Typescript = new(ComposeWriter)
	}
	w.Typescript.Init()
}
func (w *ModuleWriter) GoEntry(pkg, name string) {
	w.Golang.Declare.Ident("\"%s\"", name)
	w.Golang.Declare.Format(":%s.%s,\n", pkg, name)
}
func (w *ModuleWriter) GoEntryRaw(pkg, name string) {
	w.Golang.Declare.Format("\"%[2]s\":%[1]s.%[2]s,\n", pkg, name)
}
func (w *ModuleWriter) GoImpl(format string, args ...any) {
	w.Golang.Impl.Format(format, args)
}
func (w *ModuleWriter) GoIdent(format string, name string) {
	w.Golang.Impl.Ident(format, name)
}
func (w *ModuleWriter) GoImport(args ...string) {
	w.Golang.Declare.Imports(args...)
}
func (w *ModuleWriter) GoDefine(format string, args ...any) {
	w.Golang.Declare.Format(format, args)
}

func (w *ModuleWriter) TsImport(args ...string) {
	w.Typescript.Declare.Imports(args...)
}
func (w *ModuleWriter) TsUse(args ...string) {
	w.Typescript.Declare.Uses(args...)
}
func (w *ModuleWriter) TsDefine(format string, args ...any) {
	if w.Typescript.Declare.Len() > 0 {
		w.Typescript.Declare.Format(",")
	}
	w.Typescript.Declare.Format(format, args...)
}
func (w *ModuleWriter) TsImpl(format string, args ...any) {
	w.Typescript.Impl.Format(format, args...)
}
func (w *ModuleWriter) TsIdent(format string, name string) {
	w.Typescript.Impl.Ident(format, name)
}
func (w *ModuleWriter) TsIdentType(format string, name string, array bool) {
	w.Typescript.Impl.Type(format, name, array)
}
func (w *ModuleWriter) FlushFile(f *FileWriter, debug bool) {
	o := get()
	defer o.Free()
	{
		o.Format(`declare module "`)
		if f.golangPkg {
			o.Format("golang/%s\" {\n", f.pkgPath)
		} else {
			o.Format("go/%s\" {\n", strings.ReplaceAll(f.pkgPath, ".", "/"))
		}
		w.Typescript.Merge()
		for s := range w.Typescript.Declare.imp {
			if strings.Contains(s, ".") {
				o.Format("\t// @ts-ignore\n\timport * as %s from '%s'\n", filepath.Base(s), s)
			} else {
				o.Format("\t// @ts-ignore\n\timport * as %s from 'golang/%s'\n", filepath.Base(s), s)
			}
		}
		if len(w.Typescript.Declare.use) > 0 {
			o.Format("\t// @ts-ignore\n\timport type {%s} from 'go'\n", strings.Join(w.Typescript.Declare.use.Values(), ","))
		}
		o.Format("\n\n")
		_, _ = f.Typescript.Declare.WriteTo(o.Buffer)
		_, _ = f.Typescript.Impl.WriteTo(o.Buffer)
		o.Format("\n}")
		if !debug {
			p := filepath.Join(f.out, strings.ReplaceAll(f.pkgPath, ".", "/"), f.dts)
			if !f.override && Exists(p) {
				panic(fmt.Errorf("%s exists for %s", p, f.pkg.PkgPath))
			}
			_ = os.WriteFile(p, o.Bytes(), os.ModePerm)
		} else {
			fmt.Println(o.String())
		}
	}
	{
		o.Clear()
		w.Golang.Merge()
		pkg := f.pkg.Name
		if f.golangPkg {
			pkg = filepath.Base(f.pkg.PkgPath)
		}
		o.Format(`package %s
import ( 
	_ "embed"
`, pkg)
		if w.Golang.Declare.imp.Len() > 0 {
			for s := range w.Golang.Declare.imp {
				o.Format("\t\"%s\"\n", s)
			}
		}
		if w.Typescript.Declare.imp.Len() > 0 {
			for s := range w.Typescript.Declare.imp {
				if !strings.Contains(s, ".") {
					o.Format("\t_ \"github.com/ZenLiuCN/engine/golang/%s\"\n", s)
				} else {
					o.Format("\t_ \"github.com/ZenLiuCN/engine/%s\"\n", strings.ReplaceAll(s, ".", "/"))
				}
			}
		}
		if w.Golang.Declare.use.Len() > 0 {
			for s := range w.Golang.Declare.use {
				o.Format("\t\"%s\"\n", s)
			}
		}
		o.Format(`
	"github.com/ZenLiuCN/engine"
	"%[1]s"
)
var(
	//go:embed %[2]s
	%[3]sDefine []byte
	%[3]sDeclared = map[string]any {
`, f.pkg.PkgPath, f.dts, f.module)
		o.Append(w.Golang.Declare)
		o.Format(`
	}
)
func init() {
	engine.RegisterModule(%[1]sModule{})
}
type %[1]sModule struct{}
func (S %[1]sModule) Identity() string {
	return "golang/%[2]s"
}
func (S %[1]sModule) TypeDefine() []byte {
	return %[1]sDefine
}
func (S %[1]sModule) Exports() map[string]any {
	return %[1]sDeclared
}
`, f.module, f.pkg.PkgPath)
		o.Append(w.Golang.Impl)
		b, err := format.Source(o.Bytes())
		if err != nil {
			log.Printf("generated go source have error: %s", err)
			b = o.Bytes()
		}
		if !debug {
			p := filepath.Join(f.out, strings.ReplaceAll(f.pkgPath, ".", "/"), f.mod)
			if !f.override && Exists(p) {
				panic(fmt.Errorf("%s exists for %s", p, f.pkg.PkgPath))
			}
			_ = os.WriteFile(p, b, os.ModePerm)
		} else {
			fmt.Println(string(b))
		}
	}
	log.Printf("generated: %s", f.pkg.PkgPath)
}

func (w *ModuleWriter) Empty() bool {
	return w.Golang.Empty() && w.Typescript.Empty()
}
func (w *ModuleWriter) Append(i *ModuleWriter) {
	if !i.Golang.Empty() {
		i.Golang.Merge()
		w.Golang.Declare.Append(i.Golang.Declare)
		w.Golang.Impl.Append(i.Golang.Impl)
	}
	if !i.Typescript.Empty() {
		i.Typescript.Merge()
		w.Typescript.Declare.Append(i.Typescript.Declare)
		w.Typescript.Impl.Append(i.Typescript.Impl)
	}
}
func (w *ModuleWriter) Prepend(i *ModuleWriter) {
	i.Golang.Merge()
	i.Typescript.Merge()
	w.Golang.Declare.Prepend(i.Golang.Declare)
	w.Golang.Impl.Prepend(i.Golang.Impl)
	w.Typescript.Declare.Prepend(i.Typescript.Declare)
	w.Typescript.Impl.Prepend(i.Typescript.Impl)
}
