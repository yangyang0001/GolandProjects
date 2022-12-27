package main

import "fmt"

func main() {

	// 官网案例 hello world
	fmt.Println("----------------------------- 1、hello world 案例 -------------------------------")
	fmt.Println("hello, world")

	// 类型转换
	fmt.Println("----------------------------- 2、type(i) 类型转换 -------------------------------")
	var a int = 10
	b := float64(a)
	fmt.Printf("a = %v, a type is %T \n", a, a)
	fmt.Printf("b = %v, b type is %T \n", b, b)

	// 常量
	fmt.Println("----------------------------- 3、常量的定义和使用  -------------------------------")
	const c int = 10
	fmt.Printf("const c = %v, c type is %T \n", c, c)

	const c1, c2, c3 = "hello", 10, 100.0
	fmt.Printf("c1 = %v, c2 = %v, c3 = %v \n", c1, c2, c3)

	// 并行定义 常量
	const (
		Mon = 1
		Tue = 2
		Wed = 3
		Thu = 4
		Fri = 5
		Sat = 6
		Son = 7
	)
	fmt.Printf("Mon = %v \n", Mon)
	fmt.Printf("Tue = %v \n", Tue)
	fmt.Printf("Wed = %v \n", Wed)
	fmt.Printf("Thu = %v \n", Thu)
	fmt.Printf("Fri = %v \n", Fri)
	fmt.Printf("Sat = %v \n", Sat)
	fmt.Printf("Son = %v \n", Son)

}
