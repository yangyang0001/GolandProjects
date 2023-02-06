package main

import "fmt"

func main() {

	m0 := make(map[string]string)
	m0["a"] = "zhangsan"
	m0["b"] = "lisi"

	m1 := make(map[string][]string)
	m1["c"] = []string {"zhangsan", "lisi"}
	m1["d"] = []string {"齐达内", "克罗斯"}

	map0_each(m0)
	map1_each(m1)


}

func map0_each(m map[string]string)  {
	for key, value := range m {
		fmt.Printf("key = %v, value = %v \n", key, value)
	}
}

func map1_each(m map[string][]string)  {
	for key, value := range m {
		fmt.Printf("key = %v, value = %v \n", key, value)
	}
}

