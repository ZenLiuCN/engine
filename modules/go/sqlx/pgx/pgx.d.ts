declare module 'go/pgx' {
    export interface Config {
        textNumeric?: boolean
        rfc3339Time?: boolean
        textBigInt?: boolean
        textJson?: boolean
    }

    export interface Configure {
        textNumeric: boolean
        rfc3339Time: boolean
        textBigInt: boolean
        textJson: boolean

        Convert(a: Record<string, any>): Record<string, any>

        Parse(a: Record<string, any>, ...key: string[]): Record<string, any>
    }
    export class Pool{

    }
    export class Connection {
        constructor(dsn: string, conf?: Config);

        /**
         * mutable configure
         */
        configure(): Configure

        isClosed(): boolean

        isBusy(): boolean

        /**
         * use @ as parameter prefix
         */
        query(qry: string, args?: Record<string, any>): Rows

        /**
         * use @ as parameter prefix
         */
        exec(qry: string, args?: Record<string, any>): CommandTag

        prepare(qry: string, name?: string): Statement

        /**
         * close and cancel all execution
         */
        close()
    }

    export interface Rows {
        readonly closed: boolean

        parse(): Array<Record<string, any>>

        close()
    }

    export interface Statement {
        readonly  name: string
        readonly  sql: string
        readonly  paramOIDs: number[]
        readonly  fields: FieldDescription[]
    }

    export interface FieldDescription {
        readonly name: string
        readonly tableOID: number //uint32
        readonly tableAttributeNumber: number //uint16
        readonly dataTypeOID: number //uint32
        readonly dataTypeSize: number //int16
        readonly typeModifier: number //int32
        readonly format: number //int16
    }

    export interface Notification {
        pid: number
        channel: string
        payload: string
    }

    export interface CommandTag {
        rowsAffected(): number

        insert(): boolean

        update(): boolean

        delete(): boolean

        select(): boolean

        string(): string
    }
}
