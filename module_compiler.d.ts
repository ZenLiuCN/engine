declare module "go/compiler" {
    //@ts-ignore
    import {Code} from "go/engine"


    export function compileJs(src: string, entry: boolean): [Code, Error]
    export function compileJsCode(src: string, entry: boolean): [Code, Error]
    export function compileTs(src: string, entry: boolean): [Code, Error]
    export function compileTsCode(src: string, entry: boolean): [Code, Error]
}


