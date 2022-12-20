package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("non-generic sum: %v and %v \n", SumInt(ints), SumFloat(floats))

	fmt.Printf("use-generic sum: %v and %v \n", Sum(ints), Sum(floats))

	fmt.Printf("generic-num sum: %v and %v \n", SumNumber(ints), SumNumber(floats))
}

func SumInt(m map[string]int64) int64 {
	var sum int64
	for _, value := range m {
		sum += value
	}
	return sum
}

func SumFloat(m map[string]float64) float64 {
	var sum float64
	for _, value := range m {
		sum += value
	}
	return sum
}

/**
 * 不使用 类型声明 的 泛型函数
 */
func Sum[K comparable, V int64 | float64](m map[K]V) V {

	var sum V
	for _, value := range m {
		sum += value
	}
	return sum
}

/**
 * 使用 类型声明 的 泛型函数
 */
func SumNumber[K comparable, V Number](m map[K]V) V {

	var sum V
	for _, value := range m {
		sum += value
	}
	return sum
}
