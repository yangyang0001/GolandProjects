package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Mine struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func main() {

	fmt.Println("--------------------------------- 1、变量才能使用指针 ------------------------------")
	var i int = 10
	var m string = "hello"

	var pi *int = &i
	var pm *string = &m

	fmt.Printf("i = %v, i address = %v \n", i, pi)
	fmt.Printf("m = %v, m address = %v \n", m, pm)

	fmt.Println("--------------------------------- 2、常量是不能使用指针的 ----------------------------")
	//const j int = 100
	//const n string = "world"

	//var pj *int = &j
	//var pn *string = &n

	//fmt.Printf("j = %v, j address = %v \n", j, pj)
	//fmt.Printf("n = %v, n address = %v \n", n, pn)

	fmt.Println("--------------------------------- 3、数组 和 结构体指针 ----------------------------")
	arr := []int{1, 3, 5, 7}
	mine := Mine{1, "mine", "mine"}

	parr := &arr
	pmine := &mine

	fmt.Printf("arr  = %v, arr  address = %v \n", arr, &parr)
	fmt.Printf("mine = %v, mine address = %v \n", jsonParse(mine), &pmine)

}

func jsonParse(obj any) string {

	marshal, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	return string(marshal)
}
