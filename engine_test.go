package engine

import (
	"context"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"testing"
	"time"
)

const jsFullTimeoutAsync =
// language=javascript
`
console.log("Begin "+"For timeout")
const newCounter=()=>{
    let cnt=0
    return {
        inc:function (){
            cnt++
            console.log(cnt)
        },
        cnt:function (){
            console.log('get',cnt)
           return  cnt
        },
        toString:()=>cnt
    }
}
const counter=newCounter()
const out=()=>counter
new Promise((r, j) => {
    console.log("job 0")
    counter.inc()
    setTimeout(()=>r(1),500)
}).then(v => {
    console.log("job 1")
       counter.inc()
    return new Promise((r, j) => {
        setTimeout(()=>r(v+1),500)
    })}).then(v => {
    console.log("job 2")
       counter.inc()
    return new Promise((r, j) => {
        setTimeout(()=>r(v+1),500)
    })
}).then(v => {
    console.log("job 3")
      counter.inc()
    return new Promise((r, j) => {
       setTimeout(()=>r(v+1),500)
    })
}).then(v => {
    console.log("job 4")
        counter.inc()
    return new Promise((r, j) => {
       r(counter.cnt())
    })
})
out()
`

func TestTimeout(t *testing.T) {
	e := Get()
	defer e.Free()
	v, err := e.RunJs(jsFullTimeoutAsync)
	if err != nil {
		t.Log(err)
	}
	ctx, cc := context.WithTimeout(context.Background(), time.Second)
	defer cc()
	halts := e.AwaitWithContext(ctx)
	if !halts.IsZero() {
		e.Interrupt("shutdown for timeout")
	}
	if halts.IsZero() {
		panic(fmt.Errorf("should not done:%s,%#v ", halts.String(), e.EventLoop))
	}
	t.Log(v)
}
func TestAwait(t *testing.T) {
	e := Get()
	defer e.Free()
	v, err := e.RunJs(jsFullTimeoutAsync)
	if err != nil {
		t.Log(err)
	}
	halts := e.Await()
	if !halts.IsZero() {
		panic(fmt.Errorf("should done:%s,%#v ", halts.String(), e.EventLoop))
	}
	t.Log(v)
}
func TestContextJsSuccess(t *testing.T) {
	e := Get()
	defer e.Free()
	ctx, cc := context.WithTimeout(context.Background(), time.Second*3)
	defer cc()
	v := fn.Panic1(e.RunJsContext(jsFullTimeoutAsync, time.Millisecond, ctx)).String()
	if v != "5" {
		panic("not real done")
	}

}
func TestContextJsFailure(t *testing.T) {
	e := Get()
	defer e.Free()
	ctx, cc := context.WithTimeout(context.Background(), time.Second*2)
	defer cc()
	v, err := e.RunJsContext(jsFullTimeoutAsync, time.Millisecond, ctx)
	if err == nil {
		panic("should not success " + v.String())
	} else {
		t.Log(v, err)
	}

}
func TestAutoClose(t *testing.T) {
	e := Get()
	defer e.Free()
	fn.Panic1(e.RunJs(
		//language=javascript
		`
	import os from 'go/os'
	console.log(os)
	const f=autoClose(os.open('engine.go'))
	console.log(f.name())
`))

}
