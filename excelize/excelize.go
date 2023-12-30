package excelize

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/xuri/excelize/v2"
)

var (
	//go:embed excel.d.ts
	excelDefine []byte
)

func init() {
	engine.RegisterModule(Excel{})
}

type Excel struct {
}

func (x Excel) ExportsWithEngine(eng *engine.Engine) map[string]any {
	return map[string]any{

		"open": func(path string, opt *excelize.Options) *ExcelFile {
			if path == "" {
				if opt == nil {
					return engine.RegisterResource(eng, &ExcelFile{excelize.NewFile()})
				}
				return engine.RegisterResource(eng, &ExcelFile{excelize.NewFile(*opt)})
			}
			if opt == nil {
				return engine.RegisterResource(eng, &ExcelFile{
					File: fn.Panic1(excelize.OpenFile(path)),
				})
			}
			return engine.RegisterResource(eng, &ExcelFile{
				File: fn.Panic1(excelize.OpenFile(path, *opt)),
			})
		},
		"themeColor": func(baseColor string, tint float64) string {
			return excelize.ThemeColor(baseColor, tint)
		},
		"coordinatesToCellName": func(col, row int, absCol bool, absRow bool) string {
			return fn.Panic1(excelize.CoordinatesToCellName(col, row, absCol, absRow))
		},
		"cellNameToCoordinates": func(cell string) Coordinate {
			col, row := fn.Panic2(excelize.CellNameToCoordinates(cell))
			return Coordinate{col, row}
		},
		"columnNameToNumber": func(name string) int {
			return fn.Panic1(excelize.ColumnNameToNumber(name))
		},
		"columnNumberToName": func(num int) string {
			return fn.Panic1(excelize.ColumnNumberToName(num))
		},
		"joinCellName": func(col string, num int) string {
			return fn.Panic1(excelize.JoinCellName(col, num))
		},
		"splitCellName": func(cell string) Cell {
			col, row := fn.Panic2(excelize.SplitCellName(cell))
			return Cell{Col: col, Row: row}
		},
	}
}

func (x Excel) Identity() string {
	return "go/excel"
}

func (x Excel) Exports() map[string]any {
	return nil
}

func (x Excel) TypeDefine() []byte {
	return excelDefine
}

type (
	Coordinate struct {
		Col, Row int
	}
	Cell struct {
		Col string
		Row int
	}
)
