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

	var a int = 10
	var b float32 = 10.0
	var c float64 = 3.1415926
	var d string = "zhangsan"
	var e Mine = Mine{"xiaoming", "root123456"}

	interface_type(a)
	interface_type(b)
	interface_type(c)
	interface_type(d)
	interface_type(e)

}

func interface_type(i interface{})  {

	switch i.(type) {
	case string:
		fmt.Printf("interface type is string, value = %v \n", i)

	case int:
		fmt.Printf("interface type is int, value = %v \n", i)

	case float32:
		fmt.Printf("interface type is float32, value = %v \n", i)

	case float64:
		fmt.Printf("interface type is float64, value = %v \n", i)

	case Mine:
		fmt.Printf("interface type is Mine, json = %v \n", jsonParse(i))
	}


}

func jsonParse(i interface{}) string {
	marshal, err := json.Marshal(i)

	if err != nil {
		log.Fatal(err)
	}

	return string(marshal)
}