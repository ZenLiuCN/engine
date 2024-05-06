package engine

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/evanw/esbuild/pkg/api"
	"strings"
	"testing"
)

func TestJsSourceMap(t *testing.T) {
	s, m := compileTestTs("(() => {\n    const {div, input, button, header, details, summary, dialog, article, a, p, h3, footer, h6, textarea} = van.tags\n\n    const state = vanX.reactive({\n        user: '',\n        pushed: false,\n        token: '',\n        fetched: false,\n        vals: {},\n        actions: {},\n        errors: '',\n        loading: {},\n    })\n\n    const isOpenClass = \"modal-is-open\";\n    const openingClass = \"modal-is-opening\";\n    const closingClass = \"modal-is-closing\";\n    const animationDuration = 400 // ms\n    const isModalOpen = (modal) =>\n        modal.hasAttribute(\"open\") &&\n        modal.getAttribute(\"open\") !== \"false\"\n    const openModal = (modal) => {\n        if (isScrollbarVisible()) {\n            document.documentElement.style.setProperty(\"--scrollbar-width\",\n                `${getScrollbarWidth()}px`)\n        }\n        document.documentElement.classList.add(isOpenClass, openingClass)\n        setTimeout(() => {\n            document.documentElement.classList.remove(openingClass);\n        }, animationDuration)\n        modal.setAttribute(\"open\", true);\n    }\n    const closeModal = (modal) => {\n        document.documentElement.classList.add(closingClass)\n        setTimeout(() => {\n            document.documentElement.classList.remove(closingClass, isOpenClass)\n            document.documentElement.style.removeProperty(\"--scrollbar-width\")\n            modal.removeAttribute(\"open\");\n        }, animationDuration)\n    }\n\n    const getScrollbarWidth = () => {\n        const outer = document.createElement(\"div\")\n        outer.style.visibility = \"hidden\"\n        outer.style.overflow = \"scroll\"\n        outer.style.msOverflowStyle = \"scrollbar\"\n        document.body.appendChild(outer)\n        const inner = document.createElement(\"div\")\n        outer.appendChild(inner);\n        const scrollbarWidth = outer.offsetWidth - inner.offsetWidth\n        outer.parentNode.removeChild(outer)\n        return scrollbarWidth\n    }\n    const isScrollbarVisible = () => document.body.scrollHeight > screen.height\n    const trigger = (e) => {\n        e.preventDefault()\n        const modal = document.getElementById(e.currentTarget.getAttribute(\"data-target\"));\n        typeof modal != \"undefined\" && modal != null && isModalOpen(modal)\n            ? closeModal(modal)\n            : openModal(modal);\n    }\n    const processError=v=>v.text().then(t => {\n        if (v.status === 401) {\n            state.pushed = false\n            state.fetched = false\n            state.vals = {}\n            state.actions = {}\n            state.token = ''\n        }\n        throw new Error(`<h6>${v.statusText?v.statusText:v.status===400?'请求错误':v.status===404?'数据无效':'错误'}(${v.status})</h6><div>${t}</div>`)\n    })\n    const apiCheckJson = v => v.ok ? v.json() :processError(v)\n    const apiCheckText = v => v.ok ? v.text() : processError(v)\n    const apiModalError = err => {\n        state.errors = err.message\n        openModal(document.getElementById('modal'))\n    }\n    const apiModalErrorAct = (act) => {\n        return err => {\n            act()\n            state.errors = err.message\n            openModal(document.getElementById('modal'))\n        }\n    }\n    const app = () => div({class: 'container'},\n        dialog(\n            {id: 'modal'},\n            article({class: 'alert-dialog'},\n                header('操作错误',a({href: '#close', class: 'close', 'data-target': 'modal', onclick: (e) => trigger(e)})),\n                p({innerHTML:() => state.errors}),\n                footer(\n                    a({href: '#cancel', role: 'button', 'data-target': \"modal\", onclick: (e) => trigger(e)}, \"确定\")\n                )\n            )\n        ),\n        header({class: 'flex-center'}, h6({class: 'title'},\n            () => !state.fetched ? '身份验证' : '工具列表')),\n        () => !state.pushed ? div(\n            input({\n                type: 'text',\n                placeholder: '输入用户名',\n                value: () => state.user,\n                oninput: e => state.user = e.target.value\n            }),\n            button({\n                disabled: () => state.user === \"\",\n                onclick: () => fetch(`api/auth/push?name=${encodeURIComponent(state.user)}`)\n                    .then(apiCheckText)\n                    .then(v => state.pushed = true)\n                    .catch(apiModalError)\n\n            }, '获取秘钥')\n        ) : !state.fetched ? div(\n            input({\n                type: 'text',\n                placeholder: '输入秘钥',\n                value: () => state.token,\n                oninput: e => state.token = e.target.value\n            }),\n            button({\n                disabled: () => state.token === \"\",\n                onclick: () => fetch(`api/action/actions?signature=${encodeURIComponent(state.token)}`)\n                    .then(apiCheckJson)\n                    .then(v => {\n                        state.fetched = true\n                        state.actions = v\n                    })\n                    .catch(apiModalError)\n            }, '开始使用')\n        ) : div(\n            Object.keys(state.actions).map(x => details(\n                summary({style: 'padding-left:15px', role: \"button\"}, x),\n                h6({style: 'padding-left:5px'}, '操作参数'),\n                Object.keys(state.actions[x].param).map(p => div(\n                    input({\n                        type: state.actions[x].param[p].kind ?? 'text',\n                        placeholder: p,\n                        value: () => {\n                            const m = state.actions[x].param[p].name\n                            if (!state.vals[x]) state.vals[x] = {}\n                            return state.vals[x][m] ?? ''\n                        },\n                        oninput: e => state.vals[x][state.actions[x].param[p].name] = e.target.value\n                    })\n                )),\n                button({\n                    class: 'outline',\n                    'aria-busy': () => state.loading[x] ?? false,\n                    disabled: Object.keys(state.actions[x].param).length === 0 ? false : () => state.vals[x] === undefined || (Object.keys(state.actions[x].param).length !== 0 && Object.keys(state.vals[x]).length !== Object.keys(state.actions[x].param).length),\n                    onclick: () => {\n                        state.loading[x] = true\n                        fetch(`api/action/${encodeURIComponent(state.actions[x].action)}?${\n                            Object.keys(state.actions[x].param).length === 0 ? \"\" :\n                                Object.keys(state.vals[x]).map(k => `${encodeURIComponent(k)}=${encodeURIComponent(state.vals[x][k])}`).join('&')\n                        }&signature=${encodeURIComponent(state.token)}`)\n                            .then(apiCheckText)\n                            .then(r => {\n                                document.getElementById(x).value = r\n                                state.loading[x] = false\n                            })\n                            .catch(apiModalErrorAct(() => state.loading[x] = false))\n                    }\n                }, `执行`),\n                div(textarea({id: x, class: 'mark', placeholder: '结果', readonly: true})),\n            ))\n        )\n    )\n    van.add(document.body, app())\n})()", true)
	fmt.Printf("%#+v\n%s", NewSourceMap(m), s)
}
func compileTestTs(ts string, entry bool) (string, []byte) {
	format := api.FormatDefault
	if strings.Contains(ts, "import ") {
		format = api.FormatCommonJS
	} else if strings.Contains(ts, "export ") {
		format = api.FormatCommonJS
	} else if strings.Contains(ts, "require ") {
		format = api.FormatCommonJS
	}
	res := api.Transform(ts, api.TransformOptions{
		MinifyWhitespace:  true,
		MinifyIdentifiers: false,
		MinifySyntax:      true,

		KeepNames:   true,
		TreeShaking: api.TreeShakingTrue,

		Target:    api.ES2022,
		Platform:  api.PlatformBrowser,
		Format:    format,
		Sourcemap: api.SourceMapExternal,
		Loader:    api.LoaderTS,
	})
	if res.Errors != nil {
		panic(errors.New("Compile TS error\n" + fn.SliceJoinRune(res.Errors, '\n', func(m api.Message) string {
			return fmt.Sprintf(`%s: %s`, m.Location.LineText, m.Text)
		})))
	}
	if entry && format == api.FormatCommonJS {
		idx := bytes.Index(res.Code, []byte("module.exports=__toCommonJS(stdin_exports);"))
		if idx >= 0 {
			return string(res.Code[idx+43:]), nil
		}
		return string(res.Code), nil
	}
	return string(res.Code), res.Map
}
