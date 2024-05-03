package main

import (
	"bytes"
	"fmt"
	"github.com/ZenLiuCN/engine/define_gene/spec"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"log"
	"slices"
)

func printf(format string, args ...any) {
	if debug {
		_ = log.Output(2, "[DEBUG] "+fmt.Sprintf(format, args...))
	}
}
func tracef(n int, format string, args ...any) {
	if debug {
		_ = log.Output(n+2, "[TRACE] "+fmt.Sprintf(format, args...))
	}
}

type TypeWriter struct {
	pkg    *packages.Package //current package
	name   string            //possible type name
	define *Define           //current type
	*ModuleWriter
	*TypeRegistry
}

func (w *TypeWriter) With(pkg *packages.Package, name string, define *Define) {
	w.init()
	w.define = define
	w.name = name
	w.pkg = pkg
	w.TypeRegistry.pkg = pkg
}

func (w *TypeWriter) init() {
	if w.ModuleWriter == nil {
		w.ModuleWriter = new(ModuleWriter)
	}
	w.ModuleWriter.Init()
	if w.TypeRegistry == nil {
		w.TypeRegistry = new(TypeRegistry)
	}
}

func (w *TypeWriter) Process() {
	switch w.define.Type {
	case StructType:
		w.declareStructType()
	case MapType:
		w.declareMapType()
	case InterfaceType:
		w.declareInterfaceType()
	case ArrayType:
		w.declareArrayType()
	case ChanType:
		w.declareChanType()
	case FuncType:
		w.declareFuncType()
	case FuncDecl:
		w.declareFuncDecl()
	case IdentType:
		w.declareIdentType()
	case ConstValue:
		w.declareConstType()
	default:
		log.Printf("miss define type :%s", w.define.Type.String())
	}

}

func (w *TypeWriter) declareStructType() {
	printf("[DECLARE]  %#+v", w.define)
	t := w.define.Prim
	iw := new(InfoWriter)
	defer iw.Free()
	iw.Root = w.define
	w.writeType(Flags{FStruct}, t.Struct, iw)
	iw.Pop()
	for x := range w.define.Elt {
		if x == t {
			continue
		}
		w.writeTypeInfo(Flags{FStruct}, x, iw)
		iw.Pop()
	}
	iw.ModuleWriter.Golang.Merge()
	w.Golang.Append(iw.ModuleWriter.Golang)
	x := get()
	defer x.Free()
	w.TsUse(UseStruct)
	w.InterfacePatch(iw.ModuleWriter)
	x.Format("\texport interface %[1]s extends Struct<%[1]s>", t.Spec.Name.Name)
	if iw.Typescript.Declare.Len() > 0 {
		x.Format(",%s", iw.Typescript.Declare.Buffer.String())
	}
	x.Format("{\n%s\n\t}\n", iw.Typescript.Impl.Buffer.String())
	_, _ = w.Typescript.Impl.Write(x.Bytes())
	w.Typescript.MergeWithoutContent(iw.Typescript)
}
func (w *TypeWriter) InterfacePatch(iw *ModuleWriter) {
	//!! interface patch
	if bytes.Contains(iw.Typescript.Impl.Buffer.Bytes(), []byte("close():error")) {
		if w.pkg.PkgPath == "io" {
			iw.TsDefine("Closer")
		} else {
			iw.TsImpl("io")
			iw.TsDefine("io.Closer")
		}
	}
}

func (w *TypeWriter) declareMapType() {
	printf("[DECLARE]  %#+v", w.define)
	p := w.define.Prim
	iw := new(InfoWriter)
	defer iw.Free()
	iw.Init()
	iw.TypeRegistry = w.TypeRegistry
	iw.Root = w.define
	ik := iw.writeTypeBuffer(Flags{FStruct}, p.Map.Key)
	iw.Pop()
	iv := iw.writeTypeBuffer(Flags{FStruct}, p.Map.Value)
	iw.Pop()
	key := ik.Typescript.Impl.String()
	value := iv.Typescript.Impl.String()
	ik.Golang.Merge()
	iv.Golang.Merge()
	iw.Golang.Append(ik.Golang)
	iw.Golang.Append(iv.Golang)
	ik.Typescript.Merge()
	iv.Typescript.Merge()
	iw.Typescript.MergeWithoutContent(ik.Typescript)
	iw.Typescript.MergeWithoutContent(iv.Typescript)
	ik.Free()
	iv.Free()
	//w.writeType(Flags{FStruct}, p.Map, iw)
	//iw.Pop()
	for x := range w.define.Elt {
		if x == p {
			continue
		}
		iw.TsImpl("\t\t//@ts-ignore\n")
		w.writeTypeInfo(Flags{FStruct}, x, iw)
		iw.Pop()
	}
	iw.ModuleWriter.Golang.Merge()
	w.Golang.Append(iw.ModuleWriter.Golang)
	x := get()
	defer x.Free()
	w.TsUse(UseStruct)
	w.InterfacePatch(iw.ModuleWriter)
	x.Format("\texport interface %[1]s extends Struct<%[1]s>,Record<%s,%s>", p.Spec.Name.Name, key, value)
	if iw.Typescript.Declare.Len() > 0 {
		x.Format(",%s", iw.Typescript.Declare.Buffer.String())
	}
	x.Format("{\n%s\n\t}\n", iw.Typescript.Impl.Buffer.String())
	_, _ = w.Typescript.Impl.Write(x.Bytes())
	w.Typescript.MergeWithoutContent(iw.Typescript)
	//panic(fmt.Errorf("[DECLARE]  %#+v", w.define))
}

func (w *TypeWriter) declareInterfaceType() {
	printf("[DECLARE]  %#+v", w.define)
	t := w.define.Prim
	iw := new(InfoWriter)
	defer iw.Free()
	iw.Root = w.define
	w.writeType(Flags{FInterface}, t.Interface, iw)
	iw.Pop()
	iw.ModuleWriter.Golang.Merge()
	w.Golang.Append(iw.ModuleWriter.Golang)
	x := get()
	defer x.Free()
	w.TsUse(UseInterface)
	x.Format("\n\texport interface %[1]s extends Proto<%[1]s>", w.define.Spec.Name.Name)
	if iw.Typescript.Declare.Len() > 0 {
		x.Format(",%s", iw.Typescript.Declare.Buffer.String())
	}
	x.Format("{\n%s\n\t}\n", iw.Typescript.Impl.Buffer.String())
	_, _ = w.Typescript.Impl.Write(x.Bytes())
	w.Typescript.MergeWithoutContent(iw.Typescript)
}

func (w *TypeWriter) declareArrayType() {
	printf("[DECLARE]  %#+v", w.define)
	panic(fmt.Errorf("[DECLARE]  %#+v", w.define))
}

func (w *TypeWriter) declareChanType() {
	printf("[DECLARE]  %#+v", w.define)
	panic(fmt.Errorf("[DECLARE]  %#+v", w.define))

}

func (w *TypeWriter) declareFuncType() {
	printf("[DECLARE]  %#+v", w.define)
	p := w.define.Prim
	iw := new(InfoWriter)
	defer iw.Free()
	iw.Root = w.define
	w.writeType(Flags{FField}, p.Func, iw)
	iw.Pop()
	iw.ModuleWriter.Golang.Merge()
	w.Golang.Append(iw.ModuleWriter.Golang)
	x := get()
	defer x.Free()
	x.Format("\n\texport type %[1]s = ", p.Spec.Name.Name)
	x.Format("%s\n", iw.Typescript.Impl.Buffer.String())
	_, _ = w.Typescript.Impl.Write(x.Bytes())
	w.Typescript.MergeWithoutContent(iw.Typescript)
	//panic(fmt.Errorf("[DECLARE]  %#+v", w.define))
}

func (w *TypeWriter) declareFuncDecl() {
	printf("[DECLARE]  %#+v", w.define)
	t := w.define.Prim
	if t.FuncDecl.Name.IsExported() {
		w.TsIdent("\texport function %s", t.FuncDecl.Name.Name)
		iw := new(InfoWriter)
		iw.Root = w.define
		w.writeType(Flags{FFunction}, t.FuncDecl.Type, iw)
		iw.Pop()
		w.Append(iw.ModuleWriter)
		w.TsImpl("\n")
		w.GoEntry(w.pkg.Name, t.FuncDecl.Name.Name)
	}
}

func (w *TypeWriter) declareIdentType() {
	printf("[DECLARE]  %#+v", w.define)
	if w.define.Elt.Len() == 1 {
		w.TsUse(UseAlias)
		w.TsImpl("\texport interface %s extends ", w.define.Spec.Name.Name)
		w.TsIdentType("Alias<%s>{\n\t}\n", w.define.Prim.Ident.Name, false)
	} else {
		iw := new(InfoWriter)
		defer iw.Free()
		iw.Root = w.define
		for x := range w.define.Elt {
			if x == w.define.Prim {
				continue
			}
			w.writeTypeInfo(Flags{FInterface}, x, iw)
			iw.Pop()
		}
		iw.ModuleWriter.Golang.Merge()
		w.Golang.Append(iw.ModuleWriter.Golang)
		x := get()
		defer x.Free()
		w.TsUse(UseAlias)
		w.InterfacePatch(iw.ModuleWriter)
		x.Format("\texport interface %[1]s extends Alias<%[1]s>", w.define.Spec.Name.Name)
		if iw.Typescript.Declare.Len() > 0 {
			x.Format(",%s", iw.Typescript.Declare.Buffer.String())
		}
		x.Format("{\n%s\n\t}\n", iw.Typescript.Impl.Buffer.String())
		_, _ = w.Typescript.Impl.Write(x.Bytes())
		w.Typescript.MergeWithoutContent(iw.Typescript)
		//TODO values
	}
}

func (w *TypeWriter) declareConstType() {
	printf("[DECLARE]  %#+v", w.define)
	values := w.define.Val.Values()
	base := slices.IndexFunc(values, func(v *Value) bool {
		return v.Ident != nil || (v.Spec != nil && v.Spec.Type != nil)
	})
	if base < 0 {
		for _, value := range values {
			w.TsImpl("\texport const %s=", value.Name)
			if d, ok := w.pkg.TypesInfo.Types[value.Spec.Values[0]]; ok {
				w.TsImpl("%s\n", d.Value.String())
			} else {
				panic(fmt.Errorf("[DECLARE]  %#+v", w.define))
			}
		}
	} else {
		b := values[base]
		if b.Ident != nil && b.Ident.Obj != nil {
			for _, value := range values {
				w.TsImpl("\texport const %s:%s\n", value.Name, b.Ident.Obj.Name)
				w.GoEntryRaw(w.pkg.Name, value.Name)
			}
		} else {
			panic(fmt.Errorf("[DECLARE]  %#+v", w.define))
		}
	}

}

//region TypeWrite

func (w *TypeWriter) writeType(flag Flags, typ ast.Expr, iw *InfoWriter) {
	iw.TypeRegistry = w.TypeRegistry
	iw.Init()
	t := w.lookupType(typ)
	n := len(t)
	switch n {
	case 0:
		log.Printf("miss type %v", typ)
	default:
		node := t[len(t)-1]
		iw.Push(t)
		switch node.Node {
		case spec.AstIdent:
			iw.writeIdent(flag, node.Ident)
		case spec.AstSelectorExpr:
			iw.writeSelector(flag, node.Selector, nil)
		case spec.AstStarExpr:
			iw.writeStar(flag, node.Star)
		case spec.AstFuncType:
			iw.writeFunc(flag, node.Func)
		case spec.AstArrayType:
			iw.writeArray(flag, node.Array)
		case spec.AstMapType:
			iw.writeMap(flag, node.Map)
		case spec.AstStructType:
			iw.writeStruct(flag, node.Struct)
		case spec.AstInterfaceType:
			iw.writeInterface(flag, node.Interface)
		case spec.AstChanType:
			iw.writeChan(flag, node.Chan)
		case spec.AstEllipsis:
			iw.writeEllipsis(flag, node.Ellipsis)
		default:
			log.Printf("miss node type %v", node.Node)
		}
	}
}

func (w *TypeWriter) writeTypeInfo(flag Flags, x *Type, iw *InfoWriter) {
	iw.TypeRegistry = w.TypeRegistry
	iw.Init()
	switch x.Type {
	case IdentType:
		iw.writeIdent(flag, x.Ident)
	case SelectorType:
		iw.writeSelector(flag, x.Select, nil)
	case StarType:
		iw.writeStar(flag, x.Star)
	case FuncType:
		iw.writeFunc(flag, x.Func)
	case FuncDecl:
		iw.writeFuncDecl(flag, x.FuncDecl)
	case StarMethodType, IdentMethodType:
		iw.writeFuncDecl(append(flag, FMethod), x.FuncDecl)
	case ArrayType:
		iw.writeArray(flag, x.Array)
	case MapType:
		iw.writeMap(flag, x.Map)
	case StructType:
		iw.writeStruct(flag, x.Struct)
	case InterfaceType:
		iw.writeInterface(flag, x.Interface)
	case ChanType:
		iw.writeChan(flag, x.ChanType)
		//case ConstValue:

	default:
		log.Printf("miss  type %v", x.Type)
	}

}

//endregion
