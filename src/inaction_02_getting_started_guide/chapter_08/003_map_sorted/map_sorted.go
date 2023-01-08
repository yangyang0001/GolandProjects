package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("-------------------------- 1、map 按照 key 排序案例  --------------------------")
	var maps = map[int]string{
		3: "wangwu",
		2: "lisi",
		1: "zhangsan",
		4: "zhaoliu",
	}

	for key, val := range maps {
		fmt.Printf("before sorted, maps[%v] = %v \n", key, val)
	}

	var keys []int
	for key, _ := range maps {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("after  sorted, maps[%v] = %v \n", key, maps[key])
	}

	fmt.Println("-------------------------- 2、map 按照 val 排序案例  --------------------------")
	var vals []string
	for _, val := range maps {
		vals = append(vals, val)
	}

	sort.Strings(vals)
	fmt.Println(vals)

}
