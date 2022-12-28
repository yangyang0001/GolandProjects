package main

import (
	"fmt"
)

func main() {

	ints := []int{1, 30, 100, 20, 1000}

	maps := map[string]string{
		"A": "C罗",
		"B": "梅西",
		"C": "内马尔",
		"D": "姆巴佩",
	}
	fmt.Println("-------------------------- 1、for range arr -------------------------")
	example_method_01(ints)

	fmt.Println("-------------------------- 2、for range map -------------------------")
	example_method_02(maps)

	fmt.Println("-------------------------- 3、查看结果 案例   -------------------------")
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v) // 0 0 0 0 0
		v = 5
	}

	fmt.Println("-------------------------- 4、查看结果 案例   -------------------------")
	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i) // 死循环
	//}

	fmt.Println("-------------------------- 5、查看结果 案例   -------------------------")
	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i) // 死循环
	//}

	fmt.Println("-------------------------- 6、查看结果 案例   -------------------------")
	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s) // "", "a", "aa", "aaa", "aaaa"
		s = s + "a"
	}

	fmt.Println("-------------------------- 7、查看结果 案例   -------------------------")
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s) // 最多3行
	}
}

func example_method_01(ints []int) {

	for i, v := range ints {
		fmt.Printf("i = %v, v = %v \n", i, v)
	}

}

func example_method_02(m map[string]string) {

	for k, v := range m {
		fmt.Printf("k = %v, v = %v \n", k, v)
	}

}
