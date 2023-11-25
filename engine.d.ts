declare class Engine {
    constructor()
    /**
     * specificity modules
     * @param exclude or includes
     * @param modules names of modules
     */
    constructor(exclude: boolean, ...modules: string[])

    runScript(sc: string): any

    runJavaScript(sc: string): any

    runTypeScript(sc: string): any

    compile(src: string, ts: boolean): Code

    set(name: string, value: any)

    execute(code: Code): any

    runPromise(src:string):Promise<any>
    runPromiseCode(src:Code):Promise<any>
    runPromiseJavaScript(src:string):Promise<any>
    runPromiseTypeScript(src:string):Promise<any>

    free()
}

declare interface Code {
}