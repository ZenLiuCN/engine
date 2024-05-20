//go:build (sdk || all) && !no_sdk

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang"
	_ "github.com/ZenLiuCN/engine/modules/golang/archive/tar"
	_ "github.com/ZenLiuCN/engine/modules/golang/archive/zip"
	_ "github.com/ZenLiuCN/engine/modules/golang/bufio"
	_ "github.com/ZenLiuCN/engine/modules/golang/bytes"
	_ "github.com/ZenLiuCN/engine/modules/golang/compress/bzip2"
	_ "github.com/ZenLiuCN/engine/modules/golang/compress/flate"
	_ "github.com/ZenLiuCN/engine/modules/golang/compress/gzip"
	_ "github.com/ZenLiuCN/engine/modules/golang/compress/lzw"
	_ "github.com/ZenLiuCN/engine/modules/golang/compress/zlib"
)
