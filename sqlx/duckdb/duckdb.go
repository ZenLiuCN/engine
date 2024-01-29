package duckdb

import "C"
import (
	"context"
	"database/sql"
	"database/sql/driver"
	_ "embed"
	"fmt"
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
	ci := eng.ToValue(eng.ToConstructor(func(v []goja.Value) (rx any, err error) {
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				case error:
					err = v
				default:
					err = fmt.Errorf("%s", v)
				}
				rx = nil
			}
		}()
		dsn := v[0].Export().(string)
		if len(v) == 1 {
			rx = &Conn{
				c: fn.Panic1(duckdb.NewConnector(dsn, nil)),
				e: eng,
			}
			return
		} else {
			var q []string
			fn.Panic(eng.ExportTo(v[1], &q))
			rx = &Conn{
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
			return
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
	return engine.RegisterResource(c.e, &sqlx.SQLx{
		DB:              sqlx2.NewDb(sql.OpenDB(c.c), "duckdb"),
		Engine:          c.e,
		BigIntProcessor: &sqlx.BigIntProcessor{},
	})
}

func (c *Conn) ToAppender(schema, table string) (*Appender, error) {
	conn, err := c.c.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	app, err := duckdb.NewAppenderFromConn(conn, schema, table)
	if err != nil {
		return nil, err
	}
	appender := engine.RegisterResource(c.e, app)
	return &Appender{
		a: appender,
		e: c.e,
		c: conn,
	}, nil
}

type Appender struct {
	a *duckdb.Appender
	e *engine.Engine
	c driver.Conn
}

func (a *Appender) Error() error {
	return a.a.Error()
}

func (a *Appender) Flush() error {
	return a.a.Flush()
}

func (a *Appender) Close() error {
	a.e.RemoveResources(a.a)
	defer fn.IgnoreClose(a.c)
	return a.a.Close()
}

// AppendRow loads a row of values into the appender. The values are provided as separate arguments.
func (a *Appender) AppendRow(args ...driver.Value) error {
	return a.a.AppendRow(args...)
}
