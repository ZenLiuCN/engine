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

    /**
     * convert int64 inside object to string, fail if any property is not an int64
     * @param v object
     * @param keys the keys of big integer
     */
    export function intToString(v:Record<string, any>,keys:string[]):Record<string, any>

    /**
     * convert string inside object to int64, fail if any property is not an int64
     * @param v object
     * @param keys the keys of big integer
     */
    export function intFromString(v:Record<string, any>,keys:string[]):Record<string, any>
    /**
     * convert bigint in objects to string, fail if any property is not an int64
     * @param v object array
     * @param keys the keys of big integer properties
     */
    export function intToStringArray(v:Record<string, any>[],keys:string[]):Record<string, any>[]
    /**
     * convert bigint in objects to string, fail if any property is not an int64
     * @param v object array
     * @param keys the keys of big integer properties
     */
    export function intFromStringArray(v:Record<string, any>[],keys:string[]):Record<string, any>[]
}
