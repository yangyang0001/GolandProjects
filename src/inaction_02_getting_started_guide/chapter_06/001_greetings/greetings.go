package main

import (
	"encoding/json"
	"fmt"
)

type mineFunc func(param1 int, param2 int) int

type M interface {
	setId(id int)
	setName(name string)
	setPass(pass string)
}

type Mine struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func main() {

	fmt.Println("-------------------------- 1、基本函数调用案例  ----------------------------")
	fmt.Println("before greetings method invoke ...")
	greetings()
	fmt.Println("after  greetings method invoke ...")

	fmt.Println("-------------------------- 2、函数声明 案例 -------------------------------")

	fmt.Println("-------------------------- 3、数组参数 的方法 是引用传递 --------------------")
	ints := []int{1, 3, 5, 7, 9}
	array_update(ints)
	fmt.Printf("ints = %v \n", ints)

	fmt.Println("-------------------------- 4、map参数 的方法 是引用传递 --------------------")
	m := map[string]string{
		"A": "C罗",
		"B": "梅西",
		"C": "内马尔",
		"D": "姆巴佩",
	}
	map_update(m)
	fmt.Printf("m = %v \n", m)

	fmt.Println("-------------------------- 5、接口作为参数 案例  --------------------------")
	mine := Mine{1, "zhangsan", "zhangsan"}
	mine.setId(10)
	mine.setName("wangyan")
	mine.setPass("123456")
	fmt.Printf("mine json = %v \n", jsonParse(mine))

}

func greetings() {
	fmt.Println("greetings method invoke ...")
}

func array_update(arr []int) {
	arr[0] = 100
}

func map_update(m map[string]string) {
	delete(m, "A")
}

func (mine *Mine) setId(id int) {
	mine.Id = id
}

func (mine *Mine) setName(name string) {
	mine.Name = name
}

func (mine *Mine) setPass(pass string) {
	mine.Pass = pass
}

func jsonParse(obj any) string {
	marshal, _ := json.Marshal(obj)
	return string(marshal)
}
