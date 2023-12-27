declare module "go/sqlx" {
    /**
     * convert bit[1] to boolean value
     * @param rows row data
     * @param keys the property keys
     */
    export function bitToBool(rows:Record<string, any>[],... keys:string[]):Record<string, any>[]
    /**
     * convert boolean to bit[1] value
     * @param rows row data
     * @param keys the property keys
     */
    export function boolToBit(rows:Record<string, any>[],... keys:string[]):Record<string, any>[]

    /**
     * convert binary string to string value
     * @param rows row data
     * @param keys the property keys
     */
    export function bytesToString(rows:Record<string, any>[],... keys:string[]):Record<string, any>[]
    /**
     * convert string to binary string value
     * @param rows row data
     * @param keys the property keys
     */
    export function stringToBytes(rows:Record<string, any>[],... keys:string[]):Record<string, any>[]

    /**
     * convert bigint to string value
     * @param rows row data
     * @param keys the property keys
     */
    export function int64ToString(rows:Record<string, any>[],... keys:string[]):Record<string, any>[]
    /**
     * convert string to int64 value
     * @param rows row data
     * @param keys the property keys
     */
    export function stringToInt64(rows:Record<string, any>[],... keys:string[]):Record<string, any>[]
    /**
     *
     * @param rows the data rows
     * @param layout go time format layout
     * @param keys property key
     */
    export function parseTime(rows:Record<string, any>[],layout:string,... keys:string[]):Record<string, any>[]

    /**
     *
     * @param rows the data rows
     * @param layout go time format layout
     * @param keys property key
     */
    export function formatTime(rows:Record<string, any>[],layout:string,... keys:string[]):Record<string, any>[]
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
