package engine

import (
	"errors"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"net/url"
	"os"
	"path/filepath"
	"slices"
)

type Require struct {
	pwd      *url.URL
	cache    map[JsModule]JsModuleInstance // module instance cache
	disables []string
	*Engine
}

const ModuleRequireConstKey = "_$require"

func (r Require) Name() string {
	return "Require"
}
func (r *Require) AddDisabled(spec ...string) {
	r.disables = append(r.disables, spec...)
}
func (r *Require) GetDisabled() []string {
	return r.disables
}
func (r Require) Register(engine *Engine) {
	x := &Require{
		pwd:    WdToUrl(),
		cache:  make(map[JsModule]JsModuleInstance),
		Engine: engine,
	}
	engine.require = x
	engine.scriptRootHandler = func(p string) {
		if p == "" {
			x.pwd = WdToUrl()
		} else {
			x.pwd = PathToUrl(filepath.ToSlash(p))
		}
	}
	engine.Set("require", x.Require)
	engine.Set(ModuleRequireConstKey, x)
}

var (
	ErrDisabled = errors.New("required module is disabled")
)

func (r *Require) Require(specifier string) (*goja.Object, error) {
	if slices.Index(r.disables, specifier) >= 0 {
		return nil, ErrDisabled
	}
	if gom := resolveModule(specifier); gom != nil {
		return instanceModule(r.Engine, gom)
	}
	current := r.pwd
	defer func() {
		r.pwd = current
	}()
	fileURL, err := Loader.Resolve(r.pwd, specifier)
	if err != nil {
		return nil, err
	}
	r.pwd = Loader.Dir(fileURL)
	if specifier == "" {
		return nil, errors.New("require() can't be used with an empty specifier")
	}
	m, err := ModLoader.Resolve(current, specifier)
	if err != nil {
		return nil, err
	}
	if instance, ok := r.cache[m]; ok {
		return instance.Exports(), nil
	}
	instance := m.Instance(r.Engine)
	r.cache[m] = instance
	if err = instance.Execute(); err != nil {
		return nil, err
	}
	return instance.Exports(), nil
}

func PathToUrl(path string) *url.URL {
	if !filepath.IsAbs(path) {
		path = fn.Panic1(filepath.Abs(path))
	}
	return &url.URL{Scheme: "file", Opaque: filepath.ToSlash(path)}
}
func WdToUrl() *url.URL {
	return PathToUrl(filepath.ToSlash(fn.Panic1(os.Getwd())))
}
