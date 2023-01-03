package main

import "fmt"

func main() {

	fmt.Println("-------------------------- 1、数组声明 和 遍历  ----------------------------")
	var arr [5]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}

	for index, value := range arr {
		fmt.Printf("arr[%v] = %v \n", index, value)
	}

	fmt.Println("-------------------------- 2、省略长度的方式  -----------------------------")
	a := [...]string{"a", "b", "c", "d"}

	fmt.Printf("len(a) = %v \n", len(a))
	for i := range a {
		fmt.Printf("a[%v] = %v \n", i, a[i])
	}

	fmt.Println("-------------------------- 3、new 声明方式  ------------------------------")
	var arr1 = new([5]int)
	var arr2 = [5]int{1, 3, 5, 7, 9}
	arr2 = *arr1
	fmt.Printf("arr1 = %v \n", arr1)
	fmt.Printf("arr2 = %v \n", arr2)

}
