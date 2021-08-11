package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":12345")
	if err != nil {
		panic(err)
	}
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
