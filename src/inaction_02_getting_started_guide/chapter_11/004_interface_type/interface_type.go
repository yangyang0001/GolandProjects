package main

import (
	"fmt"
)

type NI interface {
	GetNine() string
	SetNine(nine string)
}

type Nine struct {
	nine string
}

type None struct {
	nine string
}

func main() {

	fmt.Println("------------------------------------- 1、官网给定案例1 -------------------------------------")
	var ni NI
	ni = &Nine{"zhangsan"}
	ni = &None{"wangwu"}

	if nit, nif := ni.(*Nine); nif {
		fmt.Printf("ni type is = %T \n", nit)
	} else if not, nof := ni.(*None); nof {
		fmt.Printf("ni type is = %T \n", not)
	} else {
		fmt.Println("ni is other type!")
	}

	fmt.Println("------------------------------------- 2、type_switch -----------------------------------")
	var i int = 10
	var j float32 = 10
	var k float64 = 10
	do(i)
	do(j)
	do(k)

	do("zhangsan")
	do(Nine{"zhangsan"})
	do(None{"wangwu"})
	do(true)

	fmt.Println("------------------------------------- 3、判断 某个值 是否实现了 某个接口 -------------------")

}

func (nineP *Nine) GetNine() string {
	return nineP.nine
}

func (nineP *Nine) SetNine(nine string) {
	nineP.nine = nine
}

func (noneP *None) GetNine() string {
	return noneP.nine
}

func (noneP *None) SetNine(nine string) {
	noneP.nine = nine
}

func do(i interface{}) {

	switch i.(type) {
	case int:
		fmt.Printf("1 i type is int, i = %v, \n", i)

	case float32:
		fmt.Printf("2 i type is float32, i = %v, \n", i)

	case float64:
		fmt.Printf("3 i type is float64, i = %v, \n", i)

	case string:
		fmt.Printf("4 i type is string, i = %v, \n", i)

	case Nine:
		fmt.Printf("5 i type is Nine, i = %v, \n", i)

	case None:
		fmt.Printf("6 i type is None, i = %v, \n", i)

	default:
		fmt.Printf("7 i type is other, i = %v, \n", i)
	}

}
