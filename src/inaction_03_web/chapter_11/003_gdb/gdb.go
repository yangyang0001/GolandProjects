package main

import (
	"fmt"
	"time"
)

func counting(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}
	close(c)
}

func main() {

	fmt.Println("-------------------------- main method invoke starting ----------------------------")
	bus := make(chan int)

	fmt.Println("-------------------------- goroutine method invoke starting -----------------------")
	go counting(bus)

	for count := range bus {
		time.Sleep(1 * time.Second)
		fmt.Printf("count = %v \n", count)
	}

}
