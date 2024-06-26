// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/context'{

	// @ts-ignore
	import * as time from 'golang/time'
	// @ts-ignore
	import * as go from 'go'
	// @ts-ignore
	import type {Alias,GoError,Nothing,bool,error,Ref} from 'go'
	export function afterFunc(ctx:Context,f:()=>void):()=>bool

	export function background():Context

	export interface CancelCauseFunc extends Alias<(cause:error)=>void>{

	}
	export interface CancelFunc extends Alias<()=>void>{

	}
	export const Canceled:GoError
	export function cause(c:Context)/*error*/

	export interface Context{

			deadline():[time.Time,bool]
			done():go.ChanRecv<Nothing>
			err()/*error*/
			value(key:any):any
	}
	export const DeadlineExceeded:GoError
	export function todo():Context

	export function withCancel(parent:Context):[Context,CancelFunc]

	export function withCancelCause(parent:Context):[Context,CancelCauseFunc]

	export function withDeadline(parent:Context,d:time.Time):[Context,CancelFunc]

	export function withDeadlineCause(parent:Context,d:time.Time,cause:error):[Context,CancelFunc]

	export function withTimeout(parent:Context,timeout:time.Duration):[Context,CancelFunc]

	export function withTimeoutCause(parent:Context,timeout:time.Duration,cause:error):[Context,CancelFunc]

	export function withValue(parent:Context,key:any,val:any):Context

	export function withoutCancel(parent:Context):Context

}