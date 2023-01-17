package main

import (
	"fmt"
	"os"
)

func main() {

	//fmt.Println("------------------------------------- 1、panic 案例1 -----------------------------------")
	//panic_method1()

	fmt.Println("------------------------------------- 2、panic 案例2 -----------------------------------")
	panic_method2()

}

func panic_method1()  {

	fmt.Println("panic method1 invoke start!")

	panic("panic_method1 runtime error!")

	fmt.Println("panic method1 invoke end!")

}

func panic_method2()  {

	user := os.Getenv("USER")

	if user == "" {
		panic("user is null error!")
	} else {
		fmt.Printf("user = %v \n", user)
	}

}
