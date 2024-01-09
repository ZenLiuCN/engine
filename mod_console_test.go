package engine

import (
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
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
func DoError(v int) (o int, err error) {
	if v > 0 {
		return 0, fmt.Errorf("%d", v)
	}
	return v, nil
}
func TestError(t *testing.T) {
	e := Get()
	defer e.Free()
	e.Set("dt", DoError)
	fn.Panic1(e.RunJs(
		//language=javascript
		`
	try{
    	dt(1)
	}catch (error){
    console.log(error.message)
	}
`))

}
func TestCtorError(t *testing.T) {
	e := Get()
	defer e.Free()
	e.Set("Data", e.ToConstructor(func(v []goja.Value) (any, error) {
		if len(v) != 0 {
			return nil, fmt.Errorf("bad arguments")
		}
		return "123", nil
	}))
	fn.Panic1(e.RunJs(
		//language=javascript
		`
	try{
    	let v=new Data(1)
    	console.log("success",v)
	}catch (error){
    console.log(error.message)
	}
`))

}
