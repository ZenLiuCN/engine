package main

import (
	"bytes"
	"fmt"
	"github.com/ZenLiuCN/fn"
)

type SpecWriter struct {
	Package string
	Imports fn.HashSet[string]
	buf     *bytes.Buffer
}

func WriterOf(buf *bytes.Buffer) *SpecWriter {
	return &SpecWriter{buf: buf}
}
func NewWriter() *SpecWriter {
	return &SpecWriter{buf: new(bytes.Buffer)}
}
func (s *SpecWriter) Buffer() *bytes.Buffer {
	return s.buf
}
func (s *SpecWriter) F(format string, args ...any) *SpecWriter {
	if s.buf == nil {
		s.buf = new(bytes.Buffer)
	}
	_, _ = fmt.Fprintf(s.buf, format, args...)
	return s
}
func (s *SpecWriter) Append(w *SpecWriter) *SpecWriter {
	if s.Imports == nil {
		s.Imports = fn.NewHashSet[string]()
	}
	for s2 := range w.Imports {
		s.Imports.Put(s2)
	}
	s.buf.Write(w.buf.Bytes())
	return s
}

func (s *SpecWriter) Import(path string) *SpecWriter {
	if s.Imports == nil {
		s.Imports = fn.NewHashSet[string]()
	}
	s.Imports.Put(path)
	return s
}
