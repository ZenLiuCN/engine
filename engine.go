package engine

import (
	"fmt"
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

// CallFunction invoke a function (without this)
func (s *Engine) CallFunction(fn Value, values ...any) (Value, error) {
	if f, ok := AssertFunction(fn); ok {
		v := make([]Value, len(values))
		for i, value := range values {
			v[i] = s.ToValue(value)
		}
		return f(Undefined(), v...)
	}
	return nil, fmt.Errorf("%s not a function", fn)
}

// CallMethod invoke a method (with this)
func (s *Engine) CallMethod(fn Value, self Value, values ...any) (Value, error) {
	if f, ok := AssertFunction(fn); ok {
		v := make([]Value, len(values))
		for i, value := range values {
			v[i] = s.ToValue(value)
		}
		return f(self, v...)
	}
	return nil, fmt.Errorf("%s not a function", fn)
}

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
func (s *Engine) Execute(code *Code) (Value, error) {
	return s.Runtime.RunProgram(code.Program)
}

func NewEngine(modules ...Module) (r *Engine) {
	r = &Engine{Runtime: New()}
	r.Runtime.SetFieldNameMapper(EngineFieldMapper{})
	r.Register(Modules()...)
	r.Register(modules...)
	return
}
