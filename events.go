package engine

/**
must code port from https://github.com/dop251/goja_nodejs
*/
import (
	"context"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"sync"
	"sync/atomic"
	"time"
)

type job struct {
	cancelled bool
	fn        func()
}

type Timer struct {
	job
	timer *time.Timer
}

type Interval struct {
	job
	ticker   *time.Ticker
	stopChan chan struct{}
}

type Immediate struct {
	job
}

// EventLoop process all async events and also  as The AsyncContextTracker
// there will be a background goroutine to monitor the async tasks as call StartEventLoop,
// or call StartEventLoopWhenNotStarted for not sure current EventLoop have been started.
// use StopEventLoopWait to await all async tasks are finished.
// use StopEventLoopTimeout to waiting for a fixed duration.
// use StopEventLoopNoWait to immediate shutdown EventLoop without wait tasks execution.
// use StopEventLoop wait EventLoop after current tasks execution.
type EventLoop struct {
	engine   *Engine
	jobChan  chan func()
	jobCount int32
	canRun   int32 // 1 can run ,2. stop wait all jobs 0. stop immediate.

	auxJobsLock sync.Mutex
	wakeupChan  chan struct{}

	auxJobsSpare, auxJobs []func()

	stopLock sync.Mutex
	stopCond *sync.Cond
	running  bool

	async        *AsyncContext
	asyncResumed bool
	asyncRefs    int
}
type AsyncContext struct {
	ref int
}

func (s *AsyncContext) String() string {
	return fmt.Sprintf("async<%p>[%d]", s, s.ref)
}
func (s *EventLoop) Grab() (trackingObject any) {
	ctx := s.async
	if ctx != nil {
		ctx.ref++
	} else {
		ctx = &AsyncContext{ref: 1}
		s.asyncRefs++
	}
	s.jobCount++
	//println("grab", ctx.String(), s.jobCount)
	return ctx
}

func (s *EventLoop) Resumed(trackingObject any) {
	//println("resumed", trackingObject, s.jobCount)
	if s.asyncResumed {
		panic("nested async context resumed calls")
	}
	s.async = trackingObject.(*AsyncContext)
	s.asyncResumed = true

}
func (s *EventLoop) asyncRelease() {
	s.async.ref--
	if s.async.ref < 0 {
		panic("async context reference negative")
	}
	if s.async.ref == 0 {
		s.asyncRefs--
		if s.asyncRefs < 0 {
			panic("async context refs negative")
		}
		s.jobCount--
		//println("release context", s.async, s.jobCount)
	}
}
func (s *EventLoop) Exited() {
	//println("exited", s.async, s.jobCount)
	if s.async != nil {
		s.asyncRelease()
		s.async = nil
	}
	s.asyncResumed = false
}

func NewEventLoop(engine *Engine) *EventLoop {
	loop := &EventLoop{
		engine:     engine,
		jobChan:    make(chan func()),
		wakeupChan: make(chan struct{}, 1),
	}
	loop.stopCond = sync.NewCond(&loop.stopLock)
	engine.Set("setTimeout", loop.setTimeout)
	engine.Set("setInterval", loop.setInterval)
	engine.Set("setImmediate", loop.setImmediate)
	engine.Set("clearTimeout", loop.clearTimeout)
	engine.Set("clearInterval", loop.clearInterval)
	engine.Set("clearImmediate", loop.clearImmediate)
	engine.Runtime.SetAsyncContextTracker(loop)
	return loop
}

func (s *EventLoop) schedule(call goja.FunctionCall, repeating bool) goja.Value {
	if ff, ok := goja.AssertFunction(call.Argument(0)); ok {
		delay := call.Argument(1).ToInteger()
		var args []goja.Value
		if len(call.Arguments) > 2 {
			args = append(args, call.Arguments[2:]...)
		}
		f := func() { fn.Panic1(ff(nil, args...)) }
		s.jobCount++
		if repeating {
			return s.engine.ToValue(s.addInterval(f, time.Duration(delay)*time.Millisecond))
		} else {
			return s.engine.ToValue(s.addTimeout(f, time.Duration(delay)*time.Millisecond))
		}
	}
	return nil
}

func (s *EventLoop) setTimeout(call goja.FunctionCall) goja.Value {
	return s.schedule(call, false)
}

func (s *EventLoop) setInterval(call goja.FunctionCall) goja.Value {
	return s.schedule(call, true)
}

func (s *EventLoop) setImmediate(call goja.FunctionCall) goja.Value {
	if ff, ok := goja.AssertFunction(call.Argument(0)); ok {
		var args []goja.Value
		if len(call.Arguments) > 1 {
			args = append(args, call.Arguments[1:]...)
		}
		f := func() { fn.Panic1(ff(nil, args...)) }
		s.jobCount++
		return s.engine.ToValue(s.addImmediate(f))
	}
	return nil
}

func (s *EventLoop) setRunning() {
	s.stopLock.Lock()
	defer s.stopLock.Unlock()
	if s.running {
		panic("Loop is already started")
	}
	s.running = true
	atomic.StoreInt32(&s.canRun, 1)
}

func (s *EventLoop) SetTimeout(fn func(engine *Engine), timeout time.Duration) *Timer {
	t := s.addTimeout(func() { fn(s.engine) }, timeout)
	s.addAuxJob(func() {
		s.jobCount++
	})
	return t
}

func (s *EventLoop) ClearTimeout(t *Timer) {
	s.addAuxJob(func() {
		s.clearTimeout(t)
	})
}

func (s *EventLoop) SetInterval(fn func(engine *Engine), timeout time.Duration) *Interval {
	i := s.addInterval(func() { fn(s.engine) }, timeout)
	s.addAuxJob(func() {
		s.jobCount++
	})
	return i
}

func (s *EventLoop) ClearInterval(i *Interval) {
	s.addAuxJob(func() {
		s.clearInterval(i)
	})
}
func (s *EventLoop) Running() bool {
	return s.running
}
func (s *EventLoop) Run(fn func(engine *Engine)) {
	s.setRunning()
	fn(s.engine)
	s.run(false)
}
func (s *EventLoop) registerCallback() func(func()) {
	return func(f func()) {
		s.addAuxJob(f)
	}
}
func (s *EventLoop) StartEventLoop() {
	s.setRunning()
	go s.run(true)
}
func (s *EventLoop) StartEventLoopWhenNotStarted() {
	if s.running {
		return
	}
	s.setRunning()
	go s.run(true)
}

// StopEventLoopContext stop with context for a cancelable or with deadline context
func (s *EventLoop) StopEventLoopContext(ctx context.Context) {
	dead, ok := ctx.Deadline()
	if ok {
		go func() {
			tick := time.Tick(dead.Sub(time.Now()))
			for s.running {
				select {
				case <-tick:
					s.StopEventLoopNoWait()
					return
				default:
					atomic.StoreInt32(&s.canRun, 2) //stop when no task , not use lock
					s.wakeup()
				}
			}
		}()
	} else {
		go func() {
			for s.running {
				select {
				case <-ctx.Done():
					s.StopEventLoopNoWait()
				default:
					atomic.StoreInt32(&s.canRun, 2) //stop when no task , not use lock
					s.wakeup()
				}
			}
		}()
	}

}

// StopEventLoopForContext stop with notify a context
func (s *EventLoop) StopEventLoopForContext() context.Context {
	ctx, cc := context.WithCancel(context.Background())
	go func() {
		for s.running {
			select {
			case <-ctx.Done():
				s.StopEventLoopNoWait()
			default:
				atomic.StoreInt32(&s.canRun, 2) //stop when no task , not use lock
				s.wakeup()
			}
		}
		cc()
	}()
	return ctx
}

// StopEventLoopWait all job done
func (s *EventLoop) StopEventLoopWait() int {
	for s.running {
		atomic.StoreInt32(&s.canRun, 2) //stop when no task , not use lock
		s.wakeup()
	}
	return int(s.jobCount)
}

// StopEventLoopTimeout all job done in a limit duration, the returning value may not exactly the pending job count,for no lock to wait
func (s *EventLoop) StopEventLoopTimeout(duration time.Duration) int {
	tick := time.Tick(duration)
	for s.running {
		atomic.StoreInt32(&s.canRun, 2) //stop when no task
		select {
		case <-tick:
			return s.StopEventLoop()
		default:
			s.wakeup()
		}
	}
	s.stopLock.Lock()
	s.stopCond.Wait()
	s.stopLock.Unlock()
	return int(s.jobCount)
}

// StopEventLoop wait background execution quit , returns job not executed
func (s *EventLoop) StopEventLoop() int {
	s.stopLock.Lock()
	for s.running {
		atomic.StoreInt32(&s.canRun, 0)
		s.wakeup()
		s.stopCond.Wait()
	}
	s.stopLock.Unlock()
	return int(s.jobCount)
}

// StopEventLoopNoWait without wait for background execution finish
func (s *EventLoop) StopEventLoopNoWait() {
	s.stopLock.Lock()
	if s.running {
		atomic.StoreInt32(&s.canRun, 0)
		s.wakeup()
	}
	s.stopLock.Unlock()
}

func (s *EventLoop) RunOnLoop(fn func(*Engine)) {
	s.addAuxJob(func() { fn(s.engine) })
}

func (s *EventLoop) runAux() {
	s.auxJobsLock.Lock()
	jobs := s.auxJobs
	s.auxJobs = s.auxJobsSpare
	s.auxJobsLock.Unlock()
	for i, job := range jobs {
		job()
		jobs[i] = nil
	}
	s.auxJobsSpare = jobs[:0]
}

func (s *EventLoop) run(inBackground bool) {
	s.runAux()
	e := int32(0)
	if inBackground {
		s.jobCount++
		e++
	}

LOOP:
	for s.jobCount > 0 {
		select {
		case job := <-s.jobChan:
			job()
		case <-s.wakeupChan:
			s.runAux()
			v := atomic.LoadInt32(&s.canRun)
			if v == 0 || (v == 2 && s.jobCount == e) {
				break LOOP
			}
		}
	}
	if inBackground {
		s.jobCount--
	}

	s.stopLock.Lock()
	s.running = false
	s.stopLock.Unlock()
	s.stopCond.Broadcast()
}

func (s *EventLoop) wakeup() {
	select {
	case s.wakeupChan <- struct{}{}:
	default:
	}
}

func (s *EventLoop) addAuxJob(fn func()) {
	s.auxJobsLock.Lock()
	s.auxJobs = append(s.auxJobs, fn)
	s.auxJobsLock.Unlock()
	s.wakeup()
}

func (s *EventLoop) addTimeout(f func(), timeout time.Duration) *Timer {
	t := &Timer{
		job: job{fn: f},
	}
	t.timer = time.AfterFunc(timeout, func() {
		s.jobChan <- func() {
			s.doTimeout(t)
		}
	})

	return t
}

func (s *EventLoop) addInterval(f func(), timeout time.Duration) *Interval {
	// https://nodejs.org/api/timers.html#timers_setinterval_callback_delay_args
	if timeout <= 0 {
		timeout = time.Millisecond
	}

	i := &Interval{
		job:      job{fn: f},
		ticker:   time.NewTicker(timeout),
		stopChan: make(chan struct{}),
	}

	go i.run(s)
	return i
}

func (s *EventLoop) addImmediate(f func()) *Immediate {
	i := &Immediate{
		job: job{fn: f},
	}
	s.addAuxJob(func() {
		s.doImmediate(i)
	})
	return i
}

func (s *EventLoop) doTimeout(t *Timer) {
	if !t.cancelled {
		t.fn()
		t.cancelled = true
		s.jobCount--
	}
}

func (s *EventLoop) doInterval(i *Interval) {
	if !i.cancelled {
		i.fn()
	}
}

func (s *EventLoop) doImmediate(i *Immediate) {
	if !i.cancelled {
		i.fn()
		i.cancelled = true
		s.jobCount--
	}
}

func (s *EventLoop) clearTimeout(t *Timer) {
	if t != nil && !t.cancelled {
		t.timer.Stop()
		t.cancelled = true
		s.jobCount--
	}
}

func (s *EventLoop) clearInterval(i *Interval) {
	if i != nil && !i.cancelled {
		i.cancelled = true
		close(i.stopChan)
		s.jobCount--
	}
}

func (s *EventLoop) clearImmediate(i *Immediate) {
	if i != nil && !i.cancelled {
		i.cancelled = true
		s.jobCount--
	}
}

func (i *Interval) run(loop *EventLoop) {
L:
	for {
		select {
		case <-i.stopChan:
			i.ticker.Stop()
			break L
		case <-i.ticker.C:
			loop.jobChan <- func() {
				loop.doInterval(i)
			}
		}
	}
}
