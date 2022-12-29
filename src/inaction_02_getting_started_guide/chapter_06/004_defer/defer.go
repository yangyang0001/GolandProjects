package main

import (
	"fmt"
	"io"
)

func main() {

	fmt.Println("-------------------------- 1、defer 案例1  ----------------------------")
	function1()

	fmt.Println("-------------------------- 2、defer 案例2  ----------------------------")
	f1()

	fmt.Println("-------------------------- 3、defer 案例3  ----------------------------")
	f2()

	fmt.Println("-------------------------- 4、defer 打印结果案例  -----------------------")
	f3("zhangsan")

}

func function1() {

	fmt.Println("function1 invoke start ...")

	defer function2()

	fmt.Println("function1 invoke end   ...")

}

func function2() {

	fmt.Println("function2 invoke ...")

}

func f1() {
	i := 0
	defer fmt.Printf("a method invoke i = %v \n", i)
	i++
}

func f2() {

	for i := 0; i < 10; i++ {
		defer fmt.Printf("i = %v \n", i)
	}

}

func f3(s string) (n int, err error) {

	defer func() {
		fmt.Printf("f3(%v) = %v, error = %v \n", s, n, err)
	}()

	return 7, io.EOF
}
