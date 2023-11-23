package engine

import (
	"github.com/ZenLiuCN/fn"
	. "github.com/dop251/goja"
)

type Engine struct {
	*Runtime
}

// Register register modules
func (s *Engine) Register(modules ...Module) {
	for _, module := range modules {
		if im, ok := module.(InitializeModule); ok {
			m := im.Initialize(s)
			fn.Panic(s.Runtime.Set(m.Name(), m))
		} else if tm, ok := module.(TopModule); ok {
			tm.Register(s)
		} else {
			fn.Panic(s.Runtime.Set(module.Name(), module))
		}

	}
}

// Free recycle this engine
func (s *Engine) Free() {

}

//region Override

// RunTypeScript execute typescript code
func (s *Engine) RunTypeScript(src string) (Value, error) {
	return s.Runtime.RunString(CompileTs(src))
}

// RunJavaScript execute javascript code
func (s *Engine) RunJavaScript(src string) (Value, error) {
	return s.Runtime.RunString(CompileJs(src))
}

// RunScript execute javascript (es5|es6 without import) code
func (s *Engine) RunScript(src string) (Value, error) {
	return s.Runtime.RunString(src)
}

// Execute execute compiled code
func (s *Engine) Execute(prog *Code) (Value, error) {
	return s.Runtime.RunProgram(prog.Program)
}

//endregion

func NewEngine(modules ...Module) (r *Engine) {
	r = &Engine{Runtime: New()}
	r.Runtime.SetFieldNameMapper(EngineFieldMapper{})
	r.Register(Modules()...)
	r.Register(modules...)
	return
}
