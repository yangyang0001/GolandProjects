package main

import "fmt"

func main() {

	channel := make(chan int)
	fmt.Printf("cap(channel) = %v \n", cap(channel))
	numbers := []int {1, 3, 5, 7, 9}

	go sum(numbers, channel)
	fmt.Printf("sum = %v \n", <-channel)


}

func sum(numbers []int, channel chan int)  {

	var sum = 0

	for _, number := range numbers {
		sum += number
	}

	channel <- sum
}

