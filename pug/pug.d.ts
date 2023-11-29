declare module 'go/pug' {
    // @ts-ignore
    import {Writer} from "go/io"

    /**
     * the http/template
     */
    export interface Template {
        templates(): Template[]

        option(...opt: string[]): Template

        parse(code: string): Template

        execute(w: Writer, value: any)
    }
    export interface ReplaceTokens {
        golangMode ?:boolean
        tagBgn     ?:string
        tagEnd     ?:string
        tagVoid    ?:string
        tagArgEsc  ?:string
        tagArgUne  ?:string
        tagArgStr  ?:string
        tagArgAdd  ?:string
        tagArgBgn  ?:string
        tagArgEnd  ?:string

        condIf?:     string
        condUnless?: string
        condCase?:   string
        condWhile?:  string
        condFor?:    string
        condEnd?:    string
        condForIf?:  string

        codeForElse   ?:string
        codeLongcode  ?:string
        codeBuffered  ?:string
        codeUnescaped ?:string
        codeElse      ?:string
        codeElseIf    ?:string
        codeCaseWhen  ?:string
        codeCaseDef   ?:string
        codeMixBlock  ?:string

        textStr     ?:string
        textComment ?:string

        mixinBgn?:         string
        mixinEnd?:         string
        mixinVarBgn?:      string
        mixinVar?:         string
        mixinVarRest?:     string
        mixinVarEnd?:      string
        mixinVarBlockBgn?: string
        mixinVarBlock?:    string
        mixinVarBlockEnd?: string
    }
    export function config(path: ReplaceTokens)
    export function parse(name: string, code: Uint8Array): string

    export function parseText(name: string, code: string): string

    export function parseFile(path: string): string


    export function template(name, html: string): Template
}