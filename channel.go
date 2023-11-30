package engine

import (
	_ "embed"
	"github.com/dop251/goja"
	"sync/atomic"
)

var (
	//go:embed chan.d.ts
	chanDefine []byte
)

type ChannelModule struct {
}

func (c ChannelModule) Identity() string {
	return "go/chan"
}

func (c ChannelModule) TypeDefine() []byte {
	return chanDefine
}

func (c ChannelModule) Exports() map[string]any {
	return EMPTY
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
	SendChan[T any] interface {
		Send(v T)
		Closed() bool
		Close()
	}
)
type Channel[T any] struct {
	e      *Engine
	ch     chan T
	closer chan<- struct{}
	closed atomic.Bool
}

func NewChannel[T any](ch chan T, e *Engine) *Channel[T] {
	return &Channel[T]{e: e, ch: ch, closed: atomic.Bool{}}
}

func (c *Channel[T]) Recv(f func(T)) *goja.Promise {
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

func (c *Channel[T]) Send(v T) {
	c.ch <- v
}
func (c *Channel[T]) Close() {
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
func (c *Channel[T]) Closed() bool {
	return c.closed.Load()
}

type ReadOnlyChannel[T any] struct {
	e      *Engine
	ch     <-chan T
	closer chan<- struct{}
	closed atomic.Bool
}

func NewReadOnlyChannel[T any](ch chan T, e *Engine) *ReadOnlyChannel[T] {
	return &ReadOnlyChannel[T]{e: e, ch: ch, closed: atomic.Bool{}}
}
func (c *ReadOnlyChannel[T]) Recv(f func(T)) *goja.Promise {
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
func (c *ReadOnlyChannel[T]) Close() {
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
func (c *ReadOnlyChannel[T]) Closed() bool {
	return c.closed.Load()
}

type WriteOnlyChannel[T any] struct {
	ch     chan<- T
	closed atomic.Bool
}

func NewWriteOnlyChannel[T any](ch chan T) *WriteOnlyChannel[T] {
	return &WriteOnlyChannel[T]{ch: ch, closed: atomic.Bool{}}
}
func (c *WriteOnlyChannel[T]) Send(v T) {
	c.ch <- v
}
func (c *WriteOnlyChannel[T]) Close() {
	if c.closed.Load() {
		return
	}
	c.closed.Store(true)
	close(c.ch)

}
func (c *WriteOnlyChannel[T]) Closed() bool {
	return c.closed.Load()
}
