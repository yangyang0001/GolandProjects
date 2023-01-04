package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {

	fmt.Println("-------------------------- 1、map range 案例  --------------------------")
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("map item: apital of", key, "is", capitals[key])
	}

	fmt.Println("-------------------------- 2、map days  案例  --------------------------")
	days := map[string]time.Weekday{
		"Monday":    time.Monday,
		"Tuesday":   time.Tuesday,
		"Wednesday": time.Wednesday,
		"Thursday":  time.Thursday,
		"Friday":    time.Friday,
		"Saturday":  time.Saturday,
		"Sunday":    time.Sunday,
	}

	val, exist := days["Tuesday"]
	if exist {
		fmt.Printf("exist = true, key = %v \n", val)
	} else {
		fmt.Println("Tuesday is not exist!")
	}

	val, exist = days["Hollyday"]
	if exist {
		fmt.Printf("exist = true, key = %v \n", val)
	} else {
		fmt.Println("Hollyday is not exist!")
	}

	fmt.Println("-------------------------- 3、map 类型的切片  --------------------------")
	maps := make([]map[string]string, 3)
	maps[0] = map[string]string{"zer": "0"}
	maps[1] = map[string]string{"one": "1"}
	maps[2] = map[string]string{"one": "2"}

	for index, m := range maps {
		for key, val := range m {
			fmt.Printf("index = %v, maps[%v] = %v \n", index, key, val)
		}
	}

	fmt.Println("-------------------------- 4、map 排序的实现  --------------------------")
	sortMap := map[int]string{
		0: "zhangsan",
		1: "lisi",
		2: "wangwu",
		3: "zhaoliu",
	}
	for key, val := range sortMap {
		fmt.Printf("before sort invoke, sortMap[%v] = %v \n", key, val)
	}

	var keys = make([]int, len(sortMap))
	var i = 0
	for key, _ := range sortMap {
		keys[i] = key
		i++
	}

	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("after  sort invoke, sortMap[%v] = %v \n", key, sortMap[key])
	}

	fmt.Println("-------------------------- 5、map key val 对调  --------------------------")
	sourceMap := map[int]string{
		0: "zhangsan",
		1: "lisi",
		2: "wangwu",
		3: "zhaoliu",
	}

	reverseMap := make(map[string]int, len(sourceMap))

	for key, val := range sourceMap {
		reverseMap[val] = key
	}

	for key, val := range reverseMap {
		fmt.Printf("reverseMap[%v] = %v \n", key, val)
	}

}
