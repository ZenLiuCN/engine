package engine

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestCodecModule_RawStdDecode(t *testing.T) {
	vm := Get()
	defer vm.Free()
	_ = fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Base64RawStd} from 'go/codec'
import {Bytes} from "go/buffer"
console.assert(Base64RawStd.encodeToString(new Bytes(1,2,3,4,5,6).bytes())==="AQIDBAUG")

`))

}
func TestCodecModule_Base64(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Bytes} from "go/buffer"
import codec from 'go/codec'
const bin=new Bytes(1,2,3,4,5,6)
const b64="AQIDBAUG"
const hex="010203040506"
console.assert(codec.Base64RawStd.encodeToString(bin.bytes())===b64)
console.assert(codec.Base64Std.encodeToString(bin.bytes())===b64)
console.assert(codec.Base64Url.encodeToString(bin.bytes())===b64)
console.assert(codec.Hex.encodeToString(bin.bytes())===hex)
console.assert(bin.equals(codec.Base64RawUrl.decodeString(b64)))
console.assert(bin.equals(codec.Base64RawStd.decodeString(b64)))
console.assert(bin.equals(codec.Base64Std.decodeString(b64)))
console.assert(bin.equals(codec.Base64Url.decodeString(b64)))
console.assert(bin.equals(codec.Base64RawUrl.decodeString(b64)))
console.assert(bin.equals(codec.Hex.decodeString(hex)))
`))

}
