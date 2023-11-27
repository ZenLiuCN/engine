package fetch

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestFetch_Fetch(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	fn.Panic1(vm.RunScript(
		//language=javascript
		`
fetch("https://163.com/")
.then(r=>r.text())
.then(t=>console.log('<<<\n',t,'\n>>>'))
`))
	t.Log(vm.StopEventLoopWait())
}
