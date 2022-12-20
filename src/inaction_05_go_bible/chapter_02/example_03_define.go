package main

import (
	"fmt"
)

func main() {

	// 验证了 string 类型的 零值 为 "" 空字符串!
	//var s0 string
	//var s1 string = "[]" + s0 + "[]"
	//fmt.Println(s1)
	//fmt.Println("------------------------------------------------")
	//
	//var i, j, k int
	//var m, n, o = true, 2.3, "four"
	//fmt.Println(i, j, k)
	//fmt.Println(m, n, o)
	//fmt.Println("------------------------------------------------")
	//
	//// TODO: 此种声明方式, 这里待定, 因为没有具体的案例!
	//// var p, q = os.Open("zhangsan")
	//fmt.Println("------------------------------------------------")
	//
	//// 这种声明方式应该 推荐只在 for 中使用, 因为可读性不好!
	//a, b := 1, 3
	//fmt.Println(a, b)
	//fmt.Println("------------------------------------------------")

	// 交换两个值
	var c = 8
	var d = 10
	fmt.Println(c, d)

	c, d = d, c
	fmt.Println(d, c)
	fmt.Println("------------------------------------------------")

}
