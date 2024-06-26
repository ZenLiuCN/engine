package legacy

import (
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"testing"
)

func TestPut(t *testing.T) {
	b := GetBuffer()
	if b.Detached {
		t.Fatal("already detached")
	}
	b.Free()
	if !b.Detached {
		t.Fatal("not detached")
	}
}

func TestEngineBuffer(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Buffer} from 'go/buffer'
new Buffer("123456")
`))
	if buf, ok := v.Export().(*Buffer); !ok {
		panic("not a buffer")
	} else if buf.ToString() != "123456" {
		panic("bad buffer")
	}
}
func TestBuffer_ArrayBuffer(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Buffer} from 'go/buffer'
new Buffer("123456").arrayBuffer()
`))
	if buf, ok := v.Export().(goja.ArrayBuffer); !ok {
		panic("not a array buffer")
	} else if string(buf.Bytes()) != "123456" {
		panic("bad buffer")
	}
}
func TestBuffer_Bytes(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Buffer} from 'go/buffer'
new Buffer("123456").bytes()
`))
	if _, ok := v.Export().([]byte); !ok {
		t.Fatal("not bytes")
	}
}

func TestBytes_Get(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Bytes} from 'go/buffer'
const bytes=new Bytes("123456")
console.log(bytes instanceof Bytes)
bytes[0]
`))
	if x, ok := v.Export().(int64); !ok {
		t.Fatal("not number", v)
	} else if x != '1' {
		t.Fatal("not equal", x)
	}
}
func TestBytes_Set(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Bytes} from 'go/buffer'
const bytes=new Bytes("123456")
const cap=()=>bytes
console.log(bytes instanceof Bytes)
bytes[1]=2
cap()
`))
	if b, ok := v.Export().(*Bytes); !ok {
		t.Fatal("not Bytes", v)
	} else if b.b[1] != 2 {
		t.Fatal("not equal", b)
	}
}
func TestBytes_Clone(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Bytes} from 'go/buffer'
const bytes=new Bytes("123456")
console.log(bytes instanceof Bytes)
const b2=bytes.clone()
console.log(b2 instanceof Bytes)
const out=()=>b2[1]
bytes[1]=3
out()
`))
	if x, ok := v.Export().(int64); !ok {
		t.Fatal("not number", v)
	} else if x != '2' {
		t.Fatal("not equal", x)
	}
}
func TestBytes_ToString(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Bytes} from 'go/buffer'
const bytes=new Bytes("123456")
console.log(bytes instanceof Bytes)
bytes.toText()
`))
	if b, ok := v.Export().(string); !ok {
		t.Fatal("not string", v)
	} else if b != "123456" {
		t.Fatal("not equal", b)
	}
}
func BenchmarkDyn(b *testing.B) {
	vm := Get()
	defer vm.Free()
	b.ReportAllocs()
	code := vm.Compile(
		//language=javascript
		`
import {Bytes} from 'go/buffer'
const bytes=new Bytes("123456")
bytes[0]=33
bytes.toText()
`, false, true)
	for i := 0; i < b.N; i++ {
		txt := fn.Panic1(vm.RunCode(code)).Export().(string)
		if txt != "!23456" {
			panic("not equals " + txt)
		}
	}
}
