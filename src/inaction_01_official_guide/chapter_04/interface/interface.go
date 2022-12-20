package main

import (
	"fmt"
	"math"
)

/**
 * 这里的接口, 类似于 java 中的 抽象类和接口的 综合体;
 * 表示: 一旦 某个类型或者结构体 具有了当前接口中 定义的所有方法, 那么就隐式的认为 某个类型或结构体实现了 当前接口!
**/
/*************************************** 1、第1个接口案例 ***************************************/
type MF interface {
	Abs() float64
}

// 声明struct
type Vertex struct {
	X, Y float64
}

// 声明自己的类型
type MineFloat float64

// struct 和 类型 添加方法
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (m MineFloat) Abs() float64 {
	if m < 0 {
		return float64(-m)
	}
	return float64(m)
}

/*************************************** 2、第2个接口案例 ***************************************/

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func (t *T) M() {
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(float64(f))
}

/*************************************** 4、判定接口值类型 ***************************************/
func do(i interface{}) {
	// TODO: x.(type) 只能在 switch 中使用!
	switch typeI := i.(type) {
	case int:
		fmt.Printf("case int, i = %v, type = %T \n", i, typeI)
	case string:
		fmt.Printf("case str, i = %v, type = %T \n", i, typeI)
	default:
		fmt.Printf("case oth, i = %v, type = %T \n", i, typeI)
	}
}

func main() {

	fmt.Println("*************************************** 1、第1个接口案例 ***************************************")
	var m MF

	f := MineFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	m = f
	fmt.Printf("m.Abs()  = %v \n", m.Abs())

	m = &v
	fmt.Printf("m.Abs()  = %v \n", m.Abs())

	fmt.Println("*************************************** 2、第2个接口案例 ***************************************")
	var i I
	i = &T{"Hello"}
	i.M()

	i = F(10.5)
	i.M()

	fmt.Println("*************************************** 3、空接口 案例 *****************************************")
	var j interface{}
	j = 10
	fmt.Printf("j = %v, type = %T \n", j, j)

	j = "hello"
	fmt.Printf("j = %v, type = %T \n", j, j)

	fmt.Println("*************************************** 4、判定接口值类型 ***************************************")
	var h interface{} = "hello"
	s := h.(string) // 断言 空接口 h 类型为 string, 并且将 string 值 赋值给 变量 s
	fmt.Printf("h type = %v \n", s)

	do(10)
	do(10.0)
	do("hello")

}
