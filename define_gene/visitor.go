package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ZenLiuCN/fn"
	. "github.com/ZenLiuCN/go-inspect"
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

var (
	bytesError = []byte(",error]")
)
var ctxVisitor = FnTypeVisitor[*TypeInspector[*context], *context]{
	FnVisitConst: func(I InspectorX, d Dir, name string, e *types.Const, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTConst, I.TypePath, nil)
			c.name = name
		case EXT:
			defer c.Exit(KTConst, I.TypePath, nil)
			if c.generic {
				//TODO handle generic Const
				c.reset()
				return false
			}
			if !c.typeAlias {
				c.mi(1).mf("//%s", e.Val().String()).mn()
				c.Constants[name] = e
			} else {
				c.AliasConst[name] = e
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
			defer c.Enter(KTFunc, I.TypePath, nil)
			c.name = name

		case EXT:
			defer c.Exit(KTFunc, I.TypePath, nil)
			if c.generic {
				//TODO handle generic function
				c.reset()
				return false
			}
			c.mi(1).mf("export function %s", Camel(name))
			c.Functions[name] = e

			if !c.errors {
				b := c.d.Bytes()
				if bytes.HasSuffix(b, bytesError) {
					cnt := bytes.Count(b, []byte{','})
					if cnt == 1 {
						x := bytes.LastIndex(b, []byte{'['})
						b = append(b[0:x], b[x+1:len(b)-len(bytesError)]...)
						c.d.Reset()
						c.d.Buffer().Write(b)
					}

				}

			}
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
		case EXT:
			defer c.Exit(KTType, I.TypePath, nil)
			if c.generic {
				//TODO handle generic
				c.reset()
				return false
			}
			if c.typeAlias {
				c.AliasType[name] = e
				c.mi(1).mf("export type %s=%s", name, c.d.String()).mn()
			} else {
				c.mi(1).mf("export interface %s", name)
				if c.isStruct {
					c.Structs[name] = e
				} else {
					c.Interfaces[name] = e
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
			c.Variables[name] = e
			c.mi(1).mf("export const %s:", name)
			c.declared().mn()
			c.addEntry(name, "%s.%s", c.pkg.Types.Name(), name)
			c.reset()
		}
		return true
	},
	FnVisitTypeTuple: func(I InspectorX, d Dir, o types.Object, x *types.Tuple, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTuple, I.TypePath, seen)
			switch m := I.LNth(2); m {
			case KMParam:
				c.df("(")
				c.pc = c.pc.Push(x.Len())
			case KMResult:
				c.pc = c.pc.Push(x.Len())
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
				if !c.funcAlias && (I.Len() <= 4 || I.LNth(4) == KMMethod) {
					c.df("):")
				} else {
					c.df(")=>")
				}
				c.pc, _ = c.pc.Pop()
			case KMResult:
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
	FnVisitTypeVar: func(I InspectorX, d Dir, o types.Object, x *types.Var, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KVar, I.TypePath, seen)
			switch {
			case c.path.First() == KTType && I.EndWith(KStruct, KMField): //** KMField,KTuple
				c.field = x.Name()
				c.tmp = c.d
				c.d = GetWriter()
			default:
				switch m := I.LNth(3); m {
				case KMParam:
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
				case KMResult:
				case KMField:
					c.field = x.Name()
				}
			}
		case EXT:
			defer c.Exit(KVar, I.TypePath, seen)
			defer c.resetArrayed()
			switch {
			case c.path.First() == KTType && I.EndWith(KStruct, KMField):
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
				switch I.TypePath.LNth(3) {
				case KMParam, KMResult:
					if c.pc.Last() > 1 {
						c.df(",")
					}
					c.param = ""
					c.pcsub()
				case KMField:
					c.field = ""
					c.dn()
				}
			}

		}
		return true
	},
	FnVisitTypeFunc: func(I InspectorX, d Dir, o types.Object, x *types.Func, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KFunc, I.TypePath, seen)
			c.di(2).df("%s", Camel(x.Name()))
			c.method = x.Name()
		case EXT:
			defer c.Exit(KFunc, I.TypePath, seen)
			defer c.resetArrayed()
			c.dn()
			c.method = ""

		}
		return true
	},
	FnVisitTypeBasic: func(I InspectorX, d Dir, o types.Object, x *types.Basic, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KBasic, I.TypePath, seen)
			arr := c.arrayLen.Len()
			//!! is type embedded
			if c.isTypeDefine(KBasic, seen) && !c.IsConstOrVar() {
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
			defer c.Exit(KBasic, I.TypePath, seen)
			if I.LNth(2) == KMMapKey {
				c.df(",")
			}
		}
		return true
	},
	FnVisitTypeMap: func(I InspectorX, d Dir, o types.Object, x *types.Map, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KMap, I.TypePath, seen)
			if c.isTypeDefine(KMap, seen) {
				c.mapAlias = true
				return true
			}
			c.df("Record<")

		case EXT:
			defer c.Exit(KMap, I.TypePath, seen)
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
	FnVisitTypeArray: func(I InspectorX, d Dir, o types.Object, x *types.Array, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KArray, I.TypePath, seen)
			(&c.arrayLen).PushSitu(x.Len())
		case EXT:
			defer c.Exit(KArray, I.TypePath, seen)
			(&c.arrayLen).PopSitu()
			if c.arrayed > 0 && c.arrayed >= c.path.Len() {
				return true
			}
			c.df("[/*%d*/]", x.Len())

		}
		return true
	},
	FnVisitTypeStruct: func(I InspectorX, d Dir, o types.Object, x *types.Struct, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KStruct, I.TypePath, seen)
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
			defer c.Exit(KStruct, I.TypePath, seen)
			switch {
			case c.isTypeDefineExit(KStruct, seen):
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
	FnVisitTypeUnion: func(I InspectorX, d Dir, o types.Object, x *types.Union, seen Types, c *context) bool {

		switch d {
		case ENT:
			defer c.Enter(KUnion, I.TypePath, seen)

		case EXT:
			defer c.Exit(KUnion, I.TypePath, seen)

		}
		return true
	},
	FnVisitTypeSignature: func(I InspectorX, d Dir, o types.Object, x *types.Signature, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KSignature, I.TypePath, seen)
			c.variadic = c.variadic.Push(x.Variadic())
			if c.isTypeDefine(KSignature, seen) {
				c.funcAlias = true
				return true
			}

		case EXT:
			defer c.Exit(KSignature, I.TypePath, seen)
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
	FnVisitTypeParam: func(I InspectorX, d Dir, o types.Object, x *types.TypeParam, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KTypeParam, I.TypePath, seen)
			c.generic = true
		case EXT:
			defer c.Exit(KTypeParam, I.TypePath, seen)

		}
		return true
	},
	FnVisitTypePointer: func(I InspectorX, d Dir, o types.Object, x *types.Pointer, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KPointer, I.TypePath, seen)
			if c.isTypeDefine(KPointer, seen) {
				c.use("Ref")
				c.df("Ref<")
				return true
			}
			c.use("Ref")
			c.df("Ref<")

		case EXT:
			defer c.Exit(KPointer, I.TypePath, seen)
			defer c.resetArrayed()
			c.df(">")
			if c.isTypeDefine(KPointer, seen) && !c.IsConstOrVar() {
				c.extends(c.d.Buffer().String())
				c.d.Reset()
			}

		}
		return true
	},
	FnVisitTypeSlice: func(I InspectorX, d Dir, o types.Object, x *types.Slice, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KSlice, I.TypePath, seen)
			(&c.arrayLen).PushSitu(-1)
		case EXT:
			defer c.Exit(KSlice, I.TypePath, seen)
			(&c.arrayLen).PopSitu()
			if c.arrayed > 0 && (c.arrayed >= c.path.Len() || c.path.Last() == KBasic) {
				return true
			}
			if c.isTypeDefineExit(KSlice, seen) {
				c.extends("Array<%s>", c.d.String())
				c.d.Reset()
			} else {
				c.df("[]")
			}

		}
		return true
	},
	FnVisitTypeInterface: func(I InspectorX, d Dir, o types.Object, x *types.Interface, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KInterface, I.TypePath, seen)
			switch {
			case c.isTypeDefine(KInterface, seen):
				c.isStruct = false
				return true
			case c.isTypeExtend(KInterface, seen):
				c.use("Proto")
				c.extends("Proto<%s>", c.name)
				if x.NumMethods() == 0 {
					c.df("any") //!! any == interface{}
				}
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
	FnVisitTypeChan: func(I InspectorX, d Dir, o types.Object, x *types.Chan, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KChan, I.TypePath, seen)
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
			defer c.Exit(KChan, I.TypePath, seen)
			defer c.resetArrayed()
			c.df(">")

		}
		return true
	},
	FnVisitTypeNamed: func(I InspectorX, d Dir, o types.Object, x *types.Named, seen Types, c *context) bool {
		switch d {
		case ENT:
			name := x.Obj().Name()
			defer c.Enter(KNamed, I.TypePath, seen)
			switch {
			case c.path.Len() == 1 && c.name == name && x.Obj().Pkg() == c.pkg.Types: //!! typeDefine
				return true
			case c.path.Len() == 1: //!! first alias
				c.typeAlias = true
				c.decodeNamed(I, x, c.df)
				return false
			case c.IsType() && I.LNth(2) == KMEmbedded: //!! embedded for interface
				c.decodeNamed(I, x, c.extends)
				return false
			default:
				c.decodeNamed(I, x, c.df)
				return false
			}
		case EXT:
			defer c.Exit(KNamed, I.TypePath, seen)
			defer c.resetArrayed()
			if I.LNth(2) == KMMapKey {
				c.df(",")
			}

		}
		return true
	},
}

type Bool = Stack[bool]
type Int = Stack[int]
type Int64 = Stack[int64]
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
	arrayed   int
	arrayLen  Int64
	variadic  Bool
	pc        Int
	tmp       Writer
}

func (s *state) reset() {
	s.path.CleanSitu()
	s.name = ""
	s.method = ""
	s.field = ""
	s.param = ""

	s.typeAlias = false
	s.funcAlias = false
	s.mapAlias = false
	s.generic = false
	s.isStruct = false

	s.arrayed = 0
	s.arrayLen = s.arrayLen.Clean()
	s.variadic = s.variadic.Clean()
	s.pc = s.pc.Clean()

	if s.tmp != nil {
		s.tmp.Free()
		s.tmp = nil
	}

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
func (s *state) isTypeDefine(k TypeKind, seen Types) bool {
	return s.IsType() &&
		((k == KNamed && s.path.Len() == 2) ||
			(k == KStruct && s.path.Len() == 2) ||
			(k == KInterface && s.path.Len() == 2) ||
			(s.path.Len() == 2 && s.path.Last() == KNamed) ||
			(k == KBasic && s.path.Len() > 2 && !slices.ContainsFunc(s.path[2:], func(kind TypeKind) bool {
				return kind != KSlice && kind != KArray
			})))
}
func (s *state) isTypeExtend(k TypeKind, seen Types) bool {
	return s.IsType() &&
		((k == KNamed && s.path.Len() == 3) ||
			(k == KStruct && s.path.Len() == 3) ||
			(k == KInterface && s.path.Len() == 3))
}
func (s *state) isTypeDefineExit(k TypeKind, seen Types) bool {
	return s.IsType() &&
		((k == KNamed && s.path.Len() == 2) ||
			(k == KSlice && s.path.Len() == 3) ||
			(k == KStruct && s.path.Len() == 3))

}
func (s *state) Enter(k TypeKind, path TypePath, seen Types) {
	if debug {
		log.Printf("ENTER[%s]\t%s:%s[%d]\tΔ%v\tΔ%v", k, s.name, anyOf(s.field, s.method, s.param), s.pc.Last(), path, s.path)
	}

	(&s.path).PushSitu(k)

}
func (s *state) Exit(k TypeKind, path TypePath, seen Types) {
	if debug {
		log.Printf("EXIT[%s]\t%s:%s[%d]\tΔ%v\tΔ%v", k, s.name, anyOf(s.field, s.method, s.param), s.pc.Last(), path, s.path)
	}
	//!! remove children
	kx := s.path.LastIndex(k)
	if kx == -1 {
		return
	}
	s.path = s.path[:kx]

}

type context struct {
	Functions     map[string]*types.Func
	Structs       map[string]*types.TypeName
	Interfaces    map[string]*types.TypeName
	Constants     map[string]*types.Const
	AliasConst    map[string]*types.Const
	AliasType     map[string]*types.TypeName
	Variables     map[string]*types.Var
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
	errors bool //reduce function end with error
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
	c.Functions = map[string]*types.Func{}
	c.Structs = map[string]*types.TypeName{}
	c.Interfaces = map[string]*types.TypeName{}
	c.Constants = map[string]*types.Const{}
	c.AliasConst = map[string]*types.Const{}
	c.AliasType = map[string]*types.TypeName{}
	c.Variables = map[string]*types.Var{}

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
func (c *context) flush(g *Generator) (err error) {
	x := GetWriter()
	defer x.Free()
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
	goMod := strings.Join(fn.SliceMapping(strings.Split(strings.ReplaceAll(c.pkg.Types.Path(), ".", "/"), "/"), PascalCase), "")
	if cond != "" {
		goMod += cond[1:]
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
		for name := range c.Structs {
			if strings.HasSuffix(name, "Err") || strings.HasSuffix(name, "Error") || strings.HasPrefix(name, "Err") || strings.HasPrefix(name, "Error") {
				continue
			}
			x.LF().Format("export function empty%s():%s", Safe(name), name)
			x.LF().Format("export function ref%s():Ref<%s>", Safe(name), name)
			x.LF().Format("export function refOf%s(x:%[2]s):Ref<%[2]s>", Safe(name), name)
		}
		//!! generate TypeId
		x.Format("}").LF()
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
	%[2]sDeclared = map[string]any{`, dts, goMod).LF()
		for k, v := range c.entry {
			x.Indent(2).Format("\"%s\" : %s ,", k, v).LF()
		}
		for name := range c.Structs {
			if strings.HasSuffix(name, "Err") || strings.HasSuffix(name, "Error") || strings.HasPrefix(name, "Err") || strings.HasPrefix(name, "Error") {
				continue
			}
			t := fmt.Sprintf("%s.%s", c.pkg.Types.Name(), name)
			x.LF().Format("\"empty%s\":func() (v %[2]s){\n\treturn v\n},", Safe(name), t)
			x.LF().Format("\"ref%s\":func() *%[2]s{ \n\t var x %[2]s\n\treturn &x\n},", Safe(name), t)
			x.LF().Format("\"refOf%s\":func(x %[2]s) *%[2]s{ \n\treturn &x\n},", Safe(name), t)
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
			log.Printf("generated go source have some error: %s", err)
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

func safe(s string, then func(string) string) string {
	switch s {
	case "New", "In", "Default", "Do", "For", "With":
		return s
	case "in", "do", "new":
		return strings.ToUpper(s[0:1]) + s[1:]
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
