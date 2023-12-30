package pgx

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"strconv"
	"time"
)

var (
	//go:embed pgx.d.ts
	pgxDefine []byte
)

func init() {
	engine.RegisterModule(PgxModule{})
}

type PgxModule struct {
}

func (d PgxModule) TypeDefine() []byte {
	return pgxDefine
}

func (d PgxModule) Identity() string {
	return "go/pgx"
}

func (d PgxModule) Exports() map[string]any {
	return nil
}

func (d PgxModule) ExportsWithEngine(eng *engine.Engine) map[string]any {
	ci := eng.ToValue(eng.ToConstructor(func(v []goja.Value) any {
		dsn := v[0].Export().(string)
		ctx, cc := context.WithCancel(context.Background())
		cf := new(Config)
		if len(v) > 1 {
			fn.Panic(eng.ExportTo(v[1], cf))
		}
		return &Conn{
			cf: cf,
			c:  fn.Panic1(pgx.Connect(ctx, dsn)),
			e:  eng,
			cc: cc,
		}
	}))
	return map[string]any{
		"Connection": ci,
	}
}

var (
	emc = &Config{}
)

type PoolConfig struct {
	*Config
}
type ConnPool struct {
	cf *PoolConfig
	c  *pgxpool.Pool
	e  *engine.Engine
}

func (c *ConnPool) Configure() *Config {
	return c.cf.Config
}

func (c *ConnPool) Close() error {
	c.e.RemoveResources(c)
	c.c.Close()
	return nil
}
func (c *ConnPool) Acquire() *PoolConn {
	ctx, cc := context.WithCancel(context.Background())
	return engine.RegisterResource(c.e, &PoolConn{
		cf: c.Configure(),
		c:  fn.Panic1(c.c.Acquire(ctx)),
		e:  c.e,
		cc: cc,
	})
}
func (c *ConnPool) Query(qry string, args map[string]any) *Rows {
	if args != nil {
		return engine.RegisterResource(c.e, &Rows{c: c.Configure(), r: fn.Panic1(c.c.Query(context.Background(), qry, pgx.NamedArgs(args)))})
	}

	return engine.RegisterResource(c.e, &Rows{c: c.Configure(), r: fn.Panic1(c.c.Query(context.Background(), qry))})
}
func (c *ConnPool) Exec(qry string, args map[string]any) pgconn.CommandTag {
	if args != nil {
		return fn.Panic1(c.c.Exec(context.Background(), qry, pgx.NamedArgs(args)))
	}
	return fn.Panic1(c.c.Exec(context.Background(), qry))
}

type PoolConn struct {
	cf *Config
	c  *pgxpool.Conn
	e  *engine.Engine
	cc context.CancelFunc
}

func (c *PoolConn) Configure() *Config {
	return c.cf
}
func (c *PoolConn) Close() error {
	if c.cc != nil {
		c.cc()
	}
	c.e.RemoveResources(c)
	c.c.Release()
	return nil
}
func (c *PoolConn) Query(qry string, args map[string]any) *Rows {
	if args != nil {
		return engine.RegisterResource(c.e, &Rows{c: c.cf, r: fn.Panic1(c.c.Query(context.Background(), qry, pgx.NamedArgs(args)))})
	}
	return engine.RegisterResource(c.e, &Rows{c: c.cf, r: fn.Panic1(c.c.Query(context.Background(), qry))})
}
func (c *PoolConn) Exec(qry string, args map[string]any) pgconn.CommandTag {
	if args != nil {
		return fn.Panic1(c.c.Exec(context.Background(), qry, pgx.NamedArgs(args)))
	}
	return fn.Panic1(c.c.Exec(context.Background(), qry))
}
func (c *PoolConn) Prepare(qry, name string) *pgconn.StatementDescription {
	return fn.Panic1(c.c.Conn().Prepare(context.Background(), name, qry))
}

type Config struct {
	TextNumeric bool
	RFC3339Time bool
	TextBigInt  bool
	TextJson    bool
}

func (c *Config) Convert(v map[string]any) map[string]any {
	if c == nil || (c == emc) {
		return v
	}
	for s, a := range v {
		if c.TextNumeric {
			if x, ok := a.(pgtype.Numeric); ok {
				v[s] = string(fn.Panic1(x.MarshalJSON()))
				continue
			}
		}
		if c.TextBigInt {
			if x, ok := a.(int64); ok {
				v[s] = strconv.FormatInt(x, 10)
				continue
			} else if x, ok := a.(uint64); ok {
				v[s] = strconv.FormatUint(x, 10)
				continue
			}
		}
		if c.RFC3339Time {
			if x, ok := a.(time.Time); ok {
				v[s] = x.Format(time.RFC3339)
				continue
			}
		}
		if c.TextJson {
			if x, ok := a.(map[string]any); ok {
				b := fn.Panic1(json.Marshal(x))
				v[s] = string(b)
				continue
			}
		}
	}
	return v
}
func (c *Config) Parse(v map[string]any, key ...string) map[string]any {
	if c == nil || (!c.TextNumeric && !c.TextBigInt && !c.RFC3339Time) {
		return v
	}
	for _, s := range key {
		a := v[s]
		if a == nil {
			continue
		}
		if x, ok := a.(string); ok {
			if x == "" {
				continue
			}
			var err error
			if c.TextNumeric {
				n := pgtype.Numeric{}
				err = (&n).Scan(x)
				if err == nil {
					v[s] = n
					continue
				}
			}
			if c.TextBigInt {
				var n int64
				n, err = strconv.ParseInt(x, 0, 64)
				if err == nil {
					v[s] = n
					continue
				}
			}
			if c.RFC3339Time {
				var n time.Time
				n, err = time.Parse(time.RFC3339, x)
				if err == nil {
					v[s] = n
					continue
				}
			}
			if c.TextJson {
				if x[0] == '{' {
					var n map[string]any
					err = json.Unmarshal([]byte(x), &n)
					if err != nil {
						v[s] = n
						continue
					}
				} else if x[0] == '[' {
					var n []any
					err = json.Unmarshal([]byte(x), &n)
					if err != nil {
						v[s] = n
						continue
					}
				}
			}
		}

	}
	return v
}

type Conn struct {
	cf *Config
	c  *pgx.Conn
	e  *engine.Engine
	cc context.CancelFunc
}

func (c *Conn) Configure() *Config {
	return c.cf
}
func (c *Conn) IsClosed() bool {
	return c.c.IsClosed()
}
func (c *Conn) IsBusy() bool {
	return c.c.PgConn().IsBusy()
}
func (c *Conn) Close() error {
	if c.cc != nil {
		c.cc()
	}
	c.e.RemoveResources(c)
	fn.Panic(c.c.Close(context.Background()))
	return nil
}
func (c *Conn) Query(qry string, args map[string]any) *Rows {
	if args != nil {
		return &Rows{c: c.cf, r: fn.Panic1(c.c.Query(context.Background(), qry, pgx.NamedArgs(args)))}
	}
	return &Rows{c: c.cf, r: fn.Panic1(c.c.Query(context.Background(), qry))}
}
func (c *Conn) Exec(qry string, args map[string]any) pgconn.CommandTag {
	if args != nil {
		return fn.Panic1(c.c.Exec(context.Background(), qry, pgx.NamedArgs(args)))
	}
	return fn.Panic1(c.c.Exec(context.Background(), qry))
}
func (c *Conn) Prepare(qry, name string) *pgconn.StatementDescription {
	return fn.Panic1(c.c.Prepare(context.Background(), name, qry))
}
func parseRows(r pgx.Rows, c *Config) (o []any) {
	defer r.Close()
	var cols []string
	for _, description := range r.FieldDescriptions() {
		cols = append(cols, description.Name)
	}
	pnt := false
	for r.Next() {
		val := fn.Panic1(r.Values())
		m := make(map[string]any)
		for i, col := range cols {
			m[col] = val[i]
		}
		if c != nil {
			if !pnt {
				for s, a := range m {
					fmt.Printf("%s %T \n", s, a)
				}
				pnt = true
			}
			m = c.Convert(m)
		}
		o = append(o, m)
	}
	return o
}

type Rows struct {
	c      *Config
	r      pgx.Rows
	Closed bool
}

func (c *Rows) Close() error {
	if c.Closed {
		return nil
	}
	defer func() {
		c.Closed = true
	}()
	c.r.Close()
	return nil
}
func (c *Rows) Parse() []any {
	defer func() {
		c.Closed = true
	}()
	return parseRows(c.r, c.c)
}
