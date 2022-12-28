package main

import "fmt"

func main() {
	fmt.Println("-------------------------- 1、案例1 -------------------------")
	exmaple_method_01()

	fmt.Println("-------------------------- 2、案例2 -------------------------")
	exmaple_method_02()

}

func exmaple_method_01() {
	fmt.Println("before invoke ...")
	goto Lable1
	fmt.Println("after  invoke ...")

Lable1:
	fmt.Println("hello world")
}

func exmaple_method_02() {
	a := 1
	b := 9
	goto TARGET

TARGET:
	b += a
	fmt.Printf("a = %v, b = %v \n", a, b)
}
