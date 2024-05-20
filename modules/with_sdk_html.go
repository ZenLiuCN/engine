//go:build (sdk_html || sdk || all) && !no_sdk && !no_sdk_html

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang/html"
	_ "github.com/ZenLiuCN/engine/modules/golang/html/template"
)
