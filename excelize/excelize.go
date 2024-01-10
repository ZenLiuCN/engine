package excelize

import (
	_ "embed"
	"fmt"
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
		"open": func(path string, opt *excelize.Options) (e *ExcelFile, err error) {
			defer func() {
				if r := recover(); r != nil {
					switch ex := r.(type) {
					case error:
						err = ex
						e = nil
					default:
						err = fmt.Errorf("%v", ex)
						e = nil
					}
				}
			}()
			if path == "" {
				if opt == nil {
					return engine.RegisterResource(eng, &ExcelFile{excelize.NewFile()}), nil
				}
				return engine.RegisterResource(eng, &ExcelFile{excelize.NewFile(*opt)}), nil
			}
			if opt == nil {
				return engine.RegisterResource(eng, &ExcelFile{
					File: fn.Panic1(excelize.OpenFile(path)),
				}), nil
			}
			return engine.RegisterResource(eng, &ExcelFile{
				File: fn.Panic1(excelize.OpenFile(path, *opt)),
			}), nil
		},
		"themeColor": func(baseColor string, tint float64) string {
			return excelize.ThemeColor(baseColor, tint)
		},
		"coordinatesToCellName": func(col, row int, absCol bool, absRow bool) (string, error) {
			return excelize.CoordinatesToCellName(col, row, absCol, absRow)
		},
		"cellNameToCoordinates": func(cell string) (c *Coordinate, err error) {
			c = new(Coordinate)
			c.Col, c.Row, err = excelize.CellNameToCoordinates(cell)
			return
		},
		"columnNameToNumber": func(name string) (int, error) {
			return excelize.ColumnNameToNumber(name)
		},
		"columnNumberToName": func(num int) (string, error) {
			return excelize.ColumnNumberToName(num)
		},
		"joinCellName": func(col string, num int) (string, error) {
			return excelize.JoinCellName(col, num)
		},
		"splitCellName": func(cell string) (c *Cell, err error) {
			c = new(Cell)
			c.Col, c.Row, err = excelize.SplitCellName(cell)
			return
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
