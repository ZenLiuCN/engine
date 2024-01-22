package excelize

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"

	//"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"github.com/xuri/excelize/v2"
)

type ExcelFile struct {
	*excelize.File
}

func (s *ExcelFile) Close() error {
	return s.File.Close()
}
func (s *ExcelFile) UpdateLinkedValue() error {
	return s.File.UpdateLinkedValue()
}

func (s *ExcelFile) AddVBAProject(file []byte) error {
	return s.File.AddVBAProject(file)
}

func (s *ExcelFile) Save(opt *excelize.Options) error {
	if opt != nil {
		return s.File.Save(*opt)
	} else {
		return s.File.Save()
	}
}

func (s *ExcelFile) SaveAs(name string, opt *excelize.Options) error {
	if opt != nil {
		return s.File.SaveAs(name, *opt)
	} else {
		return s.File.SaveAs(name)
	}
}

func (s *ExcelFile) Write(w *engine.Buffer, opt *excelize.Options) error {
	if opt != nil {
		return s.File.Write(w.Buffer, *opt)
	} else {
		return s.File.Write(w.Buffer)
	}
}

func (s *ExcelFile) WriteBinary(opt *excelize.Options) (b []byte, err error) {
	buf := engine.GetBytesBuffer()
	defer engine.PutBytesBuffer(buf)
	buf.Reset()
	if opt != nil {
		err = s.File.Write(buf, *opt)
	} else {
		err = s.File.Write(buf)
	}
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *ExcelFile) NewStreamWriter(sheet string) (*excelize.StreamWriter, error) {
	return s.File.NewStreamWriter(sheet)
}

func (s *ExcelFile) AddShape(sheet string, shape *excelize.Shape) error {
	return s.File.AddShape(sheet, shape)
}

func (s *ExcelFile) AddSparkline(sheet string, opt *excelize.SparklineOptions) error {
	return s.File.AddSparkline(sheet, opt)
}

// region style

func (s *ExcelFile) GetStyle(idx int) (*excelize.Style, error) {
	return s.File.GetStyle(idx)
}
func (s *ExcelFile) NewConditionalStyle(style *excelize.Style) (int, error) {
	return s.File.NewConditionalStyle(style)
}
func (s *ExcelFile) GetDefaultFont() (string, error) {
	return s.File.GetDefaultFont()
}
func (s *ExcelFile) SetDefaultFont(font string) error {
	return s.File.SetDefaultFont(font)
}
func (s *ExcelFile) NewStyle(style *excelize.Style) (int, error) {
	return s.File.NewStyle(style)
}
func (s *ExcelFile) SetCellStyle(sheet, hCell, vCell string, style int) error {
	return s.File.SetCellStyle(sheet, hCell, vCell, style)
}
func (s *ExcelFile) GetCellStyle(sheet, cell string) (style int, err error) {
	return s.File.GetCellStyle(sheet, cell)
}
func (s *ExcelFile) SetConditionalFormat(sheet, rangeRef string, opts []excelize.ConditionalFormatOptions) error {
	return s.File.SetConditionalFormat(sheet, rangeRef, opts)
}
func (s *ExcelFile) UnsetConditionalFormat(sheet, rangeRef string) error {
	return s.File.UnsetConditionalFormat(sheet, rangeRef)

}
func (s *ExcelFile) GetConditionalFormats(sheet string) (map[string][]excelize.ConditionalFormatOptions, error) {
	return s.File.GetConditionalFormats(sheet)
}

//endregion
//region vml

func (s *ExcelFile) GetComments(sheet string) ([]excelize.Comment, error) {
	return s.File.GetComments(sheet)
}

func (s *ExcelFile) AddComment(sheet string, comment *excelize.Comment) error {
	return s.File.AddComment(sheet, *comment)
}

func (s *ExcelFile) DeleteComment(sheet, cell string) error {
	return s.File.DeleteComment(sheet, cell)
}

func (s *ExcelFile) AddFormControl(sheet string, ctrl *excelize.FormControl) error {
	return s.File.AddFormControl(sheet, *ctrl)
}

func (s *ExcelFile) GetFormControls(sheet string) ([]excelize.FormControl, error) {
	return s.File.GetFormControls(sheet)
}

//endregion
//region charts

func (s *ExcelFile) AddChart(sheet, cell string, chart *excelize.Chart, combo ...*excelize.Chart) error {
	return s.File.AddChart(sheet, cell, chart, combo...)
}

func (s *ExcelFile) AddChartSheet(sheet string, chart *excelize.Chart, combo ...*excelize.Chart) error {
	return s.File.AddChartSheet(sheet, chart, combo...)
}

func (s *ExcelFile) DeleteChart(sheet, cell string) error {
	return s.File.DeleteChart(sheet, cell)
}

//endregion
//region table

func (s *ExcelFile) AddTable(sheet string, table *excelize.Table) error {
	return s.File.AddTable(sheet, table)
}

func (s *ExcelFile) GetTables(sheet string) ([]excelize.Table, error) {
	return s.File.GetTables(sheet)
}

func (s *ExcelFile) DeleteTable(name string) error {
	return s.File.DeleteTable(name)
}

func (s *ExcelFile) AutoFilter(sheet, ranged string, opt ...excelize.AutoFilterOptions) error {
	return s.File.AutoFilter(sheet, ranged, opt)
}

// endregion
// region calc

func (s *ExcelFile) CalcCellValue(sheet, cell string, opt *excelize.Options) (string, error) {
	if opt != nil {
		return s.File.CalcCellValue(sheet, cell, *opt)
	} else {
		return s.File.CalcCellValue(sheet, cell)
	}
}

//endregion
//region sheet

func (s *ExcelFile) NewSheet(sheet string) (int, error) {
	return s.File.NewSheet(sheet)
}

func (s *ExcelFile) SetSheetName(sheet, target string) error {
	return s.File.SetSheetName(sheet, target)
}

func (s *ExcelFile) GetSheetIndex(sheet string) (int, error) {
	return s.File.GetSheetIndex(sheet)
}

func (s *ExcelFile) GetSheetMap() map[int]string {
	return s.File.GetSheetMap()
}

func (s *ExcelFile) GetSheetList() []string {
	return s.File.GetSheetList()
}

func (s *ExcelFile) SetSheetBackground(sheet string, picturePath string) error {
	return s.File.SetSheetBackground(sheet, picturePath)
}

func (s *ExcelFile) SetSheetBackgroundFromBytes(sheet string, ext string, picture []byte) error {
	return s.File.SetSheetBackgroundFromBytes(sheet, ext, picture)
}

func (s *ExcelFile) DeleteSheet(sheet string) error {
	return s.File.DeleteSheet(sheet)
}

func (s *ExcelFile) CopySheet(from, to int) error {
	return s.File.CopySheet(from, to)
}

func (s *ExcelFile) SetSheetVisible(sheet string, visible bool, veryHidden goja.Value) error {
	if !engine.IsNullish(veryHidden) {
		return s.File.SetSheetVisible(sheet, visible, veryHidden.ToBoolean())
	} else {
		return s.File.SetSheetVisible(sheet, visible)
	}
}

func (s *ExcelFile) GetSheetVisible(sheet string) (bool, error) {
	return s.File.GetSheetVisible(sheet)
}

func (s *ExcelFile) SearchSheet(sheet, value string, regex goja.Value) ([]string, error) {
	if !engine.IsNullish(regex) {
		return s.File.SearchSheet(sheet, value, regex.ToBoolean())
	} else {
		return s.File.SearchSheet(sheet, value)
	}
}

func (s *ExcelFile) SetHeaderFooter(sheet string, opt *excelize.HeaderFooterOptions) error {
	return s.File.SetHeaderFooter(sheet, opt)
}

func (s *ExcelFile) ProtectSheet(sheet string, opt *excelize.SheetProtectionOptions) error {
	return s.File.ProtectSheet(sheet, opt)
}

func (s *ExcelFile) UnprotectSheet(sheet string, password goja.Value) error {
	if !engine.IsNullish(password) {
		return s.File.UnprotectSheet(sheet, password.ToString().String())
	} else {
		return s.File.UnprotectSheet(sheet)
	}
}

func (s *ExcelFile) SetPageLayout(sheet string, opt *excelize.PageLayoutOptions) error {
	return s.File.SetPageLayout(sheet, opt)
}

func (s *ExcelFile) GetPageLayout(sheet string) (*excelize.PageLayoutOptions, error) {
	e, err := s.File.GetPageLayout(sheet)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (s *ExcelFile) SetDefinedName(definedName *excelize.DefinedName) error {
	return s.File.SetDefinedName(definedName)
}

func (s *ExcelFile) DeleteDefinedName(definedName *excelize.DefinedName) error {
	return s.File.DeleteDefinedName(definedName)
}

func (s *ExcelFile) GetDefinedName() []excelize.DefinedName {
	return s.File.GetDefinedName()
}

func (s *ExcelFile) GroupSheets(sheets []string) error {
	return s.File.GroupSheets(sheets)
}

func (s *ExcelFile) UngroupSheets() error {
	return s.File.UngroupSheets()
}

func (s *ExcelFile) InsertPageBreak(sheet, cell string) error {
	return s.File.InsertPageBreak(sheet, cell)
}

func (s *ExcelFile) RemovePageBreak(sheet, cell string) error {
	return s.File.RemovePageBreak(sheet, cell)
}

func (s *ExcelFile) SetSheetDimension(sheet, rangeRef string) error {
	return s.File.SetSheetDimension(sheet, rangeRef)
}

func (s *ExcelFile) GetSheetDimension(sheet string) (string, error) {
	return s.File.GetSheetDimension(sheet)
}

func (s *ExcelFile) SetPageMargins(sheet string, opt *excelize.PageLayoutMarginsOptions) error {
	return s.File.SetPageMargins(sheet, opt)
}

func (s *ExcelFile) GetPageMargins(sheet string) (*excelize.PageLayoutMarginsOptions, error) {
	e, err := s.File.GetPageMargins(sheet)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (s *ExcelFile) SetSheetProps(sheet string, prop *excelize.SheetPropsOptions) error {
	return s.File.SetSheetProps(sheet, prop)
}

func (s *ExcelFile) GetSheetProps(sheet string) (*excelize.SheetPropsOptions, error) {
	e, err := s.File.GetSheetProps(sheet)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (s *ExcelFile) SetSheetView(sheet string, viewIndex int, prop *excelize.ViewOptions) error {
	return s.File.SetSheetView(sheet, viewIndex, prop)
}

func (s *ExcelFile) GetSheetView(sheet string, viewIndex int) (*excelize.ViewOptions, error) {
	e, err := s.File.GetSheetView(sheet, viewIndex)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
func (s *ExcelFile) SetPanes(sheet string, panes *excelize.Panes) error {
	return s.File.SetPanes(sheet, panes)
}
func (s *ExcelFile) GetPanes(sheet string) (*excelize.Panes, error) {
	e, err := s.File.GetPanes(sheet)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (s *ExcelFile) GetSheetName(index int) string {
	return s.File.GetSheetName(index)
}

func (s *ExcelFile) SetActiveSheet(index int) {
	s.File.SetActiveSheet(index)
}

func (s *ExcelFile) GetActiveSheetIndex() int {
	return s.File.GetActiveSheetIndex()
}

//endregion
//region rows

func (s *ExcelFile) GetRows(sheet string, opt *excelize.Options) ([][]string, error) {
	if opt != nil {
		return s.File.GetRows(sheet, *opt)
	} else {
		return s.File.GetRows(sheet)
	}
}

func (s *ExcelFile) Rows(sheet string) (*excelize.Rows, error) {
	return s.File.Rows(sheet)
}

func (s *ExcelFile) SetRowHeight(sheet string, row int, height float64) error {
	return s.File.SetRowHeight(sheet, row, height)
}

func (s *ExcelFile) GetRowHeight(sheet string, row int) (float64, error) {
	return s.File.GetRowHeight(sheet, row)
}

func (s *ExcelFile) SetRowVisible(sheet string, row int, visible bool) error {
	return s.File.SetRowVisible(sheet, row, visible)
}

func (s *ExcelFile) SetRowOutlineLevel(sheet string, row int, level int) error {
	return s.File.SetRowOutlineLevel(sheet, row, uint8(level))
}

func (s *ExcelFile) GetRowOutlineLevel(sheet string, row int) (int, error) {
	x, err := s.File.GetRowOutlineLevel(sheet, row)
	if err != nil {
		return 0, err
	}
	return int(x), err
}

func (s *ExcelFile) RemoveRow(sheet string, row int) error {
	return s.File.RemoveRow(sheet, row)
}

func (s *ExcelFile) InsertRows(sheet string, row, n int) error {
	return s.File.InsertRows(sheet, row, n)
}

func (s *ExcelFile) DuplicateRow(sheet string, row int) error {
	return s.File.DuplicateRow(sheet, row)
}

func (s *ExcelFile) DuplicateRowTo(sheet string, rowFrom, rowTo int) error {
	return s.File.DuplicateRowTo(sheet, rowFrom, rowTo)
}

func (s *ExcelFile) SetRowStyle(sheet string, start, end, style int) error {
	return s.File.SetRowStyle(sheet, start, end, style)
}

func (s *ExcelFile) GetRowVisible(sheet string, row int) (bool, error) {
	return s.File.GetRowVisible(sheet, row)
}

//endregion
//region cols

func (s *ExcelFile) GetCols(sheet string, opt *excelize.Options) ([][]string, error) {
	if opt == nil {
		return s.File.GetCols(sheet)
	} else {
		return s.File.GetCols(sheet, *opt)
	}
}

func (s *ExcelFile) Cols(sheet string) (*excelize.Cols, error) {
	return s.File.Cols(sheet)
}

func (s *ExcelFile) SetColVisible(sheet, col string, visible bool) error {
	return s.File.SetColVisible(sheet, col, visible)
}

func (s *ExcelFile) GetColVisible(sheet, col string) (bool, error) {
	return s.File.GetColVisible(sheet, col)
}

func (s *ExcelFile) GetColOutlineLevel(sheet, col string) (int, error) {
	e, er := s.File.GetColOutlineLevel(sheet, col)
	if er != nil {
		return 0, er
	}
	return int(e), nil
}
func (s *ExcelFile) SetColOutlineLevel(sheet, col string, level int) error {
	return s.File.SetColOutlineLevel(sheet, col, uint8(level))
}
func (s *ExcelFile) SetColStyle(sheet, col string, styleId int) error {
	return s.File.SetColStyle(sheet, col, styleId)
}
func (s *ExcelFile) GetColStyle(sheet, col string) int {
	return fn.Panic1(s.File.GetColStyle(sheet, col))
}
func (s *ExcelFile) SetColWidth(sheet, colStart, colEnd string, width float64) error {
	return s.File.SetColWidth(sheet, colStart, colEnd, width)
}

func (s *ExcelFile) GetColWidth(sheet, col string) float64 {
	return fn.Panic1(s.File.GetColWidth(sheet, col))
}

func (s *ExcelFile) InsertCols(sheet, col string, n int) error {
	return s.File.InsertCols(sheet, col, n)
}
func (s *ExcelFile) RemoveCol(sheet, col string) error {
	return s.File.RemoveCol(sheet, col)
}

//endregion
//region data validation

func (s *ExcelFile) AddDataValidation(sheet string, dv *excelize.DataValidation) error {
	return s.File.AddDataValidation(sheet, dv)
}

func (s *ExcelFile) GetDataValidations(sheet string) []*excelize.DataValidation {
	return fn.Panic1(s.File.GetDataValidations(sheet))
}

func (s *ExcelFile) DeleteDataValidation(sheet string, sqref ...string) error {
	return s.File.DeleteDataValidation(sheet, sqref...)
}

// endregion
// region properties

func (s *ExcelFile) SetAppProps(prop *excelize.AppProperties) error {
	return s.File.SetAppProps(prop)
}

func (s *ExcelFile) GetAppProps() *excelize.AppProperties {
	return fn.Panic1(s.File.GetAppProps())
}

func (s *ExcelFile) SetDocProps(prop *excelize.DocProperties) error {
	return s.File.SetDocProps(prop)
}

func (s *ExcelFile) GetDocProps() (*excelize.DocProperties, error) {
	return s.File.GetDocProps()
}

// endregion
// region cells

func (s *ExcelFile) GetCellValue(sheet, cell string, opt *excelize.Options) (string, error) {
	if opt != nil {
		return s.File.GetCellValue(sheet, cell, *opt)
	} else {
		return s.File.GetCellValue(sheet, cell)
	}
}

func (s *ExcelFile) SetCellValue(sheet, cell string, value goja.Value) error {
	if engine.IsNullish(value) {
		return s.File.SetCellValue(sheet, cell, nil)
	} else {
		return s.File.SetCellValue(sheet, cell, value.Export())
	}
}

func (s *ExcelFile) SetCellInt(sheet, cell string, value int) error {
	return s.File.SetCellInt(sheet, cell, value)
}

func (s *ExcelFile) SetCellBool(sheet, cell string, value bool) error {
	return s.File.SetCellBool(sheet, cell, value)
}

func (s *ExcelFile) SetCellFloat(sheet, cell string, value float64, precision, bitSize int) error {
	return s.File.SetCellFloat(sheet, cell, value, precision, bitSize)
}

func (s *ExcelFile) SetCellStr(sheet, cell string, value string) error {
	return s.File.SetCellStr(sheet, cell, value)
}

func (s *ExcelFile) SetCellDefault(sheet, cell string, value string) error {
	return s.File.SetCellDefault(sheet, cell, value)
}

func (s *ExcelFile) GetCellHyperLink(sheet, cell string) struct {
	link   bool
	target string
} {
	l, t := fn.Panic2(s.File.GetCellHyperLink(sheet, cell))
	return struct {
		link   bool
		target string
	}{link: l, target: t}
}

func (s *ExcelFile) SetCellHyperLink(sheet, cell, link, linkType string, opt *excelize.HyperlinkOpts) error {
	if opt != nil {
		return s.File.SetCellHyperLink(sheet, cell, link, linkType, *opt)
	} else {
		return s.File.SetCellHyperLink(sheet, cell, link, linkType)
	}

}

func (s *ExcelFile) GetCellRichText(sheet, cell string) ([]excelize.RichTextRun, error) {
	return s.File.GetCellRichText(sheet, cell)
}

func (s *ExcelFile) SetCellRichText(sheet, cell string, runs []excelize.RichTextRun) error {
	return s.File.SetCellRichText(sheet, cell, runs)
}

// func (s* ExcelFile)setSheetRow(sheet, cell string, values Array<number | bool | string | null>)

// func (s* ExcelFile)setSheetCol(sheet, cell string, values Array<number | bool | string | null>)

func (s *ExcelFile) GetCellFormula(sheet, cell string) (string, error) {
	return s.File.GetCellFormula(sheet, cell)
}

func (s *ExcelFile) SetCellFormula(sheet, cell, formula string, opt *excelize.FormulaOpts) error {
	if opt != nil {
		return s.File.SetCellFormula(sheet, cell, formula, *opt)
	} else {
		return s.File.SetCellFormula(sheet, cell, formula)
	}
}

func (s *ExcelFile) GetCellType(sheet, cell string) (int, error) {
	e, er := s.File.GetCellType(sheet, cell)
	if er != nil {
		return 0, er
	}
	return int(e), nil
}

func (s *ExcelFile) MergeCell(sheet, hCell, vCell string) error {
	return s.File.MergeCell(sheet, hCell, vCell)
}

func (s *ExcelFile) UnmergeCell(sheet, hCell, vCell string) error {
	return s.File.UnmergeCell(sheet, hCell, vCell)
}

func (s *ExcelFile) GetMergeCells(sheet string) ([]excelize.MergeCell, error) {
	return s.File.GetMergeCells(sheet)
}

//endregion
//region picture

func (s *ExcelFile) AddPicture(sheet, cell, name string, opt *excelize.GraphicOptions) error {
	return s.File.AddPicture(sheet, cell, name, opt)
}

func (s *ExcelFile) AddPictureFromBytes(sheet, cell string, pic *excelize.Picture) error {
	return s.File.AddPictureFromBytes(sheet, cell, pic)
}

func (s *ExcelFile) DeletePicture(sheet, cell string) error {
	return s.File.DeletePicture(sheet, cell)
}

func (s *ExcelFile) GetPictures(sheet, cell string) ([]excelize.Picture, error) {
	return s.File.GetPictures(sheet, cell)
}

// endregion
// region pivot table

func (s *ExcelFile) AddPivotTable(opt *excelize.PivotTableOptions) error {
	return s.File.AddPivotTable(opt)
}

// endregion
// region work book

func (s *ExcelFile) SetWorkbookProps(opt *excelize.WorkbookPropsOptions) error {
	return s.File.SetWorkbookProps(opt)
}

func (s *ExcelFile) GetWorkbookProps() (*excelize.WorkbookPropsOptions, error) {
	e, err := s.File.GetWorkbookProps()
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (s *ExcelFile) ProtectWorkbook(opt *excelize.WorkbookProtectionOptions) error {
	return s.File.ProtectWorkbook(opt)
}

func (s *ExcelFile) UnprotectWorkbook(password goja.Value) error {
	if engine.IsNullish(password) {
		return s.File.UnprotectWorkbook()
	} else {
		return s.File.UnprotectWorkbook(password.ToString().String())
	}
}

//endregion
