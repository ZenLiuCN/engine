package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ZenLiuCN/engine/define_gene/spec"
	"github.com/ZenLiuCN/fn"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"os"
	"strings"
	"sync"
	"unicode"
)

//region Writer

type Writer struct {
	*bytes.Buffer
	imp fn.HashSet[string] //package Imports
	use fn.HashSet[string] //raw go type used
}

func (t *Writer) GoString() string {
	b := new(strings.Builder)
	_, _ = fmt.Fprintf(b, "\n\timports:[%s]\n", strings.Join(t.imp.Values(), ","))
	_, _ = fmt.Fprintf(b, "\n\tuses:[%s]\n", strings.Join(t.use.Values(), ","))
	_, _ = fmt.Fprintf(b, "\n\t%s\n", t.Buffer.String())
	return b.String()
}

// Format format string
func (t *Writer) Format(format string, args ...any) *Writer {
	_, _ = fmt.Fprintf(t.Buffer, format, args...)
	return t
}

// Imports go packages (full Name for golang std )
func (t *Writer) Imports(pkg ...string) *Writer {
	for _, s := range pkg {
		t.imp.Add(s)
	}
	return t
}
func (t *Writer) Uses(typ ...string) *Writer {
	for _, s := range typ {
		t.use.Add(s)
	}
	return t
}

// Clear all
func (t *Writer) Clear() {
	t.Buffer.Reset()
	t.imp.Clear()
	t.use.Clear()
}
func (t *Writer) MergeWithoutContent(ts *Writer) *Writer {
	for s := range ts.imp {
		t.imp.Add(s)
	}
	for s := range ts.use {
		t.use.Add(s)
	}
	return t
}
func (t *Writer) Append(ts *Writer) *Writer {
	t.MergeWithoutContent(ts)
	t.Buffer.Write(ts.Bytes())
	return t
}
func (t *Writer) Prepend(ts *Writer) *Writer {
	t.MergeWithoutContent(ts)
	_, _ = t.Buffer.WriteTo(ts.Buffer)
	tx := t.Buffer
	t.Buffer = ts.Buffer
	ts.Buffer = tx
	return t
}

// Free this writer
func (t *Writer) Free() {
	t.Clear()
	pool.Put(t)
}

// Type write Ident type as format of %s
func (t *Writer) Type(format string, s string, array bool) {
	switch s {
	case "byte", "uint8":
		if array {
			t.Format(format, "Uint8Array")
		} else {
			t.Uses(s)
			t.Format(format, s)
		}
	case "uint16", "uint32":
		if array {
			t.Format(format, fmt.Sprintf("U%sArray", s[1:]))
		} else {
			t.Uses(s)
			t.Format(format, s)
		}
	case "int8", "int16", "int32":
		if array {
			t.Format(format, fmt.Sprintf("I%sArray", s[1:]))
		} else {
			t.Uses(s)
			t.Format(format, s)
		}
	case "error":
		if array {
			t.Uses(s)
			t.Format(format, "error[]")
		} else {
			t.Uses(s)
			t.Format(format, "error")
		}
	case "string":
		if array {
			t.Format(format, "string[]")
		} else {
			t.Format(format, "string")
		}
	case "any":
		if array {
			t.Format(format, "any[]")
		} else {
			t.Format(format, "any")
		}
	default:
		if array {
			if !unicode.IsUpper(rune(s[0])) {
				t.Uses(s)
			}
			t.Format(format, fmt.Sprintf("%s[]", s))
		} else {
			if !unicode.IsUpper(rune(s[0])) {
				t.Uses(s)
			}
			t.Format(format, s)
		}

	}
}

// Ident write golang Ident to a Ts Ident (low first)
func (t *Writer) Ident(format string, s string) *Writer {
	if s == "New" {
		t.Format(format, s)
		return t
	}
	buf := new(bytes.Buffer)
	u0 := false
	u := false
	rs := []rune(s)
	mx := len(rs) - 1
	for i := 0; i < len(rs); i++ {
		r := rs[i]
		switch {
		case i == 0 && unicode.IsUpper(r):
			buf.WriteRune(unicode.ToLower(r))
			u0 = true
			u = true
		case u0 && u && unicode.IsUpper(r):
			if i < mx && unicode.IsLower(rs[i+1]) {
				buf.WriteRune(r)
				u = false
			} else {
				buf.WriteRune(unicode.ToLower(r))
			}
		case unicode.IsUpper(r):
			u = true
			buf.WriteRune(r)
		case unicode.IsLower(r):
			u = false
			buf.WriteRune(r)
		default:
			buf.WriteRune(r)
		}
	}
	t.Format(format, buf.String())
	return t
}

var (
	pool = sync.Pool{New: func() any {
		return &Writer{
			Buffer: new(bytes.Buffer),
			imp:    fn.NewHashSet[string](),
			use:    fn.NewHashSet[string](),
		}
	}}
)

func get() *Writer {
	return pool.Get().(*Writer)
}

// endregion
type Flags []FLAG

func (f Flags) First() FLAG {
	return f[0]
}
func (f Flags) Last() FLAG {
	return f[len(f)-1]
}
func (f Flags) LastN(v int) FLAG {
	if len(f) > v {
		return f[len(f)-v]
	}
	return 0
}
func (f Flags) IsStruct() bool {
	return f.Last().IsStruct()
}
func (f Flags) IsInterface() bool {
	return f.Last().IsInterface()
}
func (f Flags) IsFunction() bool {
	return f.Last().IsFunction()
}
func (f Flags) IsArgument() bool {
	return f.Last().IsArgument()
}
func (f Flags) IsResult() bool {
	return f.Last().IsResult()
}
func (f Flags) IsMethod() bool {
	return f.Last().IsMethod()
}
func (f Flags) IsField() bool {
	return f.Last().IsField()
}
func (f Flags) IsMapKey() bool {
	return f.Last().IsMapKey()
}
func (f Flags) IsMapValue() bool {
	return f.Last().IsMapValue()
}
func (f Flags) IsAnonymous() bool {
	return f.Last().IsAnonymous()
}
func (f Flags) IsNamed() bool {
	return f.Last().IsNamed()
}
func (f Flags) IsNamedMany() bool {
	return f.Last().IsNamedMany()
}
func (f Flags) WithLast(flag FLAG) Flags {
	f[len(f)-1] = f.Last() | flag
	return f
}

type FLAG int

func (f FLAG) GoString() string {
	b := new(bytes.Buffer)
	if f.IsStruct() {
		b.WriteString("STRUCT|")
	}
	if f.IsInterface() {
		b.WriteString("INTERFACE|")
	}
	if f.IsFunction() {
		b.WriteString("FUNCTION|")
	}
	if f.IsMethod() {
		b.WriteString("METHOD|")
	}
	if f.IsResult() {
		b.WriteString("FIELD|")
	}
	if f.IsArgument() {
		b.WriteString("ARGUMENTS|")
	}
	if f.IsResult() {
		b.WriteString("RESULT|")
	}
	if f.IsMapKey() {
		b.WriteString("MAP_KEY|")
	}
	if f.IsMapValue() {
		b.WriteString("MAP_VALUE|")
	}
	return b.String()
}
func (f FLAG) IsStruct() bool {
	return f&FStruct > 0
}
func (f FLAG) IsInterface() bool {
	return f&FInterface > 0
}
func (f FLAG) IsFunction() bool {
	return f&FFunction > 0
}
func (f FLAG) IsArgument() bool {
	return f&FArgument > 0
}
func (f FLAG) IsResult() bool {
	return f&FResult > 0
}
func (f FLAG) IsMethod() bool {
	return f&FMethod > 0
}
func (f FLAG) IsField() bool {
	return f&FField > 0
}

func (f FLAG) IsMapKey() bool {
	return f&FMapKey > 0
}
func (f FLAG) IsMapValue() bool {
	return f&FMapValue > 0
}
func (f FLAG) IsAnonymous() bool {
	return f&FAnonymous > 0
}
func (f FLAG) IsNamed() bool {
	return f&FNamed > 0
}
func (f FLAG) IsNamedMany() bool {
	return f&FNamedMany > 0
}

const (
	FNothing FLAG = 0
	FStruct  FLAG = 1 << (iota - 1)
	FInterface
	FFunction
	FArgument
	FResult
	FMethod
	FField
	FMapKey
	FMapValue
	FNamed
	FNamedMany
	FAnonymous
)

func Exists(p string) bool {
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil {
		panic(err)
	}
	return true
}

type TypeRegistry struct {
	pkg       *packages.Package
	typeFunc  spec.Providers
	typeStack Types
	typePoint map[ast.Expr]Types
}

func (w *TypeRegistry) lookupType(expr ast.Expr) Types {
	if w.typePoint == nil {
		w.typePoint = make(map[ast.Expr]Types)
		w.typeFunc = make(spec.Providers)
		w.typeFunc.WithAllType(
			func(t *ast.Ident) {
				w.typeStack = append(w.typeStack, TypeInfo{
					Node:  spec.AstIdent,
					Ident: t,
				})
			},
			func(t *ast.SelectorExpr, target types.Object) {
				w.typeStack = append(w.typeStack, TypeInfo{
					Node:           spec.AstSelectorExpr,
					Selector:       t,
					SelectorTarget: target,
				})
			},
			func(t *ast.StarExpr) {
				w.typeStack = append(w.typeStack, TypeInfo{
					Node: spec.AstStarExpr,
					Star: t,
				})
			},
			func(t *ast.FuncType) {
				w.typeStack = append(w.typeStack, TypeInfo{
					Node: spec.AstFuncType,
					Func: t,
				})
			},
			func(t *ast.ArrayType) {
				w.typeStack = append(w.typeStack, TypeInfo{
					Node:  spec.AstArrayType,
					Array: t,
				})
			},
			func(t *ast.MapType) {
				w.typeStack = append(w.typeStack, TypeInfo{
					Node: spec.AstMapType,
					Map:  t,
				})
			},
			func(t *ast.StructType) {
				w.typeStack = append(w.typeStack, TypeInfo{
					Node:   spec.AstStructType,
					Struct: t,
				})
			},
			func(t *ast.ChanType) {
				w.typeStack = append(w.typeStack, TypeInfo{
					Node: spec.AstChanType,
					Chan: t,
				})
			},
			func(t *ast.Ellipsis) {
				w.typeStack = append(w.typeStack, TypeInfo{
					Node:     spec.AstEllipsis,
					Ellipsis: t,
				})
			},
			func(t *ast.InterfaceType) {
				w.typeStack = append(w.typeStack, TypeInfo{
					Node:      spec.AstInterfaceType,
					Interface: t,
				})
			},
		)
	}
	if d, ok := w.typePoint[expr]; ok {
		return d
	}
	cur := len(w.typeStack)
	spec.CaseTypeFunc(w.pkg, expr, w.typeFunc.Provide)
	w.typePoint[expr] = w.typeStack[cur:]
	return w.typePoint[expr]
}
