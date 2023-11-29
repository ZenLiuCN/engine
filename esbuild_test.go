package engine

import (
	"github.com/ZenLiuCN/fn"
	"github.com/evanw/esbuild/pkg/api"
	"testing"
)

func TestEsBuild(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(
		vm.RunJavaScript(
			//language=javascript
			`
import es from 'go/esbuild'
es.transform("const a=1",{
    target:1
})
`)).Export().(api.TransformResult)
	t.Log(api.FormatMessages(v.Errors, api.FormatMessagesOptions{Kind: 1}), string(v.Code))

}
