package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func main() {

	fmt.Println("--------------------------------- 1、if-else 判断语句 ------------------------------")

	boolFlag := true
	if boolFlag {
		fmt.Println("the value is true!")
	} else {
		fmt.Println("the value is false!")
	}

	flag := false
	fmt.Printf("flag = %v, book json = %v \n", flag, getBook(flag))

	fmt.Println("--------------------------------- 2、测试多返回值函数的错误 --------------------------")
	intn, err := strconv.Atoi("hello")

	if err != nil {
		fmt.Errorf("strconv atoi error is :%v", err)
		log.Fatal(err)
	} else {
		fmt.Printf("intn = %v \n", intn)
	}

}

func getBook(flag bool) string {

	if flag {
		return jsonParse(Book{1, "Go语言入门", "zhangsan"})
	} else {
		return jsonParse(Book{2, "Go in action", "wangwu"})
	}

}

func jsonParse(obj any) string {
	marshal, _ := json.Marshal(obj)
	return string(marshal)
}
