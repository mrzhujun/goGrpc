package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Address string
}
type Mother struct {
	Name    string
	Age     int
	Address string
	Child   Person
}

type Ele interface {
	start()
	stop()
}

type mobile struct {
	name string
}

func (m mobile) start() {
	fmt.Println(m.name, "开机")
}
func (m mobile) stop() {
	fmt.Println(m.name, "关机")
}

func main() {
	zhangsan := Person{
		Name:    "张三",
		Age:     18,
		Address: "重庆",
	}

	var lisi Mother
	lisi.Age = 19
	lisi.Address = "北京"
	lisi.Name = "李四"
	lisi.Child = zhangsan

	jsonByte, err := json.Marshal(lisi)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr := string(jsonByte)

	fmt.Println(jsonStr)

	var m Ele
	m = mobile{name: "手机"}
	m.start()
}
