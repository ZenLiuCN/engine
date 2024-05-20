package legacy

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestBig(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Int,Rat,zeroInt,oneRat,equals} from 'go/big'
const i=new Int("123456")
console.log(i)
console.log(typeof i)
console.log(i instanceof Int)
console.log(zeroInt() instanceof Int)
console.log(equals(i,new Int(123456)))
console.log(i.string())
const r=new Rat(1.5)
console.log(r.num())
console.log(r.denom())
console.log(r.sign())
console.log(oneRat() instanceof Rat)
console.log(oneRat().ratString())
console.log(new Rat("3/2").cmp(r))
console.log(new Rat("1/2").string())
console.log(new Rat("1/3").floatString(2))
`)).Export()

}
