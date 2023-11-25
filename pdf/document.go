package pdf

import (
	"fmt"
	"io"
)

type Document struct {
	Header    string
	Mark      string
	Pages     *Pages
	Resources []Object
	Info      *Info
	Catalog   *Catalog
	XRef      *XRef
	Trailer   *Trailer
	original  io.ReadCloser
}

func (d *Document) Validate() {
	cnt := 0
	if d.Pages == nil {
		panic(fmt.Errorf("root pages missing"))
	}
	cnt++
	d.Pages.ID = cnt
	cnt = d.Pages.Resolve(cnt)
	if d.Resources != nil {
		for _, resource := range d.Resources {
			//resolved by Pages
			if _, ok := resource.(*Pages); ok {
				continue
			}
			if _, ok := resource.(*Page); ok {
				continue
			}
			cnt++
			resource.SetId(cnt)
		}
	}
	if d.Info == nil {
		d.Info = new(Info)
	}
	cnt++
	d.Info.ID = cnt
	if d.Catalog == nil {
		d.Catalog = &Catalog{
			Pages: d.Pages.ToRef(),
		}
	}
	cnt++
	d.Catalog.ID = cnt
	if d.Trailer == nil {
		d.Trailer = new(Trailer)
	}
	if d.Info != nil {
		d.Trailer.Info = d.Info.ToRef()
	}

}
func (d *Document) Write(out io.Writer) {
	w := NewWriter(out)
	if d.Header == "" {
		d.Header = Header
	}
	w.Raw(d.Header).LineFeed()
	if d.Mark == "" {
		d.Mark = Mark
	}
	w.Raw(d.Mark).LineFeed()
	if d.XRef == nil {
		d.XRef = new(XRef)
	}
	d.Validate()
	//pages
	{
		d.XRef.Entry = append(d.XRef.Entry, w.Size())
		w.Use(d.Pages.Write)
	}
	//resources
	{
		for _, p := range d.Resources {
			d.XRef.Entry = append(d.XRef.Entry, w.Size())
			w.Use(p.Write)
		}
	}
	//info
	{
		d.XRef.Entry = append(d.XRef.Entry, w.Size())
		w.Use(d.Info.Write)
	}
	//catalog
	{
		d.XRef.Entry = append(d.XRef.Entry, w.Size())
		w.Use(d.Catalog.Write)
	}

	//xref
	xrefOffset := w.Size()
	w.Use(d.XRef.Write)

	n := len(d.XRef.Entry)
	d.Trailer.Root = d.Catalog.ToRef()
	d.Trailer.Size = n + 1
	w.Use(d.Trailer.Write)
	w.Raw(startxref).LineFeed().Int(xrefOffset).LineFeed().Raw(EOF)
}
func (d *Document) Close() error {
	if d.original != nil {
		return d.original.Close()
	}
	return nil
}

const (
	Header   = "%PDF-1.4"
	Mark     = "%Ζεν·Ληυ"
	Producer = "Ζεν·Ληυ"
	EOF      = "%%EOF"
)
