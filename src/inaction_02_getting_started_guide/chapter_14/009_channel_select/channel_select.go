package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	
	channel1 := make(chan string)
	channel2 := make(chan string)

	go SendData1(channel1)
	go SendData2(channel2)

	go ReceData(channel1, channel2)

	time.Sleep(time.Second * 100)
	
}

func SendData1(channel chan string)  {

	for i := 0; ; i++ {
		channel <- strconv.Itoa(i*2)
	}

}

func SendData2(channel chan string)  {
	for i := 0; ; i++ {
		channel <- strconv.Itoa(i+5)
	}
}


func ReceData(channel1 chan string, channel2 chan string) {

	for {
		time.Sleep(time.Second * 1)
		select {
		case element1:= <-channel1:
			fmt.Printf("element1 = %v \n", element1)
		case element2:= <-channel2:
			fmt.Printf("element2 = %v \n", element2)
		default:
			fmt.Println("there is none data!")
		}
	}

}
