package engine

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestSharedConsole_Table(t *testing.T) {
	e := Get()
	defer e.Free()
	fn.Panic1(e.RunJs(
		//language=javascript
		`
console.table([1,23,4,25])
console.table([{name:'asf',age:2},{name:'ss',age:24}])
console.table({name:'asf',age:2})
console.table({first:{name:'asf',age:2},second:{name:'ss',age:24}})
console.table({first:{name:'asf',age:2},second:{name:'ss',age:24}},['age'])
console.table([{name:'asf',age:2},{name:'ss',age:24}],["name"])
`))

}
