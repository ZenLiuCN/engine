declare module "golang/time" {
	// @ts-ignore
	import * as go from 'go'
	// @ts-ignore
	import type {error,int64,Alias,int,Struct,float64,bool,Ref} from 'go'


	export function parseInLocation(layout,value:string,loc:Ref<Location>):[Time,error]
	export function newTimer(d:Duration):Ref<Timer>
	export const Hour:Duration
	export const Microsecond:Duration
	export const Nanosecond:Duration
	export const Millisecond:Duration
	export const Second:Duration
	export const Minute:Duration
	export function after(d:Duration):go.ChanRecv<Time>
	export function unix(sec:int64,nsec:int64):Time
	export function unixMilli(msec:int64):Time
	export function unixMicro(usec:int64):Time
	export const Monday:Weekday
	export const Sunday:Weekday
	export const Tuesday:Weekday
	export const Wednesday:Weekday
	export const Thursday:Weekday
	export const Friday:Weekday
	export const Saturday:Weekday
	export function loadLocation(name:string):[Ref<Location>,error]
	export function afterFunc(d:Duration,f:()=>void):Ref<Timer>
	export function tick(d:Duration):go.ChanRecv<Time>
	export const February:Month
	export const May:Month
	export const July:Month
	export const August:Month
	export const September:Month
	export const January:Month
	export const March:Month
	export const April:Month
	export const June:Month
	export const October:Month
	export const November:Month
	export const December:Month
	export interface Month extends Alias<int>{
		string():string

	}
	export interface Weekday extends Alias<int>{
		string():string

	}
	export function since(t:Time):Duration
	export function now():Time
	export interface ParseError extends Struct<ParseError>{
		layout:string

		value:string

		layoutElem:string

		valueElem:string

		message:string
		error():string

	}
	export function parse(layout,value:string):[Time,error]
	export function sleep(d:Duration):void
	export interface Duration extends Alias<int64>{
		string():string
		microseconds():int64
		seconds():float64
		hours():float64
		round(m:Duration):Duration
		nanoseconds():int64
		milliseconds():int64
		minutes():float64
		truncate(m:Duration):Duration
		abs():Duration

	}
	export function until(t:Time):Duration
	export interface Location extends Struct<Location>{
		string():string

	}
	export interface Time extends Struct<Time>{
		yearDay():int
		unmarshalBinary(data:Uint8Array):error
		string():string
		clock():[int,int,int]
		marshalText():[Uint8Array,error]
		compare(u:Time):int
		hour():int
		zone():[string,int]
		isoWeek():[int,int]
		unixMilli():int64
		addDate(years:int,months:int,days:int):Time
		location():Ref<Location>
		date():[int,Month,int]
		second():int
		goString():string
		format(layout:string):string
		before(u:Time):bool
		month():Month
		sub(u:Time):Duration
		utc():Time
		unixNano():int64
		appendFormat(b:Uint8Array,layout:string):Uint8Array
		isZero():bool
		nanosecond():int
		in(loc:Ref<Location>):Time
		zoneBounds():[Time,Time]
		unmarshalText(data:Uint8Array):error
		after(u:Time):bool
		year():int
		minute():int
		unix():int64
		unixMicro():int64
		gobDecode(data:Uint8Array):error
		equal(u:Time):bool
		weekday():Weekday
		local():Time
		unmarshalJson(data:Uint8Array):error
		add(d:Duration):Time
		round(d:Duration):Time
		marshalJson():[Uint8Array,error]
		truncate(d:Duration):Time
		day():int
		marshalBinary():[Uint8Array,error]
		gobEncode():[Uint8Array,error]
		isDst():bool

	}
	export function newTicker(d:Duration):Ref<Ticker>
	export function date(year:int,month:Month,day,hour,min,sec,nsec:int,loc:Ref<Location>):Time
	export function fixedZone(name:string,offset:int):Ref<Location>
	export function loadLocationFromTzData(name:string,data:Uint8Array):[Ref<Location>,error]
	export function parseDuration(s:string):[Duration,error]
	export interface Timer extends Struct<Timer>{
		c:go.ChanRecv<Ref<Time>
>

		stop():bool
		reset(d:Duration):bool

	}
	export interface Ticker extends Struct<Ticker>{
		c:go.ChanRecv<Ref<Time>
>

		stop():void
		reset(d:Duration):void

	}
	export const RFC3339=9
	export const ANSIC=1
	export const UnixDate=2
	export const RubyDate=3
	export const RFC1123Z=8
	export const RFC3339Nano=10
	export const Layout=0
	export const RFC822=4
	export const RFC822Z=5
	export const Kitchen=11
	export const StampMicro=14
	export const TimeOnly=18
	export const RFC850=6
	export const RFC1123=7
	export const Stamp=12
	export const StampMilli=13
	export const StampNano=15
	export const DateTime=16
	export const DateOnly=17

}