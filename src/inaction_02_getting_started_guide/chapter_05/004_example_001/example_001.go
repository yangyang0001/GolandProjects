package main

import "fmt"

func main() {

	fmt.Println("-------------------------- 1、for 练习1 -------------------------")
	example_01()

	fmt.Println("-------------------------- 2、goto 练习 -------------------------")
	example_02()

	fmt.Println("-------------------------- 3、打印练习 ---------------------------")
	example_03()

	fmt.Println("-------------------------- 3、打印矩形 ---------------------------")
	example_04()

}

func example_01() {

	for i := 0; i < 15; i++ {
		fmt.Printf("i = %v \n", i)
	}

}

func example_02() {

	var i = 0

LOOP:
	if i < 15 {
		fmt.Printf("i = %v \n", i)
		i++
		goto LOOP
	} else {
		return
	}

}

func example_03() {
	for i := 1; i <= 25; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("G")
		}
		fmt.Println()
	}

	// fmt.Println(len("GGGGGGGGGGGGGGGGGGGGGGGGG"))
}

func example_04() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			if i == 0 || i == 9 || j == 0 || j == 19 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

	// fmt.Println(len("GGGGGGGGGGGGGGGGGGGGGGGGG"))
}
