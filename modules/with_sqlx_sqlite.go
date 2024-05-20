//go:build ((sqlite && sqlx) || all) && !no_sqlqite

package modules

import _ "modernc.org/sqlite"
