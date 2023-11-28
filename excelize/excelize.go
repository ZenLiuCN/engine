package excelize

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/xuri/excelize/v2"
)

var (
	//go:embed excel.d.ts
	excelizeDefine []byte
)

func init() {
	m := map[string]any{}

	m["open"] = func(path string, opt *excelize.Options) *ExcelFile {
		if path == "" {
			if opt == nil {
				return &ExcelFile{excelize.NewFile()}
			}
			return &ExcelFile{excelize.NewFile(*opt)}
		}
		if opt == nil {
			return &ExcelFile{
				File: fn.Panic1(excelize.OpenFile(path)),
			}
		}
		return &ExcelFile{
			File: fn.Panic1(excelize.OpenFile(path, *opt)),
		}
	}
	m["themeColor"] = func(baseColor string, tint float64) string {
		return excelize.ThemeColor(baseColor, tint)
	}
	m["coordinatesToCellName"] = func(col, row int, absCol bool, absRow bool) string {
		return fn.Panic1(excelize.CoordinatesToCellName(col, row, absCol, absRow))
	}
	m["cellNameToCoordinates"] = func(cell string) Coordinate {
		col, row := fn.Panic2(excelize.CellNameToCoordinates(cell))
		return Coordinate{col, row}
	}
	m["columnNameToNumber"] = func(name string) int {
		return fn.Panic1(excelize.ColumnNameToNumber(name))
	}
	m["columnNumberToName"] = func(num int) string {
		return fn.Panic1(excelize.ColumnNumberToName(num))
	}
	m["joinCellName"] = func(col string, num int) string {
		return fn.Panic1(excelize.JoinCellName(col, num))
	}
	m["splitCellName"] = func(cell string) Cell {
		col, row := fn.Panic2(excelize.SplitCellName(cell))
		return Cell{Col: col, Row: row}
	}

	engine.RegisterModule(Excel{m})
}

type Excel struct {
	m map[string]any
}

func (x Excel) Identity() string {
	return "go/excel"
}

func (x Excel) Exports() map[string]any {
	return x.m
}

func (x Excel) TypeDefine() []byte {
	return excelizeDefine
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
