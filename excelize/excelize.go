package excelize

import (
	"bytes"
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/xuri/excelize/v2"
	"sync"
)

var (
	//go:embed excel.d.ts
	d []byte
)

func init() {
	engine.Register(&Excel{})
}

type Excel struct {
	*engine.Engine
}

func (e Excel) Name() string {
	return "excel"
}

func (e Excel) Initialize(engine *engine.Engine) engine.Module {
	return Excel{Engine: engine}
}

func (e Excel) TypeDefine() []byte {
	return d
}

func (e Excel) Open(path string, opt *excelize.Options) *ExcelFile {
	if path == "" {
		if opt == nil {
			return &ExcelFile{excelize.NewFile(), e.Engine}
		}
		return &ExcelFile{excelize.NewFile(*opt), e.Engine}
	}
	if opt == nil {
		return &ExcelFile{
			File:   fn.Panic1(excelize.OpenFile(path)),
			Engine: e.Engine,
		}
	}
	return &ExcelFile{
		File:   fn.Panic1(excelize.OpenFile(path, *opt)),
		Engine: e.Engine,
	}
}
func (e Excel) ThemeColor(baseColor string, tint float64) string {
	return excelize.ThemeColor(baseColor, tint)
}
func (e Excel) CoordinatesToCellName(col, row int, absCol bool, absRow bool) string {
	return fn.Panic1(excelize.CoordinatesToCellName(col, row, absCol, absRow))
}
func (e Excel) CellNameToCoordinates(cell string) map[string]int {
	col, row := fn.Panic2(excelize.CellNameToCoordinates(cell))
	return map[string]int{"col": col, "row": row}
}
func (e Excel) ColumnNameToNumber(name string) int {
	return fn.Panic1(excelize.ColumnNameToNumber(name))
}
func (e Excel) ColumnNumberToName(num int) string {
	return fn.Panic1(excelize.ColumnNumberToName(num))
}
func (e Excel) JoinCellName(col string, num int) string {
	return fn.Panic1(excelize.JoinCellName(col, num))
}
func (e Excel) SplitCellName(cell string) map[string]any {
	col, row := fn.Panic2(excelize.SplitCellName(cell))
	return map[string]any{"col": col, "row": row}
}
func (e Excel) Buffer() *bytes.Buffer {
	b := buffers.Get().(*bytes.Buffer)
	b.Reset()
	return b
}

var (
	buffers = sync.Pool{New: func() any {
		buf := &bytes.Buffer{}
		return buf
	}}
)
