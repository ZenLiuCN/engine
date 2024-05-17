package larkapplication

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
		import * as larkapplication from 'github.com/larksuite/oapi-sdk-go/v3/service/application/v6'
		`))
}
