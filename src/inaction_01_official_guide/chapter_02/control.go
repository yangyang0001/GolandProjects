package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {

	fmt.Println("--------------------------- 1、for 流程控制语句 ----------------------------")
	fmt.Printf("for_method invoke, the sum is %v \n", for_method())
	fmt.Printf("for_continued invoke, the sum is %v \n", for_continued())

	fmt.Println("--------------------------- 2、if  流程控制语句 ----------------------------")
	fmt.Printf("sqrt method inovke, x = %v, result = %v \n", 9, sqrt(9))
	fmt.Printf("sqrt method inovke, x = %v, result = %v \n", -9, sqrt(-9))
	fmt.Printf("pow_if inovke, x = %v, y = %v, limit = %v, result = %v \n", 2, 3, 10, pow_if(2, 3, 10))
	fmt.Printf("pow_if inovke, x = %v, y = %v, limit = %v, result = %v \n", 2, 4, 10, pow_if(2, 4, 10))
	fmt.Printf("pow_if_else inovke, x = %v, y = %v, limit = %v, result = %v \n", 2, 4, 25, pow_if_else(2, 4, 25))
	fmt.Println("--------------------------- 3、switch流程控制语句 --------------------------")
	get_os()
	show_time()
	get_time(time.Now().Weekday())
	none_param()
	fmt.Println("--------------------------- 4、defer 控制语句 ------------------------------")
	defer_method()
	defer_stack()

}

/** ******************************************* for 案例 ******************************************* */
func for_method() int {
	var sum = 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	return sum
}

func for_continued() int {
	var sum = 1

	for sum < 1000 {
		sum += sum
	}

	return sum
}

/**
 * 死循环
 */
func for_forever() {
	for {
		fmt.Println("this is for ever method invoke ...")
	}
}

/** ******************************************* if 案例 ******************************************* */
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow_if(x, y, limit float64) float64 {
	if value := math.Pow(x, y); value < limit {
		return value
	}
	return limit
}

func pow_if_else(x, y, limit float64) float64 {
	if value := math.Pow(x, y); value < limit {
		fmt.Println("first method invoke ..")
		return value
	} else {
		fmt.Println("secon method invoke ..")
		return limit
	}
}

/** ******************************************* switch 案例 ******************************************* */
func get_os() {

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("os is linux")
	case "darwin":
		fmt.Println("os is darwin")
	default:
		fmt.Printf("os is %v \n", os)
	}
}

func show_time() {

	fmt.Printf("time.Now()           is : %v \n", time.Now())
	fmt.Printf("time.Now().Weekday() is : %v \n", time.Now().Weekday())
	fmt.Printf("time.Now().Year()    is : %v \n", time.Now().Year())
	fmt.Printf("time.Now().Month()   is : %v \n", time.Now().Month())
	fmt.Printf("time.Now().Day()     is : %v \n", time.Now().Day())
	fmt.Printf("time.Now().Hour()    is : %v \n", time.Now().Hour())
	fmt.Printf("time.Now().Minute()  is : %v \n", time.Now().Minute())
	fmt.Printf("time.Now().Second()  is : %v \n", time.Now().Second())

}

func get_time(today time.Weekday) {
	switch time.Wednesday {
	case today + 0:
		fmt.Println("today.")
	case today + 1:
		fmt.Println("tomorrow.")
	case today + 2:
		fmt.Println("in two days.")
	default:
		fmt.Println("too far away.")
	}
}

func none_param() {
	time := time.Now().Hour()
	switch {
	case time < 12:
		fmt.Println("good morning!")
	case time < 18:
		fmt.Println("good afternoon!")
	default:
		fmt.Println("good evening!")

	}
}

/** ******************************************* defer 案例 ******************************************* */
func defer_method() {

	fmt.Println("defer method invoke start ...")

	defer fmt.Println("this is defer language!")

	fmt.Println("defer method invoke end ...")

}

func defer_stack() {

	fmt.Println("defer_stack method invoke start ...")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("i = %v \n", i)
	}

	fmt.Println("defer_stack method invoke end ...")

}
