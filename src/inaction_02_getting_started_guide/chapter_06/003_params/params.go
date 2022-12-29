package main

import "fmt"

func main() {

	fmt.Println("--------------------------------- 1、变长参数只用案例 ------------------------------")
	i, j, m, n := 1, 3, 5, 7
	sum := add(i, j, m, n)
	fmt.Printf("sum = %v \n", sum)

}

func add(params ...int) int {

	var sum = 0
	for _, param := range params {
		sum += param
	}

	return sum
}