package main

import (
	"fmt"
	"reflect"
)

type Mine struct {
	Name string `json:"name" comment:"the mine name"`
	Pass string `json:"pass" comment:"the mine pass"`
}

type inner struct {
	a int
	b int
}

type outer struct {
	c int
	d int
	int
	in inner
	a  int
}

type structA struct {
	a int
}

type structB struct {
	a, b int
}

type structC struct {
	sa structA
	sb structB
}

func main() {

	fmt.Println("-------------------------- 1、结构体标签使用 --------------------------")
	mine := Mine{"zhangsan", "123456"}
	minetype := reflect.TypeOf(mine)
	numfield := minetype.NumField()
	for i := 0; i < numfield; i++ {
		field := minetype.Field(i)
		fmt.Printf("field.tag = %v, filed.index = %v \n", field.Tag, field.Index)
	}

	fmt.Println("-------------------------- 2、内嵌结构体 --------------------------")
	in := inner{1, 3}
	ou := outer{2, 4, 6, in, 10}
	fmt.Printf("ou.a = %v, ou.in.a = %v \n", ou.a, ou.in.a)

	fmt.Println("-------------------------- 3、内嵌结构体没有命名冲突问题 ------------")
	sa := structA{1}
	sb := structB{10, 10}
	sc := structC{sa, sb}
	fmt.Printf("sc.sa.a = %v, sc.sb.a = %v \n", sc.sa.a, sc.sb.a)

}
