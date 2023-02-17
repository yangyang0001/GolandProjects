package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("--------------------------------- 1、格式化时间案例 ---------------------------------")
	minetime := time.Now()
	minetime = minetime.AddDate(0, 0, 1)
	fmt.Printf("minetime = %v \n", time.Unix(minetime.Unix(), 0).Format("2006-01-02 15:04:05"))


	fmt.Println("--------------------------------- 2、定时任务案例   ---------------------------------")
	ticke := time.Tick(time.Second)
	after := time.After(time.Second * 10)

	for {
		select {
		case <-ticke:
			fmt.Printf("time = %v \n", time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
		case <-after:
			fmt.Printf("time job is end! \n")

			// 加上 return 后可以在 运行了 10s 之后 退出!
			return
		}
	}

}
