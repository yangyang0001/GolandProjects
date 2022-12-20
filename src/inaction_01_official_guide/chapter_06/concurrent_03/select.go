package main

import "fmt"

func main() {

	c := make(chan int)
	q := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("out method invoke, channle.value = %v \n", <-c)
		}

		q <- 0
	}()

	mine(c, q)

}

func mine(c, quite chan int) {
	var count int = 0

	x, y := 0, 1

	for {
		count++
		fmt.Printf("count = %v \n", count)

		select {
		case c <- x:
			fmt.Printf("01_first case invoke ..., x = %v, y = %v, x+y = %v \n", x, y, x+y)
			x, y = y, x+y
			fmt.Printf("02_first case invoke ..., x = %v, y = %v, x+y = %v \n", x, y, x+y)
		case <-quite:
			fmt.Printf("quite case invoke ..., count = %v \n", count)
			return
		default:
			fmt.Printf("default case invoke ... count = %v \n", count)
		}

		fmt.Println("--------------------------------------------------------------")
	}

}
