package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Simpler interface {
	Get() string
	Set(param string)
}

type Simple struct {
	Name string
}

func main() {

	var simpler Simpler = &Simple{"zhangsan"}
	fmt.Printf("simple json = %v \n", jsonParse(simpler))
	simpler.Set("wangwu")
	fmt.Printf("simple json = %v \n", jsonParse(simpler))

}

func (simple *Simple) Get() string {
	return simple.Name
}

func (simple *Simple) Set(name string) {
	simple.Name = name
}

func jsonParse(obj any) string {
	marshal, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return string(marshal)
}
