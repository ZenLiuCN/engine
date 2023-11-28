declare module "go/compiler" {
    //@ts-ignore
    import {Code} from "go/engine"

    export function compileJs(src: string): string

    export function compileJsCode(src: string): Code

    export function compileTs(src: string): string

    export function compileTsCode(src: string): Code
}


