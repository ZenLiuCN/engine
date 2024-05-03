package main

import (
	"fmt"
	"github.com/ZenLiuCN/engine/define_gene/spec"
	"go/ast"
	"go/types"
	"log"
	"strings"
)

type Use = string

const (
	UseStar      Use = "Ref"
	UseStruct        = "Struct"
	UseAlias         = "Alias"
	UseInterface     = "Proto"
)

type InfoWriter struct {
	Root *Define
	*ModuleWriter
	*TypeRegistry
	Types
	stack []int
}

func (w *InfoWriter) Push(t Types) {
	tracef(1, "[PUSH] %#+v << %#+v == %#+v", w.Types, w.stack, t)
	w.Types = append(w.Types, t...)
	w.stack = append(w.stack, w.Types.Len())
	tracef(1, "[PUSHED] %#+v << %#+v ", w.Types, w.stack)
}
func (w *InfoWriter) LastN(n int) TypeInfo {
	if len(w.Types) < n {
		return TypeInfo{}
	}
	return w.Types[len(w.Types)-n]
}
func (w *InfoWriter) Pop() {
	if len(w.stack) == 0 {
		return
	}
	tracef(1, "[POP] %#+v << %#+v ", w.Types, w.stack)
	n := w.stack[len(w.stack)-1]
	w.stack = w.stack[:len(w.stack)-1]
	w.Types = w.Types[:n-1]
	tracef(1, "[POP] %#+v ", w.Types)
}
func (w *InfoWriter) Init() {
	if w.ModuleWriter == nil {
		w.ModuleWriter = new(ModuleWriter)
		w.ModuleWriter.Init()
	}
}

// region Types
func (w *InfoWriter) writeFuncDecl(flag Flags, decl *ast.FuncDecl) {
	if flag.First().IsStruct() || flag.First().IsInterface() {
		if !decl.Name.IsExported() {
			return
		}
		w.TsIdent("\t\t%s", decl.Name.Name)
		w.writeFunc(flag, decl.Type)
		w.TsImpl("\n")
		w.Pop()
	} else {
		w.TsIdent("%s", decl.Name.Name)
		w.writeFunc(flag, decl.Type)
		w.Pop()
	}

}
func (w *InfoWriter) writeIdent(flag Flags, typ *ast.Ident) {
	printf("[WRITE] %#+v", typ)
	tracef(1, "[FLAGS] %#+v", flag)
	switch {
	case flag.First().IsInterface() && flag.IsField():
		if def, ok := w.TypeRegistry.pkg.TypesInfo.Types[typ]; ok && def.IsType() {
			if _, ok := def.Type.(*types.Basic); !ok {
				w.TsDefine("%s", typ.Name)
				return
			}
		}
	case flag.First().IsStruct() && flag.IsField():
		if def, ok := w.TypeRegistry.pkg.TypesInfo.Types[typ]; ok && def.IsType() {
			if _, ok := def.Type.(*types.Basic); !ok {
				named := def.Type.(*types.Named)
				if named.Obj().Pkg().Path() == w.pkg.PkgPath {
					w.TsUse(UseStar)
					if flag.IsAnonymous() {
						w.TsDefine("Ref<%s>", named.Obj().Name())
						w.TsIdent("\t\t%s:", named.Obj().Name())
					}
					w.TsImpl("Ref<%s>\n", named.Obj().Name())
					return
				} else {
					w.TsImport(named.Obj().Pkg().Path())
					if flag.IsAnonymous() {
						w.TsDefine("Ref<%s>", named.Obj().Name())
						w.TsIdent("\t\t%s:", named.Obj().Name())
					}
					w.TsImpl("Ref<%s.%s>", named.Obj().Pkg().Path(), named.Obj().Name())
					return
				}
			}
		}
	}
	if len(w.Types) == 1 {
		w.TsIdentType("%s", typ.Name, false)
		return
	}
	p := new(strings.Builder)
	n := 0
	arr := false
	if w.LastN(2).Node == spec.AstArrayType {
		arr = true
	}
	if w.LastN(2).Node == spec.AstStarExpr || (arr && w.LastN(3).Node == spec.AstStarExpr) {
		w.TsUse(UseStar)
		p.WriteString("Ref<")
		n++
	}
	if n > 0 {
		p.WriteString("%s")
		p.WriteString(strings.Repeat(">", n))
		w.TsIdentType(p.String(), typ.Name, arr)
	} else {
		w.TsIdentType("%s", typ.Name, arr)
	}

}

func (w *InfoWriter) writeSelector(flag Flags, typ *ast.SelectorExpr, target types.Object) {
	printf("[WRITE] %#+v >> %#+v", typ, target)
	switch {
	case flag.First().IsInterface() && flag.IsField():
		if target != nil {
			w.TsImport(target.Pkg().Path())
			w.TsDefine("%s.%s", target.Pkg().Name(), typ.Sel.Name)
		} else if id, ok := typ.X.(*ast.Ident); ok {
			w.TsImport(id.Name)
			w.TsDefine("%s.%s", id.Name, typ.Sel.Name)
		} else {
			panic(fmt.Errorf("miss selector: %v %v", typ, target))
		}
	default:
		if target != nil {
			w.TsImport(target.Pkg().Path())
			w.TsImpl("%s.%s", target.Pkg().Name(), typ.Sel.Name)
		} else if id, ok := typ.X.(*ast.Ident); ok {
			w.TsImport(id.Name)
			w.TsImpl("%s.%s", id.Name, typ.Sel.Name)
		} else {
			panic(fmt.Errorf("miss selector: %v %v", typ, target))
		}
	}

}

func (w *InfoWriter) writeStar(flag Flags, typ *ast.StarExpr) {
	printf("[WRITE] %#+v", typ)
	if flag.First().IsStruct() && flag.IsField() {
		if d, ok := w.pkg.TypesInfo.Types[typ.X]; ok && d.IsType() {
			named := d.Type.(*types.Named)

			if named.Obj().Pkg().Path() == w.pkg.PkgPath {
				w.TsUse(UseStar)
				if flag.IsAnonymous() {

					w.TsDefine("Ref<%s>", named.Obj().Name())
					w.TsIdent("\t\t%s:", named.Obj().Name())
				}
				w.TsImpl("Ref<%s>\n", named.Obj().Name())
				return
			} else {
				w.TsImport(named.Obj().Pkg().Path())
				if flag.IsAnonymous() {
					w.TsDefine("Ref<%s>", named.Obj().Name())
					w.TsIdent("\t\t%s:", named.Obj().Name())
				}
				w.TsImpl("Ref<%s.%s>", named.Obj().Pkg().Path(), named.Obj().Name())
				return
			}
		}
		panic("")
	}

	i := w.writeTypeBuffer(flag, typ.X)
	if i.Empty() {
		panic("miss")
	}
	w.Append(i)
	w.Pop()
	i.Free()

}

func (w *InfoWriter) writeFunc(flag Flags, typ *ast.FuncType) {
	printf("[WRITE] %#+v", typ)
	if !(flag.IsResult() || flag.IsArgument()) && (flag.First().IsFunction() || flag.First().IsInterface() || flag.IsMethod() || flag.IsFunction()) {
		if typ.TypeParams.NumFields() > 0 {
			log.Printf("IGNORE Generic Function")
			return
		}
		w.TsImpl("(")
		if typ.Params.NumFields() > 0 {
			w.writeFields(append(flag, FArgument), typ.Params)
		}
		if typ.Results.NumFields() == 0 {
			w.TsImpl("):void")
		} else {
			w.TsImpl("):")
			w.writeFields(append(flag, FResult), typ.Results)
		}
	} else {
		if typ.TypeParams.NumFields() > 0 {
			log.Printf("IGNORE Generic Function")
			return
		}
		w.TsImpl("(")
		if typ.Params.NumFields() > 0 {
			w.writeFields(append(flag, FArgument), typ.Params)
		}
		if typ.Results.NumFields() == 0 {
			w.TsImpl(")=>void")
		} else {
			w.TsImpl(")=>")
			w.writeFields(append(flag, FResult), typ.Results)
		}
	}
}

func (w *InfoWriter) writeArray(flag Flags, typ *ast.ArrayType) {
	printf("[WRITE] %#+v", typ)
	w.writeType(flag, typ.Elt)
	w.Pop()
}

func (w *InfoWriter) writeMap(flag Flags, typ *ast.MapType) {
	printf("[WRITE] %#+v", typ)
	if d, ok := w.pkg.TypesInfo.Types[typ.Key]; ok && d.IsType() && d.Type.String() == "string" {
		if (flag.First().IsStruct() || flag.First().IsInterface()) && !flag.IsField() {
			w.TsImpl("[[key:string]:")
			w.writeType(append(flag, FMapValue), typ.Value)
			w.TsImpl("]")
		} else {
			w.TsImpl("Record<string,")
			w.writeType(append(flag, FMapValue), typ.Value)
			w.TsImpl(">")
		}
		return
	}
	panic(fmt.Errorf("[WRITE] %#+v", typ))
}

func (w *InfoWriter) writeStruct(flag Flags, typ *ast.StructType) {
	printf("[WRITE] %#+v", typ)
	w.writeFields(append(flag, FField), typ.Fields)
}

func (w *InfoWriter) writeInterface(flag Flags, typ *ast.InterfaceType) {
	printf("[WRITE] %#+v", typ)
	switch {
	default:
		w.writeFields(append(flag, FField), typ.Methods)
	}

}

func (w *InfoWriter) writeChan(flag Flags, typ *ast.ChanType) {
	printf("[WRITE] %#+v", typ)
	i := w.writeTypeBuffer(flag, typ.Value)
	if i.Empty() {
		panic("miss chan type")
	}
	switch typ.Dir {
	case ast.SEND:
		w.TsImport("go")
		w.TsImpl("go.ChanSend<")
		w.Append(i)
		w.TsImpl(">")
	case ast.RECV:
		w.TsImport("go")
		w.TsImpl("go.ChanRecv<")
		w.Append(i)
		w.TsImpl(">")
	default:
		w.TsImport("go")
		w.TsImpl("go.Chan<")
		w.Append(i)
		w.TsImpl(">")
	}

}

func (w *InfoWriter) writeEllipsis(flag Flags, typ *ast.Ellipsis) {
	printf("[WRITE] %#+v", typ)
	i := w.writeTypeBuffer(flag, typ.Elt)
	if i.Empty() {
		panic(fmt.Errorf("miss type %v ", typ))
	}
	w.Append(i)
	w.Pop()
	i.Free()

}

//endregion

//region Fields

func (w *InfoWriter) writeFields(flag Flags, fields *ast.FieldList) {
	switch fields.NumFields() {
	case 0:
		printf("none fields exists %+v", fields)
	case 1:
		w.writeField(flag, 1, fields.List[0])
	default:
		if flag.IsResult() {
			w.TsImpl("[")
		}
		switch {
		case (flag.First().IsInterface() || flag.First().IsStruct()) && flag.IsField():
			x0 := w.Typescript.Impl.Len()
			x1 := x0
			for i, field := range fields.List {
				if i > 0 {
					x1 = w.Typescript.Impl.Len()
					if x1 > x0 {
						x0 = x1
						w.TsImpl("\n")
					}
				}
				w.writeField(flag, i, field)
			}
		default:
			for i, field := range fields.List {
				if i > 0 {
					w.TsImpl(",")
				}
				w.writeField(flag, i, field)
			}
		}
		if flag.IsResult() {
			w.TsImpl("]")
		}
	}

}

func (w *InfoWriter) writeField(flag Flags, n int, field *ast.Field) {
	switch len(field.Names) {
	case 0:
		w.writeAnonymousField(flag, n, field)
	case 1:
		w.writeNamedField(flag, n, field.Names[0], field)
	default:
		w.writeNamesField(flag, n, field)
	}
}

func (w *InfoWriter) writeNamesField(flag Flags, n int, field *ast.Field) {
	flag = flag.WithLast(FNamed | FNamedMany)
	i := w.writeTypeBuffer(flag, field.Type)
	if !i.Empty() {
		switch {
		case flag.IsResult():
			for ix := range field.Names {
				if ix > 0 {
					w.TsImpl(",")
				}
				w.Append(i)
			}
		case flag.IsArgument():
			for x, n := range field.Names {
				if x > 0 {
					w.TsImpl(",")
				}
				w.TsIdent("%s", n.Name)
			}
			w.TsImpl(":")
			w.Append(i)
		case flag.IsField():
			x := -1
			for _, n := range field.Names {
				if n.IsExported() {
					x++
				} else {
					continue
				}
				if x > 0 {
					w.TsImpl(",")
				}
				w.TsIdent("%s", n.Name)
			}
			if x > 0 {
				w.TsImpl(":")
				w.Append(i)
				w.TsImpl("\n")
			}
		}
	} else {
		log.Printf("miss type for field :%v", field)
	}
	w.Pop()
	i.Free()
}

func (w *InfoWriter) writeAnonymousField(flag Flags, n int, field *ast.Field) {
	switch {
	case flag.First().IsFunction() && flag.IsArgument():
		log.Printf("should not have anonymous field in function arguments")
		return
	case flag.IsArgument():
		w.TsImpl("v%d:", n)
	}
	w.writeType(flag.WithLast(FAnonymous), field.Type)
	w.Pop()
}

func (w *InfoWriter) writeNamedField(flag Flags, n int, ident *ast.Ident, field *ast.Field) {
	flag = flag.WithLast(FNamed)
	if flag.First().IsStruct() && flag.IsField() && !ident.IsExported() {
		return
	}
	i := w.writeTypeBuffer(flag, field.Type)
	if !i.Empty() {
		switch {
		case flag.IsField():
			if flag.First().IsInterface() {
				w.TsIdent("\t\t%s", ident.Name)
			} else {
				w.TsIdent("\t\t%s:", ident.Name)
			}
			w.Append(i)
			w.TsImpl("\n")
		default:
			x := w.LastIndex(spec.AstEllipsis)
			ex := x == 0 || x == 1
			if flag.IsArgument() {
				if ex {
					w.TsImpl("... %s:", ident.Name)
				} else {
					w.TsImpl("%s:", ident.Name)
				}
			}
			w.Append(i)
			if ex {
				w.TsImpl("[]")
			}
		}
	} else {
		log.Printf("miss type for field :%+v", field)
	}
	w.Pop()
	i.Free()
}

// endregion

//region Type

func (w *InfoWriter) writeType(flag Flags, typ ast.Expr) {
	t := w.lookupType(typ)
	n := len(t)
	switch n {
	case 0:
		log.Printf("miss type %v", typ)
	default:
		w.Push(t)
		node := t[len(t)-1]
		switch node.Node {
		case spec.AstIdent:
			w.writeIdent(flag, node.Ident)
		case spec.AstSelectorExpr:
			w.writeSelector(flag, node.Selector, node.SelectorTarget)
		case spec.AstStarExpr:
			w.writeStar(flag, node.Star)
		case spec.AstFuncType:
			w.writeFunc(flag, node.Func)
		case spec.AstArrayType:
			w.writeArray(flag, node.Array)
		case spec.AstMapType:
			w.writeMap(flag, node.Map)
		case spec.AstStructType:
			w.writeStruct(flag, node.Struct)
		case spec.AstInterfaceType:
			w.writeInterface(flag, node.Interface)
		case spec.AstChanType:
			w.writeChan(flag, node.Chan)
		case spec.AstEllipsis:
			w.writeEllipsis(flag, node.Ellipsis)
		default:
			log.Printf("miss node type %v", node.Node)
		}
	}
}
func (w *InfoWriter) writeTypeBuffer(flag Flags, typ ast.Expr) (define *ModuleWriter) {
	dw := new(ModuleWriter)
	dw.Init()
	t := w.lookupType(typ)
	n := len(t)
	switch n {
	case 0:
		log.Printf("miss type %v", typ)
	default:
		w.Push(t)
		ow := w.ModuleWriter
		w.ModuleWriter = dw
		defer func() {
			w.ModuleWriter = ow
		}()
		node := t[len(t)-1]
		switch node.Node {
		case spec.AstIdent:
			w.writeIdent(flag, node.Ident)
		case spec.AstSelectorExpr:
			w.writeSelector(flag, node.Selector, node.SelectorTarget)
		case spec.AstStarExpr:
			w.writeStar(flag, node.Star)
		case spec.AstFuncType:
			w.writeFunc(flag, node.Func)
		case spec.AstArrayType:
			w.writeArray(flag, node.Array)
		case spec.AstMapType:
			w.writeMap(flag, node.Map)
		case spec.AstStructType:
			w.writeStruct(flag, node.Struct)
		case spec.AstInterfaceType:
			w.writeInterface(flag, node.Interface)
		case spec.AstChanType:
			w.writeChan(flag, node.Chan)
		case spec.AstEllipsis:
			w.writeEllipsis(flag, node.Ellipsis)
		default:
			log.Printf("miss node type %v", node.Node)
		}
	}
	return dw
}

//endregion
