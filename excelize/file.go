package excelize

import (
	"bytes"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"github.com/xuri/excelize/v2"
)

type ExcelFile struct {
	*excelize.File
	*engine.Engine
}

func (s *ExcelFile) Close() {
	fn.Panic(s.File.Close())
}
func (s *ExcelFile) UpdateLinkedValue() {
	fn.Panic(s.File.UpdateLinkedValue())
}

func (s *ExcelFile) AddVBAProject(file []byte) {
	fn.Panic(s.File.AddVBAProject(file))
}

func (s *ExcelFile) Save(opt *excelize.Options) {
	if opt != nil {
		fn.Panic(s.File.Save(*opt))
	} else {
		fn.Panic(s.File.Save())
	}
}

func (s *ExcelFile) SaveAs(name string, opt *excelize.Options) {
	if opt != nil {
		fn.Panic(s.File.SaveAs(name, *opt))
	} else {
		fn.Panic(s.File.SaveAs(name))
	}
}

func (s *ExcelFile) Write(w *bytes.Buffer, opt *excelize.Options) {
	if opt != nil {
		fn.Panic(s.File.Write(w, *opt))
	} else {
		fn.Panic(s.File.Write(w))
	}
}

func (s *ExcelFile) WriteToBuffer(opt *excelize.Options) goja.Value {
	buf := engine.GetBytesBuffer()
	buf.Reset()
	if opt != nil {
		fn.Panic(s.File.Write(buf, *opt))
	} else {
		fn.Panic(s.File.Write(buf))
	}
	return s.Engine.ToValue(s.Engine.NewArrayBuffer(buf.Bytes()))
}

func (s *ExcelFile) NewStreamWriter(sheet string) *excelize.StreamWriter {
	return fn.Panic1(s.File.NewStreamWriter(sheet))
}

func (s *ExcelFile) AddShape(sheet string, shape *excelize.Shape) {
	fn.Panic(s.File.AddShape(sheet, shape))
}

func (s *ExcelFile) AddSparkline(sheet string, opt *excelize.SparklineOptions) {
	fn.Panic(s.File.AddSparkline(sheet, opt))
}

// region style

func (s *ExcelFile) GetStyle(idx int) *excelize.Style {
	return fn.Panic1(s.File.GetStyle(idx))
}
func (s *ExcelFile) NewConditionalStyle(style *excelize.Style) int {
	return fn.Panic1(s.File.NewConditionalStyle(style))
}
func (s *ExcelFile) GetDefaultFont() string {
	return fn.Panic1(s.File.GetDefaultFont())
}
func (s *ExcelFile) SetDefaultFont(font string) {
	fn.Panic(s.File.SetDefaultFont(font))
}
func (s *ExcelFile) NewStyle(style *excelize.Style) int {
	return fn.Panic1(s.File.NewStyle(style))
}
func (s *ExcelFile) SetCellStyle(sheet, hCell, vCell string, style int) {
	fn.Panic(s.File.SetCellStyle(sheet, hCell, vCell, style))
}
func (s *ExcelFile) GetCellStyle(sheet, cell string) (style int) {
	return fn.Panic1(s.File.GetCellStyle(sheet, cell))
}
func (s *ExcelFile) SetConditionalFormat(sheet, rangeRef string, opts []excelize.ConditionalFormatOptions) {
	fn.Panic(s.File.SetConditionalFormat(sheet, rangeRef, opts))
}
func (s *ExcelFile) UnsetConditionalFormat(sheet, rangeRef string) {
	fn.Panic(s.File.UnsetConditionalFormat(sheet, rangeRef))
}
func (s *ExcelFile) GetConditionalFormats(sheet string) map[string][]excelize.ConditionalFormatOptions {
	return fn.Panic1(s.File.GetConditionalFormats(sheet))
}

//endregion
//region vml

func (s *ExcelFile) GetComments(sheet string) []excelize.Comment {
	return fn.Panic1(s.File.GetComments(sheet))
}

func (s *ExcelFile) AddComment(sheet string, comment *excelize.Comment) {
	fn.Panic(s.File.AddComment(sheet, *comment))
}

func (s *ExcelFile) DeleteComment(sheet, cell string) {
	fn.Panic(s.File.DeleteComment(sheet, cell))
}

func (s *ExcelFile) AddFormControl(sheet string, ctrl *excelize.FormControl) {
	fn.Panic(s.File.AddFormControl(sheet, *ctrl))
}

func (s *ExcelFile) GetFormControls(sheet string) []excelize.FormControl {
	return fn.Panic1(s.File.GetFormControls(sheet))
}

//endregion
//region charts

func (s *ExcelFile) AddChart(sheet, cell string, chart *excelize.Chart, combo ...*excelize.Chart) {
	fn.Panic(s.File.AddChart(sheet, cell, chart, combo...))
}

func (s *ExcelFile) AddChartSheet(sheet string, chart *excelize.Chart, combo ...*excelize.Chart) {
	fn.Panic(s.File.AddChartSheet(sheet, chart, combo...))
}

func (s *ExcelFile) DeleteChart(sheet, cell string) {
	fn.Panic(s.File.DeleteChart(sheet, cell))
}

//endregion
//region table

func (s *ExcelFile) AddTable(sheet string, table *excelize.Table) {
	fn.Panic(s.File.AddTable(sheet, table))
}

func (s *ExcelFile) GetTables(sheet string) []excelize.Table {
	return fn.Panic1(s.File.GetTables(sheet))
}

func (s *ExcelFile) DeleteTable(name string) {
	fn.Panic(s.File.DeleteTable(name))
}

func (s *ExcelFile) AutoFilter(sheet, ranged string, opt ...excelize.AutoFilterOptions) {
	fn.Panic(s.File.AutoFilter(sheet, ranged, opt))
}

// endregion
// region calc
func (s *ExcelFile) CalcCellValue(sheet, cell string, opt *excelize.Options) string {
	if opt != nil {
		return fn.Panic1(s.File.CalcCellValue(sheet, cell, *opt))
	} else {
		return fn.Panic1(s.File.CalcCellValue(sheet, cell))
	}
}

//endregion
//region sheet

func (s *ExcelFile) NewSheet(sheet string) int {
	return fn.Panic1(s.File.NewSheet(sheet))
}

func (s *ExcelFile) SetSheetName(sheet, target string) {
	fn.Panic(s.File.SetSheetName(sheet, target))
}

func (s *ExcelFile) GetSheetIndex(sheet string) int {
	return fn.Panic1(s.File.GetSheetIndex(sheet))
}

func (s *ExcelFile) GetSheetMap() map[int]string {
	return s.File.GetSheetMap()
}

func (s *ExcelFile) GetSheetList() []string {
	return s.File.GetSheetList()
}

func (s *ExcelFile) SetSheetBackground(sheet string, picturePath string) {
	fn.Panic(s.File.SetSheetBackground(sheet, picturePath))
}

func (s *ExcelFile) SetSheetBackgroundFromBytes(sheet string, ext string, picture []byte) {
	fn.Panic(s.File.SetSheetBackgroundFromBytes(sheet, ext, picture))
}

func (s *ExcelFile) DeleteSheet(sheet string) {
	fn.Panic(s.File.DeleteSheet(sheet))
}

func (s *ExcelFile) CopySheet(from, to int) {
	fn.Panic(s.File.CopySheet(from, to))
}

func (s *ExcelFile) SetSheetVisible(sheet string, visible bool, veryHidden goja.Value) {
	if !engine.IsNullish(veryHidden) {
		fn.Panic(s.File.SetSheetVisible(sheet, visible, veryHidden.ToBoolean()))
	} else {
		fn.Panic(s.File.SetSheetVisible(sheet, visible))
	}

}

func (s *ExcelFile) GetSheetVisible(sheet string) bool {
	return fn.Panic1(s.File.GetSheetVisible(sheet))
}

func (s *ExcelFile) SearchSheet(sheet, value string, regex goja.Value) []string {
	if !engine.IsNullish(regex) {
		return fn.Panic1(s.File.SearchSheet(sheet, value, regex.ToBoolean()))
	} else {
		return fn.Panic1(s.File.SearchSheet(sheet, value))
	}
}

func (s *ExcelFile) SetHeaderFooter(sheet string, opt *excelize.HeaderFooterOptions) {
	fn.Panic(s.File.SetHeaderFooter(sheet, opt))
}

func (s *ExcelFile) ProtectSheet(sheet string, opt *excelize.SheetProtectionOptions) {
	fn.Panic(s.File.ProtectSheet(sheet, opt))
}

func (s *ExcelFile) UnprotectSheet(sheet string, password goja.Value) {
	if !engine.IsNullish(password) {
		fn.Panic(s.File.UnprotectSheet(sheet, password.ToString().String()))
	} else {
		fn.Panic(s.File.UnprotectSheet(sheet))
	}
}

func (s *ExcelFile) SetPageLayout(sheet string, opt *excelize.PageLayoutOptions) {
	fn.Panic(s.File.SetPageLayout(sheet, opt))
}

func (s *ExcelFile) GetPageLayout(sheet string) *excelize.PageLayoutOptions {
	e := fn.Panic1(s.File.GetPageLayout(sheet))
	return &e
}

func (s *ExcelFile) SetDefinedName(definedName *excelize.DefinedName) {
	fn.Panic(s.File.SetDefinedName(definedName))
}

func (s *ExcelFile) DeleteDefinedName(definedName *excelize.DefinedName) {
	fn.Panic(s.File.DeleteDefinedName(definedName))
}

func (s *ExcelFile) GetDefinedName() []excelize.DefinedName {
	return s.File.GetDefinedName()
}

func (s *ExcelFile) GroupSheets(sheets []string) {
	fn.Panic(s.File.GroupSheets(sheets))
}

func (s *ExcelFile) UngroupSheets() {
	fn.Panic(s.File.UngroupSheets())
}

func (s *ExcelFile) InsertPageBreak(sheet, cell string) {
	fn.Panic(s.File.InsertPageBreak(sheet, cell))
}

func (s *ExcelFile) RemovePageBreak(sheet, cell string) {
	fn.Panic(s.File.RemovePageBreak(sheet, cell))
}

func (s *ExcelFile) SetSheetDimension(sheet, rangeRef string) {
	fn.Panic(s.File.SetSheetDimension(sheet, rangeRef))
}

func (s *ExcelFile) GetSheetDimension(sheet string) string {
	return fn.Panic1(s.File.GetSheetDimension(sheet))
}

func (s *ExcelFile) SetPageMargins(sheet string, opt *excelize.PageLayoutMarginsOptions) {
	fn.Panic(s.File.SetPageMargins(sheet, opt))
}

func (s *ExcelFile) GetPageMargins(sheet string) *excelize.PageLayoutMarginsOptions {
	e := fn.Panic1(s.File.GetPageMargins(sheet))
	return &e
}

func (s *ExcelFile) SetSheetProps(sheet string, prop *excelize.SheetPropsOptions) {
	s.SetSheetProps(sheet, prop)
}

func (s *ExcelFile) GetSheetProps(sheet string) *excelize.SheetPropsOptions {
	e := fn.Panic1(s.File.GetSheetProps(sheet))
	return &e
}

func (s *ExcelFile) SetSheetView(sheet string, viewIndex int, prop *excelize.ViewOptions) {
	fn.Panic(s.File.SetSheetView(sheet, viewIndex, prop))
}

func (s *ExcelFile) GetSheetView(sheet string, viewIndex int) *excelize.ViewOptions {
	e := fn.Panic1(s.File.GetSheetView(sheet, viewIndex))
	return &e
}
func (s *ExcelFile) SetPanes(sheet string, panes *excelize.Panes) {
	fn.Panic(s.File.SetPanes(sheet, panes))
}
func (s *ExcelFile) GetPanes(sheet string) *excelize.Panes {
	e := fn.Panic1(s.File.GetPanes(sheet))
	return &e
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

func (s *ExcelFile) GetRows(sheet string, opt *excelize.Options) [][]string {
	if opt == nil {
		return fn.Panic1(s.File.GetRows(sheet, *opt))
	} else {
		return fn.Panic1(s.File.GetRows(sheet))
	}
}

func (s *ExcelFile) Rows(sheet string) *excelize.Rows {
	return fn.Panic1(s.File.Rows(sheet))
}

func (s *ExcelFile) SetRowHeight(sheet string, row int, height float64) {
	fn.Panic(s.File.SetRowHeight(sheet, row, height))
}

func (s *ExcelFile) GetRowHeight(sheet string, row int) float64 {
	return fn.Panic1(s.File.GetRowHeight(sheet, row))
}

func (s *ExcelFile) SetRowVisible(sheet string, row int, visible bool) {
	fn.Panic(s.File.SetRowVisible(sheet, row, visible))
}

func (s *ExcelFile) SetRowOutlineLevel(sheet string, row int, level int) {
	fn.Panic(s.File.SetRowOutlineLevel(sheet, row, uint8(level)))
}

func (s *ExcelFile) GetRowOutlineLevel(sheet string, row int) int {
	return int(fn.Panic1(s.File.GetRowOutlineLevel(sheet, row)))
}

func (s *ExcelFile) RemoveRow(sheet string, row int) {
	fn.Panic(s.File.RemoveRow(sheet, row))
}

func (s *ExcelFile) InsertRows(sheet string, row, n int) {
	fn.Panic(s.File.InsertRows(sheet, row, n))
}

func (s *ExcelFile) DuplicateRow(sheet string, row int) {
	fn.Panic(s.File.DuplicateRow(sheet, row))
}

func (s *ExcelFile) DuplicateRowTo(sheet string, rowFrom, rowTo int) {
	fn.Panic(s.File.DuplicateRowTo(sheet, rowFrom, rowTo))
}

func (s *ExcelFile) SetRowStyle(sheet string, start, end, style int) {
	fn.Panic(s.File.SetRowStyle(sheet, start, end, style))
}

func (s *ExcelFile) GetRowVisible(sheet string, row int) bool {
	return fn.Panic1(s.File.GetRowVisible(sheet, row))
}

//endregion
//region cols

func (s *ExcelFile) GetCols(sheet string, opt *excelize.Options) [][]string {
	if opt == nil {
		return fn.Panic1(s.File.GetCols(sheet))
	} else {
		return fn.Panic1(s.File.GetCols(sheet, *opt))
	}
}

func (s *ExcelFile) Cols(sheet string) *excelize.Cols {
	return fn.Panic1(s.File.Cols(sheet))
}

func (s *ExcelFile) SetColVisible(sheet, col string, visible bool) {
	fn.Panic(s.File.SetColVisible(sheet, col, visible))
}

func (s *ExcelFile) GetColVisible(sheet, col string) bool {
	return fn.Panic1(s.File.GetColVisible(sheet, col))
}

func (s *ExcelFile) GetColOutlineLevel(sheet, col string) int {
	return int(fn.Panic1(s.File.GetColOutlineLevel(sheet, col)))
}
func (s *ExcelFile) SetColOutlineLevel(sheet, col string, level int) {
	fn.Panic(s.File.SetColOutlineLevel(sheet, col, uint8(level)))
}
func (s *ExcelFile) SetColStyle(sheet, col string, styleId int) {
	fn.Panic(s.File.SetColStyle(sheet, col, styleId))
}
func (s *ExcelFile) GetColStyle(sheet, col string) int {
	return fn.Panic1(s.File.GetColStyle(sheet, col))
}
func (s *ExcelFile) SetColWidth(sheet, colStart, colEnd string, width float64) {
	fn.Panic(s.File.SetColWidth(sheet, colStart, colEnd, width))
}

func (s *ExcelFile) GetColWidth(sheet, col string) float64 {
	return fn.Panic1(s.File.GetColWidth(sheet, col))
}

func (s *ExcelFile) InsertCols(sheet, col string, n int) {
	fn.Panic(s.File.InsertCols(sheet, col, n))
}
func (s *ExcelFile) RemoveCol(sheet, col string) {
	fn.Panic(s.File.RemoveCol(sheet, col))
}

//endregion
//region data validation

func (s *ExcelFile) AddDataValidation(sheet string, dv *excelize.DataValidation) {
	fn.Panic(s.File.AddDataValidation(sheet, dv))
}

func (s *ExcelFile) GetDataValidations(sheet string) []*excelize.DataValidation {
	return fn.Panic1(s.File.GetDataValidations(sheet))
}

func (s *ExcelFile) DeleteDataValidation(sheet string, sqref ...string) {
	fn.Panic(s.File.DeleteDataValidation(sheet, sqref...))
}

// endregion
// region properties
func (s *ExcelFile) SetAppProps(prop *excelize.AppProperties) {
	fn.Panic(s.File.SetAppProps(prop))
}

func (s *ExcelFile) GetAppProps() *excelize.AppProperties {
	return fn.Panic1(s.File.GetAppProps())
}

func (s *ExcelFile) SetDocProps(prop *excelize.DocProperties) {
	fn.Panic(s.File.SetDocProps(prop))
}

func (s *ExcelFile) GetDocProps() *excelize.DocProperties {
	return fn.Panic1(s.File.GetDocProps())
}

// endregion
// region cells
func (s *ExcelFile) GetCellValue(sheet, cell string, opt *excelize.Options) string {
	if opt != nil {
		return fn.Panic1(s.File.GetCellValue(sheet, cell, *opt))
	} else {
		return fn.Panic1(s.File.GetCellValue(sheet, cell))
	}
}

func (s *ExcelFile) SetCellValue(sheet, cell string, value goja.Value) {
	if engine.IsNullish(value) {
		fn.Panic(s.File.SetCellValue(sheet, cell, nil))
	} else {
		fn.Panic(s.File.SetCellValue(sheet, cell, value.Export()))
	}
}

func (s *ExcelFile) SetCellInt(sheet, cell string, value int) {
	fn.Panic(s.File.SetCellInt(sheet, cell, value))
}

func (s *ExcelFile) SetCellBool(sheet, cell string, value bool) {
	fn.Panic(s.File.SetCellBool(sheet, cell, value))
}

func (s *ExcelFile) SetCellFloat(sheet, cell string, value float64, precision, bitSize int) {
	fn.Panic(s.File.SetCellFloat(sheet, cell, value, precision, bitSize))
}

func (s *ExcelFile) SetCellStr(sheet, cell string, value string) {
	fn.Panic(s.File.SetCellStr(sheet, cell, value))
}

func (s *ExcelFile) SetCellDefault(sheet, cell string, value string) {
	fn.Panic(s.File.SetCellDefault(sheet, cell, value))
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

func (s *ExcelFile) SetCellHyperLink(sheet, cell, link, linkType string, opt *excelize.HyperlinkOpts) {
	if opt != nil {
		fn.Panic(s.File.SetCellHyperLink(sheet, cell, link, linkType, *opt))
	} else {
		fn.Panic(s.File.SetCellHyperLink(sheet, cell, link, linkType))
	}
}

func (s *ExcelFile) GetCellRichText(sheet, cell string) []excelize.RichTextRun {
	return fn.Panic1(s.File.GetCellRichText(sheet, cell))
}

func (s *ExcelFile) SetCellRichText(sheet, cell string, runs []excelize.RichTextRun) {
	fn.Panic(s.File.SetCellRichText(sheet, cell, runs))
}

// func (s* ExcelFile)setSheetRow(sheet, cell string, values Array<number | bool | string | null>)

// func (s* ExcelFile)setSheetCol(sheet, cell string, values Array<number | bool | string | null>)

func (s *ExcelFile) GetCellFormula(sheet, cell string) string {
	return fn.Panic1(s.File.GetCellFormula(sheet, cell))
}

func (s *ExcelFile) SetCellFormula(sheet, cell string, formula string, opt *excelize.FormulaOpts) {
	if opt != nil {
		fn.Panic(s.File.SetCellFormula(sheet, cell, formula, *opt))
	} else {
		fn.Panic(s.File.SetCellFormula(sheet, cell, formula))
	}
}

func (s *ExcelFile) GetCellType(sheet, cell string) int {
	return int(fn.Panic1(s.File.GetCellType(sheet, cell)))
}

func (s *ExcelFile) MergeCell(sheet, hCell, vCell string) {
	fn.Panic(s.File.MergeCell(sheet, hCell, vCell))
}

func (s *ExcelFile) UnmergeCell(sheet, hCell, vCell string) {
	fn.Panic(s.File.UnmergeCell(sheet, hCell, vCell))
}

func (s *ExcelFile) GetMergeCells(sheet string) []excelize.MergeCell {
	return fn.Panic1(s.File.GetMergeCells(sheet))
}

//endregion
//region picture

func (s *ExcelFile) AddPicture(sheet, cell, name string, opt *excelize.GraphicOptions) {
	fn.Panic(s.File.AddPicture(sheet, cell, name, opt))
}

func (s *ExcelFile) AddPictureFromBytes(sheet, cell string, pic *excelize.Picture) {
	fn.Panic(s.File.AddPictureFromBytes(sheet, cell, pic))
}

func (s *ExcelFile) DeletePicture(sheet, cell string) {
	fn.Panic(s.File.DeletePicture(sheet, cell))
}

func (s *ExcelFile) GetPictures(sheet, cell string) []excelize.Picture {
	return fn.Panic1(s.File.GetPictures(sheet, cell))
}

// endregion
// region pivot table
func (s *ExcelFile) AddPivotTable(opt *excelize.PivotTableOptions) {
	fn.Panic(s.File.AddPivotTable(opt))
}

// endregion
// region work book
func (s *ExcelFile) SetWorkbookProps(opt *excelize.WorkbookPropsOptions) {
	fn.Panic(s.File.SetWorkbookProps(opt))
}

func (s *ExcelFile) GetWorkbookProps() *excelize.WorkbookPropsOptions {
	e := fn.Panic1(s.File.GetWorkbookProps())
	return &e
}

func (s *ExcelFile) ProtectWorkbook(opt *excelize.WorkbookProtectionOptions) {
	fn.Panic(s.File.ProtectWorkbook(opt))
}

func (s *ExcelFile) UnprotectWorkbook(password goja.Value) {
	if engine.IsNullish(password) {
		fn.Panic(s.File.UnprotectWorkbook())
	} else {
		fn.Panic(s.File.UnprotectWorkbook(password.ToString().String()))
	}
}

//endregion
