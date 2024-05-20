// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/time'{

	// @ts-ignore
	import * as go from 'go'
	// @ts-ignore
	import type {Ref,int,int64,float64,error,Struct,GoError,bool} from 'go'
	//"Mon Jan _2 15:04:05 2006"
	export const ANSIC:string
	export function after(d:Duration):go.ChanRecv<Time>

	export function afterFunc(d:Duration,f:()=>void):Ref<Timer>

	export const April:Month
	export const August:Month
	export function date(year:int,month:Month,day:int,hour:int,min:int,sec:int,nsec:int,loc:Ref<Location>):Time

	//"2006-01-02"
	export const DateOnly:string
	//"2006-01-02 15:04:05"
	export const DateTime:string
	export const December:Month
	export interface Duration extends int64{

	string():string
	nanoseconds():int64
	microseconds():int64
	milliseconds():int64
	seconds():float64
	minutes():float64
	hours():float64
	truncate(m:Duration):Duration
	round(m:Duration):Duration
	abs():Duration
	}
	export const February:Month
	export function fixedZone(name:string,offset:int):Ref<Location>

	export const Friday:Weekday
	export const Hour:Duration
	export const January:Month
	export const July:Month
	export const June:Month
	//"3:04PM"
	export const Kitchen:string
	//"01/02 03:04:05PM '06 -0700"
	export const Layout:string
	export function loadLocation(name:string):Ref<Location>

	export function loadLocationFromTZData(name:string,data:Uint8Array):Ref<Location>

	export const Local:Ref<Location>
	export interface Location extends Struct<Location>{

			string():string
	}
	export const March:Month
	export const May:Month
	export const Microsecond:Duration
	export const Millisecond:Duration
	export const Minute:Duration
	export const Monday:Weekday
	export interface Month extends int{

	string():string
	}
	export const Nanosecond:Duration
	export function newTicker(d:Duration):Ref<Ticker>

	export function newTimer(d:Duration):Ref<Timer>

	export const November:Month
	export function now():Time

	export const October:Month
	export function parse(layout:string,value:string):Time

	export function parseDuration(s:string):Duration

	export interface ParseError extends Struct<ParseError>,Error,GoError{

			layout:string
			value:string
			layoutElem:string
			valueElem:string
			message:string
			error():string
	}
	export function parseInLocation(layout:string,value:string,loc:Ref<Location>):Time

	//"Mon, 02 Jan 2006 15:04:05 MST"
	export const RFC1123:string
	//"Mon, 02 Jan 2006 15:04:05 -0700"
	export const RFC1123Z:string
	//"2006-01-02T15:04:05Z07:00"
	export const RFC3339:string
	//"2006-01-02T15:04:05.999999999Z07:00"
	export const RFC3339Nano:string
	//"02 Jan 06 15:04 MST"
	export const RFC822:string
	//"02 Jan 06 15:04 -0700"
	export const RFC822Z:string
	//"Monday, 02-Jan-06 15:04:05 MST"
	export const RFC850:string
	//"Mon Jan 02 15:04:05 -0700 2006"
	export const RubyDate:string
	export const Saturday:Weekday
	export const Second:Duration
	export const September:Month
	export function since(t:Time):Duration

	export function sleep(d:Duration):void

	//"Jan _2 15:04:05"
	export const Stamp:string
	//"Jan _2 15:04:05.000000"
	export const StampMicro:string
	//"Jan _2 15:04:05.000"
	export const StampMilli:string
	//"Jan _2 15:04:05.000000000"
	export const StampNano:string
	export const Sunday:Weekday
	export const Thursday:Weekday
	export function tick(d:Duration):go.ChanRecv<Time>

	export interface Ticker extends Struct<Ticker>{

			C:go.ChanRecv<Time>
			stop():void
			reset(d:Duration):void
	}
	export interface Time extends Struct<Time>{

			string():string
			goString():string
			format(layout:string):string
			appendFormat(b:Uint8Array,layout:string):Uint8Array
			after(u:Time):bool
			before(u:Time):bool
			compare(u:Time):int
			equal(u:Time):bool
			isZero():bool
			date():[int,Month,int]
			year():int
			month():Month
			day():int
			weekday():Weekday
			isoWeek():[int,int]
			clock():[int,int,int]
			hour():int
			minute():int
			second():int
			nanosecond():int
			yearDay():int
			add(d:Duration):Time
			sub(u:Time):Duration
			addDate(years:int,months:int,days:int):Time
			utc():Time
			local():Time
			In(loc:Ref<Location>):Time
			location():Ref<Location>
			zone():[string,int]
			zoneBounds():[Time,Time]
			unix():int64
			unixMilli():int64
			unixMicro():int64
			unixNano():int64
			marshalBinary():Uint8Array
			unmarshalBinary(data:Uint8Array)/*error*/
			gobEncode():Uint8Array
			gobDecode(data:Uint8Array)/*error*/
			marshalJSON():Uint8Array
			unmarshalJSON(data:Uint8Array)/*error*/
			marshalText():Uint8Array
			unmarshalText(data:Uint8Array)/*error*/
			isDST():bool
			truncate(d:Duration):Time
			round(d:Duration):Time
	}
	//"15:04:05"
	export const TimeOnly:string
	export interface Timer extends Struct<Timer>{

			C:go.ChanRecv<Time>
			stop():bool
			reset(d:Duration):bool
	}
	export const Tuesday:Weekday
	export const UTC:Ref<Location>
	export function unix(sec:int64,nsec:int64):Time

	//"Mon Jan _2 15:04:05 MST 2006"
	export const UnixDate:string
	export function unixMicro(usec:int64):Time

	export function unixMilli(msec:int64):Time

	export function until(t:Time):Duration

	export const Wednesday:Weekday
	export interface Weekday extends int{

	string():string
	}
	export function emptyTime():Time
	export function emptyRefTime():Ref<Time>
	export function refOfTime(x:Time,v:Ref<Time>)
	export function unRefTime(v:Ref<Time>):Time
	export function emptyLocation():Location
	export function emptyRefLocation():Ref<Location>
	export function refOfLocation(x:Location,v:Ref<Location>)
	export function unRefLocation(v:Ref<Location>):Location
}