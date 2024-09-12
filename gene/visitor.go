package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ZenLiuCN/fn"
	. "github.com/ZenLiuCN/go-inspect"
	"go/ast"
	"go/format"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode"
)

var ctxVisitor = FnTypeVisitor[*TypeInspector[*context], *context]{
	FnVisitConst: func(I InspectorX, d Dir, name string, e *types.Const, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTConst, I.TypePath, nil)
			c.name = name
			c.d.Indent(c.m.Indents() + 1)
		case EXT:
			defer c.Exit(KTConst, I.TypePath, nil)
			if c.generic {
				//TODO handle generic Const
				c.reset()
				return false
			}
			if !c.typeAlias {
				c.mind().mf("//%s", e.Val().String())
				c.Constants[name] = e
			} else {
				c.AliasConst[name] = e
			}
			c.mind().mf("export const %s:", name)
			c.declared()
			c.addEntry(name, "%s.%s", c.pkg.Types.Name(), name)
			c.reset()
		}
		return true
	},
	FnVisitFunc: func(I InspectorX, d Dir, name string, e *types.Func, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTFunc, I.TypePath, nil)
			c.name = name
			c.d.Indent(c.m.Indents() + 1)
		case EXT:
			defer c.Exit(KTFunc, I.TypePath, nil)
			if c.generic || c.ignore {
				//TODO handle generic function
				c.reset()
				return false
			}
			c.mind().mf("export function %s", Camel(name))
			c.Functions[name] = e
			c.declared().mn()
			c.addEntry(Camel(name), "%s.%s", c.pkg.Types.Name(), name)
			c.reset()
		}
		return true
	},
	FnVisitTypeName: func(I InspectorX, d Dir, name string, e *types.TypeName, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTType, I.TypePath, nil)
			c.name = name
			c.d.Indent(c.m.Indents() + 1)
		case EXT:
			defer c.Exit(KTType, I.TypePath, nil)
			if c.generic {
				//TODO handle generic
				c.reset()
				return false
			}
			if c.typeAlias {
				c.AliasType[name] = e
				c.mind().mf("export type %s=%s", name, c.d.String())
			} else {
				c.mind().mf("export interface %s", name)
				if c.isStruct {
					c.Structs[name] = e
				} else {
					c.Interfaces[name] = e
				}
				c.faceIntercept()
				if len(c.extended) > 0 {
					c.mf(" extends %s", strings.Join(c.extended.Values(), ","))
				}
				c.mf("{").
					mn().declared().
					mind().mf("}")
			}
			c.reset()
		}
		return true
	},
	FnVisitVar: func(I InspectorX, d Dir, name string, e *types.Var, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTVar, nil, nil)
			c.name = name
			c.d.Indent(c.m.Indents() + 1)
		case EXT:
			defer c.Exit(KTVar, nil, nil)
			if c.generic {
				//TODO handle generic Const
				c.reset()
				return false
			}
			c.Variables[name] = e
			c.mind().mf("export const %s:", name)
			c.declared()
			c.addEntry(name, "%s.%s", c.pkg.Types.Name(), name)
			c.reset()
		}
		return true
	},
	FnVisitTypeTuple: func(I InspectorX, d Dir, o types.Object, x *types.Tuple, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTuple, I.TypePath, seen)
			switch m := I.LNth(2); m {
			case KMParam:
				c.df("(")
				c.pc.PushSitu(x.Len())
			case KMResult:
				c.pc.PushSitu(x.Len())
				switch x.Len() {
				case 0:
				case 1:
				default:
					c.df("[")
				}
			}
		case EXT:
			defer c.Exit(KTuple, I.TypePath, seen)
			switch m := I.LNth(2); m {
			case KMParam:
				if !c.funcAlias && (I.LNth(4) == KMMethod || (I.Len() <= 4 && (I.First() == KTType || I.First() == KTFunc))) {
					c.df("):")
				} else {
					c.df(")=>")
				}
				c.pc.PopSitu()
			case KMResult:
				switch x.Len() {
				case 0:
					c.df("void")
				case 1:
				default:
					c.df("]")
				}
				c.pc.PopSitu()
			}
		}
		return true
	},
	FnVisitTypeVar: func(I InspectorX, d Dir, o types.Object, x *types.Var, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KVar, I.TypePath, seen)
			switch {
			//!! field writer
			case I.First() == KTType && I.EndWith(KStruct, KMField):
				c.buf(c.field)
				c.field = x.Name()
			default:
				// !! KMParam,KTuple,KVar
				switch m := I.LNth(3); m {
				case KMParam:
					c.param = x.Name()
				case KMField:
					c.field = x.Name()
				}
			}
			c.subWriter(I.TypePath, x.Type())
		case EXT:
			defer c.Exit(KVar, I.TypePath, seen)
			switch {
			//!! field writer
			case c.field != "" && I.First() == KTType && I.EndWith(KStruct, KMField):
				switch {
				//!! only expose field type should write
				case c.d.Buffer().Len() > 0:
					c.mergeWriter(I.TypePath, x.Type(), func(w Writer) {
						if c.ignore {
							c.ignore = false
							return
						}
						c.dn().df("%s:%s", Camel(c.field), w.String())
					})
				//!! ignore field
				default:
					c.mergeWriter(I.TypePath, x.Type(), assertEmpty)
				}
				c.field = c.unBuf().(string)
			//!! KParam KTuple KVar
			default:
				switch I.LNth(3) {
				case KMParam:
					c.mergeWriter(I.TypePath, x.Type(), func(w Writer) {
						switch {
						case c.pc.Last() == 1 && c.variadic.Last():
							if c.param == "" {
								c.df("...v%d:%s", c.pc.Last(), w.String())
							} else {
								c.df("...%s:%s", Safe(x.Name()), w.String())
							}
						default:
							if c.param == "" {
								c.df("v%d:%s", c.pc.Last(), w.String())
							} else {
								c.df("%s:%s", Safe(x.Name()), w.String())
							}
						}
					})
					if c.pc.Last() > 1 {
						c.df(",")
					}
					c.param = ""
					c.pcsub()
				case KMResult:
					c.mergeWriter(I.TypePath, x.Type(), func(w Writer) {
						if c.pc.Last() > 1 {
							c.df("%s,", w.String())
						} else {
							c.df("%s", w.String())
						}
					})
					c.pcsub()
				case KMField:
					panic(x)
				/*	c.mergeWriter(I.TypePath, x.Type(), func(w Writer) {
						c.dn().df("%s:%s", w.String())
					})
					c.field = ""*/
				default:
					c.mergeWriter(I.TypePath, x.Type(), func(w Writer) {
						c.dn().df("%s", w.String())
					})
				}
			}

		}
		return true
	},
	FnVisitTypeFunc: func(I InspectorX, d Dir, o types.Object, x *types.Func, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KFunc, I.TypePath, seen)
			c.subWriter(I.TypePath, x.Type())
			c.method = x.Name()
		case EXT:
			defer c.Exit(KFunc, I.TypePath, seen)
			c.mergeWriter(I.TypePath, x.Type(), func(t Writer) {
				if c.ignore || !t.NotEmpty() {
					c.ignore = false
					return
				}
				//!! Func declare
				c.di().dn().df("%s%s", Camel(x.Name()), t.String()).dd()
			})
			c.method = ""
		}
		return true
	},
	FnVisitTypeBasic: func(I InspectorX, d Dir, o types.Object, x *types.Basic, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KBasic, I.TypePath, seen)
			if I.Equals(KTType, KBasic) {
				c.df("%s", c.identType(x.Name(), false, true))
				c.typeAlias = true
				return false
			}
			arr := c.arrayLen.Last()
			switch arr {
			case 0:
				c.df("%s", c.identType(x.Name(), false, c.extending.Last()))
			case -1:
				c.df("%s", c.identType(x.Name(), true, c.extending.Last()))
			default:
				c.df("%s/*%d*/", c.identType(x.Name(), false, c.extending.Last()), arr)
			}
		case EXT:
			defer c.Exit(KBasic, I.TypePath, seen)
			if I.LNth(2) == KMMapKey {
				c.df(",")
			}
		}
		return true
	},
	FnVisitTypeMap: func(I InspectorX, d Dir, o types.Object, x *types.Map, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KMap, I.TypePath, seen)
			switch {
			case I.Equals(KTType, KNamed, KMap) || I.Equals(KTType, KMap):
				c.mapAlias = true
				c.extending.PushSitu(true)
				c.subWriter(I.TypePath, x)
			default:
				c.extending.PushSitu(false)
				c.subWriter(I.TypePath, x)
			}
		case EXT:
			defer c.Exit(KMap, I.TypePath, seen)
			defer c.extending.PopSitu()
			switch {
			case c.mapAlias:
				c.mapAlias = false
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					c.use("map")
					c.extends("map<%s>", w.String())
				})
				return true
			default:
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					c.use("map")
					c.df("map<%s>", w.String())
				})
			}
		}
		return true
	},
	FnVisitTypeArray: func(I InspectorX, d Dir, o types.Object, x *types.Array, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KArray, I.TypePath, seen)
			c.arrayLen.PushSitu(x.Len())
			if I.Equals(KTType, KNamed, KArray) || I.Equals(KTType, KArray) {
				c.extending.PushSitu(true)
			}
			c.subWriter(I.TypePath, x)
		case EXT:
			defer c.Exit(KArray, I.TypePath, seen)
			defer c.arrayLen.PopSitu()
			switch {
			case I.Equals(KTType, KNamed, KArray) || I.Equals(KTType, KArray):
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					if c.path.Last() == KBasic {
						c.use("Alias")
						c.extends("Alias<%s>", w.String())
					} else {
						c.extends("Array<%s/*%d*/>", w.String(), x.Len())
					}
				})
				c.extending.PopSitu()
			case c.extending.Last():
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					if c.path.Last() == KBasic {
						c.df("%s", w.String())
					} else {
						c.df("Array<%s/*%d*/>", w.String(), x.Len())
					}
				})
			default:
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					if c.path.Last() == KBasic {
						c.df("%s", w.String())
					} else {
						c.df("%s[/*%d*/]", w.String(), x.Len())
					}
				})
			}
		}
		return true
	},
	FnVisitTypeStruct: func(I InspectorX, d Dir, o types.Object, x *types.Struct, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KStruct, I.TypePath, seen)
			switch {
			//!! embedded (first level)
			case c.field != "" && I.StartWith(KTType, KNamed, KStruct, KMField) && I.EndWith(KStruct) && I.Len() == 6:
				if x.NumFields() == 0 {
					c.use("Nothing")
					c.df("Nothing")
					return false
				}
				c.buf(c.field)
				c.field = ""
				c.subWriter(I.TypePath, x)
				c.di()
			//!! embedded (more level)
			case c.field != "" && c.buffed.Len() > 0:
				if x.NumFields() == 0 {
					return false
				}
				c.buf(c.field)
				c.field = ""
				c.subWriter(I.TypePath, x)
				c.di()
			//!! define
			case I.Equals(KTType, KNamed, KStruct):
				if x.NumFields() == 0 {
					c.use("Alias")
					c.use("Nothing")
					c.extends("Alias<Nothing>")
					return false
				}
				c.subWriter(I.TypePath, x)
				c.di()
				//!! anonymous empty struct
			//!! anonymous nothing struct without sub-writer
			case x.NumFields() == 0:
				c.use("Nothing")
				c.df("Nothing")
				return false
			default:
				c.subWriter(I.TypePath, x)
			}
		case EXT:
			defer c.Exit(KStruct, I.TypePath, seen)
			switch {
			//!! anonymous nothing struct, without sub-writer
			case x.NumFields() == 0:

			//!! embedded first level
			case c.field == "" && I.StartWith(KTType, KNamed, KStruct, KMField) && I.EndWith(KStruct) && I.Len() == 6:
				c.field = c.unBuf().(string)
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					if !w.NotEmpty() { // ** ignore if nothing
						return
					}
					c.use("Struct")
					if c.path.Last() == KArray || c.path.Last() == KSlice {
						x := w.String()
						c.df("Struct<{").dn().df("%s", x[:len(x)-2]).dn().df("}>[]")
					} else {
						c.df("Struct<{").dn().df("%s", w.String()).dn().df("}>")
					}
				})
			//!! embedded more level
			case c.field == "" && c.buffed.Len() > 1:
				c.field = c.unBuf().(string)
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					if !w.NotEmpty() { // ** ignore if nothing
						return
					}
					c.use("Struct")
					if c.path.Last() == KArray || c.path.Last() == KSlice {
						x := w.String()
						c.df("Struct<{").dn().df("%s", x[:len(x)-2]).dn().df("}>[]")
					} else {
						c.df("Struct<{").dn().df("%s", w.String()).dn().df("}>")
					}

				})
			//!! define
			case I.Equals(KTType, KNamed, KStruct):
				switch {
				case strings.Contains(c.name, "Err"):
					c.use("Struct")
					c.extends("Struct<%s>", c.name)
					c.extends("Error")
					c.isStruct = true
					c.mergeWriter(I.TypePath, x, func(w Writer) {
						c.df("%s", w.String())
					})
				default:
					c.use("Struct")
					c.extends("Struct<%s>", c.name)
					c.isStruct = true
					for _, face := range I.Implements(o.Type()) {
						if face.Pkg == c.pkg.Types {
							if !c.extended.Exists(face.Name) {
								c.extends(face.Name)
							}
						} else {
							name := fmt.Sprintf("%s.%s", face.Pkg.Name(), face.Name)
							if !c.extended.Exists(name) && !c.ignored(face.Pkg) {
								c.imported(face.Pkg, face.Pkg.Name())
								c.extends(name)
							}
						}
					}
					c.mergeWriter(I.TypePath, x, func(w Writer) {
						c.df("%s", w.String())
					})
				}
			default:
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					c.d.Append(w)
				})
			}
		}
		return true
	},
	FnVisitTypeUnion: func(I InspectorX, d Dir, o types.Object, x *types.Union, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KUnion, I.TypePath, seen)

		case EXT:
			defer c.Exit(KUnion, I.TypePath, seen)

		}
		return true
	},
	FnVisitTypeSignature: func(I InspectorX, d Dir, o types.Object, x *types.Signature, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KSignature, I.TypePath, seen)
			c.variadic = c.variadic.Push(x.Variadic())
			if I.Equals(KTType, KNamed, KFunc, KSignature) || I.Equals(KTType, KNamed, KSignature) {
				c.funcAlias = true //!! not need Extending for function
			}
			c.subWriter(I.TypePath, x)
		case EXT:
			defer c.Exit(KSignature, I.TypePath, seen)
			c.variadic, _ = c.variadic.Pop()
			if c.funcAlias {
				c.funcAlias = false
				c.use("Alias")
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					c.extends("Alias<%s>", w.String())
				})
				return true
			}
			c.mergeWriter(I.TypePath, x, func(w Writer) {
				c.df(w.String())
			})
			//!! check error
			if c.method == "Error" || c.method == "Close" {
				return true
			}
			c.checkErrorResult()
		}
		return true
	},
	FnVisitTypeParam: func(I InspectorX, d Dir, o types.Object, x *types.TypeParam, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTypeParam, I.TypePath, seen)
			c.generic = true
		case EXT:
			defer c.Exit(KTypeParam, I.TypePath, seen)

		}
		return true
	},
	FnVisitTypePointer: func(I InspectorX, d Dir, o types.Object, x *types.Pointer, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KPointer, I.TypePath, seen)
			switch {
			//!! extends
			case I.Equals(KTType, KNamed, KPointer):
				c.extending.PushSitu(true)
			default:
			}
			c.subWriter(I.TypePath, x)
		case EXT:
			defer c.Exit(KPointer, I.TypePath, seen)
			switch {
			case I.Equals(KTType, KNamed, KPointer):
				c.extending.PopSitu()
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					c.use("Ref")
					c.extends("Ref<%s>", w.String())
				})
			default:
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					c.use("Ref")
					c.df("Ref<%s>", w.String())
				})
			}
		}
		return true
	},
	FnVisitTypeSlice: func(I InspectorX, d Dir, o types.Object, x *types.Slice, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KSlice, I.TypePath, seen)
			defer c.arrayLen.PushSitu(-1)
			if I.Equals(KTType, KNamed, KSlice) || I.Equals(KTType, KSlice) {
				c.extending.PushSitu(true)
			}
			c.subWriter(I.TypePath, x)
		case EXT:
			defer c.Exit(KSlice, I.TypePath, seen)
			defer c.arrayLen.PopSitu()
			switch {
			case I.Equals(KTType, KNamed, KSlice) || I.Equals(KTType, KSlice):
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					if c.path.Last() == KBasic {
						c.use("Alias")
						c.extends("Alias<%s>", w.String())
					} else {
						c.extends("Array<%s>", w.String())
					}
				})
				c.extending.PopSitu()
			case c.extending.Last():
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					if c.path.Last() == KBasic {
						c.df("%s", w.String())
					} else {
						c.df("Array<%s>", w.String())
					}
				})
			default:
				c.mergeWriter(I.TypePath, x, func(w Writer) {
					if c.path.Last() == KBasic {
						c.df("%s", w.String())
					} else {
						c.df("%s[]", w.String())
					}
				})
			}
		}
		return true
	},
	FnVisitTypeInterface: func(I InspectorX, d Dir, o types.Object, x *types.Interface, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KInterface, I.TypePath, seen)
			switch {
			//!! declared
			case I.Equals(KTType, KNamed, KInterface) || I.Equals(KTType, KInterface):
				c.isStruct = false
				if x.NumMethods() == 0 && x.NumEmbeddeds() == 0 {
					return false
				}
				return true
			//!! alias
			case I.Len() == 3 && I.LNth(2) == KNamed:
				c.use("Proto")
				c.extends("Proto<%s>", c.name)
				return false
			default:
				if x.NumMethods() == 0 {
					c.df("any") //!! any == interface{}
				}
			}
		case EXT:
			defer c.Exit(KInterface, I.TypePath, seen)
		}
		return true
	},
	FnVisitTypeChan: func(I InspectorX, d Dir, o types.Object, x *types.Chan, seen *Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KChan, I.TypePath, seen)
			c.subWriter(I.TypePath, x)
		case EXT:
			defer c.Exit(KChan, I.TypePath, seen)
			c.mergeWriter(I.TypePath, x, func(w Writer) {
				c.imported(nil, "go")
				switch x.Dir() {
				case types.SendOnly:
					c.df("go.ChanSend<%s>", w.String())
				case types.RecvOnly:
					c.df("go.ChanRecv<%s>", w.String())
				default:
					c.df("go.Chan<%s>", w.String())
				}
			})
		}
		return true
	},
	FnVisitTypeNamed: func(I InspectorX, d Dir, o types.Object, x *types.Named, seen *Types, c *context) bool {
		name := x.Obj().Name()
		switch d {
		case ENT:
			defer c.Enter(KNamed, I.TypePath, seen)
			switch {
			//!! typeDefine
			case c.path.Len() == 1 && c.name == name && x.Obj().Pkg() == c.pkg.Types:
				return true
			//!! first alias
			case c.path.Len() == 1:
				if c.ignored(x.Obj().Pkg()) {
					c.ignore = true
					return false
				}
				c.typeAlias = true
				c.decodeNamed(I, x, c.df)
				return false
			//!! embedded for interface
			case I.First() == KTType && I.LNth(2) == KMEmbedded:
				if c.ignored(x.Obj().Pkg()) {
					c.ignore = true
					return false
				}
				c.decodeNamed(I, x, c.extends)
				return false
			//!! unexported
			case name != "error" && !ast.IsExported(name):
				switch {
				case I.LastIndex(KSignature) > 3 || I.StartWith(KTType, KSignature) || I.StartWith(KTType, KNamed, KSignature):
					c.use("reserved")
					c.df("reserved")
					return false
				//!! field with unexported but with exported methods
				case I.Len() <= 7 && c.field != "" && c.method == "" && x.NumMethods() > 0:
					c.subWriter(I.TypePath, x)
					return true
				default:
					c.ignore = true
					return false
				}
			case c.ignored(x.Obj().Pkg()):
				c.ignore = true
				return false
			default:
				c.decodeNamed(I, x, c.df)
				return false

			}
		case ONC:
			switch {
			//!! alias
			case I.Len() == 2 && I.First() == KTType && c.path.Last() == KBasic:
				c.use("Alias")
				c.extends("Alias<%s>", c.d.String())
				c.d.Reset()
			}
		case EXT:
			defer c.Exit(KNamed, I.TypePath, seen)
			if !c.ignore {
				switch {
				//!! field with unexported struct but with exported methods
				case I.Len() <= 7 && name != "error" && !ast.IsExported(name) && c.method == "" && c.field != "" && x.NumMethods() > 0:
					c.mergeWriter(I.TypePath, x, func(w Writer) {
						c.df("{").dn().df("%s", w.String()).dn().df("}")
					})
				case I.LNth(2) == KMMapKey:
					c.df(",")
				}
			}

		}
		return true
	},
}

type Bool = fn.Stack[bool]
type Writers = fn.StackAny[WriterInfo]
type WriterInfo struct {
	path   fn.Stack[TypeKind]
	writer Writer
	types  types.Type
}
type Any = fn.Stack[any]
type Int = fn.Stack[int]
type Int64 = fn.Stack[int64]
type state struct {
	path   TypePath
	name   string
	method string
	field  string
	param  string

	typeAlias bool
	funcAlias bool
	mapAlias  bool
	generic   bool
	isStruct  bool
	ignore    bool // some part of current field type or method types have ignored
	buffed    Any
	arrayLen  Int64
	variadic  Bool
	extending Bool
	pc        Int
	tmp       Writers
}

func (c *state) buf(v any) {
	if c.buffed == nil {
		c.buffed = fn.Stack[any]{}
	}
	c.buffed.PushSitu(v)
}
func (c *state) unBuf() any {
	if c.buffed.Len() == 0 {
		panic("empty buffered")
	}
	return c.buffed.PopSitu()
}
func (c *state) reset() {
	c.path.CleanSitu()
	c.name = ""
	c.method = ""
	c.field = ""
	c.param = ""

	c.typeAlias = false
	c.funcAlias = false
	c.mapAlias = false
	c.generic = false
	c.isStruct = false
	c.ignore = false

	c.buffed.CleanSitu()
	c.arrayLen.CleanSitu()
	c.extending.CleanSitu()
	c.variadic.CleanSitu()
	c.pc.CleanSitu()
	c.tmp.CleanSitu()

}

/*
	func (s *state) resetArrayed() {
		if s.arrayed >= s.path.Len() {
			s.arrayed = 0
		}
	}
*/
func (c *state) pcsub() {
	if c.pc.Empty() {
		return
	}
	c.pc[c.pc.MaxIdx()]--
}

func (c *state) Enter(k TypeKind, path TypePath, seen *Types) {
	if debug {
		log.Printf("ENTER[%s]\t%s:%s[%d]\tΔ%v\tΔ%v", k, c.name, anyOf(c.field, c.method, c.param), c.pc.Last(), path, c.path)
	}
	switch k {
	case KSlice, KArray:
	default:
		c.arrayLen.Push(0)
	}
	c.path.PushSitu(k)

}
func (c *state) Exit(k TypeKind, path TypePath, seen *Types) {
	if debug {
		log.Printf("EXIT[%s]\t%s:%s[%d]\tΔ%v\tΔ%v", k, c.name, anyOf(c.field, c.method, c.param), c.pc.Last(), path, c.path)
	}
	if c.tmp.Len() > 0 && c.tmp.Last().path.Equals(path...) {
		panic(fmt.Errorf("not reset writer %s", path))
	}
	switch k {
	case KVar, KMap, KTuple, KFunc, KSignature, KSlice, KPointer, KArray: //!! branch clean
		if path.Len() > 2 {
			kx := c.path.LastIndex(k)
			if kx == -1 {
				return
			}
			c.path = c.path[:kx]
		}
	default:
		if path.Len() < 3 {
			//!! remove children branch
			kx := c.path.LastIndex(k)
			if kx == -1 {
				return
			}
			c.path = c.path[:kx]
		}
	}
	c.arrayLen.PopSitu()
}

type context struct {
	root          string
	Functions     map[string]*types.Func
	Structs       map[string]*types.TypeName
	Interfaces    map[string]*types.TypeName
	Constants     map[string]*types.Const
	AliasConst    map[string]*types.Const
	AliasType     map[string]*types.TypeName
	Variables     map[string]*types.Var
	pkg           *packages.Package
	uses          fn.HashSet[string]
	imports       fn.HashSet[Imported]
	entry         map[string]string
	overrideEntry map[string]string
	wrapper       Writer
	m             Writer
	extended      fn.HashSet[string]
	d             Writer
	FnTypeVisitor[*TypeInspector[*context], *context]

	state
	errors  bool //reduce function end subWriter error
	ignores []string
}
type Imported struct {
	Pkg  *types.Package
	Name string
}

func (c *context) subWriter(path TypePath, t types.Type) {
	if debug {
		log.Printf("acquire tmp writers: %d", c.tmp.Len())
	}
	if c.tmp == nil {
		c.tmp = fn.StackAny[WriterInfo]{}
	}
	n := c.d.Indents()
	if c.tmp.Len() > 0 && (path.Len() <= c.tmp.Last().path.Len() || !path.StartWith(c.tmp.Last().path...)) {
		panic(fmt.Errorf("miss push writer: %s vs %s", path, c.tmp.Last().path))
	}
	c.tmp.PushSitu(WriterInfo{
		path:   slices.Clone(path),
		types:  t,
		writer: c.d,
	})
	c.d = GetWriter()
	c.d.Indent(n)
}
func (c *context) mergeWriter(path TypePath, types types.Type, act func(w Writer)) {
	if debug {
		log.Printf("free tmp writers: %d", c.tmp.Len())
	}
	if c.tmp.Len() == 0 {
		panic("not have temp writer")
	}
	t := c.d
	x := c.tmp.PopSitu()
	if x.types != types || !x.path.Equals(path...) {
		panic("miss remove writer")
	}
	c.d = x.writer
	if act != nil {
		act(t)
	}
	t.Release()
}
func (c *context) decodeNamed(i InspectorX, x *types.Named, f func(string, ...any) *context) bool {
	if x.Obj().Exported() || x.Obj().Name() == "error" {
		if x.Obj().Name() == "error" {
			if c.pc.Len() > 0 {
				c.use("error")
				f("error")
			} else {
				c.use("GoError")
				f("GoError")
			}
		} else if x.Obj().Pkg().Path() == c.pkg.Types.Path() {
			f("%s", x.Obj().Name())
		} else {
			c.imported(x.Obj().Pkg(), x.Obj().Pkg().Name())
			f("%s.%s", x.Obj().Pkg().Name(), x.Obj().Name())
		}
		return true
	} else { //!! unexported type
		is := i.Implements(x)
		if len(is) > 0 {
			slices.SortFunc(is, func(a, b Face) int {
				return a.Interface.NumMethods() - b.Interface.NumMethods()
			})
			b := new(bytes.Buffer)
			for _, face := range is {
				if face.Pkg.Name() == c.pkg.Types.Name() {
					if b.Len() > 0 {
						b.WriteByte('&')
					}
					b.WriteString(face.Name)
				}
			}
			f(b.String())
			if b.Len() > 0 {
				return true
			}
			c.ignore = true
			return false
		}
	}

	return false
}

func (c *context) init() {
	c.state = state{
		path:      TypePath{},
		name:      "",
		method:    "",
		field:     "",
		param:     "",
		typeAlias: false,
		funcAlias: false,
		mapAlias:  false,
		generic:   false,
		isStruct:  false,
		ignore:    false,
		buffed:    Any{},
		arrayLen:  Int64{},
		extending: Bool{},
		variadic:  Bool{},
		pc:        Int{},
		tmp:       Writers{},
	}
	c.Functions = map[string]*types.Func{}
	c.Structs = map[string]*types.TypeName{}
	c.Interfaces = map[string]*types.TypeName{}
	c.Constants = map[string]*types.Const{}
	c.AliasConst = map[string]*types.Const{}
	c.AliasType = map[string]*types.TypeName{}
	c.Variables = map[string]*types.Var{}

	c.uses = fn.HashSet[string]{}
	c.imports = fn.HashSet[Imported]{}
	c.extended = fn.HashSet[string]{}
	c.entry = map[string]string{}
	c.overrideEntry = map[string]string{}
	c.wrapper = GetWriter()
	c.m = GetWriter()
	c.d = GetWriter()
	c.m.Indent(1)
	c.FnTypeVisitor = ctxVisitor
}
func (c *context) identType(s string, array, ext bool) string {
	s = strings.TrimPrefix(s, "untyped ")
	switch s {
	case "byte", "uint8":
		if array {
			return "Uint8Array"
		} else {
			c.uses.Add(s)
			return s
		}
	case "uint16", "uint32":
		if array {
			return fmt.Sprintf("U%sArray", s[1:])
		} else {
			c.uses.Add(s)
			return s
		}
	case "int8", "int16", "int32":
		if array {
			return fmt.Sprintf("I%sArray", s[1:])
		} else {
			c.uses.Add(s)
			return s
		}
	case "error":
		if array {
			c.uses.Add(s)
			if ext {
				return "Array<error>"
			}
			return s + "[]"
		} else {
			c.uses.Add(s)
			return s
		}
	case "Pointer": //!! unsafe.Pointer is a Basic type
		if array {
			c.uses.Add(s)
			if ext {
				return "Array<Pointer>"
			}
			return "Pointer[]"
		} else {
			c.uses.Add(s)
			return s
		}
	case "string":
		if array {
			if ext {
				return "Array<string>"
			}
			return "string[]"
		} else {
			if ext {
				c.uses.Add("Alias")
				return "Alias<string>"
			}
			return "string"
		}
	case "any":
		if array {
			if ext {
				return "Array<any>"
			}
			return "any[]"
		} else {
			return "any"
		}
	default:
		if array {
			if !unicode.IsUpper(rune(s[0])) {
				c.uses.Add(s)
			}
			if ext {
				return fmt.Sprintf("Array<%s>", s)
			}
			return fmt.Sprintf("%s[]", s)
		} else {
			if !unicode.IsUpper(rune(s[0])) {
				c.uses.Add(s)
			}
			return s
		}

	}
}

func (c *context) imported(pkg *types.Package, s string) {
	c.imports.Add(Imported{
		Pkg:  pkg,
		Name: s,
	})
}

func (c *context) addEntry(name string, format string, args ...any) {
	c.entry[name] = fmt.Sprintf(format, args...)
}

func (c *context) declared() *context {
	c.m.Append(c.d)
	c.d.Buffer().Reset()
	return c
}

func (c *context) use(s string) {
	c.uses.Add(s)
}
func (c *context) extends(format string, args ...any) *context {
	c.extended.Add(fmt.Sprintf(format, args...))
	return c
}

func (c *context) di() *context {
	c.d.Indent(c.d.Indents() + 1)
	return c
}
func (c *context) dd() *context {
	c.d.Indent(c.d.Indents() - 1)
	return c
}
func (c *context) dn() *context {
	c.d.IndLF()
	return c
}
func (c *context) dl() *context {
	c.d.LF()
	return c
}
func (c *context) df(format string, args ...any) *context {
	c.d.Format(format, args...)
	return c
}

func (c *context) mf(s string, args ...any) *context {
	c.m.Format(s, args...)
	return c
}

func (c *context) mn() *context {
	c.m.LF()
	return c
}
func (c *context) mind() *context {
	c.m.IndLF()
	return c
}
func (c *context) reset() {
	c.state.reset()
	c.extended.Clear()
	c.d.Reset()
}
func (c *context) faceIntercept() {
	bin := c.d.Buffer().Bytes()
	if !c.extended.Exists("Closer") &&
		!c.extended.Exists("io.Closer") &&
		bytes.Contains(bin, []byte("close():error")) &&
		c.name != "Closer" {
		if c.pkg.Types.Name() == "io" {
			c.extends("Closer")
		} else {
			c.imported(nil, "io")
			c.extends("io.Closer")
		}
	}
	if !c.extended.Exists("error") &&
		bytes.Contains(bin, []byte("error():string")) &&
		c.name != "error" {
		c.use("GoError")
		c.extends("GoError")
	}
}

func (c *context) ignored(pkg *types.Package) bool {
	if pkg == nil || len(c.ignores) == 0 {
		return false
	}
	return slices.Contains(c.ignores, pkg.Path())
}

var (
	bytesError         = []byte(",error]")
	byteVoidFuncResult = []byte("=>void/*error*/")
	byteErrorComment   = []byte("/*error*/")
	byteArrow          = []byte("=>")

	byteResult    = []byte("):")
	byteError     = []byte("error")
	byteErrorFunc = []byte("error():error")
	byteCloseFunc = []byte("close():error")
)

func (c *context) checkErrorResult() {
	if !c.errors {
		b := c.d.Bytes()
		if bytes.HasSuffix(b, bytesError) {
			cnt := bytes.Count(b[bytes.LastIndexByte(b, ')'):], []byte{','})
			if cnt == 1 {
				x := bytes.LastIndex(b, byteResult) // !! check ): first
				if x == -1 {
					x = bytes.LastIndex(b, byteArrow)
				}
				if x == -1 {
					panic("not a function ")
				}
				pre := b[0 : x+2]
				last := b[x+2:]

				y := bytes.LastIndexByte(last, ',')
				b = append(pre, last[1:y]...)
				c.d.Reset()
				c.d.Buffer().Write(b)
			} else {
				b = append(b[:len(b)-len(bytesError)], ']')
				c.d.Reset()
				c.d.Buffer().Write(b)
			}
		} else if bytes.HasSuffix(b, byteError) && !bytes.HasSuffix(b, byteErrorFunc) && !bytes.HasSuffix(b, byteCloseFunc) {
			x := bytes.LastIndex(b, byteArrow)
			y := bytes.LastIndex(b, byteResult)
			if x == -1 || x < y || y > 0 { //fix: for parameter contains lambda
				x = y //bytes.LastIndex(b, byteResult)
				if x != -1 {
					b = append(b[0:x+1], byteErrorComment...)
				}
			} else {
				b = append(b[0:x], byteVoidFuncResult...)
			}
			if x == -1 {
				panic("not a function ")
			}

			c.d.Reset()
			c.d.Buffer().Write(b)
		}

	}
}
func (c *context) flush(g *Generator) (err error) {
	x := GetWriter()
	defer x.Release()
	module := ""
	cond := ""
	if g.suffix != "" {
		cond = "_" + g.suffix
	}
	pkgPath := c.pkg.Types.Path()
	filePath := pkgPath
	if strings.IndexByte(filePath, '.') >= 0 {
		filePath = strings.ReplaceAll(filePath, ".", "/")
	}
	pkg := c.pkg.Types.Name()
	dts := strings.Join(fn.SliceMapping(strings.Split(strings.ReplaceAll(c.pkg.Types.Path(), ".", "/"), "/"), strings.ToLower), "_") + cond + ".d.ts"
	mod := dts[:len(dts)-5] + ".go"
	modTest := dts[:len(dts)-5] + "_test.go"
	goMod := strings.Join(fn.SliceMapping(strings.Split(strings.ReplaceAll(strings.ReplaceAll(c.pkg.Types.Path(), "-", "/"), ".", "/"), "/"), PascalCase), "")
	if cond != "" {
		goMod += cond[1:]
	}
	if strings.ContainsRune(c.pkg.Types.Path(), '.') {
		module = c.pkg.Types.Path()
	} else {
		module = "golang/" + c.pkg.Types.Path()
	}
	//!! gen d.ts
	{
		x.Format("// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection").LF().
			Format("// Code generated by define_gene; DO NOT EDIT.").LF().
			Format("declare module '%s'{", module).LF()
		x.Indent(1)
		imported := fn.NewHashSet[string]()
		for s := range c.imports {
			switch {
			case strings.Contains(s.Name, ".") || (s.Pkg != nil && strings.Contains(s.Pkg.Path(), ".")):
				path := s.Name
				name := s.Name
				if s.Pkg != nil {
					path = s.Pkg.Path()
				}
				if i := strings.LastIndexByte(s.Name, '/'); i >= 0 {
					name = s.Name[i+1:]
				}
				if imported.Add(path) {
					x.IndLF().Format("// @ts-ignore").
						IndLF().Format("import * as %s from '%s'", name, path)
				}
			case s.Name == "go":
				x.IndLF().Format("// @ts-ignore").
					IndLF().Format("import * as go from 'go'")
			default:
				name := s.Name
				path := s.Name
				if s.Pkg != nil {
					path = s.Pkg.Path()
				}
				if i := strings.LastIndexByte(name, '/'); i >= 0 {
					name = name[i+1:]
				}
				if imported.Add(path) {
					x.IndLF().Format("// @ts-ignore").
						IndLF().Format("import * as %s from 'golang/%s'", name, path)
				}
			}
		}
		if c.uses.Len() > 0 {
			if slices.Contains(c.uses.Values(), "Ref") {
				x.IndLF().Format("// @ts-ignore").
					IndLF().Format("import type {%s} from 'go'", strings.Join(c.uses.Values(), ","))
			} else {
				x.IndLF().Format("// @ts-ignore").
					IndLF().Format("import type {%s,Ref} from 'go'", strings.Join(c.uses.Values(), ","))
			}
		} else {
			x.IndLF().Format("// @ts-ignore").
				IndLF().Format("import type {Ref} from 'go'")
		}
		x.Append(c.m)
		//!! generate struct constructor
		for name := range c.Structs {
			if strings.HasSuffix(name, "Err") || strings.HasSuffix(name, "Error") || strings.HasPrefix(name, "Err") || strings.HasPrefix(name, "Error") {
				continue
			}
			if _, ok := c.Functions[fmt.Sprintf("New%s", name)]; !ok {
				x.IndLF().Format("export function empty%s():%s", Safe(name), name)
				x.IndLF().Format("export function emptyRef%s():Ref<%s>", Safe(name), name)
				x.IndLF().Format("export function refOf%s(x:%[2]s,v:Ref<%[2]s>)", Safe(name), name)
				x.IndLF().Format("export function unRef%s(v:Ref<%[2]s>):%[2]s", Safe(name), name)
			}
		}
		//!! generate TypeId
		x.Format("\n}")
		if g.print {
			println(x.Buffer().String())
		} else {
			p := filepath.Join(g.out, filePath, dts)
			if err = write(x.Buffer().Bytes(), p, g.over); err != nil {
				return
			}
		}

	}

	var b []byte

	//!! gen go source
	{
		x.Buffer().Reset()
		if cond != "" {
			x.Format("//go:build %s", cond[1:])
		}
		maps := ""
		if len(c.overrideEntry) > 0 {
			maps = "\"maps\""
		}
		x.Format(`
// Code generated by define_gene; DO NOT EDIT.
package %[1]s
import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	%[2]s
	`, pkg, maps)
		for s := range c.imports {
			name := s.Name
			if s.Pkg != nil {
				name = s.Pkg.Path()
			}
			if strings.ContainsRune(name, '.') {
				x.LF().Format("_ \"%s%s\"", c.root, strings.ReplaceAll(name, ".", "/"))
			} else if name != "go" {
				x.LF().Format("_ \"github.com/ZenLiuCN/engine/modules/golang/%s\"", name)
			}
		}
		if len(c.entry) > 0 || len(c.overrideEntry) > 0 {
			x.Format("\n\"%[1]s\"", pkgPath)
		}
		x.Format(`
)
var (
	//go:embed %[1]s
	%[2]sDefine []byte
	%[2]sDeclared = map[string]any{`, dts, goMod).LF()
		for k, v := range c.entry {
			x.Indent(2).Format("\"%s\" : %s ,", k, v).LF()
		}
		for name := range c.Structs {
			if strings.HasSuffix(name, "Err") || strings.HasSuffix(name, "Error") || strings.HasPrefix(name, "Err") || strings.HasPrefix(name, "Error") {
				continue
			}
			t := fmt.Sprintf("%s.%s", c.pkg.Types.Name(), name)
			if _, ok := c.Functions[fmt.Sprintf("New%s", name)]; !ok {
				x.LF().Format("\"empty%s\":engine.Empty[%[2]s],", Safe(name), t)
				x.LF().Format("\"emptyRef%s\":engine.EmptyRefer[%[2]s],", Safe(name), t)
				x.LF().Format("\"refOf%s\":engine.ReferOf[%[2]s],", Safe(name), t)
				x.LF().Format("\"unRef%s\":engine.UnRefer[%[2]s],", Safe(name), t)
			}
		}
		x.Format("\t}\n)")
		x.Format(`
func init(){
	engine.RegisterModule(%[1]sModule{})
}
type %[1]sModule struct{}

func (S %[1]sModule) Identity() string {
	return "%[2]s"
}
func (S %[1]sModule) TypeDefine() []byte {
	return %[1]sDefine
}
func (S %[1]sModule) Exports() map[string]any {
	return %[1]sDeclared
}
`, goMod, module)
		if len(c.overrideEntry) > 0 {
			x.Format(`
func (e %sModule) ExportsWithEngine(e *Engine) map[string]any{
	m:=maps.Clone([1]sDeclared)

`)
			for k, v := range c.overrideEntry {
				x.Indent(1).Format("m[\"%s\"] = v", k, v).LF()
			}
			x.Format("\treturn m\t}\n")
		}
		if c.wrapper.NotEmpty() {
			x.Append(c.wrapper)
		}
		b, err = format.Source(x.Buffer().Bytes())
		if err != nil {
			log.Printf("generated go source have some error: %w", err)
			b = x.Buffer().Bytes()
		}
		if g.print {
			println(string(b))
		} else {
			p := filepath.Join(g.out, filePath, mod)
			if err = write(b, p, g.over); err != nil {
				return
			}
		}
	}

	//!! gen go test source
	{
		x.Reset()
		x.Format(`
package %[1]s

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestSimple%[2]s(t *testing.T) {
	v := engine.Get()
	defer v.Free()
	fn.Panic1(v.RunTs(
		//language=typescript
		%[3]s))
}

`, pkg, cond, fmt.Sprintf("`"+`
		import * as %s from '%s'
		`+"`", c.pkg.Types.Name(), module))
		b, err = format.Source(x.Buffer().Bytes())
		if err != nil {
			log.Printf("generated go test source have some error: %s", err)
			b = x.Buffer().Bytes()
		}
		if g.print {
			println(string(b))
		} else {
			p := filepath.Join(g.out, filePath, modTest)
			if err = write(b, p, g.overTest); err != nil {
				return
			}
		}
	}
	log.Println("generated", mod)
	return
}
func anyOf(an ...string) (r string) {
	for _, s := range an {
		if s != "" {
			r += ":"
			r += s
		}
	}
	return
}
func write(data []byte, p string, override bool) (err error) {
	if !override && Exists(p) {
		log.Printf("skip exists file '%s'", p)
		return
	}
	err = os.MkdirAll(filepath.Dir(p), os.ModePerm)
	if err != nil {
		return
	}
	return os.WriteFile(p, data, os.ModePerm)
}
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}
	return err == nil
}

func safe(s string, then func(string) string) string {
	switch s {
	case "New", "In", "Default", "Do", "For", "With":
		return s
	case "in", "do", "new":
		return strings.ToUpper(s[0:1]) + s[1:]
	case "class", "Class":
		return "clazz"
	case "NaN":
		return "nan"
	default:
		return then(s)
	}
}
func Safe(s string) string {
	return safe(s, fn.Identity[string])
}
func Camel(s string) string {
	return safe(s, moduleNameCase)
}
func moduleNameCase(s string) string {
	if len(s) <= 1 {
		return s
	}
	for _, i2 := range s[1:] {
		if i2 >= 48 && i2 <= 57 {
			continue
		}
		goto x
	}
	return s
x:
	return CamelCase(s)
}

func assertEmpty(w Writer) {
	if w.NotEmpty() {
		panic("should empty of writer: " + w.String())
	}
}
