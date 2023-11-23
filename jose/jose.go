package jose

import (
	"github.com/ZenLiuCN/engine"
)

type Jose struct {
}

func (j Jose) Name() string {
	return "jose"
}

func (j Jose) Initialize(engine *engine.Engine) engine.Module {
	panic("not impl")
}
