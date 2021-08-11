package client_proxy

import (
	"mxshop_servs/json_rpc_new/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(network, address string) HelloServiceStub {
	client, err := rpc.Dial(network, address)
	if err != nil {
		panic("连接失败")
	}
	//rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	return HelloServiceStub{client}
}

func (h *HelloServiceStub) Hello(request string, reply *string) error {
	err := h.Call(handler.HelloService+".Hello", request, reply)
	if err != nil {
		panic(err)
	}
	return nil
}
