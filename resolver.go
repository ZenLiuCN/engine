package engine

import (
	"context"
	"errors"
	"fmt"
	"github.com/dop251/goja"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	Loader Resolver = &BaseResolver{cache: make(map[*url.URL]*CacheSource)}
)

type (
	Source struct {
		Data []byte
		URL  *url.URL
	}
	Resolver interface {
		Dir(old *url.URL) *url.URL
		Load(specifier *url.URL, originalModuleSpecifier string) (*Source, error)
		Resolve(pwd *url.URL, specifier string) (*url.URL, error)
	}
	BaseResolver struct {
		cache map[*url.URL]*CacheSource
	}
	CacheSource struct {
		Src *Source
		Err error
	}
)

func (s *BaseResolver) Fetch(u *url.URL) (*Source, error) {
	if c, ok := s.cache[u]; ok {
		return c.Src, c.Err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()
	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusNotFound:
			return nil, fmt.Errorf("not found: %s", u)
		default:
			return nil, fmt.Errorf("wrong status code (%d) for: %s", res.StatusCode, u)
		}
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		s.cache[u] = &CacheSource{
			Src: nil,
			Err: err,
		}
		return nil, err
	}
	src := &Source{URL: u, Data: data}
	s.cache[u] = &CacheSource{
		Src: src,
		Err: nil,
	}
	return src, nil
}
func (s *BaseResolver) ResolveFilePath(pwd *url.URL, moduleSpecifier string) (*url.URL, error) {
	if pwd.Opaque != "" {
		parts := strings.SplitN(pwd.Opaque, "/", 2)
		if moduleSpecifier[0] == '/' {
			return &url.URL{Opaque: path.Join(parts[0], moduleSpecifier)}, nil
		}
		return &url.URL{Opaque: path.Join(parts[0], path.Join(path.Dir(parts[1]+"/"), moduleSpecifier))}, nil
	}
	if filepath.VolumeName(moduleSpecifier) != "" {
		moduleSpecifier = "/" + moduleSpecifier
	}
	finalPwd := pwd
	if pwd.Opaque != "" {
		if !strings.HasSuffix(pwd.Opaque, "/") {
			finalPwd = &url.URL{Opaque: pwd.Opaque + "/"}
		}
	} else if !strings.HasSuffix(pwd.Path, "/") {
		finalPwd = &url.URL{}
		*finalPwd = *pwd
		finalPwd.Path += "/"
	}
	return finalPwd.Parse(moduleSpecifier)
}
func (s *BaseResolver) LoadFile(u *url.URL) (*Source, error) {
	if c, ok := s.cache[u]; ok {
		return c.Src, c.Err
	}
	var pathOnFs string
	switch {
	case u.Opaque != "":
		if runtime.GOOS != "windows" {
			pathOnFs = filepath.Join(string(os.PathSeparator), u.Opaque)
		} else {
			pathOnFs = u.Opaque
		}
	case u.Scheme == "":
		pathOnFs = path.Clean(u.String())
	default:
		pathOnFs = path.Clean(u.String()[len(u.Scheme)+len(":/"):])
	}
	pathOnFs, err := url.PathUnescape(filepath.FromSlash(pathOnFs))
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(pathOnFs)
	if err == nil {
		src := &Source{URL: u, Data: data}
		s.cache[u] = &CacheSource{
			Src: src,
			Err: nil,
		}
		return src, nil
	}
	if !errors.Is(err, fs.ErrNotExist) {
		s.cache[u] = &CacheSource{
			Src: nil,
			Err: err,
		}
		return nil, err
	}
	return nil, fs.ErrNotExist
}
func (s *BaseResolver) Dir(old *url.URL) *url.URL {
	if old.Opaque != "" {
		return &url.URL{Opaque: path.Join(old.Opaque, "../")}
	}
	return old.ResolveReference(&url.URL{Path: "./"})
}
func (s *BaseResolver) Load(specifier *url.URL, originalModuleSpecifier string) (*Source, error) {
	r, err := s.LoadFile(specifier)
	if err == nil || !errors.Is(err, fs.ErrNotExist) {
		return r, err
	}
	scheme := specifier.Scheme
	if scheme == "" {
		if specifier.Opaque == "" {
			return nil, fmt.Errorf("bad module %s", originalModuleSpecifier)
		}
		scheme = "https"
	}

	var result *Source
	result, err = s.Fetch(specifier)
	if err != nil {
		return nil, fmt.Errorf("loading remote module %s from %s fail: %s", originalModuleSpecifier, specifier, err)
	}
	result.URL = specifier
	return result, nil
}
func (s *BaseResolver) Resolve(pwd *url.URL, specifier string) (*url.URL, error) {
	if specifier == "" {
		return nil, errors.New("local or remote path required")
	}
	if specifier[0] == '.' || specifier[0] == '/' || filepath.IsAbs(specifier) {
		return s.ResolveFilePath(pwd, specifier)
	}
	if strings.Contains(specifier, "://") {
		u, err := url.Parse(specifier)
		if err != nil {
			return nil, err
		}
		if u.Scheme != "file" && u.Scheme != "https" && u.Scheme != "http" {
			return nil,
				fmt.Errorf("only supported schemes for imports are file and https or http, %s has `%s`",
					specifier, u.Scheme)
		}
		if u.Scheme == "file" && pwd.Scheme != "file" {
			return nil, fmt.Errorf("origin (%s) not allowed to load local file: %s", pwd, specifier)
		}
		return u, err
	}
	return &url.URL{Opaque: specifier}, nil
}

type (
	JsModule interface {
		Instance(engine *Engine) JsModuleInstance
	}
	JsModuleInstance interface {
		Execute() error
		Exports() *goja.Object
	}
	cjs struct {
		prg *goja.Program
		url *url.URL
	}
	cjsModule struct {
		mod *cjs
		obj *goja.Object
		*Engine
	}
)

func (c *cjs) Instance(engine *Engine) JsModuleInstance {
	return &cjsModule{
		mod:    c,
		obj:    nil,
		Engine: engine,
	}
}
func (c *cjsModule) Execute() error {
	exports := c.Engine.NewObject()
	c.obj = c.Engine.NewObject()
	err := c.obj.Set("exports", exports)
	if err != nil {
		return fmt.Errorf("error prepare import commonJS, couldn't set exports property of module: %w",
			err)
	}
	err = c.Engine.Runtime.Set("module", c.obj)
	if err != nil {
		return fmt.Errorf("error prepare import commonJS, couldn't set module: %w",
			err)
	}
	f, err := c.Engine.RunProgram(c.mod.prg)
	if err != nil {
		return err
	}
	if call, ok := goja.AssertFunction(f); ok {
		if _, err = call(exports, c.obj, exports); err != nil {
			return err
		}
	}
	err = c.Engine.Runtime.Set("module", nil)
	if err != nil {
		return fmt.Errorf("error after import commonJS, couldn't remove module: %w",
			err)
	}
	return nil
}

func (c *cjsModule) Exports() *goja.Object {
	exports := c.obj.Get("exports")
	if IsNullish(exports) {
		return nil
	}
	return exports.ToObject(c.Engine.Runtime)
}

type (
	ModCache struct {
		mod JsModule
		err error
	}
	ModResolver interface {
		Resolve(basePWD *url.URL, arg string) (JsModule, error)
	}
	BaseModResolver struct {
		cache map[string]ModCache
	}
)

var (
	ModLoader ModResolver = &BaseModResolver{cache: map[string]ModCache{}}
)

func (s *BaseModResolver) Resolve(basePWD *url.URL, arg string) (JsModule, error) {
	if cached, ok := s.cache[arg]; ok {
		return cached.mod, cached.err
	}
	specifier, err := Loader.Resolve(basePWD, arg)
	if err != nil {
		return nil, err
	}
	if cached, ok := s.cache[specifier.String()]; ok {
		return cached.mod, cached.err
	}
	// Fall back to loading
	data, err := Loader.Load(specifier, arg)
	if err != nil {
		s.cache[specifier.String()] = ModCache{err: err}
		return nil, err
	}
	mod, err := CompileCJS(data)
	s.cache[specifier.String()] = ModCache{mod: mod, err: err}
	return mod, err
}

func CompileCJS(data *Source) (m JsModule, err error) {
	if strings.HasSuffix(data.URL.String(), ".ts") {
		m, err = compileTs(data)
		if err == nil {
			return
		}
	}
	return compileJs(data)
}
func compileTs(data *Source) (m JsModule, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch er := r.(type) {
			case error:
				err = er
			case string:
				err = errors.New(er)
			default:
				err = fmt.Errorf("%v", er)
			}
		}
	}()
	p := CompileSource(string(data.Data), true, false)
	return &cjs{
		prg: p.Program,
		url: data.URL,
	}, nil
}
func compileJs(data *Source) (m JsModule, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch er := r.(type) {
			case error:
				err = er
			case string:
				err = errors.New(er)
			default:
				err = fmt.Errorf("%v", er)
			}
		}
	}()
	p := CompileSource(string(data.Data), false, false)
	return &cjs{
		prg: p.Program,
		url: data.URL,
	}, nil
}
