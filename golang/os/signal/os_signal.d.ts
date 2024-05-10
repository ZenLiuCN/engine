// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/os/signal'{
	// @ts-ignore
	import * as os from 'golang/os'
	// @ts-ignore
	import * as go from 'go'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import type {bool} from 'go'
	export function ignore(...sig:os.Signal[]):void
	export function ignored(sig:os.Signal):bool
	export function notify(c:go.ChanSend<os.Signal>,...sig:os.Signal[]):void
	export function notifyContext(parent:context.Context,...signals:os.Signal[]):[context.Context,context.CancelFunc]
	export function reset(...sig:os.Signal[]):void
	export function stop(c:go.ChanSend<os.Signal>):void
}
