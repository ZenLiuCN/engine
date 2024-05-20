package legacy

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"io"
	http "net/http"
	"net/url"
	"strings"
)

var (
	//go:embed module_http.d.ts
	httpDefine []byte
	httpMap    = map[string]any{
		"get": func(url string) *HttpResponse {
			return newResponse(http.Get(url))
		},
		"head": func(url string) *HttpResponse {
			return newResponse(http.Head(url))
		},
		"post": func(url, contentType string, body goja.Value) *HttpResponse {
			switch v := body.Export().(type) {
			case string:
				return newResponse(http.Post(url, contentType, strings.NewReader(v)))
			case []byte:
				return newResponse(http.Post(url, contentType, bytes.NewReader(v)))
			case io.Reader:
				return newResponse(http.Post(url, contentType, v))
			default:
				j, err := json.Marshal(v)
				if err != nil {
					return &HttpResponse{
						err: err,
					}
				}
				return newResponse(http.Post(url, contentType, bytes.NewReader(j)))
			}
		},
		"postJson": func(url string, body goja.Value) *HttpResponse {
			switch v := body.Export().(type) {
			case string:
				return newResponse(http.Post(url, "application/json", strings.NewReader(v)))
			case []byte:
				return newResponse(http.Post(url, "application/json", bytes.NewReader(v)))
			default:
				j, err := json.Marshal(v)
				if err != nil {
					return &HttpResponse{
						err: err,
					}
				}
				return newResponse(http.Post(url, "application/json", bytes.NewReader(j)))
			}
		},
		"postForm": func(url string, form map[string][]string) *HttpResponse {
			return newResponse(http.PostForm(url, form))
		},
		"values": func() url.Values {
			return map[string][]string{}
		},
		"valuesToHeader": func(v url.Values) http.Header {
			return http.Header(v)
		},
		"headerToValues": func(v http.Header) url.Values {
			return url.Values(v)
		},
		"request": func(req *http.Request) *HttpResponse {
			return newResponse(http.DefaultClient.Do(req))
		},
		"requestOf": func(m, u string, b goja.Value) (r *http.Request, err error) {
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
			switch v := b.Export().(type) {
			case string:
				return http.NewRequest(m, u, strings.NewReader(v))
			case []byte:
				return http.NewRequest(m, u, bytes.NewReader(v))
			case io.Reader:
				return http.NewRequest(m, u, v)
			default:
				return http.NewRequest(m, u, bytes.NewReader(fn.Panic1(json.Marshal(v))))
			}
		},
	}
)

type HttpResponse struct {
	err error
	*http.Response
	closed bool
}

func (s *HttpResponse) Close() {
	if !s.closed && s.Response != nil && s.Response.Body != nil {
		_ = s.Response.Body.Close()
	}
}
func (s *HttpResponse) HasError() bool {
	return s.err != nil
}
func (s *HttpResponse) GetError() error {
	if s.err != nil {
		return s.err
	}
	return nil
}

func (s *HttpResponse) Json() (r any, err error) {
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
	defer fn.IgnoreClose(s.Body)
	s.closed = true
	bin, err := io.ReadAll(s.Body)
	if err != nil {
		panic(err)
	}
	if bin[0] == '[' {
		v := make([]any, 0)
		fn.Panic(json.Unmarshal(bin, &v))
		return v, nil
	} else if bin[0] == '{' {
		v := make(map[string]any)
		fn.Panic(json.Unmarshal(bin, &v))
		return v, nil
	} else if bin[0] == '"' {
		v := ""
		fn.Panic(json.Unmarshal(bin, &v))
		return v, nil
	} else if bin[0] == 't' || bin[0] == 'f' {
		v := true
		fn.Panic(json.Unmarshal(bin, &v))
		return v, nil
	} else {
		v := 0.0
		fn.Panic(json.Unmarshal(bin, &v))
		return v, nil
	}
}
func (s *HttpResponse) Text() (r string, err error) {
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
	defer fn.IgnoreClose(s.Body)
	s.closed = true
	bin, err := io.ReadAll(s.Body)
	if err != nil {
		panic(err)
	}
	return string(bin), nil
}
func (s *HttpResponse) Binary() []byte {
	defer fn.IgnoreClose(s.Body)
	s.closed = true
	return fn.Panic1(io.ReadAll(s.Body))
}
func (s *HttpResponse) HasBody() bool {
	return s.Response != nil && s.Body != nil
}
func newResponse(resp *http.Response, err error) *HttpResponse {
	return &HttpResponse{
		err:      err,
		Response: resp,
	}
}

type HttpModule struct {
}

func (c HttpModule) TypeDefine() []byte {
	return httpDefine
}

func (c HttpModule) Identity() string {
	return "go/http"
}

func (c HttpModule) Exports() map[string]any {
	return httpMap
}
