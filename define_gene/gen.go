package main

//https://cs.opensource.google/go/x/tools/+/master:cmd/stringer/stringer.go;drc=daf94608b5e2caf763ba634b84e7a5ba7970e155;l=382
import (
	"bytes"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/urfave/cli/v2"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Usage:   "show version",
		Aliases: []string{"v"},
	}
	err := (&cli.App{
		UseShortOptionHandling: true,
		Name:                   "Define Generator",
		Version:                "v0.0.1",
		Usage:                  "Generate engine define from go source",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "type",
				Usage:   "type names of current file or directory",
				Aliases: []string{"t"},
			},
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "debug generator",
				Aliases: []string{"d"},
			},
			&cli.StringSliceFlag{
				Name:    "tags",
				Usage:   "tags to apply",
				Aliases: []string{"g"},
			},
		},
		Suggest:              true,
		EnableBashCompletion: true,
		Action: func(c *cli.Context) error {
			var file []string
			if c.Args().Len() == 0 {
				file = append(file, ".")
			} else {
				file = append(file, c.Args().Slice()...)
			}
			tags := c.StringSlice("tags")
			var dir string
			if len(file) == 1 && IsDir(file[0]) {
				dir = file[0]
			} else if len(tags) != 0 {
				log.Fatal("--tags can only applies with directory")
			} else {
				dir = filepath.Dir(file[0])
			}
			g := new(Generator)
			{
				g.dir = dir
				g.tags = tags
				g.files = file
				g.types = c.StringSlice("type")
				g.out = c.String("out")
				if c.Bool("debug") {
					g.log = log.Printf
				}
			}
			g.generate()
			return nil
		},
	}).Run(os.Args)
	if err != nil {
		panic(err)
	}
}

type Generator struct {
	dir   string
	tags  []string
	files []string
	types []string
	out   string
	log   func(format string, a ...any)
}

//go:generate stringer -type=Kind
type Kind int

const (
	StructType Kind = iota
	MapType
	PrimitiveType
	InterfaceType
	ArrayType
	AliasType
	ChanType
	FuncType
	MethodDecl
	FuncDecl
	IdentTypeFunc
	IdentType
)

type Type struct {
	File *ast.File
	Type Kind

	Spec       *ast.TypeSpec
	Interface  *ast.InterfaceType
	Ident      *ast.Ident
	Struct     *ast.StructType
	Map        *ast.MapType
	Array      *ast.ArrayType
	Func       *ast.FuncType
	FuncDecl   *ast.FuncDecl
	IdentRecv  *ast.Ident
	Recv       *ast.Field
	MethodDecl *ast.FuncDecl
	StarRecv   *ast.StarExpr
	ChanType   *ast.ChanType
}

const glob = "global"

//region Declare

type WalkDecl struct {
	found func(string, *Type)
}

func (w WalkDecl) set(name string, t *Type) {
	w.found(name, t)
}
func (w WalkDecl) FuncDecl(d *ast.FuncDecl) {
	w.set(glob, &Type{
		FuncDecl: d,
		Type:     FuncDecl,
	})
}
func (w WalkDecl) ChanTypeSpec(s *ast.TypeSpec, t *ast.ChanType) {
	w.set(glob, &Type{
		Spec:     s,
		ChanType: t,
		Type:     ChanType,
	})
}
func (w WalkDecl) MethodDecl(d *ast.FuncDecl, r *ast.Field) bool {
	return true
}

func (w WalkDecl) IdentTypeSpec(s *ast.TypeSpec, t *ast.Ident) {
	if t.Obj != nil {
		switch t.Obj.Kind {
		case ast.Fun:
			w.set(t.Name, &Type{
				Ident: t,
				Type:  IdentTypeFunc,
			})
		case ast.Typ:
			w.set(t.Name, &Type{
				Spec:  s,
				Ident: t,
				Type:  AliasType,
			})
		}
	} else {
		w.set(s.Name.Name, &Type{
			Spec:  s,
			Ident: t,
			Type:  PrimitiveType,
		})
	}
}

func (w WalkDecl) StructTypeSpec(s *ast.TypeSpec, t *ast.StructType) {
	w.set(s.Name.Name, &Type{
		Struct: t,
		Spec:   s,
		Type:   StructType,
	})
}

func (w WalkDecl) InterfaceTypeSpec(s *ast.TypeSpec, t *ast.InterfaceType) {
	w.set(s.Name.Name, &Type{
		Interface: t,
		Spec:      s,
		Type:      InterfaceType,
	})
}

func (w WalkDecl) MapTypeSpec(s *ast.TypeSpec, t *ast.MapType) {
	w.set(s.Name.Name, &Type{
		Map:  t,
		Spec: s,
		Type: MapType,
	})
}

func (w WalkDecl) ArrayTypeSpec(s *ast.TypeSpec, t *ast.ArrayType) {
	w.set(s.Name.Name, &Type{
		Array: t,
		Spec:  s,
		Type:  ArrayType,
	})
}

func (w WalkDecl) FuncTypeSpec(s *ast.TypeSpec, t *ast.FuncType) {
	w.set(s.Name.Name, &Type{
		Func: t,
		Spec: s,
		Type: FuncType,
	})
}

func (w WalkDecl) MethodDeclStarRecv(d *ast.FuncDecl, r *ast.Field, t *ast.StarExpr) {
	w.set(t.X.(*ast.Ident).Name, &Type{
		MethodDecl: d,
		Recv:       r,
		StarRecv:   t,
		Type:       MethodDecl,
	})
}

func (w WalkDecl) MethodDeclIdentRecv(d *ast.FuncDecl, r *ast.Field, t *ast.Ident) {
	w.set(t.Name, &Type{
		MethodDecl: d,
		Recv:       r,
		IdentRecv:  t,
		Type:       MethodDecl,
	})
}

//endregion

func (g *Generator) generate() {
	var flags []string
	if len(g.tags) > 0 {
		flags = append(flags, fmt.Sprintf("-tags=%s", strings.Join(g.tags, " ")))
	}
	p := ParseTypeInfo(flags, g.files, g.log)
	if len(p) > 1 {
		panic("only one package each time")
	}
	pkg := p[0]
	ts := map[string][]*Type{}
	for _, file := range pkg.Syntax {
		CaseDecl(file, &ExportedDeclCases{&WalkDecl{func(s string, t *Type) {
			t.File = file
			v := ts[s]
			v = append(v, t)
			ts[s] = v
		}}})
	}
	//!! write d.ts
	imported := fn.NewHashSet[string]()
	fw := NewWriter() //file
	fw.F("declare module 'go/%s' {\n", pkg.Name)
	iw := NewWriter()
	dw := NewWriter() // declare
	ew := NewWriter() // elements
	tw := NewWriter() // element
	for name, types := range ts {
		slices.SortFunc(types, func(a, b *Type) int {
			return int(a.Type - b.Type)
		})
		closure := name != glob
		closed := false
		if closure {
			switch types[0].Type {
			case MapType, StructType, InterfaceType:
				dw.F("\t export interface %s", name)
				closed = true
			default:
				if len(types) > 1 && !(types[1].Type == MethodDecl && types[1].MethodDecl.Name.Name == "string") { //!! exclude alias with only Stringer
					dw.F("\t export interface %s", name)
					closed = true
				}
			}
		}
		for _, t := range types {
			if closure {
				tw.F("\t")
			}
			switch t.Type {
			case FuncDecl:
				d := t.FuncDecl
				tw.F("\texport function %s(", GoFuncToJsFunc(d.Name.Name))
				for i, pa := range d.Type.Params.List {
					if i > 0 {
						tw.F(",")
					}
					tw.F("%s:", pa.Names[0].Name)
					TypeResolve(imported, pkg, t.File, pa.Type, tw, false)
				}
				tw.F(")")
				switch d.Type.Results.NumFields() {
				case 0:
				case 1:
					tw.F(":")
					TypeResolve(imported, pkg, t.File, d.Type.Results.List[0].Type, tw, false)
				default:
					tw.F(":(")
					for i, pa := range d.Type.Results.List {
						if i > 0 {
							tw.F("|")
						}
						TypeResolve(imported, pkg, t.File, pa.Type, tw, false)
					}
					tw.F(")[]")
				}
				tw.F("\n")
			case MethodDecl:
				d := t.MethodDecl
				tw.F("\t%s(", GoFuncToJsFunc(d.Name.Name))
				for i, pa := range d.Type.Params.List {
					if i > 0 {
						tw.F(",")
					}
					if _, ok := pa.Type.(*ast.Ellipsis); ok {
						tw.F("...%s:", pa.Names[0].Name)
					} else {
						tw.F("%s:", pa.Names[0].Name)
					}
					TypeResolve(imported, pkg, t.File, pa.Type, tw, false)
				}
				tw.F(")")
				switch d.Type.Results.NumFields() {
				case 0:
				case 1:
					tw.F(":")
					TypeResolve(imported, pkg, t.File, d.Type.Results.List[0].Type, tw, false)
				default:
					tw.F(":(")
					for i, pa := range d.Type.Results.List {
						if i > 0 {
							tw.F("|")
						}
						TypeResolve(imported, pkg, t.File, pa.Type, tw, false)
					}
					tw.F(")[]")
				}
				tw.F("\n")
			case FuncType:
				d := t.Spec.Type.(*ast.FuncType)
				tw.F("\texport type %s(", GoFuncToJsFunc(t.Spec.Name.Name))
				for i, pa := range d.Params.List {
					if i > 0 {
						tw.F(",")
					}
					if len(pa.Names) == 0 {
						tw.F("v%d:", i)
					} else {
						for i, n := range pa.Names {
							if i > 0 {
								tw.F(",")
							}
							tw.F("%s", n.Name)
						}
						tw.F(":")
					}
					TypeResolve(imported, pkg, t.File, pa.Type, tw, false)
				}
				tw.F(")")
				switch d.Results.NumFields() {
				case 0:
				case 1:
					tw.F(":")
					TypeResolve(imported, pkg, t.File, d.Results.List[0].Type, tw, false)
				default:
					tw.F(":(")
					for i, pa := range d.Results.List {
						if i > 0 {
							tw.F("|")
						}
						TypeResolve(imported, pkg, t.File, pa.Type, tw, false)
					}
					tw.F(")[]")
				}
				tw.F("\n")
			case StructType:
				st := t.Spec.Type.(*ast.StructType)
				if st.Fields.NumFields() != 0 {
					for i, f := range st.Fields.List {
						//!! embedded
						if len(f.Names) == 0 {
							TypeResolve(imported, pkg, t.File, f.Type, tw, true)
						} else {
							x := -1
							for _, name := range f.Names {
								if ast.IsExported(name.Name) {
									x++
								}
							}
							if x >= 0 {
								if i > 0 {
									tw.F("\n")
									tw.F("\t\t")
								} else {
									tw.F("\t")
								}
								x = -1
								for _, name := range f.Names {
									if !ast.IsExported(name.Name) {
										continue
									}
									x++
									if x > 0 {
										tw.F(",")
									}
									tw.F("%s", GoFuncToJsFunc(name.Name))
								}
								tw.F(":")
								TypeResolve(imported, pkg, t.File, f.Type, tw, false)
							}
						}
					}
				}
				tw.F("\n")
			case MapType:
				tw.F("\t")
				TypeResolve(imported, pkg, t.File, t.Spec.Type, tw, true)
				tw.F("\n")
			case PrimitiveType:
				fmt.Printf("%s %s \n", t.Spec.Name, t.Spec.Type.(*ast.Ident))
			case InterfaceType:
				fmt.Printf("Interface %s \n", t.Spec.Name)
			case ArrayType:
				fmt.Printf("Array %s \n", t.Spec.Name)
			case AliasType:
				fmt.Printf("Alias %s %s \n", t.Spec.Name, t.Spec.Type.(*ast.Ident))
			default:
				println(fmt.Sprintf("miss %#+v", t))
			}
			ew.Append(tw)
			tw.Reset()
		}
		if closed {
			dw.F("{\n")
			dw.Append(ew)
			ew.Reset()
			dw.F("\n\t}\n")
		} else {
			dw.Append(ew)
			ew.Reset()
			dw.F("\n")
		}
		iw.Append(dw)
		dw.Reset()
	}
	for k := range imported {
		fw.F("\timport * as %[1]s from 'go/%[1]s'\n", k)
	}
	fw.Append(iw)
	fw.F("\n}")
	fmt.Println(fw.String())
}

func GoFuncToJsFunc(n string) string {
	buf := new(bytes.Buffer)
	u0 := false
	u := false
	rs := []rune(n)
	mx := len(rs) - 1
	for i := 0; i < len(rs); i++ {
		r := rs[i]
		switch {
		case i == 0 && unicode.IsUpper(r):
			buf.WriteRune(unicode.ToLower(r))
			u0 = true
			u = true
		case u0 && u && unicode.IsUpper(r):
			if i < mx && unicode.IsLower(rs[i+1]) {
				buf.WriteRune(r)
				u = false
			} else {
				buf.WriteRune(unicode.ToLower(r))
			}
		case unicode.IsUpper(r):
			u = true
			buf.WriteRune(r)
		case unicode.IsLower(r):
			u = false
			buf.WriteRune(r)
		default:
			buf.WriteRune(r)
		}
	}

	return buf.String()
}
func TypeResolve(imported fn.HashSet[string], pkg *packages.Package, file *ast.File, typ ast.Expr, w *SpecWriter, field bool) AstNode {
	ca := &WalkWriter{
		Imported:       imported,
		pkg:            pkg,
		field:          field,
		SpecWriter:     w,
		DepthTypeCases: &DepthTypeCases{},
	}
	ca.u = []TypeCases{ca.DepthTypeCases, ca}
	CaseType(pkg, typ, ca.u)
	l, _ := ca.Last()
	return l.Node
}
func GoIdentToTs(s string, array bool) string {
	switch s {
	case "byte", "uint8", "uint16", "uint32":
		if array {
			return fmt.Sprintf("U%sArray", s[1:])
		} else {
			return fmt.Sprintf("/*%s*/number", s)
		}
	case "int8", "int16", "int32":
		if array {
			return fmt.Sprintf("I%sArray", s[1:])
		} else {
			return fmt.Sprintf("/*%s*/number", s)
		}
	case "rune":
		if array {
			return "/*rune[]*/string"
		} else {
			return "/*rune*/number"
		}
	case "int", "int64":
		if array {
			return fmt.Sprintf("/*%s*/number[]", s)
		} else {
			return fmt.Sprintf("/*%s*/number", s)
		}
	case "uint64", "uint":
		if array {
			return fmt.Sprintf("/*%s*/number[]", s)
		} else {
			return fmt.Sprintf("/*%s*/number", s)
		}
	case "bool":
		if array {
			return "boolean[]"
		} else {
			return "boolean"
		}
	case "error":
		if array {
			return "Error[]"
		}
		return "Error"
	default:
		if array {
			return s + "[]"
		} else {
			return s
		}

	}
}

//region Writer

type WalkWriter struct {
	Imported fn.HashSet[string]
	pkg      *packages.Package
	field    bool
	decl     *SpecWriter
	*SpecWriter
	*DepthTypeCases
	u UnionTypeCases
}

func (w *WalkWriter) isArray() bool {
	d, ok := w.Last()
	if !ok {
		return false
	}
	return d.Node == AstArrayType
}
func (w *WalkWriter) IdentType(t *ast.Ident) {
	w.F("%s", GoIdentToTs(t.Name, w.isArray()))
}

func (w *WalkWriter) SelectorType(t *ast.SelectorExpr, target types.Object) {
	w.Imported.Add(target.Pkg().Name())
	w.F("%s.%s", target.Pkg().Name(), target.Name())

}

func (w *WalkWriter) StarType(t *ast.StarExpr) {
	w.F("/*Pointer*/")
	CaseType(w.pkg, t.X, w.u)
	w.Pop()

}

func (w *WalkWriter) FuncType(t *ast.FuncType) {
	w.F("((")
	if t.Params.NumFields() != 0 {
		for i, f := range t.Params.List {
			if i > 0 {
				w.F(",")
			}
			w.F("v%d:", i)
			CaseType(w.pkg, f.Type, w.u)
			w.Pop()
		}
	}
	w.F(")")
	w.F("=>")
	if t.Results.NumFields() > 1 {
		w.F("(")
		for i, f := range t.Results.List {
			if i > 0 {
				w.F("|")
			}
			CaseType(w.pkg, f.Type, w.u)
			w.Pop()
		}
		w.F(")[]")
	} else if t.Results.NumFields() == 1 {
		CaseType(w.pkg, t.Results.List[0].Type, w.u)
		w.Pop()
	} else {
		w.F("void")
	}
	w.F(")")
}

func (w *WalkWriter) ArrayType(t *ast.ArrayType) {
	CaseType(w.pkg, t.Elt, w.u)
	l, ok := w.Pop()
	if ok && l.Node != AstIdent {
		w.F("[]")
	}
}

func (w *WalkWriter) MapType(t *ast.MapType) {
	if w.field {
		//w.F("[key:%s]:%s", GoIdentToTs(k.Name, false), GoIdentToTs(v.Name, false))
		w.F("[key:")
		CaseType(w.pkg, t.Key, w.u)
		w.Pop()
		w.F("]:")
		CaseType(w.pkg, t.Value, w.u)
		w.Pop()
	} else {
		//w.F("Record<%s,%s>", GoIdentToTs(k.Name, false), GoIdentToTs(v.Name, false))
		w.F("Record<")
		CaseType(w.pkg, t.Key, w.u)
		w.Pop()
		w.F(",")
		CaseType(w.pkg, t.Value, w.u)
		w.Pop()
		w.F(">")
	}
}

func (w *WalkWriter) StructType(t *ast.StructType) {
	if w.field {
		if t.Fields.NumFields() != 0 {
			for i, f := range t.Fields.List {
				if len(f.Names) == 0 {
					//TODO embedded
				} else {
					x := -1
					for _, name := range f.Names {
						if ast.IsExported(name.Name) {
							x++
						}
					}
					if x >= 0 {
						if i > 0 {
							w.F("\t\t")
						}
						x = -1
						for _, name := range f.Names {
							if !ast.IsExported(name.Name) {
								continue
							}
							x++
							if x > 0 {
								w.F(",")
							}
							w.F("%s", GoFuncToJsFunc(name.Name))
						}
						CaseType(w.pkg, f.Type, w.u)
						w.Pop()
					}
				}
			}
		}
	} else {
		if t.Fields.NumFields() == 0 {
			w.F("/*struct{}{}*/void")
			w.Pop()
			return
		}
		panic(fmt.Sprintf("should not access structType: %#+v", w))
	}
}

func (w *WalkWriter) ChanType(t *ast.ChanType) {
	log.Printf("chan type found: %#+v\n", t)
	w.Imported.Add("chan")
	if t.Dir == ast.RECV {
		w.F("chan.ChanRecv<")
	} else if t.Dir == ast.SEND {
		w.F("chan.ChanSend<")
	} else {
		w.F("chan.Chan<")
	}
	CaseType(w.pkg, t.Value, w.u)
	w.Pop()
	w.F(">")
}
func (w *WalkWriter) Ellipsis(t *ast.Ellipsis) {
	if w.field {
		panic(fmt.Errorf("found ellipsis in field: %#+v", w))
	}
	CaseType(w.pkg, t.Elt, w.u)
	l, ok := w.Pop()
	if ok && l.Node != AstIdent {
		w.F("[]")
	}
}

//endregion
