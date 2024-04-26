package fsnotifier

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestNotifier(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	fn.Panic1(vm.RunJs(
		//language=javascript
		`
import {Watcher} from 'go/fsNotify'
const n=new Watcher()
n.addTree("./.idea")
n.onEvent((ev,err)=>{
    if (err){
        console.error(err)
        n.close()
    }else if (ev){
        console.log(
           {
            path:ev.path(),
			isFile:ev.isFile(),
			isCreated:ev.isCreated(),
			isRemoved:ev.isRemoved(),
			isModified:ev.isModified(),
			isRenamed:ev.isRenamed(),
			isAttributeChanged:ev.isAttributeChanged()
           }
        )
    }
})
n.await()
`))
}
