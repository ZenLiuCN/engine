package engine

import (
	_ "embed"
	"errors"
	"github.com/dop251/goja"
	"sync/atomic"
)

type GoError struct {
	Err error
}

func (g *GoError) String() string {
	if g == nil || g.Err == nil {
		return "<nil>"
	}
	return g.Err.Error()
}
func (g *GoError) Same(err GoError) bool {
	return errors.Is(g.Err, err.Err)
}
func (g *GoError) Error() string {
	return g.String()
}

var (
	//go:embed module_go.d.ts
	goDefine []byte
)

type GoModule struct {
}

func (c GoModule) Identity() string {
	return "go"
}

func (c GoModule) TypeDefine() []byte {
	return goDefine
}

func (c GoModule) Exports() map[string]any {
	return EmptyMap
}

type (
	Chan[T any] interface {
		Recv(func(T)) *goja.Promise
		Send(v T)
		Closed() bool
		Close()
	}
	ReadChan[T any] interface {
		Recv(func(T)) *goja.Promise
		Closed() bool
		Close()
	}
	WriteChan[T any] interface {
		Send(v T)
		Closed() bool
		Close()
	}
)
type rwch[T any] struct {
	e      *Engine
	ch     chan T
	closer chan<- struct{}
	closed atomic.Bool
}

func NewChan[T any](ch chan T, e *Engine) Chan[T] {
	return &rwch[T]{e: e, ch: ch, closed: atomic.Bool{}}
}

func (c *rwch[T]) Recv(f func(T)) *goja.Promise {
	p, r, j := c.e.NewPromise()
	if c.closer != nil {
		j("already have receiver")
	}
	ch := make(chan struct{})
	c.closer = ch
	go func() {
		defer func() {
			r(nil)
		}()
		for {
			select {
			case <-ch:
				return
			case v, ok := <-c.ch:
				if ok {
					f(v)
				} else {
					c.closed.Store(true)
					return
				}
			}
		}
	}()
	return p
}

func (c *rwch[T]) Send(v T) {
	c.ch <- v
}
func (c *rwch[T]) Close() {
	if c.closed.Load() {
		return
	}
	if c.closer != nil {
		c.closer <- struct{}{}
	}
	c.closed.Store(true)
	close(c.ch)
	if c.closer != nil {
		close(c.closer)
		c.closer = nil
	}
}
func (c *rwch[T]) Closed() bool {
	return c.closed.Load()
}

type roch[T any] struct {
	e      *Engine
	ch     <-chan T
	closer chan<- struct{}
	closed atomic.Bool
}

func NewChanReadOnly[T any](ch <-chan T, e *Engine) ReadChan[T] {
	return &roch[T]{e: e, ch: ch, closed: atomic.Bool{}}
}
func (c *roch[T]) Recv(f func(T)) *goja.Promise {
	p, r, j := c.e.NewPromise()
	if c.closer != nil {
		j("already have receiver")
	}
	ch := make(chan struct{})
	c.closer = ch
	go func() {
		defer func() {
			r(nil)
		}()
		for {
			select {
			case <-ch:
				return
			case v, ok := <-c.ch:
				if ok {
					f(v)
				} else {
					c.closed.Store(true)
					return
				}
			}
		}
	}()
	return p
}
func (c *roch[T]) Close() {
	if c.closed.Load() {
		return
	}
	if c.closer != nil {
		c.closer <- struct{}{}
	}
	c.closed.Store(true)
	if c.closer != nil {
		close(c.closer)
		c.closer = nil
	}
}
func (c *roch[T]) Closed() bool {
	return c.closed.Load()
}

type wch[T any] struct {
	ch     chan<- T
	closed atomic.Bool
}

func NewChanWriteOnly[T any](ch chan<- T) WriteChan[T] {
	return &wch[T]{ch: ch, closed: atomic.Bool{}}
}
func (c *wch[T]) Send(v T) {
	c.ch <- v
}
func (c *wch[T]) Close() {
	if c.closed.Load() {
		return
	}
	c.closed.Store(true)
	close(c.ch)

}
func (c *wch[T]) Closed() bool {
	return c.closed.Load()
}