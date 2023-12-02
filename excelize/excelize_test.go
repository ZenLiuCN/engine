package excelize

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
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
	const xls=excel.open()
	let sheet=xls.getSheetList()[0]
	xls.setSheetName(sheet,"报表")
	sheet="报表"
	xls.setCellInt(sheet,"A1",1123)
	xls.writeToBuffer()
`))
	fn.Panic(os.WriteFile("out.xlsx", buf.Export().(goja.ArrayBuffer).Bytes(), os.ModePerm))
}
