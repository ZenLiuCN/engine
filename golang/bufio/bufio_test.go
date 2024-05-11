package bufio

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
		import * as bufio from 'golang/bufio'
		const r=bufio.refReader()
		console.log(r.buffered())
		const [c,e]=r.discard(2)
		console.log(c,e)
		`))
}
