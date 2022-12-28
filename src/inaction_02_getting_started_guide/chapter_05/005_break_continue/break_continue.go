package main

import "fmt"

func main() {

	fmt.Println("-------------------------- 1、break method -------------------------")
	break_method()

	fmt.Println("-------------------------- 2、continue method ----------------------")
	continue_method()

	fmt.Println("-------------------------- 3、break 等价于 continue 方法 ------------")
	break_equal_continue_method()

}

func break_method() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			break
		}
		fmt.Printf("i = %v \n", i)
	}
}

func continue_method() {

	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		} else {
			fmt.Printf("i = %v \n", i)
		}

	}
}

func break_equal_continue_method() {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 {
				// 这里的 break 和 continue 是一致的!
				//continue
				break
			} else {
				fmt.Printf("i = %v, j = %v \n", i, j)
			}
		}
	}

}
