package pgx

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestSimple(t *testing.T) {
	v := engine.Get()
	defer v.Free()
	fn.Panic1(v.RunJs(
		//language=javascript
		`
	import {Connection} from 'go/pgx'
	const conn=new Connection('postgres://medtree:medtree2345678@192.168.8.94:65433/profits?',{textNumeric:true,rfc3339Time:true,textBigInt:true,textJson:true})
	const v=conn.query('select * from _family_user_202309260920 limit 5').parse()
	console.table(v)
	conn.close()
`))
}
