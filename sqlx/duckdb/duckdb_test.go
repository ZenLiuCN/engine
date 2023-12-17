package duckdb

import (
	"github.com/ZenLiuCN/engine"
	_ "github.com/ZenLiuCN/engine/sqlx"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestSimple(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {SQLX} from "go/sqlx"
const db=new SQLX("duckdb","?access_mode=READ_WRITE")
console.log(db)
console.log(db.query("SELECT current_setting('access_mode')"))
console.log(db.exec("CREATE TABLE users(name VARCHAR, age INTEGER, height FLOAT, awesome BOOLEAN, bday DATE)"))
console.log(db.exec("INSERT INTO users VALUES('marc', 99, 1.91, true, '1970-01-01')"))
console.log(db.exec("INSERT INTO users VALUES('macgyver', 70, 1.85, true, '1951-01-23')"))
console.log(db.query('	SELECT name, age, height, awesome, bday	FROM users	WHERE (name = :name1 OR name = :name2) AND age > :age AND awesome = :awesome',	{
  name1:  "macgyver", name2:"marc", age:30, awesome:true,
}))
db.close()
`))
}
func TestPgFWD(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {SQLX} from "go/sqlx"
const db=new SQLX("duckdb","?access_mode=READ_WRITE")
console.log(db)
console.log(db.query("SELECT current_setting('access_mode')"))
console.log(db.exec("INSTALL postgres from 'http://extensions.duckdb.org'"))  //FROM 'http://extensions.duckdb.org/v0.9.2/windows_amd64/postgres_scanner.duckdb_extension.gz'
console.log("install done")
console.log(db.exec("ATTACH 'dbname=profit user=postgres password=123456' AS pg (TYPE postgres)"))
console.log("attach done")
console.log(db.exec("USE pg"))
console.log("use done")
console.log(db.query("select * from pg.fact_profit limit 1"))
db.close()
`))
}
