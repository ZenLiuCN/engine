package chrome

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/css"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/target"
	cd "github.com/chromedp/chromedp"
	"github.com/dop251/goja"
)

var (
	//go:embed chrome.d.ts
	chromeDefine []byte
)

func init() {
	engine.RegisterModule(ModuleChrome{})
}

type ModuleChrome struct {
}

func (c ModuleChrome) TypeDefine() []byte {
	return chromeDefine
}

func (c ModuleChrome) Identity() string {
	return "go/chrome"
}

func (c ModuleChrome) Exports() map[string]any {
	return nil
}

func (c ModuleChrome) ExportsWithEngine(eng *engine.Engine) map[string]any {
	return map[string]any{
		//convert
		"toBrowserContextID": func(id string) cdp.BrowserContextID {
			return cdp.BrowserContextID(id)
		},
		"toTargetID": func(id string) target.ID {
			return target.ID(id)
		},
		"toFrameID": func(id string) cdp.FrameID {
			return cdp.FrameID(id)
		},
		"toTargetSessionID": func(id string) target.SessionID {
			return target.SessionID(id)
		},
		"toNodeID": func(id int64) cdp.NodeID {
			return cdp.NodeID(id)
		},
		"fromNodeID": func(id cdp.NodeID) int64 {
			return int64(id)
		},
		//context option
		"withTargetID":               cd.WithTargetID,
		"withExistingBrowserContext": cd.WithExistingBrowserContext,
		"withBrowserOption":          cd.WithBrowserOption,
		//"withNewBrowserContext": cd.WithNewBrowserContext,

		//browser option
		"withDialTimeout": cd.WithDialTimeout,
		//exec option
		"execPath":         cd.ExecPath,
		"flag":             cd.Flag,
		"env":              cd.Env,
		"userDataDir":      cd.UserDataDir,
		"proxyServer":      cd.ProxyServer,
		"ignoreCertErrors": cd.IgnoreCertErrors,
		"windowSize":       cd.WindowSize,
		"userAgent":        cd.UserAgent,
		"noSandbox":        cd.NoSandbox,
		"noFirstRun":       cd.NoFirstRun,
		"headless":         cd.Headless,
		"disableGPU":       cd.DisableGPU,
		"noModifyURL":      cd.NoModifyURL,
		"combinedOutput":   cd.CombinedOutput,
		"wsUrlReadTimeout": cd.WSURLReadTimeout,
		//Chrome
		"Chrome": eng.ToConstructor(func(v []goja.Value) any {
			if len(v) == 0 {
				return engine.RegisterResource(eng, NewChromeDefault())
			}
			v0 := v[0]
			if url, ok := v0.Export().(string); ok {
				if len(v) == 1 {
					return engine.RegisterResource(eng, NewChromeUrl(url, nil))
				}
				if ar, ok := v[1].Export().([]any); !ok {
					panic("not an array of RemoteOption")
				} else {
					var remotes []cd.RemoteAllocatorOption
					for _, value := range ar {
						switch t := value.(type) {
						case cd.RemoteAllocatorOption:
							remotes = append(remotes, t)
						default:
							panic("not a RemoteOption")
						}
					}
					var opt []cd.ContextOption
					for i, value := range v {
						if i == 0 || i == 1 {
							continue
						}
						if o, ok := value.Export().(cd.ContextOption); ok {
							opt = append(opt, o)
						} else {
							panic("not an ContextOption")
						}
					}
					return engine.RegisterResource(eng, NewChromeUrl(url, remotes, opt...))
				}
			}
			if o0, ok := v0.Export().(cd.ContextOption); ok {
				opt := []cd.ContextOption{o0}
				for i, value := range v {
					if i == 0 {
						continue
					}
					if o, ok := value.Export().(cd.ContextOption); ok {
						opt = append(opt, o)
					} else {
						panic("not an ContextOption")
					}
				}
				return engine.RegisterResource(eng, NewChromeDefault(opt...))
			}
			if eao, ok := v0.Export().([]cd.ExecAllocatorOption); ok {
				var opt []cd.ContextOption
				for i, value := range v {
					if i == 0 {
						continue
					}
					if o, ok := value.Export().(cd.ContextOption); ok {
						opt = append(opt, o)
					} else {
						panic("not an ContextOption")
					}
				}
				return engine.RegisterResource(eng, NewChromeOptions(eao, opt...))
			}
			panic("invalid arguments")
		}),
		//navigate
		"navigate":        cd.Navigate,
		"navigateBack":    cd.NavigateBack,
		"navigateForward": cd.NavigateForward,
		"reload":          cd.Reload,
		"stopLoading":     cd.Stop,
		"location": ActionWithValueFunc[string](func(v *ValueAction[string]) cd.Action {
			return cd.Location(&v.value)
		}),
		"title": ActionWithValueFunc[string](func(v *ValueAction[string]) cd.Action {
			return cd.Title(&v.value)
		}),
		"navigationEntries": ActionWithValueFunc(func(v *ValueAction[NaviEntries]) cd.Action {
			return cd.NavigationEntries(&v.value.CurrentId, &v.value.entries)
		}),
		"navigateToHistoryEntry": cd.NavigateToHistoryEntry,
		//poll
		"pollingInterval": cd.WithPollingInterval,
		"pollingMutation": cd.WithPollingMutation,
		"pollingTimeout":  cd.WithPollingTimeout,
		"pollingInFrame":  cd.WithPollingInFrame,
		"pollingArgs":     cd.WithPollingArgs,
		"poll":            cd.Poll,
		"pollFunction":    cd.PollFunction,
		//Query
		"FromNode":       cd.FromNode,
		"ByFunc":         cd.ByFunc,
		"ByQuery":        cd.ByQuery,
		"ByQueryAll":     cd.ByQueryAll,
		"ByID":           cd.ByID,
		"BySearch":       cd.BySearch,
		"ByJSPath":       cd.ByJSPath,
		"ByNodeID":       cd.ByNodeID,
		"NodeReady":      cd.NodeReady,
		"NodeVisible":    cd.NodeVisible,
		"NodeNotVisible": cd.NodeNotVisible,
		"NodeEnabled":    cd.NodeEnabled,
		"NodeSelected":   cd.NodeSelected,
		"NodeNotPresent": cd.NodeNotPresent,
		"fromNode":       cd.FromNode,
		"atLeast":        cd.AtLeast,
		"retryInterval":  cd.RetryInterval,
		"waitReady":      cd.WaitReady,
		"waitVisible":    cd.WaitVisible,
		"waitNotVisible": cd.WaitNotVisible,
		"waitEnabled":    cd.WaitEnabled,
		"waitSelected":   cd.WaitSelected,
		"waitNotPresent": cd.WaitNotPresent,
		"query":          cd.Query,
		"focus":          cd.Focus,
		"blur":           cd.Blur,
		"clear":          cd.Clear,
		"click":          cd.Click,
		"doubleClick":    cd.DoubleClick,
		"submit":         cd.Submit,
		"reset":          cd.Reset,
		"scrollIntoView": cd.ScrollIntoView,
		"nodes": func(sel string, opts ...cd.QueryOption) *ValueAction[[]*cdp.Node] {
			return ActionWithValue(func(v *ValueAction[[]*cdp.Node]) cd.Action {
				return cd.Nodes(sel, &v.value, opts...)
			})
		},
		"nodeIDs": func(sel string, opts ...cd.QueryOption) *ValueAction[[]cdp.NodeID] {
			return ActionWithValue(func(v *ValueAction[[]cdp.NodeID]) cd.Action {
				return cd.NodeIDs(sel, &v.value, opts...)
			})
		},
		"outerHTML": func(sel string, opts ...cd.QueryOption) *ValueAction[string] {
			return ActionWithValue(func(v *ValueAction[string]) cd.Action {
				return cd.OuterHTML(sel, &v.value, opts...)
			})
		},
		"innerHTML": func(sel string, opts ...cd.QueryOption) *ValueAction[string] {
			return ActionWithValue(func(v *ValueAction[string]) cd.Action {
				return cd.InnerHTML(sel, &v.value, opts...)
			})
		},
		"value": func(sel string, opts ...cd.QueryOption) *ValueAction[string] {
			return ActionWithValue(func(v *ValueAction[string]) cd.Action {
				return cd.Value(sel, &v.value, opts...)
			})
		},
		"text": func(sel string, opts ...cd.QueryOption) *ValueAction[string] {
			return ActionWithValue(func(v *ValueAction[string]) cd.Action {
				return cd.Text(sel, &v.value, opts...)
			})
		},
		"textContent": func(sel string, opts ...cd.QueryOption) *ValueAction[string] {
			return ActionWithValue(func(v *ValueAction[string]) cd.Action {
				return cd.TextContent(sel, &v.value, opts...)
			})
		},
		"dimensions": func(sel string, opts ...cd.QueryOption) *ValueAction[*dom.BoxModel] {
			return ActionWithValue(func(v *ValueAction[*dom.BoxModel]) cd.Action {
				return cd.Dimensions(sel, &v.value, opts...)
			})
		}, "attributes": func(sel string, opts ...cd.QueryOption) *ValueAction[map[string]string] {
			return ActionWithValue(func(v *ValueAction[map[string]string]) cd.Action {
				return cd.Attributes(sel, &v.value, opts...)
			})
		},
		"attributesAll": func(sel string, opts ...cd.QueryOption) *ValueAction[[]map[string]string] {
			return ActionWithValue(func(v *ValueAction[[]map[string]string]) cd.Action {
				return cd.AttributesAll(sel, &v.value, opts...)
			})
		},
		"setValue": func(sel, value string, opts ...cd.QueryOption) cd.Action {
			return cd.SetValue(sel, value, opts...)

		},
		"setAttributes": func(sel string, value map[string]string, opts ...cd.QueryOption) cd.Action {
			return cd.SetAttributes(sel, value, opts...)
		},
		"setAttributeValue": func(sel string, name, value string, opts ...cd.QueryOption) cd.Action {
			return cd.SetAttributeValue(sel, name, value, opts...)
		},
		"removeAttribute": func(sel, value string, opts ...cd.QueryOption) cd.Action {
			return cd.RemoveAttribute(sel, value, opts...)

		},
		"javascriptAttribute": func(sel, name string, opts ...cd.QueryOption) *ValueAction[any] {
			return ActionWithValue(func(v *ValueAction[any]) cd.Action {
				return cd.JavascriptAttribute(sel, name, v.value, opts...)
			})
		},
		"setJavascriptAttribute": func(sel string, name string, value string, opts ...cd.QueryOption) cd.Action {
			return cd.SetJavascriptAttribute(sel, name, value, opts...)

		},
		"sendKeys": func(sel string, name string, opts ...cd.QueryOption) cd.Action {
			return cd.SendKeys(sel, name, opts...)

		},
		"setUploadFiles": func(sel string, name []string, opts ...cd.QueryOption) cd.Action {
			return cd.SetUploadFiles(sel, name, opts...)

		},
		"computedStyle": func(sel string, opts ...cd.QueryOption) *ValueAction[[]*css.ComputedStyleProperty] {
			return ActionWithValue(func(v *ValueAction[[]*css.ComputedStyleProperty]) cd.Action {
				return cd.ComputedStyle(sel, &v.value, opts...)
			})
		},
		"matchedStyle": func(sel string, opts ...cd.QueryOption) *ValueAction[*css.GetMatchedStylesForNodeReturns] {
			return ActionWithValue(func(v *ValueAction[*css.GetMatchedStylesForNodeReturns]) cd.Action {
				return cd.MatchedStyle(sel, &v.value, opts...)
			})
		},
		//Screenshot
		"screenshot": func(sel string, opts ...cd.QueryOption) *ValueAction[[]byte] {
			return ActionWithValue(func(v *ValueAction[[]byte]) cd.Action {
				return cd.Screenshot(sel, &v.value, opts...)
			})
		},
		"captureScreenshot": ActionWithValueFunc(func(v *ValueAction[[]byte]) cd.Action {
			return cd.CaptureScreenshot(&v.value)
		}),
		"fullScreenshot": func(quality int) *ValueAction[[]byte] {
			return ActionWithValue(func(v *ValueAction[[]byte]) cd.Action {
				return cd.FullScreenshot(&v.value, quality)
			})
		},
		//Eval
		"EvalWithCommandLineAPI": cd.EvalWithCommandLineAPI,
		"EvalIgnoreExceptions":   cd.EvalIgnoreExceptions,
		"EvalAsValue":            cd.EvalAsValue,
		"evalObjectGroup":        cd.EvalObjectGroup,
		"evaluate": func(expr string, opts ...cd.EvaluateOption) *ValueAction[any] {
			return ActionWithValue(func(v *ValueAction[any]) cd.Action {
				return cd.Evaluate(expr, &v.value, opts...)
			})
		},
		"evaluateAsDevTools": func(expr string, opts ...cd.EvaluateOption) *ValueAction[any] {
			return ActionWithValue(func(v *ValueAction[any]) cd.Action {
				return cd.EvaluateAsDevTools(expr, &v.value, opts...)
			})
		},
		//Call
		"callFunctionOn": func(function string, args ...any) *ValueAction[any] {
			return ActionWithValue(func(v *ValueAction[any]) cd.Action {
				return cd.CallFunctionOn(function, &v.value, nil, args...)
			})
		},
		//Emulate
	}
}
func ActionWithValueFunc[T any](f func(*ValueAction[T]) cd.Action) func() *ValueAction[T] {

	return func() *ValueAction[T] {
		v := new(ValueAction[T])
		v.Action = f(v)
		return v
	}
}
func ActionWithValue[T any](f func(*ValueAction[T]) cd.Action) *ValueAction[T] {
	v := new(ValueAction[T])
	v.Action = f(v)
	return v
}

type NaviEntries struct {
	CurrentId int64
	entries   []*page.NavigationEntry
}
