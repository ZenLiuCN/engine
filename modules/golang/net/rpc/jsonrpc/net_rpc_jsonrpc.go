// Code generated by define_gene; DO NOT EDIT.
package jsonrpc

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/rpc"
	"net/rpc/jsonrpc"
)

var (
	//go:embed net_rpc_jsonrpc.d.ts
	NetRpcJsonrpcDefine   []byte
	NetRpcJsonrpcDeclared = map[string]any{
		"newClient":      jsonrpc.NewClient,
		"newClientCodec": jsonrpc.NewClientCodec,
		"newServerCodec": jsonrpc.NewServerCodec,
		"serveConn":      jsonrpc.ServeConn,
		"dial":           jsonrpc.Dial,
	}
)

func init() {
	engine.RegisterModule(NetRpcJsonrpcModule{})
}

type NetRpcJsonrpcModule struct{}

func (S NetRpcJsonrpcModule) Identity() string {
	return "golang/net/rpc/jsonrpc"
}
func (S NetRpcJsonrpcModule) TypeDefine() []byte {
	return NetRpcJsonrpcDefine
}
func (S NetRpcJsonrpcModule) Exports() map[string]any {
	return NetRpcJsonrpcDeclared
}
