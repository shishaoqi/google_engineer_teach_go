package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP	string
}

type Serverslice struct {
	Servers []Server
}


func main() {
	// 解析到结构体
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`

	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	fmt.Printf("相应值的默认格式 %v\n", s) // 与 fmt.Println(s) 一样
	fmt.Printf("打印结构体时，会添加字段名 %+v\n", s)

	// ================== 解析到 interface =========================
	fmt.Println("解析到 interface")
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`) // []byte 字节切片
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Println(f)

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k,"is float64",vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
