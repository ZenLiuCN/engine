package pdf

import (
	"github.com/ZenLiuCN/engine"
	_ "github.com/ZenLiuCN/engine/modules/golang/strings"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestSimple(t *testing.T) {
	v := engine.Get()
	defer v.Free()
	fn.Panic1(v.RunTs(
		//language=typescript
		`
				import * as pdf from 'github.com/dslipak/pdf'
				import * as strings from 'golang/strings'
				import {typeOf} from 'go'
				let r=pdf.open("./testdata/sample.pdf") 
				const page=r.numPage()
				console.log('total',page)
				const bu=strings.refBuilder()
				console.log(typeOf(bu).identity())
				for(let i=1;i <=page;i++){
					const texts=r.page(i).content().text
					let y=0
					texts.forEach(x=>{
                        if (x.Y>y){
                            y=x.Y
                            console.log(bu.string())
                            bu.reset()
                        }
                        bu.writeString(x.S)
					})
				}
              
				`))
}
