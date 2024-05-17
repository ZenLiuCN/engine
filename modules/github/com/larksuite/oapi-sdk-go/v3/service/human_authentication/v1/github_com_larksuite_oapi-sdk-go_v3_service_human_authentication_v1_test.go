package larkhuman_authentication

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
		import * as larkhuman_authentication from 'github.com/larksuite/oapi-sdk-go/v3/service/human_authentication/v1'
		`))
}
