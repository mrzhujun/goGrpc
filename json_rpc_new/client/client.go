package main

import (
	"fmt"
	"mxshop_servs/json_rpc_new/client_proxy"
)

func main() {
	client := client_proxy.NewHelloServiceClient("tcp", ":12345")
	var reply string
	err := client.Hello("hello", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
