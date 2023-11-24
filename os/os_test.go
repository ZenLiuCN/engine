package os

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestConstant(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	println(fn.Panic1(
		vm.RunJavaScript(
			//language=javascript
			`
f=()=>1
console.log(os.root)
console.log(os.ls())
f()
`)).Export().(int64) == 1)

}
