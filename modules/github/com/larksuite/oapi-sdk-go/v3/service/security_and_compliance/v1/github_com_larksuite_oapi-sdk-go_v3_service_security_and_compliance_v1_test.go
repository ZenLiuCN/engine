package larksecurity_and_compliance

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
		import * as larksecurity_and_compliance from 'github.com/larksuite/oapi-sdk-go/v3/service/security_and_compliance/v1'
		`))
}
