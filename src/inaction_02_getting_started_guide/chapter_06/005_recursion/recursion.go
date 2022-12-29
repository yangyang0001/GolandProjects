package main

import "fmt"

func main() {

	fmt.Println("------------------------------------- 1、斐波那契数列 -------------------------------------")
	for i := 1; i <= 10; i++ {
		result := fabinacci(i)
		fmt.Printf("fabinacci(%v) = %v \n", i, result)
	}

	fmt.Println("------------------------------------- 2、求 N 的阶乘 -------------------------------------")
	for i := 1; i <= 10; i++ {
		result := factorial(i)
		fmt.Printf("factorial(%v) = %v \n", i, result)
	}

}

// 斐波那契数列, 求 第N项的值!
func fabinacci(n int) int {

	if n <= 2 {
		return 1
	}

	return fabinacci(n-1) + fabinacci(n-2)

}

// 阶乘
func factorial(n int) int {

	if n == 1 {
		return 1
	}

	return factorial(n-1) * n
}
