package engine

import (
	"fmt"
)

func TypeOf(v any) GenericType {
	return GenericType(fmt.Sprintf("%T", v))
}
func TypeUsage(t GenericType) *TypeInfo {
	v, ok := TypeRegistry[t]
	if ok {
		return &v
	}
	return nil
}

type TypeInfo struct {
	StructType bool
	SliceType  bool
	Slice      any
}

var (
	TypeRegistry map[GenericType]TypeInfo
)

func init() {
	//TODO common types
	TypeRegistry = map[GenericType]TypeInfo{}
	RegisterSliceType[[]int32]()
}

type GenericType string

func (t GenericType) ToString() string {
	return string(t)
}
func RegisterType[T any]() bool {
	var x T
	t := GenericType(fmt.Sprintf("%T", x))
	if _, ok := TypeRegistry[t]; ok {
		return false
	}
	v := TypeInfo{
		Slice: nil,
	}
	TypeRegistry[t] = v
	return true
}
func RegisterSliceType[T ~[]X, X any]() bool {
	var x T
	t := GenericType(fmt.Sprintf("%T", x))
	if _, ok := TypeRegistry[t]; ok {
		return false
	}
	v := TypeInfo{
		SliceType: true,
		Slice: func(size ...int) T {
			switch len(size) {
			case 0:
				return make(T, 0)
			case 1:
				return make(T, size[0])
			default:
				return make(T, size[0], size[1])
			}
		},
	}
	TypeRegistry[t] = v
	return true
}
