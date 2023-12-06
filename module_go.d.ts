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
}
