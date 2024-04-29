package main

import (
	"bytes"
	"fmt"
	"github.com/ZenLiuCN/fn"
)

type SpecWriter struct {
	Package string
	Imports fn.HashSet[string]
	*bytes.Buffer
}

func WriterOf(buf *bytes.Buffer) *SpecWriter {
	return &SpecWriter{Buffer: buf}
}
func NewWriter() *SpecWriter {
	return &SpecWriter{Buffer: new(bytes.Buffer)}
}

func (s *SpecWriter) F(format string, args ...any) *SpecWriter {
	if s.Buffer == nil {
		s.Buffer = new(bytes.Buffer)
	}
	_, _ = fmt.Fprintf(s.Buffer, format, args...)
	return s
}
func (s *SpecWriter) Append(w *SpecWriter) *SpecWriter {
	if s.Imports == nil {
		s.Imports = fn.NewHashSet[string]()
	}
	for s2 := range w.Imports {
		s.Imports.Put(s2)
	}
	s.Buffer.Write(w.Buffer.Bytes())
	return s
}

func (s *SpecWriter) Import(path string) *SpecWriter {
	if s.Imports == nil {
		s.Imports = fn.NewHashSet[string]()
	}
	s.Imports.Put(path)
	return s
}
