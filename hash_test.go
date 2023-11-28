package engine

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestHash_Hash(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunJavaScript(
		//language=javascript
		`
import codec from 'go/codec'
import hash from 'go/hash'
const bin=new Bytes(1,2,3,4,5,6)
const b64="AQIDBAUG1B2M2Y8AsgTpgAmY7PhCfg=="
console.assert(codec.base64StdEncode(hash.MD5.get().sum(bin.bytes()))===b64)
`))

}
