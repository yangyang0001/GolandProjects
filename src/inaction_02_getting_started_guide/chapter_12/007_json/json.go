package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Mine struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func main() {

	fmt.Println("----------------------------------- 1、json 序列化 ----------------------------------")
	mine := Mine{"zhangsan", "123456"}
	fmt.Printf("temp = %v \n", jsonParse(mine))

	fmt.Println("----------------------------------- 2、json 反序列化 --------------------------------")
	var m Mine
	temp := "{\"name\":\"lisi\",\"pass\":\"root123\"}"
	json.Unmarshal([]byte(temp), &m)
	fmt.Printf("m = %v \n", m)



}

func jsonParse(obj interface{}) string {

	marshal, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	return string(marshal)

}
