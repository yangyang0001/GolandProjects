package main

import (
	"fmt"
	"time"
)

func main() {


	fmt.Println("-------------------------------------- 1、channel 阻塞等待计算 ----------------------------------")
	params := []int {1, 3, 5, 7, 9}
	channel := make(chan int)

	go sum(params, channel)

	fmt.Println("sum before running ...")

	time.Sleep(time.Second * 2)

	fmt.Printf("sum = %v \n", <-channel)

	time.Sleep(time.Second * 2)

	fmt.Println("sum after  running ...")

	fmt.Println("-------------------------------------- 2、channel 作为方法返回值 --------------------------------")
	tempchannel := getchannel()
	go usechannel(tempchannel)

	time.Sleep(time.Second * 2)

}

func sum(params []int, channel chan int) {

	var sum = 0
	for _, param := range params {
		sum += param
	}

	channel <- sum
}

func getchannel() chan int {

	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	return channel

}

func usechannel(channel chan int)  {
	for {
		fmt.Printf("channel element = %v \n", <-channel)
	}
}
