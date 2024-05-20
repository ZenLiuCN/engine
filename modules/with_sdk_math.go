//go:build (sdk_math || sdk || all) && !no_sdk && !no_sdk_math

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang/math"
	_ "github.com/ZenLiuCN/engine/modules/golang/math/big"
	_ "github.com/ZenLiuCN/engine/modules/golang/math/bits"
	_ "github.com/ZenLiuCN/engine/modules/golang/math/cmplx"
	_ "github.com/ZenLiuCN/engine/modules/golang/math/rand"
)
