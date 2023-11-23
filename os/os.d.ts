declare interface ExecOption {
    /**
     * command, can be empty when use with sleep to just do a sleep.
     */
    cmd: string
    /**
     * arguments
     */
    args?: string[]
    /**
     * customer working directory
     */
    workingDir?: string
    /**
     * show window (only support windows)
     */
    showWindow?: boolean
    /**
     * expand paths in args
     */
    pathPath?: boolean
    /**
     * sleep after execute, when with use await, sleep after execute finished.
     */
    sleep?: string
    /**
     * optional execute
     */
    optional?: boolean
    /**
     * await execute finish
     */
    await?: boolean

}

declare interface Os {
    /**
     * root directory of executable
     */
    readonly root: string
    /**
     * extension of executable
     */
    readonly ext: string
    /**
     * executable name with extension
     */
    readonly executable: string
    /**
     * name of executable without extension
     */
    readonly name: string
    readonly pathSeparator: string
    readonly pathListSeparator: string


    /**
     * Expand a relative path to absolute path.
     * first '@' expand to root
     * '$NAME or %NAME%' expand as environment variable.
     * @param path path value
     */
    expand(path: string): string

    /**
     * prepend value to environment variable
     * @param name variable name
     * @param values values
     */
    pre(name: string, ...values: string[])

    /**
     * append value to environment variable
     * @param name variable name
     * @param values values
     */
    ap(name: string, ...values: string[])

    /**
     * prepend path to environment variable
     * @see os.expand
     * @param name variable name
     * @param paths values
     */
    prep(name: string, ...paths: string[])

    /**
     * append path to environment variable
     * @see os.expand
     * @param name variable name
     * @param paths values
     */
    app(name: string, ...paths: string[])

    set(name: string, ...values: string[])

    /**
     * set path to environment variable, override exists one
     * @see os.expand
     * @param name variable name
     * @param paths values
     */
    setp(name: string, ...paths: string[])

    /**
     * put value to environment variable, not override exists one
     * @param name variable name
     * @param values
     */
    put(name: string, ...values: string[])

    /**
     * put path to environment variable, not override exists one
     * @see os.expand
     * @param name variable name
     * @param paths values
     */
    putp(name: string, ...paths: string[])

    /**
     * get environment variable
     * @param name the name
     * @returns empty string if not exists one
     */
    var(name: string): string

    /**
     * eval script file
     * @see os.expand
     * @param path of typescript file or javascript file
     */
    evalFile(path: string): any

    /**
     * eval script files
     * @see os.expand
     * @param paths of typescript files or javascript files
     */
    evalFiles(...paths: string[]): any[]

    exec(option: ExecOption)

    mkdir(path: string)

    mkdirAll(path: string)

    exists(path: string): boolean

    write(path: string, data: ArrayBuffer)

    writeText(path: string, data: string)

    read(path: string): ArrayBuffer

    readText(path: string): string
    chdir(path: string)
    pwd():string
    ls(path?:string):Array<{dir:boolean,name:string,mode:string,size:number,modified:Date}>
}

declare const os: Os

