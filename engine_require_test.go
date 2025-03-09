package engine

import (
	"fmt"
	"github.com/ZenLiuCN/fn"
	"os"
	"strconv"
	"strings"
	"testing"
)

func validate(im, msg string) {
	defer func() {
		switch x := recover().(type) {
		case error:
			panic(fmt.Errorf("%s %w", msg, x))
		case nil:
		default:
			panic(fmt.Errorf("%s %+v", msg, x))
		}
	}()
	e := NewEngine()
	e.Debug = true
	defer e.Free()
	call := fn.Panic1(e.RunJs(
		//language=javascript
		fmt.Sprintf(`
import {Some,v} from '%s'
console.log('import')
console.log(Some())
const x=()=>v
x()
`, im))).Export().(int64)
	e.Await()
	if call != 1 {
		panic("should only call once: " + strconv.Itoa(int(call)))
	}
	_, _ = fmt.Fprintf(os.Stdout, "done: %s \n", msg)
}
func TestRequireFunc(t *testing.T) {
	code := []byte(
		//language=javascript
		`
export let v=0
export function Some(){
    console.log("run",v)
    return v++
}
`)
	_ = os.WriteFile("func.js", code, os.ModePerm)
	defer func() {
		_ = os.Remove("func.js")
	}()
	_ = os.Mkdir("func_", os.ModePerm)
	_ = os.WriteFile("func_/index.js", code, os.ModePerm)
	_ = os.WriteFile("func_/func.js", code, os.ModePerm)
	defer func() {
		_ = os.RemoveAll("func_")
	}()
	validate("func", "pwd without extension")
	validate("./func", "relative without extension")
	validate("./func_", "relative with index")
	validate("func_/", "pwd with index")
	validate("func_/index", "pwd with index")
	validate("func_/func", "pwd with filename")

}
func TestRequireWithFailure(t *testing.T) {

	e := NewEngine()
	e.Debug = true
	defer e.Free()
	_, err := e.RunJs(
		//language=javascript
		fmt.Sprintf(`
import {Some,v} from '%s'
console.log('import')
console.log(Some())
const x=()=>v
x()
`, "func"))
	e.Await()
	if err == nil || !strings.HasSuffix(err.Error(), "vm.js[2:21]\timport {Some,v} from ♦'func'") {
		panic(err)
	}

}
