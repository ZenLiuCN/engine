package sqlx

import (
	"context"
	_ "embed"
	"github.com/ZenLiuCN/engine"
	_ "github.com/ZenLiuCN/engine/sqlx/mysql"
	"github.com/ZenLiuCN/fn"
	"github.com/dop251/goja"
	_ "github.com/glebarez/go-sqlite"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var (
	//go:embed sqlx.d.ts
	d []byte
)

func init() {
	engine.Register(&SQLx{})
}

type SQLx struct {
	*sqlx.DB
	*engine.Engine
}

func (s *SQLx) TypeDefine() []byte {
	return d
}

func (s *SQLx) Name() string {
	return "SQLX"
}

func (s *SQLx) Register(engine *engine.Engine) {
	fn.Panic(engine.Runtime.Set("SQLX", func(call goja.ConstructorCall) *goja.Object {
		db := fn.Panic1(sqlx.Connect(call.Argument(0).ToString().String(), call.Argument(1).ToString().String()))
		i := &SQLx{DB: db, Engine: engine}
		o := engine.ToValue(i).(*goja.Object)
		fn.Panic(o.SetPrototype(call.This.Prototype()))
		return o
	}))
}

func (s *SQLx) Close() {
	fn.Panic(s.DB.Close())
	s.DB = nil
}

func (s *SQLx) Query(query string, args map[string]any) goja.Value {
	r := fn.Panic1(s.DB.NamedQuery(query, args))
	defer r.Close()
	var g []any
	for r.Next() {
		v := make(map[string]any)
		fn.Panic(r.MapScan(v))
		g = append(g, v)
	}
	return s.Engine.NewArray(g...)
}
func (s *SQLx) Exec(query string, args map[string]any) map[string]int64 {
	r := fn.Panic1(s.DB.NamedExec(query, args))
	v, err := r.RowsAffected()
	if err != nil {
		v = 0
	}
	i, err := r.LastInsertId()
	if err != nil {
		i = 0
	}
	return map[string]int64{
		"lastInsertId": i,
		"rowsAffected": v,
	}
}
func (s *SQLx) Prepare(query string) *Stmt {
	r := fn.Panic1(s.DB.PrepareNamed(query))
	i := &Stmt{r, s.Engine}
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
}

func (s *TX) Commit() {
	fn.Panic(s.Tx.Commit())
}
func (s *TX) Rollback() {
	fn.Panic(s.Tx.Rollback())
}
func (s *TX) Prepare(qry string) *Stmt {
	r := fn.Panic1(s.Tx.PrepareNamed(qry))
	i := &Stmt{r, s.Engine}
	return i
}
func (s *TX) Stmt(stmt *Stmt) *Stmt {
	r := s.Tx.NamedStmt(stmt.NamedStmt)
	i := &Stmt{r, s.Engine}
	return i
}
func (s *TX) Query(query string, args map[string]any) goja.Value {
	r := fn.Panic1(s.NamedQuery(query, args))
	defer r.Close()
	var g []any
	for r.Next() {
		v := make(map[string]any)
		fn.Panic(r.MapScan(v))
		g = append(g, v)
	}
	return s.Engine.NewArray(g...)
}
func (s *TX) Exec(query string, args map[string]any) map[string]int64 {
	r := fn.Panic1(s.NamedExec(query, args))
	v, err := r.RowsAffected()
	if err != nil {
		v = 0
	}
	i, err := r.LastInsertId()
	if err != nil {
		i = 0
	}
	return map[string]int64{
		"lastInsertId": i,
		"rowsAffected": v,
	}
}

type Stmt struct {
	*sqlx.NamedStmt
	*engine.Engine
}

func (s *Stmt) Query(args map[string]any) goja.Value {
	r := fn.Panic1(s.NamedStmt.Queryx(args))
	defer r.Close()
	var g []any
	for r.Next() {
		v := make(map[string]any)
		fn.Panic(r.MapScan(v))
		g = append(g, v)
	}
	return s.Engine.NewArray(g...)
}
func (s *Stmt) Exec(args map[string]any) map[string]int64 {
	r := s.NamedStmt.MustExec(args)
	v, err := r.RowsAffected()
	if err != nil {
		v = 0
	}
	i, err := r.LastInsertId()
	if err != nil {
		i = 0
	}
	return map[string]int64{
		"lastInsertId": i,
		"rowsAffected": v,
	}
}
func (s *Stmt) Close() {
	fn.Panic(s.NamedStmt.Close())
}
