package main

import "fmt"

const (
	a = iota
	b = "zhangsan"
	c
	d = iota
	e
	f = 3.1415926
	g = iota
)

const h = iota

const (
	i, j, k = iota, iota, iota
)

/*
	验证 iota 的值, 以及 go 语言中的, 赋值规则!

	iota 遇到 const 重置为 0
	iota 同一行中的值是一致的
	iota 下一行的值自动+1

*/
func main() {

	fmt.Printf("a = %v \n", a)		// a = 0
	fmt.Printf("b = %v \n", b)		// b = zhangsan
	fmt.Printf("c = %v \n", c)		// c = zhangsan
	fmt.Printf("d = %v \n", d)		// d = 3
	fmt.Printf("e = %v \n", e)		// e = 4
	fmt.Printf("f = %v \n", f)		// f = 3.1415926


	fmt.Printf("g = %v \n", g)		// g = 6


	fmt.Printf("i = %v \n", i)		// i = 0
	fmt.Printf("j = %v \n", j)		// j = 0
	fmt.Printf("k = %v \n", k)		// k = 0

}
