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

// EventLoop process all async events and also  as The AsyncContextTracker
// there will be a background goroutine to monitor the async tasks as call StartEventLoop,
// or call TryStartEventLoop for not sure current EventLoop have been started.
// use Await to await all async tasks are finished.
// use AwaitTimeout to waiting for a fixed duration.
// use StopEventLoopNoWait to immediate shutdown EventLoop without wait tasks execution.
// use StopEventLoop wait EventLoop after current tasks execution.
type EventLoop struct {
	engine   *Engine
	jobChan  chan func()
	jobCount int32 // include real job and AsyncTracker task
	canRun   int32 // 1 can run ,2. stop wait all jobs 0. stop immediate.

	auxJobsLock sync.Mutex
	wakeupChan  chan struct{}

	auxJobsSpare, auxJobs []func()

	stopCond sync.Cond
	running  bool
	tracker  *AsyncTracker
}

func (e *EventLoop) jobInc() {
	debugN(2, "job inc ", e.jobCount)
	e.jobCount++
}
func (e *EventLoop) jobDec() {
	debugN(2, "job dec ", e.jobCount)
	e.jobCount--
	e.stopCond.Broadcast() //avoid deadlock
}

func NewEventLoop(engine *Engine) *EventLoop {
	loop := &EventLoop{
		engine:     engine,
		jobChan:    make(chan func()),
		wakeupChan: make(chan struct{}, 1),
	}
	loop.stopCond = sync.Cond{L: new(sync.Mutex)}
	loop.tracker = new(AsyncTracker)
	loop.tracker.jobInc = loop.jobInc
	loop.tracker.jobDec = loop.jobDec

	engine.Set("setTimeout", loop.setTimeout)
	engine.Set("setInterval", loop.setInterval)
	engine.Set("setImmediate", loop.setImmediate)
	engine.Set("clearTimeout", loop.clearTimeout)
	engine.Set("clearInterval", loop.clearInterval)
	engine.Set("clearImmediate", loop.clearImmediate)
	engine.Runtime.SetAsyncContextTracker(loop.tracker)
	return loop
}

// RunOnLoop  run function on event loop
func (e *EventLoop) RunOnLoop(fn func(*Engine)) {
	e.addAuxJob(func() { fn(e.engine) })
}
func (e *EventLoop) SetTimeout(fn func(engine *Engine), timeout time.Duration) *Timer {
	t := e.addTimeout(func() { fn(e.engine) }, timeout)
	e.addAuxJob(func() {
		e.jobInc()
	})
	return t
}

func (e *EventLoop) ClearTimeout(t *Timer) {
	e.addAuxJob(func() {
		e.clearTimeout(t)
	})
}

func (e *EventLoop) SetInterval(fn func(engine *Engine), timeout time.Duration) *Interval {
	i := e.addInterval(func() { fn(e.engine) }, timeout)
	e.addAuxJob(func() {
		e.jobInc()
	})
	return i
}

func (e *EventLoop) ClearInterval(i *Interval) {
	e.addAuxJob(func() {
		e.clearInterval(i)
	})
}

// StartEventLoop fail if already started
func (e *EventLoop) StartEventLoop() {
	e.setRunning()
	go e.run(true)
}

// TryStartEventLoop no effect if already started
func (e *EventLoop) TryStartEventLoop() {
	e.stopCond.L.Lock()
	if e.running {
		debugN(2, "loop already running", e.running)
		e.stopCond.L.Unlock()
		return
	}
	e.stopCond.L.Unlock()
	debugN(2, "event loop start")
	e.setRunning()
	go e.run(true)
}

// StopEventLoop wait background execution quit (not all tasks) , returns job not executed
func (e *EventLoop) StopEventLoop() HaltJobs {
	e.stopCond.L.Lock()
	for e.running {
		atomic.StoreInt32(&e.canRun, 0)
		e.wakeup()
		e.stopCond.Wait()
	}
	e.stopCond.L.Unlock()
	return HaltJobs{int(e.jobCount), e.tracker.asyncContexts}
}

// StopEventLoopNoWait without wait for background execution finish,
// the result may not accuracy values
func (e *EventLoop) StopEventLoopNoWait() HaltJobs {
	e.stopCond.L.Lock()
	if e.running {
		atomic.StoreInt32(&e.canRun, 0)
		e.wakeup()
	}
	e.stopCond.L.Unlock()
	return HaltJobs{int(e.jobCount), e.tracker.asyncContexts}
}

// Await all job done, see also AwaitContext, AwaitTimeout and AwaitWithContext .
func (e *EventLoop) Await() HaltJobs {
	e.stopCond.L.Lock()
	for e.running {
		//println(e.jobCount, e.tracker.asyncContexts, len(e.auxJobs))
		atomic.StoreInt32(&e.canRun, 2) //stop when no task , not use lock
		e.wakeup()
		//println(e.jobCount, e.tracker.asyncContexts, len(e.auxJobs))
		e.stopCond.Wait()
	}
	e.stopCond.L.Unlock()
	return HaltJobs{int(e.jobCount), e.tracker.asyncContexts}
}

// AwaitWithContext stop with context for a cancelable or with deadline context, see also Await, AwaitTimeout and AwaitContext.
// should use goroutine to execute script to avoid blocking current thread
func (e *EventLoop) AwaitWithContext(ctx context.Context) HaltJobs {
	e.stopCond.L.Lock()
	defer e.stopCond.L.Unlock()
	debug("AwaitWithContext", ctx, "running", e.running)
	err := waitCond(ctx, &e.stopCond, func() {
		atomic.StoreInt32(&e.canRun, 2)
		debug("wakeup")
		e.wakeup()
	}, func() bool {
		return !e.running
	})
	debug("AwaitWithContext", "done", ctx, e.running)
	if err != nil {
		//wait shutdown task routine
		atomic.StoreInt32(&e.canRun, 0)
		debug("await for stop ", ctx, e.running)
		for e.running {
			e.wakeup()
			e.stopCond.Wait()
		}
		debug("done wait for stop ", ctx, e.running)
	}
	return HaltJobs{int(e.jobCount), e.tracker.asyncContexts}
}

func (e *EventLoop) registerCallback() func(func()) {
	return func(f func()) {
		e.addAuxJob(f)
	}
}
func (e *EventLoop) schedule(call goja.FunctionCall, repeating bool) goja.Value {
	if ff, ok := goja.AssertFunction(call.Argument(0)); ok {
		delay := call.Argument(1).ToInteger()
		var args []goja.Value
		if len(call.Arguments) > 2 {
			args = append(args, call.Arguments[2:]...)
		}
		f := func() { fn.Panic1(ff(nil, args...)) }
		e.jobInc()
		if repeating {
			return e.engine.ToValue(e.addInterval(f, time.Duration(delay)*time.Millisecond))
		} else {
			return e.engine.ToValue(e.addTimeout(f, time.Duration(delay)*time.Millisecond))
		}
	}
	return nil
}

func (e *EventLoop) setTimeout(call goja.FunctionCall) goja.Value {
	return e.schedule(call, false)
}

func (e *EventLoop) setInterval(call goja.FunctionCall) goja.Value {
	return e.schedule(call, true)
}

func (e *EventLoop) setImmediate(call goja.FunctionCall) goja.Value {
	if ff, ok := goja.AssertFunction(call.Argument(0)); ok {
		var args []goja.Value
		if len(call.Arguments) > 1 {
			args = append(args, call.Arguments[1:]...)
		}
		f := func() { fn.Panic1(ff(nil, args...)) }
		e.jobInc()
		return e.engine.ToValue(e.addImmediate(f))
	}
	return nil
}

func (e *EventLoop) setRunning() {
	e.stopCond.L.Lock()
	defer e.stopCond.L.Unlock()
	if e.running {
		panic("Loop is already started")
	}
	e.running = true
	atomic.StoreInt32(&e.canRun, 1)
}
func (e *EventLoop) runAux() {
	e.auxJobsLock.Lock()
	jobs := e.auxJobs
	e.auxJobs = e.auxJobsSpare
	e.auxJobsLock.Unlock()
	for i, job := range jobs {
		job()
		jobs[i] = nil
	}
	e.auxJobsSpare = jobs[:0]
}

func (e *EventLoop) run(inBackground bool) {
	e.runAux()
	be := int32(0)
	if inBackground {
		e.jobInc()
		be++
	}
LOOP:
	for e.jobCount > 0 {
		select {
		case job := <-e.jobChan:
			debug("executor job", "jobs:", e.jobCount, "initial jobs:", be)
			job()
		case <-e.wakeupChan:
			e.runAux()
			v := atomic.LoadInt32(&e.canRun)
			debug("wakeup executor", "flag:", v, "jobs:", e.jobCount, "initial jobs:", be)
			if v == 0 || (v == 2 && e.jobCount == be) {
				break LOOP
			}
		}
	}
	if inBackground {
		e.jobCount--
	}

	e.stopCond.L.Lock()
	e.running = false
	e.stopCond.L.Unlock()
	e.stopCond.Broadcast()
	debug("shutdown executor")
}

func (e *EventLoop) wakeup() {
	select {
	case e.wakeupChan <- struct{}{}:
	default:
	}
}

func (e *EventLoop) addAuxJob(fn func()) {
	e.auxJobsLock.Lock()
	e.auxJobs = append(e.auxJobs, fn)
	e.auxJobsLock.Unlock()
	e.wakeup()
}

func (e *EventLoop) addTimeout(f func(), timeout time.Duration) *Timer {
	t := &Timer{
		job: job{fn: f},
	}
	t.timer = time.AfterFunc(timeout, func() {
		e.jobChan <- func() {
			e.doTimeout(t)
		}
	})

	return t
}

func (e *EventLoop) addInterval(f func(), timeout time.Duration) *Interval {
	// https://nodejs.org/api/timers.html#timers_setinterval_callback_delay_args
	if timeout <= 0 {
		timeout = time.Millisecond
	}

	i := &Interval{
		job:      job{fn: f},
		ticker:   time.NewTicker(timeout),
		stopChan: make(chan struct{}),
	}

	go i.run(e)
	return i
}

func (e *EventLoop) addImmediate(f func()) *Immediate {
	i := &Immediate{
		job: job{fn: f},
	}
	e.addAuxJob(func() {
		e.doImmediate(i)
	})
	return i
}

func (e *EventLoop) doTimeout(t *Timer) {
	if !t.cancelled {
		t.fn()
		t.cancelled = true
		e.jobDec()
	}
}

func (e *EventLoop) doInterval(i *Interval) {
	if !i.cancelled {
		i.fn()
	}
}

func (e *EventLoop) doImmediate(i *Immediate) {
	if !i.cancelled {
		i.fn()
		i.cancelled = true
		e.jobDec()
	}
}

func (e *EventLoop) clearTimeout(t *Timer) {
	if t != nil && !t.cancelled {
		t.timer.Stop()
		t.cancelled = true
		e.jobDec()
	}
}

func (e *EventLoop) clearInterval(i *Interval) {
	if i != nil && !i.cancelled {
		i.cancelled = true
		close(i.stopChan)
		e.jobDec()
	}
}

func (e *EventLoop) clearImmediate(i *Immediate) {
	if i != nil && !i.cancelled {
		i.cancelled = true
		e.jobDec()
	}
}

type job struct {
	cancelled bool
	fn        func()
}

type Timer struct {
	job
	timer *time.Timer
}

type Immediate struct {
	job
}

type Interval struct {
	job
	ticker   *time.Ticker
	stopChan chan struct{}
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

type AsyncContext struct {
	ref int
}

func (s *AsyncContext) String() string {
	return fmt.Sprintf("async<%p>[%d]", s, s.ref)
}

type AsyncTracker struct {
	async         *AsyncContext
	asyncResumed  bool
	asyncContexts int
	jobDec        func()
	jobInc        func()
}

func (s *AsyncTracker) Grab() (trackingObject any) {
	ctx := s.async
	if ctx != nil {
		ctx.ref++
	} else {
		ctx = &AsyncContext{ref: 1}
		s.asyncContexts++
	}
	s.jobInc()
	//println("grab", ctx.String(), s.jobCount)
	return ctx
}

func (s *AsyncTracker) Resumed(trackingObject any) {
	//println("resumed", trackingObject, s.jobCount)
	if s.asyncResumed {
		panic("nested async context resumed calls")
	}
	s.async = trackingObject.(*AsyncContext)
	s.asyncResumed = true

}
func (s *AsyncTracker) asyncRelease() {
	s.async.ref--
	if s.async.ref < 0 {
		panic("async context reference negative")
	}
	if s.async.ref == 0 {
		s.asyncContexts--
		if s.asyncContexts < 0 {
			panic("async context refs negative")
		}
		s.jobDec()
		//println("release context", s.async, s.jobCount)
	}
}
func (s *AsyncTracker) Exited() {
	//println("exited", s.async, s.jobCount)
	if s.async != nil {
		s.asyncRelease()
		s.async = nil
	}
	s.asyncResumed = false
}

type HaltJobs struct {
	Job       int // Job remains after execution, include async task
	AsyncTask int // AsyncTask remains after execution
}

// IsZero dose all job are done
func (s HaltJobs) IsZero() bool {
	return s.Job == 0
}

// RealJob real job in queue
func (s HaltJobs) RealJob() int {
	return s.Job - s.AsyncTask
}
func (s HaltJobs) String() string {
	return fmt.Sprintf("HaltJobs{jobs:%d,asyncTask:%d}", s.Job, s.AsyncTask)
}

// waitCond https://pkg.go.dev/context#example-AfterFunc-Cond
func waitCond(ctx context.Context, cond *sync.Cond, waitJob func(), condition func() bool) error {
	df := context.AfterFunc(ctx, func() {
		cond.L.Lock()
		defer cond.L.Unlock()
		cond.Broadcast()
	})
	defer df()
	debug("WaitCond", "before loop")
	for !condition() {
		debug("WaitCond", "job")
		waitJob() //the order is important
		cond.Wait()
		debug("WaitCond", "after await")
		if ctx.Err() != nil {
			debug("WaitCond", "context dead")
			return ctx.Err()
		}
	}
	debug("WaitCond", "done wait loop")
	return nil
}
