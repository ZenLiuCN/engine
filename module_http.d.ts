declare module 'go/http' {
    // @ts-ignore
    import {Time} from 'go/time'
    // @ts-ignore
    import {Reader} from 'go/io'

    export interface Values {
        [key: string]: string[]
        // @ts-ignore
        get(key: string): string
        // @ts-ignore
        set(key, value: string)
        // @ts-ignore
        add(key, value: string)
        // @ts-ignore
        del(key: string)
        // @ts-ignore
        has(key: string): boolean
    }

    export interface Header {
        [key: string]: string[]

        // @ts-ignore
        add(key, value: string)
        // @ts-ignore
        set(key, value: string)
        // @ts-ignore
        get(key: string): string
        // @ts-ignore
        values(key: string): string[]
        // @ts-ignore
        del(key: string)
        // @ts-ignore
        clone(): Header
    }

    export interface Cookie {
        name: string
        value: string

        path?: string    // optional
        domain?: string    // optional
        expires?: Time // optional
        rawExpires?: string    // for reading cookies only

        // MaxAge=0 means no 'Max-Age' attribute specified.
        // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
        // MaxAge>0 means Max-Age attribute present and given in seconds
        maxAge?: number

        secure: boolean

        httpOnly: boolean
        /**
         * 1 default 2 Lax 3 Strict 4 None
         */
        sameSite?: 1 | 2 | 3 | 4
        raw?: string
        unparsed?: string []// Raw text of unparsed attribute-value pairs
    }

    export interface Request {
        readonly   method: string
        readonly url: Url

        header?: Header
        form?: Values
        postForm?: Values

        userAgent(): string

        referer(): string

        setBasicAuth(username, password: string)

        formValue(key: string): string

        postFormValue(key: string): string

        cookie(name: string): Cookie

        cookies(): Cookie[]

        addCookie(c: Cookie)
    }

    export interface Url {

    }

    export interface Response {
        readonly status: string // e.g. "200 OK"
        readonly statusCode: number    // e.g. 200
        readonly proto: string // e.g. "HTTP/1.0"
        readonly protoMajor: number    // e.g. 1
        readonly protoMinor: number    // e.g. 0

        readonly header: Header
        readonly request: Request

        /** close response*/
        close()

        hasError(): boolean

        getError(): Error | undefined


        hasBody(): boolean


        json(): any

        text(): string

        binary(): Uint8Array
    }

    export function get(url: string): Response

    export function post(url, contentType: string, body: Reader | Uint8Array | string): Response

    export function postJson(url, body: Uint8Array | Record<string, any> | Array<any> | string): Response

    export function values(): Values

    export function valuesToHeader(v: Values): Header

    export function headerToValues(v: Header): Values

    export function request(req: Request): Response

    type METHOD = 'OPTIONS' | 'GET' | 'HEAD' | 'POST' | 'PUT' | 'DELETE' | 'TRACE' | 'CONNECT' | 'PATCH'

    export function requestOf(method: METHOD, url: string, body?: Reader | Uint8Array | Record<string, any> | Array<any> | string): Request


}
