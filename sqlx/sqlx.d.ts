declare module "go/sqlx" {
    // @ts-ignore
    import {Duration} from "go/time"

    export class SQLX {
        /**
         * whether support auto convert int64 to big.Int
         */
        readonly BigInt: boolean

        /**
         *
         * @param driver the driver name
         * @param dsn the dsn for driver
         * @param conf with bigint=true  convert 64bit integer to big.Int
         */
        constructor(driver: string, dsn: string, conf?: { bigint: boolean })

        query(qry: string, args?: Record<string, any>): Array<any>

        exec(qry: string, args?: Record<string, any>): Result

        /**
         * do batch insert
         * @param qry a insert query with columns (must with columns) and value (only one line needed)
         * @param args array of object of args
         */
        batch(qry: string, args: Record<string, any>[]): Result

        prepare(qry: string): Stmt

        begin(): TX

        close()

        /** set max idle connections */
        setMaxIdleConns(n: number)

        setMaxOpenConns(n: number)

        setConnMaxIdleTime(n: Duration)

        setConnMaxLifetime(n: Duration)

        stats(): DBStats
    }

    export interface DBStats {
        readonly maxOpenConnections: number
        readonly openConnections: number
        readonly inUse: number
        readonly idle: number
        readonly waitCount: number
        readonly waitDuration: Duration
        readonly maxIdleClosed: number
        readonly maxIdleTimeClosed: number
        readonly maxLifetimeClosed: number
    }

    export interface TX {
        readonly BigInt: boolean

        commit()

        rollback()

        query(qry: string, args?: Record<string, any>): Array<any>

        exec(qry: string, args?: Record<string, any>): Result

        prepare(qry: string): Stmt

        stmt(stmt: Stmt): Stmt
    }

    export interface Stmt {
        readonly BigInt: boolean

        query(args?: Record<string, any>): Array<any>

        exec(args?: Record<string, any>): Array<any>

        close()
    }

    export interface Result {
        lastInsertedId: number
        rowsAffected: number
    }

}
