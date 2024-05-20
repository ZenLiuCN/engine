//go:build (sdk_sync || sdk || all) && !no_sdk && !no_sdk_sync

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang/sync"
	_ "github.com/ZenLiuCN/engine/modules/golang/sync/atomic"
)
