package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"math/rand"
	"strings"
)

var (
	ColumnMarker = "♦"
)

func (s *Engine) useMapping(b []byte) string {
	rnd := fmt.Sprintf("%d.map", rand.Int())
	s.Runtime.SetParserOptions(parser.WithSourceMapLoader(func(path string) ([]byte, error) {
		if path == rnd {
			return b, nil
		}
		return nil, nil
	}))
	return "\n//# sourceMappingURL=" + rnd
}
func (s *Engine) freeMapping() {
	s.SourceMap = nil
	s.Runtime.SetParserOptions([]parser.Option{}...)

}

type SourceMapping map[string]*SourceMap

func (s SourceMapping) register(src *SourceMap) {
	for _, source := range src.Sources {
		s[source] = src
	}
}
func (s SourceMapping) dump(b *bytes.Buffer, stacks []goja.StackFrame) {
	for _, stack := range stacks {
		loc := stack.Position()
		f := stack.SrcName()
		if f != "<native>" {
			if o, ok := s[f]; ok {
				if src := o.Code(f, loc.Line-1, loc.Column); src != "" {
					b.WriteString("\n")
					b.WriteString(f)
					b.WriteString(fmt.Sprintf("[%d:%d]", loc.Line, loc.Column))
					b.WriteString("\t")
					b.WriteString(src)
					continue
				}
			}
		}
		b.WriteByte('\n')
		stack.Write(b)
	}
}

func (s SourceMapping) one() *SourceMap {
	for _, sourceMap := range s {
		return sourceMap
	}
	return nil
}

type SourceMap struct {
	File           string   `json:"file"`
	SourceRoot     string   `json:"sourceRoot"`
	Sources        []string `json:"sources"`
	SourcesContent []string `json:"sourcesContent"`
	buf            [][][]rune
}

func (s *SourceMap) Code(src string, line, col int) string {
	n := -1
	for i, source := range s.Sources {
		if source == src {
			n = i
			break
		}
	}
	if n < 0 {
		return ""
	}
	if s.buf == nil {
		s.buf = make([][][]rune, len(s.Sources))
	}
	if s.buf[n] == nil {
		for _, s2 := range strings.Split(s.SourcesContent[n], "\n") {
			s.buf[n] = append(s.buf[n], []rune(s2))
		}
	}
	b := s.buf[n]
	if len(b) < line {
		return ""
	}
	l := b[line]
	if len(l) < col {
		return string(l) + ColumnMarker
	}
	return string(l[:col]) + ColumnMarker + string(l[col:])
}
func NewSourceMap(bin []byte) (v SourceMapping) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("error on generate source mapping :%#+v", r)
			v = nil
		}
	}()
	if len(bin) == 0 {
		return
	}
	var src = new(SourceMap)
	fn.Panic(json.Unmarshal(bin, src))
	v = make(SourceMapping)
	v.register(src)
	return
}
