package engine

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestEngineModuleSimple(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunScript(
		//language=javascript
		`const e=new Engine()
	console.log(e.runScript('1'))
	e.free()
	`))

}
func TestEngineTTL(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunScript(
		//language=javascript
		`
	setTimeout(()=>console.log("done"),100)
	console.log(typeof setTimeout)
	`))
	//println(vm.StopEventLoopTimeout(time.Millisecond * 200))
	println(vm.StopEventLoopWait())
}
