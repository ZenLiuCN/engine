declare module "go" {
    export interface Stringer {
        string(): string
    }

    export interface GoStringer {
        goString(): string
    }

    export interface GoError extends Stringer {
        error(): string

        same(error: GoError): boolean
    }

    export type Err = GoError | null

    export interface Chan<T> extends ReadOnlyChan<T>, WriteOnlyChan<T> {


    }

    export interface ReadOnlyChan<T> {

        recv(handle: (t: T) => void): Promise<void>

        closed(): boolean

        close()
    }

    export interface WriteOnlyChan<T> {

        send(v: T)

        closed(): boolean

        close()
    }

    export interface Maybe<V> {
        Value?: V
        Error?: GoError
    }
    export function toRaw(v:any):any

    /**
     * extract bigint from object
     * @param v object
     * @param key the key of big integer
     */
    export function bigint(v:Record<string, any>,key:string):string
    /**
     * extract bigint from object
     * @param v object array
     * @param key the key of big integer
     */
    export function bigintOf(v:Record<string, any>[],key:string):string[]
}
