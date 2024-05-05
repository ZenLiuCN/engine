package time

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
	import * as time from 'golang/time'
	const now=time.now()
	console.log(now)
	console.log(now.add(time.duration(2 , time.Second)))
	console.log(now.add(time.duration(2 , time.Second)).sub(now).string())
	time.afterFunc(time.duration(2,time.Second),()=>console.log("done"))
	time.sleep(time.duration(3,time.Second))
`))
}
