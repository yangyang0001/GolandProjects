package main

import (
	"fmt"
	"time"
)

func main() {


	fmt.Println("----------------------- 1、阻塞测试, 在一个方法中只能处理一个 channel 否则阻塞 ---------------------")
	channel0 := make(chan int)
	go send_channel(channel0)
	go rece_channel(channel0)
	time.Sleep(time.Second * 1)
	fmt.Println("----------------------- 2、多次使用单通道方法, 处理不同的 channel 进行 select ---------------------")

	channel1 := make(chan int)
	channel2 := make(chan int)
	go send_channel(channel1)
	go send_channel(channel2)
	go channel_select(channel1, channel2)
	time.Sleep(time.Second * 2)

}

func send_channel(channel0 chan int)  {
	for i := 0; i < 10; i++ {
		channel0 <- i
	}
}

func rece_channel(channel0 chan int)  {

	for val0 := range channel0 {
		fmt.Printf("channel0 val0 = %v \n", val0)
	}

}

func channel_select(channel1, channel2 chan int)  {

	for {
		select {
		case val1 := <-channel1:
			fmt.Printf("channel1 val1 = %v \n", val1)

		case val2 := <-channel2:
			fmt.Printf("channel2 val2 = %v \n", val2)

		default:
			fmt.Println("all the channels is blocking!")
		}
		time.Sleep(time.Millisecond * 100)
	}

}