package internal

import (
	"errors"
	"golang.org/x/tools/go/packages"
)

type Inspector interface {
	inspect(*packages.Package)
	initialize(*packages.Config)
}
type BaseInspector[T Inspector] struct {
	Parser func(*packages.Config, ...string) ([]*packages.Package, error)

	Inspector T
}

func (s *BaseInspector[T]) initialize(conf *packages.Config) {
	if s.Parser == nil {
		s.Parser = packages.Load
	}
	s.Inspector.initialize(conf)
}
func (s *BaseInspector[T]) Inspect(configure func(*packages.Config), checks func([]*packages.Package) error, patterns ...string) (err error) {

	var pkg []*packages.Package
	if pkg, err = s.parse(configure, patterns...); err != nil {
		return err
	}
	if checks != nil {
		if err = checks(pkg); err != nil {
			return err
		}
	}
	for _, k := range pkg {
		s.Inspector.inspect(k)
	}

	return
}
func (s *BaseInspector[T]) parse(cf func(*packages.Config), patterns ...string) (pkg []*packages.Package, err error) {
	conf := &packages.Config{Tests: false}
	s.initialize(conf)
	if cf != nil {
		cf(conf)
	}
	if pkg, err = s.Parser(conf, patterns...); err != nil {
		return
	}
	for _, p := range pkg {
		if len(p.Errors) > 0 {
			if err == nil {
				err = errors.New(p.Errors[0].Msg)
			} else {
				err = errors.Join(err, errors.New(p.Errors[0].Msg))
			}
		}
	}
	return
}
