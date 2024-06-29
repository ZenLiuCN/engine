package bbolt

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestSimple(t *testing.T) {
	v := engine.Get()
	defer v.Free()
	fn.Panic1(v.RunTs(
		//language=typescript
		`
		import * as bbolt from 'go.etcd.io/bbolt'
		import {toFileMode} from 'golang/io/fs'
		import {bytesFromString as bfs} from 'go'
		const db=bbolt.open("data.bolt",toFileMode(0o0600) ,null)
		db.update(tx=>{
            let bkt=tx.createBucketIfNotExists(bfs('data'))
            bkt.put(bfs('some'),bfs(JSON.stringify({
            	a:1,
            	b:'ab',
            	c:false
            })))
		})
		db.close()
		`))
}
