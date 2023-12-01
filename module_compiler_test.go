package engine

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestCompileSimple(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunJavaScript(
		//language=javascript
		`
import cm from 'go/compiler'
console.log(cm.compileJs("function am(){return 1} am()"))
console.log(cm.compileTs("function am():number{return 1} am()"))
`))
}
