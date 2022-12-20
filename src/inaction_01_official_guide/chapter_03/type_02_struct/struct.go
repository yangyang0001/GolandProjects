package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Point struct {
	X int
	Y int
}

func main() {

	/** ******************************************* 结构体 json ******************************************* */
	zhangsan := Student{1, "zhangsan", "zhangsan"}
	zhangsanBytes, err := json.Marshal(zhangsan)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("zhangsan json is : %v \n", string(zhangsanBytes))
	}

	students := []Student{
		{1, "zhangsan", "123456"},
		{2, "lisi", "234567"},
		{3, "wangwu", "345678"},
		{4, "zhaoliu", "456789"},
	}

	studentsBytes, err := json.Marshal(students)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("students json is : %v \n", string(studentsBytes))
	}

	/** ******************************************* 结构体赋值 ******************************************* */
	point0 := Point{3, 3}
	point0.Y = 4
	fmt.Printf("point0.X = %v, point0.Y = %v \n", point0.X, point0.Y)

	p0 := &point0
	p0.Y = 10
	fmt.Printf("point0.Y = %v \n", point0.Y)

	/** ******************************************* 结构体的 声明方式 ******************************************* */
	point1 := Point{1, 3}
	point2 := Point{X: 1}
	point3 := Point{}
	point4 := &Point{11, 22}

	point1Bytes, err := json.Marshal(point1)
	fmt.Printf("point1 json is %v \n", string(point1Bytes))

	point2Bytes, err := json.Marshal(point2)
	fmt.Printf("point2 json is %v \n", string(point2Bytes))

	point3Bytes, err := json.Marshal(point3)
	fmt.Printf("point3 json is %v \n", string(point3Bytes))

	fmt.Printf("point4.X = %v, point4.Y = %v \n", point4.X, point4.Y)

	point4X := &point4.X
	*point4X = 100
	point4.Y = 1000

	point4Bytes, err := json.Marshal(point4)
	fmt.Printf("point4 json is :%v \n", string(point4Bytes))

}
