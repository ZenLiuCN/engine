declare module "go/compiler" {
    //@ts-ignore
    import {Code, Mapping} from "go/engine"


    export function compileJs(src: string, entry: boolean): [Code, Error]

    export function compileJsWithMapping(src: string, entry: boolean): [string, Mapping,Uint8Array, Error]

    export function compileJsCode(src: string, entry: boolean): [Code, Error]

    export function compileTs(src: string, entry: boolean): [Code, Error]

    export function compileTsWithMapping(src: string, entry: boolean): [string, Mapping,Uint8Array, Error]

    export function compileTsCode(src: string, entry: boolean): [Code, Error]
}


