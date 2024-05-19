package lark

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
		import * as lark from 'github.com/larksuite/base-sdk-go/v3'
		import * as btv1 from 'github.com/larksuite/base-sdk-go/v3/service/base/v1'
		import * as context from 'golang/context'
		const cli=lark.newClient( "pt-ZS4oNx07xtipF1KMaFRSAah8shVf--7tZP4wv2SPAQAABMCAhxzAAgEI6457","PoqOb09Lwa8jiOs5QGzc6vIQnSe")
		const ref=btv1.newListAppTableRecordReqBuilder()
		.appToken('PoqOb09Lwa8jiOs5QGzc6vIQnSe')
		.tableId('tblTWi1nzfIMr3hu')
		.viewId('veweVpwWjk')
		.fieldNames('["任务","人员","项目","需求","资料","BUG","状态","说明","需要协助的问题"]')

		.build()
		const res=cli.base.appTableRecord.list(context.background(),ref)
		console.log(res.apiResp.string())
		console.log("res",res.data.items)
		`))
}
