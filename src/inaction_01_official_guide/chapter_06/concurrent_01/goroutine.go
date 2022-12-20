package main

import (
	"fmt"
	"time"
)

func main() {
	go SayHello("maosan") // 以目前的知识了解到: 这个只能写在 第一行, 否则不执行!
	SayHello("zhangsan")
}

func SayHello(name string) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Hello %v, i = %v, time = %v \n", name, i, NowTime())
	}
}

func NowTime() string {
	// 2006-01-02 15:04:05 这个时间 据说是 golang 语言诞生的时间, 这是固定写法, 不能随意修改!
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}
