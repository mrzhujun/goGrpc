package main

import "fmt"

func Add(a, b int) int {
	total := a + b
	return total
}

type Company struct {
	Name    string
	Address string
}

type Employee struct {
	Name    string
	company Company
}

type PrintResult struct {
	Info string
	Err  error
}

func RpcPrintln(employee Employee) PrintResult {
	//http协议来说 有一个问题：一次性 一旦对方返回了结果 连接端口
	//http2.0 长链接  grpc

	//RPC中的第二个点 传输协议、数据编码协议
	//http1.x协议  http2.0协议
	//http协议底层使用的也是tcp协议  http现在主流的事1.x 这种协议有性能问题 一次性 一旦结果返回连接就断开
	//1.自己基于tcp/udp协议去封装一层协议 myhttp 没有通用性，http2.0 既有http的特性也有长链接的特性 grpc

	/*
		客户端
		1.建立连接 tcp/http
		2.将Employee对象序列化成json字符串 - 序列化
		3.发送json字符串 - 调用成功后实际上接受到的是一个二进制数据
		4.等待服务器发送结果
		5.将服务器返回的数据解析成 PrintResult对象 - 反序列化
		服务端
		1.监听网络端口
		2.读取数据 - 二进制的json数据
		3.将数据反序列化Employee对象
		4.开始处理业务服务
		5.将处理的结果PrintResult序列化成json二进制数据 - 序列化
		6.将数据返回

		序列化和反序列化是可以选择的，不一定要采用json、xml、protobuf、msgpack
	*/
}

func main() {
	//fmt.Println(Add(1,2))

	//讲这个打印的工作放在另一台服务器上，我需要讲本地的内存对象struct、这样不行，可行的方式是将struct序列化成json
	fmt.Println(Employee{
		Name: "Bobby",
		company: Company{
			Name:    "慕课网",
			Address: "北京市",
		},
	})

	//远程的服务器需要将二进制对象反解析成struct对象
	//搞这么麻烦，直接全部使用json不香吗？这种做法在浏览器和gin服务之间是可行的，但如果你是一个大型的分布式系统，无法维护
}
