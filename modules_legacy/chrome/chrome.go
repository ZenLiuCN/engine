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
func (c *Chrome) Close() error {
	return cd.Cancel(c.ctx)
}
func (c *Chrome) Targets() Maybe[[]*target.Info] {
	return MaybeBoth(cd.Targets(c.ctx))
}

// Shutdown close chrome without await for full close
func (c *Chrome) Shutdown() {
	c.cc()
}

func (c *Chrome) Submit(act ...cd.Action) error {
	return cd.Run(c.ctx, act...)
}
func (c *Chrome) CreateBrowser(url string, opts ...cd.BrowserOption) Maybe[*Browser] {
	if v, e := cd.NewBrowser(c.ctx, url, opts...); e != nil {
		return MaybeError[*Browser](e)
	} else {
		return MaybeOk(&Browser{v})
	}

}
func NewChromeDefault(opt ...cd.ContextOption) *Chrome {
	ctx, cc := cd.NewContext(context.Background(), opt...)
	return &Chrome{ctx, cc}
}

func NewChromeUrl(url string, remotes []cd.RemoteAllocatorOption, opt ...cd.ContextOption) *Chrome {
	ctx, cc := cd.NewRemoteAllocator(context.Background(), url, remotes...)
	ctx1, _ := cd.NewContext(ctx, opt...)
	return &Chrome{ctx1, cc}
}
func NewChromeOptions(execOption []cd.ExecAllocatorOption, opt ...cd.ContextOption) *Chrome {
	ctx, cc := cd.NewExecAllocator(context.Background(), execOption...)
	ctx1, _ := cd.NewContext(ctx, opt...)
	return &Chrome{ctx1, cc}
}

type Browser struct {
	*cd.Browser
}
