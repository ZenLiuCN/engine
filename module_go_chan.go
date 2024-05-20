package engine

import "github.com/dop251/goja"

type GoChan[T any] struct {
	e      *Engine
	raw    chan T
	closer chan struct{}
}

func (g *GoChan[T]) Raw() chan T {
	return g.raw
}
func (g *GoChan[T]) AsSendOnly() GoSendChan[T] {
	return GoSendChan[T]{g.e, g.raw}
}
func (g *GoChan[T]) AsRecvOnly() GoRecvChan[T] {
	return GoRecvChan[T]{g.e, g.raw, nil}
}
func (g *GoChan[T]) Recv(f func(T)) *goja.Promise {
	p, r, j := g.e.NewPromise()
	if g.closer != nil {
		j("already have receiver")
	}
	ch := make(chan struct{})
	g.closer = ch
	go func() {
		defer func() {
			r(nil)
		}()
		for {
			select {
			case <-ch:
				return
			case v, ok := <-g.raw:
				if ok {
					f(v)
				} else {
					return
				}
			}
		}
	}()
	return p
}
func (g *GoChan[T]) Close() {
	if g.closer != nil {
		g.closer <- struct{}{}
		close(g.closer)
		g.closer = nil
	}
	close(g.raw)
}
func (g *GoChan[T]) Stop() bool {
	if g.closer != nil {
		g.closer <- struct{}{}
		close(g.closer)
		g.closer = nil
		return true
	}
	return false
}

type GoRecvChan[T any] struct {
	e      *Engine
	raw    <-chan T
	closer chan struct{}
}

func (g *GoRecvChan[T]) Raw() <-chan T {
	return g.raw
}
func (g *GoRecvChan[T]) Recv(f func(T)) *goja.Promise {
	p, r, j := g.e.NewPromise()
	if g.closer != nil {
		j("already have receiver")
	}
	ch := make(chan struct{})
	g.closer = ch
	go func() {
		defer func() {
			r(nil)
		}()
		for {
			select {
			case <-ch:
				return
			case v, ok := <-g.raw:
				if ok {
					f(v)
				} else {
					return
				}
			}
		}
	}()
	return p
}
func (g *GoRecvChan[T]) Stop() bool {
	if g.closer != nil {
		g.closer <- struct{}{}
		close(g.closer)
		g.closer = nil
		return true
	}
	return false
}

type GoSendChan[T any] struct {
	e   *Engine
	raw chan<- T
}

func (g *GoSendChan[T]) Raw() chan<- T {
	return g.raw
}
func (g *GoSendChan[T]) Send(v T) {
	g.raw <- v
}
func (g *GoSendChan[T]) Close() {
	close(g.raw)
}
