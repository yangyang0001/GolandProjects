package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func main() {

	var client goredis.Client
	// 设置端口为redis默认端口
	client.Addr = "192.168.188.49:6379"

	// string 操作
	client.Set("a", []byte("hello"))
	value, err := client.Get("a")
	if err != nil {
		fmt.Printf("get error = %v \n", err)
	}
	fmt.Printf("key = %v, value = %v \n", "a", string(value))
	// client.Del("a")

	// list 操作
	values := []string{"a", "b", "c", "d", "e"}
	for _, v := range values {
		client.Rpush("list", []byte(v))
	}
	dbvals,_ := client.Lrange("list", 0, 4)
	for index, v := range dbvals {
		fmt.Printf("index = %v, value = %v \n", index, string(v))
	}
	// client.Del("l")
}
