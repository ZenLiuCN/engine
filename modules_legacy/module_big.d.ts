declare module 'go/big' {
    export function zeroInt(): Int

    export function oneRat(): Rat

    export function equals(a, b: Int): boolean

    export function bigger(a, b: Int): boolean

    export function smaller(a, b: Int): boolean

    export class Int {
        constructor(n: number | string)

        /**
         * -1 for negative
         *
         * 0 for zero
         *
         * 1 for positive
         */
        sign(): -1 | 0 | 1

        neg(): Int

        add(x, y: Int): Int

        sub(x, y: Int): Int

        mul(x, y: Int): Int

        quo(x, y: Int): Int

        rem(x, y: Int): Int

        div(x, y: Int): Int

        mod(x, y: Int): Int

        divMod(x, y, m: Int): Int

        modSqrt(x, p: Int): Int

        sqrt(x: Int): Int

        exp(x, y, m: Int): Int

        gcd(x, y, a, b: Int): Int

        cmp(y: Int): -1 | 0 | 1

        cmpAbs(y: Int): -1 | 0 | 1

        setString(s: string, base: number): Int | undefined

        bitLen(): number

        int64(): number

        isInt64(): boolean

        quoRem(x, y, r: Int): Int

        text(base: number): string

        string(): string

        bytes(): Uint8Array
    }

    export class Rat {
        /**
         *
         * @param a  integer or float or string
         * @param b  optional denom, must integer
         */
        constructor(a: number | string, b?: number)

        sign(): -1 | 0 | 1

        setFloat64(f: number): Rat

        setFrac(a, b: Int): Rat

        float64(): number

        float32(): number

        isInt(): boolean

        num(): Int

        denom(): Int

        set(v: Rat): Rat

        abs(v: Rat): Rat

        neg(v: Rat): Rat

        add(x, y: Rat): Rat

        sub(x, y: Rat): Rat

        mul(x, y: Rat): Rat

        quo(x, y: Rat): Rat

        rem(x, y: Rat): Rat

        div(x, y: Rat): Rat

        mod(x, y: Rat): Rat

        divMod(x, y, m: Rat): Rat

        modSqrt(x, p: Rat): Rat

        sqrt(x: Rat): Rat

        exp(x, y, m: Rat): Rat

        cmp(y: Rat): -1 | 0 | 1

        cmpAbs(y: Rat): -1 | 0 | 1

        ratString():string
        string():string
        floatString(precision:number):string
    }
}