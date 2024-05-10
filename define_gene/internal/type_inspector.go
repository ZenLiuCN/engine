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
	PopEach        bool //pop each Type when done
	TypePath            //the path from root to leaf, which is cutoff branches when current node is [ KMField ], [ KMParam ], [ KMResult ]; not include current visit node
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

func TypeObjectVisit[X any](pkg *packages.Package, name string, o types.Object, visitor TypeVisitor[*TypeInspector[X], X], withUnexported bool) {
	s := &TypeInspector[X]{WithUnexported: withUnexported}
	s.Pkg = pkg
	s.Visitor = visitor
	switch e := o.(type) {
	case *types.Const:
		if s.Visitor.VisitConst(s, ENT, name, e, s.X) {
			s.visitType(o.Type(), o, nil)
		}
		s.Visitor.VisitConst(s, EXT, name, e, s.X)
	case *types.Func:
		if s.Visitor.VisitFunc(s, ENT, name, e, s.X) {
			s.visitType(o.Type(), o, nil)
		}
		s.Visitor.VisitFunc(s, EXT, name, e, s.X)
	case *types.Var:
		if s.Visitor.VisitVar(s, ENT, name, e, s.X) {
			s.visitType(o.Type(), o, nil)
		}
		s.Visitor.VisitVar(s, EXT, name, e, s.X)
	case *types.TypeName:
		if s.Visitor.VisitTypeName(s, ENT, name, e, s.X) {
			s.visitType(o.Type(), o, nil)
		}
		s.Visitor.VisitTypeName(s, EXT, name, e, s.X)
	default:
		debugf("%s => %T\n", name, e)
	}
}
func TypeVisit[X any](pkg *packages.Package, o types.Object, p types.Type, visitor TypeVisitor[*TypeInspector[X], X], withUnexported bool) {
	v := &TypeInspector[X]{WithUnexported: withUnexported}
	v.Pkg = pkg
	v.Visitor = visitor
	v.visitType(p, o, nil)
}
func (s *TypeInspector[X]) initialize(conf *packages.Config) {
	if s.Visitor == nil {
		s.Visitor = new(FnTypeVisitor[*TypeInspector[X], X])
	}
	if conf != nil {
		conf.Mode |= packages.NeedTypes | packages.NeedTypesInfo
	}

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
			s.OnEnter(KTConst)
			if s.Visitor.VisitConst(s, ENT, name, e, s.X) {
				s.visitType(o.Type(), o, nil)
			}
			s.Visitor.VisitConst(s, EXT, name, e, s.X)
			s.OnExit(KTConst)
		case *types.Func:
			s.OnEnter(KTFunc)
			if s.Visitor.VisitFunc(s, ENT, name, e, s.X) {
				s.visitType(o.Type(), o, nil)
			}
			s.Visitor.VisitFunc(s, EXT, name, e, s.X)
			s.OnExit(KTFunc)
		case *types.Var:
			s.OnEnter(KTVar)
			if s.Visitor.VisitVar(s, ENT, name, e, s.X) {
				s.visitType(o.Type(), o, nil)
			}
			s.Visitor.VisitVar(s, EXT, name, e, s.X)
			s.OnExit(KTVar)
		case *types.TypeName:
			s.OnEnter(KTType)
			if s.Visitor.VisitTypeName(s, ENT, name, e, s.X) {
				s.visitType(o.Type(), o, nil)
			}
			s.Visitor.VisitTypeName(s, EXT, name, e, s.X)
			s.OnExit(KTType)
		default:
			debugf("%s => %T\n", name, e)
		}

	}
}

func (s *TypeInspector[X]) visitType(t types.Type, o types.Object, seen Types) {
	if i := slices.Index(seen, t); i > -1 && i != 0 {
		return
	}
	switch x := t.(type) {
	case *types.Basic:
		s.visitBasic(x, o, seen)
	case *types.Map:
		s.visitMap(x, o, seen)
	case *types.Array:
		s.visitArray(x, o, seen)
	case *types.Struct:
		s.visitStruct(x, o, seen)
	case *types.Tuple:
		s.visitTuple(x, o, seen)
	case *types.Union:
		s.visitUnion(x, o, seen)
	case *types.Signature:
		s.visitSignature(x, o, seen)
	case *types.TypeParam:
		s.visitTypeParam(x, o, seen)
	case *types.Pointer:
		s.visitPointer(x, o, seen)
	case *types.Slice:
		s.visitSlice(x, o, seen)
	case *types.Interface:
		s.visitInterface(x, o, seen)
	case *types.Chan:
		s.visitChan(x, o, seen)
	case *types.Named:
		s.visitNamed(x, o, seen)
	default:
		debugf("undefined %T", x)
	}
}

func (s *TypeInspector[X]) visitSignature(x *types.Signature, o types.Object, seen Types) {
	s.OnEnter(KSignature)
	if s.Visitor.VisitTypeSignature(s, ENT, o, x, seen, s.X) {
		s.OnEnter(KMParam)
		s.visitType(x.Params(), o, append(seen, x))
		s.OnExit(KMParam)
		s.OnEnter(KMResult)
		s.visitType(x.Results(), o, append(seen, x))
		s.OnExit(KMResult)
	}
	s.Visitor.VisitTypeSignature(s, EXT, o, x, seen, s.X)
	s.OnExit(KSignature)
}

func (s *TypeInspector[X]) isExported(txt string) bool {
	if !s.WithUnexported {
		n := strings.ReplaceAll(txt, ".", "/")
		i := strings.LastIndex(n, "/")
		return exported(n[i+1:])
	}
	return true
}

func (s *TypeInspector[X]) visitNamed(x *types.Named, o types.Object, seen Types) {
	s.OnEnter(KNamed)
	if s.Visitor.VisitTypeNamed(s, ENT, o, x, seen, s.X) {
		seen2 := append(seen, x)
		s.visitType(x.Underlying(), o, seen2)
		for i1, n1 := 0, x.NumMethods(); i1 < n1; i1++ {
			m1 := x.Method(i1)
			if s.isExported(m1.Name()) {
				s.OnEnter(KMMethod)
				if s.Visitor.VisitTypeFunc(s, ENT, o, m1, seen2, s.X) {
					s.visitType(m1.Type(), o, seen2)
				}
				s.Visitor.VisitTypeFunc(s, EXT, o, m1, seen2, s.X)
				s.OnExit(KMMethod)
			}
		}
	}
	s.Visitor.VisitTypeNamed(s, EXT, o, x, seen, s.X)
	s.OnExit(KNamed)
}

func (s *TypeInspector[X]) visitChan(x *types.Chan, o types.Object, seen Types) {
	s.OnEnter(KChan)
	if s.Visitor.VisitTypeChan(s, ENT, o, x, seen, s.X) {
		s.visitType(x.Elem(), o, append(seen, x))
	}
	s.Visitor.VisitTypeChan(s, EXT, o, x, seen, s.X)
	s.OnExit(KChan)
}

func (s *TypeInspector[X]) visitInterface(x *types.Interface, o types.Object, seen Types) {
	s.OnEnter(KInterface)
	if s.Visitor.VisitTypeInterface(s, ENT, o, x, seen, s.X) {
		s.visitInterfaceEmbedded(x, o, append(seen, x))
		s.visitInterfaceMethods(x, o, append(seen, x))
	}
	s.Visitor.VisitTypeInterface(s, EXT, o, x, seen, s.X)
	s.OnExit(KInterface)
}
func (s *TypeInspector[X]) visitInterfaceEmbedded(x *types.Interface, o types.Object, seen Types) {
	for i, n := 0, x.NumEmbeddeds(); i < n; i++ {
		switch m := x.EmbeddedType(i).(type) {
		case *types.Named:
			if s.isExported(m.String()) {
				s.OnEnter(KMEmbedded)
				s.visitNamed(m, o, seen)
				s.OnExit(KMEmbedded)
			}
		case *types.Union:
			s.OnEnter(KMEmbedded)
			s.visitUnion(m, o, seen)
			s.OnExit(KMEmbedded)
		}

	}
}

func (s *TypeInspector[X]) visitInterfaceMethods(x *types.Interface, o types.Object, seen Types) {
	for i, n := 0, x.NumExplicitMethods(); i < n; i++ {
		m := x.ExplicitMethod(i)
		if s.isExported(m.Name()) {
			s.OnEnter(KMMethod)
			if s.Visitor.VisitTypeFunc(s, ENT, o, m, seen, s.X) {
				s.visitType(m.Type(), o, seen)
			}
			s.Visitor.VisitTypeFunc(s, EXT, o, m, seen, s.X)
			s.OnExit(KMMethod)
		}
	}
}

func (s *TypeInspector[X]) visitSlice(x *types.Slice, o types.Object, seen Types) {
	s.OnEnter(KSlice)
	if s.Visitor.VisitTypeSlice(s, ENT, o, x, seen, s.X) {
		if seen.LNth(1) == x {
			s.visitType(x.Elem(), o, seen)
		} else {
			s.visitType(x.Elem(), o, append(seen, x))
		}

	}
	s.Visitor.VisitTypeSlice(s, EXT, o, x, seen, s.X)
	s.OnExit(KSlice)
}

func (s *TypeInspector[X]) visitPointer(x *types.Pointer, o types.Object, seen Types) {
	s.OnEnter(KPointer)
	if s.Visitor.VisitTypePointer(s, ENT, o, x, seen, s.X) {
		s.visitType(x.Elem(), o, append(seen, x))
	}
	s.Visitor.VisitTypePointer(s, EXT, o, x, seen, s.X)
	s.OnExit(KPointer)
}

func (s *TypeInspector[X]) visitTypeParam(x *types.TypeParam, o types.Object, seen Types) {
	s.OnEnter(KTypeParam)
	if s.Visitor.VisitTypeParam(s, ENT, o, x, append(seen, x), s.X) {
		s.visitType(x.Underlying(), o, append(seen, x))
	}
	s.Visitor.VisitTypeParam(s, EXT, o, x, seen, s.X)
	s.OnExit(KTypeParam)
}

func (s *TypeInspector[X]) visitUnion(x *types.Union, o types.Object, seen Types) {
	s.OnEnter(KUnion)
	if s.Visitor.VisitTypeUnion(s, ENT, o, x, seen, s.X) {
		n := x.Len()
		for i := 0; i < n; i++ {
			s.OnEnter(KTerm)
			s.visitType(x.Term(i).Type(), o, append(seen, x))
			s.OnExit(KTerm)
		}
	}
	s.Visitor.VisitTypeUnion(s, EXT, o, x, seen, s.X)
	s.OnExit(KUnion)
}

func (s *TypeInspector[X]) visitTuple(x *types.Tuple, o types.Object, seen Types) {
	s.OnEnter(KTuple)
	if s.Visitor.VisitTypeTuple(s, ENT, o, x, seen, s.X) {
		n := x.Len()
		seen2 := append(seen, x)
		for i := 0; i < n; i++ {
			s.OnEnter(KVar)
			if s.Visitor.VisitTypeVar(s, ENT, o, x.At(i), seen2, s.X) {
				s.visitType(x.At(i).Type(), o, seen2)
			}
			s.Visitor.VisitTypeVar(s, EXT, o, x.At(i), seen2, s.X)
			s.OnExit(KVar)
		}
	}
	s.Visitor.VisitTypeTuple(s, EXT, o, x, seen, s.X)
	s.OnExit(KTuple)
}

func (s *TypeInspector[X]) visitStruct(x *types.Struct, o types.Object, seen Types) {
	s.OnEnter(KStruct)
	if s.Visitor.VisitTypeStruct(s, ENT, o, x, seen, s.X) {
		seen2 := append(seen, x)
		for i := 0; i < x.NumFields(); i++ {
			if s.isExported(x.Field(i).Name()) {
				s.OnEnter(KMField)
				if s.Visitor.VisitTypeVar(s, ENT, o, x.Field(i), seen2, s.X) {
					s.visitType(x.Field(i).Type(), x.Field(i), seen2)
				}
				s.Visitor.VisitTypeVar(s, EXT, o, x.Field(i), seen2, s.X)
				s.OnExit(KMField)
			}
		}
	}
	s.Visitor.VisitTypeStruct(s, EXT, o, x, seen, s.X)
	s.OnExit(KStruct)
}

func (s *TypeInspector[X]) visitArray(x *types.Array, o types.Object, seen Types) {
	s.OnEnter(KArray)
	if s.Visitor.VisitTypeArray(s, ENT, o, x, seen, s.X) {
		s.visitType(x.Elem(), o, append(seen, x))
	}
	s.Visitor.VisitTypeArray(s, EXT, o, x, seen, s.X)
	s.OnExit(KArray)
}

func (s *TypeInspector[X]) visitMap(x *types.Map, o types.Object, seen Types) {
	s.OnEnter(KMap)
	if s.Visitor.VisitTypeMap(s, ENT, o, x, seen, s.X) {
		s.OnEnter(KMMapKey)
		s.visitType(x.Key(), o, append(seen, x))
		s.OnExit(KMMapKey)
		s.OnEnter(KMMapValue)
		s.visitType(x.Elem(), o, append(seen, x))
		s.OnExit(KMMapValue)
	}
	s.Visitor.VisitTypeMap(s, EXT, o, x, seen, s.X)
	s.OnExit(KMap)
}

func (s *TypeInspector[X]) visitBasic(x *types.Basic, o types.Object, seen Types) {
	s.OnEnter(KBasic)
	if s.Visitor.VisitTypeBasic(s, ENT, o, x, seen, s.X) {

	}
	s.Visitor.VisitTypeBasic(s, EXT, o, x, seen, s.X)
	s.OnExit(KBasic)
}

func (s *TypeInspector[X]) OnEnter(kind TypeKind) {
	s.TypePath.PushSitu(kind)
}
func (s *TypeInspector[X]) OnExit(kind TypeKind) {
	switch kind {
	case KMField, KMMethod, KMEmbedded, KMParam, KMResult:
		s.TypePath = s.TypePath[:s.TypePath.LastIndex(kind)]
	default:
		if s.PopEach {
			s.TypePath.PopSitu()
		}
	}
}
