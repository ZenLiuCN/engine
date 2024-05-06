package engine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"strings"
)

type SourceMapping map[Location]Sources
type Location [2]int
type Sources struct {
	Location
	Source  string
	Content string
}
type SourceMap struct {
	Version        int      `json:"version"`
	File           string   `json:"file"`
	SourceRoot     string   `json:"sourceRoot"`
	Sources        []string `json:"sources"`
	SourcesContent []string `json:"sourcesContent"`
	Names          []string `json:"names"`
	Mappings       string   `json:"mappings"`
	buf            map[int][]string
}

func (s *SourceMap) sliceSource(si, line, col int) string {
	if s.buf == nil {
		s.buf = make(map[int][]string, len(s.Sources))
	}
read:
	if lines, ok := s.buf[si]; ok {
		n := len(lines)
		if line >= n {
			return ""
		} else if line == 0 {
			n := lines[0]
			b := new(strings.Builder)
			b.Grow(len(n) + 25 + len(lines[1]))
			b.WriteString(fmt.Sprintf("%d →", 0))
			b.WriteString(n[:col])
			b.WriteRune('☆')
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
			b.WriteRune('☆')
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
	var v []string
	sc := bufio.NewScanner(strings.NewReader(s.SourcesContent[si]))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		v = append(v, sc.Text())
	}
	s.buf[si] = v
	goto read
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
	var src SourceMap
	fn.Panic(json.Unmarshal(bin, &src))
	m := []byte(src.Mappings)
	n := len(m)
	ss := len(src.Sources)
	ns := len(src.Names)
	cur := 0
	var gl, gc, si, ol, oc, on, delta int
	var i int
	var ok bool
	read := func(name string) int {
		if delta, i, ok = decodeVLQ(m[cur:]); !ok {
			panic(fmt.Errorf("missing %s", name))
		}
		cur += i
		return delta
	}
	mayRead := func(name string) bool {
		if delta, i, ok = decodeVLQ(m[cur:]); ok {
			return true
		}
		cur += i
		return false
	}
	for cur < n {
		if m[cur] == ';' {
			gl++
			gc = 0
			cur++
			continue
		}
		read("generated column")
		gc += delta
		if gc < 0 {
			panic(fmt.Errorf("invalid generated column: %d", gc))
		}
		if cur == n {
			break
		}
		switch m[cur] {
		case ',':
			cur++
			continue
		case ';':
			continue

		}

		si += read("source index")
		if si < 0 || si >= ss {
			panic(fmt.Errorf("invalid source index: %d", si))
		}
		ol += read("original line")
		if ol < 0 {
			panic(fmt.Errorf("invalid soriginal line: %d", ol))
		}
		oc += read("original column")
		if oc < 0 {
			panic(fmt.Errorf("invalid original column: %d", oc))
		}
		if mayRead("original name") {
			on += delta
			if on < 0 || on >= ns {
				panic(fmt.Errorf("invalid original name: %d", on))
			}
		}
		if cur < n {
			switch m[cur] {
			case ',':
				cur++
			case ';':
				panic(fmt.Errorf("invalid mapping data: %d", cur))
			}
		}
		if v == nil {
			v = make(map[Location]Sources)
		}
		v[Location{gl, gc}] = Sources{
			Location: Location{ol, oc},
			Source:   src.Sources[si],
			Content:  (&src).sliceSource(si, ol, oc),
		}

	}
	return
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
