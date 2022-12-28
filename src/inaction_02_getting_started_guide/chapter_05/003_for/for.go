package main

import (
	"fmt"
)

func main() {

	fmt.Println("-------------------------- 1、计数器迭代 ----------------------------")
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %v \n", i)
	}

	fmt.Println("-------------------------- 2、实现字符串倒转 -------------------------")
	reve := reverse("hello")
	fmt.Printf("revese method invoke, result = %v \n", reve)

	fmt.Println("-------------------------- 3、变量的交换案例 -------------------------")
	var i, j = 100, 88
	i, j = j, i
	fmt.Printf("i = %v, j = %v \n", i, j)
}

func reverse(str string) string {

	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)

}
