package main

import "fmt"

// iota 的使用, 初始值为0, 每一行都自动加1
const (
	A = iota
	B = iota
	C = iota
	D = iota
)

// 如果常量没有使用赋值, 则自动使用上一行的赋值语句
const (
	E = iota
	F
	G
	H
)

// 综合上面的两种特性, 同一行中的 iota 值是一致的, 下一行的 iota 值自动 +1; 如果没有赋值语句, 则自动使用上一行的赋值语句!
const (
	I, J = iota + 1, iota + 2 // iota = 0
	K    = "zhangsan"
	L
	M = iota // 一共 4 行, 这里 iota = 3
)

func main() {

	fmt.Println("----------------------------- 1、使用 iota 每一行自增特性 -------------------------------")
	fmt.Printf("A = %v \n", A)
	fmt.Printf("B = %v \n", B)
	fmt.Printf("C = %v \n", C)
	fmt.Printf("D = %v \n", D)

	fmt.Println("----------------------------- 2、常量的特性, 自动使用上一行的赋值 ------------------------")
	fmt.Printf("E = %v \n", E)
	fmt.Printf("F = %v \n", F)
	fmt.Printf("G = %v \n", G)
	fmt.Printf("H = %v \n", H)

	fmt.Println("----------------------------- 3、同一行的 iota 值一致, 下一行自动 +1 --------------------")
	fmt.Printf("I = %v, J = %v \n", I, J)
	fmt.Printf("K = %v \n", K)
	fmt.Printf("L = %v \n", L)
	fmt.Printf("M = %v \n", M)

}
