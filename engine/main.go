package main

import (
	"context"
	"fmt"
	"github.com/ZenLiuCN/engine"
	_ "github.com/ZenLiuCN/engine/chrome"
	_ "github.com/ZenLiuCN/engine/excelize"
	_ "github.com/ZenLiuCN/engine/fetch"
	_ "github.com/ZenLiuCN/engine/gse"
	_ "github.com/ZenLiuCN/engine/minify"
	_ "github.com/ZenLiuCN/engine/pug"
	_ "github.com/ZenLiuCN/engine/sqlx"
	_ "github.com/ZenLiuCN/engine/sqlx/duckdb"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	. "github.com/urfave/cli/v2"
	"os"
	"sync"
	"time"
)

func main() {
	err := (&App{
		Name:        "Engine",
		Description: "ESM engine with golang extensions",
		Usage:       "engine [flags] <script> [args ...]",
		Version:     "0.1.0",
		Flags: []Flag{
			&BoolFlag{Name: "define", Aliases: []string{"d"}, Usage: "Dump extension definitions (.d.ts)\n Eg: engine -d <PATH TO STORE FILES>"},
			&BoolFlag{Name: "typescript", Aliases: []string{"t"}, Usage: "typescript mode, script file will auto detect by extension."},
			&BoolFlag{Name: "source", Aliases: []string{"s"}, Usage: "execute source from commandline, this mode not support commandline args for script."},
			&DurationFlag{Name: "warmup", Aliases: []string{"w"}, Usage: "warmup time before time limit", DefaultText: "1s", Value: time.Second},
			&DurationFlag{Name: "timeout", Aliases: []string{"i"}, Usage: "limit execute time"},
		},
		Action: func(c *Context) error {
			if c.Bool("define") {
				p := c.Args().First()
				if p == "" {
					p = "."
				}
				engine.DumpDefines(p)
				return nil
			}
			ts := c.Bool("typescript")
			warm := c.Duration("warmup")
			timeout := c.Duration("timeout")
			withTimeout := timeout == 0
			if c.Bool("source") {
				return executeStdIn(c, ts, warm, timeout, withTimeout)
			}
			args := c.Args().Slice()
			if len(args) == 0 {
				return fmt.Errorf("missing script file, use engine -h to show helps")
			}
			vm := engine.Get()
			defer vm.Free()
			wg := &sync.WaitGroup{}
			wg.Add(1)
			cc := fn.WithSignal(func(ctx context.Context) {
				defer wg.Done()
				if len(args) > 1 {
					vm.Set("args", args[1:])
				} else {
					vm.Set("args", []string{})
				}
				var v goja.Value
				if withTimeout {
					cx, cc := context.WithTimeout(ctx, timeout)
					defer cc()
					v = fn.Panic1(vm.RunCodeContext(engine.CompileFile(args[0], true), warm, cx))
				} else {
					v = fn.Panic1(vm.RunCodeContext(engine.CompileFile(args[0], true), warm, ctx))
				}
				if !engine.IsNullish(v) {
					fmt.Printf("%v\n", v.Export())
				}
			})
			wg.Wait()
			defer cc()
			return nil
		},
		Authors: []*Author{{
			Name: "ZenLiu",
		}},
		UseShortOptionHandling: true,
		Suggest:                true,
		EnableBashCompletion:   true,
	}).Run(os.Args)
	if err != nil {
		println(err.Error())
	}
}
func executeStdIn(c *Context, ts bool, warm time.Duration, timeout time.Duration, withTimeout bool) error {
	if !c.Args().Present() {
		return fmt.Errorf("missing script source, use engine -h to show helps")
	}
	vm := engine.Get()
	defer vm.Free()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	cc := fn.WithSignal(func(ctx context.Context) {
		defer wg.Done()
		var v goja.Value
		if withTimeout {
			cx, cc := context.WithTimeout(ctx, timeout)
			defer cc()
			v = fn.Panic1(vm.RunCodeContext(engine.CompileSource(fn.SliceJoinRune(c.Args().Slice(), '\n', fn.Identity[string]), ts, true), warm, cx))
		} else {
			v = fn.Panic1(vm.RunCodeContext(engine.CompileSource(fn.SliceJoinRune(c.Args().Slice(), '\n', fn.Identity[string]), ts, true), warm, ctx))
		}
		if !engine.IsNullish(v) {
			fmt.Printf("%v\n", v.Export())
		}
	})
	wg.Wait()
	defer cc()
	return nil
}
