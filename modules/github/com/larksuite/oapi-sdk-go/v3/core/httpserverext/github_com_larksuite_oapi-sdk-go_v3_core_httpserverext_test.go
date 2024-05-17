package httpserverext

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
		import * as httpserverext from 'github.com/larksuite/oapi-sdk-go/v3/core/httpserverext'
		`))
}
