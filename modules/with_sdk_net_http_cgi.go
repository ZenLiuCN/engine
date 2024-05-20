//go:build (sdk_net_http_cgi || sdk || all) && !no_sdk && !no_sdk_net_http_cgi

package modules

import (
	_ "github.com/ZenLiuCN/engine/modules/golang/net/http/cgi"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/http/fcgi"
)
