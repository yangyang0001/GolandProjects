package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
)

type Mine struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func main() {

	fmt.Println("----------------------------- 1、并行定义变量的方式 -------------------------------")
	var a, b, c int = 10, 20, 30
	fmt.Printf("a = %v, b = %v, c = %v \n", a, b, c)

	fmt.Println("----------------------------- 2、使用 var () 定义变量 ----------------------------")
	var (
		d int     = 100
		e int     = 200
		f float64 = 10.88
	)
	fmt.Printf("d = %v, e = %v, f = %v \n", d, e, f)

	fmt.Println("----------------------------- 3、使用变量展示几个环境变量 --------------------------")
	var (
		GOROOT string = os.Getenv("GOROOT")
		GOPATH string = os.Getenv("GOPATH")
		PATH   string = os.Getenv("PATH")
		GOOS   string = runtime.GOOS
	)
	fmt.Printf("GOROOT = %v \n", GOROOT)
	fmt.Printf("GOPATH = %v \n", GOPATH)
	fmt.Printf("PATH   = %v \n", PATH)
	fmt.Printf("GOOS   = %v \n", GOOS)

	fmt.Println("----------------------------- 4、基本类型的地址获取案例 ----------------------------")
	var i int = 10
	var j string = "zhangsan"
	fmt.Printf("i = %v, i address is :%v \n", i, &i)
	fmt.Printf("j = %v, j address is :%v \n", j, &j)

	fmt.Println("----------------------------- 5、数组 和 结构体地址获取案例 ------------------------")
	/**
	验证了 数组参数的方法 是 值传递; 因此以后使用值传递的方法时, 要特别小心!
	arr address is :0xc000012030, arr = [100 200 300 400 500]
	update method invoke start, arr address is :0xc000012038, arr = [100 200 300 400 500]
	update method invoke start, arr address is :0xc000012038, arr = [8888 200 300 400 500]

	验证了 结构体参数的方法 就是 值传递! 因此以后使用值传递方法时, 要特别小心!
	mine address is :0xc000012030, json is :{"id":1,"name":"zhangsan","pass":"zhangsan"}
	setName method invoke start, mine address is :0xc000012070, json is :{"id":1,"name":"zhangsan","pass":"zhangsan"}
	setName method invoke end  , mine address is :0xc000012070, json is :{"id":1,"name":"Yang","pass":"zhangsan"}
	*/
	var arr []int = []int{100, 200, 300, 400, 500}
	var arraddress *[]int = &arr
	fmt.Printf("arr address is :%v, arr = %v \n", &arraddress, arr)
	update(arr)

	var mine Mine = Mine{1, "zhangsan", "zhangsan"}
	var add *Mine = &mine
	fmt.Printf("mine address is :%v, json is :%v \n", &add, jsonParse(mine))
	setName(mine, "Yang")

}

func update(arr []int) {
	var add *[]int = &arr
	fmt.Printf("update method invoke start, arr address is :%v, arr = %v \n", &add, arr)
	arr[0] = 8888
	fmt.Printf("update method invoke end  , arr address is :%v, arr = %v \n", &add, arr)

}

func setName(mine Mine, name string) {
	var add *Mine = &mine
	fmt.Printf("setName method invoke start, mine address is :%v, json is :%v \n", &add, jsonParse(mine))
	mine.Name = name
	fmt.Printf("setName method invoke end  , mine address is :%v, json is :%v \n", &add, jsonParse(mine))
}

func jsonParse(obj any) string {
	bytes, err := json.Marshal(obj)

	if err != nil {
		fmt.Errorf("jsonParse error is %v \n", err)
		log.Fatal(err)
	}

	return string(bytes)
}

// 优先级高于 main()
func init() {
	fmt.Println("init method invoke ...")
}
