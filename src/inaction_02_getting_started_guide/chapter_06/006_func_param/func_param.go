package main

import "fmt"

func main() {
	sum := function1(10, 20, add)
	fmt.Printf("sum = %v \n", sum)

}

func function1(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}
