//go:build ((pgx && sqlx) || all) && !no_pgx

package modules

import _ "github.com/ZenLiuCN/engine/modules/go/sqlx/pgx"
