//go:build (lark_base || all) && !no_lark_base

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/base-sdk-go/v3"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/base-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/base-sdk-go/v3/service/base/v1"
	_ "github.com/ZenLiuCN/engine/modules/github/com/larksuite/base-sdk-go/v3/service/drive/v1"
)
