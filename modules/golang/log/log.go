// Code generated by define_gene; DO NOT EDIT.
package log

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	"log"
)

var (
	//go:embed log.d.ts
	LogDefine   []byte
	LogDeclared = map[string]any{
		"print":         log.Print,
		"setFlags":      log.SetFlags,
		"fatalln":       log.Fatalln,
		"Ldate":         log.Ldate,
		"Llongfile":     log.Llongfile,
		"Lmsgprefix":    log.Lmsgprefix,
		"writer":        log.Writer,
		"flags":         log.Flags,
		"Lshortfile":    log.Lshortfile,
		"Ltime":         log.Ltime,
		"printf":        log.Printf,
		"fatalf":        log.Fatalf,
		"panic":         log.Panic,
		"setPrefix":     log.SetPrefix,
		"Default":       log.Default,
		"prefix":        log.Prefix,
		"LstdFlags":     log.LstdFlags,
		"output":        log.Output,
		"panicf":        log.Panicf,
		"panicln":       log.Panicln,
		"LUTC":          log.LUTC,
		"setOutput":     log.SetOutput,
		"fatal":         log.Fatal,
		"Lmicroseconds": log.Lmicroseconds,
		"New":           log.New,
		"println":       log.Println,

		"emptyLogger":    engine.Empty[log.Logger],
		"emptyRefLogger": engine.EmptyRefer[log.Logger],
		"refOfLogger":    engine.ReferOf[log.Logger],
		"unRefLogger":    engine.UnRefer[log.Logger]}
)

func init() {
	engine.RegisterModule(LogModule{})
}

type LogModule struct{}

func (S LogModule) Identity() string {
	return "golang/log"
}
func (S LogModule) TypeDefine() []byte {
	return LogDefine
}
func (S LogModule) Exports() map[string]any {
	return LogDeclared
}