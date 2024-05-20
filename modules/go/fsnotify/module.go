package fsnotify

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/dop251/goja"
	"github.com/fsnotify/fsnotify"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	//go:embed fs-notify.d.ts
	define []byte
)

func init() {
	engine.RegisterModule(Notifier{})
}

type Notifier struct {
}

func (p Notifier) Identity() string {
	return "go/fsNotify"
}

func (p Notifier) TypeDefine() []byte {
	return define
}

func (p Notifier) Exports() map[string]any {

	return nil
}
func (p Notifier) ExportsWithEngine(eng *engine.Engine) map[string]any {
	var ctor = eng.ToConstructor(func(v []goja.Value) (o any, err error) {
		var w *fsnotify.Watcher
		if len(v) == 0 {
			w, err = fsnotify.NewWatcher()
		} else if len(v) == 1 {
			w, err = fsnotify.NewBufferedWatcher(uint(v[0].ToInteger()))
		}
		if err != nil {
			return
		}
		return &Notify{e: eng, Watcher: w, closer: make(chan struct{})}, nil
	})
	return map[string]any{
		"Watcher": ctor,
	}
}

type Notify struct {
	e *engine.Engine
	*fsnotify.Watcher
	closer  chan struct{}
	tree    []string
	running bool
}

func (n *Notify) AddTree(p string) error {
	return filepath.Walk(p, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			err = n.Add(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
func (n *Notify) RemoveTree(p string) error {
	return filepath.Walk(p, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			err = n.Remove(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
func (n *Notify) OnEvent(fn goja.Value) {
	if n.running {
		return
	}
	n.running = true
	go func(e *engine.Engine) {
		defer func() {
			n.closer <- struct{}{}
		}()
		for {
			select {
			case err, ok := <-n.Errors:
				if !ok {
					return
				}
				_, _ = e.CallFunction(fn, e.Undefined(), err)
			case ev, ok := <-n.Events:
				if !ok {
					return
				}
				_, _ = e.CallFunction(fn, NotifyEvent{ev}, e.Undefined())
			}
		}
	}(n.e)
}
func (n *Notify) Await() {
	<-n.closer
	n.running = false
}

func (n *Notify) Close() error {
	n.tree = nil
	n.closer <- struct{}{}
	return n.Watcher.Close()
}

type NotifyEvent struct {
	fsnotify.Event
}

func (e NotifyEvent) Path() string {
	return e.Name
}
func (e NotifyEvent) IsFile() (bool, error) {
	if e.IsRemoved() {
		return false, nil
	}
	s, err := os.Stat(e.Name)
	if err != nil {
		return false, nil
	}
	return !s.IsDir(), nil
}
func (e NotifyEvent) IsCreated() bool {
	return e.Has(fsnotify.Create)
}
func (e NotifyEvent) IsModified() bool {
	return e.Has(fsnotify.Write)
}
func (e NotifyEvent) IsRenamed() bool {
	return e.Has(fsnotify.Rename)
}
func (e NotifyEvent) IsRemoved() bool {
	return e.Has(fsnotify.Remove)
}
func (e NotifyEvent) IsAttributeChanged() bool {
	return e.Has(fsnotify.Chmod)
}
