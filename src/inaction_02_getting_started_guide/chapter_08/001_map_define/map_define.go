package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("-------------------------- 1、make map  ----------------------------")
	var mapLit map[string]int
	var mapAss map[string]int

	mapLit = map[string]int{
		"one": 1,
		"two": 2,
	}
	mapAss = mapLit

	mapCre := make(map[string]float32)
	mapCre["key1"] = 4.5
	mapCre["key2"] = 3.14159
	mapAss["two"] = 3 // value 集合为 1, 3

	for key, val := range mapLit {
		fmt.Printf("mapLit[%v] = %v \n", key, val)
	}

	for key, val := range mapAss {
		fmt.Printf("mapAss[%v] = %v \n", key, val)
	}

	for key, val := range mapCre {
		fmt.Printf("mapCre[%v] = %v \n", key, val)
	}

	fmt.Println("-------------------------- 2、map val为 func 的案例  -----------------")
	map_func := map[string]func() int{
		"one": func() int { return 1 },
		"two": func() int { return 2 },
	}

	fmt.Println(map_func)

	fmt.Println("-------------------------- 3、map 的容量 ----------------------------")
	var mine map[string]string = map[string]string{}
	for i := 0; i < 5; i++ {
		key := "name" + strconv.Itoa(i)
		val := "pass" + strconv.Itoa(i)
		mine[key] = val
		fmt.Printf("mine = %v, len(mine) = %v \n", mine, len(mine))
	}

	fmt.Println("-------------------------- 4、判断 map key 是否存在 -----------------")
	var player map[string]string = map[string]string{
		"C罗":  "皇马",
		"梅西":  "巴黎",
		"姆巴佩": "巴黎",
		"内马尔": "巴黎",
	}

	value, exist := player["梅西"]
	if exist {
		fmt.Printf("exist = true, value = %v \n", value)
	} else {
		fmt.Printf("exist = false, value = nil \n")
	}

	fmt.Println("-------------------------- 5、map 删除 key 案例  --------------------")
	delete(player, "梅西")
	delete(player, "C罗")
	delete(player, "哈维")
	delete(player, "魔笛")

	for key, val := range player {
		fmt.Printf("player[%v] = %v \n", key, val)
	}
}
