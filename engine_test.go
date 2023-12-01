package engine

import (
	"github.com/ZenLiuCN/fn"
	"os"
	"testing"
	"time"
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
func TestEs2015(t *testing.T) {
	e := NewEngine()
	println(IsNullish(fn.Panic1(e.RunScript(
		//language=javascript
		`
	class A {
    constructor(age){
        this.age=age
    }
	}
console.log(typeof A)
const x=1
console.log(Number.EPSILON)
console.log(Number.MIN_SAFE_INTEGER)
console.log(Number.MAX_SAFE_INTEGER)
console.log(Math.trunc(1.2))
console.log(Math.sign(1.2))
console.log({[x+"b"]:1})
console.log(new A(1))
`))))
	e.Await()
}
func TestEs2016(t *testing.T) {
	e := NewEngine()
	println(IsNullish(fn.Panic1(e.RunScript(
		//language=javascript
		`
	console.log(2**12)
	console.log([1,2].includes(1))
`))))

}
func TestEs2017(t *testing.T) {
	e := NewEngine()
	println(IsNullish(fn.Panic1(e.RunScript(
		//language=javascript
		`
const b=()=>new Promise((r,j)=>r(1))
const c= async ()=> await b()
console.log(typeof b)
console.log(b.name)
c().then(v=>console.log('future',v))
// console.log(await c()) TOP Level not supported yet
console.log(Object.values({a:1,b:2}))
console.log(Object.entries({a:1,b:2}))
console.log(Object.getOwnPropertyDescriptors({a:1,b:2}))
const str = 'adata'
console.log(str.padStart(10))
console.log(str.padEnd(10))

`))))
	e.Await()
}
func TestEs2018(t *testing.T) {
	e := NewEngine()
	println(IsNullish(fn.Panic1(e.RunScript(
		//language=javascript
		`
/*const asyncIterable = {
  [Symbol.asyncIterator]() {
    return {
      i: 0,
      next() {
        if (this.i < 3) {
          return Promise.resolve({ value: this.i++, done: false })
        }
        return Promise.resolve({ done: true })
      }
    }
  }
}
(async function () {
  for await (num of asyncIterable) {
    console.log(num)
  }
})()*/
console.log("asyncIterable not supported yet")
const n = { name: 'som' }
const a = { age: 182 }
const person = { ...n, ...a }
console.log(person)
const syn=()=>new Promise((r,j)=>r(1))
syn()
.then(v=>console.log(v))
.finally(()=>{
    console.log("finally will never end")
    return new Promise((r,j)=>r())
})
`))))
	e.AwaitTimeout(time.Second)
}

func TestEs2019(t *testing.T) {
	e := NewEngine()
	println(IsNullish(fn.Panic1(e.RunScript(
		//language=javascript
		`
console.log([1,2,3,[4,5,6]].flat())
const a=()=>console.log(1)
console.log(Object.fromEntries( [["a",1],["b",2]]))
console.log(a.toString())
`))))
	e.Await()
}
func TestEs2020(t *testing.T) {
	e := NewEngine()
	println(IsNullish(fn.Panic1(e.RunScript(
		//language=javascript
		`
console.log(globalThis)
// import('./promise.js').then(r=>console.log("supported"))
// console.log(BigInt("0x1fffffffffffff"))
console.log(null ?? 10) 
console.log(undefined ?? 10) 
console.log(false ?? 10) 
const a={b:{c:'c'}}
console.log(a.b?.b?.c)
`))))
	e.Await()
}
func TestEs2021(t *testing.T) {
	e := NewEngine()
	println(IsNullish(fn.Panic1(e.RunScript(
		//language=javascript
		`
const str = 'sdfoooodbb'
const newStr = str.replaceAll('o', 'z')
console.log(newStr)
// console.log(1_0000_0000) 

`))))
	e.Await()
}
func TestEs2022(t *testing.T) {
	e := NewEngine()
	println(IsNullish(fn.Panic1(e.RunScript(
		//language=javascript
		`
class C1 {
  myName = '11一碗'
  #dsdf='123'
  get x(){
      return this.#dsdf
  }
}
class C2 {
  constructor() {
    this.myName = "fdafzhah汉族"
  }
}
console.log(new C1(),new C2())
console.log(new C1().x)
//TOP level await not supported yet
const b=()=>new Promise((r,j)=>r(1))
const c= async ()=> await b()
console.log(typeof b)
console.log(b.name)
c().then(v=>console.log('future',v))
// console.log(await c()) TOP Level not supported yet
console.log(Object.hasOwn(new C1(),'#dsdf'))
console.log(Object.hasOwn(new C1(),'myName'))
console.log([1,2,3].at(-1))
`))))
	e.Await()
}
func TestSimpleRequire(t *testing.T) {
	fn.Panic(os.WriteFile("simple.js", []byte(`
export function some() {
return 1
}`), os.ModePerm))
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
	_ = os.Remove("simple.js")

}
