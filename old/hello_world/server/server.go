package main

import (
	"net"
	"net/rpc"
)

type HelloServer struct{}

func (h HelloServer) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {
	listen, _ := net.Listen("Hello", ":1234")
	_ = rpc.RegisterName("Hello", HelloServer{})

	conn, _ := listen.Accept()
	rpc.ServeConn(conn)
}
