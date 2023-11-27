package fetch

import (
	"fmt"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
	"time"
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
	if n := vm.StopEventLoopWait(); n > 0 {
		panic(fmt.Sprintf("should complete async task: %d", n))
	}
}
func TestFetch_Fetch_timeout(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	fn.Panic1(vm.RunScript(
		//language=javascript
		`
fetch("https://163.com/")
.then(r=>r.text())
.then(t=>console.log('<<<\n',t,'\n>>>'))
`))
	if n := vm.StopEventLoopTimeout(time.Millisecond * 1); n != 2 {
		panic(fmt.Sprintf("should not complete async task: %d", n))
	}
}
