package main

import "fmt"

func main() {

	fmt.Println("-------------------------- 1、copy() 方法  ----------------------------")
	slice_from := []int{1, 3, 5, 7, 9}
	//slice_dest := make([]int, 3)
	slice_dest := make([]int, 10)

	n := copy(slice_dest, slice_from)
	fmt.Printf("copy method invoke, n = %v \n", n)
	for index, val := range slice_dest {
		fmt.Printf("slice_dest[%v] = %v \n", index, val)
	}

	fmt.Println("-------------------------- 2、append 方法  --------------------------")
	slice3 := []int{1, 2, 3}
	slice3 = append(slice3, 4, 5, 6)
	fmt.Printf("append method invoke, slice3 = %v \n", slice3)

	fmt.Println("-------------------------- 3、扩展案例   ----------------------------")
	var s []int = []int{1}
	fmt.Printf("before len(s) = %v, cap(s) = %v, s = %v \n", len(s), cap(s), s)

	for _, val := range s {
		fmt.Printf("before val = %v \n", val)
	}

	factor := 10
	s = make([]int, len(s)*factor)
	fmt.Printf("after  len(s) = %v, cap(s) = %v, s = %v \n", len(s), cap(s), s)
	for _, val := range s {
		fmt.Printf("after  val = %v \n", val)
	}
}
