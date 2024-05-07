package internal

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"slices"
	"strings"
)

type TypeInspector[X any] struct {
	WithUnexported bool
	Visitor        TypeVisitor[*TypeInspector[X], X]
	Pkg            *packages.Package
	Face           map[Face]struct{}
	X              X
}
type Face struct {
	Pkg       *types.Package
	Name      string
	Interface *types.Interface
}

func NewTypeInspector[X any](withUnexported bool) *BaseInspector[*TypeInspector[X]] {
	v := new(BaseInspector[*TypeInspector[X]])
	v.Inspector = &TypeInspector[X]{WithUnexported: withUnexported}
	return v
}

var (
	exported = ast.IsExported
)

func (s *TypeInspector[X]) initialize(conf *packages.Config) {
	if s.Visitor == nil {
		s.Visitor = new(FnTypeVisitor[*TypeInspector[X], X])
	}
	conf.Mode |= packages.NeedTypes | packages.NeedTypesInfo
}
func (s *TypeInspector[X]) LoadFaces() {
	if s.Face == nil {
		s.Face = map[Face]struct{}{}

	}
	scope := s.Pkg.Types.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		if !obj.Exported() {
			continue
		}
		if _, ok := obj.(*types.TypeName); ok {
			if f, ok := obj.Type().(*types.Named); ok {
				if f, ok := f.Underlying().(*types.Interface); ok {
					s.Face[Face{
						Pkg:       s.Pkg.Types,
						Name:      name,
						Interface: f,
					}] = struct{}{}
				}
			}
		}
	}
	for _, p2 := range s.Pkg.Types.Imports() {
		scope = p2.Scope()
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			if !obj.Exported() {
				continue
			}
			if _, ok := obj.(*types.TypeName); ok {
				if f, ok := obj.Type().(*types.Named); ok {
					if f, ok := f.Underlying().(*types.Interface); ok {
						s.Face[Face{
							Pkg:       p2,
							Name:      name,
							Interface: f,
						}] = struct{}{}
					}
				}
			}
		}
	}
}
func (s *TypeInspector[X]) Implements(p types.Type) (m []Face) {
	s.LoadFaces()
	px := types.NewPointer(p)
	for face := range s.Face {
		sam := types.Implements(px, face.Interface) || types.Implements(p, face.Interface)
		tracef(0, "cmp %s vs %s %t", face.Name, p.String(), sam)
		if sam {
			m = append(m, face)
		}
	}

	return reduce(m)
}

func reduce(m []Face) (r []Face) {
	var x []int
	for _, face := range m {
		for i, n := 0, face.Interface.NumEmbeddeds(); i < n; i++ {
			e := face.Interface.EmbeddedType(i)
			f := e.Underlying().(*types.Interface)
			for i, f2 := range m {
				if types.IdenticalIgnoreTags(f, f2.Interface) {
					x = append(x, i)
				}
				if len(x) == len(m)-1 {
					break
				}
			}
		}
	}
	for i, face := range m {
		if !slices.Contains(x, i) {
			r = append(r, face)
		}
	}
	return
}

func (s *TypeInspector[X]) inspect(p *packages.Package) {
	s.Pkg = p
	scope := p.Types.Scope()
	for _, name := range scope.Names() {
		if !exported(name) {
			continue
		}
		o := scope.Lookup(name)
		switch e := o.(type) {
		case *types.Const:
			if s.Visitor.VisitConst(s, ENT, name, e, s.X) {
				s.visitType(o.Type(), o, nil, nil)
			}
			s.Visitor.VisitConst(s, EXT, name, e, s.X)
		case *types.Func:
			if s.Visitor.VisitFunc(s, ENT, name, e, s.X) {
				s.visitType(o.Type(), o, nil, nil)
			}
			s.Visitor.VisitFunc(s, EXT, name, e, s.X)
		case *types.Var:
			if s.Visitor.VisitVar(s, ENT, name, e, s.X) {
				s.visitType(o.Type(), o, nil, nil)
			}
			s.Visitor.VisitVar(s, EXT, name, e, s.X)
		case *types.TypeName:
			if s.Visitor.VisitTypeName(s, ENT, name, e, s.X) {
				s.visitType(o.Type(), o, nil, nil)
			}
			s.Visitor.VisitTypeName(s, EXT, name, e, s.X)
		default:
			debugf("%s => %T\n", name, e)
		}

	}
}

func (s *TypeInspector[X]) visitType(t types.Type, o types.Object, mods Mods, seen Types) {
	if i := slices.Index(seen, t); i > -1 && i != 0 {
		return
	}
	switch x := t.(type) {
	case *types.Basic:
		s.visitBasic(x, o, mods, seen)
	case *types.Map:
		s.visitMap(x, o, mods, seen)
	case *types.Array:
		s.visitArray(x, o, mods, seen)
	case *types.Struct:
		s.visitStruct(x, o, mods, seen)
	case *types.Tuple:
		s.visitTuple(x, o, mods, seen)
	case *types.Union:
		s.visitUnion(x, o, mods, seen)
	case *types.Signature:
		s.visitSignature(x, o, mods, seen)
	case *types.TypeParam:
		s.visitTypeParam(x, o, mods, seen)
	case *types.Pointer:
		s.visitPointer(x, o, mods, seen)
	case *types.Slice:
		s.visitSlice(x, o, mods, seen)
	case *types.Interface:
		s.visitInterface(x, o, mods, seen)
	case *types.Chan:
		s.visitChan(x, o, mods, seen)
	case *types.Named:
		s.visitNamed(x, o, mods, seen)
	default:
		debugf("undefined %T", x)
	}
}

func (s *TypeInspector[X]) visitSignature(x *types.Signature, o types.Object, mods Mods, seen Types) {
	if len(mods) == 0 {
		if x.Recv() == nil {
			mods = append(mods, ModFunction)
		} else {
			mods = append(mods, ModMethod)
		}
	}
	if s.Visitor.VisitTypeSignature(s, ENT, o, x, mods, seen, s.X) {
		s.visitType(x.Params(), o, append(mods, ModParam), append(seen, x))
		s.visitType(x.Results(), o, append(mods, ModResult), append(seen, x))
	}
	s.Visitor.VisitTypeSignature(s, EXT, o, x, mods, seen, s.X)
}

func (s *TypeInspector[X]) isExported(txt string) bool {
	if !s.WithUnexported {
		n := strings.ReplaceAll(txt, ".", "/")
		i := strings.LastIndex(n, "/")
		return exported(n[i+1:])
	}
	return true
}

func (s *TypeInspector[X]) visitNamed(x *types.Named, o types.Object, mods Mods, seen Types) {
	if s.Visitor.VisitTypeNamed(s, ENT, o, x, mods, seen, s.X) {
		seen2 := append(seen, x)
		s.visitType(x.Underlying(), o, append(mods, ModeNamedElt), seen2)
		for i1, n1 := 0, x.NumMethods(); i1 < n1; i1++ {
			m1 := x.Method(i1)
			if s.isExported(m1.Name()) {
				if s.Visitor.VisitTypeFunc(s, ENT, o, m1, append(mods, ModeNamedElt), seen2, s.X) {
					s.visitType(m1.Type(), o, append(mods, ModeNamedElt, ModMethod), seen2)
				}
				s.Visitor.VisitTypeFunc(s, EXT, o, m1, append(mods, ModeNamedElt), seen2, s.X)
			}
		}
	}
	s.Visitor.VisitTypeNamed(s, EXT, o, x, mods, seen, s.X)
}

func (s *TypeInspector[X]) visitChan(x *types.Chan, o types.Object, mods Mods, seen Types) {
	if s.Visitor.VisitTypeChan(s, ENT, o, x, mods, seen, s.X) {
		s.visitType(x.Elem(), o, append(mods, ModChanElt), append(seen, x))
	}
	s.Visitor.VisitTypeChan(s, EXT, o, x, mods, seen, s.X)
}

func (s *TypeInspector[X]) visitInterface(x *types.Interface, o types.Object, mods Mods, seen Types) {
	if s.Visitor.VisitTypeInterface(s, ENT, o, x, mods, seen, s.X) {
		s.visitInterfaceEmbedded(x, o, mods, append(seen, x))
		s.visitInterfaceMethods(x, o, mods, append(seen, x))
	}
	s.Visitor.VisitTypeInterface(s, EXT, o, x, mods, seen, s.X)
}
func (s *TypeInspector[X]) visitInterfaceEmbedded(x *types.Interface, o types.Object, mods Mods, seen Types) {
	for i, n := 0, x.NumEmbeddeds(); i < n; i++ {
		m := x.EmbeddedType(i).(*types.Named)
		if s.isExported(m.String()) {
			s.visitNamed(m, o, append(mods, ModEmbedded), seen)
		}
	}
}

func (s *TypeInspector[X]) visitInterfaceMethods(x *types.Interface, o types.Object, mods Mods, seen Types) {
	for i, n := 0, x.NumExplicitMethods(); i < n; i++ {
		m := x.ExplicitMethod(i)
		if s.isExported(m.Name()) {
			if s.Visitor.VisitTypeFunc(s, ENT, o, m, mods, seen, s.X) {
				s.visitType(m.Type(), o, append(mods, ModMethod), seen)
			}
			s.Visitor.VisitTypeFunc(s, EXT, o, m, mods, seen, s.X)
		}
	}
}

func (s *TypeInspector[X]) visitSlice(x *types.Slice, o types.Object, mods Mods, seen Types) {
	if s.Visitor.VisitTypeSlice(s, ENT, o, x, mods, seen, s.X) {
		if seen.LNth(1) == x {
			s.visitType(x.Elem(), o, append(mods, ModSliceElt), seen)
		} else {
			s.visitType(x.Elem(), o, append(mods, ModSliceElt), append(seen, x))
		}

	}
	s.Visitor.VisitTypeSlice(s, EXT, o, x, mods, seen, s.X)
}

func (s *TypeInspector[X]) visitPointer(x *types.Pointer, o types.Object, mods Mods, seen Types) {
	if s.Visitor.VisitTypePointer(s, ENT, o, x, mods, seen, s.X) {
		s.visitType(x.Elem(), o, append(mods, ModPointerElt), append(seen, x))
	}
	s.Visitor.VisitTypePointer(s, EXT, o, x, mods, seen, s.X)
}

func (s *TypeInspector[X]) visitTypeParam(x *types.TypeParam, o types.Object, mods Mods, seen Types) {
	if s.Visitor.VisitTypeParam(s, ENT, o, x, mods, append(seen, x), s.X) {
		s.visitType(x.Underlying(), o, append(mods, ModTypeParam), append(seen, x))
	}
	s.Visitor.VisitTypeParam(s, EXT, o, x, mods, seen, s.X)
}

func (s *TypeInspector[X]) visitUnion(x *types.Union, o types.Object, mods Mods, seen Types) {
	if s.Visitor.VisitTypeUnion(s, ENT, o, x, mods, seen, s.X) {
		n := x.Len()
		for i := 0; i < n; i++ {
			s.visitType(x.Term(i).Type(), o, mods, append(seen, x))
		}
	}
	s.Visitor.VisitTypeUnion(s, EXT, o, x, mods, seen, s.X)
}

func (s *TypeInspector[X]) visitTuple(x *types.Tuple, o types.Object, mods Mods, seen Types) {
	if s.Visitor.VisitTypeTuple(s, ENT, o, x, mods, seen, s.X) {
		n := x.Len()
		mods2 := append(mods, ModTuple)
		seen2 := append(seen, x)
		for i := 0; i < n; i++ {
			s.Visitor.VisitTypeVar(s, ENT, o, x.At(i), mods2, seen2, s.X)
			s.visitType(x.At(i).Type(), o, mods2, seen2)
			s.Visitor.VisitTypeVar(s, EXT, o, x.At(i), mods2, seen2, s.X)
		}
	}
	s.Visitor.VisitTypeTuple(s, EXT, o, x, mods, seen, s.X)
}

func (s *TypeInspector[X]) visitStruct(x *types.Struct, o types.Object, mods Mods, seen Types) {
	if s.Visitor.VisitTypeStruct(s, ENT, o, x, mods, seen, s.X) {
		seen2 := append(seen, x)
		for i := 0; i < x.NumFields(); i++ {
			if s.isExported(x.Field(i).Name()) {
				if !s.Visitor.VisitTypeVar(s, ENT, o, x.Field(i), append(mods, ModField), seen2, s.X) {
					continue
				}
				s.visitType(x.Field(i).Type(), x.Field(i), append(mods, ModField), seen2)
				if !s.Visitor.VisitTypeVar(s, EXT, o, x.Field(i), append(mods, ModField), seen2, s.X) {
					continue
				}
			}
		}
	}
	s.Visitor.VisitTypeStruct(s, EXT, o, x, mods, seen, s.X)
}

func (s *TypeInspector[X]) visitArray(x *types.Array, o types.Object, mods Mods, seen Types) {
	if s.Visitor.VisitTypeArray(s, ENT, o, x, mods, seen, s.X) {
		s.visitType(x.Elem(), o, append(mods, ModArrayElt), append(seen, x))
	}
	s.Visitor.VisitTypeArray(s, EXT, o, x, mods, seen, s.X)
}

func (s *TypeInspector[X]) visitMap(x *types.Map, o types.Object, mods Mods, seen Types) {
	if s.Visitor.VisitTypeMap(s, ENT, o, x, mods, seen, s.X) {
		s.visitType(x.Key(), o, append(mods, ModMapKey), append(seen, x))
		s.visitType(x.Elem(), o, append(mods, ModMapValue), append(seen, x))
	}
	s.Visitor.VisitTypeMap(s, EXT, o, x, mods, seen, s.X)
}

func (s *TypeInspector[X]) visitBasic(x *types.Basic, o types.Object, mods Mods, seen Types) {
	if s.Visitor.VisitTypeBasic(s, ENT, o, x, mods, seen, s.X) {

	}
	s.Visitor.VisitTypeBasic(s, EXT, o, x, mods, seen, s.X)
}
