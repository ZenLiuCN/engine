declare module "go/time" {
    // @ts-ignore
    import {Stringer} from "go";
    export const UTC: Location

    export interface Location extends Stringer {

    }

    export function fixedZone(name: string, offset: number): Location

    export function loadLocation(name: string): Location | null

    export function date(year: number, month: Month, day, hour, min, sec, nsec: number, loc: Location): Time

    export interface Time extends Stringer {
        after(u: Time): boolean

        before(u: Time): boolean

        compare(u: Time): number

        equal(u: Time): boolean

        isZero(): boolean

        year(): number

        month(): Month

        day(): number

        weekday(): Weekday

        hour(): number

        minute(): number

        second(): number

        nanosecond(): number

        yearDay(): number

        addDate(years: number, months: number, days: number): Time

        utc(): Time

        local(): Time

        unix(): number

        unixMilli(): number

        unixMicro(): number

        unixNano(): number

        isDST(): boolean

        truncate(d: Duration): Time

        round(d: Duration): Time

        format(layout: string): string
    }

    export const January: Month
    export const February: Month
    export const March: Month
    export const April: Month
    export const May: Month
    export const June: Month
    export const July: Month
    export const August: Month
    export const September: Month
    export const October: Month
    export const November: Month
    export const December: Month

    export function monthNumber(m: Month): 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12

    export function numberMonth(m: 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12): Month

    export interface Month extends Stringer {

    }

    export interface Weekday extends Stringer {

    }

    export const Sunday: Weekday
    export const Monday: Weekday
    export const Tuesday: Weekday
    export const Wednesday: Weekday
    export const Thursday: Weekday
    export const Friday: Weekday
    export const Saturday: Weekday

    export function weekdayNumber(m: Month): 0 | 1 | 2 | 3 | 4 | 5 | 6

    export function numberWeekday(m: 0 | 1 | 2 | 3 | 4 | 5 | 6): Weekday

    export const Nanosecond: Duration
    export const Microsecond: Duration
    export const Millisecond: Duration
    export const Second: Duration
    export const Minute: Duration
    export const Hour: Duration

    export interface Duration extends Stringer {
        truncate(m: Duration): Duration

        round(m: Duration): Duration

        abs(): Duration

        seconds(): number

        minutes(): number

        hours(): number

        nanoseconds(): number

        microseconds(): number

        milliseconds(): number
    }

    export function unix(sec, nsec: number): Time

    export function unixMilli(msec: number): Time

    export function unixMicro(usec: number): Time

    export function now(): Time

    /**
     * Layout:
     *
     *  Year: "2006" "06"
     *
     *  Month: "Jan" "January" "01" "1"
     *
     *  Day of the week: "Mon" "Monday"
     *
     *  Day of the month: "2" "_2" "02"
     *
     *  Day of the year: "__2" "002"
     *
     *  Hour: "15" "3" "03" (PM or AM)
     *
     *  Minute: "4" "04"
     *
     *  Second: "5" "05"
     *
     *  AM/PM mark: "PM"
     *
     *  "-0700"     ±hhmm
     *
     * "-07:00"    ±hh:mm
     *
     * "-07"       ±hh
     *
     * "-070000"   ±hhmmss
     *
     * "-07:00:00" ±hh:mm:ss
     *
     * "Z0700"      Z or ±hhmm
     *
     * "Z07:00"     Z or ±hh:mm
     *
     * "Z07"        Z or ±hh
     *
     * "Z070000"    Z or ±hhmmss
     *
     * "Z07:00:00"  Z or ±hh:mm:ss
     *
     * "15:04:05,000" or "15:04:05.000" formats or parses with millisecond precision.
     *
     * @param layout the golang spec layout
     * @param time time text
     */
    export function parse(layout: string, time: string): Time

    /**
     * number with units  "ns", "us" (or "µs"), "ms", "s", "m", "h". eg: 300ms, -1.5h, 2h45m
     * @param value
     */
    export function duration(value: string): Duration

    export function times(value: Duration, n: number): Duration

    export function since(t: Time): Duration

    export function until(t: Time): Duration
    export function sleep(t: Duration)
}