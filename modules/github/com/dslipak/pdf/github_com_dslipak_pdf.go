// Code generated by define_gene; DO NOT EDIT.
package pdf

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/fmt"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	"github.com/dslipak/pdf"
)

var (
	//go:embed github_com_dslipak_pdf.d.ts
	GithubComDslipakPdfDefine   []byte
	GithubComDslipakPdfDeclared = map[string]any{
		"open":               pdf.Open,
		"Bool":               pdf.Bool,
		"Integer":            pdf.Integer,
		"newReaderEncrypted": pdf.NewReaderEncrypted,
		"Real":               pdf.Real,
		"Stream":             pdf.Stream,
		"interpret":          pdf.Interpret,
		"Name":               pdf.Name,
		"Null":               pdf.Null,
		"newReader":          pdf.NewReader,
		"String":             pdf.String,
		"Array":              pdf.Array,
		"Dict":               pdf.Dict,
		"ErrInvalidPassword": pdf.ErrInvalidPassword,

		"emptyValue": func() (v pdf.Value) {
			return v
		},
		"refValue": func() *pdf.Value {
			var x pdf.Value
			return &x
		},
		"refOfValue": func(x pdf.Value) *pdf.Value {
			return &x
		},
		"emptyFont": func() (v pdf.Font) {
			return v
		},
		"refFont": func() *pdf.Font {
			var x pdf.Font
			return &x
		},
		"refOfFont": func(x pdf.Font) *pdf.Font {
			return &x
		},
		"emptyPage": func() (v pdf.Page) {
			return v
		},
		"refPage": func() *pdf.Page {
			var x pdf.Page
			return &x
		},
		"refOfPage": func(x pdf.Page) *pdf.Page {
			return &x
		},
		"emptyPoint": func() (v pdf.Point) {
			return v
		},
		"refPoint": func() *pdf.Point {
			var x pdf.Point
			return &x
		},
		"refOfPoint": func(x pdf.Point) *pdf.Point {
			return &x
		},
		"emptyRow": func() (v pdf.Row) {
			return v
		},
		"refRow": func() *pdf.Row {
			var x pdf.Row
			return &x
		},
		"refOfRow": func(x pdf.Row) *pdf.Row {
			return &x
		},
		"emptyStack": func() (v pdf.Stack) {
			return v
		},
		"refStack": func() *pdf.Stack {
			var x pdf.Stack
			return &x
		},
		"refOfStack": func(x pdf.Stack) *pdf.Stack {
			return &x
		},
		"emptyText": func() (v pdf.Text) {
			return v
		},
		"refText": func() *pdf.Text {
			var x pdf.Text
			return &x
		},
		"refOfText": func(x pdf.Text) *pdf.Text {
			return &x
		},
		"emptyColumn": func() (v pdf.Column) {
			return v
		},
		"refColumn": func() *pdf.Column {
			var x pdf.Column
			return &x
		},
		"refOfColumn": func(x pdf.Column) *pdf.Column {
			return &x
		},
		"emptyContent": func() (v pdf.Content) {
			return v
		},
		"refContent": func() *pdf.Content {
			var x pdf.Content
			return &x
		},
		"refOfContent": func(x pdf.Content) *pdf.Content {
			return &x
		},
		"emptyOutline": func() (v pdf.Outline) {
			return v
		},
		"refOutline": func() *pdf.Outline {
			var x pdf.Outline
			return &x
		},
		"refOfOutline": func(x pdf.Outline) *pdf.Outline {
			return &x
		},
		"emptyReader": func() (v pdf.Reader) {
			return v
		},
		"refReader": func() *pdf.Reader {
			var x pdf.Reader
			return &x
		},
		"refOfReader": func(x pdf.Reader) *pdf.Reader {
			return &x
		},
		"emptyRect": func() (v pdf.Rect) {
			return v
		},
		"refRect": func() *pdf.Rect {
			var x pdf.Rect
			return &x
		},
		"refOfRect": func(x pdf.Rect) *pdf.Rect {
			return &x
		}}
)

func init() {
	engine.RegisterModule(GithubComDslipakPdfModule{})
}

type GithubComDslipakPdfModule struct{}

func (S GithubComDslipakPdfModule) Identity() string {
	return "github.com/dslipak/pdf"
}
func (S GithubComDslipakPdfModule) TypeDefine() []byte {
	return GithubComDslipakPdfDefine
}
func (S GithubComDslipakPdfModule) Exports() map[string]any {
	return GithubComDslipakPdfDeclared
}