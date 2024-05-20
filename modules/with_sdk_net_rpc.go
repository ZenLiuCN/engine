//go:build (sdk_net_rpc || sdk || all) && !no_sdk && !no_sdk_net_rpc

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang/net"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/rpc"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/rpc/jsonrpc"
)
