package engine

import (
	"bytes"
	"context"
	"fmt"
	"github.com/dop251/goja"
	"log/slog"
	"strings"
)

type Console struct {
	*slog.Logger
}

func NewConsole(logger *slog.Logger) *Console {
	return &Console{Logger: logger}
}

func (s Console) Name() string {
	return "console"
}
func (s Console) Assert(cond bool, args ...goja.Value) {
	if !cond {
		s.Error(args...)
		panic(fmt.Errorf("%#v", args))
	}
}
func (s Console) log(level slog.Level, args ...goja.Value) {
	var msg strings.Builder
	for i := 0; i < len(args); i++ {
		if i > 0 {
			msg.WriteString(" ")
		}
		msg.WriteString(ValueString(args[i]))
	}
	s.Logger.Log(context.Background(), level, msg.String())
}

func (s Console) Log(args ...goja.Value) {
	s.Info(args...)
}

func (s Console) Debug(args ...goja.Value) {
	s.log(slog.LevelDebug, args...)
}

func (s Console) Info(args ...goja.Value) {
	s.log(slog.LevelInfo, args...)
}

func (s Console) Warn(args ...goja.Value) {
	s.log(slog.LevelWarn, args...)
}

func (s Console) Error(args ...goja.Value) {
	s.log(slog.LevelError, args...)
}

type BufferConsole struct {
	*bytes.Buffer
}

func NewBufferConsoleOf(buf *bytes.Buffer) *BufferConsole {
	return &BufferConsole{buf}
}
func NewBufferConsole() *BufferConsole {
	return &BufferConsole{GetBytesBuffer()}
}
func (s *BufferConsole) Name() string {
	return "console"
}
func (s *BufferConsole) log(level slog.Level, args ...goja.Value) {
	s.Buffer.WriteRune('[')
	s.Buffer.WriteString(level.String())
	s.Buffer.WriteRune(']')
	s.Buffer.WriteRune('\t')
	for i := 0; i < len(args); i++ {
		if i > 0 {
			s.Buffer.WriteRune(' ')
		}
		s.Buffer.WriteString(ValueString(args[i]))
	}
	s.Buffer.WriteRune('\n')

}
func (s *BufferConsole) Assert(cond bool, args ...goja.Value) {
	if !cond {
		s.Error(args...)
		panic(fmt.Errorf("%#v", args))
	}
}
func (s *BufferConsole) Log(args ...goja.Value) {
	s.Info(args...)
}

func (s *BufferConsole) Debug(args ...goja.Value) {
	s.log(slog.LevelDebug, args...)
}

func (s *BufferConsole) Info(args ...goja.Value) {
	s.log(slog.LevelInfo, args...)
}

func (s *BufferConsole) Warn(args ...goja.Value) {
	s.log(slog.LevelWarn, args...)
}

func (s *BufferConsole) Error(args ...goja.Value) {
	s.log(slog.LevelError, args...)
}
