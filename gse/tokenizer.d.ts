declare module "go/gse" {
    export const version: string;

    export class Tokenizer {
        constructor()

        readonly load: boolean
        dictSep: string
        dictPath: string
        notLoadHMM: boolean
        alphaNum: boolean
        alpha: boolean
        num: boolean
        loadNoFreq: boolean
        minTokenFreq: number
        textFreq: string
        skipLog: boolean
        moreLog: boolean
        skipPos: boolean
        noStop: boolean
        stopWordMap: { [key: string]: boolean }


        //region dict util
        /**
         * after configuration
         */
        init()

        dictionary(): Dictionary

        toToken(text: string, freq: number, pos?: string): Token

        addToken(text: string, freq: number, pos?: string)

        addTokenForce(text: string, freq: number, pos?: string)

        reAddToken(text: string, freq: number, pos?: string)

        removeToken(text: string)

        empty()

        loadDictMap(m: { [key: string]: string }[])

        loadDict(...files: string[])


        getCurrentFilePath(): string

        getIdfPath(...files: string[]): string[]

        read(file: string)

        size(size: number, text, freqText: string): number

        //readN
        //reader
        calcToken()

        //endregion
        //region dict_1.16
        loadDictEmbed(dict?: EmbedDict | string)

        loadStopEmbed(dict?: EmbedDict | string)

        loadStopStr(dict: string)

        loadDictStr(dict: string)

        //endregion
        //region segmenter
        segment(bytes: Uint8Array): Segment[]

        modeSegment(bytes: Uint8Array, mode?: boolean): Segment[]

        splitTextToWords(text: Uint8Array): Uint8Array[]

        //endregion
        //region gse
        cut(text: string, hmm?: boolean): string[]

        cutSearch(text: string, hmm?: boolean): string[]

        cutAll(v: string): string[]

        cutDAG(text: string, ...reg: RegExp[]): string[]

        cutDAGNoHMM(text: string): string[]

        cutStr(text: string, sep?: string): string[]

        loadModel(...prob: { [key: string]: number }[])

        hmmCut(str: string, ...reg: RegExp[]): string[]

        hmmCutMode(str: string, ...prob: { [key: string]: number }[]): string[]

        slice(s: string, searchMode?: boolean): string[]

        string(s: string, searchMode?: boolean): string

        pos(s: string, searchMode?: boolean): SegPos[]

        posStr(s: SegPos[], separator?: string): string

        //endregion
        //region dag
        find(s: string): Found

        value(text: string): Value

        analyze(text: string[], t1: string, by?: boolean): AnalyzeToken[]

        suggestFreq(...str: string[]): number

        //endregion
        //region stop
        loadStopArr(dict: string[])

        loadStop(...files: string[])

        addStop(text: string)

        addStopArr(...text: string[])

        removeStop(text: string)

        isStop(text: string)

        emptyStop()

        //endregion
        //region trim
        trimPunct(s: string[]): string[]

        trimPosPunct(s: SegPos[]): SegPos[]

        trimWithPos(s: SegPos[], ...pos: string[]): SegPos[]

        stop(s: string[]): string[]

        trim(s: string[]): string[]

        trimSymbol(s: string[]): string[]

        trimPos(s: SegPos[]): SegPos[]

        cutStop(s: string, hmm?: boolean): string[]

        cutTrim(s: string, hmm?: boolean): string[]

        posTrim(s: string, search: boolean, ...pos: string[]): SegPos[]

        posTrimArr(s: string, search: boolean, ...pos: string[]): string[]

        posTrimStr(s: string, search: boolean, ...pos: string[]): string

        cutTrimHtml(s: string, hmm?: boolean): string[]

        cutTrimHtmls(s: string, hmm?: boolean): string

        cutUrl(s: string, hmm?: boolean): string[]

        cutUrls(s: string, hmm?: boolean): string

        //endregion
    }

    export interface Found {
        readonly pos: string
        readonly freq: number
        readonly exists: boolean
    }

    export interface SegPos {
        text: string
        pos: string
    }

    export interface Value {
        readonly value: number
        readonly id: number
    }

    export interface Segment {
        start(): number

        end(): number

        token(): Token
    }

    export interface AnalyzeToken {
        start: number
        end: number

        position: number
        len: number

        type: string

        text: string
        freq: number
        pos: string
    }

    export interface Token {
        text(): string

        freq(): number

        pos(): string

        segments(): Segment[]

        equals(str: string): boolean
    }

    export interface Dictionary {
        maxTokenLen(): number

        numTokens(): number

        totalFreq(): number

        addToken(tk: Token)

        removeToken(tk: Token)

        lookupTokens(text: Uint8Array[], tk: Token[]): number
    }

    //region dict_1.16
    export type EmbedDict = 'ja' | 'zh' | 'zh_t' | 'zh_s' | 'zhIdf'

    export function newEmbed(...dict: string[]): Tokenizer

    //endregion
    //region dict_util
    /**
     * switch alphabet case, then return current status
     */
    export function toLower(): boolean;

    export function dictPaths(dir, file: string): string[]

    export function isJp(segText: string): boolean

    //endregion
    //region dictionary
    export function newDict(): Dictionary

    //endregion
    //region trim
    export function splitNum(text: string): string[]

    export function splitNums(text: string): string

    export function filterEmoji(text: string): string

    export function filterSymbol(text: string): string

    export function filterHtml(text: string): string

    export function filterLang(text: string): string

    export function range(text: string): string[]

    export function rangeText(text: string): string

    //endregion
    //region stop
    export function setStopWords(map: { [key: string]: boolean });

    export function getStopWords(): { [key: string]: boolean }

    //endregion
    //region segmenter
    export function splitWords(text: Uint8Array): Uint8Array[]

    //endregion
    //region SegUtils
    export function toString(seg: Segment[], searchMod?: boolean): string

    export function toSlice(seg: Segment[], searchMod?: boolean): string[]

    export function toPos(seg: Segment[], searchMod?: boolean): SegPos[]

    export function join(seg: Uint8Array[]): string

    //endregion
    export interface WithTokenizer {
        withGse(seg: Tokenizer)
    }

    export class TextRanker implements WithTokenizer {
        withGse(seg: Tokenizer);

        loadDict(...file: string[])

        textRank(text: string, topK: number): TagSegment[]

        textRankWithPOS(text: string, topK: number, allowPOS: string[]): TagSegment[]
    }

    export class TagExtractor implements WithTokenizer {
        constructor()

        withGse(seg: Tokenizer);


        loadDict(...file: string[])

        loadIdf(...file: string[])

        loadIdfStr(str: string)

        loadStopWords(...file: string[])

        extractTags(text: string, topK: number): TagSegment[]
    }

    export interface TagSegment {
        readonly text: string

        readonly  weight: number

        string(): string
    }

    export class PosTokenizer implements WithTokenizer {
        constructor()

        withGse(seg: Tokenizer);


        loadDict(...file: string[])

        cut(text: string, hmm?: boolean): SegPos[]

        trimPunct(ps: SegPos[]): SegPos[]

        trim(ps: SegPos[]): SegPos[]

        trimWithPose(ps: SegPos[], ...pos: string[]): SegPos[]


    }
}
