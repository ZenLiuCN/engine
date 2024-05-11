package excelize

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
		import * as excel from 'github.com/xuri/excelize/v2'
		const xls=excel.newFile()
		xls.setActiveSheet(0)
		xls.setCellInt(xls.getSheetName(0),"A1",1)
		xls.saveAs("./out.xlsx")
		`))
}
