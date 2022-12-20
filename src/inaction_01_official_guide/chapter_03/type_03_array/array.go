package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func main() {

	fmt.Println("******************************************* 1、使用数组案例 *******************************************")
	var a [2]string
	a[0] = "hello"
	a[1] = "world"

	b := []string{"welcome", "to", "the", "array"}
	fmt.Printf("array a is : %v \n", a)
	fmt.Printf("array b is : %v \n", b)

	fmt.Println("******************************************* 2、使用切片案例 *******************************************")
	ints := []int{1, 2, 3, 4, 5, 6}
	subs := ints[1:3]
	fmt.Printf("subs is : %v \n", subs)

	// 切片 和 数组共享一段数据, 这块修改谁必然会影响 另外一个!
	subs[0] = 100
	fmt.Printf("ints is : %v \n", ints)
	fmt.Printf("subs is : %v \n", subs)

	fmt.Println("******************************************* 3、结构体数组和切片案例 *************************************")
	students := []Student{
		{1, "maosan", "123456"},
		{2, "gousi", "234567"},
		{3, "niuqi", "345678"},
	}
	fmt.Printf("students json is : %v\n", jsonParse(students))

	substudents := students[0:2]
	fmt.Printf("substudents json is : %v \n", jsonParse(substudents))

	substudents[0].Name = "root"
	substudents[0].Pass = "root"
	fmt.Printf("students json is : %v\n", jsonParse(students))
	fmt.Printf("substudents json is : %v \n", jsonParse(substudents))

	fmt.Println("******************************************* 4、切片的 容量cap 和 长度 len *******************************")
	fmt.Printf("substudents  cap is : %v \n", cap(substudents))
	fmt.Printf("substudents  len is : %v \n", len(substudents))

	fmt.Println("******************************************* 5、空切片 *************************************************")
	var s []int
	fmt.Printf("s is : %v, cap = %v, len = %v \n", s, cap(s), len(s))

	fmt.Println("******************************************* 6、使用 make 创建 切片 *************************************")
	m := make([]int, 5)
	fmt.Printf("m len is : %v \n", len(m))

	fmt.Println("******************************************* 7、创建并使用 二维数组 **************************************")
	board := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	for _, row := range board {
		for _, value := range row {
			fmt.Printf(value + " ")
		}
		fmt.Println()
	}

	fmt.Println("******************************************* 8、切片中 append 值 ***************************************")
	var arr []int
	arr = append(arr, 1, 3, 5)
	fmt.Printf("arr is %v \n", arr)

	fmt.Println("******************************************* 9、range 遍历元素 *****************************************")
	for index, value := range arr {
		fmt.Printf("index = %v, value = %v \n", index, value)
	}

}

func jsonParse(obj any) string {

	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
