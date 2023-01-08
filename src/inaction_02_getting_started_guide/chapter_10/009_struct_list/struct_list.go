package main

import "fmt"

type List []int

func main() {

	var list List
	for i := 0; i < 10; i++ {
		list = append(list, i)
		fmt.Printf("len(list) = %v, cap(list) = %v, list = %v \n", len(list), cap(list), list)
	}

}
