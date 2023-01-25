package main

import (
	"fmt"
)

func main() {

	channel := make(chan string)

	go SendData(channel)

	go ReceData(channel)

	// 协程的生命周期 必须建立在 主线程的 生命周期之内!
	time.Sleep(time.Second * 1000)

}

func SendData(sendchan chan string)  {
	fmt.Println("send data method invoke start ...")

	sendchan <- "zhangsan"
	sendchan <- "lisi"
	sendchan <- "wangwu"
	sendchan <- "zhaoliu"
}

func ReceData(recechan chan string)  {
	fmt.Println("receive data method invoke start ...")

	var data string

	for {
		data = <- recechan
		fmt.Printf("data = %v \n", data)
	}
}

