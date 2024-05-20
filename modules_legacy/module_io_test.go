package legacy

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestIoModule(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
import io from 'go/io'
import {Buffer,Bytes} from 'go/buffer'
const b=new Buffer()
console.assert(io.copy(b.toWriter(),new Buffer(Uint8Array.of(1,2,3,4,5)).toReader())===5)
console.assert(io.copyBuffer(b.toWriter(),new Buffer(Uint8Array.of(1,2,3,4,5)).toReader(),Uint8Array.of(1,2,3))===5)
console.assert(new Bytes(io.readAll(b.toReader())).equals(new Bytes(1,2,3,4,5,1,2,3,4,5)))
`))
}
