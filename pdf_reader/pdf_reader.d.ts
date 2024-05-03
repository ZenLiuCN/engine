declare module "go/pdf/reader" {
    //@ts-ignore
    import * as io from "go/io"
    //@ts-ignore
    import type {bool, float64, int, int64, Ptr, Struct} from "go"

    interface Stringer {
        string(): string
    }

    export interface Text extends Struct<Text> {
        font: string
        fontSize: float64
        x: float64
        y: float64
        w: float64
        s: string
    }

    export interface Point extends Struct<Point> {
        x: float64
        y: float64
    }

    export interface Rect extends Struct<Rect> {
        min: Point
        max: Point
    }

    export interface Content extends Struct<Content> {
        text: Text[]
        rect: Rect[]
    }

    export interface Column extends Struct<Column> {
        position: int64
        content: TextVertical
    }

    export interface Row extends Struct<Row> {
        position: int64
        content: TextHorizontal
    }

    export interface Outline extends Struct<Outline> {
        title: string
        child: Outline[]
    }

    export interface Font extends Struct<Font> {
        v:Value
    }

    export interface Page extends Struct<Page> {
        readonly v: Value

        resources(): Value

        fonts(): string[]

        font(name: string): Font
    }

    export type Rows = Ptr<Row>[]
    export type Columns = Ptr<Column>[]
    export type TextVertical = Text[]
    export type TextHorizontal = Text[]

    export interface Reader extends Struct<Reader> {
        page(num: int): Page

        outline(num: int): Outline

        numPage(): int

        getPlainText(): [io.Reader, Error]

        trailer(): Value
    }


    export type ValueKind =
    /*Null*/ 0 |
        /*Bool*/ 1 |
        /*Integer*/ 2 |
        /*Real*/ 3 |
        /*String*/ 4 |
        /*Name*/ 5 |
        /*Dict*/ 6 |
        /*Array*/ 7 |
        /*Stream*/ 8


    export interface Value extends Struct<Value>, Stringer {
        isNull(): boolean

        kind(): ValueKind

        bool(): bool

        int64(): int64

        float64(): float64

        rawString(): string

        text(): string

        textFromUTF16(): string

        name(): string

        keys(): string[]

        key(k: string): Value

        index(i: int): Value

        len(): int

        reader(): io.ReadCloser
    }

    export function open(path: string): (Reader | Error)[]

    export function newReader(r: io.ReaderAt, size:/*int64*/number): (Reader | Error)[]

    export function newReaderEncrypted(r: io.ReaderAt, size:/*int64*/number, pwd: () => string): (Reader | Error)[]
}