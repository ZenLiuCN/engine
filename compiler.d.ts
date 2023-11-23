declare interface Compiler {
    compileJs(src: string): string

    compileTs(src: string): string
}

declare const compiler: Compiler

