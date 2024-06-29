//go:build (bbolt || all) && !no_bbolt

package modules

import _ "github.com/ZenLiuCN/engine/modules/go/etcd/io/bbolt"
