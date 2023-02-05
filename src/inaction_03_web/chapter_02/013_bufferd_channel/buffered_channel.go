package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("----------------------------------- 1、buffered channel 案例1 -----------------------------------")
	channel := make(chan int, 2)
	channel <- 1
	channel <- 2

	go channel_each(channel)
	time.Sleep(time.Second * 2)

	fmt.Println("----------------------------------- 2、buffered channel 案例2 -----------------------------------")
	buffered_channel := make(chan int, 10)
	go fibonacci(cap(buffered_channel), buffered_channel)
	go fibonacci_show(buffered_channel)
	time.Sleep(time.Second * 2)


}


func channel_each(channel chan int)  {
	for val := range channel {
		fmt.Printf("val = %v \n", val)
	}
}


func fibonacci(n int, channel chan int)  {
	var x, y int = 1, 1
	for i := 1; i <= n; i++ {
		channel <- x
		x, y = y, x+y
	}
}

func fibonacci_show(buffered_channel chan int)  {
	var index int = 1
	for val := range buffered_channel {
		fmt.Printf("fibonacci index = %v, value = %v \n", index, val)
		index++
	}
}