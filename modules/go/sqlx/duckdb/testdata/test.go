package main

import (
	"github.com/ZenLiuCN/engine"
	_ "github.com/ZenLiuCN/engine/modules/go/sqlx"
	_ "github.com/ZenLiuCN/engine/modules/go/sqlx/duckdb"
	"github.com/ZenLiuCN/fn"
)

func main() {
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
