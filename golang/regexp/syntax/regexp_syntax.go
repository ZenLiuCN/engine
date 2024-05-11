// Code generated by define_gene; DO NOT EDIT.
package syntax

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"regexp/syntax"
)

var (
	//go:embed regexp_syntax.d.ts
	RegexpSyntaxDefine   []byte
	RegexpSyntaxDeclared = map[string]any{
		"DotNL":                    syntax.DotNL,
		"InstRuneAnyNotNL":         syntax.InstRuneAnyNotNL,
		"parse":                    syntax.Parse,
		"EmptyNoWordBoundary":      syntax.EmptyNoWordBoundary,
		"ErrInvalidRepeatOp":       syntax.ErrInvalidRepeatOp,
		"ErrInvalidRepeatSize":     syntax.ErrInvalidRepeatSize,
		"InstMatch":                syntax.InstMatch,
		"InstRune1":                syntax.InstRune1,
		"OpAlternate":              syntax.OpAlternate,
		"OpWordBoundary":           syntax.OpWordBoundary,
		"EmptyBeginText":           syntax.EmptyBeginText,
		"ErrMissingRepeatArgument": syntax.ErrMissingRepeatArgument,
		"FoldCase":                 syntax.FoldCase,
		"Literal":                  syntax.Literal,
		"OpStar":                   syntax.OpStar,
		"WasDollar":                syntax.WasDollar,
		"UnicodeGroups":            syntax.UnicodeGroups,
		"EmptyEndText":             syntax.EmptyEndText,
		"InstFail":                 syntax.InstFail,
		"OpBeginText":              syntax.OpBeginText,
		"compile":                  syntax.Compile,
		"EmptyEndLine":             syntax.EmptyEndLine,
		"InstAlt":                  syntax.InstAlt,
		"OpQuest":                  syntax.OpQuest,
		"ErrInternalError":         syntax.ErrInternalError,
		"ErrInvalidCharClass":      syntax.ErrInvalidCharClass,
		"ErrInvalidPerlOp":         syntax.ErrInvalidPerlOp,
		"ErrInvalidUTF8":           syntax.ErrInvalidUTF8,
		"OneLine":                  syntax.OneLine,
		"OpCapture":                syntax.OpCapture,
		"OpEndText":                syntax.OpEndText,
		"InstCapture":              syntax.InstCapture,
		"Perl":                     syntax.Perl,
		"Simple":                   syntax.Simple,
		"ErrInvalidEscape":         syntax.ErrInvalidEscape,
		"ErrNestingDepth":          syntax.ErrNestingDepth,
		"ErrTrailingBackslash":     syntax.ErrTrailingBackslash,
		"OpBeginLine":              syntax.OpBeginLine,
		"ErrInvalidCharRange":      syntax.ErrInvalidCharRange,
		"OpCharClass":              syntax.OpCharClass,
		"OpEndLine":                syntax.OpEndLine,
		"OpNoWordBoundary":         syntax.OpNoWordBoundary,
		"isWordChar":               syntax.IsWordChar,
		"ErrInvalidNamedCapture":   syntax.ErrInvalidNamedCapture,
		"ErrMissingBracket":        syntax.ErrMissingBracket,
		"ErrMissingParen":          syntax.ErrMissingParen,
		"ErrUnexpectedParen":       syntax.ErrUnexpectedParen,
		"InstAltMatch":             syntax.InstAltMatch,
		"InstEmptyWidth":           syntax.InstEmptyWidth,
		"InstNop":                  syntax.InstNop,
		"OpRepeat":                 syntax.OpRepeat,
		"MatchNL":                  syntax.MatchNL,
		"OpAnyChar":                syntax.OpAnyChar,
		"OpConcat":                 syntax.OpConcat,
		"emptyOpContext":           syntax.EmptyOpContext,
		"InstRuneAny":              syntax.InstRuneAny,
		"OpPlus":                   syntax.OpPlus,
		"EmptyBeginLine":           syntax.EmptyBeginLine,
		"InstRune":                 syntax.InstRune,
		"EmptyWordBoundary":        syntax.EmptyWordBoundary,
		"ErrLarge":                 syntax.ErrLarge,
		"OpAnyCharNotNL":           syntax.OpAnyCharNotNL,
		"ClassNL":                  syntax.ClassNL,
		"NonGreedy":                syntax.NonGreedy,
		"OpEmptyMatch":             syntax.OpEmptyMatch,
		"OpLiteral":                syntax.OpLiteral,
		"OpNoMatch":                syntax.OpNoMatch,
		"POSIX":                    syntax.POSIX,
		"PerlX":                    syntax.PerlX,

		"emptyRegexp": func() (v syntax.Regexp) {
			return v
		},
		"refRegexp": func() *syntax.Regexp {
			var x syntax.Regexp
			return &x
		},
		"refOfRegexp": func(x syntax.Regexp) *syntax.Regexp {
			return &x
		},
		"emptyInst": func() (v syntax.Inst) {
			return v
		},
		"refInst": func() *syntax.Inst {
			var x syntax.Inst
			return &x
		},
		"refOfInst": func(x syntax.Inst) *syntax.Inst {
			return &x
		},
		"emptyProg": func() (v syntax.Prog) {
			return v
		},
		"refProg": func() *syntax.Prog {
			var x syntax.Prog
			return &x
		},
		"refOfProg": func(x syntax.Prog) *syntax.Prog {
			return &x
		}}
)

func init() {
	engine.RegisterModule(RegexpSyntaxModule{})
}

type RegexpSyntaxModule struct{}

func (S RegexpSyntaxModule) Identity() string {
	return "golang/regexp/syntax"
}
func (S RegexpSyntaxModule) TypeDefine() []byte {
	return RegexpSyntaxDefine
}
func (S RegexpSyntaxModule) Exports() map[string]any {
	return RegexpSyntaxDeclared
}
