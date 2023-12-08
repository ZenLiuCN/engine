declare module "go/sqlx" {
    // @ts-ignore
    import {Duration} from "go/time"

    export class SQLX {
        constructor(driver: string, dsn: string)

        query(qry: string, args?: Record<string, any>): Array<any>

        exec(qry: string, args?: Record<string, any>): Result

        prepare(qry: string): Stmt

        begin(): TX

        close()

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
        commit()

        rollback()

        query(qry: string, args?: Record<string, any>): Array<any>

        exec(qry: string, args?: Record<string, any>): Result

        prepare(qry: string): Stmt

        stmt(stmt: Stmt): Stmt
    }

    export interface Stmt {
        query(args?: Record<string, any>): Array<any>

        exec(args?: Record<string, any>): Array<any>

        close()
    }

    export interface Result {
        lastInsertedId: number
        rowsAffected: number
    }

}
