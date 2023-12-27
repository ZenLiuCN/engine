package duckdb

import "C"
import (
	"context"
	"database/sql"
	"database/sql/driver"
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/engine/sqlx"
	"github.com/ZenLiuCN/fn"
	"github.com/ZenLiuCN/go-duckdb"
	_ "github.com/ZenLiuCN/go-duckdb"
	"github.com/dop251/goja"
	sqlx2 "github.com/jmoiron/sqlx"
)

var (
	//go:embed duckdb.d.ts
	duckdbDefine []byte
)

func init() {
	engine.RegisterModule(DuckDBModule{})
}

type DuckDBModule struct {
}

func (d DuckDBModule) TypeDefine() []byte {
	return duckdbDefine
}

func (d DuckDBModule) Identity() string {
	return "go/duckdb"
}

func (d DuckDBModule) Exports() map[string]any {
	return nil
}

func (d DuckDBModule) ExportsWithEngine(eng *engine.Engine) map[string]any {
	ci := eng.ToValue(eng.ToConstructor(func(v []goja.Value) any {
		dsn := v[0].Export().(string)
		if len(v) == 1 {
			return &Conn{
				c: fn.Panic1(duckdb.NewConnector(dsn, nil)),
				e: eng,
			}
		} else {
			var q []string
			fn.Panic(eng.ExportTo(v[1], &q))
			return &Conn{
				c: fn.Panic1(duckdb.NewConnector(dsn, func(execer driver.ExecerContext) error {
					for _, qry := range q {
						_, err := execer.ExecContext(context.Background(), qry, nil)
						if err != nil {
							return err
						}
					}
					return nil
				})),
				e: eng,
			}
		}
	}))
	return map[string]any{
		"Connector": ci,
	}
}

type Conn struct {
	c driver.Connector
	e *engine.Engine
}

func (c *Conn) ToSQLX() *sqlx.SQLx {
	return &sqlx.SQLx{
		DB:     sqlx2.NewDb(sql.OpenDB(c.c), "duckdb"),
		Engine: c.e,
		BigInt: false,
	}
}

func (c *Conn) ToAppender(schema, table string) *Appender {
	conn := fn.Panic1(c.c.Connect(context.Background()))
	appender := fn.Panic1(duckdb.NewAppenderFromConn(conn, schema, table))
	return &Appender{
		a: appender,
		e: c.e,
		c: conn,
	}
}

type Appender struct {
	a *duckdb.Appender
	e *engine.Engine
	c driver.Conn
}

func (a *Appender) Error() *engine.GoError {
	return engine.GoErrorOf(a.a.Error())
}

func (a *Appender) Flush() *engine.GoError {
	return engine.GoErrorOf(a.a.Flush())
}

func (a *Appender) Close() *engine.GoError {
	defer fn.IgnoreClose(a.c)
	return engine.GoErrorOf(a.a.Close())
}

// AppendRow loads a row of values into the appender. The values are provided as separate arguments.
func (a *Appender) AppendRow(args ...driver.Value) *engine.GoError {
	return engine.GoErrorOf(a.a.AppendRow(args...))
}
