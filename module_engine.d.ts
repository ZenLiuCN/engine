declare module "go/engine"{
    export class Engine {
        constructor()
        /**
         * specificity modules
         * @param exclude or includes
         * @param modules names of modules
         */
        constructor(exclude: boolean, ...modules: string[])
        disableModules(...modules:string[]):boolean
        runString(sc: string): Value

        runJs(sc: string): Value

        runTs(sc: string): Value

        compile(src: string, ts: boolean): Code

        set(name: string, value: any)

        execute(code: Code): Value

        free()
    }
    export interface Value{
        export():any
    }
    export interface Code {
    }
}