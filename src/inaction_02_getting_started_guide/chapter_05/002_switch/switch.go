package main

import (
	"fmt"
	"time"
)

func main() {
	switch_method_0()

	switch_method_1()

	switch_method_2()

	switch_method_3(33.0)

	season()
}

func switch_method_0() {
	week := time.Now().Weekday()

	switch week {
	case time.Monday:
		fmt.Println("time now is monday!")

	case time.Tuesday:
		fmt.Println("time now is tuesday!")

	case time.Wednesday:
		fmt.Println("time now is wednesday!")

	case time.Thursday:
		fmt.Println("time now is thursday!")

	case time.Friday:
		fmt.Println("time now is friday!")

	case time.Saturday:
		fmt.Println("time now is saturday!")

	case time.Sunday:
		fmt.Println("time now is sunday!")

	default:
		fmt.Println("time is error!")
	}
}

func switch_method_1() {
	var number int = 100

	switch number {
	case 98, 99:
		fmt.Println("number equals 98 or 99!")

	case 100:
		fmt.Println("number equals 100!")

	default:
		fmt.Println("number not equals 98 to 100!")
	}
}

func switch_method_2() {
	var number int = 7

	switch {
	case number < 0:
		fmt.Println("number < 0")

	case number >= 0 && number <= 10:
		fmt.Println("number between 0 and 10")

	default:
		fmt.Println("Number is 10 or greater")
	}
}

func switch_method_3(obj any) {

	switch obj.(type) {
	case int:
		fmt.Printf("obj type is int, obj = %v \n", obj)
	case string:
		fmt.Printf("obj type is string, obj = %v \n", obj)
	default:
		fmt.Printf("obj type is other, obj = %v \n", obj)
	}

}

func season() {

	month := time.Now().Month()
	switch month {
	case time.January, time.February, time.March:
		fmt.Println("time now is in Spring!")

	case time.April, time.May, time.June:
		fmt.Println("time now is in Summer!")

	case time.July, time.August, time.September:
		fmt.Println("time now is in Autumn!")

	case time.October, time.November, time.December:
		fmt.Println("time now is in Winter!")

	default:
		fmt.Println("time is other!")
	}

}
