package main

import "fmt"

type IDuck interface {
	Quack()
	Walk()
	ImplementsIDuck()
}

type Bird struct {
}

func main() {

	fmt.Println("--------------------------------------- 1、接口与动态类型 ---------------------------------------")
	var iduck IDuck = &Bird{}
	DuckDance(iduck)

	fmt.Println("--------------------------------------- 2、函数的重载案例 -------------------------------------")
	sum := Add(1, 3, 5, 10)
	fmt.Printf("add method invoke, sum = %v \n", sum)
}

func (*Bird) Quack() {
	fmt.Println("I am quaking!")
}
func (*Bird) Walk() {
	fmt.Println("I am walking!")
}

func (*Bird) ImplementsIDuck() {

}

func DuckDance(iduck IDuck) {
	iduck.Quack()
	iduck.Walk()
}

func Add(params ...int) int {
	var sum int = 0
	for _, param := range params {
		sum += param
	}
	return sum
}
