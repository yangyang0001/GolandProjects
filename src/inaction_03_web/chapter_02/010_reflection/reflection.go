package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Mine struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func main() {
	mine := Mine{"zhangsan", "111111"}
	fmt.Printf("mine json = %v \n", jsonParse(mine))

	typeof := reflect.TypeOf(mine)
	fmt.Printf("typeof  = %v \n", typeof)

	valueof := reflect.ValueOf(mine)
	fmt.Printf("valueof = %v \n", valueof)

	canset := reflect.ValueOf(mine).CanSet()
	fmt.Printf("mine element canset = %v \n", canset)

	mpcset := reflect.ValueOf(&mine).CanSet()
	fmt.Printf("&mine element mpcset = %v \n", mpcset)

	minevalue := reflect.ValueOf(&mine)
	if minevalue.Elem().CanSet() {
		minevalue.Elem().Field(0).SetString("lisi")
	}
	fmt.Printf("mine json = %v \n", jsonParse(mine))

}

func jsonParse(obj interface{}) string {
	marshal, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	return string(marshal)
}

