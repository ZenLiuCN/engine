package engine

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestConstant(t *testing.T) {
	vm := Get()
	defer vm.Free()
	println(fn.Panic1(
		vm.RunJs(
			//language=javascript
			`
import os from "go/os"
const f=()=>1
console.log(os.root)
console.log(os.ls())
f()
`)).Export().(int64) == 1)

}
