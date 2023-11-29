package engine

import (
	"fmt"
	"strings"
	"testing"
)

type builder struct {
	b strings.Builder
}

func (b *builder) f(format string, args ...any) *builder {
	b.b.WriteString(fmt.Sprintf(format, args...))
	return b
}
func (b *builder) r(r rune) *builder {
	b.b.WriteRune(r)
	return b
}

func (b *builder) s(r string) *builder {
	b.b.WriteString(r)
	return b
}
func (b *builder) o(cond bool, r rune) *builder {
	if cond {
		b.b.WriteRune(r)
	}
	return b
}
func (b *builder) t(cond bool, r string) *builder {
	if cond {
		b.b.WriteString(r)
	}
	return b
}
func (b *builder) q(n int, format string, args ...any) *builder {
	for i := 0; i < n; i++ {
		if i > 0 {
			b.r(',')
		}
		b.f(format, append([]any{i}, args...)...)
	}
	return b
}
func (b *builder) Reset() {
	b.b.Reset()
}
func (b *builder) String() string {
	return b.b.String()
}
func TestTransformGen(t *testing.T) {
	b := &builder{strings.Builder{}}

	for i := 0; i < 5; i++ {
		for j := 1; j < 5; j++ {
			b.f("func TransformErr%d%d", i, j).o(i+j > 0, '[').q(i, "A%d").o(i > 0, ',').q(j, "B%d").o(j > 0, ',').t(i+j > 0, " C any]").
				s("(f func(").q(i, "A%d").s(")(").q(j, "B%d").o(j > 0, ',').s("error),t func(").q(j, "B%d").o(j > 0, ',').s("error)(C,error))func(").
				q(i, "A%d").s(`)(C,*GoError){
		return func(`).q(i, "a%[1]d A%[1]d").o(i > 0, ',').s(`)(C,*GoError){
		`).s("c,err:=t(f(").q(i, "a%[1]d").s(`))
		if err!=nil{
			return c,&GoError{Error: err}
		}
		return c,nil
	}
}
`)
			println(b.String())
			b.Reset()
			b.f("func Transform%d%d", i, j).o(i+j > 0, '[').q(i, "A%d").o(i > 0, ',').q(j, "B%d").o(j > 0, ',').t(i+j > 0, " C any]").
				s("(f func(").q(i, "A%d").s(")(").q(j, "B%d").s("),t func(").q(j, "B%d").s(")C)func(").q(i, "A%d").s(`)C{
		return func(`).q(i, "a%[1]d A%[1]d").s(`)C{
		return t(f(`).q(i, "a%[1]d").s(`))
	}
}
`)
			println(b.String())
			b.Reset()
		}
	}
}
