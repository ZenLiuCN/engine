//go:build (sdk_net_all || sdk || all) && !no_sdk && !no_sdk_net_all

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang/net"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/http"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/http/cgi"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/http/cookiejar"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/http/fcgi"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/http/pprof"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/mail"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/netip"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/rpc"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/rpc/jsonrpc"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/smtp"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/textproto"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/url"
)
