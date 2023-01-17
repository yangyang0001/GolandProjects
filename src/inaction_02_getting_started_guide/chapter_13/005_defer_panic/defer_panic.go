package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	fmt.Printf("start time = %v \n", start)

	fmt.Println("defer panic running start ...")

	ExecuteDefer()

	fmt.Println("defer panic running end!")

	cost  := time.Now().Sub(start)
	fmt.Printf("cost  time = %v \n", cost)


	/**
		执行结果:
		defer panic running start ...

		ExecutePanic num = 3
		ExecutePanic num = 2
		ExecutePanic num = 1
		ExecutePanic num = 0

		panicing error, num cannot more than 3, num = 4
		defer panic running end!
	*/
}

func ExecuteDefer()  {

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panicing error, num cannot more than 3, num = %v \n", e)
		}
	}()

	ExecutePanic(0);

	fmt.Println("ExecuteDefer method invoke end!")
}

func ExecutePanic(num int)  {

	if num > 3 {
		panic(num)
	}

	defer fmt.Printf("ExecutePanic num = %v \n", num)

	ExecutePanic(num + 1)

	fmt.Println("ExecutePanic method invoke end!")
}







