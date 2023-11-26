package engine

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestCryptoSimple(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunScript(
		//language=javascript
		`const enc = Symmetric.cipher({
        cipher: Symmetric.aes(),
        padding: Symmetric.pkcs5(),
        block:Symmetric.ecb(true),
        encrypt: true
    }
).crypto(binFrom("1234567812345678"), binFrom("12345678901234"))
console.log(codec.hexEncode(enc))
`))
}
