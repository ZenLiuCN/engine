package main

import (
	"flag"
	"github.com/ZenLiuCN/engine"
	_ "github.com/ZenLiuCN/engine/excelize"
	_ "github.com/ZenLiuCN/engine/fetch"
	_ "github.com/ZenLiuCN/engine/os"
	_ "github.com/ZenLiuCN/engine/sqlx"
	"github.com/ZenLiuCN/fn"
	"os"
)

var (
	typed  bool
	source bool
	define bool
)

func main() {
	flag.BoolVar(&typed, "t", false, "use typescript source")
	flag.BoolVar(&source, "s", false, "use script source,default false")
	flag.BoolVar(&define, "d", false, "export define to file,default false")
	flag.Parse()
	args := flag.Args()
	switch {
	case len(args) == 1 && define:
		_ = os.WriteFile(args[0], engine.TypeDefines(), os.ModePerm)
	case len(args) == 1 && !source:
		vm := engine.Get()
		defer vm.Free()
		v := fn.Panic1(vm.Execute(engine.CompileFile(args[0])))
		vm.StopEventLoopWait()
		if !engine.IsNullish(v) {
			println(v.String())
		}
	case len(args) > 0:
		vm := engine.Get()
		defer vm.Free()
		v := fn.Panic1(vm.Execute(engine.CompileSource(fn.SliceJoinRune(args, '\n', fn.Identity[string]), typed)))
		vm.StopEventLoopWait()
		if !engine.IsNullish(v) {
			println(v.String())
		}
	default:
		println("usage: " + engine.GetExecutable() + " [FLAG] source|file")
		flag.Usage()
	}

}
