package engine

import (
	"reflect"
)

func init() {
	//TODO common types
	TypeRegistry = map[TypeId]TypeInfo{}
	TypeAlias = map[TypeId]TypeId{}
	RegisterType[[]rune]()
	RegisterTypeAlias(rune(0), int32(0))
}

type TypeId [1]reflect.Type

func (t TypeId) typo() reflect.Type {
	return t[0]
}
func (t TypeId) Valid() bool {
	return t[0] != nil && t[0].Kind() != reflect.Invalid
}
func (t TypeId) Kind() reflect.Kind {
	if t[0] == nil {
		return reflect.Invalid
	}
	return t[0].Kind()
}
func (t TypeId) String() string {
	return t.Identity()
}
func (t TypeId) Identity() string {
	x := t[0]
	if x == nil {
		return "undefined"
	}
	return x.String()
}
func ElementOf(x TypeId) TypeId {
	if x[0] == nil {
		return x
	}
	t := x.typo()
	switch t.Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan, reflect.Map, reflect.Pointer:
		return TypeId{t.Elem()}
	default:
		return TypeId{}
	}
}
func SliceOf(x TypeId) (v TypeId) {
	if x.Valid() {
		t := x.typo()
		v = TypeId{reflect.SliceOf(t)}
	}
	return
}
func MapOf(k, v TypeId) (r TypeId) {
	if k.Valid() && v.Valid() {
		v = TypeId{reflect.MapOf(k.typo(), v.typo())}
	}
	return
}
func TypeOf(v any) (x TypeId) {
	t := reflect.TypeOf(v)
	//if t == nil || t.PkgPath() == "github.com/dop251/goja" {
	if t == nil {
		return
	}
	x = TypeId{t}
	return
}
func TypeUsage(t TypeId) TypeUse {
	v, ok := TypeRegistry[t]
	if ok {
		return &v
	}
	if a, ok := TypeAlias[t]; ok {
		v, ok = TypeRegistry[a]
		if ok {
			return &v
		}
	}
	i := ReflectTypeInfo{id: t}
	return &i
}

type TypeUse interface {
	Id() TypeId
	Slice() any
	Instance() any
}
type TypeInfo struct {
	id       TypeId
	slice    any //function(cap,len)[]T
	instance any //function(cap,len)[]T
}

func (t TypeInfo) Id() TypeId {
	return t.id
}
func (t TypeInfo) Slice() any {
	return t.slice
}
func (t TypeInfo) Instance() any {
	return t.instance
}

type ReflectTypeInfo struct {
	id       TypeId
	slice    any
	instance any
}

func (t *ReflectTypeInfo) Id() TypeId {
	return t.id
}
func (t *ReflectTypeInfo) Slice() any {
	if t.slice == nil {
		t.slice = slice(t.id)
	}
	return t.slice
}
func (t *ReflectTypeInfo) Instance() any {
	if t.instance == nil {
		t.instance = instance(t.id)
	}
	return t.instance
}
func slice(id TypeId) func(n ...int) any {
	return func(n ...int) any {
		switch len(n) {
		case 0:
			return reflect.MakeSlice(id.typo(), 0, 4).Interface()
		case 1:
			return reflect.MakeSlice(id.typo(), 0, n[0]).Interface()
		default:
			return reflect.MakeSlice(id.typo(), n[1], n[0]).Interface()
		}
	}
}
func instance(id TypeId) func() any {
	return func() any {
		return reflect.New(id.typo())
	}
}

var (
	TypeRegistry map[TypeId]TypeInfo
	TypeAlias    map[TypeId]TypeId
)

func RegisterTypeAlias(src any, target any) {
	st := TypeOf(src)
	if _, ok := TypeAlias[st]; ok {
		return
	}
	TypeAlias[st] = TypeOf(target)
}
func RegisterType[T any]() bool {
	var x T
	t := TypeOf(x)
	if _, ok := TypeRegistry[t]; ok {
		return false
	}
	v := TypeInfo{
		id: t,
		slice: func(p ...int) []T {
			switch len(p) {
			case 0:
				return make([]T, 0)
			case 1:
				return make([]T, 0, p[0])
			default:

				return make([]T, p[1], p[0])
			}
		},
		instance: func() T {
			var v T
			return v
		},
	}
	TypeRegistry[t] = v
	return true
}
