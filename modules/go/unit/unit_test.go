package unit

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestChmod(t *testing.T) {
	vm := engine.Get()
	defer vm.Free()
	fn.Panic1(vm.RunTs(
		//language=typescript
		`
import * as unit from 'go/unit'
unit.chmod('D:/dev/project/workspace/authorized_workspace/engine/modules/go/unit/unit.go',666)
console.table(unit.ls('D:/dev/project/workspace/authorized_workspace/engine/modules/go/'))
`))
}
