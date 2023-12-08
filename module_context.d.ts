

declare module 'go/context' {
    export interface Context {

    }

    export interface CancelFunc {
        ()
    }
    export interface ContextWithCancel {
        Context: Context
        Cancel: CancelFunc
    }
    // @ts-ignore
    import {Duration} from 'go/time'

    export function background(): Context

    export function withCancel(c: Context):ContextWithCancel

    export function withTimeout(c: Context, t: Duration): ContextWithCancel

}