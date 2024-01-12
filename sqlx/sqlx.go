package sqlx

import (
	"context"
	"database/sql"
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	"strconv"
	"time"

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

func (S SQLXModule) ExportsWithEngine(eng *engine.Engine) map[string]any {
	return map[string]any{
		"SQLX": eng.ToConstructor(func(v []goja.Value) (any, error) {
			bigint := false
			if len(v) == 3 {
				opt := v[2].Export().(map[string]any)
				x := opt["bigint"]
				if x != nil && x.(bool) {
					bigint = true
				}
			}
			db, err := sqlx.Connect(v[0].ToString().String(), v[1].ToString().String())
			if err != nil {
				return nil, err
			}
			return engine.RegisterResource(eng, &SQLx{DB: db, Engine: eng, BigInt: bigint}), nil
		}),
		"bitToBool": func(row []map[string]any, key ...string) ([]map[string]any, error) {
			return mapAll(func(v []byte) (any, error) {
				if len(v) == 1 {
					return v[0] != 0, nil
				} else {
					return v, nil
				}
			}, row, key...)
		},
		"boolToBit": func(row []map[string]any, key ...string) ([]map[string]any, error) {
			return mapAll(func(v bool) (any, error) {
				if v {
					return []byte{1}, nil
				} else {
					return []byte{0}, nil
				}
			}, row, key...)
		},
		"bytesToString": func(row []map[string]any, key ...string) ([]map[string]any, error) {
			return mapAll(func(v []byte) (any, error) {
				return string(v), nil
			}, row, key...)
		},
		"stringToBytes": func(row []map[string]any, key ...string) ([]map[string]any, error) {
			return mapAll(func(v string) (any, error) {
				return []byte(v), nil
			}, row, key...)
		},
		"int64ToString": func(row []map[string]any, key ...string) ([]map[string]any, error) {
			return mapAll(func(v int64) (any, error) {
				return strconv.FormatInt(v, 10), nil
			}, row, key...)
		},
		"stringToInt64": func(row []map[string]any, key ...string) ([]map[string]any, error) {
			return mapAll(func(v string) (any, error) {
				return strconv.ParseInt(v, 0, 64)
			}, row, key...)
		},
		"formatTime": func(row []map[string]any, format string, key ...string) ([]map[string]any, error) {
			return mapAll(func(v time.Time) (any, error) {
				return v.Format(format), nil
			}, row, key...)
		},
		"parseTime": func(row []map[string]any, format string, key ...string) ([]map[string]any, error) {
			return mapAll(func(v string) (any, error) {
				return time.Parse(format, v)
			}, row, key...)
		},
	}
}

func mapAll[T any](fn func(T) (any, error), row []map[string]any, key ...string) ([]map[string]any, error) {
	var err error
	for _, m := range row {
		for _, k := range key {
			v := m[k]
			if val, ok := v.(T); ok {
				m[k], err = fn(val)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return row, nil
}

type SQLx struct {
	*sqlx.DB
	*engine.Engine
	BigInt bool
}

func (s *SQLx) Close() error {
	if s.DB == nil {
		return nil
	}
	s.Engine.RemoveResources(s)
	defer func() {
		s.DB = nil
	}()
	return s.DB.Close()
}

func (s *SQLx) Query(query string, args map[string]any) (goja.Value, error) {
	var r *sqlx.Rows
	var err error
	if args != nil && len(args) > 0 {
		if s.BigInt {
			for k, v := range args {
				if b, ok := v.(*big.Int); ok {
					args[k] = b.Int64()
				}
			}
		}
		r, err = s.DB.NamedQuery(query, args)
		if err != nil {
			return nil, err
		}
	} else {
		r, err = s.DB.Queryx(query)
		if err != nil {
			return nil, err
		}
	}
	defer fn.IgnoreClose(r)
	var g []any
	for r.Next() {
		v := make(map[string]any)
		err = r.MapScan(v)
		if err != nil {
			return nil, err
		}
		if s.BigInt {
			for k, val := range v {
				if t, ok := val.(int64); ok {
					v[k] = big.NewInt(t)
				}
			}
		}
		g = append(g, v)
	}
	return s.Engine.NewArray(g...), nil
}
func (s *SQLx) Exec(query string, args map[string]any) (res Result, err error) {
	var r sql.Result

	if args != nil && len(args) > 0 {
		if s.BigInt {
			for k, v := range args {
				if b, ok := v.(*big.Int); ok {
					args[k] = b.Int64()
				}
			}
		}
		r, err = s.DB.NamedExec(query, args)
		if err != nil {
			return Result{}, err
		}
	} else {
		r, err = s.DB.Exec(query)
		if err != nil {
			return
		}
	}
	res.RowsAffected, err = r.RowsAffected()
	if err != nil {
		err = nil
	}
	res.LastInsertId, err = r.LastInsertId()
	if err != nil {
		err = nil
	}
	return
}

type Result struct {
	LastInsertId int64
	RowsAffected int64
}

func (s *SQLx) Batch(query string, args []map[string]any) (res Result, err error) {
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
	r, err = s.DB.NamedExec(query, args)
	if err != nil {
		return
	}
	res.RowsAffected, err = r.RowsAffected()
	if err != nil {
		err = nil
	}
	res.LastInsertId, err = r.LastInsertId()
	if err != nil {
		err = nil
	}
	return
}
func (s *SQLx) Prepare(query string) (*Stmt, error) {
	r, err := s.DB.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	i := &Stmt{r, s.Engine, s.BigInt}
	return i, nil
}
func (s *SQLx) Begin() (*TX, error) {
	r, err := s.DB.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	i := &TX{Tx: r, Engine: s.Engine}
	return i, nil
}

type TX struct {
	*sqlx.Tx
	*engine.Engine
	BigInt bool
}

func (s *TX) Commit() error {
	return s.Tx.Commit()
}
func (s *TX) Rollback() error {
	return s.Tx.Rollback()
}
func (s *TX) Prepare(qry string) (*Stmt, error) {
	r, err := s.Tx.PrepareNamed(qry)
	if err != nil {
		return nil, err
	}
	i := &Stmt{r, s.Engine, s.BigInt}
	return i, nil
}
func (s *TX) Stmt(stmt *Stmt) *Stmt {
	r := s.Tx.NamedStmt(stmt.NamedStmt)
	i := &Stmt{r, s.Engine, s.BigInt}
	return i
}
func (s *TX) Query(query string, args map[string]any) (goja.Value, error) {
	if s.BigInt {
		for k, v := range args {
			if b, ok := v.(*big.Int); ok {
				args[k] = b.Int64()
			}
		}
	}
	r, err := s.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	var g []any
	for r.Next() {
		v := make(map[string]any)
		err = r.MapScan(v)
		if err != nil {
			return nil, err
		}
		if s.BigInt {
			for k, val := range v {
				if t, ok := val.(int64); ok {
					v[k] = big.NewInt(t)
				}
			}
		}
		g = append(g, v)
	}
	return s.Engine.NewArray(g...), nil
}
func (s *TX) Exec(query string, args map[string]any) (Result, error) {
	if s.BigInt {
		for k, v := range args {
			if b, ok := v.(*big.Int); ok {
				args[k] = b.Int64()
			}
		}
	}
	r, err := s.NamedExec(query, args)
	if err != nil {
		return Result{}, err
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
	}, nil
}

type Stmt struct {
	*sqlx.NamedStmt
	*engine.Engine
	BigInt bool
}

func (s *Stmt) Query(args map[string]any) (goja.Value, error) {
	if s.BigInt {
		for k, v := range args {
			if b, ok := v.(*big.Int); ok {
				args[k] = b.Int64()
			}
		}
	}
	r, err := s.NamedStmt.Queryx(args)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	var g []any
	for r.Next() {
		v := make(map[string]any)
		err = r.MapScan(v)
		if err != nil {
			return nil, err
		}
		if s.BigInt {
			for k, val := range v {
				if t, ok := val.(int64); ok {
					v[k] = big.NewInt(t)
				}
			}
		}
		g = append(g, v)
	}
	return s.Engine.NewArray(g...), nil
}
func (s *Stmt) Exec(args map[string]any) (Result, error) {
	if s.BigInt {
		for k, v := range args {
			if b, ok := v.(*big.Int); ok {
				args[k] = b.Int64()
			}
		}
	}
	r, err := s.NamedStmt.Exec(args)
	if err != nil {
		return Result{}, err
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
	}, nil
}
func (s *Stmt) Close() error {
	s.Engine.RemoveResources(s.NamedStmt)
	return s.NamedStmt.Close()
}
