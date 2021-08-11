package main

import (
	"mxshop_servs/json_rpc_new/handler"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (h HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		panic(err)
	}
	err = rpc.RegisterName(handler.HelloService, &HelloService{})
	if err != nil {
		panic(err)
	}
	for {
		conn, _ := listen.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
