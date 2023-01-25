package main

import (
	"fmt"
	"time"
)

func main() {

	//channel := make(chan int)
	//go SendData(channel)
	//go ReceData(channel)
	//
	//time.Sleep(time.Second * 1000)

	channel := make(chan int)

	go f1(channel)
	go f0(channel)
	time.Sleep(time.Second * 10)

}

func SendData(channel chan int)  {
	fmt.Println("send data method invoke start!")
	for i := 0; ; i++ {
		channel <- i
	}
}

func ReceData(channel chan int)  {
	fmt.Println("receive data method invoke start!")
	for {
		fmt.Printf("channel = %v \n", <-channel)
		time.Sleep(time.Second * 1)
	}
}

func f1(channel chan int)  {
	fmt.Printf("channel = %v \n", <-channel)
}

func f0(channel chan int)  {
	channel <- 2
}

