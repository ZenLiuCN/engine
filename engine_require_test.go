package engine

import (
	"github.com/ZenLiuCN/fn"
	"os"
	"strconv"
	"testing"
)

func TestRequireFunc(t *testing.T) {
	_ = os.WriteFile("func.js", []byte(
		//language=javascript
		`
export let v=0
export function Some(){
    console.log("run",v)
    return v++
}
`), os.ModePerm)
	defer func() {
		_ = os.Remove("func.js")
	}()
	e := NewEngine()
	call := fn.Panic1(e.RunJs(
		//language=javascript
		`
import {Some,v} from './func.js'
console.log('import')
console.log(Some())
const x=()=>v
x()
`)).Export().(int64)
	e.Await()
	if call != 1 {
		panic("should only call once: " + strconv.Itoa(int(call)))
	}
}
