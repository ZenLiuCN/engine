package excelize

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"os"
	"testing"
)

func TestExcel(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	buf := fn.Panic1(
		vm.RunJs(
			//language=javascript
			`
import * as excel from 'go/excel'
	const xls=excel.open()
	let sheet=xls.getSheetList()[0]
	xls.setSheetName(sheet,"报表")
	sheet="报表"
	xls.setCellInt(sheet,"A1",1123)
	xls.writeBinary()
`))
	fn.Panic(os.WriteFile("out.xlsx", buf.Export().([]byte), os.ModePerm))
}
