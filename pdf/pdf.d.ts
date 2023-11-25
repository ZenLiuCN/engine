declare type Orientation='portrait' | 'landscape'
declare interface SizeType{Wd:number,Ht:number}
declare interface PointType{X:number,Y:number}
declare class PDFWriter {
    constructor(
        orientation: Orientation,
        unit: 'pt' | 'mm' | 'cm' | 'inch',
        size: 'A3' | 'A4' | 'A5' | 'Letter' | 'Legal' | 'Tabloid',
        fontDirectory: string
    )
    constructor(conf: PDFConfig)

    pointConvert(pt: number): number
    pointToUnitConvert(pt: number): number
    unitToPointConvert(pt: number): number
    addFont(family, style, file: string)
    addFontFromBytes(family, style :string, jsonFileBytes, zFileBytes :Uint8Array)
    // AddFontFromReader(familyStr, style :string, r io.Reader)
    addLayer(name:string, visible:boolean):number
    addLink():number
    addPage()
    addPageFormat(orientation:Orientation, size :SizeType)
    addSpotColor(name: string, c, m, y, k :number)
    aliasNbPages(alias:string)
    arcTo(x, y, rx, ry, degRotate, degStart, degEnd :number)
    arc(x, y, rx, ry, degRotate, degStart, degEnd :number, style::string)
    beginLayer(id :number)
    beziergon(points :PointType[], style: string)
    bookmark(txt: string, level :number, y :number)
    cellFormat(w, h :number, txtStr, border: string, ln :number, align:string, fill :boolean, link :number, linkStr :string)
    cellf(w, h :number, fmt:string,  ...args:any[])
cell(w, h :number, txt :string)
circle(x, y, r :number, style :string)
// ClearError()
clipCircle(x, y, r :number, outline :boolean)
clipEllipse(x, y, rx, ry :number, outline :boolean)
clipEnd()
clipPolygon(points :PointType[], outline :boolean)
clipRect(x, y, w, h :number, outline :boolean)
clipRoundedRect(x, y, w, h, r :number, outline :boolean)
clipText(x, y :number, txt :string, outline :boolean)
close()
closePath()
// CreateTemplateCustom(corner :PointType, size :SizeType, fn func(*Tpl)): Template
// CreateTemplate(fn func(*Tpl)) Template
curveBezierCubicTo(cx0, cy0, cx1, cy1, x, y :number)
curveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1 :number, styleStr :string)
curveCubic(x0, y0, cx0, cy0, x1, y1, cx1, cy1 :number, styleStr :string)
curveTo(cx, cy, x, y :number)
curve(x0, y0, cx, cy, x1, y1 :number, styleStr :string)
drawPath(styleStr :string)
ellipse(x, y, rx, ry, degRotate :number, styleStr :string)
endLayer()
err() :boolean

getAlpha(): {alpha :number, blendMode :string}
getAutoPageBreak() :{auto :boolean, margin :number}
GetCellMargin() :number
GetConversionRatio() :number
GetDrawColor() :{r:number, g:number, b:number}
GetDrawSpotColor() :{name :string, c, m, y, k :number}
GetFillColor() :{r:number, g:number, b:number}
GetFillSpotColor():{name :string, c, m, y, k :number}
GetFontDesc(familyStr, styleStr :string) FontDescType
GetFontSize() (ptSize, unitSize :number)
GetImageInfo(imageStr :string) (info *ImageInfoType)
GetLineWidth() :number
GetMargins() (left, top, right, bottom :number)
GetPageSizeStr(sizeStr :string) (size SizeType)
GetPageSize() (width, height :number)
GetStringWidth(s :string) :number
GetTextColor() (:number, :number, :number)
GetTextSpotColor() (name :string, c, m, y, k byte)
GetX() :number
GetXY() (:number, :number)
GetY() :number
HTMLBasicNew() (html HTMLBasicType)
Image(imageNameStr :string, x, y, w, h :number, flow :boolean, tp :string, link :number, linkStr :string)
ImageOptions(imageNameStr :string, x, y, w, h :number, flow :boolean, options ImageOptions, link :number, linkStr :string)
ImageTypeFromMime(mimeStr :string) (tp :string)
LinearGradient(x, y, w, h :number, r1, g1, b1, r2, g2, b2 :number, x1, y1, x2, y2 :number)
LineTo(x, y :number)
Line(x1, y1, x2, y2 :number)
LinkString(x, y, w, h :number, linkStr :string)
Link(x, y, w, h :number, link :number)
Ln(h :number)
MoveTo(x, y :number)
MultiCell(w, h :number, txtStr, borderStr, alignStr :string, fill :boolean)
Ok() :boolean
OpenLayerPane()
OutputAndClose(w io.WriteCloser) error
OutputFileAndClose(fileStr :string) error
Output(w io.Writer) error
PageCount() :number
PageNo() :number
PageSize(pageNum :number) (wd, ht :number, unitStr :string)
PointConvert(pt :number) (u :number)
PointToUnitConvert(pt :number) (u :number)
Polygon(points []PointType, styleStr :string)
RadialGradient(x, y, w, h :number, r1, g1, b1, r2, g2, b2 :number, x1, y1, x2, y2, r :number)
RawWriteBuf(r io.Reader)
RawWriteStr(str :string)
Rect(x, y, w, h :number, styleStr :string)
RegisterAlias(alias, replacement :string)
RegisterImage(fileStr, tp :string) (info *ImageInfoType)
RegisterImageOptions(fileStr :string, options ImageOptions) (info *ImageInfoType)
RegisterImageOptionsReader(imgName :string, options ImageOptions, r io.Reader) (info *ImageInfoType)
RegisterImageReader(imgName, tp :string, r io.Reader) (info *ImageInfoType)
SetAcceptPageBreakFunc(fnc func() :boolean)
SetAlpha(alpha :number, blendModeStr :string)
SetAuthor(authorStr :string, isUTF8 :boolean)
SetAutoPageBreak(auto :boolean, margin :number)
SetCatalogSort(flag :boolean)
SetCellMargin(margin :number)
SetCompression(compress :boolean)
SetCreationDate(tm time.Time)
SetCreator(creatorStr :string, isUTF8 :boolean)
SetDashPattern(dashArray []:number, dashPhase :number)
SetDisplayMode(zoomStr, layoutStr :string)
SetLang(lang :string)
SetDrawColor(r, g, b :number)
SetDrawSpotColor(nameStr :string, tint byte)
SetError(err error)
SetErrorf(fmtStr :string, args ...interface{})
SetFillColor(r, g, b :number)
SetFillSpotColor(nameStr :string, tint byte)
SetFont(familyStr, styleStr :string, size :number)
SetFontLoader(loader FontLoader)
SetFontLocation(fontDirStr :string)
SetFontSize(size :number)
SetFontStyle(styleStr :string)
SetFontUnitSize(size :number)
SetFooterFunc(fnc func())
SetFooterFuncLpi(fnc func(lastPage :boolean))
SetHeaderFunc(fnc func())
SetHeaderFuncMode(fnc func(), homeMode :boolean)
SetHomeXY()
SetJavascript(script :string)
SetKeywords(keywordsStr :string, isUTF8 :boolean)
SetLeftMargin(margin :number)
SetLineCapStyle(styleStr :string)
SetLineJoinStyle(styleStr :string)
SetLineWidth(width :number)
SetLink(link :number, y :number, page :number)
SetMargins(left, top, right :number)
SetPageBoxRec(t :string, pb PageBox)
SetPageBox(t :string, x, y, wd, ht :number)
SetPage(pageNum :number)
SetProtection(actionFlag byte, userPassStr, ownerPassStr :string)
SetRightMargin(margin :number)
SetSubject(subjectStr :string, isUTF8 :boolean)
SetTextColor(r, g, b :number)
SetTextSpotColor(nameStr :string, tint byte)
SetTitle(titleStr :string, isUTF8 :boolean)
SetTopMargin(margin :number)
SetUnderlineThickness(thickness :number)
SetXmpMetadata(xmpStream []byte)
SetX(x :number)
SetXY(x, y :number)
SetY(y :number)
SplitLines(txt []byte, w :number) [][]byte
String() :string
SVGBasicWrite(sb *SVGBasicType, scale :number)
Text(x, y :number, txtStr :string)
TransformBegin()
TransformEnd()
TransformMirrorHorizontal(x :number)
TransformMirrorLine(angle, x, y :number)
TransformMirrorPoint(x, y :number)
TransformMirrorVertical(y :number)
TransformRotate(angle, x, y :number)
TransformScale(scaleWd, scaleHt, x, y :number)
TransformScaleX(scaleWd, x, y :number)
TransformScaleXY(s, x, y :number)
TransformScaleY(scaleHt, x, y :number)
TransformSkew(angleX, angleY, x, y :number)
TransformSkewX(angleX, x, y :number)
TransformSkewY(angleY, x, y :number)
Transform(tm TransformMatrix)
TransformTranslate(tx, ty :number)
TransformTranslateX(tx :number)
TransformTranslateY(ty :number)
UnicodeTranslatorFromDescriptor(cpStr :string) (rep func(:string) :string)
UnitToPointConvert(u :number) (pt :number)
UseTemplateScaled(t Template, corner PointType, size SizeType)
UseTemplate(t Template)
WriteAligned(width, lineHeight :number, textStr, alignStr :string)
Writef(h :number, fmtStr :string, args ...interface{})
Write(h :number, txtStr :string)
WriteLinkID(h :number, displayStr :string, linkID :number)
WriteLinkString(h :number, displayStr, targetStr :string)
}

declare interface PDFConfig {
    orientation: 'portrait' | 'landscape',
    unit: 'pt' | 'mm' | 'cm' | 'inch',
    size?: 'A3' | 'A4' | 'A5' | 'Letter' | 'Legal' | 'Tabloid',
    fontDirectory: :string
    width?: number
    height?: number
}