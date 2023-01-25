package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)

	go senddata(channel)
	go recedata(channel)

	time.Sleep(time.Second * 2)

}

func senddata(channel chan string)  {
	channel <- "Washington"
	channel <- "Tripoli"
	channel <- "London"
	channel <- "Beijing"
	channel <- "Tokio"
	// 只有发送方, 才能主动关闭 channel
	close(channel)
}

func recedata(channel chan string)  {
	//for {
	//	res, open := <-channel
	//	if open {
	//		fmt.Printf("res = %v \n", res)
	//	} else {
	//		fmt.Println("this string channel is closed!")
	//		break
	//	}
	//}

	// for-range 是能够自动检测 channel 是否关闭的!
	for res := range channel {
		fmt.Printf("res = %v \n", res)
	}
}


