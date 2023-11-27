declare interface ExecOption extends ProcOption {
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

declare interface ProcOption {
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


}

declare interface SubProc {
    /**
     * start and not wait process exit
     * @returns error or empty string
     */
    start(): string

    /**
     * wait started process exit
     * @returns error or empty string
     */
    wait(): string

    /**
     * release process resources, when not use wait()
     * free() also called
     * @returns error or empty string
     */
    release(): string

    /**
     * run process and wait for exit
     * @returns error or empty string
     */
    run(): string

    /**
     * run and fetch process output
     * @returns  console encoding binary data and error string
     */
    output(): { data: Uint8Array, error: string }

    /**
     * run and fetch process output include stdError
     * @returns  console encoding binary data and error string
     */
    combinedOutput(): { data: Uint8Array, error: string }

    /**
     * use with start|run method to read from stdout
     * ,call once before invoke  start|run
     * @returns  console encoding binary data and error string
     */
    readStdout(): { data: Uint8Array, error: string }

    /**
     * use with start|run method to read from stderr
     * ,call once before invoke  start|run
     * @returns  console encoding binary data and error string
     */
    readStderr(): { data: Uint8Array, error: string }

    /**
     * use with start|run method to write to stdin
     * ,call once with null before invoke  start|run
     * @returns  written bytes count or error string
     */
    writeStdin(data: Uint8Array | null): { write: number, error: string }

    /**
     * false if not stated yet
     */
    exited(): boolean

    /**
     * false if not stated yet
     */
    success(): boolean

    /**
     * -1 if not stated yet
     */
    sysTime(): number

    /**
     * -1 if not stated yet
     */
    userTime(): number

    /**
     * -1 if not stated yet
     */
    sysNanoTime(): number

    /**
     * -1 if not stated yet
     */
    userNanoTime(): number

    /**
     * kill the process
     * @returns error or empty string
     */
    kill(): string

    /**
     * free all used pipes for stdout stderr and stdin
     */
    free()

    /**
     * convert local console encoding buf data to utf8 encode data
     * @param buf
     */
    fromConsole(buf: Uint8Array): Uint8Array

    /**
     * convert utf8 buf data to local console encoding
     * @param buf
     */
    toConsole(buf: Uint8Array): Uint8Array
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

    proc(option: ProcOption): SubProc

    mkdir(path: string)

    mkdirAll(path: string)

    exists(path: string): boolean

    write(path: string, data: Uint8Array)

    writeText(path: string, data: string)

    read(path: string): Uint8Array

    readText(path: string): string

    chdir(path: string)

    pwd(): string

    ls(path?: string): Array<{ dir: boolean, name: string, mode: string, size: number, modified: string }>
}

declare const os: Os

