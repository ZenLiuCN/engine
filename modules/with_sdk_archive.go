//go:build (sdk_archive || sdk || all) && !no_sdk && !no_sdk_archive

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang/archive/tar"
	_ "github.com/ZenLiuCN/engine/modules/golang/archive/zip"
)
