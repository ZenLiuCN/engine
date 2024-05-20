// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/os/user'{

	// @ts-ignore
	import type {Ref,error,Struct,GoError,int} from 'go'
	export function current():Ref<User>

	export interface Group extends Struct<Group>{

			gid:string
			name:string
	}
	export function lookup(username:string):Ref<User>

	export function lookupGroup(name:string):Ref<Group>

	export function lookupGroupId(gid:string):Ref<Group>

	export function lookupId(uid:string):Ref<User>

	export interface UnknownGroupError extends string,GoError{

	error():string
	}
	export interface UnknownGroupIdError extends GoError,string{

	error():string
	}
	export interface UnknownUserError extends string,GoError{

	error():string
	}
	export interface UnknownUserIdError extends GoError,int{

	error():string
	}
	export interface User extends Struct<User>{

			uid:string
			gid:string
			username:string
			name:string
			homeDir:string
			groupIds():string[]
	}
	export function emptyGroup():Group
	export function emptyRefGroup():Ref<Group>
	export function refOfGroup(x:Group,v:Ref<Group>)
	export function unRefGroup(v:Ref<Group>):Group
	export function emptyUser():User
	export function emptyRefUser():Ref<User>
	export function refOfUser(x:User,v:Ref<User>)
	export function unRefUser(v:Ref<User>):User
}