package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ShopRpcServ struct{}

func (s ShopRpcServ) Order(request string, reply *string) error {
	*reply = request + "：成功下单"
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	_ = rpc.RegisterName("OrderServer", ShopRpcServ{})

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
