package engine

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestHasher_RawStdDecode(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
codec.base64RawStdEncode(binFrom(1,2,3,4,5,6))

`))
	if x, ok := v.Export().(string); !ok {
		panic(v)
	} else if x != "AQIDBAUG" {
		t.Fatal(x, "AQIDBAUG")
	}

}
func TestHasher_Base64(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunScript(
		//language=javascript
		`
const bin=binFrom(1,2,3,4,5,6)
const b64="AQIDBAUG"
const hex="010203040506"
console.assert(codec.base64RawStdEncode(bin)===b64)
console.assert(codec.base64StdEncode(bin)===b64)
console.assert(codec.base64UrlEncode(bin)===b64)
console.assert(codec.hexEncode(bin)===hex)
console.assert(binEquals(codec.base64RawUrlDecode(b64),bin))
console.assert(binEquals(codec.base64RawStdDecode(b64),bin))
console.assert(binEquals(codec.base64StdDecode(b64),bin))
console.assert(binEquals(codec.base64UrlDecode(b64),bin))
console.assert(binEquals(codec.base64RawUrlDecode(b64),bin))
console.assert(binEquals(codec.hexDecode(hex),bin))
`))

}
func TestHasher_Hash(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunScript(
		//language=javascript
		`
const bin=binFrom(1,2,3,4,5,6)
const b64="AQIDBAUG1B2M2Y8AsgTpgAmY7PhCfg=="
console.assert(codec.base64StdEncode(hasher.MD5.get().sum(bin))===b64)
`))

}
