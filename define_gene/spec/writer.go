package spec

import (
	"bytes"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"io"
)

type (
	Writer interface {
		Imports() fn.HashSet[string]
		BytesBuffer() *bytes.Buffer

		WriteGoGenerated(extra string) Writer
		F(format string, args ...any) Writer
		Append(w Writer) Writer
		Import(path string) Writer
		MergeImports(path fn.HashSet[string]) Writer

		Bytes() []byte
		AvailableBuffer() []byte
		Len() int
		Cap() int
		Available() int
		Truncate(n int)
		Reset()
		Grow(n int)
		Next(n int) []byte
		ReadBytes(delim byte) (line []byte, err error)
		ReadString(delim byte) (line string, err error)

		fmt.Stringer
		io.Writer
		io.StringWriter
		io.ReadWriter
		io.WriterTo
		io.RuneScanner
		io.ByteScanner
		io.ReaderFrom
	}
	SpecWriter struct {
		imports fn.HashSet[string]
		*bytes.Buffer
	}
)

func (s *SpecWriter) MergeImports(v fn.HashSet[string]) Writer {
	if s.imports == nil {
		s.imports = fn.NewHashSet[string]()
	}
	for s2, _ := range v {
		s.imports.Add(s2)
	}
	return s
}
func (s *SpecWriter) Imports() fn.HashSet[string] {
	if s.imports == nil {
		s.imports = fn.NewHashSet[string]()
	}
	return s.imports
}
func (s *SpecWriter) Reset() {
	s.Buffer.Reset()
	s.imports.Clear()
}
func (s *SpecWriter) BytesBuffer() *bytes.Buffer {
	return s.Buffer
}
func (s *SpecWriter) WriteGoGenerated(extra string) Writer {
	s.Buffer.WriteString(fmt.Sprintf("\n// Code generated %s; DO NOT EDIT.\n", extra))
	return s
}
func (s *SpecWriter) F(format string, args ...any) Writer {
	if s.Buffer == nil {
		s.Buffer = new(bytes.Buffer)
	}
	_, _ = fmt.Fprintf(s.Buffer, format, args...)
	return s
}
func (s *SpecWriter) Append(w Writer) Writer {
	if s.imports == nil {
		s.imports = fn.NewHashSet[string]()
	}
	for s2 := range w.Imports() {
		s.imports.Add(s2)
	}
	s.Buffer.Write(w.Bytes())
	return s
}
func (s *SpecWriter) Import(path string) Writer {
	if s.imports == nil {
		s.imports = fn.NewHashSet[string]()
	}
	s.imports.Add(path)
	return s
}
func WriterOf(buf *bytes.Buffer) *SpecWriter {
	return &SpecWriter{Buffer: buf}
}
func NewWriter() *SpecWriter {
	return &SpecWriter{Buffer: new(bytes.Buffer)}
}
