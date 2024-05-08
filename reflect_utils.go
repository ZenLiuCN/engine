package engine

import (
	"reflect"
	"unsafe"
)

func Reader[V any, K any](fieldIndex int) func(*V) K {
	return func(t *V) K {
		v := reflect.ValueOf(t).Elem()
		f := v.Field(fieldIndex)
		rf := reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
		return rf.Interface().(K)
	}
}
