package main

import (
	"fmt"
	"time"
)

type TwoInts struct {
	a int
	b int
}

type IntVector []int

type employee struct {
	salary float64
}

type MineTime struct {
	time.Time
}

type B struct {
	thing int
}

func main() {

	fmt.Println("------------------------------ 1、methdos1 ------------------------------")
	two1 := TwoInts{12, 10}

	fmt.Printf("two1 sum = %v \n", two1.AddThem())
	fmt.Printf("two1 add them sum = %v \n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("two2 sum = %v \n", two2.AddThem())

	fmt.Println("------------------------------ 2、methdos2 ------------------------------")
	v := IntVector{1, 3, 5}
	fmt.Printf("v sum = %v \n", v.Sum())

	fmt.Println("------------------------------ 3、example1 ------------------------------")
	e := employee{100.0}
	e.giveRaise(0.1)
	fmt.Printf("e = %v \n", e)

	fmt.Println("------------------------------ 4、struct 和 struct.method 必须在一个包中 --")

	fmt.Println("------------------------------ 5、封装 time -----------------------------")
	minetime := MineTime{time.Now()}
	chars := minetime.first3Chars()
	fmt.Printf("minetime first3Chars invoke, chars = %v \n", chars)

	fmt.Println("------------------------------ 6、struct B 案例 -------------------------")
	var b1 B
	b1.change()
	fmt.Printf("b1 thing = %v \n", b1.write())

	b2 := new(B)
	b2.change()
	fmt.Printf("b2 thing = %v \n", b2.write())
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

func (v *IntVector) Sum() int {
	var sum = 0
	for _, val := range *v {
		sum += val
	}
	return sum
}

func (e *employee) giveRaise(percent float64) {
	e.salary *= (1 + percent)
}

func (t MineTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func (b *B) change() {
	b.thing = 1
}

func (b B) write() string {
	return fmt.Sprint(b.thing)
}
