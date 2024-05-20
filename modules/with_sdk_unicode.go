//go:build (sdk_unicode || sdk || all) && !no_sdk && !no_sdk_unicode

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang/unicode"
	_ "github.com/ZenLiuCN/engine/modules/golang/unicode/utf16"
	_ "github.com/ZenLiuCN/engine/modules/golang/unicode/utf8"
)
