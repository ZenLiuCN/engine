package main

//https://cs.opensource.google/go/x/tools/+/master:cmd/stringer/stringer.go;drc=daf94608b5e2caf763ba634b84e7a5ba7970e155;l=382
import (
	"bytes"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/urfave/cli/v2"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"path/filepath"
	"slices"
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
	FuncType
	MethodDecl
	FuncDecl
)

type Type struct {
	File   *ast.File
	Decl   ast.Decl
	Object *ast.Object
	Spec   *ast.TypeSpec
	Type   Kind
}

const glob = "global"

func (g *Generator) generate() {
	p := Parse(g.tags, g.files, g.log)
	if len(p) > 1 {
		panic("only one package each time")
	}
	pkg := p[0]
	ts := map[string][]*Type{}
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			set := func(name string, typ *Type) {
				v := ts[name]
				typ.Decl = decl
				typ.File = file
				ts[name] = append(v, typ)
			}
			switch d := decl.(type) {
			case *ast.FuncDecl:
				if !ast.IsExported(d.Name.Name) {
					continue
				}
				if d.Recv == nil {
					set(glob, &Type{
						Object: d.Name.Obj,
						Type:   FuncDecl,
					})
				} else if d.Recv.NumFields() != 1 {
					panic(fmt.Errorf("more than one reciver: %#+v", d))
				} else {
					r := d.Recv.List[0]
					if len(r.Names) == 0 && !ast.IsExported(r.Type.(*ast.Ident).Name) {
						//!! Unexported Type's Exported method
						continue
					}
					switch x := r.Type.(type) {
					case *ast.Ident:
						if ast.IsExported(x.Name) {
							set(x.Name, &Type{
								Object: d.Name.Obj,
								Type:   MethodDecl,
							})
						}
					case *ast.StarExpr:
						n := x.X.(*ast.Ident).Name
						if ast.IsExported(n) {
							set(n, &Type{
								Object: d.Name.Obj,
								Type:   MethodDecl,
							})
						}
					default:
						fmt.Printf("missing %#+v", x)
					}

				}
			case *ast.GenDecl:
				switch d.Tok {
				case token.TYPE:
					for _, spec := range d.Specs {
						if s, ok := spec.(*ast.TypeSpec); ok {
							if !ast.IsExported(s.Name.Name) {
								continue
							}
							switch t := s.Type.(type) {
							case *ast.Ident:
								if t.Obj != nil {
									switch t.Obj.Kind {
									case ast.Fun:
										set(t.Name, &Type{
											Object: t.Obj,
											Type:   FuncType,
										})
									case ast.Typ:
										set(t.Name, &Type{
											Object: t.Obj,
											Type:   AliasType,
										})
									}
								} else {
									set(s.Name.Name, &Type{
										Spec:   s,
										Object: s.Name.Obj,
										Type:   PrimitiveType,
									})
								}
							case *ast.StructType:
								set(s.Name.Name, &Type{
									Object: s.Name.Obj,
									Spec:   s,
									Type:   StructType,
								})
							case *ast.InterfaceType:
								set(s.Name.Name, &Type{
									Object: s.Name.Obj,
									Spec:   s,
									Type:   InterfaceType,
								})
							case *ast.MapType:
								set(s.Name.Name, &Type{
									Object: s.Name.Obj,
									Spec:   s,
									Type:   MapType,
								})
							case *ast.FuncType:
								set(s.Name.Name, &Type{
									Object: s.Name.Obj,
									Spec:   s,
									Type:   FuncType,
								})
							case *ast.ArrayType:
								set(s.Name.Name, &Type{
									Object: s.Name.Obj,
									Spec:   s,
									Type:   ArrayType,
								})
							default:
								fmt.Printf("miss %#+v\n", t)
							}
						}
					}
				default:
					continue
				}
			}
		}
	}
	//!! write d.ts
	imported := fn.NewHashSet[string]()
	gw := NewWriter()
	gw.F("declare module 'go/%s' {\n", pkg.Name)
	iw := NewWriter()
	dw := NewWriter()
	tw := NewWriter()
	for name, types := range ts {
		slices.SortFunc(types, func(a, b *Type) int {
			return int(a.Type - b.Type)
		})
		zero := types[0]
		closure := false
		switch zero.Type {
		case InterfaceType, StructType, MapType, ArrayType:
			dw.F("\n\texport interface %s {\n", name)
			closure = true
		}
		for _, t := range types {
			if closure {
				tw.F("\t")
			}
			switch t.Type {
			case FuncDecl:
				d := t.Decl.(*ast.FuncDecl)
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
				d := t.Decl.(*ast.FuncDecl)
				tw.F("\t%s(", GoFuncToJsFunc(d.Name.Name))
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
			case FuncType:
				d := t.Spec.Type.(*ast.FuncType)
				tw.F("\texport function %s(", GoFuncToJsFunc(t.Spec.Name.Name))
				for i, pa := range d.Params.List {
					if i > 0 {
						tw.F(",")
					}
					if len(pa.Names) == 0 {
						tw.F("v:")
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
			}
			dw.Append(tw)
			tw.Reset()
		}
		if closure {
			dw.F("\n\t}\n")
		} else {
			dw.F("\n")
		}
		iw.Append(dw)
		dw.Reset()
	}
	for k := range imported {
		gw.F("\timport * as %[1]s from 'go/%[1]s'\n", k)
	}
	gw.Append(iw)
	gw.F("\n}")
	fmt.Println(gw.String())
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
func TypeResolve(imported fn.HashSet[string], pkg *packages.Package, file *ast.File, typ ast.Expr, w *SpecWriter, field bool) {
	kd, obj, pg := LookupType(pkg, file, typ)
	n := len(kd) - 1
	if n < 0 {
		panic(fmt.Errorf("unresolved type: %#+v", typ))
	}
	switch kd[n] {
	case TypeKindFunc:
		switch {
		case n == 0:
			w.F("((")
			if typ.(*ast.FuncType).Params.NumFields() != 0 {
				for i, field := range typ.(*ast.FuncType).Params.List {
					if i > 0 {
						w.F(",")
					}
					TypeResolve(imported, pkg, file, field.Type, w, false)
				}
			}
			w.F(")")
			w.F("=>")
			if typ.(*ast.FuncType).Results.NumFields() > 1 {
				w.F("(")
				for i, field := range typ.(*ast.FuncType).Results.List {
					if i > 0 {
						w.F("|")
					}
					TypeResolve(imported, pkg, file, field.Type, w, false)
				}
				w.F(")[]")
			} else if typ.(*ast.FuncType).Results.NumFields() == 1 {
				TypeResolve(imported, pkg, file, typ.(*ast.FuncType).Results.List[0].Type, w, false)
			} else {
				w.F("void")
			}
			w.F(")")
		default:
			fmt.Printf("\nmissing func: %#+v , %#+v , %#+v for %#+v \n", kd, obj, pkg, typ)
		}
	case TypeKindArray:
		switch {
		case n == 0:
			w.F("%s[]", obj)
		default:
			fmt.Printf("\nmissing array: %#+v , %#+v , %#+v for %#+v \n", kd, obj, pkg, typ)
		}
	case TypeKindSelector:
		switch {
		case n == 0:
			imported[pg.Name()] = fn.Nothing
			w.F("%s.%s", pg.Name(), obj.Name())
		case n == 1 && kd[n-1] == TypeKindStar:
			imported[pg.Name()] = fn.Nothing
			w.F("%s.%s/*Pointer*/", pg.Name(), obj.Name())
		default:
			fmt.Printf("\nmissing selector: %#+v , %#+v , %#+v for %#+v \n", kd, obj, pkg, typ)
		}
	case TypeKindIdent:
		switch {
		case n == 0:
			w.F("%s", GoIdentToTs(typ.(*ast.Ident).Name, false))
		case n == 1 && kd[n-1] == TypeKindStar:
			w.F("%s/*pointer*/", GoIdentToTs(typ.(*ast.StarExpr).X.(*ast.Ident).Name, false))
		case n == 1 && kd[n-1] == TypeKindArray:
			w.F("%s", GoIdentToTs(typ.(*ast.ArrayType).Elt.(*ast.Ident).Name, true))
		case n == 2 && kd[n-1] == TypeKindStar && kd[n-2] == TypeKindArray:
			w.F("%s", GoIdentToTs(typ.(*ast.ArrayType).Elt.(*ast.StarExpr).X.(*ast.Ident).Name, true))
		case n == 2 && kd[0] == TypeKindMap && kd[n-1] == TypeKindIdent:
			m := typ.(*ast.MapType)
			k := m.Key.(*ast.Ident)
			v := m.Value.(*ast.Ident)
			if field {
				w.F("[key:%s]:%s", GoIdentToTs(k.Name, false), GoIdentToTs(v.Name, false))
			} else {
				w.F("Record<%s,%s>", GoIdentToTs(k.Name, false), GoIdentToTs(v.Name, false))
			}
		case n == 3 && kd[0] == TypeKindMap && kd[1] == TypeKindIdent && kd[2] == TypeKindArray:
			m := typ.(*ast.MapType)
			k := m.Key.(*ast.Ident)
			v := m.Value.(*ast.ArrayType).Elt.(*ast.Ident)
			if field {
				w.F("[key:%s]:%s", GoIdentToTs(k.Name, false), GoIdentToTs(v.Name, true))
			} else {
				w.F("Record<%s,%s>", GoIdentToTs(k.Name, false), GoIdentToTs(v.Name, true))
			}
		default:
			fmt.Printf("\nmissing gen ident: %#+v , %#+v , %#+v for %#+v \n", kd, obj, pkg, typ)
		}
	case TypeKindStruct:
		switch {
		case n == 0:
			t := typ.(*ast.StructType)
			if field {
				if t.Fields.NumFields() != 0 {
					for i, f := range t.Fields.List {
						if len(f.Names) == 0 {
							TypeResolve(imported, pkg, file, f.Type, w, true)
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
							}
						}
					}
				}
			} else {

			}
		case n == 1 && kd[0] == TypeKindStar:
		default:
			fmt.Printf("\nfound struct %#+v , %#+v , %#+v for %#+v \n", kd, obj, pkg, typ)
		}

	default:
		fmt.Printf("\nmissing %#+v , %#+v , %#+v for %#+v \n", kd, obj, pkg, typ)
	}
}
func GoIdentToTs(s string, array bool) string {
	switch s {
	case "byte", "uint8", "uint16", "uint32":
		if array {
			return fmt.Sprintf("U%sArray", s[1:])
		} else {
			return "number/*positive*/"
		}
	case "int8", "int16", "int32":
		if array {
			return fmt.Sprintf("I%sArray", s[1:])
		} else {
			return "number/*int*/"
		}
	case "rune":
		if array {
			return "string"
		} else {
			return "number/*rune*/"
		}
	case "int", "int64":
		if array {
			return "number[]/*int*/"
		} else {
			return "number/*int*/"
		}
	case "uint64", "uint":
		if array {
			return "number[]/*uint*/"
		} else {
			return "number/*uint*/"
		}
	case "bool":
		if array {
			return "boolean[]"
		} else {
			return "boolean"
		}
	default:
		if array {
			return s + "[]"
		} else {
			return s
		}

	}
}
