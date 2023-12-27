declare module 'go/duckdb' {
    // @ts-ignore
    import {SQLX} from 'go/sqlx'
    // @ts-ignore
    import {GoError} from 'go'

    export class Connector {
        constructor(dsn: string, ...bootQueries: string[])

        toSQLX(): SQLX

        toAppender(schema, table: string): Appender
    }

    export interface Appender {
        close(): GoError

        flush(): GoError

        appendRow(...values: any[]): GoError
    }
}