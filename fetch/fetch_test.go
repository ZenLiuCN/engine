package fetch

import (
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
	if vm.StopEventLoopWait() > 0 {
		panic("should complete async task")
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
	if vm.StopEventLoopTimeout(time.Millisecond*1) <= 0 {
		panic("should not complete async task")
	}
}
