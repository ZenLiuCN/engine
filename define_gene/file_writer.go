package main

import (
	"github.com/ZenLiuCN/engine/define_gene/spec"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"log"
	"strings"
	"unicode"
)

type FileWriter struct {
	pkg       *packages.Package
	debug     bool
	pkgPath   string
	golangPkg bool
	override  bool   //dose override exists file
	out       string //output path
	dts       string //d.ts name
	mod       string //go module file name
	module    string //go module name
	*ModuleWriter
	types   []*Type
	values  []*Value
	defines map[string]*Define
	m       spec.Providers
	current *ast.File

	registry *TypeRegistry
}

func (f *FileWriter) With(pkg *packages.Package, alias string) {
	f.init()
	f.pkg = pkg
	f.golangPkg = !strings.Contains(pkg.PkgPath, ".")
	if alias == "" {
		f.pkgPath = f.pkg.PkgPath
	} else {
		f.pkgPath = alias
	}
	{
		p := strings.ReplaceAll(f.pkgPath, "/", "_")
		f.dts = "golang_" + p + ".d.ts"
		f.mod = "golang_" + p + ".go"
		vs := strings.Split(f.pkgPath, "/")
		b := new(strings.Builder)
		for _, v := range vs {
			if v == "" {
				continue
			}
			b.WriteRune(unicode.ToUpper(rune(v[0])))
			b.WriteString(v[1:])
		}
		f.module = b.String()
	}

}
func (f *FileWriter) Process() {
	f.gather()
	f.computeDefines()
	f.writeTypes()
}
func (f *FileWriter) init() {
	f.pkg = nil
	if f.ModuleWriter == nil {
		f.ModuleWriter = new(ModuleWriter)
	}
	f.ModuleWriter.Init()

	if f.types != nil {
		f.types = f.types[:0]
	}
	if f.values != nil {
		f.values = f.values[:0]
	}
	if f.registry == nil {
		f.registry = &TypeRegistry{}
	}
}

func (f *FileWriter) Flush() {
	f.ModuleWriter.FlushFile(f, f.debug)
}

func (f *FileWriter) initFunctor() {
	f.m = make(spec.Providers, spec.FuncFnMax)
	f.m.WithAllSpec(func(d *ast.FuncDecl) {
		f.types = append(f.types, &Type{
			File:     f.current,
			Type:     FuncDecl,
			FuncDecl: d,
		})
	}, func(d *ast.FuncDecl, r *ast.Field) bool {
		return true
	}, func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.Ident) {
		f.types = append(f.types, &Type{
			File:  f.current,
			Type:  IdentType,
			Decl:  d,
			Spec:  s,
			Ident: t,
		})
	}, func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.StructType) {
		f.types = append(f.types, &Type{
			File:   f.current,
			Type:   StructType,
			Decl:   d,
			Spec:   s,
			Struct: t,
		})
	}, func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.InterfaceType) {
		f.types = append(f.types, &Type{
			File:      f.current,
			Type:      InterfaceType,
			Decl:      d,
			Spec:      s,
			Interface: t,
		})
	}, func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.MapType) {
		f.types = append(f.types, &Type{
			File: f.current,
			Type: MapType,
			Decl: d,
			Spec: s,
			Map:  t,
		})
	}, func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.ArrayType) {
		f.types = append(f.types, &Type{
			File:  f.current,
			Type:  ArrayType,
			Decl:  d,
			Spec:  s,
			Array: t,
		})
	}, func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.FuncType) {
		f.types = append(f.types, &Type{
			File: f.current,
			Type: FuncType,
			Decl: d,
			Spec: s,
			Func: t,
		})
	}, func(d *ast.FuncDecl, r *ast.Field, t *ast.StarExpr) {
		f.types = append(f.types, &Type{
			File:     f.current,
			Type:     StarMethodType,
			FuncDecl: d,
			Recv:     r,
			Star:     t,
		})
	}, func(d *ast.FuncDecl, r *ast.Field, t *ast.Ident) {
		f.types = append(f.types, &Type{
			File:     f.current,
			Type:     IdentMethodType,
			FuncDecl: d,
			Recv:     r,
			Ident:    t,
		})
	}, func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.ChanType) {
		f.types = append(f.types, &Type{
			File:     f.current,
			Type:     ChanType,
			Decl:     d,
			Spec:     s,
			ChanType: t,
		})
	}, func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.SelectorExpr) {
		f.types = append(f.types, &Type{
			File:   f.current,
			Type:   SelectorType,
			Decl:   d,
			Spec:   s,
			Select: t,
		})
	}, func(d *ast.GenDecl, s *ast.TypeSpec, t *ast.StarExpr) {
		f.types = append(f.types, &Type{
			File: f.current,
			Type: StarType,
			Decl: d,
			Spec: s,
			Star: t,
		})
	})
	f.m.WithAllConst(func(d *ast.GenDecl, s *ast.ValueSpec, t *ast.Ident, unused bool) {
		f.values = append(f.values, &Value{
			File:  f.current,
			Name:  s.Names[0].Name,
			Ident: t,
			Spec:  s,
			Decl:  d,
		})
	}, func(d *ast.GenDecl, s *ast.ValueSpec, unused bool) {
		f.values = append(f.values, &Value{
			File: f.current,
			Name: s.Names[0].Name,
			Spec: s,
			Decl: d,
		})
	})
}

func (f *FileWriter) gather() {
	f.initFunctor()
	for _, syntax := range f.pkg.Syntax {
		f.current = syntax
		spec.CaseDeclsFunc(syntax, spec.ConstantDeclMode|spec.TypeDeclMode, f.m.Provide)
	}
}

func (f *FileWriter) computeDefines() {
	f.defines = map[string]*Define{}
	for _, t := range f.types {
		switch t.Type {
		case FuncDecl:
			recv := t.FuncDecl.Name.Name
			if ast.IsExported(recv) {
				if d, ok := f.defines[recv]; ok {
					log.Printf("exists defined function %v for %v\n", d, t)
				} else {
					f.defines[recv] = (&Define{
						File: t.File,
						Type: FuncDecl,
						Spec: t.Spec,
						Decl: t.Decl,
						Prim: t,
					}).AddElt(t)
				}
			}
		case IdentMethodType, StarMethodType:
			recv := f.resolveRecv(t.Recv.Type)
			if ast.IsExported(recv) {
				if d, ok := f.defines[recv]; ok {
					d.AddElt(t)
				} else {
					f.defines[recv] = (&Define{
						File: t.File,
						Type: IdentMethodType,
						Spec: t.Spec,
						Decl: t.Decl,
						Prim: t,
					}).AddElt(t)
				}
			}
		case FuncType:
			recv := t.Spec.Name.Name
			if ast.IsExported(recv) {
				if d, ok := f.defines[recv]; ok {
					log.Printf("exists defined FuncType %v for %v\n", d, t)
				} else {
					f.defines[recv] = (&Define{
						File: t.File,
						Type: FuncType,
						Spec: t.Spec,
						Decl: t.Decl,
						Prim: t,
					}).AddElt(t)
				}
			}
		case StructType:
			recv := t.Spec.Name.Name
			if ast.IsExported(recv) {
				if d, ok := f.defines[recv]; ok {
					if d.Type == IdentMethodType || d.Type == StarMethodType {
						d.AddElt(t)
						d.Prim = t
						d.Type = StructType
					} else {
						log.Printf("exists defined StructType %v for %v\n", d, t)
					}
				} else {
					f.defines[recv] = (&Define{
						File: t.File,
						Type: StructType,
						Spec: t.Spec,
						Decl: t.Decl,
						Prim: t,
					}).AddElt(t)
				}
			}
		case MapType:
			recv := t.Spec.Name.Name
			if ast.IsExported(recv) {
				if d, ok := f.defines[recv]; ok {
					log.Printf("exists defined MapType %v for %v\n", d, t)
				} else {
					f.defines[recv] = (&Define{
						File: t.File,
						Type: MapType,
						Spec: t.Spec,
						Decl: t.Decl,
						Prim: t,
					}).AddElt(t)
				}
			}
		case InterfaceType:
			recv := t.Spec.Name.Name
			if ast.IsExported(recv) {
				if d, ok := f.defines[recv]; ok {
					log.Printf("exists defined InterfaceType %v for %v\n", d, t)
				} else {
					f.defines[recv] = (&Define{
						File: t.File,
						Type: InterfaceType,
						Spec: t.Spec,
						Decl: t.Decl,
						Prim: t,
					}).AddElt(t)
				}
			}
		case ArrayType:
			recv := t.Spec.Name.Name
			if ast.IsExported(recv) {
				if d, ok := f.defines[recv]; ok {
					log.Printf("exists defined ArrayType %v for %v\n", d, t)
				} else {
					f.defines[recv] = (&Define{
						File: t.File,
						Type: ArrayType,
						Spec: t.Spec,
						Decl: t.Decl,
						Prim: t,
					}).AddElt(t)
				}
			}
		case IdentType:
			recv := t.Spec.Name.Name
			if ast.IsExported(recv) {
				if d, ok := f.defines[recv]; ok {
					if d.Prim.Type == IdentMethodType || d.Prim.Type == StarMethodType {
						d.Prim = t
						d.Type = IdentType
						d.Spec = t.Spec
						d.Decl = t.Decl
						d.AddElt(t)
					} else {
						log.Printf("exists defined IdentType %v for %v\n", d, t)
					}
				} else {
					f.defines[recv] = (&Define{
						File: t.File,
						Type: IdentType,
						Spec: t.Spec,
						Decl: t.Decl,
						Prim: t,
					}).AddElt(t)
				}
			}
		case StarType:
			recv := t.Spec.Name.Name
			if ast.IsExported(recv) {
				if d, ok := f.defines[recv]; ok {
					log.Printf("exists defined StarType %v for %v\n", d, t)
				} else {
					f.defines[recv] = (&Define{
						File: t.File,
						Type: StarType,
						Spec: t.Spec,
						Decl: t.Decl,
						Prim: t,
						Val:  nil,
					}).AddElt(t)
				}
			}
		default:
			log.Printf("miss type %#v\n", t)
		}
	}
	for _, t := range f.values {
		if ast.IsExported(t.Name) {
			if t.Ident != nil {
				recv := t.Ident.Name
				if d, ok := f.defines[recv]; ok {
					d.AddVal(t)
				} else {
					log.Printf("new defined ValueSpec %v \n", t)
					f.defines[recv] = (&Define{
						File: t.File,
						Type: ConstValue,
						Decl: t.Decl,
					}).AddVal(t)
				}
			} else {
				if d, ok := f.locateOrCreateValue(t); !ok {
					f.defines[t.Name] = d
				}
			}
		}
	}
}

// locate exists define or put all value into new define
func (f *FileWriter) locateOrCreateValue(t *Value) (*Define, bool) {
	for _, define := range f.defines {
		if define.Decl == t.Decl {
			return define, true
		}
	}
	d := &Define{
		File: t.File,
		Type: ConstValue,
		Decl: t.Decl,
		Val:  nil,
	}
	d.AddVal(t)
	for _, s := range t.Decl.Specs {
		for _, value := range f.values {
			if value.Spec == s {
				d.AddVal(value)
			}
		}
	}
	return d, false
}

func (f *FileWriter) resolveRecv(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return f.resolveRecv(t.X)
	default:
		log.Printf("miss recv %#+v\n", expr)
		return ""
	}
}

func (f *FileWriter) writeTypes() {
	for name, define := range f.defines {
		printf("[DEFINE]  %#+v", define)
		tw := new(TypeWriter)
		tw.TypeRegistry = f.registry
		tw.With(f.pkg, name, define)
		tw.Process()
		f.Append(tw.ModuleWriter)
		tw.Free()
	}
}
