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
    setTimeout(()=>r(1),1000)
}).then(v => {
    console.log("job",v)
       counter.inc()
    return new Promise((r, j) => {
        setTimeout(()=>r(v+1),1000)
    })}).then(v => {
    console.log("job",v)
       counter.inc()
    return new Promise((r, j) => {
        setTimeout(()=>r(v+1),1000)
    })
}).then(v => {
    console.log("job",v)
      counter.inc()
    return new Promise((r, j) => {
       setTimeout(()=>r(v+1),2000)
    })
}).then(v => {
    console.log("job",v)
        counter.inc()
    return new Promise((r, j) => {
       r(counter.cnt())
    })
})
out()
`
const jsPartlyTimeoutAsync =
// language=javascript
`
let n=0
for(var i = 0; i < 100; i++) {
  n+=i
}
console.log("Begin "+"For timeout with "+n)
new Promise((r, j) => {
    console.log("job 0")
    setTimeout(()=>r(1),1000)
}).then(v => {
    console.log("job",v)
    return new Promise((r, j) => {
        setTimeout(()=>r(v+1),1000)
    })}).then(v => {
    console.log("job",v)
    return new Promise((r, j) => {
        setTimeout(()=>r(v+1),1000)
    })
}).then(v => {
    console.log("job",v)
    return new Promise((r, j) => {
       setTimeout(()=>r(v+1),2000)
    })
}).then(v => {
    console.log("job",v)
    return new Promise((r, j) => {
       setTimeout(()=>r(v+1),2000)
    })
})`

func TestTimeout(t *testing.T) {
	e := Get()
	defer e.Free()
	v, err := e.RunJs(jsFullTimeoutAsync)
	if err != nil {
		t.Log(err)
	}
	halts := e.AwaitTimeout(time.Second * 5)
	if !halts.IsZero() {
		e.Interrupt("shutdown for timeout")
	}
	if halts.IsZero() {
		panic(fmt.Errorf("should not done:%s,%#v ", halts.String(), e.EventLoop))
	}
	t.Log(v)
}
func TestAsyncDone(t *testing.T) {
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
func TestTimeoutDelayJsSuccess(t *testing.T) {
	e := Get()
	defer e.Free()
	v := fn.Panic1(e.RunJsTimeoutDelay(jsFullTimeoutAsync, time.Microsecond*500, time.Second*8)).String()
	if v != "5" {
		panic("not real done")
	}

}
func TestTimeoutDelayJsFailure(t *testing.T) {
	e := Get()
	defer e.Free()
	v, err := e.RunJsTimeoutDelay(jsFullTimeoutAsync, time.Microsecond*10, time.Second*5)
	if err == nil {
		panic("should not success " + v.String())
	} else {
		t.Log(v, err)
	}

}
func TestContextJs(t *testing.T) {
	e := Get()
	defer e.Free()
	ctx, cc := context.WithTimeout(context.Background(), time.Second*5)
	defer cc()
	v, err := e.RunJsContext(jsFullTimeoutAsync, ctx)
	if err == nil {
		panic("should not success " + v.String())
	} else {
		t.Log(v, err)
	}

}
func TestContextJsSuccess(t *testing.T) {
	e := Get()
	defer e.Free()
	ctx, cc := context.WithTimeout(context.Background(), time.Second*8)
	defer cc()
	t.Log(fn.Panic1(e.RunJsContext(jsFullTimeoutAsync, ctx)))

}
