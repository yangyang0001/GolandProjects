package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	channel := make(chan string)

	go send(channel)

	go rece(channel)

	time.Sleep(time.Second * 2)

}

func send(channel chan string)  {

	for i := 0; i < 10; i++ {
		channel <- strconv.Itoa(i)
	}

}

func rece(channel chan string)  {
	for val := range channel {
		fmt.Printf("val = %v \n", val)
	}
}




