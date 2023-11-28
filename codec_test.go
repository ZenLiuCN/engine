package engine

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestCodecModule_RawStdDecode(t *testing.T) {
	vm := Get()
	defer vm.Free()
	_ = fn.Panic1(vm.RunJavaScript(
		//language=javascript
		`
import {base64RawStdEncode} from 'go/codec'
console.assert(base64RawStdEncode(new Bytes(1,2,3,4,5,6).bytes())==="AQIDBAUG")

`))

}
func TestCodecModule_Base64(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunJavaScript(
		//language=javascript
		`
import codec from 'go/codec'
const bin=new Bytes(1,2,3,4,5,6)
const b64="AQIDBAUG"
const hex="010203040506"
console.assert(codec.base64RawStdEncode(bin.bytes())===b64)
console.assert(codec.base64StdEncode(bin.bytes())===b64)
console.assert(codec.base64UrlEncode(bin.bytes())===b64)
console.assert(codec.hexEncode(bin.bytes())===hex)
console.assert(bin.equals(codec.base64RawUrlDecode(b64)))
console.assert(bin.equals(codec.base64RawStdDecode(b64)))
console.assert(bin.equals(codec.base64StdDecode(b64)))
console.assert(bin.equals(codec.base64UrlDecode(b64)))
console.assert(bin.equals(codec.base64RawUrlDecode(b64)))
console.assert(bin.equals(codec.hexDecode(hex)))
`))

}
