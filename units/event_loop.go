package units

import (
	"github.com/ZenLiuCN/engine"
	"runtime"
)

func debug(args ...any) {

}
func debugN(n uint, args ...any) {

}

type Engine = engine.Engine
type Job = func()
type EventLoop struct {
	engine   *Engine
	jobQueue []Job
	jobs     chan Job
	quit     chan struct{}
}

func NewEventLoop(engine *Engine) *EventLoop {
	return &EventLoop{engine: engine}
}
func (el *EventLoop) Start() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	select {
	case job := <-el.jobs:
		job()
	case <-el.quit:
		break
	}
}
