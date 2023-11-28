declare module "go/execl" {

    export function open(path?: string, option?: Option): ExcelFile

    export function themeColor(baseColor: string, tint: number): string

    export function newDataValidation(allowBlank: boolean): DataValidation

    export function coordinatesToCellName(col, row: number, absCol, absRow: boolean): string

    export function cellNameToCoordinates(cell: string): { readonly col: number, readonly row: number }

    export function columnNameToNumber(name: string): number

    export function columnNumberToName(num: number): string

    export function splitCellName(cell: number): { readonly col: string, readonly row: number }

    export function joinCellName(col: string, row: number): string

    export interface Option {
        maxCalcIterations?: number
        password?: string
        rawCellValue?: boolean
        unzipSizeLimit?: number
        unzipXMLSizeLimit?: number
        shortDatePattern?: string
        longDatePattern?: string
        longTimePattern?: string
        /**
         *  0 unspecified,  1 EnUS, 2 ZhCN
         */
        CultureName?: 0 | 1 | 2
    }

    export interface ExcelFile {
        readonly workBook: WorkBook
        readonly path: string
        readonly sheetCount: number

        close()

        updateLinkedValue()

        addVBAProject(file: ArrayBuffer)

        save(opt?: Option)

        saveAs(name: string, opt?: Option)

        write(w: Buffer, opt?: Option)

        writeBinary(opt?: Option): Uint8Array

        newStreamWriter(sheet: string): StreamWriter

        addShape(sheet: string, shape: Shape)

        addSparkline(sheet: string, opt: SparklineOption)

        //region style
        getStyle(idx: number): Style

        newConditionalStyle(style: Style): number

        getDefaultFont(): string

        setDefaultFont(font: string)

        newStyle(style: Style): number

        getConditionalFormats(sheet: string): Record<string, ConditionalFormatOption[]>

        setConditionalFormat(sheet, rangeRef: string, opts: ConditionalFormatOption[])

        unsetConditionalFormat(sheet, rangeRef: string)

        //endregion
        //region vml
        getComments(sheet: string): Comment[]

        addComment(sheet: string, comment: Comment)

        deleteComment(sheet, cell: string)

        addFormControl(sheet: string, ctrl: FormControl)

        getFormControls(sheet: string): FormControl[]

        //endregion
        //region charts
        addChart(sheet, cell: string, chart: Chart, ...combo: Chart[])

        addChartSheet(sheet: string, chart: Chart, ...combo: Chart[])

        deleteChart(sheet, cell: string)

        //endregion
        //region table
        addTable(sheet: string, table: Table)

        getTables(sheet: string): Table[]

        deleteTable(name: string)

        autoFilter(sheet, range: string, opt: Array<AutoFilterOption>)

        //endregion
        //region calc
        calcCellValue(sheet, cell: string, opt?: Option): string

        //endregion
        //region sheet
        newSheet(sheet: string): number

        setSheetName(sheet, target: string)

        getSheetIndex(sheet: string): number

        getSheetMap(): { [key: number]: string }


        getSheetList(): string[]

        setSheetBackground(sheet: string, picturePath: string)

        setSheetBackgroundFromBytes(sheet, ext: string, picture: ArrayBuffer)

        deleteSheet(sheet: string)

        copySheet(from, to: number)

        setSheetVisible(sheet: string, visible: boolean, veryHidden?: boolean)

        getSheetVisible(sheet: string): boolean

        searchSheet(sheet, value: string, regex?: boolean): string[]

        setHeaderFooter(sheet: string, opt: HeaderFooterOption)

        protectSheet(sheet: string, opt: SheetProtectionOption)

        unprotectSheet(sheet: string, password?: string)

        setPageLayout(sheet: string, opt?: PageLayoutOption)

        getPageLayout(sheet: string): PageLayoutOption

        setDefinedName(definedName: DefinedName)

        deleteDefinedName(definedName: DefinedName)

        getDefinedName(): DefinedName[]

        groupSheets(sheets: string[])

        ungroupSheets()

        insertPageBreak(sheet, cell: string)

        removePageBreak(sheet, cell: string)

        setSheetDimension(sheet, rangeRef: string)

        getSheetDimension(sheet: string): string

        setPageMargins(sheet: string, opt: PageLayoutMarginsOption)

        getPageMargins(sheet: string): PageLayoutMarginsOption

        setSheetProps(sheet: string, prop: SheetPropsOption)

        getSheetProps(sheet: string): SheetPropsOption

        setSheetView(sheet: string, viewIndex: number, prop: ViewOption)

        getSheetView(sheet: string, viewIndex: number): ViewOption

        setPanes(sheet: string, panes: Panes)

        getPanes(sheet: string): Panes

        getSheetName(index: number): string

        setActiveSheet(index: number)

        getActiveSheetIndex(): number

        //endregion
        //region rows
        getRows(sheet: string, opt?: Option): string[][]

        rows(sheet: string): Rows

        setRowHeight(sheet: string, row: number, height: number)

        getRowHeight(sheet: string, row: number): number

        setRowVisible(sheet: string, row: number, visible: boolean)

        setRowOutlineLevel(sheet: string, row: number, level: number)

        getRowOutlineLevel(sheet: string, row: number): number

        removeRow(sheet: string, row: number)

        insertRows(sheet: string, row, n: number)

        duplicateRow(sheet: string, row: number)

        duplicateRowTo(sheet: string, rowFrom, rowTo: number)

        setRowStyle(sheet: string, start, end, style: number)

        getRowVisible(sheet: string, row: number): boolean

        //endregion
        //region cols
        getCols(sheet: string, opt?: Option): string[][]

        cols(sheet: string): Cols

        setColVisible(sheet, col: string, visible: boolean)

        getColVisible(sheet, col: string): boolean

        getColOutlineLevel(sheet, col: string): number //int8
        setColOutlineLevel(sheet, col: string, level: number) //int8
        setColStyle(sheet, col: string, styleId: number) //int
        getColStyle(sheet, col: string): number //int
        setColWidth(sheet, colStart, colEnd: string, width: number)

        getColWidth(sheet, col: string): number

        insertCols(sheet, col: string, n: number) //int
        removeCol(sheet, col: string)

        //endregion
        //region data validation
        addDataValidation(sheet: string, dv: DataValidation);

        getDataValidations(sheet: string): DataValidation[]

        deleteDataValidation(sheet: string, ...sqref: string[])

        //endregion
        //region properties
        setAppProps(prop: AppProperties)

        getAppProps(): AppProperties

        setDocProps(prop: DocProperties)

        getDocProps(): DocProperties

        //endregion
        //region cells
        getCellValue(sheet, cell: string, opt?: Option): string

        setCellValue(sheet, cell: string, value: number | boolean | string | null)

        setCellInt(sheet, cell: string, value: number)

        setCellBool(sheet, cell: string, value: boolean)

        setCellFloat(sheet, cell: string, value: number, precision, bitSize: number)

        setCellStr(sheet, cell: string, value: string)

        setCellDefault(sheet, cell: string, value: string)

        getCellHyperLink(sheet, cell: string): { link: boolean, target: string }

        setCellHyperLink(sheet, cell, link, linkType: string, opt?: HyperLinkOption)

        getCellRichText(sheet, cell: string): RichText[]

        setCellRichText(sheet, cell: string, runs: RichText[])

        // setSheetRow(sheet, cell: string, values: Array<number | boolean | string | null>)

        // setSheetCol(sheet, cell: string, values: Array<number | boolean | string | null>)

        getCellFormula(sheet, cell: string): string

        setCellFormula(sheet, cell: string, formula: string, opt?: FormulaOption)

        setCellStyle(sheet, hCell, vCell: string, style: number)

        getCellStyle(sheet, cell: string): number


        getCellType(sheet, cell: string): number //CellType

        mergeCell(sheet, hCell, vCell: string)

        unmergeCell(sheet, hCell, vCell: string)

        getMergeCells(sheet: string): MergeCell[]

        //endregion
        //region picture
        addPicture(sheet, cell, name: string, opt?: GraphicOption)


        addPictureFromBytes(sheet, cell: string, pic: Picture)

        deletePicture(sheet, cell: string)

        getPictures(sheet, cell: string): Picture[]

        //endregion
        //region pivot table
        addPivotTable(opt: PivotTableOption)

        //endregion
        //region work book
        setWorkbookProps(opt: WorkBookPropsOption)

        getWorkbookProps(): WorkBookPropsOption

        protectWorkbook(opt: WorkBookProtectionOption)

        unprotectWorkbook(password?: string)

        //endregion
    }

    export interface HyperLinkOption {
        display?: string
        tooltip?: string
    }

    export interface FormulaOption {
        type?: "array" | "dataTable" | "normal" | "shared" // Formula type
        ref?: string // Shared formula ref
    }

    export interface WorkBookPropsOption {
        date1904?: boolean
        filterPrivacy?: boolean
        codeName?: string
    }

    export interface WorkBookProtectionOption {
        algorithmName: string
        password: string
        lockStructure: boolean
        lockWindows: boolean
    }

    export interface FormControl {
        Cell: string
        Macro: string
        Width: number//uint
        Height: number//uint
        Checked: boolean
        CurrentVal: number//uint
        MinVal: number//uint
        MaxVal: number//uint
        IncChange: number//uint
        PageChange: number//uint
        Horizontally: boolean
        CellLink: string
        Text: string
        Paragraph: RichText[]
        /**
         *     Note,
         *     Button,
         *     OptionButton,
         *     SpinButton,
         *     CheckBox,
         *     GroupBox,
         *     Label,
         *     ScrollBar,
         */
        Type: 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7
        Format: GraphicOption
    }

    export interface Comment {
        author: string
        authorID: number
        cell: string
        text: string
        paragraph: RichText[]
    }

    export interface Style {
        border?: Border[]
        fill?: Fill
        font?: Font
        alignment?: Alignment
        protection?: Protection
        numFmt?: number//int
        decimalPlaces?: number//int
        customNumFmt?: string
        negRed?: boolean
    }

    export interface Alignment {
        horizontal?: 'left' | 'center' | 'right' | 'fill' | 'justify' | 'centerContinuous' | 'distributed'
        indent?: number//int
        justifyLastLine?: boolean
        readingOrder?: 0 | 1 | 2//uint64
        relativeIndent?: number//int
        shrinkToFit?: boolean
        textRotation?: number//int
        vertical?: 'center' | 'top' | 'justify' | 'distributed'
        wrapText?: boolean
    }

    export interface Border {
        type: 'left' | 'top' | 'right' | 'bottom' | 'diagonalDown' | 'diagonalUp'
        color?: string
        /**
         * Index | Name          | Weight | Style
         *
         * ------+---------------+--------+-------------
         *
         * 0     | None          | 0      |
         *
         * 1     | Continuous    | 1      | -----------
         *
         * 2     | Continuous    | 2      | -----------
         *
         * 3     | Dash          | 1      | - - - - - -
         *
         * 4     | Dot           | 1      | . . . . . .
         *
         * 5     | Continuous    | 3      | -----------
         *
         * 6     | Double        | 3      | ===========
         *
         * 7     | Continuous    | 0      | -----------
         *
         * 8     | Dash          | 2      | - - - - - -
         *
         * 9     | Dash Dot      | 1      | - . - . - .
         *
         * 10    | Dash Dot      | 2      | - . - . - .
         *
         * 11    | Dash Dot Dot  | 1      | - . . - . .
         *
         * 12    | Dash Dot Dot  | 2      | - . . - . .
         *
         * 13    | SlantDash Dot | 2      | / - . / - .
         */
        style?: 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13//int
    }

    export interface Protection {
        hidden: boolean
        locked: boolean
    }

    export interface DefinedName {
        name: string
        comment?: string
        refersTo?: string
        scope: string
    }

    export interface Shape {
        cell: string
        /**
         *
         * accentBorderCallout1 (Callout 1 with Border and Accent Shape)
         *
         * accentBorderCallout2 (Callout 2 with Border and Accent Shape)
         *
         * accentBorderCallout3 (Callout 3 with Border and Accent Shape)
         *
         * accentCallout1 (Callout 1 Shape)
         *
         * accentCallout2 (Callout 2 Shape)
         *
         * accentCallout3 (Callout 3 Shape)
         *
         * actionButtonBackPrevious (Back or Previous Button Shape)
         *
         * actionButtonBeginning (Beginning Button Shape)
         *
         * actionButtonBlank (Blank Button Shape)
         *
         * actionButtonDocument (Document Button Shape)
         *
         * actionButtonEnd (End Button Shape)
         *
         * actionButtonForwardNext (Forward or Next Button Shape)
         *
         * actionButtonHelp (Help Button Shape)
         *
         * actionButtonHome (Home Button Shape)
         *
         * actionButtonInformation (Information Button Shape)
         *
         * actionButtonMovie (Movie Button Shape)
         *
         * actionButtonReturn (Return Button Shape)
         *
         * actionButtonSound (Sound Button Shape)
         *
         * arc (Curved Arc Shape)
         *
         * bentArrow (Bent Arrow Shape)
         *
         * bentConnector2 (Bent Connector 2 Shape)
         *
         * bentConnector3 (Bent Connector 3 Shape)
         *
         * bentConnector4 (Bent Connector 4 Shape)
         *
         * bentConnector5 (Bent Connector 5 Shape)
         *
         * bentUpArrow (Bent Up Arrow Shape)
         *
         * bevel (Bevel Shape)
         *
         * blockArc (Block Arc Shape)
         *
         * borderCallout1 (Callout 1 with Border Shape)
         *
         * borderCallout2 (Callout 2 with Border Shape)
         *
         * borderCallout3 (Callout 3 with Border Shape)
         *
         * bracePair (Brace Pair Shape)
         *
         * bracketPair (Bracket Pair Shape)
         *
         * callout1 (Callout 1 Shape)
         *
         * callout2 (Callout 2 Shape)
         *
         * callout3 (Callout 3 Shape)
         *
         * can (Can Shape)
         *
         * chartPlus (Chart Plus Shape)
         *
         * chartStar (Chart Star Shape)
         *
         * chartX (Chart X Shape)
         *
         * chevron (Chevron Shape)
         *
         * chord (Chord Shape)
         *
         * circularArrow (Circular Arrow Shape)
         *
         * cloud (Cloud Shape)
         *
         * cloudCallout (Callout Cloud Shape)
         *
         * corner (Corner Shape)
         *
         * cornerTabs (Corner Tabs Shape)
         *
         * cube (Cube Shape)
         *
         * curvedConnector2 (Curved Connector 2 Shape)
         *
         * curvedConnector3 (Curved Connector 3 Shape)
         *
         * curvedConnector4 (Curved Connector 4 Shape)
         *
         * curvedConnector5 (Curved Connector 5 Shape)
         *
         * curvedDownArrow (Curved Down Arrow Shape)
         *
         * curvedLeftArrow (Curved Left Arrow Shape)
         *
         * curvedRightArrow (Curved Right Arrow Shape)
         *
         * curvedUpArrow (Curved Up Arrow Shape)
         *
         * decagon (Decagon Shape)
         *
         * diagStripe (Diagonal Stripe Shape)
         *
         * diamond (Diamond Shape)
         *
         * dodecagon (Dodecagon Shape)
         *
         * donut (Donut Shape)
         *
         * doubleWave (Double Wave Shape)
         *
         * downArrow (Down Arrow Shape)
         *
         * downArrowCallout (Callout Down Arrow Shape)
         *
         * ellipse (Ellipse Shape)
         *
         * ellipseRibbon (Ellipse Ribbon Shape)
         *
         * ellipseRibbon2 (Ellipse Ribbon 2 Shape)
         *
         * flowChartAlternateProcess (Alternate Process Flow Shape)
         *
         * flowChartCollate (Collate Flow Shape)
         *
         * flowChartConnector (Connector Flow Shape)
         *
         * flowChartDecision (Decision Flow Shape)
         *
         * flowChartDelay (Delay Flow Shape)
         *
         * flowChartDisplay (Display Flow Shape)
         *
         * flowChartDocument (Document Flow Shape)
         *
         * flowChartExtract (Extract Flow Shape)
         *
         * flowChartInputOutput (Input Output Flow Shape)
         *
         * flowChartInternalStorage (Internal Storage Flow Shape)
         *
         * flowChartMagneticDisk (Magnetic Disk Flow Shape)
         *
         * flowChartMagneticDrum (Magnetic Drum Flow Shape)
         *
         * flowChartMagneticTape (Magnetic Tape Flow Shape)
         *
         * flowChartManualInput (Manual Input Flow Shape)
         *
         * flowChartManualOperation (Manual Operation Flow Shape)
         *
         * flowChartMerge (Merge Flow Shape)
         *
         * flowChartMultidocument (Multi-Document Flow Shape)
         *
         * flowChartOfflineStorage (Offline Storage Flow Shape)
         *
         * flowChartOffpageConnector (Off-Page Connector Flow Shape)
         *
         * flowChartOnlineStorage (Online Storage Flow Shape)
         *
         * flowChartOr (Or Flow Shape)
         *
         * flowChartPredefinedProcess (Predefined Process Flow Shape)
         *
         * flowChartPreparation (Preparation Flow Shape)
         *
         * flowChartProcess (Process Flow Shape)
         *
         * flowChartPunchedCard (Punched Card Flow Shape)
         *
         * flowChartPunchedTape (Punched Tape Flow Shape)
         *
         * flowChartSort (Sort Flow Shape)
         *
         * flowChartSummingJunction (Summing Junction Flow Shape)
         *
         * flowChartTerminator (Terminator Flow Shape)
         *
         * foldedCorner (Folded Corner Shape)
         *
         * frame (Frame Shape)
         *
         * funnel (Funnel Shape)
         *
         * gear6 (Gear 6 Shape)
         *
         * gear9 (Gear 9 Shape)
         *
         * halfFrame (Half Frame Shape)
         *
         * heart (Heart Shape)
         *
         * heptagon (Heptagon Shape)
         *
         * hexagon (Hexagon Shape)
         *
         * homePlate (Home Plate Shape)
         *
         * horizontalScroll (Horizontal Scroll Shape)
         *
         * irregularSeal1 (Irregular Seal 1 Shape)
         *
         * irregularSeal2 (Irregular Seal 2 Shape)
         *
         * leftArrow (Left Arrow Shape)
         *
         * leftArrowCallout (Callout Left Arrow Shape)
         *
         * leftBrace (Left Brace Shape)
         *
         * leftBracket (Left Bracket Shape)
         *
         * leftCircularArrow (Left Circular Arrow Shape)
         *
         * leftRightArrow (Left Right Arrow Shape)
         *
         * leftRightArrowCallout (Callout Left Right Arrow Shape)
         *
         * leftRightCircularArrow (Left Right Circular Arrow Shape)
         *
         * leftRightRibbon (Left Right Ribbon Shape)
         *
         * leftRightUpArrow (Left Right Up Arrow Shape)
         *
         * leftUpArrow (Left Up Arrow Shape)
         *
         * lightningBolt (Lightning Bolt Shape)
         *
         * line (Line Shape)
         *
         * lineInv (Line Inverse Shape)
         *
         * mathDivide (Divide Math Shape)
         *
         * mathEqual (Equal Math Shape)
         *
         * mathMinus (Minus Math Shape)
         *
         * mathMultiply (Multiply Math Shape)
         *
         * mathNotEqual (Not Equal Math Shape)
         *
         * mathPlus (Plus Math Shape)
         *
         * moon (Moon Shape)
         *
         * nonIsoscelesTrapezoid (Non-Isosceles Trapezoid Shape)
         *
         * noSmoking (No Smoking Shape)
         *
         * notchedRightArrow (Notched Right Arrow Shape)
         *
         * octagon (Octagon Shape)
         *
         * parallelogram (Parallelogram Shape)
         *
         * pentagon (Pentagon Shape)
         *
         * pie (Pie Shape)
         *
         * pieWedge (Pie Wedge Shape)
         *
         * plaque (Plaque Shape)
         *
         * plaqueTabs (Plaque Tabs Shape)
         *
         * plus (Plus Shape)
         *
         * quadArrow (Quad-Arrow Shape)
         *
         * quadArrowCallout (Callout Quad-Arrow Shape)
         *
         * rect (Rectangle Shape)
         *
         * ribbon (Ribbon Shape)
         *
         * ribbon2 (Ribbon 2 Shape)
         *
         * rightArrow (Right Arrow Shape)
         *
         * rightArrowCallout (Callout Right Arrow Shape)
         *
         * rightBrace (Right Brace Shape)
         *
         * rightBracket (Right Bracket Shape)
         *
         * round1Rect (One Round Corner Rectangle Shape)
         *
         * round2DiagRect (Two Diagonal Round Corner Rectangle Shape)
         *
         * round2SameRect (Two Same-side Round Corner Rectangle Shape)
         *
         * roundRect (Round Corner Rectangle Shape)
         *
         * rtTriangle (Right Triangle Shape)
         *
         * smileyFace (Smiley Face Shape)
         *
         * snip1Rect (One Snip Corner Rectangle Shape)
         *
         * snip2DiagRect (Two Diagonal Snip Corner Rectangle Shape)
         *
         * snip2SameRect (Two Same-side Snip Corner Rectangle Shape)
         *
         * snipRoundRect (One Snip One Round Corner Rectangle Shape)
         *
         * squareTabs (Square Tabs Shape)
         *
         * star10 (Ten Pointed Star Shape)
         *
         * star12 (Twelve Pointed Star Shape)
         *
         * star16 (Sixteen Pointed Star Shape)
         *
         * star24 (Twenty Four Pointed Star Shape)
         *
         * star32 (Thirty Two Pointed Star Shape)
         *
         * star4 (Four Pointed Star Shape)
         *
         * star5 (Five Pointed Star Shape)
         *
         * star6 (Six Pointed Star Shape)
         *
         * star7 (Seven Pointed Star Shape)
         *
         * star8 (Eight Pointed Star Shape)
         *
         * straightConnector1 (Straight Connector 1 Shape)
         *
         * stripedRightArrow (Striped Right Arrow Shape)
         *
         * sun (Sun Shape)
         *
         * swooshArrow (Swoosh Arrow Shape)
         *
         * teardrop (Teardrop Shape)
         *
         * trapezoid (Trapezoid Shape)
         *
         * triangle (Triangle Shape)
         *
         * upArrow (Up Arrow Shape)
         *
         * upArrowCallout (Callout Up Arrow Shape)
         *
         * upDownArrow (Up Down Arrow Shape)
         *
         * upDownArrowCallout (Callout Up Down Arrow Shape)
         *
         * uturnArrow (U-Turn Arrow Shape)
         *
         * verticalScroll (Vertical Scroll Shape)
         *
         * wave (Wave Shape)
         *
         * wedgeEllipseCallout (Callout Wedge Ellipse Shape)
         *
         * wedgeRectCallout (Callout Wedge Rectangle Shape)
         *
         * wedgeRoundRectCallout (Callout Wedge Round Rectangle Shape)
         */
        type: string
        macro: string
        width: number//uint
        height: number//uint
        format: GraphicOption
        fill: Fill
        line: Line
        paragraph: RichText[]
    }

    export interface StreamWriter {
        addTable(table: Table)

        setRow(cell: string, value: any[], opt?: RowOption)

        setColWidth(min, max, width: number) //int,int,float64
        insertPageBreak(cell: string)

        setPanes(panes: Panes)

        mergeCell(hCell, vCell: string)

        flush()
    }

    export interface Panes {
        freeze: boolean
        split: boolean
        xSplit: number
        ySplit: number
        topLeftCell: string
        activePane: string
        selection: Selection[]
    }

    /**
     *  Fields           | Description
     *
     * ------------------+-----------------------------------------------------------
     *
     *  AlignWithMargins | Align header footer margins with page margins
     *
     *  DifferentFirst   | Different first-page header and footer indicator
     *
     *  DifferentOddEven | Different odd and even page headers and footers indicator
     *
     *  ScaleWithDoc     | Scale header and footer with document scaling
     *
     *  OddFooter        | Odd Page Footer
     *
     *  OddHeader        | Odd Header
     *
     *  EvenFooter       | Even Page Footer
     *
     *  EvenHeader       | Even Page Header
     *
     *  FirstFooter      | First Page Footer
     *
     *  FirstHeader      | First Page Header
     *
     * ------------------+-----------------------------------------------------------
     *
     *
     * The following formatting codes can be used in 6 string type fields:
     * OddHeader, OddFooter, EvenHeader, EvenFooter, FirstFooter, FirstHeader
     *
     *    ------------------+-----------------------------------------------------------
     *     Formatting Code        | Description
     *
     *    ------------------------+------------------------------------------------------------------------
     *
     *     &&                     | The character "&"
     *
     *     &font-size             | Size of the text font, where font-size is a decimal font size in points
     *
     *     &"font name,font type" | A text font-name string, font name, and a text font-type string,font type
     *
     *     &"-,Regular"           | Regular text format. Toggles bold and italic modes to off
     *
     *     &A                     | Current worksheet's tab name

     *
     *     &B or &"-,Bold"        | Bold text format, from off to on, or vice versa. The default mode is of
     *
     *     &D                     | Current date
     *
     *     &C                     | Center section
     *
     *     &E                     | Double-underline text format
     *
     *     &F                     | Current workbook's file name
     *
     *     &G                     | Drawing object as background (Not support currently)
     *
     *     &H                     | Shadow text format
     *
     *     &I or &"-,Italic"      | Italic text format
     *
     *     &K                     | Text font color, An RGB Color is specified as RRGGBB, A Theme Color is specified as TTSNNN where TT is the theme color Id , S is either "+" or "-" of the tint/shade value, and NNN is the tint/shade value
     *
     *     &L                     | Left section
     *
     *     &N                     | Total number of pages
     *
     *     &O                     | Outline text format
     *
     *     &P[[+|-]n]             | Without the optional suffix, the current page number in decimal
     *
     *     &R                     | Right section
     *
     *     &S                     | Strike through text format
     *
     *     &T                     | Current time
     *
     *     &U                     | Single-underline text format. If double-underline mode is on, the next, occurrence in a section specifier toggles double-underline mode to off;, otherwise, it toggles single-underline mode, from off to on, or vice, versa. The default mode is off,
     *
     *     &X                     | Superscript text format
     *
     *     &Y                     | Subscript text format
     *
     *     &Z                     | Current workbook's file path
     *
     *     -----------------------+-----------------------------------------------------------
     */
    export interface HeaderFooterOption {
        alignWithMargins?: boolean
        differentFirst?: boolean
        differentOddEven?: boolean
        scaleWithDoc?: boolean
        oddHeader?: string
        oddFooter?: string
        evenHeader?: string
        evenFooter?: string
        firstHeader?: string
        firstFooter?: string
    }

    export interface SheetProtectionOption {
        algorithmName?: 'XOR' | 'MD4' | 'MD5' | 'SHA-1' | 'SHA-256' | 'SHA-384' | 'SHA-512'
        autoFilter?: boolean
        deleteColumns?: boolean
        deleteRows?: boolean
        editObjects?: boolean
        editScenarios?: boolean
        formatCells?: boolean
        formatColumns?: boolean
        formatRows?: boolean
        insertColumns?: boolean
        insertHyperlinks?: boolean
        insertRows?: boolean
        password?: string
        pivotTables?: boolean
        selectLockedCells?: boolean
        selectUnlockedCells?: boolean
        sort?: boolean
    }

    export interface ConditionalFormatOption {
        /**
         *  Type          | Parameters
         *
         * ---------------+---------------
         *
         *  cell          | Criteria
         *                | Value
         *                | MinValue
         *                | MaxValue
         *
         *  date          | Criteria
         *                | Value
         *                | MinValue
         *                | MaxValue
         *
         *  time_period   | Criteria
         *
         *  text          | Criteria
         *                | Value
         *
         *  average       | Criteria
         *
         *  duplicate     | (none)
         *
         *  unique        | (none)
         *
         *  top           | Criteria
         *                | Value
         *
         *  bottom        | Criteria
         *                | Value
         *
         *  blanks        | (none)
         *
         *  no_blanks     | (none)
         *
         *  errors        | (none)
         *
         *  no_errors     | (none)
         *
         *  2_color_scale | MinType
         *                | MaxType
         *                | MinValue
         *                | MaxValue
         *                | MinColor
         *                | MaxColor
         *
         *  3_color_scale | MinType
         *                | MidType
         *                | MaxType
         *                | MinValue
         *                | MidValue
         *                | MaxValue
         *                | MinColor
         *                | MidColor
         *                | MaxColor
         *
         *  data_bar      | MinType
         *                | MaxType
         *                | MinValue
         *                | MaxValue
         *                | BarBorderColor
         *                | BarColor
         *                | BarDirection
         *                | BarOnly
         *                | BarSolid
         *
         *  icon_set      | IconStyle
         *                | ReverseIcons
         *                | IconsOnly
         *
         *  formula       | Criteria
         *
         */
        type: string
        aboveAverage?: boolean
        percent?: boolean
        format?: number //int
        /**
         * between
         *
         * not between
         *
         * equal to                 | ==
         *
         * not equal to             | !=
         *
         * greater than             | >
         *
         * less than                | <
         *
         * greater than or equal to | >=
         *
         * less than or equal to    | <=
         *
         */
        criteria?: 'between' | 'not between' | 'equal to' | '=' | 'not equal to' | '!=' | 'greater than' | '>' | 'less than' | '<' | 'greater than or equal to' | '>=' | "less than or equal to" | '<='
        value?: string

        minType?: 'min' | 'num' | 'percent' | 'percentile' | 'formula'
        midType?: 'num' | 'percent' | 'percentile' | 'formula'
        maxType?: 'num' | 'percent' | 'percentile' | 'formula' | 'max'
        minValue?: string
        midValue?: string
        maxValue?: string
        minColor?: string
        midColor?: string
        maxColor?: string
        barColor?: string
        barBorderColor?: string
        barDirection?: 'context' | 'leftToRight' | 'rightToLeft'
        barOnly?: boolean
        barSolid?: boolean
        iconStyle?: '3Arrows' | '3ArrowsGray' | '3Flags' | '3Signs' | '3Symbols' | '3Symbols2' | '3TrafficLights1' | '3TrafficLights2' | '4Arrows' | '4ArrowsGray' | '4Rating' | '4RedToBlack' | '4TrafficLights' | '5Arrows' | '5ArrowsGray' | '5Quarters' | '5Rating'
        reverseIcons?: boolean
        iconsOnly?: boolean
        stopIfTrue?: boolean
    }

    export interface SparklineOption {
        location: string[]
        range: string[]
        max?: number//int
        custMax?: number//int
        min?: number//int
        custMin?: number//int
        type?: 'line' | 'column' | 'win_loss'
        weight?: number//float64
        dateAxis?: boolean
        markers?: boolean
        high?: boolean
        low?: boolean
        first?: boolean
        last?: boolean
        negative?: boolean
        axis?: boolean
        hidden?: boolean
        reverse?: boolean
        /**
         * [0 ,35]
         */
        style?: number//int
        seriesColor?: string
        negativeColor?: string
        markersColor?: string
        firstColor?: string
        lastColor?: string
        hightColor?: string
        lowColor?: string
        emptyCells?: string
    }

    export interface PageLayoutOption {
        /**
         *  Index | Paper Size
         *
         * -------+-----------------------------------------------
         *
         *    1   | Letter paper (8.5 in. by 11 in.)
         *
         *    2   | Letter small paper (8.5 in. by 11 in.)
         *
         *    3   | Tabloid paper (11 in. by 17 in.)
         *
         *    4   | Ledger paper (17 in. by 11 in.)
         *
         *    5   | Legal paper (8.5 in. by 14 in.)
         *
         *    6   | Statement paper (5.5 in. by 8.5 in.)
         *
         *    7   | Executive paper (7.25 in. by 10.5 in.)
         *
         *    8   | A3 paper (297 mm by 420 mm)
         *
         *    9   | A4 paper (210 mm by 297 mm)
         *
         *    10  | A4 small paper (210 mm by 297 mm)
         *
         *    11  | A5 paper (148 mm by 210 mm)
         *
         *    12  | B4 paper (250 mm by 353 mm)
         *
         *    13  | B5 paper (176 mm by 250 mm)
         *
         *    14  | Folio paper (8.5 in. by 13 in.)
         *
         *    15  | Quarto paper (215 mm by 275 mm)
         *
         *    16  | Standard paper (10 in. by 14 in.)
         *
         *    17  | Standard paper (11 in. by 17 in.)
         *
         *    18  | Note paper (8.5 in. by 11 in.)
         *
         *    19  | #9 envelope (3.875 in. by 8.875 in.)
         *
         *    20  | #10 envelope (4.125 in. by 9.5 in.)
         *
         *    21  | #11 envelope (4.5 in. by 10.375 in.)
         *
         *    22  | #12 envelope (4.75 in. by 11 in.)
         *
         *    23  | #14 envelope (5 in. by 11.5 in.)
         *
         *    24  | C paper (17 in. by 22 in.)
         *
         *    25  | D paper (22 in. by 34 in.)
         *
         *    26  | E paper (34 in. by 44 in.)
         *
         *    27  | DL envelope (110 mm by 220 mm)
         *
         *    28  | C5 envelope (162 mm by 229 mm)
         *
         *    29  | C3 envelope (324 mm by 458 mm)
         *
         *    30  | C4 envelope (229 mm by 324 mm)
         *
         *    31  | C6 envelope (114 mm by 162 mm)
         *
         *    32  | C65 envelope (114 mm by 229 mm)
         *
         *    33  | B4 envelope (250 mm by 353 mm)
         *
         *    34  | B5 envelope (176 mm by 250 mm)
         *
         *    35  | B6 envelope (176 mm by 125 mm)
         *
         *    36  | Italy envelope (110 mm by 230 mm)
         *
         *    37  | Monarch envelope (3.875 in. by 7.5 in.).
         *
         *    38  | 6 3/4 envelope (3.625 in. by 6.5 in.)
         *
         *    39  | US standard fanfold (14.875 in. by 11 in.)
         *
         *    40  | German standard fanfold (8.5 in. by 12 in.)
         *
         *    41  | German legal fanfold (8.5 in. by 13 in.)
         *
         *    42  | ISO B4 (250 mm by 353 mm)
         *
         *    43  | Japanese postcard (100 mm by 148 mm)
         *
         *    44  | Standard paper (9 in. by 11 in.)
         *
         *    45  | Standard paper (10 in. by 11 in.)
         *
         *    46  | Standard paper (15 in. by 11 in.)
         *
         *    47  | Invite envelope (220 mm by 220 mm)
         *
         *    50  | Letter extra paper (9.275 in. by 12 in.)
         *
         *    51  | Legal extra paper (9.275 in. by 15 in.)
         *
         *    52  | Tabloid extra paper (11.69 in. by 18 in.)
         *
         *    53  | A4 extra paper (236 mm by 322 mm)
         *
         *    54  | Letter transverse paper (8.275 in. by 11 in.)
         *
         *    55  | A4 transverse paper (210 mm by 297 mm)
         *
         *    56  | Letter extra transverse paper (9.275 in. by 12 in.)
         *
         *    57  | SuperA/SuperA/A4 paper (227 mm by 356 mm)
         *
         *    58  | SuperB/SuperB/A3 paper (305 mm by 487 mm)
         *
         *    59  | Letter plus paper (8.5 in. by 12.69 in.)
         *
         *    60  | A4 plus paper (210 mm by 330 mm)
         *
         *    61  | A5 transverse paper (148 mm by 210 mm)
         *
         *    62  | JIS B5 transverse paper (182 mm by 257 mm)
         *
         *    63  | A3 extra paper (322 mm by 445 mm)
         *
         *    64  | A5 extra paper (174 mm by 235 mm)
         *
         *    65  | ISO B5 extra paper (201 mm by 276 mm)
         *
         *    66  | A2 paper (420 mm by 594 mm)
         *
         *    67  | A3 transverse paper (297 mm by 420 mm)
         *
         *    68  | A3 extra transverse paper (322 mm by 445 mm)
         *
         *    69  | Japanese Double Postcard (200 mm x 148 mm)
         *
         *    70  | A6 (105 mm x 148 mm)
         *
         *    71  | Japanese Envelope Kaku #2
         *
         *    72  | Japanese Envelope Kaku #3
         *
         *    73  | Japanese Envelope Chou #3
         *
         *    74  | Japanese Envelope Chou #4
         *
         *    75  | Letter Rotated (11in x 8 1/2 11 in)
         *
         *    76  | A3 Rotated (420 mm x 297 mm)
         *
         *    77  | A4 Rotated (297 mm x 210 mm)
         *
         *    78  | A5 Rotated (210 mm x 148 mm)
         *
         *    79  | B4 (JIS) Rotated (364 mm x 257 mm)
         *
         *    80  | B5 (JIS) Rotated (257 mm x 182 mm)
         *
         *    81  | Japanese Postcard Rotated (148 mm x 100 mm)
         *
         *    82  | Double Japanese Postcard Rotated (148 mm x 200 mm)
         *
         *    83  | A6 Rotated (148 mm x 105 mm)
         *
         *    84  | Japanese Envelope Kaku #2 Rotated
         *
         *    85  | Japanese Envelope Kaku #3 Rotated
         *
         *    86  | Japanese Envelope Chou #3 Rotated
         *
         *    87  | Japanese Envelope Chou #4 Rotated
         *
         *    88  | B6 (JIS) (128 mm x 182 mm)
         *
         *    89  | B6 (JIS) Rotated (182 mm x 128 mm)
         *
         *    90  | (12 in x 11 in)
         *
         *    91  | Japanese Envelope You #4
         *
         *    92  | Japanese Envelope You #4 Rotated
         *
         *    93  | PRC 16K (146 mm x 215 mm)
         *
         *    94  | PRC 32K (97 mm x 151 mm)
         *
         *    95  | PRC 32K(Big) (97 mm x 151 mm)
         *
         *    96  | PRC Envelope #1 (102 mm x 165 mm)
         *
         *    97  | PRC Envelope #2 (102 mm x 176 mm)
         *
         *    98  | PRC Envelope #3 (125 mm x 176 mm)
         *
         *    99  | PRC Envelope #4 (110 mm x 208 mm)
         *
         *    100 | PRC Envelope #5 (110 mm x 220 mm)
         *
         *    101 | PRC Envelope #6 (120 mm x 230 mm)
         *
         *    102 | PRC Envelope #7 (160 mm x 230 mm)
         *
         *    103 | PRC Envelope #8 (120 mm x 309 mm)
         *
         *    104 | PRC Envelope #9 (229 mm x 324 mm)
         *
         *    105 | PRC Envelope #10 (324 mm x 458 mm)
         *
         *    106 | PRC 16K Rotated
         *
         *    107 | PRC 32K Rotated
         *
         *    108 | PRC 32K(Big) Rotated
         *
         *    109 | PRC Envelope #1 Rotated (165 mm x 102 mm)
         *
         *    110 | PRC Envelope #2 Rotated (176 mm x 102 mm)
         *
         *    111 | PRC Envelope #3 Rotated (176 mm x 125 mm)
         *
         *    112 | PRC Envelope #4 Rotated (208 mm x 110 mm)
         *
         *    113 | PRC Envelope #5 Rotated (220 mm x 110 mm)
         *
         *    114 | PRC Envelope #6 Rotated (230 mm x 120 mm)
         *
         *    115 | PRC Envelope #7 Rotated (230 mm x 160 mm)
         *
         *    116 | PRC Envelope #8 Rotated (309 mm x 120 mm)
         *
         *    117 | PRC Envelope #9 Rotated (324 mm x 229 mm)
         *
         *    118 | PRC Envelope #10 Rotated (458 mm x 324 mm)
         */
        size?: number// *int
        orientation?: 'landscape' | 'portrait'// *string
        firstPageNumber?: number// *uint
        /**
         * 10 .. 400 (400%)
         */
        adjustTo?: number// *uint
        /**
         * Number of pages to fit
         */
        fitToHeight?: number// *int
        /**
         * Number of pages to fit
         */
        fitToWidth?: number// *int
        blackAndWhite?: boolean// *bool
    }

    export interface ViewOption {

        defaultGridColor?: boolean// *bool
        rightToLeft?: boolean// *bool
        showFormulas?: boolean// *bool
        showGridLines?: boolean// *bool
        showRowColHeaders?: boolean// *bool
        showRuler?: boolean// *bool
        showZeros?: boolean// *bool
        topLeftCell?: string// *string
        view?: '' | 'normal' | 'pageLayout' | 'pageBreakPreview'// *string
        /**
         * 10 to 400
         */
        zoomScale?: number// *float64
    }

    export interface SheetPropsOption {
        codeName?: string// *string
        enableFormatConditionsCalculation?: boolean// *bool
        published?: boolean// *bool
        autoPageBreaks?: boolean// *bool
        fitToPage?: boolean// *bool
        tabColorIndexed?: number// *int
        tabColorRGB?: string// *string
        tabColorTheme?: number// *int
        tabColorTint?: number// *float64
        outlineSummaryBelow?: boolean// *bool
        outlineSummaryRight?: boolean// *bool
        baseColWidth?: number// *uint8
        defaultColWidth?: number// *float64
        defaultRowHeight?: number// *float64
        customHeight?: boolean// *bool
        zeroHeight?: boolean// *bool
        thickTop?: boolean// *bool
        thickBottom?: boolean// *bool
    }

    export interface PageLayoutMarginsOption {
        bottom?: number//float64
        footer?: number//float64
        header?: number//float64
        left?: number//float64
        right?: number//float64
        top?: number//float64
        horizontally?: boolean//bool
        vertically?: boolean//bool
    }

    export interface Selection {
        sqRef: string
        activeCell: string
        pane: string
    }

    export interface Rows {
        next(): boolean

        getRowOpts(): RowOption

        error(): string

        close()

        columns(opt?: Option): string[]
    }

    export interface Cell {
        styleID: number //int
        formula: string
        value: any
    }

    export interface RowOption {
        height: number
        hidden: boolean
        styleID: number //int
        outlineLevel: number //int
    }

    export interface PivotTableOption {

        dataRange: string
        pivotTableRange: string
        rows: PivotTableField[]
        columns: PivotTableField[]
        data: PivotTableField[]
        filter: PivotTableField[]
        rowGrandTotals: boolean
        colGrandTotals: boolean
        showDrill: boolean
        useAutoFormatting: boolean
        pageOverThenDown: boolean
        mergeItem: boolean
        compactData: boolean
        showError: boolean
        showRowHeaders: boolean
        showColHeaders: boolean
        showRowStripes: boolean
        showColStripes: boolean
        showLastColumn: boolean
        pivotTableStyleName: string
    }

    export interface PivotTableField {
        compact: boolean
        data: string
        name: string
        outline: boolean
        subtotal: string
        defaultSubtotal: boolean
    }

    export interface Picture {
        extension: string
        file: ArrayBuffer
        format: GraphicOption

    }

    export interface MergeCell {
        getCellValue(): string

        getStartAxis(): string

        getEndAxis(): string
    }

    export interface DocProperties {
        category: string
        contentStatus: string
        created: string
        creator: string
        description: string
        identifier: string
        keywords: string
        lastModifiedBy: string
        modified: string
        revision: string
        subject: string
        title: string
        language: string
        version: string
    }

    export interface AppProperties {
        application: string
        scaleCrop: boolean
        docSecurity: number//int
        company: string
        linksUpToDate: boolean
        hyperlinksChanged: boolean
        appVersion: string
    }

    export interface DataValidation {
        allowBlank: boolean
        error?: string
        errorStyle?: string
        errorTitle?: string
        operator: string
        prompt?: string
        promptTitle?: string
        showDropDown: boolean
        showErrorMessage: boolean
        showInputMessage: boolean
        sqref: string
        type: string
        formula1: string
        formula2: string

        /**
         *
         * @param style 1:Stop 2:Warning 3:Information
         * @param title
         * @param message
         */
        setError(style: 1 | 2 | 3, title, message: string)

        setInput(title, message: string)

        setDropList(keys: string[])

        /**
         *
         * @param f1
         * @param f2
         * @param t 1:None,Custom,Date,Decimal,typeList,TextLength,Time,8:Whole
         * @param o 1:Between,Equal,GreaterThan,GreaterThanOrEqual,LessThan,LessThanOrEqual,NotBetween,8:NotEqual
         */
        setRange(f1, f2: string | number, t: 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8, o: 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8)

        setSqrefDropList(sqref: string)

        setSqref(sqref: string)
    }

    export interface Cols {
        next(): boolean

        error(): string

        rows(opt?: Option): string[]
    }

    export interface Table {
        range: string
        name?: string
        /**
         *  TableStyleLight1 ... TableStyleLight21
         *  TableStyleMedium1 ... TableStyleMedium28
         *  TableStyleDark1 ... TableStyleDark11
         */
        styleName?: string
        showColumnStripes?: boolean
        showFirstColumn?: boolean
        showHeaderRow?: boolean
        showLastColumn?: boolean
        showRowStripes?: boolean
    }

    export interface AutoFilterOption {
        column: string
        expression: string
    }

    export interface Chart {

        /**
         * ID | Enumeration                 | Chart
         *
         * ---+-----------------------------+------------------------------
         *
         * 0  | Area                        | 2D area chart
         *
         * 1  | AreaStacked                 | 2D stacked area chart
         *
         * 2  | AreaPercentStacked          | 2D 100% stacked area chart
         *
         * 3  | Area3D                      | 3D area chart
         *
         * 4  | Area3DStacked               | 3D stacked area chart
         *
         * 5  | Area3DPercentStacked        | 3D 100% stacked area chart
         *
         * 6  | Bar                         | 2D clustered bar chart
         *
         * 7  | BarStacked                  | 2D stacked bar chart
         *
         * 8  | BarPercentStacked           | 2D 100% stacked bar chart
         *
         * 9  | Bar3DClustered              | 3D clustered bar chart
         *
         * 10 | Bar3DStacked                | 3D stacked bar chart
         *
         * 11 | Bar3DPercentStacked         | 3D 100% stacked bar chart
         *
         * 12 | Bar3DConeClustered          | 3D cone clustered bar chart
         *
         * 13 | Bar3DConeStacked            | 3D cone stacked bar chart
         *
         * 14 | Bar3DConePercentStacked     | 3D cone percent bar chart
         *
         * 15 | Bar3DPyramidClustered       | 3D pyramid clustered bar chart
         *
         * 16 | Bar3DPyramidStacked         | 3D pyramid stacked bar chart
         *
         * 17 | Bar3DPyramidPercentStacked  | 3D pyramid percent stacked bar chart
         *
         * 18 | Bar3DCylinderClustered      | 3D cylinder clustered bar chart
         *
         * 19 | Bar3DCylinderStacked        | 3D cylinder stacked bar chart
         *
         * 20 | Bar3DCylinderPercentStacked | 3D cylinder percent stacked bar chart
         *
         * 21 | Col                         | 2D clustered column chart
         *
         * 22 | ColStacked                  | 2D stacked column chart
         *
         * 23 | ColPercentStacked           | 2D 100% stacked column chart
         *
         * 24 | Col3DClustered              | 3D clustered column chart
         *
         * 25 | Col3D                       | 3D column chart
         *
         * 26 | Col3DStacked                | 3D stacked column chart
         *
         * 27 | Col3DPercentStacked         | 3D 100% stacked column chart
         *
         * 28 | Col3DCone                   | 3D cone column chart
         *
         * 29 | Col3DConeClustered          | 3D cone clustered column chart
         *
         * 30 | Col3DConeStacked            | 3D cone stacked column chart
         *
         * 31 | Col3DConePercentStacked     | 3D cone percent stacked column chart
         *
         * 32 | Col3DPyramid                | 3D pyramid column chart
         *
         * 33 | Col3DPyramidClustered       | 3D pyramid clustered column chart
         *
         * 34 | Col3DPyramidStacked         | 3D pyramid stacked column chart
         *
         * 35 | Col3DPyramidPercentStacked  | 3D pyramid percent stacked column chart
         *
         * 36 | Col3DCylinder               | 3D cylinder column chart
         *
         * 37 | Col3DCylinderClustered      | 3D cylinder clustered column chart
         *
         * 38 | Col3DCylinderStacked        | 3D cylinder stacked column chart
         *
         * 39 | Col3DCylinderPercentStacked | 3D cylinder percent stacked column chart
         *
         * 40 | Doughnut                    | doughnut chart
         *
         * 41 | Line                        | line chart
         *
         * 42 | Line3D                      | 3D line chart
         *
         * 43 | Pie                         | pie chart
         *
         * 44 | Pie3D                       | 3D pie chart
         *
         * 45 | PieOfPie                    | pie of pie chart
         *
         * 46 | BarOfPie                    | bar of pie chart
         *
         * 47 | Radar                       | radar chart
         *
         * 48 | Scatter                     | scatter chart
         *
         * 49 | Surface3D                   | 3D surface chart
         *
         * 50 | WireframeSurface3D          | 3D wireframe surface chart
         *
         * 51 | Contour                     | contour chart
         *
         * 52 | WireframeContour            | wireframe contour chart
         *
         * 53 | Bubble                      | bubble chart
         *
         * 54 | Bubble3D                    | 3D bubble chart
         */
        type: number
        series: ChartSeries[]

        format: GraphicOption
        dimension: {
            width: number
            height: number
        }
        legend: {
            position: 'none' | 'top' | 'bottom' | 'left' | 'right' | 'top_right'
            showLegendKey: boolean
        }
        title?: RichText[]

        varyColors?: boolean
        xAxis: ChartAxis
        yAxis: ChartAxis
        plotArea: {
            secondPlotValues: number //int
            showBubbleSize: boolean
            showCatName: boolean
            showLeaderLines: boolean
            showPercent: boolean
            showSerName: boolean
            showVal: boolean
            numFmt: ChartNumFmt
        }
        showBlanksAs: 'gap' | 'span' | 'zero'
        holeSize: number //int
    }

    export interface RichText {
        font?: Font
        text: string
    }

    export interface Font {
        bold?: boolean
        italic?: boolean
        underline?: "none" | "single" | "double" | 'words' | 'sng' | 'dbl' | 'heavy' | 'dotted' | 'dottedHeavy' | 'dash' | 'dashHeavy' | 'dashLong' | 'dashLongHeavy' | 'dotDash' | 'dotDashHeavy' | 'dotDotDash' | 'dotDotDashHeavy' | 'wavy' | 'wavyHeavy' | 'wavyDbl'
        family?: string
        size?: number
        strike?: boolean
        color?: string
        colorIndexed?: number
        colorTheme?: number
        colorTint?: number
        vertAlign?: string
    }

    export interface ChartAxis {
        none: boolean
        majorGridLines: boolean
        minorGridLines: boolean
        majorUnit: number
        tickLabelSkip: number
        reverseOrder: boolean
        secondary: boolean
        maximum?: number
        minimum?: number
        font: Font
        logBase: number
        numFmt: ChartNumFmt
        title: RichText[]
    }

    export interface ChartNumFmt {
        customNumFmt: string
        sourceLinked: boolean
    }

    export interface GraphicOption {
        altText: string
        printObject: boolean
        locked: boolean
        lockAspectRatio: boolean
        autoFit: boolean
        offsetX: number
        offsetY: number
        scaleX: number
        scaleY: number
        hyperlink: string
        hyperlinkType: string
        positioning: string
    }

    export interface ChartSeries {
        /**
         * "Sheet1!$A$2" or just raw text
         */
        name: string
        /**
         * "Sheet1!$B$1:$D$1"
         */
        categories: string
        sizes?: string
        /**
         * "Sheet1!$B$2:$D$2"
         */
        values: string
        fill?: Fill
        line?: Line
        marker?: {
            symbol: 'circle' | 'dash' | 'diamond' | 'dot' | 'none' | 'picture' | 'plus' | 'square' | 'star' | 'triangle' | 'x' | 'auto'
            size: number//int
        }
    }

    export interface Line {
        smooth: boolean
        width: number
    }

    export interface Fill {
        type?: string
        /**
         *     Index | Style           | Index | Style
         *
         *    -------+-----------------+-------+-----------------
         *
         *     0     | None            | 10    | darkTrellis
         *
         *     1     | solid           | 11    | lightHorizontal
         *
         *     2     | mediumGray      | 12    | lightVertical
         *
         *     3     | darkGray        | 13    | lightDown
         *
         *     4     | lightGray       | 14    | lightUp
         *
         *     5     | darkHorizontal  | 15    | lightGrid
         *
         *     6     | darkVertical    | 16    | lightTrellis
         *
         *     7     | darkDown        | 17    | gray125
         *
         *     8     | darkUp          | 18    | gray0625
         *
         *     9     | darkGrid        |       |
         *
         *
         */
        pattern?: 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18
        color?: string[]
        /**
         *    Index | Style           | Index | Style
         *
         *   -------+-----------------+-------+-----------------
         *
         *    0-2   | Horizontal      | 9-11  | Diagonal down
         *
         *    3-5   | Vertical        | 12-15 | From corner
         *
         *    6-8   | Diagonal Up     | 16    | From center
         *
         */
        shading?: 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16
    }

    export interface WorkBook {
        readonly workBook

        close()

        addChart(sheet: string)
    }


}