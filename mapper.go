package engine

import (
	"reflect"
	"strings"
	"unicode"
)

var (
	reservedFieldName = map[string]string{
		"OCSP": "ocsp",
	}

	reservedMethodNames = map[string]string{
		"JSON": "json",
		"HTML": "html",
		"URL":  "url",
		"OCSP": "ocsp",
	}
)

func FieldName(_ reflect.Type, f reflect.StructField) string {
	if f.PkgPath != "" {
		return ""
	}
	if tag := f.Tag.Get("js"); tag != "" {
		if tag == "-" {
			return ""
		}
		return tag
	}
	if exception, ok := reservedFieldName[f.Name]; ok {
		return exception
	}
	return toLowerCase(f.Name)
}

// MethodName Returns the JS name for an exported method.
func MethodName(_ reflect.Type, m reflect.Method) string {
	if m.Name[0] == 'X' {
		return m.Name[1:]
	}
	if exception, ok := reservedMethodNames[m.Name]; ok {
		return exception
	}
	return toLowerCase(m.Name)
}

type EngineFieldMapper struct{}

// https://godoc.org/github.com/dop251/goja#FieldNameMapper

func (EngineFieldMapper) FieldName(t reflect.Type, f reflect.StructField) string {
	return FieldName(t, f)
}

// https://godoc.org/github.com/dop251/goja#FieldNameMapper

func (EngineFieldMapper) MethodName(t reflect.Type, m reflect.Method) string { return MethodName(t, m) }

func toLowerCase(name string) string {
	runes := []rune(name)
	n := 0
	for i := 0; i < len(runes); i++ {
		if !unicode.IsUpper(runes[i]) {
			if i > 1 {
				n = i - 1
			} else {
				n = i
			}
			break
		}
	}
	out := strings.ToLower(string(runes[:n])) + string(runes[n:])
	return out
}
