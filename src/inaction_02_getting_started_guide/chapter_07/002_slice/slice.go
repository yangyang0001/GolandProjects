package main

import "fmt"

func main() {

	fmt.Println("-------------------------- 1、切片的声明和案例  ----------------------------")
	arr := []int{11, 32, 55, 37, 29, 100}
	sli := arr[1:3]
	for index, val := range sli {
		fmt.Printf("arr[%v] = %v \n", index, val)
	}

	fmt.Printf("slice len = %v, cap = %v \n", len(sli), cap(sli))

	fmt.Println("-------------------------- 2、使用切片代表数组  ----------------------------")
	strings := []string{"zhangsan", "lisi", "wangwu", "zhaoliu"}
	strslic := strings[:]

	for index, val := range strslic {
		fmt.Printf("strslic[%v] = %v \n", index, val)
	}

	fmt.Println("-------------------------- 3、声明数组切片案例  ----------------------------")
	ints := []int{1, 3, 5, 7, 9}[:]
	for index, val := range ints {
		fmt.Printf("ints[%v] = %v \n", index, val)
	}

	fmt.Println("-------------------------- 4、官网中的案例实现  ----------------------------")
	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for index, val := range slice1 {
		fmt.Printf("first slice1[%v] = %v \n", index, val)
	}

	fmt.Printf("first the length of arr1 is %d \n", len(arr1))       // 6
	fmt.Printf("first the length of slice1 is %d \n", len(slice1))   // 3
	fmt.Printf("first the capacity of slice1 is %d \n", cap(slice1)) // 4

	slice1 = slice1[0:4]
	for index, val := range slice1 {
		fmt.Printf("second slice1[%v] = %v \n", index, val)
	}
	fmt.Printf("second the length of slice1 is %d \n", len(slice1))   // 4
	fmt.Printf("second the capacity of slice1 is %d \n", cap(slice1)) // 4

	fmt.Println("-------------------------- 4、将切片传递给函数  ----------------------------")
	var array0 = [5]int{0, 1, 2, 3, 4}
	s := sum(array0[:])
	fmt.Printf("s = %v \n", s)

	fmt.Println("-------------------------- 5、使用 make 创建切片  ----------------------------")
	var slice_1 = make([]int, 10)

	for i := 0; i < len(slice_1); i++ {
		slice_1[i] = 5 * i
	}

	// print the slice:
	for i := 0; i < len(slice_1); i++ {
		fmt.Printf("five slice_1[%v] = %v \n", i, slice_1[i])
	}
	fmt.Printf("five the length of slice_1 is %d \n", len(slice_1))
	fmt.Printf("five the capacity of slice_1 is %d \n", cap(slice_1))
}

func sum(arr []int) int {
	var s = 0
	for _, val := range arr {
		s += val
	}

	return s
}
