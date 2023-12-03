# Engine

A javascript Engine base on [Goja](https://github.com/dop251/goja) inspired by [k6](https://github.com/grafana/k6)

## Limitations and modifications

+ Top level async/await not supported by Goja
+ ES6+ partially implemented by Goja
+ Async operations should wait for ends with `engine.StopEventLoopXXX`
+ Module includes remote/local/GoModule support by translate to CommonJs (EsBuild)
## Sample
### no time limit
```go

package main

import (
   "github.com/ZenLiuCN/engine"
)
func main() {
   vm := engine.Get()
   defer vm.Free()
   v, err := vm.RunJs(
      `
console.log("Begin "+"For timeout")
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
})`)
   halts := vm.Await() //manual await task done
   if !halts.IsZero(){
	   panic("task not done")
   }
   if err!=nil{
      panic(err)
   }
   println(v)
}
```
### manual
```go
package main

import (
   "github.com/ZenLiuCN/engine"
   "time"
)
func main() {
   vm := engine.Get()
   defer vm.Free()
   v, err := vm.RunJs(
      `
console.log("Begin "+"For timeout")
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
})`)
   halts := vm.AwaitTimeout(time.Second * 5) //manual await task done for limited time
   if !halts.IsZero(){
	   panic("task not done within 5 seconds")
   }
   if err!=nil{
      panic(err)
   }
   println(v)
}

```
### automatic

```go
package main

import (
	"github.com/ZenLiuCN/engine"
	"time"
)

func main() {
	vm := engine.Get()
	defer vm.Free()
	v, err := vm.RunJsTimeout(`import {Second, sleep} from 'go/time'

new Promise((r, j) => {
    sleep(Second)
    r(1)
}).then(v => {
    console.log(v)
    return new Promise((r, j) => {
        sleep(Second)
        r(v+1)
    })
}).then(v => {
    console.log(v)
    return new Promise((r, j) => {
        sleep(Second)
        r(v+1)
    })
}).then(v => {
    console.log(v)
    return new Promise((r, j) => {
        sleep(Second*2)
        r(v+1)
    })
}).then(v => {
    console.log(v)
    return new Promise((r, j) => {
        sleep(Second*2)
        r(v+1)
    })
})
`, time.Second*8)
	if err != nil {
		panic(err) // if error is ErrTimeout, the value is HaltJobs
	}
	println(v)
}

```
## Extensions

### Engine

Simple wrapper for goja.Runtime with extended abilities.

### Code

Simple wrapper for goja.Program

### Naming

1. use struct tag `js:"name"` to mapping fields to js property
2. use struct tag `js:"-"` to ignore export field
3. default strategy for both functions methods and fields
    + `UUUlllUll` => `uuUllUll`
    + `UlllUll` => `ullUll`
    + `XlllUll` => `llUll`
    + `XUllUll` => `UllUll`

### Resolver with `require`

Simple support for CommonJs, ES script and TypeScript also compiled as CJS script, inspire by k6

### Module System

1. Module: a simple global instance
2. InitializeModule:  a module instanced when register to an Engine
3. TopModule: a module register some top level function or objects
4. TypeDefined: a module contains `d.ts` data
5. GoModule: a module for expose golang into js (use by import module)

### compiler module

+ `go/compiler` built-in compiler for both typescript and javascript
+ `go/esbuild` expose esbuild to js

**components**

+ [esbuild](https://github.com/evanw/esbuild)

### engine module

+ `go/engine`: use engine in scripts, _maybe not to use too many russian dolls_

### console module

+ global : slog console or byte buffer console

### buffer module

+ `go/buffer`: golang byte slice and bytes.Buffer

### hash module

+ `go/hash`:golang hash functions
+ `go/codec`:golang codec functions, include base64 ...

### crypto module

+ `go/crypto` : golang crypto

### os module

+ `go/os` : golang os , not full functions

### io module

+ `go/io` : golang io module
### chan module

+ `go/chan` : golang chan type only

## Modules

pluggable modules

### sqlx

+ `go/sqlx`: sqlx with `pgx` `mysql` and `sqlite` driver

**components**

+ [sqlx](https://github.com/jmoiron/sqlx),
+ [sqlite](https://github.com/glebarez/go-sqlite)
+ [mysql](https://github.com/go-sql-driver/mysql) *modified* see source at `sqlx/mysql`

### Excelize

+ `go/excel`: excel reading or generate

**components**

+ [excelize](https://github.com/xuri/excelize/)

### fetch

+ `go/fetch`: base function done, improve for compact with browser fetch.

### pug

+ `go/pug`: jade(pug) compiler

 **components**

+ [jade](https://github.com/Joker/jade)

### minify

+ `go/minify`: file minify

**components**

+ [minify](https://github.com/tdewolff/minify)

### pdf

dev

### jose

draft

### http

draft
