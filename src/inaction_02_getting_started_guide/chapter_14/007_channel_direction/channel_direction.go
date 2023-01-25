package main

import (
	"fmt"
	"time"
)

func main() {

	getdatachan := make(<-chan int)

	recdatachan := make(chan<- int)

	go send_only(getdatachan)

	go rece_only(recdatachan)

	time.Sleep(time.Second * 10)

}

func send_only(channel <-chan int)  {
	fmt.Printf("channel = %v \n", channel)
}

func rece_only(channel chan<- int)  {
	for i := 0; i < 10; i++ {
		channel <- i
	}
}
