package main

//https://cs.opensource.google/go/x/tools/+/master:cmd/stringer/stringer.go;drc=daf94608b5e2caf763ba634b84e7a5ba7970e155;l=382
import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go/ast"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"slices"
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
	for name, types := range ts {
		fmt.Printf("%s\n", name)
		slices.SortFunc(types, func(a, b *Type) int {
			return int(a.Type - b.Type)
		})
		for _, t := range types {
			fmt.Printf("\t\t")
			switch t.Type {
			case FuncDecl:
				fmt.Printf("\tfunction %s", t.Decl.(*ast.FuncDecl).Name)
			case MethodDecl:
				fmt.Printf("\tmethod %s", t.Decl.(*ast.FuncDecl).Name)
			case FuncType:
				fmt.Printf("\tFunc %s", t.Spec.Name)
			case StructType:
				fmt.Printf("\tStruct %s", t.Spec.Name)
			case MapType:
				fmt.Printf("\tMap %s", t.Spec.Name)
			case PrimitiveType:
				fmt.Printf("\t%s %s", t.Spec.Name, t.Spec.Type.(*ast.Ident))
			case InterfaceType:
				fmt.Printf("\tInterface %s", t.Spec.Name)
			case ArrayType:
				fmt.Printf("\tArray %s", t.Spec.Name)
			case AliasType:
				fmt.Printf("\tAlias %s %s", t.Spec.Name, t.Spec.Type.(*ast.Ident))
			}
			fmt.Printf("\n")
		}
	}
}
