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
func TestConnectError(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {SQLX} from "go/sqlx"

let now = new Date();
now.setDate(now.getDate() - 1);
const oneDayAgo = now.toISOString().split('T')[0];
const offSetTo = oneDayAgo+" 16:00:00";
now.setDate(now.getDate() - 1);
const twoDaysAgo = now.toISOString().split('T')[0];
const offSetFrom = twoDaysAgo+" 16:00:00";
let prd = null;
let agency_pg = null;

try{
    agency_pg = new SQLX('pgx', 'postgres://medtree:medtree2345678@192.168.8.94:65433/dw?search_path=test');
    let v=agency_pg.query("SELECT * FROM SOME_TABLE LIMIT 1")
    prd = new SQLX('mysql', 'root:123456@tcp(127.0.0.1:330)/test1?parseTime=true');
}catch (error){
    console.log("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx");
    console.log("sync error: ",error);

}finally {
    if (prd != null){
        prd.close();
    }
    if (agency_pg != null){
        agency_pg.close();
    }
}
`))
}
func TestBigInt(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {SQLX} from "go/sqlx"
const db=new SQLX('mysql', 'root:12345678@tcp(192.168.8.94:3326)/shop_medicine?parseTime=true',{bigintText:true,bigintFields:['id']})
const data=db.query('SELECT * FROM shop_medicine_rec_pharma limit 5')
data.forEach(v=>console.log(typeof v.id))
console.table(data)
console.log(db.bigInt())
console.log(db.bigIntText())
console.log(db.bigIntFields())
`))
}
