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
	"strings"
	"unicode"
)

func safe(s string, then func(string) string) string {
	switch s {
	case "New":
		return s
	case "new":
		return "New"
	case "class":
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
				c.reset()
				return false
			}
			if !c.typeAlias {
				c.mi(1).mf("//%s", e.Val().String()).mn()
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
			c.mi(1).mf("export function %s", Camel(name))
		case EXT:
			defer c.Exit(KTFunc, nil, nil)
			if c.generic {
				c.reset()
				return false
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
			defer c.Enter(KTType, nil, nil)
			c.name = name
		case EXT:
			defer c.Exit(KTType, nil, nil)
			if c.generic {
				c.reset()
				return false
			}
			if c.typeAlias {
				c.mi(1).mf("export type %s=%s", name, c.d.String()).mn()
			} else {
				c.mi(1).mf("export interface %s", name)
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
			if c.generic {
				c.reset()
				return false
			}
			defer c.Exit(KTVar, nil, nil)
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
			if c.path.First() == KTType && mods.LNth(1).IsField() {
				c.di(2).df("%s:", Camel(x.Name()))
				c.field = x.Name()
			} else {
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
			defer func() { (&c.path).PopSitu() }() //!! remove TVar itself
			if c.path.First() == KTType && mods.LNth(1).IsField() {
				c.dn()
				c.field = ""
			} else {
				switch m := mods.LNth(2); true {
				case m.IsParam() || m.IsResult():
					if c.pc.Last() > 1 {
						c.df(",")
					}
					c.param = ""
					c.pcsub()
					if c.pc.Last() == 0 && m.IsParam() && c.variadic.Last() {
						c.df("[]")
					}
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

		}
		return true
	},
	FnVisitTypeBasic: func(I InspectorX, d Dir, o types.Object, x *types.Basic, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KBasic, mods, seen)
			arr := 0
			for {
				n := c.path.LNth(arr + 1)
				if n == KArray || n == KSlice {
					arr++
				} else {
					break
				}
			}
			if c.isTypeDefine(KBasic, mods, seen) && !c.IsConstOrVar() {
				if arr == 1 || arr == 0 {
					c.extends("%s", c.identType(x.Name(), arr > 0))
				} else {
					arr--
					c.extends("%s%s%s", strings.Repeat("Array<", arr), c.identType(x.Name(), true), strings.Repeat(">", arr))
				}
				return true
			}
			if arr == 1 || arr == 0 {
				c.df("%s", c.identType(x.Name(), arr > 0))
			} else {
				arr--
				c.df("%s%s%s", strings.Repeat("Array<", arr), c.identType(x.Name(), true), strings.Repeat(">", arr))
			}
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
			if c.isTypeDefine(KArray, mods, seen) {

			} else {

			}
		case EXT:
			defer c.Exit(KArray, mods, seen)
			if c.path.Last() == KBasic {
				return true
			}
			if c.isTypeDefine(KArray, mods, seen) && c.path.Len() == 3 {
				c.extends("Array<%s>", c.d.String())
				c.d.Reset()
			} else {
				c.df("[]")
			}

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
			case c.isTypeDefine(KStruct, mods, seen):
				switch {
				case strings.Contains(c.name, "Err"):
					c.use("Struct")
					c.extends("Struct<%s>", c.name)
					c.extends("Error")
				default:
					c.use("Struct")
					c.extends("Struct<%s>", c.name)
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

		case EXT:
			defer c.Exit(KSlice, mods, seen)
			if c.path.Last() == KBasic {
				return true
			}
			if c.isTypeDefine(KSlice, mods, seen) {
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
				c.use("Proto")
				c.extends("Proto<%s>", c.name)
				return false
			default:
				c.df("any") //!! any == interface{}
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
			c.df(">")

		}
		return true
	},
	FnVisitTypeNamed: func(I InspectorX, d Dir, o types.Object, x *types.Named, mods Mods, seen Types, c *context) bool {
		switch d {
		case ENT:
			defer c.Enter(KNamed, mods, seen)
			switch {
			case c.path.Len() == 1 && c.name == x.Obj().Name() && x.Obj().Pkg() == c.pkg.Types:
				return true
			case c.path.Len() == 1: //!! first alias
				c.typeAlias = true
				c.decodeNamed(x, c.df)
				return false
			case c.IsType() && mods.LNth(1).IsEmbedded(): //!! embedded for interface
				c.decodeNamed(x, c.extends)
				return false
			default:
				c.decodeNamed(x, c.df)
				return false
			}
		case EXT:
			defer c.Exit(KNamed, mods, seen)
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
type state struct {
	name   string
	method string
	field  string
	param  string

	typeAlias bool
	funcAlias bool
	mapAlias  bool
	generic   bool

	path     Path
	variadic Bool
	pc       Int
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

	s.pc = s.pc.Clean()
	s.path = s.path.Clean()
	s.variadic = s.variadic.Clean()

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
		((k == KNamed && s.path.Last() == k && mods.Empty()) ||
			(mods.Len() == 1 && len(seen) == 1 && mods.Last().IseNamedElt()))
}

type context struct {
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

func (c *context) decodeNamed(x *types.Named, f func(string, ...any) *context) {
	if x.Obj().Name() == "error" {
		c.use("error")
		f("error")
	} else if x.Obj().Pkg().Path() == c.pkg.Types.Path() {
		f("%s", x.Obj().Name())
	} else {
		c.imported(x.Obj().Pkg().Path())
		f("%s.%s", x.Obj().Pkg().Name(), x.Obj().Name())
	}
}

func (s *state) Enter(k Kind, mods Mods, seen Types) {
	if debug {
		fmt.Printf("ENTER[%s]\t%s:%s[%d]\tΔ%s\tΔ%s", k, s.name, anyOf(s.field, s.method, s.param), s.pc.Last(), mods, s.path)
	}

	(&s.path).PushSitu(k)

}
func (s *state) Exit(k Kind, mods Mods, seen Types) {
	if debug {
		fmt.Printf("EXIT[%s]\t%s:%s[%d]\tΔ%s\tΔ%s", k, s.name, anyOf(s.field, s.method, s.param), s.pc.Last(), mods, s.path)
	}
	l := s.path.Last()
	if l == 0 {
		return
	}
	if l != k || l <= KTVar {
		(&s.path).PopSitu()
	}

}
func (c *context) init() {
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
func (c *context) identType(s string, array bool) string {
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
			return s + "[]"
		} else {
			c.uses.Add(s)
			return s
		}
	case "string":
		if array {
			return "string[]"
		} else {
			return "string"
		}
	case "any":
		if array {
			return "any[]"
		} else {
			return "any"
		}
	default:
		if array {
			if !unicode.IsUpper(rune(s[0])) {
				c.uses.Add(s)
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
func (c *context) flush(g *Generator) (err error) {
	x := GetWriter()
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
		x.Format(`
// Code generated by define_gene; DO NOT EDIT.
package %[1]s
import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"%[2]s"`, pkg, pkgPath)
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
	m:=make(map[string]any,len([1]sDeclared)+10)
	for k,v:=range [1]sDeclared {
		m[k]=v
	}
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
	return
}
