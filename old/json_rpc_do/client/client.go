package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	reply := new(string)
	err = client.Call("OrderServer.Order", "bbq", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(*reply)
}
