declare module 'go/duckdb' {
    // @ts-ignore
    import {SQLX} from 'go/sqlx'
    // @ts-ignore
    import {error} from 'go'

    export class Connector {
        constructor(dsn: string, ...bootQueries: string[])

        toSQLX(): SQLX

        toAppender(schema, table: string): Appender
    }

    export interface Appender {
        close():void /*error*/

        flush():void /*error*/

        appendRow(...values: any[]):void /*error*/
    }
}
