package main

import "fmt"

func main() {

	fmt.Println("*************************************** 1、第1个信道案例 ***************************************")
	arr := []int{7, 2, 8, -9, 4, 100}
	channel := make(chan int)

	go sum(arr[1:3], channel, "first") // 10
	go sum(arr[4:6], channel, "secon") // 104		先执行!

	x, y := <-channel, <-channel

	// x = 104, y = 10, x+y = 114
	fmt.Printf("x = %v, y = %v, x+y = %v \n", x, y, x+y)

	fmt.Println("*************************************** 2、斐波那契数列 ***************************************")
	go fibonacci(10, channel)
	for val := range channel {
		fmt.Println(val)
	}

}

func sum(s []int, channel chan int, name string) {
	fmt.Printf("sum method invoke, name = %v \n", name)

	var res int = 0

	for _, val := range s {
		res += val
	}

	channel <- res
}

func fibonacci(n int, channel chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		channel <- x
		x, y = y, x+y
	}
	close(channel)
}
