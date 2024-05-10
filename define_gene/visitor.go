package main

import (
	"bytes"
	"errors"
	"fmt"
	. "github.com/ZenLiuCN/engine/define_gene/internal"
	"github.com/ZenLiuCN/fn"
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

func safe(s string, then func(string) string) string {
	switch s {
	case "New", "In", "Default", "Do", "For":
		return s
	case "new":
		return "New"
	case "class", "Class":
		return "clazz"
	default:
		return then(s)
	}
}
func Safe(s string) string {
	return safe(s, fn.Identity[string])
}
func Camel(s string) string {
	return safe(s, CamelCase)
}

var ctxVisitor = FnTypeVisitor[*TypeInspector[*context], *context]{
	FnVisitConst: func(I InspectorX, d Dir, name string, e *types.Const, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTConst, nil, nil)
			c.name = name
		case EXT:
			defer c.Exit(KTConst, nil, nil)
			if c.generic {
				//TODO handle generic Const
				c.reset()
				return false
			}
			if !c.typeAlias {
				c.mi(1).mf("//%s", e.Val().String()).mn()
				c.Constants[name] = e.String()
			} else {
				c.Alias[name] = c.d.String()
			}
			c.mi(1).mf("export const %s:", name)
			c.declared()
			c.addEntry(name, "%s.%s", c.pkg.Types.Name(), name)
			c.mn().reset()
		}
		return true
	},
	FnVisitFunc: func(I InspectorX, d Dir, name string, e *types.Func, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTFunc, nil, nil)
			c.name = name

		case EXT:
			defer c.Exit(KTFunc, nil, nil)
			if c.generic {
				//TODO handle generic function
				c.reset()
				return false
			}
			c.mi(1).mf("export function %s", Camel(name))
			c.Functions[name] = e.String()
			c.declared().mn()
			c.addEntry(Camel(name), "%s.%s", c.pkg.Types.Name(), name)
			c.reset()
		}
		return true
	},
	FnVisitTypeName: func(I InspectorX, d Dir, name string, e *types.TypeName, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTType, nil, nil)
			c.name = name
		case EXT:
			defer c.Exit(KTType, nil, nil)
			if c.generic {
				//TODO handle generic
				c.reset()
				return false
			}
			if c.typeAlias {
				c.Alias[name] = c.d.String()
				c.mi(1).mf("export type %s=%s", name, c.d.String()).mn()
			} else {
				c.mi(1).mf("export interface %s", name)
				if c.isStruct {
					c.Structs[name] = ""
				} else {
					c.Interfaces[name] = ""
				}
				c.faceIntercept()
				if len(c.extended) > 0 {
					c.mf(" extends %s", strings.Join(c.extended.Values(), ","))
				}
				c.mf("{").mn()
				c.declared().
					mi(1).mf("}").mn()
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

		case EXT:
			defer c.Exit(KTVar, nil, nil)
			if c.generic {
				//TODO handle generic Const
				c.reset()
				return false
			}
			c.Variables[name] = c.d.String()
			c.mi(1).mf("export const %s:", name)
			c.declared().mn()
			c.addEntry(name, "%s.%s", c.pkg.Types.Name(), name)
			c.reset()
		}
		return true
	},
	FnVisitTypeTuple: func(I InspectorX, d Dir, o types.Object, x *types.Tuple, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTuple, mods, seen)
			switch m := mods.LNth(1); true {
			case m.IsParam():
				c.df("(")
				c.pc = c.pc.Push(x.Len())
			case m.IsResult():
				c.pc = c.pc.Push(x.Len())
				switch x.Len() {
				case 0:
				case 1:
				default:
					c.df("[")
				}
			}
		case EXT:
			defer c.Exit(KTuple, mods, seen)
			switch m := mods.LNth(1); true {
			case m.IsParam():
				if !c.funcAlias && (mods.Len() < 3 || mods.LNth(2).IsMethod()) {
					c.df("):")
				} else {
					c.df(")=>")
				}
				c.pc, _ = c.pc.Pop()
			case m.IsResult():
				switch x.Len() {
				case 0:
					c.df("void")
				case 1:
				default:
					c.df("]")
				}
				c.pc, _ = c.pc.Pop()
			}
		}
		return true
	},
	FnVisitTypeVar: func(I InspectorX, d Dir, o types.Object, x *types.Var, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KVar, mods, seen)
			switch {
			case c.path.First() == KTType && mods.LNth(1).IsField():
				c.field = x.Name()
				c.tmp = c.d
				c.d = GetWriter()
			default:
				switch m := mods.LNth(2); true {
				case m.IsParam() || m.IsResult():
					if m.IsParam() {
						c.param = x.Name()
						if c.pc.Last() == 1 && c.variadic.Last() {
							if c.param == "" {
								c.df("...v%d:", c.pc.Last())
							} else {
								c.df("...%s:", Safe(x.Name()))
							}
						} else {
							if c.param == "" {
								c.df("v%d:", c.pc.Last())
							} else {
								c.df("%s:", Safe(x.Name()))
							}
						}
					}
				case m.IsField():
					c.field = x.Name()
				}
			}
		case EXT:
			defer c.Exit(KVar, mods, seen)
			defer c.resetArrayed()
			switch {
			case c.path.First() == KTType && mods.LNth(1).IsField():
				if c.d.Buffer().Len() > 0 { //!! only expose field type should write
					t := c.d
					c.d = c.tmp
					c.tmp = nil
					c.di(2).df("%s:%s", Camel(c.field), t.String())
					t.Free()
					c.dn()
				} else {
					t := c.d
					c.d = c.tmp
					c.tmp = nil
					t.Free()
				}
				c.field = ""
			default:
				switch m := mods.LNth(2); true {
				case m.IsParam() || m.IsResult():
					if c.pc.Last() > 1 {
						c.df(",")
					}
					c.param = ""
					c.pcsub()
					/*			if c.pc.Last() == 0 && m.IsParam() && c.variadic.Last() && c.arrayed == 0 {
								c.df("[]")
							}*/
				case m.IsField():
					c.field = ""
					c.dn()
				}
			}

		}
		return true
	},
	FnVisitTypeFunc: func(I InspectorX, d Dir, o types.Object, x *types.Func, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KFunc, mods, seen)
			c.di(2).df("%s", Camel(x.Name()))
			c.method = x.Name()
		case EXT:
			defer c.Exit(KFunc, mods, seen)
			c.dn()
			c.method = ""
			if c.arrayed >= c.path.Len() {
				c.arrayed = 0
			}
		}
		return true
	},
	FnVisitTypeBasic: func(I InspectorX, d Dir, o types.Object, x *types.Basic, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KBasic, mods, seen)
			arr := c.arrayLen.Len()
			//!! is type embedded
			if c.isTypeDefine(KBasic, mods, seen) && !c.IsConstOrVar() {
				if arr == 1 || arr == 0 {
					if c.arrayLen.Last() == -1 || arr == 0 {
						c.extends("%s", c.identType(x.Name(), arr > 0, true))
					} else {
						c.extends("%s/*%d*/", c.identType(x.Name(), arr > 0, true), c.arrayLen.Last())
					}
				} else {

					b := GetWriter()
					for _, i := range c.arrayLen[:arr-1] {
						if i == -1 {
							b.Format("Array<")
						} else {
							b.Format("Array</*%d*/", i)
						}
					}
					lst := c.arrayLen.Last()
					s := ""
					if lst > 0 {
						s = fmt.Sprintf("/*%d*/", lst)
					}
					arr--
					c.extends("%s%s%s%s", b.String(), c.identType(x.Name(), true, true), s, strings.Repeat(">", arr))
					b.Free()

				}
				c.arrayed = c.path.Len() + 1 //!! KBasic push after process
				return true
			}
			if arr == 1 || arr == 0 {
				if c.arrayLen.Last() == -1 || arr == 0 {
					c.df("%s", c.identType(x.Name(), arr > 0, false))
				} else {
					c.df("%s/*%d*/", c.identType(x.Name(), arr > 0, false), c.arrayLen.Last())
				}
			} else {

				b := GetWriter()
				for _, i := range c.arrayLen[:arr-1] {
					if i == -1 {
						b.Format("Array<")
					} else {
						b.Format("Array</*%d*/", i)
					}
				}
				arr--
				c.df("%s%s%s", b.String(), c.identType(x.Name(), true, false), strings.Repeat(">", arr))
				b.Free()
			}
			c.arrayed = c.path.Len() + 1 //!! KBasic push after process
		case EXT:
			defer c.Exit(KBasic, mods, seen)
			if mods.LNth(1).IsMapKey() {
				c.df(",")
			}
		}
		return true
	},
	FnVisitTypeMap: func(I InspectorX, d Dir, o types.Object, x *types.Map, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KMap, mods, seen)
			if c.isTypeDefine(KMap, mods, seen) {
				c.mapAlias = true
				return true
			}
			c.df("Record<")

		case EXT:
			defer c.Exit(KMap, mods, seen)
			defer c.resetArrayed()
			if c.mapAlias {
				c.mapAlias = false
				c.extends("Record<%s>", c.d.Buffer().String())
				c.d.Buffer().Reset()
				return true
			}
			c.df(">")

		}
		return true
	},
	FnVisitTypeArray: func(I InspectorX, d Dir, o types.Object, x *types.Array, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KArray, mods, seen)
			(&c.arrayLen).PushSitu(x.Len())
		case EXT:
			defer c.Exit(KArray, mods, seen)
			(&c.arrayLen).PopSitu()
			if c.arrayed > 0 && c.arrayed >= c.path.Len() {
				return true
			}
			c.df("[/*%d*/]", x.Len())

		}
		return true
	},
	FnVisitTypeStruct: func(I InspectorX, d Dir, o types.Object, x *types.Struct, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KStruct, mods, seen)
			switch {
			case x.NumFields() == 0:

				c.use("Nothing")
				if c.path.Len() <= 2 {
					c.use("Alias")
					c.extends("Alias<Nothing>")
				} else {
					c.df("Nothing")
				}

			}
		case EXT:
			defer c.Exit(KStruct, mods, seen)
			switch {
			case c.isTypeDefineExit(KStruct, mods, seen):
				switch {
				case strings.Contains(c.name, "Err"):
					c.use("Struct")
					c.extends("Struct<%s>", c.name)
					c.extends("Error")
					c.isStruct = true
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
							if !c.extended.Exists(name) {
								c.imported(face.Pkg.Path())
								c.extends(name)
							}
						}
					}
				}
			}
		}
		return true
	},
	FnVisitTypeUnion: func(I InspectorX, d Dir, o types.Object, x *types.Union, mods Mods, seen Types, c *context) bool {

		switch d {
		case ENT:
			defer c.Enter(KUnion, mods, seen)

		case EXT:
			defer c.Exit(KUnion, mods, seen)

		}
		return true
	},
	FnVisitTypeSignature: func(I InspectorX, d Dir, o types.Object, x *types.Signature, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KSign, mods, seen)
			c.variadic = c.variadic.Push(x.Variadic())
			if c.isTypeDefine(KSign, mods, seen) {
				c.funcAlias = true
				return true
			}

		case EXT:
			defer c.Exit(KSign, mods, seen)
			c.variadic, _ = c.variadic.Pop()
			if c.funcAlias {
				c.funcAlias = false
				c.use("Alias")
				c.extends("Alias<%s>", c.d.String())
				c.d.Reset()
				return true
			}

		}
		return true
	},
	FnVisitTypeParam: func(I InspectorX, d Dir, o types.Object, x *types.TypeParam, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTypeParam, mods, seen)
			c.generic = true
		case EXT:
			defer c.Exit(KTypeParam, mods, seen)

		}
		return true
	},
	FnVisitTypePointer: func(I InspectorX, d Dir, o types.Object, x *types.Pointer, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KPointer, mods, seen)
			if c.isTypeDefine(KPointer, mods, seen) {
				c.use("Ref")
				c.df("Ref<")
				return true
			}
			c.use("Ref")
			c.df("Ref<")

		case EXT:
			defer c.Exit(KPointer, mods, seen)
			defer c.resetArrayed()
			c.df(">")
			if c.isTypeDefine(KPointer, mods, seen) && !c.IsConstOrVar() {
				c.extends(c.d.Buffer().String())
				c.d.Reset()
			}

		}
		return true
	},
	FnVisitTypeSlice: func(I InspectorX, d Dir, o types.Object, x *types.Slice, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KSlice, mods, seen)
			(&c.arrayLen).PushSitu(-1)
		case EXT:
			defer c.Exit(KSlice, mods, seen)
			(&c.arrayLen).PopSitu()
			if c.arrayed > 0 && (c.arrayed >= c.path.Len() || c.path.Last() == KBasic) {
				return true
			}
			if c.isTypeDefineExit(KSlice, mods, seen) {
				c.extends("Array<%s>", c.d.String())
				c.d.Reset()
			} else {
				c.df("[]")
			}

		}
		return true
	},
	FnVisitTypeInterface: func(I InspectorX, d Dir, o types.Object, x *types.Interface, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KInterface, mods, seen)
			switch {
			case c.isTypeDefine(KInterface, mods, seen):
				return true
			case c.isTypeExtend(KInterface, mods, seen):
				c.use("Proto")
				c.extends("Proto<%s>", c.name)
				return false
			default:
				if x.NumMethods() == 0 {
					c.df("any") //!! any == interface{}
				}

			}
		case EXT:
			defer c.Exit(KInterface, mods, seen)
		}
		return true
	},
	FnVisitTypeChan: func(I InspectorX, d Dir, o types.Object, x *types.Chan, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KChan, mods, seen)
			c.imported("go")
			switch x.Dir() {
			case types.SendOnly:
				c.df("go.ChanSend<")
			case types.RecvOnly:
				c.df("go.ChanRecv<")
			default:
				c.df("go.Chan<")
			}

		case EXT:
			defer c.Exit(KChan, mods, seen)
			defer c.resetArrayed()
			c.df(">")

		}
		return true
	},
	FnVisitTypeNamed: func(I InspectorX, d Dir, o types.Object, x *types.Named, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			name := x.Obj().Name()
			defer c.Enter(KNamed, mods, seen)
			switch {
			case c.path.Len() == 1 && c.name == name && x.Obj().Pkg() == c.pkg.Types: //!! typeDefine
				return true
			case c.path.Len() == 1: //!! first alias
				c.typeAlias = true
				c.decodeNamed(I, x, c.df)
				return false
			case c.IsType() && mods.LNth(1).IsEmbedded(): //!! embedded for interface
				c.decodeNamed(I, x, c.extends)
				return false
			default:
				c.decodeNamed(I, x, c.df)
				return false
			}
		case EXT:
			defer c.Exit(KNamed, mods, seen)
			defer c.resetArrayed()
			if mods.LNth(1).IsMapKey() {
				c.df(",")
			}

		}
		return true
	},
}

type Path = Stack[Kind]

//go:generate stringer -type=Kind
type Kind int

const (
	KNone Kind = iota
	KTFunc
	KTType
	KTConst
	KTVar
	KVar
	KFunc

	KNamed
	KStruct
	KMap
	KPointer
	KArray
	KSlice
	KInterface
	KChan
	KBasic
	KTypeParam
	KSign
	KUnion
	KTuple
)

type Bool = Stack[bool]
type Int = Stack[int]
type Int64 = Stack[int64]
type state struct {
	name   string
	method string
	field  string
	param  string

	typeAlias bool
	funcAlias bool
	mapAlias  bool
	generic   bool
	isStruct  bool
	arrayed   int
	arrayLen  Int64
	path      Path
	variadic  Bool
	pc        Int
	tmp       Writer
}

func (s *state) reset() {
	s.name = ""
	s.method = ""
	s.field = ""
	s.param = ""

	s.typeAlias = false
	s.funcAlias = false
	s.mapAlias = false
	s.generic = false
	s.arrayed = 0
	s.arrayLen = s.arrayLen.Clean()

	s.pc = s.pc.Clean()
	s.path = s.path.Clean()
	s.variadic = s.variadic.Clean()

}
func (s *state) resetArrayed() {
	if s.arrayed >= s.path.Len() {
		s.arrayed = 0
	}
}
func (s *state) pcsub() {
	if s.pc.Empty() {
		return
	}
	s.pc[s.pc.MaxIdx()]--
}
func (s *state) IsConstOrVar() bool {
	f := s.path.First()
	return f == KTVar || f == KTConst
}
func (s *state) IsConst() bool {
	return s.path.First() == KTConst
}
func (s *state) IsFunc() bool {
	return s.path.First() == KTFunc
}
func (s *state) IsVar() bool {
	return s.path.First() == KTVar
}
func (s *state) IsType() bool {
	return s.path.First() == KTType
}
func (s *state) isTypeDefine(k Kind, mods Mods, seen Types) bool {
	return s.IsType() &&
		((k == KNamed && s.path.Len() == 2) ||
			(k == KStruct && s.path.Len() == 2) ||
			(k == KInterface && s.path.Len() == 2) ||
			(s.path.Len() == 2 && s.path.Last() == KNamed) ||
			(k == KBasic && s.path.Len() > 2 && !slices.ContainsFunc(s.path[2:], func(kind Kind) bool {
				return kind != KSlice && kind != KArray
			})))
}
func (s *state) isTypeExtend(k Kind, mods Mods, seen Types) bool {
	return s.IsType() &&
		((k == KNamed && s.path.Len() == 3) ||
			(k == KStruct && s.path.Len() == 3) ||
			(k == KInterface && s.path.Len() == 3))
}
func (s *state) isTypeDefineExit(k Kind, mods Mods, seen Types) bool {
	return s.IsType() &&
		((k == KNamed && s.path.Len() == 2) ||
			(k == KSlice && s.path.Len() == 3) ||
			(k == KStruct && s.path.Len() == 3))

}
func (s *state) Enter(k Kind, mods Mods, seen Types) {
	if debug {
		log.Printf("ENTER[%s]\t%s:%s[%d]\tΔ%s\tΔ%s", k, s.name, anyOf(s.field, s.method, s.param), s.pc.Last(), mods, s.path)
	}

	(&s.path).PushSitu(k)

}
func (s *state) Exit(k Kind, mods Mods, seen Types) {
	if debug {
		log.Printf("EXIT[%s]\t%s:%s[%d]\tΔ%s\tΔ%s", k, s.name, anyOf(s.field, s.method, s.param), s.pc.Last(), mods, s.path)
	}
	//!! remove children
	kx := s.path.LastIndex(k)
	if kx == -1 {
		return
	}
	s.path = s.path[:kx]

}

type context struct {
	Functions     map[string]string
	Structs       map[string]string
	Interfaces    map[string]string
	Constants     map[string]string
	Alias         map[string]string
	Variables     map[string]string
	pkg           *packages.Package
	uses          fn.HashSet[string]
	imports       fn.HashSet[string]
	entry         map[string]string
	overrideEntry map[string]string
	wrapper       Writer
	m             Writer
	extended      fn.HashSet[string]
	d             Writer
	FnTypeVisitor[*TypeInspector[*context], *context]

	state
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
			c.imported(x.Obj().Pkg().Path())
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
		}
	}
	return false
}

func (c *context) init() {
	c.Alias = map[string]string{}
	c.Functions = map[string]string{}
	c.Structs = map[string]string{}
	c.Interfaces = map[string]string{}
	c.Constants = map[string]string{}
	c.Variables = map[string]string{}
	c.uses = fn.HashSet[string]{}
	c.imports = fn.HashSet[string]{}
	c.extended = fn.HashSet[string]{}
	c.entry = map[string]string{}
	c.overrideEntry = map[string]string{}
	c.wrapper = GetWriter()
	c.m = GetWriter()
	c.d = GetWriter()

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

func (c *context) imported(s string) {
	c.imports.Add(s)
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
func (c *context) di(n int) *context {
	c.d.Indent(n)
	return c
}
func (c *context) df(format string, args ...any) *context {
	c.d.Format(format, args...)
	return c
}
func (c *context) mi(n int) *context {
	c.m.Indent(n)
	return c
}
func (c *context) mf(s string, args ...any) *context {
	c.m.Format(s, args...)
	return c
}

func (c *context) dn() *context {
	c.d.LF()

	return c
}
func (c *context) mn() *context {
	c.m.LF()
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
			c.imported("io")
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

func anyOf(an ...string) (r string) {
	for _, s := range an {
		if s != "" {
			r += ":"
			r += s
		}
	}
	return
}
func firstOf(an ...string) (r string) {
	for _, s := range an {
		if s != "" {
			return s
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
func (c *context) flush(g *Generator) (err error) {
	x := GetWriter()
	defer x.Free()
	module := ""
	cond := ""
	if g.suffix != "" {
		cond = "_" + g.suffix
	}
	pkgPath := c.pkg.Types.Path()
	pkg := c.pkg.Types.Name()
	dts := strings.Join(fn.SliceMapping(strings.Split(strings.ReplaceAll(c.pkg.Types.Path(), ".", "/"), "/"), strings.ToLower), "_") + cond + ".d.ts"
	mod := dts[:len(dts)-5] + ".go"
	modTest := dts[:len(dts)-5] + "_test.go"
	name := strings.Join(fn.SliceMapping(strings.Split(c.pkg.Types.Path(), "/"), PascalCase), "")
	if cond != "" {
		name += cond[1:]
	}
	if strings.ContainsRune(c.pkg.Types.Path(), '.') {
		module = strings.ReplaceAll(c.pkg.Types.Path(), ".", "/")
	} else {
		module = "golang/" + c.pkg.Types.Path()
	}
	//!! gen d.ts
	{
		x.Format("// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection").LF().
			Format("// Code generated by define_gene; DO NOT EDIT.").LF().
			Format("declare module '%s'{", module).LF()
		for s := range c.imports {
			switch {
			case strings.Contains(s, "."):
				name := s
				if i := strings.LastIndexByte(s, '/'); i >= 0 {
					name = s[i+1:]
				}
				x.Indent(1).Format("// @ts-ignore").LF().
					Indent(1).Format("import * as %s from '%s'", name, s).LF()
			case s == "go":
				x.Indent(1).Format("// @ts-ignore").LF().
					Indent(1).Format("import * as go from 'go'").LF()
			default:
				name := s
				if i := strings.LastIndexByte(s, '/'); i >= 0 {
					name = s[i+1:]
				}
				x.Indent(1).Format("// @ts-ignore").LF().
					Indent(1).Format("import * as %s from 'golang/%s'", name, s).LF()
			}
		}
		if c.uses.Len() > 0 {
			x.Indent(1).Format("// @ts-ignore").LF().
				Indent(1).Format("import type {%s} from 'go'", strings.Join(c.uses.Values(), ",")).LF()
		}
		x.Append(c.m)
		//!! generate struct constructor
		/*		for name := range c.Structs {
				x.LF().Format("export function %s():%s", Safe(name), name)
			}*/
		//!! generate TypeId
		x.Format("}").LF()
		if g.print {
			println(x.Buffer().String())
		} else {
			p := filepath.Join(g.out, pkgPath, dts)
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
	%[3]s
	"%[2]s"`, pkg, pkgPath, maps)
		for s := range c.imports {
			if strings.ContainsRune(s, '.') {
				x.LF().Format("_ \"github.com/ZenLiuCN/engine/%s\"", strings.ReplaceAll(s, ".", "/"))
			} else if s != "go" {
				x.LF().Format("_ \"github.com/ZenLiuCN/engine/golang/%s\"", s)
			}
		}
		x.Format(`
)
var (
	//go:embed %[1]s
	%[2]sDefine []byte
	%[2]sDeclared = map[string]any{`, dts, name).LF()
		for k, v := range c.entry {
			x.Indent(2).Format("\"%s\" : %s ,", k, v).LF()
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
`, name, module)
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
			log.Printf("generated go source have some error: %s", err)
			b = x.Buffer().Bytes()
		}
		if g.print {
			println(string(b))
		} else {
			p := filepath.Join(g.out, pkgPath, mod)
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
		import * as time from '%s'
		`+"`", module))
		b, err = format.Source(x.Buffer().Bytes())
		if err != nil {
			log.Printf("generated go test source have some error: %s", err)
			b = x.Buffer().Bytes()
		}
		if g.print {
			println(string(b))
		} else {
			p := filepath.Join(g.out, pkgPath, modTest)
			if err = write(b, p, g.overTest); err != nil {
				return
			}
		}
	}
	log.Println("generated", mod)
	return
}
