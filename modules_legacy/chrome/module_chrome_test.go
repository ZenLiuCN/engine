package chrome

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestSimple(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	fn.Panic1(
		vm.RunJs(
			//language=javascript
			`
	import * as ch from 'go/chrome'
	const chr=new ch.Chrome('ws://127.0.0.1:65533/')
    console.log(chr.targets())
    chr.shutdown()
`))

}
func TestTarget(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	fn.Panic1(
		vm.RunJs(
			//language=javascript
			`
	import * as ch from 'go/chrome'
	const chr=new ch.Chrome('ws://127.0.0.1:65533/',[],ch.withTargetID(ch.toTargetID('C58889EBAEAB8DAB9FDE9EE19DDAEA19')))
	const selector='#placeholder-wrapper-6AuUB2x7 > div.page-block.root-block > div > div.page-block-children > div > div.render-unit-wrapper > div:nth-child(1) > div.list-wrapper.ordered-list > div.list-children > div > div:nth-child(1) > div.list-wrapper.ordered-list > div.list-children > div > div > div > div > div > div > div > span:nth-child(1)'
    const innerHtml=ch.text(selector)
    const err=chr.submit(innerHtml,ch.click(selector))
    if (err) console.error(err.error())
    else console.log(innerHtml.value())
    chr.shutdown()
`))

}
