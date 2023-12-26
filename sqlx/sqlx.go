package sqlx

import (
	"context"
	"database/sql"
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/ZenLiuCN/engine/sqlx/mysql_2023-12-22"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"math/big"
	_ "modernc.org/sqlite"
)

var (
	//go:embed sqlx.d.ts
	sqlxDefine []byte
)

func init() {
	engine.RegisterModule(SQLXModule{})
}

type SQLXModule struct {
}

func (S SQLXModule) Identity() string {
	return "go/sqlx"
}

func (S SQLXModule) TypeDefine() []byte {
	return sqlxDefine
}

func (S SQLXModule) Exports() map[string]any {
	return nil
}

func (S SQLXModule) ExportsWithEngine(engine *engine.Engine) map[string]any {
	return map[string]any{
		"SQLX": engine.ToConstructor(func(v []goja.Value) any {
			bigint := false
			if len(v) == 3 {
				opt := v[2].Export().(map[string]any)
				x := opt["bigint"]
				if x != nil && x.(bool) {
					bigint = true
				}
			}
			db := fn.Panic1(sqlx.Connect(v[0].ToString().String(), v[1].ToString().String()))
			return &SQLx{DB: db, Engine: engine, BigInt: bigint}
		}),
		"bitToBool": func(row []map[string]any, key ...string) []map[string]any {
			return mapAll(func(v []byte) any {
				if len(v) == 1 {
					return v[0] != 0
				} else {
					return v
				}
			}, row, key...)
		},
		"boolToBit": func(row []map[string]any, key ...string) []map[string]any {
			return mapAll(func(v bool) any {
				if v {
					return []byte{1}
				} else {
					return []byte{0}
				}
			}, row, key...)
		},
		"bytesToString": func(row []map[string]any, key ...string) []map[string]any {
			return mapAll(func(v []byte) any {
				return string(v)
			}, row, key...)
		},
		"stringToBytes": func(row []map[string]any, key ...string) []map[string]any {
			return mapAll(func(v string) any {
				return []byte(v)
			}, row, key...)
		},
	}
}

func mapAll[T any](fn func(T) any, row []map[string]any, key ...string) []map[string]any {
	for _, m := range row {
		for _, k := range key {
			v := m[k]
			if val, ok := v.(T); ok {
				m[k] = fn(val)
			}
		}
	}
	return row
}

type SQLx struct {
	*sqlx.DB
	*engine.Engine
	BigInt bool
}

func (s *SQLx) Close() {
	fn.Panic(s.DB.Close())
	s.DB = nil
}

func (s *SQLx) Query(query string, args map[string]any) goja.Value {
	var r *sqlx.Rows
	if args != nil && len(args) > 0 {
		if s.BigInt {
			for k, v := range args {
				if b, ok := v.(*big.Int); ok {
					args[k] = b.Int64()
				}
			}
		}
		r = fn.Panic1(s.DB.NamedQuery(query, args))
	} else {
		r = fn.Panic1(s.DB.Queryx(query))
	}
	defer r.Close()
	var g []any
	for r.Next() {
		v := make(map[string]any)
		fn.Panic(r.MapScan(v))
		if s.BigInt {
			for k, val := range v {
				if t, ok := val.(int64); ok {
					v[k] = big.NewInt(t)
				}
			}
		}
		g = append(g, v)
	}
	return s.Engine.NewArray(g...)
}
func (s *SQLx) Exec(query string, args map[string]any) Result {
	var r sql.Result
	if args != nil && len(args) > 0 {
		if s.BigInt {
			for k, v := range args {
				if b, ok := v.(*big.Int); ok {
					args[k] = b.Int64()
				}
			}
		}
		r = fn.Panic1(s.DB.NamedExec(query, args))
	} else {
		r = fn.Panic1(s.DB.Exec(query))
	}
	v, err := r.RowsAffected()
	if err != nil {
		v = 0
	}
	i, err := r.LastInsertId()
	if err != nil {
		i = 0
	}
	return Result{
		LastInsertId: i,
		RowsAffected: v,
	}
}

type Result struct {
	LastInsertId int64
	RowsAffected int64
}

func (s *SQLx) Batch(query string, args []map[string]any) Result {
	var r sql.Result
	if s.BigInt {
		for i, arg := range args {
			for k, v := range arg {
				if b, ok := v.(*big.Int); ok {
					args[i][k] = b.Int64()
				}
			}
		}
	}
	r = fn.Panic1(s.DB.NamedExec(query, args))
	v, err := r.RowsAffected()
	if err != nil {
		v = 0
	}
	i, err := r.LastInsertId()
	if err != nil {
		i = 0
	}
	return Result{
		LastInsertId: i,
		RowsAffected: v,
	}
}
func (s *SQLx) Prepare(query string) *Stmt {
	r := fn.Panic1(s.DB.PrepareNamed(query))
	i := &Stmt{r, s.Engine, s.BigInt}
	return i
}
func (s *SQLx) Begin() *TX {
	r := fn.Panic1(s.DB.BeginTxx(context.Background(), nil))
	i := &TX{Tx: r, Engine: s.Engine}
	return i
}

type TX struct {
	*sqlx.Tx
	*engine.Engine
	BigInt bool
}

func (s *TX) Commit() {
	fn.Panic(s.Tx.Commit())
}
func (s *TX) Rollback() {
	fn.Panic(s.Tx.Rollback())
}
func (s *TX) Prepare(qry string) *Stmt {
	r := fn.Panic1(s.Tx.PrepareNamed(qry))
	i := &Stmt{r, s.Engine, s.BigInt}
	return i
}
func (s *TX) Stmt(stmt *Stmt) *Stmt {
	r := s.Tx.NamedStmt(stmt.NamedStmt)
	i := &Stmt{r, s.Engine, s.BigInt}
	return i
}
func (s *TX) Query(query string, args map[string]any) goja.Value {
	if s.BigInt {
		for k, v := range args {
			if b, ok := v.(*big.Int); ok {
				args[k] = b.Int64()
			}
		}
	}
	r := fn.Panic1(s.NamedQuery(query, args))
	defer r.Close()
	var g []any
	for r.Next() {
		v := make(map[string]any)
		fn.Panic(r.MapScan(v))
		if s.BigInt {
			for k, val := range v {
				if t, ok := val.(int64); ok {
					v[k] = big.NewInt(t)
				}
			}
		}
		g = append(g, v)
	}
	return s.Engine.NewArray(g...)
}
func (s *TX) Exec(query string, args map[string]any) Result {
	if s.BigInt {
		for k, v := range args {
			if b, ok := v.(*big.Int); ok {
				args[k] = b.Int64()
			}
		}
	}
	r := fn.Panic1(s.NamedExec(query, args))
	v, err := r.RowsAffected()
	if err != nil {
		v = 0
	}
	i, err := r.LastInsertId()
	if err != nil {
		i = 0
	}
	return Result{
		LastInsertId: i,
		RowsAffected: v,
	}
}

type Stmt struct {
	*sqlx.NamedStmt
	*engine.Engine
	BigInt bool
}

func (s *Stmt) Query(args map[string]any) goja.Value {
	if s.BigInt {
		for k, v := range args {
			if b, ok := v.(*big.Int); ok {
				args[k] = b.Int64()
			}
		}
	}
	r := fn.Panic1(s.NamedStmt.Queryx(args))
	defer r.Close()
	var g []any
	for r.Next() {
		v := make(map[string]any)
		fn.Panic(r.MapScan(v))
		if s.BigInt {
			for k, val := range v {
				if t, ok := val.(int64); ok {
					v[k] = big.NewInt(t)
				}
			}
		}
		g = append(g, v)
	}
	return s.Engine.NewArray(g...)
}
func (s *Stmt) Exec(args map[string]any) Result {
	if s.BigInt {
		for k, v := range args {
			if b, ok := v.(*big.Int); ok {
				args[k] = b.Int64()
			}
		}
	}
	r := s.NamedStmt.MustExec(args)
	v, err := r.RowsAffected()
	if err != nil {
		v = 0
	}
	i, err := r.LastInsertId()
	if err != nil {
		i = 0
	}
	return Result{
		LastInsertId: i,
		RowsAffected: v,
	}
}
func (s *Stmt) Close() {
	fn.Panic(s.NamedStmt.Close())
}
