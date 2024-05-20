package engine

import (
	"github.com/dop251/goja"
	"testing"
)

var (
	fun = Reader[goja.Exception, []goja.StackFrame](1)
)

func TestReader(t *testing.T) {
	vm := Get()
	_, err := vm.Runtime.RunString("throw new Error(123)")
	switch e := err.(type) {
	case *goja.Exception:
		frames := fun(e)
		t.Logf("%#v", frames)
	}
}
