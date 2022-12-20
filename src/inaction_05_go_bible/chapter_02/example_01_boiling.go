package main

import "fmt"

const boilingF = 212.0 // boilingF 声明的常量; 这种声明类似于 java 语言的全局变量, 可以在本文件中随意引用!

func main() {

	// 函数内部的变量, 这两个变量 只能在本法中使用, 数据局部变量!
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point is %g ℉ or %g ℃", f, c)
}
