package sqlx

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestBatch(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {SQLX} from "go/sqlx"
const db=new SQLX("sqlite",":memory:")
console.log(db)
console.log(db.exec("CREATE TABLE users(name VARCHAR, age INTEGER, height FLOAT, awesome BOOLEAN, bday DATE)"))
console.log(db.batch("INSERT INTO users (name,age,height,awesome,bday) VALUES(:name,:age,:height,:awesome,:bday)",[
    { name:'marc', age:99, height:1.91, awesome:true, bday:'1970-01-01'},
    { name:'macgyver', age:70, height:1.85, awesome:true, bday:'1951-01-23'}
]))
console.log(db.query('	SELECT name, age, height, awesome, bday	FROM users	WHERE (name = :name1 OR name = :name2) AND age > :age AND awesome = :awesome',	{
  name1:  "macgyver", name2:"marc", age:30, awesome:true,
}))
db.close()
`))
}
