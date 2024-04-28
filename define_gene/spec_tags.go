package main

import (
	"go/ast"
	"strings"
)

type SpecTags map[string][]string

func NewTags(s string) SpecTags {
	if s == "" {
		return nil
	}
	t := make(SpecTags)
	for _, seg := range strings.Split(s, " ") {
		v := strings.Split(seg, ":")
		if len(v) == 1 {
			t[v[0]] = nil
		} else {
			val := strings.Split(v[1][1:len(v[1])-1], ",")
			for _, s2 := range val {
				ss := strings.TrimSpace(s2)
				if len(ss) > 0 {
					t[v[0]] = append(t[v[0]], ss)
				}
			}
		}
	}
	return t
}
func NewTagsOf(s *ast.BasicLit) SpecTags {
	if s == nil {
		return nil
	}
	return NewTags(s.Value[1 : len(s.Value)-1])
}
