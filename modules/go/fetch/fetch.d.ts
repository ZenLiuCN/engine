declare module "go/fetch"{
    export interface FetchOptions {
        method: string
        /**
         * no effect
         */
        mode?: string
        /**
         * no effect
         */
        cache?: string
        /**
         * with cookie jar or not
         */
        credentials?: 'same-origin' | 'include' | 'omit'
        headers?: { [key: string]: string }
        redirect?: 'follow' | 'error'
        /**
         * no effect
         */
        referrerPolicy?: string

        body?: string | ArrayBuffer | Record<any, any>
    }

    export interface FetchResponse {
        readonly status: number
        readonly statusText: string
        readonly type: string
        readonly url: string
        readonly headers: { [key: string]: string }
        readonly ok: boolean
        readonly bodyUsed: boolean

        arrayBuffer(): Promise<ArrayBuffer>

        json(): Promise<any>

        text(): Promise<string>
    }

    export function fetch(url: string, opt?: FetchOptions): Promise<FetchResponse>
}
