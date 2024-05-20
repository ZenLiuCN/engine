//go:build (sdk_tzdata || sdk || all) && !no_sdk && !no_sdk_tzdata

package time

import (
	_ "time/tzdata"
)
