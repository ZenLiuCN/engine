

declare module "go/unit" {


    /**
     * root directory of executable
     */
    export const root: string
    /**
     * extension of executable
     */
    export const ext: string
    /**
     * executable name with extension
     */
    export const executable: string
    /**
     * name of executable without extension
     */
    export const name: string
    export const pathSeparator: string
    export const pathListSeparator: string


    /**
     * Expand a relative path to absolute path.
     * first '@' expand to root
     * '$NAME or %NAME%' expand as environment variable.
     * @param path path value
     */
    export function expand(path: string): string

    /**
     * prepend value to environment variable
     * @param name variable name
     * @param values values
     */
    export function pre(name: string, ...values: string[])

    /**
     * append value to environment variable
     * @param name variable name
     * @param values values
     */
    export function ap(name: string, ...values: string[])

    /**
     * prepend path to environment variable
     * @see os.expand
     * @param name variable name
     * @param paths values
     */
    export function prep(name: string, ...paths: string[])

    /**
     * append path to environment variable
     * @see os.expand
     * @param name variable name
     * @param paths values
     */
    export function app(name: string, ...paths: string[])

    export function set(name: string, ...values: string[])

    /**
     * set path to environment variable, override exists one
     * @see os.expand
     * @param name variable name
     * @param paths values
     */
    export function setp(name: string, ...paths: string[])

    /**
     * put value to environment variable, not override exists one
     * @param name variable name
     * @param values
     */
    export function put(name: string, ...values: string[])

    /**
     * put path to environment variable, not override exists one
     * @see os.expand
     * @param name variable name
     * @param paths values
     */
    export function putp(name: string, ...paths: string[])

    /**
     * get environment variable
     * @param name the name
     * @returns empty string if not exists one
     */
    export function variable(name: string): string

    /**
     * eval script file
     * @see os.expand
     * @param path of typescript file or javascript file
     */
    export function evalFile(path: string): any

    /**
     * eval script files
     * @see os.expand
     * @param paths of typescript files or javascript files
     */
    export function evalFiles(...paths: string[]): any[]



    export function mkdir(path: string)

    export function mkdirAll(path: string)

    export function exists(path: string): boolean

    export function write(path: string, data: Uint8Array)

    export function writeText(path: string, data: string)

    export function read(path: string): Uint8Array

    export function readText(path: string): string

    export function chdir(path: string)

    export function pwd(): string
    export function mkdirAll(path: string, perm: number)//511=0777

    export function rename(path, newPath: string)
    export function userCacheDir(): string

    export function userConfigDir(): string

    export function userHomeDir(): string
    export function chown(path: string, uid, gid: number)
    export function chmod(path: string, flag: number)
    export function getUID(): number

    export function getGID(): number

    export function getPagesize(): number

    export function getPid(): number

    export function remove(name: string)

    export function removeAll(name: string)

    export function hostname(): string

    export function tempDir(): string
    export function ls(path?: string): Array<{
        name: string,
        mode: string,
        size: number,
        modified: string }>

    export function stat(path: string): undefined | {
        name: string
        mode: string
        size: number
        modified: string /*golang/time.Time*/
    }

    export interface ExecOption extends ProcOption {
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

    export interface ProcOption {
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

    export interface SubProc {
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

    export function exec(option: ExecOption)

    export function proc(option: ProcOption): SubProc

    /**
     * eval script file
     * @see os.expand
     * @param path of typescript file or javascript file
     */
    export function evalFile(path: string): any

    /**
     * eval script files
     * @see os.expand
     * @param paths of typescript files or javascript files
     */
    export function evalFiles(...paths: string[]): any[]
}
