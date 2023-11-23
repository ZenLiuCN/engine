package engine

import (
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestSimple(t *testing.T) {
	e := NewEngine()
	println(IsNullish(fn.Panic1(e.RunJavaScript("console.log('123'+'2324')"))))
	println(IsNullish(fn.Panic1(e.RunTypeScript(`
	const adder=(i:number)=>i+1
	const added=(u:number)=>{let x=u;return (i:number)=>x+i}
	console.log(adder(1234))
	const add=added(5)
	console.log(add(1))
	console.log(add(2))
	console.log(add(3))
`))))

}
func TestSimpleRequire(t *testing.T) {
	e := Get()
	defer e.Free()
	println(IsNullish(fn.Panic1(e.RunTypeScript(`
	import {some} from './simple.js'
	console.log(typeof some)
	console.log(some())
`))))
	println(IsNullish(fn.Panic1(e.RunJavaScript(`
	import {some} from './simple.js'
	console.log(typeof some)
	console.log(some())
`))))

}
