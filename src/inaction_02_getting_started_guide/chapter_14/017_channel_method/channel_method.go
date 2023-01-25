package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("--------------------------------------- 1、无参通道 ------------------------------------")
	channel0 := make(chan func())
	go send0(channel0)
	go rece0(channel0)
	time.Sleep(time.Second)

	fmt.Println("--------------------------------------- 2、有参通道 ------------------------------------")
	channel1 := make(chan func(int, int))
	go send1(channel1)
	go rece1(channel1)
	time.Sleep(time.Second)

}

func send0(channel chan func())  {
	channel <- mine
}

func rece0(channel chan func())  {
	mine := <-channel
	mine()
}

func send1(channel chan func(int, int))  {
	channel <- add
}

func rece1(channel chan func(int, int))  {
	add := <-channel
	add(100, 200)
}


func mine()  {
	fmt.Println("mine method invoke ...")
}

func add(a, b int) {
	fmt.Printf("sum = %v \n", a + b)
}
