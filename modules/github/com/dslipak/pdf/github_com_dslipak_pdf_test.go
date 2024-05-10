package pdf

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestSimple(t *testing.T) {
	v := engine.Get()
	defer v.Free()
	fn.Panic1(v.RunTs(
		//language=typescript
		`
				import * as pdf from 'github/com/dslipak/pdf'
				let r=pdf.open("./testdata/sample.pdf") 
				const page=r.numPage()
				console.log('total',page)
				for(let i=1;i <=page;i++){
					const v=r.page(i)
					console.log(v)
					const content=v.content()
					console.log(content.text.length)
					const texts=content.text
					texts.forEach(x=>console.log(x.S))
				}
				`))
}
