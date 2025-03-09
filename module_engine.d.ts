declare module "go/engine" {
    // @ts-ignore
    export class Engine {
        constructor()

        //dbg mode will execute script with mapping data
        dbg: boolean

        /**
         * specificity modules
         * @param exclude or includes
         * @param modules names of modules
         */
        constructor(exclude: boolean, ...modules: string[])

        disableModules(...modules: string[]): boolean
        set(name: string, value: any)
        runString(sc: string): [Value, Error]
        runJs(sc: string): [Value, Error]
        runTs(sc: string): [Value, Error]
        runCode(code: Code): [Value, Error]
        runCodeWithMapping(code: Code): [Value, Error]

        //call this when not use this engine
        free()
    }

    export interface Value {
        export(): any
    }

    export interface Code {
    }
}