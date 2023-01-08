package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type struct1 struct {
	i1 int     `json:"i_1"`
	f1 float32 `json:"f_1"`
	st string  `json:"st"`
}

type Person struct {
	FirsName string `json:"firs_name"`
	LastName string `json:"last_name"`
}

func main() {

	fmt.Println("-------------------------- 1、声明初始化结构体  --------------------------")
	ms := struct1{10, 10.5, "Chris"}
	fmt.Printf("ms = %v, json = %v \n", ms, jsonParse(ms))

	fmt.Println("-------------------------- 2、结构体指针方法   --------------------------")
	person := Person{"zhang", "san"}
	updatePerson(&person)
	fmt.Printf("person json = %v \n", jsonParse(person))
}

func updatePerson(p *Person) {
	p.FirsName = strings.ToUpper(p.FirsName)
	p.LastName = strings.ToUpper(p.LastName)
}

func jsonParse(obj any) string {

	marshal, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	return string(marshal)
}
