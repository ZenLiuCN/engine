package main

//https://cs.opensource.google/go/x/tools/+/master:cmd/stringer/stringer.go;drc=daf94608b5e2caf763ba634b84e7a5ba7970e155;l=382
import (
	"bytes"
	"fmt"
	"github.com/urfave/cli/v2"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
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
	p := ParseTypeInfo(flags, g.files, IdentityLoadMode, g.log)
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
	fw := NewWriter() //file
	fw.F("declare module 'go/%s' {\n", pkg.Name)
	iw := NewWriter()
	dw := NewWriter()  // declared
	exw := NewWriter() // extend
	ew := NewWriter()  // elements of declared
	tw := NewWriter()  // element
	for name, types := range ts {
		slices.SortFunc(types, func(a, b *Type) int {
			return int(a.Type - b.Type)
		})
		closure := name != glob
		closed := false
		if closure {
			switch types[0].Type {
			case MapType, StructType, InterfaceType:
				dw.F("\t export interface /*%s*/ %s", types[0].Type, name)
				closed = true
			default:
				if len(types) > 1 && !(types[1].Type == MethodDecl && types[1].MethodDecl.Name.Name == "string") { //!! exclude alias with only Stringer
					dw.F("\t export interface %s", name)
					closed = true
				}
			}
		}
		for _, t := range types {
			switch t.Type {
			case FuncDecl:
				d := t.FuncDecl
				tw.F("\texport function %s ", GoFuncToJsFunc(d.Name.Name))
				FuncDeclWrite(d.Type, NewWalkWriter(pkg, false, exw, tw, nil))
				tw.F("\n")
			case MethodDecl:
				if closure {
					tw.F("\t")
				}
				d := t.MethodDecl
				tw.F("\t %s ", GoFuncToJsFunc(d.Name.Name))
				FuncDeclWrite(d.Type, NewWalkWriter(pkg, false, exw, tw, nil))
				tw.F("\n")
			case FuncType:
				//!! write as top level
				d := t.Spec.Type.(*ast.FuncType)
				iw.F("\texport type %s=", GoFuncToJsFunc(t.Spec.Name.Name))
				FuncTypeWrite(d, NewWalkWriter(pkg, false, exw, iw, nil))
				iw.F("\n")
			case StructType:
				if closure {
					tw.F("\t")
				}
				StructTypeWrite(t.Struct, NewWalkWriter(pkg, true, exw, tw, nil))
				tw.F("\n")
			case MapType:
				if closure {
					tw.F("\t")
				}
				tw.F("\t")
				MapTypeWrite(t.Map, NewWalkWriter(pkg, true, exw, tw, nil))
				tw.F("\n")
			case PrimitiveType:
				if closure {
					tw.F("\t")
				}
				fmt.Printf("Primitive %s %s \n", t.Spec.Name, t.Spec.Type.(*ast.Ident))
			case InterfaceType:
				if closure {
					tw.F("\t")
				}
				switch t.Interface.Methods.NumFields() {
				case 0:
				default:
					for _, field := range t.Interface.Methods.List {
						if len(field.Names) > 0 {
							for i, ident := range field.Names {
								if i > 0 {
									tw.F(",")
								}
								tw.F("%s", GoFuncToJsFunc(ident.Name))
							}
							switch t := field.Type.(type) {
							case *ast.SelectorExpr:
								TypeResolve(pkg, field.Type, tw, exw, false, nil)
							case *ast.FuncType:
								FuncDeclWrite(t, NewWalkWriter(pkg, false, exw, tw, nil))
							default:
								panic(fmt.Errorf("miss %#+v", t))
							}

						} else {
							TypeResolve(pkg, field.Type, tw, exw, true, nil)
						}
					}

				}
			case ArrayType:
				if closure {
					tw.F("\t")
				}
				fmt.Printf("Array %s \n", t.Spec.Name)
			case AliasType:
				if closure {
					tw.F("\t")
				}
				fmt.Printf("Alias %s %s \n", t.Spec.Name, t.Spec.Type.(*ast.Ident))
			default:
				println(fmt.Sprintf("miss %#+v", t))
			}
			ew.Append(tw)
			tw.Reset()
		}
		if exw.Len() > 0 {
			dw.F(" extends %s", exw.String())
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
	for k := range iw.Imports() {
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
func TypeResolve(pkg *packages.Package, typ ast.Expr, w Writer, d Writer, field bool, ca *WalkWriter) (*WalkWriter, AstNode) {
	if ca == nil {
		ca = &WalkWriter{
			pkg:            pkg,
			field:          field,
			decl:           d,
			Writer:         w,
			DepthTypeCases: &DepthTypeCases{},
		}
		ca.u = []TypeCases{ca.DepthTypeCases, ca}
	} else if w != nil || d != nil || ca.field != field {
		f := ca.field
		ca.field = field
		defer func() {
			ca.field = f
		}()
		if w != nil {
			o := ca.Writer
			ca.Writer = w
			defer func() {
				ca.Writer = o
			}()
		}
		if d != nil {
			o := ca.decl
			ca.decl = d
			defer func() {
				ca.decl = o
			}()
		}
	}

	CaseType(pkg, typ, ca.u)
	l, _ := ca.Last()
	return ca, l.Node
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

const (
	AstIdentError = AstNodeBuiltInMax + 1
)

type WalkWriter struct {
	pkg   *packages.Package
	field bool
	decl  Writer
	Writer
	*DepthTypeCases
	u UnionTypeCases
}

func NewWalkWriter(pkg *packages.Package, field bool, decl Writer, writer Writer, depthTypeCases *DepthTypeCases) (r *WalkWriter) {
	r = &WalkWriter{pkg: pkg, field: field, decl: decl, Writer: writer, DepthTypeCases: depthTypeCases}
	if r.DepthTypeCases == nil {
		r.DepthTypeCases = &DepthTypeCases{}
	}
	r.u = []TypeCases{r.DepthTypeCases, r}
	return r
}

func (w *WalkWriter) isArray() bool {
	d, ok := w.Last()
	if !ok {
		return false
	}
	return d.Node == AstArrayType
}
func (w *WalkWriter) IdentType(t *ast.Ident) {
	if t.Name == "error" {
		w.DepthTypeCases.ReplaceLast(Layer{
			Node:  AstIdentError,
			Ident: t,
		})
	} else if v, ok := w.pkg.TypesInfo.Uses[t]; ok {
		if v.Pkg() != nil && v.Pkg().Name() != w.pkg.Name { //!! not universe or current scope
			panic(fmt.Sprintf("used %#+v", v))
		}
	} else if v, ok := w.pkg.TypesInfo.Defs[t]; ok {
		panic(fmt.Sprintf("defined %#+v", v))
	}

	w.F("%s", GoIdentToTs(t.Name, w.isArray()))
}
func (w *WalkWriter) InterfaceType(t *ast.InterfaceType) {
	panic(fmt.Errorf("not expect interface type: %#+v", t))
}

func (w *WalkWriter) SelectorType(t *ast.SelectorExpr, target types.Object) {
	w.Imports().Add(target.Pkg().Name())
	if w.field {
		if w.decl.Len() != 0 {
			w.decl.F(",")
		}
		w.decl.F("%s.%s", target.Pkg().Name(), target.Name())
	} else {
		w.F("%s.%s", target.Pkg().Name(), target.Name())
	}

}

func (w *WalkWriter) StarType(t *ast.StarExpr) {
	w.F("/*Pointer*/")
	CaseType(w.pkg, t.X, w.u)
	w.Pop()

}

func (w *WalkWriter) FuncType(t *ast.FuncType) {
	/*w.F("((")
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
	w.F(")")*/
	FuncTypeWrite(t, w)
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
	w.Imports().Add("chan")
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

// endregion
var (
	writer = sync.Pool{New: func() any {
		return NewWriter()
	}}
)

func GetWriter() Writer {
	return writer.Get().(Writer)
}
func FreeWriter(w Writer) {
	w.Reset()
	writer.Put(w)
}
func FuncTypeWrite(d *ast.FuncType, w *WalkWriter) {
	temp := GetWriter()
	defer func() {
		FreeWriter(temp)
	}()
	w.F("((")
	if d.Params.NumFields() > 0 {
		for i, pa := range d.Params.List {
			if i > 0 {
				w.F(",")
			}
			_, n := TypeResolve(w.pkg, pa.Type, temp, nil, false, w)
			switch n {
			case AstIdentError, AstStarExpr:
				switch len(pa.Names) {
				case 0:
					w.F("v%d:%s", i, temp.String())
				case 1:
					w.F("%s:%s", pa.Names[0].Name, temp.String())
				default:
					for x, n := range pa.Names {
						if x > 0 {
							w.F(",")
						}

						w.F("%s", n)
					}
					w.F(":%s", temp.String())
				}
			case AstEllipsis:
				switch len(pa.Names) {
				case 0:
					w.F("...v%d:%s", i, temp.String())
				case 1:
					w.F("...%s:%s", pa.Names[0].Name, temp.String())
				default:
					panic(fmt.Errorf("ellipsis with multiple names: %#+v", pa))
				}
			default:
				switch len(pa.Names) {
				case 0:
					w.F("v%d:%s", i, temp.String())
				case 1:
					w.F("%s:%s", pa.Names[0].Name, temp.String())
				default:
					for x, n := range pa.Names {
						if x > 0 {
							w.F(",")
						}
						w.F("%s", n)
					}
					w.F(":%s", temp.String())
				}
			}
			w.Pop()
			temp.Reset()
		}
	}
	w.F(")")
	switch d.Results.NumFields() {
	case 0:
		w.F("=>void")
	case 1:
		w.F("=>")
		_, n := TypeResolve(w.pkg, d.Results.List[0].Type, temp, nil, false, w)
		switch n {
		case AstIdentError, AstStarExpr:
			w.F("(%s|undefined)", temp.String())
		default:
			w.F("%s", temp.String())
		}
		w.Pop()
		temp.Reset()
	default:
		w.F("=>(")
		for i, pa := range d.Results.List {
			if i > 0 {
				w.F("|")
			}
			_, n := TypeResolve(w.pkg, pa.Type, temp, nil, false, w)
			switch n {
			case AstIdentError, AstStarExpr:
				w.F("(%s|undefined)", temp.String())
			default:
				w.F("%s", temp.String())
			}
			w.Pop()
			temp.Reset()
		}
		w.F(")[]")
	}
	w.F(")")
}
func FuncDeclWrite(d *ast.FuncType, w *WalkWriter) {
	temp := GetWriter()
	defer func() {
		FreeWriter(temp)
	}()
	w.F("(")
	if d.Params.NumFields() > 0 {
		for i, pa := range d.Params.List {
			if i > 0 {
				w.F(",")
			}
			_, n := TypeResolve(w.pkg, pa.Type, temp, nil, false, w)
			switch n {
			case AstIdentError, AstStarExpr:
				switch len(pa.Names) {
				case 0:
					w.F("v%d:%s", i, temp.String())
				case 1:
					w.F("%s:%s", pa.Names[0].Name, temp.String())
				default:
					for x, n := range pa.Names {
						if x > 0 {
							w.F(",")
						}
						w.F("%s", n)
					}
					w.F(":%s", temp.String())
				}
			case AstEllipsis:
				switch len(pa.Names) {
				case 0:
					w.F("...v%d:%s", i, temp.String())
				case 1:
					w.F("...%s:%s", pa.Names[0].Name, temp.String())
				default:
					panic(fmt.Errorf("ellipsis with multiple names: %#+v", pa))
				}
			default:
				switch len(pa.Names) {
				case 0:
					w.F("v%d:%s", i, temp.String())
				case 1:
					w.F("%s:%s", pa.Names[0].Name, temp.String())
				default:
					for x, n := range pa.Names {
						if x > 0 {
							w.F(",")
						}
						w.F("%s", n)
					}
					w.F(":%s", temp.String())
				}
			}
			w.Pop()
			temp.Reset()
		}
	}
	w.F(")")
	switch d.Results.NumFields() {
	case 0:
		w.F(":void")
	case 1:
		w.F(":")
		_, n := TypeResolve(w.pkg, d.Results.List[0].Type, temp, nil, false, w)
		switch n {
		case AstIdentError, AstStarExpr:
			w.F("(%s|undefined)", temp.String())
		default:
			w.F("%s", temp.String())
		}
		w.Pop()
		temp.Reset()
	default:
		w.F(":(")
		for i, pa := range d.Results.List {
			if i > 0 {
				w.F("|")
			}
			_, n := TypeResolve(w.pkg, pa.Type, temp, nil, false, w)
			switch n {
			case AstIdentError, AstStarExpr:
				w.F("(%s|undefined)", temp.String())
			default:
				w.F("%s", temp.String())
			}
			w.Pop()
			temp.Reset()
		}
		w.F(")[]")
	}
}
func StructTypeWrite(st *ast.StructType, w *WalkWriter) {
	temp := GetWriter()
	defer FreeWriter(temp)
	if st.Fields.NumFields() != 0 {
		for i, f := range st.Fields.List {
			switch len(f.Names) {
			case 0: //!! embedded
				_, _ = TypeResolve(w.pkg, f.Type, nil, nil, true, w)
			default:
				x := -1
				for _, name := range f.Names {
					if ast.IsExported(name.Name) {
						x++
					}
				}
				if x >= 0 {
					if i > 0 {
						w.F("\n")
						w.F("\t\t")
					} else {
						w.F("\t")
					}
					_, r := TypeResolve(w.pkg, f.Type, temp, nil, false, w)
					x = -1
					switch r {
					case AstIdentError, AstStarExpr:
						for _, name := range f.Names {
							if !ast.IsExported(name.Name) {
								continue
							}
							x++
							if x > 0 {
								w.F(",")
							}
							w.F("%s?", GoFuncToJsFunc(name.Name))
						}
						w.F(":(%s|undefined)", temp.String())
					default:
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
						w.F(":%s", temp.String())
					}
					temp.Reset()
				}
			}
		}
	}
}
func MapTypeWrite(st *ast.MapType, w *WalkWriter) {
	_, _ = TypeResolve(w.pkg, st, nil, nil, w.field, w)
}
