//go:build (sdk_text || sdk || all) && !no_sdk && !no_sdk_text

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang/text/scanner"
	_ "github.com/ZenLiuCN/engine/modules/golang/text/tabwriter"
	_ "github.com/ZenLiuCN/engine/modules/golang/text/template"
	_ "github.com/ZenLiuCN/engine/modules/golang/text/template/parse"
)
