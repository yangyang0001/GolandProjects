package main

import (
	"fmt"
)

func main() {

	var channel = SequenceChannel();
	fmt.Printf("sequence = %v \n", GetSequence(channel))
	fmt.Printf("sequence = %v \n", GetSequence(channel))
	fmt.Printf("sequence = %v \n", GetSequence(channel))
	fmt.Printf("sequence = %v \n", GetSequence(channel))

	//for {
	//	sequence := GetSequence(channel)
	//	fmt.Printf("sequence = %v \n", sequence)
	//	time.Sleep(time.Second)
	//}

}

func SequenceChannel() chan int {

	channel := make(chan int)
	count := 0

	go func() {
		for {
			channel <- count
			count++
		}
	}()

	return channel
}

func GetSequence(channel chan int) int {
	return <- channel
}
