//go:build ((mysql && sqlx) || all) && !no_mysql

package modules

import _ "github.com/ZenLiuCN/engine/modules/go/sqlx/mysql_2023-12-22"
