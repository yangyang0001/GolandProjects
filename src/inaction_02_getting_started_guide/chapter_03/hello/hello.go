package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func main() {

	student0 := Student{"zhangsan", "123456"}

	fmt.Printf("Hello %v \n", student0)
	fmt.Printf("Hello %+v \n", student0)
	fmt.Printf("Hello %#v \n", student0)
	fmt.Printf("Hello %T \n", student0)
	fmt.Printf("Hello %v \n", jsonParse(student0))

}

func jsonParse(obj any) string {

	bytes, err := json.Marshal(obj)

	if err != nil {
		log.Fatal("jsonParse error is :", err)
	}

	return string(bytes)
}
