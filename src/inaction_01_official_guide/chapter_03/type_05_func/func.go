package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("******************************************* 1、func 作为参数案例 *****************************************")

	powValue := compute(math.Pow, 10, 3)
	fmt.Printf("powValue = %v \n", powValue)

	addValue := compute(add, 10, 8)
	fmt.Printf("addValue = %v \n", addValue)

	fmt.Println("******************************************* 2、func 作为返回值案例 ***************************************")

	fun := mineAdd()
	add := fun(10, 8)
	fmt.Printf("add value = %v \n", add)

}

func compute(fn func(float64, float64) float64, x float64, y float64) float64 {
	return fn(x, y)
}

func add(x float64, y float64) float64 {
	return x + y
}

func mineAdd() func(x float64, y float64) float64 {
	return func(x float64, y float64) float64 {
		return x + y
	}
}
