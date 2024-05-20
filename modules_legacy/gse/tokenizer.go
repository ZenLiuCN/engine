package gse

import (
	_ "embed"
	"fmt"
	"github.com/ZenLiuCN/engine"
	"github.com/dop251/goja"
	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/idf"
	"github.com/go-ego/gse/hmm/pos"
	"regexp"
)

var (
	//go:embed tokenizer.d.ts
	tokenizerDefine []byte
)

func init() {
	engine.RegisterModule(TokenizerModule{})
}

type TokenizerModule struct {
}

func (S TokenizerModule) Identity() string {
	return "go/gse"
}

func (S TokenizerModule) TypeDefine() []byte {
	return tokenizerDefine
}

func (S TokenizerModule) Exports() map[string]any {
	return nil
}

func (S TokenizerModule) ExportsWithEngine(eng *engine.Engine) map[string]any {
	return map[string]any{
		"version": gse.Version,
		"toLower": func() bool {
			gse.ToLower = !gse.ToLower
			return gse.ToLower
		},
		"newEmbed":     gse.NewEmbed,
		"dictPaths":    gse.DictPaths,
		"isJp":         gse.IsJp,
		"newDict":      gse.NewDict,
		"splitNum":     gse.SplitNum,
		"splitNums":    gse.SplitNums,
		"filterEmoji":  gse.FilterEmoji,
		"filterSymbol": gse.FilterSymbol,
		"filterHtml":   gse.FilterHtml,
		"filterLang":   gse.FilterLang,
		"range":        gse.Range,
		"rangeText":    gse.RangeText,
		"splitWords":   gse.SplitWords,
		"toString":     gse.ToString,
		"toSlice":      gse.ToSlice,
		"toPos":        gse.ToPos,
		"join":         gse.Join,
		"setStopWords": func(m map[string]bool) {
			gse.StopWordMap = m
		},
		"getStopWords": func() map[string]bool {
			return gse.StopWordMap
		},
		"Tokenizer": eng.ToConstructor(func(v []goja.Value) (any, error) {
			seg := &Tokenizer{new(gse.Segmenter)}
			return seg, nil
		}),
		"TagExtractor": eng.ToConstructor(func(v []goja.Value) (any, error) {
			seg := &TagExtractor{new(idf.TagExtracter)}
			return seg, nil
		}),
		"TextRanker": eng.ToConstructor(func(v []goja.Value) (any, error) {
			seg := new(idf.TextRanker)
			return seg, nil
		}),
		"PosTokenizer": eng.ToConstructor(func(v []goja.Value) (any, error) {
			seg := &PosTokenizer{new(pos.Segmenter)}
			return seg, nil
		}),
	}
}

//region Tokenizer

type Tokenizer struct {
	*gse.Segmenter
}

func (t *Tokenizer) CutDAG(str string, reg ...string) []string {
	var regs []*regexp.Regexp
	for _, s := range reg {
		regs = append(regs, regexp.MustCompile(s))
	}
	return t.Segmenter.CutDAG(str, regs...)
}
func (t *Tokenizer) Value(str string) (Value, error) {
	v, i, e := t.Segmenter.Value(str)
	if e != nil {
		return Value{}, e
	}
	return Value{v, i}, nil
}
func (t *Tokenizer) Find(str string) Found {
	v, i, e := t.Segmenter.Find(str)

	return Found{v, i, e}
}

type Found struct {
	Freq  float64
	Pos   string
	Exist bool
}
type Value struct {
	Value int
	Id    int
}

//endregion

//region TagExtractor

type TagExtractor struct {
	*idf.TagExtracter
}

func (s *TagExtractor) ExtractTags(txt string, topK int) (r []*Seg) {
	v := s.TagExtracter.ExtractTags(txt, topK)
	for _, segment := range v {
		r = append(r, &Seg{
			segment.Text(),
			segment.Weight(),
		})
	}
	return

}

type Seg struct {
	Text   string
	Weight float64
}

func (s Seg) String() string {
	return fmt.Sprintf(`{"text":"%s","weight":%f}`, s.Text, s.Weight)
}

//endregion

type PosTokenizer struct {
	*pos.Segmenter
}
