declare module "go/sqlx"{
    export class SQLX {
        constructor(driver: string, dsn: string)

        query(qry: string, args?: Record<string, any>): Array<any>

        exec(qry: string, args?: Record<string, any>): Result

        prepare(qry: string): Stmt

        begin(): TX

        close()
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
