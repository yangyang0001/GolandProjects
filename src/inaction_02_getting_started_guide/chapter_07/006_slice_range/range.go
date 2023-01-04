package main

import "fmt"

func main() {

	fmt.Println("-------------------------- 1、for range 案例  ----------------------------")
	var ints = []int{1, 3, 5, 7, 9}
	var sum = 0
	for _, val := range ints {
		sum += val
	}
	fmt.Printf("sum = %v \n", sum)

	fmt.Println("-------------------------- 2、二维切片 案例  ----------------------------")
	var arrs = [][]string{
		{"zhangsan", "lisi"},
		{"wangwu", "zhaoliu"},
		{"maosan", "gosi"},
	}

	for row, arr := range arrs {
		for col, val := range arr {
			fmt.Printf("arrs[%v][%v] = %v \n", row, col, val)
		}
	}

	fmt.Println("-------------------------- 3、range 修改值 ----------------------------")
	var items = []int{1, 3, 5, 7, 9}
	for _, item := range items {
		item = item * 2
	}
	for index, item := range items {
		fmt.Printf("items[%v] = %v \n", index, item)
	}

	fmt.Println("-------------------------- 4、切片边缘问题 ----------------------------")
	var arr = [6]int{1, 3, 5, 7, 9, 11}
	slice1 := arr[6:6] // 包含又不包含, 就是空切片!
	// slice2 := arr[6:7]						// 编译报错! 因为越界了!

	fmt.Printf("slice1 = %v \n", slice1)

}
