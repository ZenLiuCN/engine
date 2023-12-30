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
func TestOsWriteFile(t *testing.T) {
	vm := Get()
	defer vm.Free()
	println(fn.Panic1(
		vm.RunJs(
			//language=javascript
			`
import os from "go/os"
import buffer from "go/buffer"
const f=()=>1
console.log(os.root)
console.table(os.ls())
os.writeFile('temp.ts',new buffer.Bytes(1,23,44).bytes(),777)
os.rename("temp.ts","temp1.ts")
os.remove("temp1.ts")
console.log(os.hostname())
f()
`)).Export().(int64) == 1)

}
