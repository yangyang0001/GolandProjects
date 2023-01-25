package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Printf("num_cpu = %v \n", runtime.NumCPU())

	cores := runtime.NumCPU()
	channel := make(chan int)

	for i := 0; i < cores; i++ {
		go send(channel)
	}

	for i := 0; i < cores; i++ {
		go rece(channel)
	}

	time.Sleep(time.Second * 2)

}

func send(channel chan int)  {
	channel <- 10
}

func rece(channel chan int)  {
	fmt.Printf("channel element = %v \n", <-channel)
}
