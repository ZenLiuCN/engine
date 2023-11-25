package engine

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
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
new Buffer("123456")
`))
	if buf, ok := v.Export().(*Buffer); !ok {
		panic("not a buffer")
	} else if buf.String() != "123456" {
		panic("bad buffer")
	}
}
func TestBuffer_ArrayBuffer(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
new Buffer("123456").arrayBuffer()
`))
	if buf, ok := v.Export().(goja.ArrayBuffer); !ok {
		panic("not a array buffer")
	} else if string(buf.Bytes()) != "123456" {
		panic("bad buffer")
	}
}
func TestBuffer_EachChar(t *testing.T) {
	vm := Get()
	defer vm.Free()
	fn.Panic1(vm.RunScript(
		//language=javascript
		`
new Buffer("123456").eachChar(u=>{
    console.log(u)
    return true
})
`))

}
func TestBuffer_MapChar(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
new Buffer("123456汉").mapChar(u=>{
    console.log(u)
    return true
})
`))
	if x, ok := v.Export().([]any); !ok {
		t.Fatal(v.String())
	} else if len(x) != 7 {
		t.Fatal(x...)
	} else {
		t.Log(x...)
	}

}
func TestBuffer_MapU8(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
new Buffer("123456中").mapU8(u=>{
    console.log(u)
    return true
})
`))
	if x, ok := v.Export().([]any); !ok {
		t.Fatal(v.String())
	} else if len(x) != 9 {
		t.Fatal(x...)
	} else {
		t.Log(x...)
	}

}
func TestBuffer_Binary(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
new Buffer("123456中").binary().map(v=>v.toString(10))
`))
	if x, ok := v.Export().([]any); !ok {
		t.Fatal(v.String())
	} else if len(x) != 9 {
		t.Fatal(x)
	} else {
		t.Log(x)
	}
}
func TestBuffer_Binary2(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
new Buffer("123456中").binary().map(v=>v)
`))
	if x, ok := v.Export().([]any); !ok {
		t.Fatal(v.String())
	} else if len(x) != 9 {
		t.Fatal(x)
	} else {
		t.Log(x)
	}
}
func TestBuffer_Bytes(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
new Buffer("123456中").bytes().map(v=>v)
`))
	if x, ok := v.Export().([]any); !ok {
		t.Fatal(v.String())
	} else if len(x) != 9 {
		t.Fatal(x)
	} else {
		t.Log(x)
	}
}
func TestBuffer_toBinary(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
 binFrom(1,2,3,4,5,6)
`))
	if x, ok := v.Export().([]byte); !ok {
		t.Fatal(v.String())
	} else if len(x) != 6 {
		t.Fatal(x)
	} else {
		t.Log(x)
	}
}
func TestBuffer_toBinary2(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
 binFrom([1,2,3,4,5,6])
`))
	if x, ok := v.Export().([]byte); !ok {
		t.Fatal(v.String())
	} else if len(x) != 6 {
		t.Fatal(x)
	} else {
		t.Log(x)
	}
}
func TestBuffer_LengthBinary(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
 binLen(binFrom([1,2,3,4,5,6]))
`))
	if x, ok := v.Export().(int64); !ok {
		t.Fatal(v.String())
	} else if x != 6 {
		t.Fatal(x)
	}
}
func TestBuffer_AtBinary(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
 binAt(binFrom([1,2,3,4,5,6]),2)
`))
	if x, ok := v.Export().(int64); !ok {
		t.Fatal(v.String())
	} else if x != 3 {
		t.Fatal(x)
	}
}
func TestBuffer_SliceBinary(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
 binSlice(binFrom([1,2,3,4,5,6]),2,3)
`))
	if x, ok := v.Export().([]byte); !ok {
		t.Fatal(v.String())
	} else if len(x) != 1 {
		t.Fatal(x)
	} else if x[0] != 3 {
		t.Fatal(x)
	}
}
func TestBuffer_MapBinary(t *testing.T) {
	vm := Get()
	defer vm.Free()
	v := fn.Panic1(vm.RunScript(
		//language=javascript
		`
 binMap(binFrom([1,2,3,4,5,6]),(v,i)=>({val:v,index:i}))
`))
	if x, ok := v.Export().([]any); !ok {
		t.Fatal(v.String())
	} else if len(x) != 6 {
		t.Fatal(x)
	} else {
		var val []map[string]any
		fn.Panic(vm.ExportTo(v, &val))
		t.Logf("%#v", val)
	}

}
