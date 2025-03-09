package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"strings"
)

const (
	mark = "♦"
)

func (e *Engine) useMapping() {
	if e.SourceMap == nil {
		e.SourceMap = &SourceMapping{sources: make(map[string]*SourceMap), raw: make(map[string][]byte)}
		e.SourceMap.config(e)
	}
}
func (e *Engine) compileWithMapping(source string, entry bool) {

}
func (e *Engine) freeMapping() {
	if e.SourceMap != nil {
		e.SourceMap = nil
		e.Runtime.SetParserOptions(make([]parser.Option, 0)...)
	}
}

type SourceMapping struct {
	sources map[string]*SourceMap
	raw     map[string][]byte
}

func (s *SourceMapping) config(e *Engine) {
	//e.SourceMap = s
	e.Runtime.SetParserOptions(parser.WithSourceMapLoader(func(path string) ([]byte, error) {
		if b, ok := s.raw[path]; ok {
			return b, nil
		}
		return nil, nil
	}))
}

func (s *SourceMapping) dump(b *bytes.Buffer, stacks []goja.StackFrame) {
	for _, stack := range stacks {
		loc := stack.Position()
		f := stack.SrcName()
		if f != "<native>" {
			if o, ok := s.sources[f]; ok {
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

type SourceMap struct {
	Version        int      `json:"version"`
	File           string   `json:"file"`
	SourceRoot     string   `json:"sourceRoot"`
	Sources        []string `json:"sources"`
	SourcesContent []string `json:"sourcesContent"`
	buf            [][]string
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
		s.buf = make([][]string, len(s.Sources))
	}
	if s.buf[n] == nil {
		s.buf[n] = strings.Split(s.SourcesContent[n], "\n")
	}
	b := s.buf[n]
	if len(b) < line {
		return ""
	}
	l := b[line]
	if len(l) < col {
		return l + mark
	}
	return l[:col] + mark + l[col:]
}

/*
	func (s *SourceMap) sliceSource(si, line, col int) string {
		if s.buf == nil {
			s.buf = make(map[int][]string, len(s.Sources))
		}

read:

		if lines, ok := s.buf[si]; ok {
			if SingularLine {
				return dumpOneLine(lines, line, col)
			}
			return dumpTwoLines(lines, line, col)

		}
		var v []string
		sc := bufio.NewScanner(strings.NewReader(s.SourcesContent[si]))
		sc.Split(bufio.ScanLines)
		for sc.Scan() {
			v = append(v, sc.Text())
		}
		s.buf[si] = v
		goto read
	}

	func dumpOneLine(lines []string, line int, col int) string {
		n := len(lines)
		if line >= n {
			return ""
		} else if line == 0 {
			n := lines[0]
			b := new(strings.Builder)
			b.Grow(len(n) + 25)
			b.WriteString("0 →")
			b.WriteString(n[:col])
			b.WriteRune(ColumnMarker)
			b.WriteString(n[col:])
			return b.String()
		} else {
			ls := lines[line]
			b := new(strings.Builder)
			b.Grow(len(ls) + 25)
			b.WriteString(fmt.Sprintf("%d →", line))
			b.WriteString(ls[:col])
			b.WriteRune(ColumnMarker)
			b.WriteString(ls[col:])
			return b.String()
		}
	}

	func dumpTwoLines(lines []string, line int, col int) string {
		n := len(lines)
		if line >= n {
			return ""
		} else if line == 0 {
			n := lines[0]
			b := new(strings.Builder)
			b.Grow(len(n) + 25 + len(lines[1]))
			b.WriteString(fmt.Sprintf("%d →", 0))
			b.WriteString(n[:col])
			b.WriteRune(ColumnMarker)
			b.WriteString(n[col:])
			b.WriteRune('\n')
			b.WriteString(lines[1])
			return b.String()
		} else if line == n-1 {
			ls := lines[n-1]
			b := new(strings.Builder)
			b.Grow(len(ls) + 25 + len(lines[n-2]))
			b.WriteString(lines[n-2])
			b.WriteRune('\n')
			b.WriteString(fmt.Sprintf("%d →", 0))
			b.WriteString(ls[:col])
			b.WriteRune(ColumnMarker)
			b.WriteString(ls[col:])
			return b.String()
		} else {
			ls := lines[line-1 : line+2]
			x := 0
			for _, l := range ls {
				x += len(l)
			}
			x += 25
			b := new(strings.Builder)
			b.Grow(x)
			b.WriteString(ls[0])
			b.WriteRune('\n')
			b.WriteString(fmt.Sprintf("%d →", 0))
			b.WriteString(ls[1][:col])
			b.WriteRune('☆')
			b.WriteString(ls[1][col:])
			b.WriteRune('\n')
			b.WriteString(ls[2])
			return b.String()
		}
	}
*/
func NewSourceMap(bin []byte) (v *SourceMap) {
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
	return src
}

func decodeVLQ(encoded []byte) (value int, i int, ok bool) {
	n := len(encoded)
	if n == 0 {
		return
	}
	shift := 0
	var vlq int
	for {
		if i >= n {
			return 0, 0, false
		}
		index := strings.IndexByte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/", encoded[i])
		if index < 0 {
			return 0, 0, false
		}
		vlq |= (index & 31) << shift
		i++
		shift += 5
		if (index & 32) == 0 {
			break
		}
	}
	value = vlq >> 1
	if (vlq & 1) != 0 {
		value = -value
	}
	return value, i, true
}
