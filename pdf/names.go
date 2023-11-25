package pdf

// Trailer
const (
	trailer   = "trailer"
	xref      = "xref"
	startxref = "startxref"
	root      = "Root"
	size      = "Size"
	info      = "Info"
	id        = "ID"
)

// Document Information
const (
	title        = "Title"
	subject      = "Subject"
	keywords     = "Keywords"
	author       = "Author"
	creator      = "Creator"
	producer     = "Producer"
	creationDate = "CreationDate"
	modDate      = "ModDate"
)

// Catalog
const (
	catalog           = "Catalog"
	pages             = "Pages"
	pageLabels        = "PageLabels"
	names             = "Names"
	dests             = "Dests"
	viewerPreferences = "ViewerPreferences"
	pageLayout        = "PageLayout"
	pageMode          = "PageMode"
	outlines          = "Outlines"
	metadata          = "Metadata"
)

// Page
const (
	parent    = "Parent"
	resources = "Resources"
	contents  = "Contents"
	rotate    = "Rotate"
	page      = "Page"
	mediaBox  = "MediaBox"
	cropBox   = "CropBox"
)

// Page Tree
const (
	kids  = "Kids"
	count = "Count"
)

const (
	font     = "Font"
	baseFont = "BaseFont"
	subtype  = "Subtype"

	prev             = "Prev"
	linearized       = "Linearized"
	annotation       = "Annot"
	action           = "Action"
	uRI              = "URI"
	structParent     = "StructParent"
	rect             = "Rect"
	border           = "Border"
	dest             = "Dest"
	filter           = "Filter"
	image            = "Image"
	colorSpace       = "ColorSpace"
	height           = "Height"
	width            = "Width"
	bitsPerComponent = "BitsPerComponent"
	xObject          = "XObject"
	deviceRGB        = "DeviceRGB"
	sMask            = "SMask"

	toUnicode       = "ToUnicode"
	encoding        = "Encoding"
	descendantFonts = "DescendantFonts"

	typeName = "Type"

	length = "Length"
)

// Filters
const (
	ASCIIHexDecode  = "ASCIIHexDecode"
	ASCII85Decode   = "ASCII85Decode"
	LZWDecode       = "LZWDecode"
	FlateDecode     = "FlateDecode"
	RunLengthDecode = "RunLengthDecode"
	CCITTFaxDecode  = "CCITTFaxDecode"
	JBIG2Decode     = "JBIG2Decode"
	DCTDecode       = "DCTDecode"
	JPXDecode       = "JPXDecode"
)

// Binary
var (
	space        = []byte{' '}
	trueValue    = []byte("true")
	falseValue   = []byte("false")
	dictBegin    = []byte("<< ")
	dictEnd      = []byte(">>")
	arrayBegin   = []byte{'['}
	arrayEnd     = []byte{']'}
	textBegin    = []byte{'('}
	textEnd      = []byte{')'}
	streamBegin  = []byte("stream\n")
	streamEnd    = []byte("\nendstream")
	objectBegin  = []byte("obj\n")
	objectEnd    = []byte("endobj")
	escape       = []byte{'\\'}
	lineFeed     = []byte{'\n'}
	escLineFeed  = []byte{'\\', 'n'}
	escReturns   = []byte{'\\', 'r'}
	escTab       = []byte{'\\', 't'}
	escBackspace = []byte{'\\', 'b'}
	escFormFeed  = []byte{'\\', 'f'}
	null         = []byte("null")
)
