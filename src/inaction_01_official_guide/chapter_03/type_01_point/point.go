package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Printf("i = %v, &i = %v; j = %v, &j = %v \n", i, &i, j, &j)

	p, q := &i, &j
	fmt.Printf("*p = %v, *q = %v \n", *p, *q)

	// 通过指针 重新赋值!
	*p = 100
	fmt.Printf("i = %v \n", i)

	*q = *q / 37
	fmt.Printf("j = %v \n", j)
}
