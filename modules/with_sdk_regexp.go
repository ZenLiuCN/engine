//go:build (sdk_regexp || sdk || all) && !no_sdk && !no_sdk_regexp

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang/regexp"
	_ "github.com/ZenLiuCN/engine/modules/golang/regexp/syntax"
)
