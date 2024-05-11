// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/unicode'{
	// @ts-ignore
	import type {Ref,uint32,Struct,rune,bool,int,uint16} from 'go'
	export const ASCII_Hex_Digit:Ref<RangeTable>
	export const Adlam:Ref<RangeTable>
	export const Ahom:Ref<RangeTable>
	export const Anatolian_Hieroglyphs:Ref<RangeTable>
	export const Arabic:Ref<RangeTable>
	export const Armenian:Ref<RangeTable>
	export const Avestan:Ref<RangeTable>
	export const AzeriCase:SpecialCase
	export const Balinese:Ref<RangeTable>
	export const Bamum:Ref<RangeTable>
	export const Bassa_Vah:Ref<RangeTable>
	export const Batak:Ref<RangeTable>
	export const Bengali:Ref<RangeTable>
	export const Bhaiksuki:Ref<RangeTable>
	export const Bidi_Control:Ref<RangeTable>
	export const Bopomofo:Ref<RangeTable>
	export const Brahmi:Ref<RangeTable>
	export const Braille:Ref<RangeTable>
	export const Buginese:Ref<RangeTable>
	export const Buhid:Ref<RangeTable>
	export const C:Ref<RangeTable>
	export const Canadian_Aboriginal:Ref<RangeTable>
	export const Carian:Ref<RangeTable>
	export interface CaseRange extends Struct<CaseRange>{
		lo:uint32
		hi:uint32
	}
	export const CaseRanges:CaseRange[]
	export const Categories:Record<string,Ref<RangeTable>>
	export const Caucasian_Albanian:Ref<RangeTable>
	export const Cc:Ref<RangeTable>
	export const Cf:Ref<RangeTable>
	export const Chakma:Ref<RangeTable>
	export const Cham:Ref<RangeTable>
	export const Cherokee:Ref<RangeTable>
	export const Chorasmian:Ref<RangeTable>
	export const Co:Ref<RangeTable>
	export const Common:Ref<RangeTable>
	export const Coptic:Ref<RangeTable>
	export const Cs:Ref<RangeTable>
	export const Cuneiform:Ref<RangeTable>
	export const Cypriot:Ref<RangeTable>
	export const Cypro_Minoan:Ref<RangeTable>
	export const Cyrillic:Ref<RangeTable>
	export const Dash:Ref<RangeTable>
	export const Deprecated:Ref<RangeTable>
	export const Deseret:Ref<RangeTable>
	export const Devanagari:Ref<RangeTable>
	export const Diacritic:Ref<RangeTable>
	export const Digit:Ref<RangeTable>
	export const Dives_Akuru:Ref<RangeTable>
	export const Dogra:Ref<RangeTable>
	export const Duployan:Ref<RangeTable>
	export const Egyptian_Hieroglyphs:Ref<RangeTable>
	export const Elbasan:Ref<RangeTable>
	export const Elymaic:Ref<RangeTable>
	export const Ethiopic:Ref<RangeTable>
	export const Extender:Ref<RangeTable>
	export const FoldCategory:Record<string,Ref<RangeTable>>
	export const FoldScript:Record<string,Ref<RangeTable>>
	export const Georgian:Ref<RangeTable>
	export const Glagolitic:Ref<RangeTable>
	export const Gothic:Ref<RangeTable>
	export const Grantha:Ref<RangeTable>
	export const GraphicRanges:Ref<RangeTable>[]
	export const Greek:Ref<RangeTable>
	export const Gujarati:Ref<RangeTable>
	export const Gunjala_Gondi:Ref<RangeTable>
	export const Gurmukhi:Ref<RangeTable>
	export const Han:Ref<RangeTable>
	export const Hangul:Ref<RangeTable>
	export const Hanifi_Rohingya:Ref<RangeTable>
	export const Hanunoo:Ref<RangeTable>
	export const Hatran:Ref<RangeTable>
	export const Hebrew:Ref<RangeTable>
	export const Hex_Digit:Ref<RangeTable>
	export const Hiragana:Ref<RangeTable>
	export const Hyphen:Ref<RangeTable>
	export const IDS_Binary_Operator:Ref<RangeTable>
	export const IDS_Trinary_Operator:Ref<RangeTable>
	export const Ideographic:Ref<RangeTable>
	export const Imperial_Aramaic:Ref<RangeTable>
	export function In(r:rune,...ranges:Ref<RangeTable>[]):bool
	export const Inherited:Ref<RangeTable>
	export const Inscriptional_Pahlavi:Ref<RangeTable>
	export const Inscriptional_Parthian:Ref<RangeTable>
	export function is(rangeTab:Ref<RangeTable>,r:rune):bool
	export function isControl(r:rune):bool
	export function isDigit(r:rune):bool
	export function isGraphic(r:rune):bool
	export function isLetter(r:rune):bool
	export function isLower(r:rune):bool
	export function isMark(r:rune):bool
	export function isNumber(r:rune):bool
	export function isOneOf(ranges:Ref<RangeTable>[],r:rune):bool
	export function isPrint(r:rune):bool
	export function isPunct(r:rune):bool
	export function isSpace(r:rune):bool
	export function isSymbol(r:rune):bool
	export function isTitle(r:rune):bool
	export function isUpper(r:rune):bool
	export const Javanese:Ref<RangeTable>
	export const Join_Control:Ref<RangeTable>
	export const Kaithi:Ref<RangeTable>
	export const Kannada:Ref<RangeTable>
	export const Katakana:Ref<RangeTable>
	export const Kawi:Ref<RangeTable>
	export const Kayah_Li:Ref<RangeTable>
	export const Kharoshthi:Ref<RangeTable>
	export const Khitan_Small_Script:Ref<RangeTable>
	export const Khmer:Ref<RangeTable>
	export const Khojki:Ref<RangeTable>
	export const Khudawadi:Ref<RangeTable>
	export const L:Ref<RangeTable>
	export const Lao:Ref<RangeTable>
	export const Latin:Ref<RangeTable>
	export const Lepcha:Ref<RangeTable>
	export const Letter:Ref<RangeTable>
	export const Limbu:Ref<RangeTable>
	export const Linear_A:Ref<RangeTable>
	export const Linear_B:Ref<RangeTable>
	export const Lisu:Ref<RangeTable>
	export const Ll:Ref<RangeTable>
	export const Lm:Ref<RangeTable>
	export const Lo:Ref<RangeTable>
	export const Logical_Order_Exception:Ref<RangeTable>
	export const Lower:Ref<RangeTable>
	//1
	export const LowerCase:int
	export const Lt:Ref<RangeTable>
	export const Lu:Ref<RangeTable>
	export const Lycian:Ref<RangeTable>
	export const Lydian:Ref<RangeTable>
	export const M:Ref<RangeTable>
	export const Mahajani:Ref<RangeTable>
	export const Makasar:Ref<RangeTable>
	export const Malayalam:Ref<RangeTable>
	export const Mandaic:Ref<RangeTable>
	export const Manichaean:Ref<RangeTable>
	export const Marchen:Ref<RangeTable>
	export const Mark:Ref<RangeTable>
	export const Masaram_Gondi:Ref<RangeTable>
	//127
	export const MaxASCII:rune
	//3
	export const MaxCase:int
	//255
	export const MaxLatin1:rune
	//1114111
	export const MaxRune:rune
	export const Mc:Ref<RangeTable>
	export const Me:Ref<RangeTable>
	export const Medefaidrin:Ref<RangeTable>
	export const Meetei_Mayek:Ref<RangeTable>
	export const Mende_Kikakui:Ref<RangeTable>
	export const Meroitic_Cursive:Ref<RangeTable>
	export const Meroitic_Hieroglyphs:Ref<RangeTable>
	export const Miao:Ref<RangeTable>
	export const Mn:Ref<RangeTable>
	export const Modi:Ref<RangeTable>
	export const Mongolian:Ref<RangeTable>
	export const Mro:Ref<RangeTable>
	export const Multani:Ref<RangeTable>
	export const Myanmar:Ref<RangeTable>
	export const N:Ref<RangeTable>
	export const Nabataean:Ref<RangeTable>
	export const Nag_Mundari:Ref<RangeTable>
	export const Nandinagari:Ref<RangeTable>
	export const Nd:Ref<RangeTable>
	export const New_Tai_Lue:Ref<RangeTable>
	export const Newa:Ref<RangeTable>
	export const Nko:Ref<RangeTable>
	export const Nl:Ref<RangeTable>
	export const No:Ref<RangeTable>
	export const Noncharacter_Code_Point:Ref<RangeTable>
	export const Number:Ref<RangeTable>
	export const Nushu:Ref<RangeTable>
	export const Nyiakeng_Puachue_Hmong:Ref<RangeTable>
	export const Ogham:Ref<RangeTable>
	export const Ol_Chiki:Ref<RangeTable>
	export const Old_Hungarian:Ref<RangeTable>
	export const Old_Italic:Ref<RangeTable>
	export const Old_North_Arabian:Ref<RangeTable>
	export const Old_Permic:Ref<RangeTable>
	export const Old_Persian:Ref<RangeTable>
	export const Old_Sogdian:Ref<RangeTable>
	export const Old_South_Arabian:Ref<RangeTable>
	export const Old_Turkic:Ref<RangeTable>
	export const Old_Uyghur:Ref<RangeTable>
	export const Oriya:Ref<RangeTable>
	export const Osage:Ref<RangeTable>
	export const Osmanya:Ref<RangeTable>
	export const Other:Ref<RangeTable>
	export const Other_Alphabetic:Ref<RangeTable>
	export const Other_Default_Ignorable_Code_Point:Ref<RangeTable>
	export const Other_Grapheme_Extend:Ref<RangeTable>
	export const Other_ID_Continue:Ref<RangeTable>
	export const Other_ID_Start:Ref<RangeTable>
	export const Other_Lowercase:Ref<RangeTable>
	export const Other_Math:Ref<RangeTable>
	export const Other_Uppercase:Ref<RangeTable>
	export const P:Ref<RangeTable>
	export const Pahawh_Hmong:Ref<RangeTable>
	export const Palmyrene:Ref<RangeTable>
	export const Pattern_Syntax:Ref<RangeTable>
	export const Pattern_White_Space:Ref<RangeTable>
	export const Pau_Cin_Hau:Ref<RangeTable>
	export const Pc:Ref<RangeTable>
	export const Pd:Ref<RangeTable>
	export const Pe:Ref<RangeTable>
	export const Pf:Ref<RangeTable>
	export const Phags_Pa:Ref<RangeTable>
	export const Phoenician:Ref<RangeTable>
	export const Pi:Ref<RangeTable>
	export const Po:Ref<RangeTable>
	export const Prepended_Concatenation_Mark:Ref<RangeTable>
	export const PrintRanges:Ref<RangeTable>[]
	export const Properties:Record<string,Ref<RangeTable>>
	export const Ps:Ref<RangeTable>
	export const Psalter_Pahlavi:Ref<RangeTable>
	export const Punct:Ref<RangeTable>
	export const Quotation_Mark:Ref<RangeTable>
	export const Radical:Ref<RangeTable>
	export interface Range16 extends Struct<Range16>{
		lo:uint16
		hi:uint16
		stride:uint16
	}
	export interface Range32 extends Struct<Range32>{
		lo:uint32
		hi:uint32
		stride:uint32
	}
	export interface RangeTable extends Struct<RangeTable>{
		r16:Range16[]
		r32:Range32[]
		latinOffset:int
	}
	export const Regional_Indicator:Ref<RangeTable>
	export const Rejang:Ref<RangeTable>
	//65533
	export const ReplacementChar:rune
	export const Runic:Ref<RangeTable>
	export const S:Ref<RangeTable>
	export const STerm:Ref<RangeTable>
	export const Samaritan:Ref<RangeTable>
	export const Saurashtra:Ref<RangeTable>
	export const Sc:Ref<RangeTable>
	export const Scripts:Record<string,Ref<RangeTable>>
	export const Sentence_Terminal:Ref<RangeTable>
	export const Sharada:Ref<RangeTable>
	export const Shavian:Ref<RangeTable>
	export const Siddham:Ref<RangeTable>
	export const SignWriting:Ref<RangeTable>
	export function simpleFold(r:rune):rune
	export const Sinhala:Ref<RangeTable>
	export const Sk:Ref<RangeTable>
	export const Sm:Ref<RangeTable>
	export const So:Ref<RangeTable>
	export const Soft_Dotted:Ref<RangeTable>
	export const Sogdian:Ref<RangeTable>
	export const Sora_Sompeng:Ref<RangeTable>
	export const Soyombo:Ref<RangeTable>
	export const Space:Ref<RangeTable>
	export interface SpecialCase extends Array<CaseRange>{
		toUpper(r:rune):rune
		toTitle(r:rune):rune
		toLower(r:rune):rune
	}
	export const Sundanese:Ref<RangeTable>
	export const Syloti_Nagri:Ref<RangeTable>
	export const Symbol:Ref<RangeTable>
	export const Syriac:Ref<RangeTable>
	export const Tagalog:Ref<RangeTable>
	export const Tagbanwa:Ref<RangeTable>
	export const Tai_Le:Ref<RangeTable>
	export const Tai_Tham:Ref<RangeTable>
	export const Tai_Viet:Ref<RangeTable>
	export const Takri:Ref<RangeTable>
	export const Tamil:Ref<RangeTable>
	export const Tangsa:Ref<RangeTable>
	export const Tangut:Ref<RangeTable>
	export const Telugu:Ref<RangeTable>
	export const Terminal_Punctuation:Ref<RangeTable>
	export const Thaana:Ref<RangeTable>
	export const Thai:Ref<RangeTable>
	export const Tibetan:Ref<RangeTable>
	export const Tifinagh:Ref<RangeTable>
	export const Tirhuta:Ref<RangeTable>
	export const Title:Ref<RangeTable>
	//2
	export const TitleCase:int
	export function to(_case:int,r:rune):rune
	export function toLower(r:rune):rune
	export function toTitle(r:rune):rune
	export function toUpper(r:rune):rune
	export const Toto:Ref<RangeTable>
	export const TurkishCase:SpecialCase
	export const Ugaritic:Ref<RangeTable>
	export const Unified_Ideograph:Ref<RangeTable>
	export const Upper:Ref<RangeTable>
	//0
	export const UpperCase:int
	//1114112
	export const UpperLower:rune
	export const Vai:Ref<RangeTable>
	export const Variation_Selector:Ref<RangeTable>
	//"15.0.0"
	export const Version:string
	export const Vithkuqi:Ref<RangeTable>
	export const Wancho:Ref<RangeTable>
	export const Warang_Citi:Ref<RangeTable>
	export const White_Space:Ref<RangeTable>
	export const Yezidi:Ref<RangeTable>
	export const Yi:Ref<RangeTable>
	export const Z:Ref<RangeTable>
	export const Zanabazar_Square:Ref<RangeTable>
	export const Zl:Ref<RangeTable>
	export const Zp:Ref<RangeTable>
	export const Zs:Ref<RangeTable>

export function emptyCaseRange():CaseRange
export function refCaseRange():Ref<CaseRange>
export function refOfCaseRange(x:CaseRange):Ref<CaseRange>
export function emptyRange16():Range16
export function refRange16():Ref<Range16>
export function refOfRange16(x:Range16):Ref<Range16>
export function emptyRange32():Range32
export function refRange32():Ref<Range32>
export function refOfRange32(x:Range32):Ref<Range32>
export function emptyRangeTable():RangeTable
export function refRangeTable():Ref<RangeTable>
export function refOfRangeTable(x:RangeTable):Ref<RangeTable>
}
