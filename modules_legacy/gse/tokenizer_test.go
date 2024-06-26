package gse

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/ZenLiuCN/engine/modules/go/sqlx/mysql_2023-12-22"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "modernc.org/sqlite"
)

func TestTokenizer(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Tokenizer,TagExtractor}from 'go/gse'
const t=new Tokenizer()
t.alphaNum=true
t.skipLog=true
console.log(t)
t.loadDict()
console.log(t.load)
console.log(t.cut("特伦特·奥唐纳导演负责制作",true))
const tag=new TagExtractor()
tag.withGse(t)
tag.loadIdf()
console.log(tag.extractTags("特伦特·奥唐纳导演负责制作",5).map(v=>v.string()).join("/"))
`))
}
