package main

//https://cs.opensource.google/go/x/tools/+/master:cmd/stringer/stringer.go;drc=daf94608b5e2caf763ba634b84e7a5ba7970e155;l=382
import (
	"bytes"
	"errors"
	"fmt"
	. "github.com/ZenLiuCN/engine/define_gene/spec"
	"github.com/ZenLiuCN/fn"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

type Generator struct {
	dir   string
	tags  []string
	files []string
	types []string
	out   string
	log   func(format string, a ...any)
	over  bool
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

type Value struct {
	name      string
	identType *ast.Ident
	spec      *ast.ValueSpec
	decl      *ast.GenDecl
}
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
	Select     *ast.SelectorExpr
	Star       *ast.StarExpr
}

const glob = "global"

// region Declare
type ConstWalker struct {
	found func(*Value)
}

func (c *ConstWalker) IdentConstSpec(d *ast.GenDecl, s *ast.ValueSpec, t *ast.Ident) {
	for _, name := range s.Names {
		if ast.IsExported(name.Name) {
			c.found(&Value{
				name:      name.Name,
				identType: t,
				decl:      d,
				spec:      s,
			})
		}
	}
}

func (c *ConstWalker) NilConstSpec(d *ast.GenDecl, s *ast.ValueSpec) {
	for _, name := range s.Names {
		if ast.IsExported(name.Name) {
			c.found(&Value{
				name: name.Name,
				decl: d,
				spec: s,
			})
		}
	}
}

type WalkDecl struct {
	found func(string, *Type)
}

func (w WalkDecl) StarExprSpec(s *ast.TypeSpec, t *ast.StarExpr) {
	w.set(s.Name.Name, &Type{
		Type: AliasType,
		Spec: s,
		Star: t,
	})
}
func (w WalkDecl) SelectorTypeSpec(s *ast.TypeSpec, t *ast.SelectorExpr) {
	w.set(s.Name.Name, &Type{

		Type:   AliasType,
		Spec:   s,
		Select: t,
	})
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
	if pkg.Name == "" {
		return
	}
	ts := map[string][]*Type{}
	var cs []*Value
	for _, file := range pkg.Syntax {
		CaseTypeDecl(file, &ExportedTypeDeclCases{&WalkDecl{func(s string, t *Type) {
			t.File = file
			v := ts[s]
			v = append(v, t)
			ts[s] = v
		}}})
		CaseConstDecl(file, &ExportedConstDeclCases{
			&ConstWalker{func(value *Value) {
				cs = append(cs, value)
			}},
		})
	}
	mw := NewWriter()
	dts := "go_" + pkg.Name + ".d.ts"
	mn := "go" + string(unicode.ToUpper(rune(pkg.Name[0]))) + pkg.Name[1:]
	mw.F(`package golang
import ( 
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"%[1]s"
)
var(
	//go:embed %[2]s
	%[3]sDefine []byte
	%[3]sDeclared=map[string]any{

`, pkg.PkgPath, dts, mn)
	//!! write d.ts
	fw := NewWriter() //file
	fw.F("declare module 'golang/%s' {\n", pkg.Name)
	iw := NewWriter()
	dw := NewWriter()  // declared
	exw := NewWriter() // extend
	ew := NewWriter()  // elements of declared
	tw := NewWriter()  // element
	lookupValues := func(name string, t *Type) []*Value {
		var values []*Value
		prefix := name
		for _, c := range cs {
			if strings.HasPrefix(c.name, prefix) {
				values = append(values, c)
			} else if c.identType != nil && c.identType.Name == name {
				for i := 1; i < len(c.name); i++ {
					if unicode.IsUpper(rune(c.name[i])) {
						prefix = c.name[:i]
						break
					}
				}
				for _, c := range cs {
					if strings.HasPrefix(c.name, prefix) {
						values = append(values, c)
					}
				}
				if len(values) == 0 {
					decl := c.decl
					for _, c := range cs {
						if c.decl == decl {
							values = append(values, c)
						}
					}
				}
				break
			}
		}
		return values
	}
	for name, types := range ts {
		slices.SortFunc(types, func(a, b *Type) int {
			return int(a.Type - b.Type)
		})
		closure := name != glob
		closed := false
		if closure {
			switch types[0].Type {
			case MapType, StructType, InterfaceType:
				dw.F("\t /*%s*/ export interface  %s", types[0].Type, name)
				closed = true
			default:
				if len(types) > 1 && !(types[1].Type == MethodDecl && types[1].MethodDecl.Name.Name == "String") { //!! exclude alias with only Stringer
					dw.F("\t export interface %s", name)
					closed = true
				} else if len(types) == 2 && (types[1].Type == MethodDecl && types[1].MethodDecl.Name.Name == "String") {
					types = types[:1]
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
				mw.F("\"%s\":%s.%s ,\n", GoFuncToJsFunc(d.Name.Name), pkg.Name, d.Name.Name)
			case MethodDecl:
				if closure {
					tw.F("\t")
				}
				d := t.MethodDecl
				if types[0].Type == MapType {
					tw.F("\t\t // @ts-ignore\n")
				}
				tw.F("\t\t %s ", GoFuncToJsFunc(d.Name.Name))
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
				if dw.Len() == 0 {
					values := lookupValues(name, t)
					if len(values) == 0 {
						dw.F("/*miss values for %v */\n", t)
						continue
					}
					decl := values[0]
					n := 0
					var inc func()
					if decl.identType == nil {
						dw.F("\texport type %s =", name)
						for i, value := range values {
							if i > 0 {
								dw.F("|")
							}
							dw.F("/*%s*/%d", value.name, value.spec.Names[0].Obj.Data)
						}
						dw.F("\n")
					} else {
						if expr, ok := decl.spec.Values[0].(*ast.BinaryExpr); ok {
							var inv int
							if lit, ok := expr.Y.(*ast.BasicLit); ok {
								inv = fn.Panic1(strconv.Atoi(lit.Value))
							} else {
								inv = fn.Panic1(strconv.Atoi(expr.X.(*ast.BasicLit).Value))
							}
							switch expr.Op {
							case token.ADD:
								inc = func() {
									n += inv
								}
							case token.SHL:
								inc = func() {
									n = inv
								}
							}
						} else if _, ok := decl.spec.Values[0].(*ast.Ident); ok {
							n = 0
							inc = func() {
								n++
							}
						}
						if inc != nil {
							dw.F("\texport type %s =", name)
							for i, value := range values {
								if i > 0 {
									inc()
									dw.F("|")
								}
								dw.F("/*%s*/%d", value.name, n)
							}

						} else {
							for _, value := range values {
								dw.F("\texport const %s :%s\n", value.name, name)
								mw.F("\"%s\":%s.%s,\n", value.name, pkg.Name, value.name)
							}
							dw.F("export interface %s \n", name)
							closed = true
						}
					}
					break
				} else {
					values := lookupValues(name, t)
					if len(values) > 0 {
						for _, value := range values {
							iw.F("\t\texport const %s:%s \n", value.name, name)
						}
					} else if strings.HasSuffix(name, "Error") {

					} else {
						fmt.Printf("miss values for alias: %s \n", name)
					}
				}
			case InterfaceType:
				switch t.Interface.Methods.NumFields() {
				case 0:
				default:
					for _, field := range t.Interface.Methods.List {
						if len(field.Names) > 0 {
							switch len(field.Names) {
							case 1:
								tw.F("%s", GoFuncToJsFunc(field.Names[0].Name))
							default:
								for i, ident := range field.Names {
									if i > 0 {
										tw.F(",")
									}
									tw.F("%s", GoFuncToJsFunc(ident.Name))
								}
							}
							switch t := field.Type.(type) {
							case *ast.SelectorExpr:
								TypeResolve(pkg, field.Type, tw, exw, false, nil)
							case *ast.FuncType:
								FuncDeclWrite(t, NewWalkWriter(pkg, false, exw, tw, nil))
							default:
								panic(fmt.Errorf("miss %#+v", t))
							}
							tw.F("\n")
						} else {
							tw.F("/* TODO miss %+v */\n", field)
							//TypeResolve(pkg, field.Type, tw, exw, true, nil)
						}
					}
				}
			case ArrayType:
				ArrayTypeWrite(t.Array, NewWalkWriter(pkg, true, exw, tw, nil))
			case AliasType:
				switch x := t.Spec.Type.(type) {
				case *ast.Ident:
					dw.F("\texport type %s = %s", name, x.Name)
				case *ast.SelectorExpr:
					dw.Import(x.X.(*ast.Ident).Name)
					dw.F("\texport type %s = %s.%s", name, x.X.(*ast.Ident).Name, x.Sel.Name)
				default:
					fmt.Printf("Alias %s \n", t.Spec.Name)
				}

			default:
				println(fmt.Sprintf("miss %#+v", t))
			}
			ew.Append(tw)
			tw.Reset()
		}
		if exw.Len() > 0 {
			dw.F(" extends %s", exw.String())
			exw.Reset()
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
		fw.F("\t//@ts-ignore\n\timport * as %[1]s from 'golang/%[1]s'\n", k)
	}
	fw.Append(iw)
	fw.F("\n}")
	pDTS := filepath.Join(g.out, dts)
	if !g.over && Exists(pDTS) {
		panic(fmt.Errorf("%s exists for %s", pDTS, pkg.PkgPath))
	}
	_ = os.WriteFile(pDTS, fw.Bytes(), os.ModePerm)
	mw.F(`
	}
)
func init() {
	engine.RegisterModule(%[1]sModule{})
}
type %[1]sModule struct{}
func (S %[1]sModule) Identity() string {
	return "golang/%[2]s"
}
func (S %[1]sModule) TypeDefine() []byte {
	return %[1]sDefine
}
func (S %[1]sModule) Exports() map[string]any {
	return %[1]sDeclared
}
`, mn, pkg.Name)
	pDG := filepath.Join(g.out, strings.Replace(dts, ".d.ts", ".go", 1))
	if !g.over && Exists(pDG) {
		panic(fmt.Errorf("%s exists for %s", pDG, pkg.PkgPath))
	}
	_ = os.WriteFile(pDG, mw.Bytes(), os.ModePerm)
}

func GoFuncToJsFunc(n string) string {
	if n == "New" {
		return n
	}
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
	case "byte", "uint8":
		if array {
			return "Uint8Array"
		} else {
			return fmt.Sprintf("/*%s*/number", s)
		}
	case "uint16", "uint32":
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
	case "int", "int64", "float64", "float32":
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
	case "uintptr":
		if array {
			return "number/*uintptr*/[]"
		}
		return "number/*uintptr*/"
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
	d, ok := w.LastN(2)
	if !ok {
		return false
	}
	for _, layer := range d {
		if layer.Node == AstArrayType {
			return true
		}
	}
	return false
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
	if w.field {
		switch x := t.Elt.(type) {
		case *ast.Ident:
			w.decl.F("%s", GoIdentToTs(x.Name, true))
		case *ast.ArrayType:
			temp := GetWriter()
			old := w.decl
			w.decl = temp
			CaseType(w.pkg, t.Elt, w.u)
			w.decl.F("Array<%s>", temp.String())
			old.MergeImports(temp.Imports())
			w.decl = old
			FreeWriter(temp)
		default:
			fmt.Printf("miss array type of field: %#+v\n", t)
		}
	}
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
	defer FreeWriter(temp)
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
			w.MergeImports(temp.Imports())
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
		w.MergeImports(temp.Imports())
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
			w.MergeImports(temp.Imports())
			temp.Reset()
		}
		w.F(")[]")
	}
	w.F(")")
}
func FuncDeclWrite(d *ast.FuncType, w *WalkWriter) {
	temp := GetWriter()
	defer FreeWriter(temp)
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
			w.MergeImports(temp.Imports())
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
		w.MergeImports(temp.Imports())
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
			w.MergeImports(temp.Imports())
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
					w.MergeImports(temp.Imports())
					temp.Reset()
				}
			}
		}

	}
}
func MapTypeWrite(st *ast.MapType, w *WalkWriter) {
	_, _ = TypeResolve(w.pkg, st, nil, nil, w.field, w)
}
func ArrayTypeWrite(st *ast.ArrayType, w *WalkWriter) {
	_, _ = TypeResolve(w.pkg, st, nil, nil, w.field, w)
}
func Exists(p string) bool {
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil {
		panic(err)
	}
	return true
}
