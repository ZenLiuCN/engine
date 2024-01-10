package fetch

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

var (
	//go:embed fetch.d.ts
	fetchDefine []byte
)

func init() {

	engine.RegisterModule(&Fetch{})
}

type Fetch struct {
}

func (f Fetch) Exports() map[string]any {
	return nil
}
func (f Fetch) ExportsWithEngine(e *engine.Engine) map[string]any {
	return map[string]any{
		"fetch": func(url string, option *Option) *goja.Promise {
			jar := new(cookiejar.Jar)
			if option != nil {
				return option.Request(e, jar, url)
			}
			p, ack, rj := e.NewPromise()
			go func() {
				res, err := http.Get(url)
				if err != nil {
					rj(err)
					return
				}
				ack(parse(e, res))
			}()
			return p
		},
	}
}

func (f Fetch) TypeDefine() []byte {
	return fetchDefine
}

func (f Fetch) Identity() string {
	return "go/fetch"
}

type Response struct {
	engine     *engine.Engine
	res        *http.Response
	Status     int
	StatusText string
	Type       string
	Url        string
	Headers    map[string]string
	Ok         bool
	BodyUsed   bool
}

func (r *Response) ArrayBuffer() *goja.Promise {
	p, ack, rj := r.engine.NewPromise()
	if r.BodyUsed {
		rj(http.ErrBodyReadAfterClose)
		return p
	}
	if r.res.Body == nil {
		ack(r.engine.Undefined())
		return p
	}
	go func() {
		defer fn.IgnoreClose(r.res.Body)
		r.BodyUsed = true
		bin, err := io.ReadAll(r.res.Body)
		if err != nil {
			rj(err)
		}
		ack(r.engine.ToValue(r.engine.NewArrayBuffer(bin)))
	}()
	return p
}

/*
	func (r *FetchResponse) Blob() *goja.Promise {
		p, ack, rj := r.engine.NewPromise()
		if r.BodyUsed {
			rj(http.ErrBodyReadAfterClose)
			return p
		}
		go func() {
			defer r.res.Body.Close()
			r.BodyUsed = true
			bin, err := io.ReadAll(r.res.Body)
			if err != nil {
				rj(err)
			}
			ack(r.engine.ToValue(r.engine.NewArrayBuffer(bin)))
		}()
		return p
	}
*/
func (r *Response) Json() *goja.Promise {
	p, ack, rj := r.engine.NewPromise()
	if r.BodyUsed {
		rj(http.ErrBodyReadAfterClose)
		return p
	}
	if r.res.Body == nil {
		ack(r.engine.Undefined())
		return p
	}
	go func() {
		defer fn.IgnoreClose(r.res.Body)
		r.BodyUsed = true
		bin, err := io.ReadAll(r.res.Body)
		if err != nil {
			rj(err)
		}
		if bin[0] == '[' {
			v := make([]any, 0)
			err := json.Unmarshal(bin, &v)
			if err != nil {
				rj(err)
			} else {
				ack(r.engine.ToValue(v))
			}
		} else if bin[0] == '{' {
			v := make(map[string]any)
			err := json.Unmarshal(bin, &v)
			if err != nil {
				rj(err)
			} else {
				ack(r.engine.ToValue(v))
			}
		} else if bin[0] == '"' {
			v := ""
			err := json.Unmarshal(bin, &v)
			if err != nil {
				rj(err)
			} else {
				ack(r.engine.ToValue(v))
			}
		} else if bin[0] == 't' || bin[0] == 'f' {
			v := true
			err := json.Unmarshal(bin, &v)
			if err != nil {
				rj(err)
			} else {
				ack(r.engine.ToValue(v))
			}
		} else {
			v := 0.0
			err := json.Unmarshal(bin, &v)
			if err != nil {
				rj(err)
			} else {
				ack(r.engine.ToValue(v))
			}
		}
	}()
	return p
}
func (r *Response) Text() *goja.Promise {
	p, ack, rj := r.engine.NewPromise()
	if r.BodyUsed {
		rj(http.ErrBodyReadAfterClose)
		return p
	}
	if r.res.Body == nil {
		ack(r.engine.Undefined())
		return p
	}
	go func() {
		defer fn.IgnoreClose(r.res.Body)
		r.BodyUsed = true
		bin, err := io.ReadAll(r.res.Body)
		if err != nil {
			rj(err)
		}
		ack(r.engine.ToValue(string(bin)))
	}()
	return p
}

func parse(e *engine.Engine, response *http.Response) *Response {
	return &Response{
		engine:     e,
		res:        response,
		Headers:    nil,
		Status:     response.StatusCode,
		StatusText: response.Status,
		Type:       "basic",
		Url:        response.Request.URL.String(),
		BodyUsed:   false,
		Ok:         response.StatusCode < 299 && response.StatusCode > 200,
	}
}

type Option struct {
	Method         string
	Mode           string
	Cache          string
	Credentials    string
	Headers        map[string]string
	Redirect       string
	ReferrerPolicy string
	Body           goja.Value
}

func (o *Option) Request(e *engine.Engine, jar *cookiejar.Jar, uri string) *goja.Promise {
	p, ack, rj := e.NewPromise()
	client := new(http.Client)
	switch o.Redirect {
	case "error":
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrNotSupported
		}
	case "manual":

	}
	switch o.Credentials {
	case "include", "same-origin":
		client.Jar = jar
	}
	req := new(http.Request)
	switch o.Method {
	case "GET":
		if !engine.IsNullish(o.Body) {
			rj(http.ErrBodyNotAllowed)
			return p
		} else {
			req.Method = o.Method
		}
	default:
		req.Method = o.Method
	}
	var err error
	req.URL, err = url.Parse(uri)
	if err != nil {
		rj(err)
		return p
	}
	req.Header = make(http.Header)
	contentType := ""
	for s, s2 := range o.Headers {
		req.Header.Add(s, s2)
		if strings.EqualFold(s, "Content-Type") {
			contentType = s2
		}
	}
	switch {
	case engine.IsNullish(o.Body):

	case strings.EqualFold(contentType, "application/json"):
		req.Body = io.NopCloser(strings.NewReader(o.Body.ToString().String()))
	case contentType == "":
		switch v := o.Body.Export().(type) {
		case string:
			req.Body = io.NopCloser(strings.NewReader(v))
		case goja.ArrayBuffer:
			req.Body = io.NopCloser(bytes.NewReader(v.Bytes()))
		case []byte:
			req.Body = io.NopCloser(bytes.NewReader(v))
		case goja.JsonEncodable:
			b, err := json.Marshal(v.JsonEncodable())
			if err != nil {
				rj(err)
			}
			req.Body = io.NopCloser(bytes.NewReader(b))
		default:
			req.Body = io.NopCloser(strings.NewReader(fmt.Sprintf("%s", v)))
		}
	}
	go func() {
		res, err := client.Do(req)
		if err != nil {
			rj(err)
			return
		}
		ack(parse(e, res))
	}()
	return p
}
