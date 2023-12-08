package minify

import (
	_ "embed"
	"fmt"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/gabriel-vasile/mimetype"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	//go:embed minify.d.ts
	minifyDefine []byte
	minifyMap    = map[string]any{
		"setESVersion": ESVersion,
		"minifier":     Get,
		"fileToFile": func(in, out, mime string) error {
			if mime == "" {
				return GuessFileToFile(in, out)
			}
			return FileToFile(in, out, mime)
		},
		"fileToFolder": func(in, out, mime string) error {
			if mime == "" {
				return GuessFileToFolder(in, out)
			}
			return FileToFolder(in, out, mime)
		},
		"folderToFolder": func(in, out, mime string) error {
			if mime == "" {
				return GuessFolderToFolder(in, out)
			}
			return FolderToFolder(in, out, mime)
		},
	}
)

func init() {
	engine.RegisterModule(Minify{})
}

type Minify struct {
}

func (p Minify) Identity() string {
	return "go/minify"
}

func (p Minify) TypeDefine() []byte {
	return minifyDefine
}

func (p Minify) Exports() map[string]any {

	return minifyMap
}

var (
	m *minify.M
)

func ESVersion(v int) {
	Get().AddRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), &js.Minifier{Version: v})
}
func Get() *minify.M {
	if m != nil {
		return m
	}
	m = minify.New()
	m.Add("text/css", &css.Minifier{KeepCSS2: true})
	m.Add("text/html", &html.Minifier{KeepConditionalComments: true, KeepDocumentTags: true, KeepQuotes: true})
	m.Add("image/svg+xml", &svg.Minifier{})
	m.AddRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), &js.Minifier{Version: 2019})
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
	return m
}
func GuessFileToFile(in, out string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%s", v)
			}
		}
	}()
	if !requiredFile(in) {
		panic(in + " not a valid file")
	}
	if !nothingOrFile(out) {
		panic(out + " not a valid file")
	}
	mime := fn.Panic1(detect(in))

	fo := fn.Panic1(os.OpenFile(out, os.O_TRUNC|os.O_CREATE, os.ModePerm))
	defer fo.Close()

	fi := fn.Panic1(os.OpenFile(in, os.O_RDONLY, os.ModePerm))
	defer fi.Close()

	return Get().Minify(mime.String(), fo, fi)
}
func GuessFileToFolder(in, out string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%s", v)
			}
		}
	}()
	if !nothingOrFolder(out) {
		panic(out + " not a valid folder")
	}
	if !exists(out) {
		fn.Panic(os.MkdirAll(out, os.ModePerm))
	}
	return GuessFileToFile(in, filepath.Join(out, filepath.Base(in)))
}

func GuessFolderToFolder(in, out string) (err error) {

	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%s", v)
			}
		}
	}()

	if !nothingOrFolder(out) {
		panic(out + " not a valid folder")
	}
	if !requiredFolder(in) {
		panic(in + " not a valid folder")
	}
	if !exists(out) {
		fn.Panic(os.MkdirAll(out, os.ModePerm))
	}
	return filepath.Walk(in, func(path string, info fs.FileInfo, err error) (erri error) {
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				case error:
					erri = v
				default:
					erri = fmt.Errorf("%s", v)
				}
			}
		}()
		if !info.IsDir() {
			mime, err := detect(path)
			if err != nil {
				log.Printf("ingore file  %s which can't detect mime-type: %s \n", path, err)
				return nil
			}
			rel := filepath.Join(out, fn.Panic1(filepath.Rel(in, path)))
			relP, _ := filepath.Split(rel)
			if !exists(relP) {
				fn.Panic(os.MkdirAll(out, os.ModePerm))
			}
			return FileToFile(path, rel, mime.String())
		}
		return nil
	})
}
func Guess(in, out string) error {

	ii, ie := os.Stat(in)
	if ie != nil {
		return ie
	}
	oi, oe := os.Stat(out)
	if oe != nil && !os.IsNotExist(oe) {
		return oe
	} else if !oi.IsDir() && ii.IsDir() {
		return fmt.Errorf("cant minify folder %s to file %s", in, out)
	}
	sameExt := filepath.Ext(out) == filepath.Ext(in)
	switch {
	case !ii.IsDir() && sameExt:
		return GuessFileToFile(in, out)
	case !ii.IsDir() && !sameExt:
		return GuessFileToFolder(in, out)
	default:
		return GuessFolderToFolder(in, out)
	}
}
func Files(in, out, mime string) error {
	ii, ie := os.Stat(in)
	if ie != nil {
		return ie
	}
	oi, oe := os.Stat(out)
	if oe != nil && !os.IsNotExist(oe) {
		return oe
	} else if !oi.IsDir() && ii.IsDir() {
		return fmt.Errorf("cant minify folder %s to file %s", in, out)
	}
	sameExt := filepath.Ext(out) == filepath.Ext(in)
	if mime == "" {
		switch {
		case !ii.IsDir() && sameExt:
			return GuessFileToFile(in, out)
		case !ii.IsDir() && !sameExt:
			return GuessFileToFolder(in, out)
		default:
			return GuessFolderToFolder(in, out)
		}
	} else {
		switch {
		case !ii.IsDir() && sameExt:
			return FileToFile(in, out, mime)
		case !ii.IsDir() && !sameExt:
			return FileToFolder(in, out, mime)
		default:
			return FolderToFolder(in, out, mime)
		}
	}
}

func FileToFile(in, out, mime string) (err error) {

	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%s", v)
			}
		}
	}()
	if !requiredFile(in) {
		panic(in + " not a valid file")
	}
	if !nothingOrFile(out) {
		panic(out + " not a valid file")
	}
	mimeT := fn.Panic1(detect(in))
	if mimeT.String() == mime {
		fo := fn.Panic1(os.OpenFile(out, os.O_TRUNC|os.O_CREATE, os.ModePerm))
		defer fn.IgnoreClose(fo)
		fi := fn.Panic1(os.OpenFile(in, os.O_RDONLY, os.ModePerm))
		defer fn.IgnoreClose(fi)
		return Get().Minify(mime, fo, fi)
	}
	return fmt.Errorf("%s of input %s not same as target mime %s", mimeT, in, mime)

}
func FileToFolder(in, out, mime string) (err error) {

	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%s", v)
			}
		}
	}()
	if !nothingOrFolder(out) {
		panic(out + " not a valid folder")
	}
	if !exists(out) {
		fn.Panic(os.MkdirAll(out, os.ModePerm))
	}
	return FileToFile(in, filepath.Join(out, filepath.Base(in)), mime)
}

func FolderToFolder(in, out, mime string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%s", v)
			}
		}
	}()
	if !nothingOrFolder(out) {
		panic(out + " not a valid folder")
	}
	if !requiredFolder(in) {
		panic(in + " not a valid folder")
	}
	if !exists(out) {
		fn.Panic(os.MkdirAll(out, os.ModePerm))
	}
	return filepath.Walk(in, func(path string, info fs.FileInfo, err error) (erri error) {
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				case error:
					erri = v
				default:
					erri = fmt.Errorf("%s", v)
				}
			}
		}()
		if !info.IsDir() {
			mimeT, err := detect(path)
			if err != nil {
				log.Printf("ingore file  %s which can't detect mime-type: %s \n", path, err)
				return nil
			}
			if mimeT.String() != mime {
				log.Printf("ingore file  %s[%s] which not match target mime-type: %s \n", path, mimeT, mime)
				return nil
			}
			rel := filepath.Join(out, fn.Panic1(filepath.Rel(in, path)))
			relP, _ := filepath.Split(rel)
			if !exists(relP) {
				fn.Panic(os.MkdirAll(out, os.ModePerm))
			}
			return FileToFile(path, rel, mime)
		}
		return nil
	})
}

func detect(path string) (*mimetype.MIME, error) {
	mime, err := mimetype.DetectFile(path)
	if err != nil {
		return nil, err
	}
	if strings.HasSuffix(mime.String(), "; charset=utf-8") {
		mime = mimetype.Lookup(strings.TrimSuffix(mime.String(), "; charset=utf-8"))
	}
	if strings.HasPrefix(mime.String(), "text/plain") && strings.EqualFold(filepath.Ext(path), ".js") {
		return mimetype.Lookup("application/javascript"), err
	}
	return mime, err
}
func requiredFile(name string) bool {
	i, e := os.Stat(name)
	if e != nil && os.IsNotExist(e) {
		return false
	} else if e != nil {
		panic(e)
	}
	return !i.IsDir()
}
func nothingOrFile(name string) bool {
	i, e := os.Stat(name)
	if e != nil && os.IsNotExist(e) {
		return true
	} else if e != nil {
		panic(e)
	}
	return !i.IsDir()
}
func nothingOrFolder(name string) bool {
	i, e := os.Stat(name)
	if e != nil && os.IsNotExist(e) {
		return true
	} else if e != nil {
		panic(e)
	}
	return i.IsDir()
}
func requiredFolder(name string) bool {
	i, e := os.Stat(name)
	if e != nil && os.IsNotExist(e) {
		return false
	} else if e != nil {
		panic(e)
	}
	return i.IsDir()
}
func exists(name string) bool {
	_, e := os.Stat(name)
	if e != nil && os.IsNotExist(e) {
		return false
	} else if e != nil {
		panic(e)
	}
	return true
}
