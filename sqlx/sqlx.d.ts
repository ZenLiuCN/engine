declare class SQLX {
    constructor(driver: string, dsn: string)

    query(qry: string, args?: Record<string, any>): Array<any>

    exec(qry: string, args?: Record<string, any>): Result

    prepare(qry: string): Stmt

    begin(): TX

    close()
}

declare interface TX {
    commit()

    rollback()

    query(qry: string, args?: Record<string, any>): Array<any>

    exec(qry: string, args?: Record<string, any>): Result

    prepare(qry: string): Stmt

    stmt(stmt: Stmt): Stmt
}

declare interface Stmt {
    query(args?: Record<string, any>): Array<any>

    exec(args?: Record<string, any>): Array<any>

    close()
}

declare interface Result {
    lastInsertedId: number
    rowsAffected: number
}

