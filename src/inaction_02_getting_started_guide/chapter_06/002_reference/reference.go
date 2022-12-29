package main

import "fmt"

func main() {

	fmt.Println("--------------------------------- 1、引用传递函数使用案例 ------------------------------")
	i := 10
	update_int(&i)
	fmt.Printf("i = %v \n", i)

}

func update_int(pint *int) {
	*pint = 100
}
