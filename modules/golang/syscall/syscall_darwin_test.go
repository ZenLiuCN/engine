package syscall

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestSimple_darwin(t *testing.T) {
	v := engine.Get()
	defer v.Free()
	fn.Panic1(v.RunTs(
		//language=typescript
		`
		import * as time from 'golang/syscall'
		`))
}
