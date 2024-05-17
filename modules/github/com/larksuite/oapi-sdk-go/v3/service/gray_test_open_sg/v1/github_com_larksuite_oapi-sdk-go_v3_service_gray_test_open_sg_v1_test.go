package larkgray_test_open_sg

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
		import * as larkgray_test_open_sg from 'github.com/larksuite/oapi-sdk-go/v3/service/gray_test_open_sg/v1'
		`))
}
