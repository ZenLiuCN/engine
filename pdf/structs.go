package pdf

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

type Type int

const (
	TypeNull Type = iota
	TypeBoolean
	TypeInt
	TypeReal
	TypeText
	TypeName
	TypeBinary
	TypeDate
	TypeArray
	TypeDict
	TypeObject
)

type ElementType int

const (
	ElementTrailer ElementType = iota
	ElementPage
	ElementPages
	ElementRect
	ElementFonts
	ElementFont
	ElementXRef
	ElementStream
)

type Value interface {
	Type() Type
	Write(writer Writer)
}
type Element interface {
	Value
	ElementType() ElementType
	ObjectType() Name
}
type Binary []byte

func (b Binary) Type() Type {
	return TypeBinary
}
func (b Binary) Write(writer Writer) {
	writer.Binary(b)
}

type Name string

func (b Name) Type() Type {
	return TypeName
}
func (b Name) Write(writer Writer) {
	writer.Name(string(b))
}

type Text string

func (b Text) Type() Type {
	return TypeText
}
func (b Text) Write(writer Writer) {
	writer.Text(string(b))
}

type Int int

func (b Int) Type() Type {
	return TypeInt
}
func (b Int) Write(writer Writer) {
	writer.Int(int(b))
}

type Real float64

func (b Real) Type() Type {
	return TypeReal
}
func (b Real) Write(writer Writer) {
	writer.Real(float64(b))
}

type Boolean bool

func (b *Boolean) Type() Type {
	return TypeBoolean
}
func (b *Boolean) Write(writer Writer) {
	writer.Boolean(bool(*b))
}

type Date time.Time

func (b Date) Type() Type {
	return TypeDate
}
func (b Date) Write(writer Writer) {
	writer.Date(time.Time(b))
}

type Array []Value

func (b Array) Type() Type {
	return TypeArray
}
func (b Array) Write(writer Writer) {
	writer.Array(func(w Writer) {
		for i, value := range b {
			if i > 0 {
				writer.Space()
			}
			value.Write(writer)
		}
	})
}

type ObjectArray struct {
}

func (b *ObjectArray) Type() Type {
	return TypeArray
}

type Dict map[Name]Value

func (b Dict) Type() Type {
	return TypeArray
}
func (b Dict) Write(writer Writer) {
	writer.Dict(func(w Writer) {
		for k, value := range b {
			w.Entry(string(k), value).LineFeed()
		}
	}).LineFeed()
}

type Ref struct {
	ID    int
	Gene  int
	Refer Object
}

func (b *Ref) Type() Type {
	return TypeArray
}
func (b *Ref) Write(writer Writer) {
	if b.ID == 0 && b.Refer != nil {
		writer.Ref(b.Refer.Id(), b.Refer.Gen())
	} else {
		writer.Ref(b.ID, b.Gene)
	}
}

type Object interface {
	Value
	Id() int
	Gen() int
	SetId(int)
	SetGen(int)
	ToRef() *Ref
}
type Base struct {
	ID   int
	Gene int
}

func (b *Base) Type() Type {
	return TypeObject
}
func (b *Base) Id() int {
	return b.ID
}
func (b *Base) Gen() int {
	return b.Gene
}
func (b *Base) SetId(v int) {
	b.ID = v
}
func (b *Base) SetGen(v int) {
	b.Gene = v
}

type Page struct {
	Base
	Parent    *Ref  //required
	Resources *Ref  //optional inherited
	Contents  Array //optional
	Rotate    int   //optional
	MediaBox  *Rect //optional inherited
	CropBox   *Rect //optional
}

func (p *Page) ObjectType() Name {
	return page
}
func (p *Page) ElementType() ElementType {
	return ElementPage
}
func (p *Page) Write(writer Writer) {
	writer.Object(p.ID, p.Gene, func(w Writer) {
		w.Dict(func(w Writer) {
			w.Entry(typeName, Name(page)).Space()
			if p.Parent != nil {
				w.Entry(parent, p.Parent).Space()
			}
			if p.Resources != nil {
				w.Entry(resources, p.Resources).Space()
			}
			if p.Contents != nil {
				w.Entry(contents, p.Contents).Space()
			}
			if p.Rotate != 0 {
				w.Entry(rotate, Int(p.Rotate)).Space()
			}
			if p.MediaBox != nil {
				w.Entry(mediaBox, p.MediaBox).Space()
			}
			if p.CropBox != nil {
				w.Entry(cropBox, p.CropBox).Space()
			}
		}).LineFeed()
	}).LineFeed()

}
func (p *Page) ToRef() *Ref {
	return &Ref{
		ID:    p.ID,
		Gene:  p.Gene,
		Refer: p,
	}
}

type Pages struct {
	Base
	Parent *Ref //optional
	Count  int  //auto resolve
	Kids   []*Ref
}

func (p *Pages) Resolve(objCnt int) int {
	p.Count = len(p.Kids)
	for _, kid := range p.Kids {
		switch x := kid.Refer.(type) {
		case *Pages:
			objCnt++
			x.ID = objCnt
			objCnt = x.Resolve(objCnt)
		case *Page:
			x.ID = objCnt
		default:
			panic(fmt.Errorf("can't have %T in pages", x))
		}
	}
	return objCnt
}
func (p *Pages) ObjectType() Name {
	return pages
}
func (p *Pages) ElementType() ElementType {
	return ElementPages
}
func (p *Pages) Write(writer Writer) {
	writer.Object(p.ID, p.Gene, func(w Writer) {
		w.Dict(func(w Writer) {
			w.Entry(typeName, Name(pages)).Space()
			if p.Parent != nil {
				w.Entry(parent, p.Parent).Space()
			}
			if p.Count != 0 {
				w.Entry(rotate, Int(p.Count)).Space()
			}
			if p.Kids != nil {
				w.Name(kids).Space().Array(func(w Writer) {
					for i, kid := range p.Kids {
						if i > 0 {
							w.Space()
						}
						w.Use(kid.Write)
					}
				}).Space()
			}
		}).LineFeed()
	}).LineFeed()
}
func (p *Pages) ToRef() *Ref {
	return &Ref{
		ID:    p.ID,
		Gene:  p.Gene,
		Refer: p,
	}
}

type Rect struct {
	ObjectArray
	X, Y, W, H int
}

func (b *Rect) Type() Type {
	return TypeArray
}
func (b *Rect) ElementType() ElementType {
	return ElementRect
}
func (b *Rect) Write(writer Writer) {
	writer.Array(func(w Writer) {
		w.
			Int(b.X).Space().
			Int(b.Y).Space().
			Int(b.W).Space().
			Int(b.H)
	})

}

type Font struct {
	Base    Name
	Subtype Name
}

func (b *Font) Type() Type {
	return TypeDict
}
func (b *Font) ElementType() ElementType {
	return ElementFont
}
func (b *Font) Write(writer Writer) {
	writer.Dict(func(w Writer) {
		w.Name(typeName).Space().Name(font).Space()
		if b.Base != "" {
			w.Name(baseFont).Space().Name(string(b.Base)).Space()
		}
		if b.Subtype != "" {
			w.Name(subtype).Space().Name(string(b.Subtype)).Space()
		}
	}).LineFeed()

}

type Fonts struct {
	Base
	Fonts map[Name]*Font
}

func (p *Fonts) ToRef() *Ref {
	return &Ref{
		ID:    p.ID,
		Gene:  p.Gene,
		Refer: p,
	}
}
func (p *Fonts) ObjectType() Name {
	return font
}
func (p *Fonts) ElementType() ElementType {
	return ElementFonts
}
func (p *Fonts) Write(writer Writer) {
	writer.Object(p.ID, p.Gene, func(w Writer) {
		w.Dict(func(w Writer) {
			w.Name(font).LineFeed().
				Dict(func(w Writer) {
					for name, f := range p.Fonts {
						writer.Name(string(name)).Space().Use(f.Write)
					}
				}).LineFeed()
		}).LineFeed()
	}).LineFeed()

}

type Catalog struct {
	Base
	Pages *Ref
}

func (p *Catalog) ToRef() *Ref {
	return &Ref{
		ID:    p.ID,
		Gene:  p.Gene,
		Refer: p,
	}
}
func (p *Catalog) Write(writer Writer) {
	writer.Object(p.ID, p.Gene, func(w Writer) {
		w.Dict(func(w Writer) {
			w.Entry(typeName, Name(catalog)).Space()
			w.Entry(pages, p.Pages).Space()
		}).LineFeed()
	}).LineFeed()
}

func (p *Catalog) ObjectType() Name {
	return catalog
}

type XRef struct {
	Begin int
	Entry []int
}

func (b *XRef) ElementType() ElementType {
	return ElementXRef
}
func (b *XRef) Write(writer Writer) {
	writer.Raw(xref).LineFeed().
		Int(b.Begin).Space().Int(b.Begin + len(b.Entry) + 1).LineFeed().
		Raw(PaddingZero(10, 0)).Space().Raw(PaddingZero(5, 65535)).Space().RawRune('f').LineFeed()
	for _, i := range b.Entry {
		writer.Raw(PaddingZero(10, i)).Space().Raw(PaddingZero(5, 0)).Space().RawRune('n').LineFeed()
	}
}

type Trailer struct {
	Size int
	Root *Ref
	Info *Ref
}

func (b *Trailer) ElementType() ElementType {
	return ElementTrailer
}
func (b *Trailer) Write(writer Writer) {
	writer.Raw(trailer).LineFeed().Dict(func(w Writer) {
		w.
			Entry(root, b.Root).Space().
			Entry(size, Int(b.Size)).Space()
		if b.Info != nil {
			w.Entry(info, b.Info).Space()
		}
	}).LineFeed()
}

func PaddingZero(n int, v int) string {
	s := strings.Builder{}
	vx := fmt.Sprintf("%d", v)
	x := n - len(vx)
	if x > 0 {
		s.WriteString(strings.Repeat("0", x))
	}
	s.WriteString(vx)
	return s.String()
}

type Info struct {
	Base
	Title, Subject, Keywords, Author string
	CreationDate, ModDate            time.Time
	Creator, Producer                string
}

func (p *Info) ElementType() ElementType {
	return ElementStream
}
func (p *Info) Write(writer Writer) {
	writer.Object(p.ID, p.Gene, func(w Writer) {
		w.Dict(func(w Writer) {
			if p.Title != "" {
				w.Entry(title, Text(p.Title))
			}
			if p.Subject != "" {
				w.Entry(subject, Text(p.Subject))
			}
			if p.Keywords != "" {
				w.Entry(keywords, Text(p.Keywords))
			}
			if p.Author != "" {
				w.Entry(author, Text(p.Author))
			}
			if !p.CreationDate.IsZero() {
				w.Entry(creationDate, Date(p.CreationDate))
			}
			if !p.ModDate.IsZero() {
				w.Entry(modDate, Date(p.ModDate))
			}
			if p.Creator != "" {
				w.Entry(creator, Text(p.Creator))
			} else {
				w.Entry(creator, Text(Producer))
			}
			if p.Producer != "" {
				w.Entry(producer, Text(p.Producer))
			} else {
				w.Entry(producer, Text(Producer))
			}
		}).LineFeed()
	}).LineFeed()

}
func (p *Info) ToRef() *Ref {
	return &Ref{
		ID:    p.ID,
		Gene:  p.Gene,
		Refer: p,
	}
}

type Stream struct {
	Base
	Filter string
	Stream *bytes.Buffer
}

func (p *Stream) ToRef() *Ref {
	return &Ref{
		ID:    p.ID,
		Gene:  p.Gene,
		Refer: p,
	}
}

func (p *Stream) ElementType() ElementType {
	return ElementStream
}
func (p *Stream) Write(writer Writer) {
	writer.Object(p.ID, p.Gene, func(w Writer) {
		w.Dict(func(w Writer) {
			w.Entry(length, Int(p.Stream.Len())).Space()
			if p.Filter != "" {
				w.Entry(filter, Name(p.Filter)).Space()
			}
		}).LineFeed()
		writer.Stream(p.Stream).LineFeed()
	}).LineFeed()

}

type DictAny struct {
	Base
	Content Dict
}

func (p *DictAny) ElementType() ElementType {
	return ElementStream
}
func (p *DictAny) Write(writer Writer) {
	writer.Object(p.ID, p.Gene, func(w Writer) {
		w.Dict(func(w Writer) {
			if p.Content != nil {
				for name, value := range p.Content {
					w.Entry(string(name), value)
				}
			}
		}).LineFeed()
	}).LineFeed()

}
func (p *DictAny) ToRef() *Ref {
	return &Ref{
		ID:    p.ID,
		Gene:  p.Gene,
		Refer: p,
	}
}
