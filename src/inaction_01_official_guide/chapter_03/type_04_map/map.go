package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Score struct {
	Name string  `json:"name"`
	High float64 `json:"high"`
}

func main() {
	fmt.Println("******************************************* 1、map 遍历 key, value *****************************************")
	score := map[string]Score{
		"zhangsan": {"zhangsan", 181.0},
		"lisi":     {"lisi", 175.5},
		"zhaoliu":  {"zhaoliu", 168.5},
	}

	for key, value := range score {
		fmt.Printf("key is %v, value is %v \n", key, jsonParse(value))
	}

	fmt.Println("******************************************* 2、修改 map 中的 值 ********************************************")
	map0 := map[string]int{
		"zhangsan": 90,
		"lisi":     80,
		"wangwu":   100,
	}
	map0["zhangsan"] = 99
	fmt.Println(map0["zhaoliu"])

	fmt.Println("******************************************* 3、make 构造 map 类型 *****************************************")
	m := make(map[string]string)
	fmt.Printf("m is %v, m.len = %v \n", m, len(m))

}

func jsonParse(obj any) string {
	objBytes, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	return string(objBytes)
}
