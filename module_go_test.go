package engine

import (
	"github.com/ZenLiuCN/fn"
	"testing"
	"time"
)

func TestChannel_bi(t *testing.T) {
	vm := Get()
	defer vm.Free()
	ch := make(chan int)
	vm.Set("ch", NewChan(ch, vm))
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
	import {Chan} from 'go'
	/** @type {Chan<number>}*/
	const cc=ch
	cc.recv((v)=>console.log(v)).then(()=>console.log("closed"))
`))
	tick := time.Tick(time.Millisecond)
	go func() {
		i := 0
		for {
			select {
			case <-tick:
				close(ch)
				return
			default:
				i++
				ch <- i
			}
		}
	}()
	vm.Await()
}
func TestChannel_out(t *testing.T) {
	vm := Get()
	defer vm.Free()
	ch := make(chan int)
	vm.Set("ch", NewChanWriteOnly(ch))
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
	import {WriteOnlyChan} from 'go'
	/** @type {WriteOnlyChan<number>}*/
	const cc=ch
	const i=setInterval(()=>cc.send(1),10)
	setTimeout(()=>{
        clearInterval(i)
        cc.close()
	},100)
`))
	go func() {
		for {
			select {
			case i, ok := <-ch:
				if ok {
					println(i)
				} else {
					return
				}
			}
		}
	}()
	vm.Await()
}
func TestChannel_in(t *testing.T) {
	vm := Get()
	defer vm.Free()
	ch := make(chan int)
	vm.Set("ch", NewChanReadOnly(ch, vm))
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
	import {ReadOnlyChan} from 'go'
	/** @type {ReadOnlyChan<number>}*/
	const cc=ch
	cc.recv((v)=>console.log(v)).then(()=>console.log("closed"))
`))
	tick := time.Tick(time.Millisecond)
	go func() {
		i := 0
		for {
			select {
			case <-tick:
				close(ch)
				return
			default:
				i++
				ch <- i
			}
		}
	}()
	vm.Await()
}

func TestConvert(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunTs(
		//language=typescript
		`
					
		import {runesFromString, stringFromRunes, typeOf, usage} from 'go'
		import type {rune} from 'go'
					const r=runesFromString("123ABCΔ")
					console.log(r);
					console.log(stringFromRunes(r))
					r.push(49)
					console.log(stringFromRunes(r))
					console.log(typeOf(r))
					const slice:rune[]=usage(typeOf(r)).slice(2)
					console.log(slice )
					slice[0]=49
					slice[1]=49
					console.log(slice)
					console.log(typeOf(slice),typeof slice)
					console.log(stringFromRunes(slice))
					console.log(slice.map(v=>v))
					console.log(typeOf(slice.map(v=>v)))
				`))

	vm.Await()
}
