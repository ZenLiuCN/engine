//go:build (sdk_encoding || sdk || all) && !no_sdk && !no_sdk_encoding

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/ascii85"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/asn1"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/base32"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/base64"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/binary"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/csv"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/gob"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/hex"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/json"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/pem"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/xml"
)
