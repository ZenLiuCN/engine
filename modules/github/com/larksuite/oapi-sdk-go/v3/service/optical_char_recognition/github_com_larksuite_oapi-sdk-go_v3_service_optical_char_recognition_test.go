package optical_char_recognition

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestSimple(t *testing.T) {
	v := engine.Get()
	defer v.Free()
	fn.Panic1(v.RunTs(
		//language=typescript
		`
		import * as optical_char_recognition from 'github.com/larksuite/oapi-sdk-go/v3/service/optical_char_recognition'
		`))
}
