package main

import (
	"fmt"
)

// 本案例使用的是 channel 的阻塞, 来构建一个 channel 链!
// 	channel0 <- (1 + <-channel1)
// 		|			   channel1 <- (1 + <-channel2)
// 		|								  channel2 <- (1 + <-channel3)
// 		|													 ........
// 		|
// 		|
//	leftmost channel 定义最左 通道, 以备后续使用! 最终展示 <-leftmost 的值!
func main() {

	chain()

}


func chain() {

	leftmost := make(chan int)
	var left, right chan int = nil, leftmost

	for i := 0; i < 1000; i++ {
		left, right = right, make(chan int)
		// 通过 协程使用 channel, 必须使用 go!
		go dochain(left, right)
	}

	right <- 0
	result := <-leftmost

	fmt.Printf("result = %v \n", result)

}

func dochain(left chan int, right chan int)  {
	left <- 1 + <-right
}
