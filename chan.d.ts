declare module "go/chan" {
    export interface Chan<T> extends ReadOnlyChan<T>, WriteOnlyChan<T> {


    }

    export interface ReadOnlyChan<T> {

        recv(handle: (t: T) => void):Promise<void>

        closed(): boolean

        close()
    }

    export interface WriteOnlyChan<T> {

        send(v: T)

        closed(): boolean

        close()
    }
}
