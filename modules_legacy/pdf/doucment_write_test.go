package pdf

import (
	"bytes"
	"os"
	"testing"
)

func TestSimpleWrite(t *testing.T) {
	d := new(Document)
	d.Pages = new(Pages)
	d.Resources = []Object{
		&DictAny{},
		&Stream{
			Stream: bytes.NewBuffer([]byte("200 150 m 600 450 l S")),
		},
	}
	d.Resources = append(d.Resources, &Page{
		Parent:    d.Pages.ToRef(),
		Resources: d.Resources[0].ToRef(),
		Contents:  Array{d.Resources[1].ToRef()},
		Rotate:    0,
		MediaBox: &Rect{
			X: 0,
			Y: 0,
			W: 792,
			H: 612,
		},
		CropBox: nil,
	})
	d.Pages.Kids = append(d.Pages.Kids, d.Resources[2].ToRef())
	buf := bytes.NewBuffer(nil)
	d.Write(buf)
	_ = os.WriteFile("test.pdf", buf.Bytes(), os.ModePerm)
}

func BenchmarkSimpleWrite(b *testing.B) {
	buf := bytes.NewBuffer(nil)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		d := new(Document)
		d.Pages = new(Pages)
		d.Resources = []Object{
			&DictAny{},
			&Stream{
				Stream: bytes.NewBuffer([]byte("200 150 m 600 450 l S")),
			},
		}
		d.Resources = append(d.Resources, &Page{
			Parent:    d.Pages.ToRef(),
			Resources: d.Resources[0].ToRef(),
			Contents:  Array{d.Resources[1].ToRef()},
			Rotate:    0,
			MediaBox: &Rect{
				X: 0,
				Y: 0,
				W: 792,
				H: 612,
			},
			CropBox: nil,
		})
		d.Pages.Kids = append(d.Pages.Kids, d.Resources[2].ToRef())
		d.Write(buf)
		buf.Reset()
	}
}
