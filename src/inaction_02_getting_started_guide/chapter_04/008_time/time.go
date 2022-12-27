package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("--------------------------------- 1、格式化时间 ------------------------------")
	// 2006-01-02 15:04:05 这个时间 据说是 golang 语言诞生的时间, 这是固定写法, 不能随意修改!
	fmt.Printf("now time = :%v \n", time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))

}
