//go:build (all || (duckdb && sqlx)) && !no_duckdb

package modules

import _ "github.com/ZenLiuCN/engine/modules/go/sqlx/duckdb"
