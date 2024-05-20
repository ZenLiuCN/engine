//go:build (sdk_container || sdk || all) && !no_sdk && !no_sdk_container

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang/container/heap"
	_ "github.com/ZenLiuCN/engine/modules/golang/container/list"
	_ "github.com/ZenLiuCN/engine/modules/golang/container/ring"
)
