package main

import (
	"context"
	"flag"
	"github.com/ZenLiuCN/engine"
	_ "github.com/ZenLiuCN/engine/excelize"
	_ "github.com/ZenLiuCN/engine/fetch"
	_ "github.com/ZenLiuCN/engine/minify"
	_ "github.com/ZenLiuCN/engine/pug"
	_ "github.com/ZenLiuCN/engine/sqlx"
	"github.com/ZenLiuCN/fn"
	"sync"
	"time"
)

var (
	typed  bool
	source bool
	define bool
)

func main() {
	flag.BoolVar(&typed, "t", false, "use typescript source")
	flag.BoolVar(&source, "s", false, "use script source,default false")
	flag.BoolVar(&define, "d", false, "export defines to files,default false")
	flag.Parse()
	args := flag.Args()
	switch {
	case define:
		engine.DumpDefines(".")
	case len(args) == 1 && !source:
		vm := engine.Get()
		defer vm.Free()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		cc := fn.WithSignal(func(ctx context.Context) {
			defer wg.Done()
			v := fn.Panic1(vm.RunCodeContext(engine.CompileFile(args[0]), time.Millisecond*100, ctx))
			if !engine.IsNullish(v) {
				println(v.String())
			}
		})
		wg.Wait()
		defer cc()
	case len(args) > 0:
		vm := engine.Get()
		defer vm.Free()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		cc := fn.WithSignal(func(ctx context.Context) {
			defer wg.Done()
			v := fn.Panic1(vm.RunCodeContext(engine.CompileSource(fn.SliceJoinRune(args, '\n', fn.Identity[string]), typed), time.Millisecond*100, ctx))
			if !engine.IsNullish(v) {
				println(v.String())
			}
		})
		wg.Wait()
		defer cc()
	default:
		println("usage: " + engine.GetExecutable() + " [FLAG] source|file")
		flag.Usage()
	}

}
