// Code generated by define_gene; DO NOT EDIT.
package errors

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"errors"
)

var (
	//go:embed errors.d.ts
	ErrorsDefine   []byte
	ErrorsDeclared = map[string]any{
		"unwrap":         errors.Unwrap,
		"as":             errors.As,
		"ErrUnsupported": errors.ErrUnsupported,
		"is":             errors.Is,
		"join":           errors.Join,
		"New":            errors.New,
	}
)

func init() {
	engine.RegisterModule(ErrorsModule{})
}

type ErrorsModule struct{}

func (S ErrorsModule) Identity() string {
	return "golang/errors"
}
func (S ErrorsModule) TypeDefine() []byte {
	return ErrorsDefine
}
func (S ErrorsModule) Exports() map[string]any {
	return ErrorsDeclared
}