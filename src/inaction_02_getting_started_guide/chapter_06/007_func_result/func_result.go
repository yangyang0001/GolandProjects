package main

import (
	"fmt"
	"time"
)

func main() {

	mine := getAdd()
	i, j := 10, 20

	start := time.Now()
	sum := mine(i, j)
	end := time.Now()

	cost := end.Sub(start)

	fmt.Printf("i = %v, j = %v, i + j = %v, cost = %v \n", i, j, sum, cost)

}

func getAdd() func(int, int) int {
	return func(a int, b int) int {
		return a + b
	}
}
