package main

import (
	"deepblue.com/inaction_00/greetings"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("-------------------------- 1、Hello World 程序  ----------------------------")
	fmt.Println("Hello, World")

	fmt.Println("-------------------------- 2、调用greetings程序 ----------------------------")
	message := greetings.Hello("zhangsan")
	fmt.Println(message)

	fmt.Println("-------------------------- 3、返回 错误处理 案例 ----------------------------")
	message, err := greetings.HelloERR("zhangsan")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(message)
	}

	fmt.Println("-------------------------- 4、练习 数组 案例 ----------------------------")
	arr := []int{11, 22, 33}
	arreach(arr)

	fmt.Println("-------------------------- 5、练习 map 案例 ----------------------------")
	m := map[string]string{
		"a": "zhangsan",
		"b": "lisi",
		"c": "wangwu",
	}
	mapeach(m)

	fmt.Println("-------------------------- 6、多返回值 案例 ----------------------------")
	names := []string{"zhangsan", "lisi", "wangwu", "zhangliu"}
	resul, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	} else {
		mapeach(resul)
	}

	fmt.Println("-------------------------- 7、随机值 案例 ----------------------------")
	fmt.Println("[0 - 10) 的随机值 :", random(10))

	fmt.Println("-------------------------- 8、返回随机问候语 ----------------------------")
	result, err := greetings.RandomHello(names)
	if err != nil {
		log.Fatal(err)
	} else {
		mapeach(result)
	}
}

func arreach(arr []int) {
	for index, value := range arr {
		fmt.Println("index is :", index, ", value is :", value)
	}
}

func mapeach(m map[string]string) {
	for key, value := range m {
		fmt.Println("key is :", key, ", value is :", value)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func random(number int) int {
	return rand.Intn(number)
}
