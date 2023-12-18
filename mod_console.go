package engine

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"github.com/dop251/goja"
	"log/slog"
	"slices"
	"strings"
	"time"
)

type console interface {
	Assert(cond bool, args ...goja.Value)
	Log(args ...goja.Value)
	Debug(args ...goja.Value)
	Info(args ...goja.Value)
	Warn(args ...goja.Value)
	Error(args ...goja.Value)
	Clear()
	Count(label string)
	CountReset(label string)
	Dir(v goja.Value)
	Time(label string)
	TimeEnd(label string)
	TimeLog(label string, args ...goja.Value)
	Group(label string)
	GroupCollapsed(label string)
	GroupEnd()
	Table(value goja.Value, columns []string)
}
type sharedConsole struct {
	counter map[string]int
	time    map[string]time.Time
	log     func(...any)
	group   []string
}

func newSharedConsole(log func(...any)) *sharedConsole {
	return &sharedConsole{
		log:     log,
		counter: map[string]int{},
		time:    map[string]time.Time{},
	}
}
func (s *sharedConsole) Count(label string) {
	var n int = 0
	if label == "" {
		label = "default"
	}
	n, _ = s.counter[label]
	n++
	s.counter[label] = n
	s.log(fmt.Sprintf("%s %d", label, n))
}
func (s *sharedConsole) CountReset(label string) {
	if label == "" {
		label = "default"
	}
	s.counter[label] = 0
}
func (s *sharedConsole) Dir(v goja.Value) {
	s.log(v)
}
func (s *sharedConsole) Time(label string) {
	if label == "" {
		label = "default"
	}
	s.time[label] = time.Now()
}
func (s *sharedConsole) TimeEnd(label string) {
	if label == "" {
		label = "default"
	}
	delete(s.time, label)
}
func (s *sharedConsole) TimeLog(label string, val ...goja.Value) {
	if label == "" {
		label = "default"
	}
	t, ok := s.time[label]
	if !ok {
		t = time.Now()
	}
	v := time.Now().Sub(t).String()
	s.log(append([]any{v}, toAny(val)...))
}
func (s *sharedConsole) Group(label string) {
	s.group = append(s.group, label)
}
func (s *sharedConsole) GroupCollapsed(label string) {
	s.group = append(s.group, label)
}
func (s *sharedConsole) GroupEnd() {
	s.group = s.group[:len(s.group)-1]
}
func (s *sharedConsole) Table(value goja.Value, columns []string) {
	s.log(dumpColumns(value, columns))
}
func dumpColumns(v goja.Value, col []string) string {
	m := strings.Builder{}
	m.WriteString("\n(index)")
	haveCol := len(col) > 0
	if len(col) > 0 {
		for _, column := range col {
			m.WriteString("\t")
			m.WriteString(column)
		}
		m.WriteString("\n")
	}
	switch t := v.Export().(type) {
	case []any:
		f := t[0]
		switch fv := f.(type) {
		case map[string]any:
			var cols []string
			if !haveCol {
				for s := range fv {
					cols = append(cols, s)
				}
				slices.Sort(cols)
				for _, s := range cols {
					m.WriteString("\t")
					m.WriteString(s)
				}
				m.WriteRune('\n')
			}
			for i, a := range t {
				m.WriteString(fmt.Sprintf("%d", i))
				mp := a.(map[string]any)
				if haveCol {
					for s, a2 := range mp {
						if haveCol && slices.Index(col, s) >= 0 {
							m.WriteString(fmt.Sprintf("\t%s", dumpOne(a2)))
						}
					}
				} else {
					for _, s := range cols {
						m.WriteString(fmt.Sprintf("\t%s", dumpOne(mp[s])))
					}
				}
				m.WriteRune('\n')
			}
		default:
			if !haveCol {
				m.WriteString("\tValue\n")
			}
			for i, a := range t {
				m.WriteString(fmt.Sprintf("%d\t", i))
				m.WriteString(fmt.Sprintf("%s", dumpOne(a)))
				m.WriteRune('\n')
			}
		}
	case map[string]any:
		var x any
		for _, a := range t {
			x = a
			break
		}
		switch f := x.(type) {
		case map[string]any:
			var cols []string
			if !haveCol {
				for s := range f {
					cols = append(cols, s)
				}
				slices.Sort(cols)
				for _, s := range cols {
					m.WriteString("\t")
					m.WriteString(s)
				}
				m.WriteRune('\n')
			}
			for i, a := range t {
				m.WriteString(fmt.Sprintf("%s", i))
				mp := a.(map[string]any)
				if haveCol {
					for s, a2 := range mp {
						if haveCol && slices.Index(col, s) >= 0 {
							m.WriteString(fmt.Sprintf("\t%s", dumpOne(a2)))
						}
					}
				} else {
					for _, s := range cols {
						m.WriteString(fmt.Sprintf("\t%s", dumpOne(mp[s])))
					}
				}
				m.WriteRune('\n')
			}
		default:
			if !haveCol {
				m.WriteString("\tValue\n")
			}
			for i, a := range t {
				m.WriteString(fmt.Sprintf("%s\t", i))
				m.WriteString(fmt.Sprintf("%s", dumpOne(a)))
				m.WriteRune('\n')
			}
		}
	default:
		return fmt.Sprintf("%v", t)
	}
	return m.String()
}
func dump(v ...any) string {
	var msg strings.Builder
	for i := 0; i < len(v); i++ {
		if i > 0 {
			msg.WriteString(" ")
		}
		if val, ok := v[i].(goja.Value); ok {
			msg.WriteString(ValueString(val))
		} else {
			msg.WriteString(fmt.Sprintf("%s", v[i]))
		}
	}
	return msg.String()
}
func dumpOne(v any) string {
	if val, ok := v.(goja.Value); ok {
		return ValueString(val)
	} else {
		return fmt.Sprintf("%v", v)
	}

}
func toAny[T any](v []T) []any {
	x := make([]any, len(v))
	for i, t := range v {
		x[i] = t
	}
	return x
}

type Console struct {
	*sharedConsole
	*slog.Logger
}

func (s *Console) TypeDefine() []byte {
	return nil
}
func (s *Console) Name() string {
	return "console"
}

func NewConsole(logger *slog.Logger) *Console {
	s := &Console{Logger: logger}
	s.sharedConsole = newSharedConsole(s.Print)
	return s
}
func (s *Console) Print(v ...any) {
	msg := dump(v...)
	s.Logger.Info(msg)
}
func (s *Console) log(level slog.Level, args ...goja.Value) {
	msg := dump(toAny(args)...)
	msg = strings.Repeat("\t", len(s.group)) + msg
	s.Logger.Log(context.Background(), level, msg)
}

func (s *Console) Assert(cond bool, args ...goja.Value) {
	if !cond {
		s.Error(args...)
		panic(fmt.Errorf("%#v", args))
	}
}

func (s *Console) Log(args ...goja.Value) {
	s.log(slog.LevelInfo, args...)
}
func (s *Console) Trace(args ...goja.Value) {
	s.log(slog.LevelDebug, args...)
}

func (s *Console) Debug(args ...goja.Value) {
	s.log(slog.LevelDebug, args...)
}

func (s *Console) Info(args ...goja.Value) {
	s.log(slog.LevelInfo, args...)
}

func (s *Console) Warn(args ...goja.Value) {
	s.log(slog.LevelWarn, args...)
}

func (s *Console) Error(args ...goja.Value) {
	s.log(slog.LevelError, args...)
}

func (s *Console) Clear() {

}

type BufferConsole struct {
	*sharedConsole
	*bytes.Buffer
}

func NewBufferConsoleOf(buf *bytes.Buffer) *BufferConsole {
	s := &BufferConsole{Buffer: buf}
	s.sharedConsole = newSharedConsole(s.Print)
	return s
}
func NewBufferConsole() *BufferConsole {
	s := &BufferConsole{Buffer: GetBytesBuffer()}
	s.sharedConsole = newSharedConsole(s.Print)
	return s
}
func (s *BufferConsole) Name() string {
	return "console"
}
func (s *BufferConsole) TypeDefine() []byte {
	return nil
}
func (s *BufferConsole) log(level string, args ...goja.Value) {
	if level != "" {
		s.Buffer.WriteRune('[')
		s.Buffer.WriteString(level)
		s.Buffer.WriteRune(']')
		s.Buffer.WriteRune('\t')
	}
	s.Buffer.WriteString(strings.Repeat("\t", len(s.group)))
	s.Buffer.WriteString(dump(toAny(args)...))
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
func (s *BufferConsole) Trace(args ...goja.Value) {
	s.log("Trace", args...)
}
func (s *BufferConsole) Debug(args ...goja.Value) {
	s.log("Debug", args...)
}

func (s *BufferConsole) Info(args ...goja.Value) {
	s.log("Info", args...)
}

func (s *BufferConsole) Warn(args ...goja.Value) {
	s.log("Warn", args...)
}

func (s *BufferConsole) Error(args ...goja.Value) {
	s.log("Error", args...)
}
func (s *BufferConsole) Clear() {
	s.Buffer.Reset()
}
func (s *BufferConsole) Print(v ...any) {
	s.Buffer.WriteString(dump(v...))
	s.Buffer.WriteRune('\n')
}
