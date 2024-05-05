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

type Generator struct {
	dir      string
	tags     []string
	files    []string
	out      string
	log      func(format string, a ...any)
	over     bool
	overTest bool
	print    bool
	trace    bool
}

func (g *Generator) generate() (err error) {
	debug = g.log != nil
	if g.trace {
		Debug = true
		Trace = true
	}
	var flags []string
	if len(g.tags) > 0 {
		flags = append(flags, fmt.Sprintf("-tags=%s", strings.Join(g.tags, " ")))
	}
	c := new(context)
	c.init()
	ip := NewTypeInspector(false)
	ip.Inspector.Visitor = c
	err = ip.Inspect(
		func(config *packages.Config) {
			config.BuildFlags = flags
			//config.Mode |= packages.NeedImports
		},
		func(i []*packages.Package) error {
			if len(i) != 1 {
				return fmt.Errorf("only one package each run")
			}
			c.pkg = i[0]
			return nil
		}, g.files...)
	if err != nil {
		return err
	}

	return c.flush(g)
}

func safe(s string, then func(string) string) string {
	switch s {
	case "New":
		return s
	case "class":
		return "clazz"
	default:
		return then(s)
	}
}
func Camel(s string) string {
	return safe(s, CamelCase)
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
	FnTypeVisitor[*TypeInspector]

	name      string
	named     bool
	method    string
	field     string
	param     string
	variadic  []bool
	consts    bool
	types     bool
	funcs     bool
	funcAlias bool
	mapAlias  bool
	pc        int
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

	c.FnTypeVisitor = FnTypeVisitor[*TypeInspector]{
		FnVisitConst: func(I *TypeInspector, d Dir, name string, e *types.Const) bool {
			switch d {
			case ENT:
				c.name = name
				c.consts = true
			case EXT:
				if !c.named {
					c.mi(1).mf("//%s", e.Val().String()).mn()
				}
				c.mi(1).mf("export const %s:", name)
				c.declared()
				c.addEntry(name, "%s.%s", c.pkg.Types.Name(), name)
				c.mn().reset()
			}
			return true
		},
		FnVisitFunc: func(I *TypeInspector, d Dir, name string, e *types.Func) bool {
			switch d {
			case ENT:
				c.name = name
				c.funcs = true
				c.mi(1).mf("export function %s", Camel(name))
			case EXT:
				c.declared().mn()
				c.addEntry(Camel(name), "%s.%s", c.pkg.Types.Name(), name)
				c.reset()
			}
			return true
		},
		FnVisitTypeName: func(I *TypeInspector, d Dir, name string, e *types.TypeName) bool {
			switch d {
			case ENT:
				c.name = name
				c.types = true
				c.mi(1).mf("export interface %s", name)
			case EXT:
				c.faceIntercept()
				if len(c.extended) > 0 {
					c.mf(" extends %s", strings.Join(c.extended.Values(), ","))
				}
				c.mf("{").mn()
				c.declared().
					mi(1).mf("}").mn()
				c.reset()

			}
			return true
		},
		FnVisitVar: func(I *TypeInspector, d Dir, name string, e *types.Var) bool {
			switch d {
			case ENT:
				c.name = name
				c.consts = true
				c.mi(1).mf("export const %s:", name)
			case EXT:
				c.declared().mn()
				c.addEntry(name, "%s.%s", c.pkg.Types.Name(), name)
				c.reset()
			}
			return true
		},
		FnVisitTypeTuple: func(I *TypeInspector, d Dir, o types.Object, x *types.Tuple, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				switch m := mods.LNth(1); true {
				case m.IsParam():
					c.df("(")
					c.pc = x.Len()
				case m.IsResult():
					c.pc = x.Len()
					switch x.Len() {
					case 0:
					case 1:
					default:
						c.df("[")
					}
				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				switch m := mods.LNth(1); true {
				case m.IsParam():
					if !c.funcAlias && (mods.Len() < 3 || mods.LNth(2).IsMethod()) {
						c.df("):")
					} else {
						c.df(")=>")
					}

				case m.IsResult():
					switch x.Len() {
					case 0:
						c.df("void")
					case 1:
					default:
						c.df("]")
					}

				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			}
			return true
		},
		FnVisitTypeVar: func(I *TypeInspector, d Dir, o types.Object, x *types.Var, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				if c.types && mods.LNth(1).IsField() {
					c.di(2).df("%s:", Camel(x.Name()))
					c.field = x.Name()
				} else {
					switch m := mods.LNth(2); true {
					case m.IsParam() || m.IsResult():
						if m.IsParam() {
							c.param = x.Name()
							if c.pc == 1 && c.variadic[len(c.variadic)-1] {
								if c.param == "" {
									c.df("...v%d:", c.pc)
								} else {
									c.df("...%s:", x.Name())
								}
							} else {
								if c.param == "" {
									c.df("v%d:", c.pc)
								} else {
									c.df("%s:", x.Name())
								}
							}

						}
					case m.IsField():

					}
				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				if c.types && mods.LNth(1).IsField() {
					c.dn()
					c.field = ""
				} else {
					switch m := mods.LNth(2); true {
					case m.IsParam() || m.IsResult():
						if c.pc > 1 {
							c.df(",")
						}
						c.param = ""
						c.pc--
						if c.pc == 0 && m.IsParam() && c.variadic[len(c.variadic)-1] {
							c.df("[]")
						}
					case m.IsField():
						c.dn()
					}
				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			}
			return true
		},
		FnVisitTypeFunc: func(I *TypeInspector, d Dir, o types.Object, x *types.Func, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				c.di(2).df("%s", Camel(x.Name()))
				c.method = x.Name()
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				c.dn()
				c.method = ""
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			}
			return true
		},
		FnVisitTypeBasic: func(I *TypeInspector, d Dir, o types.Object, x *types.Basic, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				if c.types && len(seen) == 1 {
					c.use("Alias")
					c.extends("Alias<%s>", c.identType(x.Name(), false))
					//TODO add caster function
				} else {
					c.df("%s", c.identType(x.Name(), false))
				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				if mods.LNth(1).IsMapKey() {
					c.df(",")
				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			}
			return true
		},
		FnVisitTypeMap: func(I *TypeInspector, d Dir, o types.Object, x *types.Map, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				if c.types && len(seen) == 1 {
					c.mapAlias = true
					return true
				}
				c.df("Record<")
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				if c.mapAlias {
					c.mapAlias = false
					c.extends("Record<%s>", c.d.Buffer().String())
					c.d.Buffer().Reset()
					return true
				}
				c.df(">")
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)

			}
			return true
		},
		FnVisitTypeArray: func(I *TypeInspector, d Dir, o types.Object, x *types.Array, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)

			}
			return true
		},
		FnVisitTypeStruct: func(I *TypeInspector, d Dir, o types.Object, x *types.Struct, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				if c.types && len(seen) == 1 {
					if strings.HasSuffix(c.name, "Error") {
						c.use("Struct")
						c.extends("Struct<%s>", c.name)
						c.extends("Error")
					} else {
						for _, face := range I.Implements(o.Type()) {
							if face.Pkg == c.pkg.Types {
								c.extends("%s", face.Name)
							} else {
								c.imported(face.Pkg.Path())
								c.extends("%s.%s", face.Pkg.Name(), face.Name)
							}
						}
						c.use("Struct")
						c.extends("Struct<%s>", c.name)
					}
				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			}
			return true
		},
		FnVisitTypeUnion: func(I *TypeInspector, d Dir, o types.Object, x *types.Union, mods Mods, seen Types) bool {

			switch d {
			case ENT:
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			}
			return true
		},
		FnVisitTypeSignature: func(I *TypeInspector, d Dir, o types.Object, x *types.Signature, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				c.variadic = append(c.variadic, x.Variadic())
				if c.types && len(seen) == 1 && mods.Len() == 1 && mods.LNth(1).IsFunction() {
					c.funcAlias = true
					return true
				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				c.variadic = c.variadic[:len(c.variadic)-1]
				if c.funcAlias {
					c.funcAlias = false
					c.extends("Alias<%s>", c.d.Buffer().String())
					c.d.Reset()
					return true
				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)

			}
			return true
		},
		FnVisitTypeParam: func(I *TypeInspector, d Dir, o types.Object, x *types.TypeParam, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)

			}
			return true
		},
		FnVisitTypePointer: func(I *TypeInspector, d Dir, o types.Object, x *types.Pointer, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				c.use("Ref")
				c.df("Ref<")
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				c.df(">")
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			}
			return true
		},
		FnVisitTypeSlice: func(I *TypeInspector, d Dir, o types.Object, x *types.Slice, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			}
			return true
		},
		FnVisitTypeInterface: func(I *TypeInspector, d Dir, o types.Object, x *types.Interface, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				if c.types && len(seen) == 1 {
					c.use("Proto")
					c.extends("Proto<%s>", c.name)
				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)

			}
			return true
		},
		FnVisitTypeChan: func(I *TypeInspector, d Dir, o types.Object, x *types.Chan, mods Mods, seen Types) bool {

			switch d {
			case ENT:
				c.imported("go")
				switch x.Dir() {
				case types.SendOnly:
					c.df("go.ChanSend<")
				case types.RecvOnly:
					c.df("go.ChanRecv<")
				default:
					c.df("go.Chan<")
				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			case EXT:
				c.df(">")
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)

			}
			return true
		},
		FnVisitTypeNamed: func(I *TypeInspector, d Dir, o types.Object, x *types.Named, mods Mods, seen Types) bool {
			switch d {
			case ENT:
				if c.consts {
					c.named = true
				}
				if c.types && len(seen) == 0 { //!! self
					return true
				} else if c.types && mods.LNth(1).IsEmbedded() { //!! embedded
					if x.Obj().Pkg().Path() == c.pkg.Types.Path() { //!! same package
						c.extends("%s", x.Obj().Name())
					} else {
						c.imported(x.Obj().Pkg().Path())
						c.extends("%s.%s", x.Obj().Pkg().Name(), x.Obj().Name())
					}
					debugf("* %s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
					return false
				}
				if x.Obj().Name() == "error" {
					c.use("error")
					c.df("error")
				} else if x.Obj().Pkg().Path() == c.pkg.Types.Path() {
					c.df("%s", x.Obj().Name())
				} else {
					c.imported(x.Obj().Pkg().Path())
					c.df("%s.%s", x.Obj().Pkg().Name(), x.Obj().Name())
				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
				return false
			case EXT:
				if mods.LNth(1).IsMapKey() {
					c.df(",")
				}
				debugf("%s : %T <<= %T %s || %s", d, o, x, c.summary(), mods)
			}
			return true
		},
	}
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
func (c *context) extends(format string, args ...any) {
	c.extended.Add(fmt.Sprintf(format, args...))
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
func (c *context) summary() string {
	if !debug {
		return ""
	}
	return fmt.Sprintf(":%s:%s:%2d:%4d", c.name, anyOf(c.field, c.method, c.param), c.pc, c.d.Buffer().Len())
}

func (c *context) reset() {
	c.named = false
	c.method = ""
	c.field = ""
	c.param = ""
	c.consts = false
	c.types = false
	c.funcs = false
	c.pc = 0
	c.extended.Clear()
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
			c.use("io")
			c.extends("io.Closer")
		}
	}
}
func (c *context) flush(g *Generator) (err error) {
	x := GetWriter()
	module := ""
	pkgPath := c.pkg.Types.Path()
	pkg := c.pkg.Types.Name()
	dts := strings.Join(fn.SliceMapping(strings.Split(strings.ReplaceAll(c.pkg.Types.Path(), ".", "/"), "/"), strings.ToLower), "_") + ".d.ts"
	mod := dts[:len(dts)-5] + ".go"
	modTest := dts[:len(dts)-5] + "_test.go"
	name := strings.Join(fn.SliceMapping(strings.Split(c.pkg.Types.Path(), "/"), PascalCase), "")
	if strings.ContainsRune(c.pkg.Types.Path(), '.') {
		module = strings.ReplaceAll(c.pkg.Types.Path(), ".", "/")
	} else {
		module = "golang/" + c.pkg.Types.Path()
	}
	//!! gen d.ts
	{

		x.Format("declare module '%s'{", module).LF()
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
		x.Format(`
package %[1]s
import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"%[2]s"`, pkg, pkgPath)
		for s := range c.imports {
			if strings.ContainsRune(s, '.') {
				x.LF().Format("_ \"github.com/ZenLiuCN/engine/%s\"", strings.ReplaceAll(s, ".", "/"))
			} else {
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

func TestSimple(t *testing.T) {
	v := engine.Get()
	defer v.Free()
	fn.Panic1(v.RunTs(
		//language=typescript
		%[2]s))
}

`, pkg, fmt.Sprintf("`"+`
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
func anyOf(an ...string) string {
	for _, s := range an {
		if s != "" {
			return s
		}
	}
	return ""
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
