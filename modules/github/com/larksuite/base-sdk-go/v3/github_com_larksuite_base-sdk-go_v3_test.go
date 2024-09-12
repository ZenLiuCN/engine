package lark

import (
	"github.com/ZenLiuCN/engine"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/base-sdk-go/v3/service/base/v1"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestSimple(t *testing.T) {
	v := engine.Get()
	defer v.Free()
	fn.Panic1(v.RunTs(
		//language=typescript
		`
		import * as lark from 'github.com/larksuite/base-sdk-go/v3'
		import * as btv1 from 'github.com/larksuite/base-sdk-go/v3/service/base/v1'
		import * as context from 'golang/context'
		const fields=["任务","人员","项目","需求","资料","BUG","状态","说明","需要协助的问题"]
		const cli=lark.newClient( "","")
		const ref=btv1.newListAppTableRecordReqBuilder()
		.appToken('')
		.tableId('')
		.viewId('')
		.fieldNames('["任务","人员","项目","需求","资料","BUG","状态","说明","需要协助的问题"]')
		.build()
		const res=cli.base.appTableRecord.list(context.background(),ref)
		const result=res.data.items.map(it=>fields.map(x=>it.fields[x]))
		console.log("res",result)
		`))
}
