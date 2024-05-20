package legacy

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestCryptoSimple(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Bytes} from "go/buffer"
import {cipher,aes,pkcs5,cbc} from "go/crypto"
import {base64StdEncode} from "go/codec"
const enc = cipher({
        cipher: aes(),
        padding: pkcs5(),
        block:cbc(new Bytes("1234567812345678").bytes(),true),
        encrypt: true
    }
).crypto(new Bytes("1234567812345678").bytes(), new Bytes("12345678901234").bytes())
console.log(base64StdEncode(enc))
`))
}
