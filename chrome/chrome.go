package chrome

import (
	"context"
	. "github.com/ZenLiuCN/engine"
	"github.com/chromedp/cdproto/target"
	cd "github.com/chromedp/chromedp"
)

type ValueAction[T any] struct {
	value T
	cd.Action
}

func (c *ValueAction[T]) Value() T {
	return c.value
}

type Chrome struct {
	ctx context.Context
	cc  context.CancelFunc
}

// Close the chrome application
func (c *Chrome) Close() *GoError {
	return GoErrorOf(cd.Cancel(c.ctx))
}

// Shutdown close chrome without await for full close
func (c *Chrome) Shutdown() {
	c.cc()
}

func (c *Chrome) Submit(act ...cd.Action) *GoError {
	return GoErrorOf(cd.Run(c.ctx, act...))
}
func (c *Chrome) CreateBrowser(url string, opts ...cd.BrowserOption) Maybe[*Browser] {
	if v, e := cd.NewBrowser(c.ctx, url, opts...); e != nil {
		return MaybeError[*Browser](e)
	} else {
		return MaybeOk(&Browser{v})
	}

}
func NewChromeDefault() *Chrome {
	ctx, cc := cd.NewContext(context.Background())
	return &Chrome{ctx, cc}
}
func NewChromeTarget(targetId string) *Chrome {
	ctx, cc := cd.NewContext(context.Background(), cd.WithTargetID(target.ID(targetId)))
	return &Chrome{ctx, cc}
}

func NewChromeOptions(options ...cd.ExecAllocatorOption) *Chrome {
	ctx, cc := cd.NewExecAllocator(context.Background(), options...)
	return &Chrome{ctx, cc}
}

type Browser struct {
	*cd.Browser
}
