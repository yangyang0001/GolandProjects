package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type Mine interface {
	SayHi()
	Sing(song string)
}

type Human struct {
	Name string
	Age int
	Height float32
	Weight float32
}

type Student struct {
	Human
	school string
}

type Singer struct {
	Human
	company string
}

func main() {

	xiaoming := Student{Human{"xiaoming", 10, 180.0, 175.0}, "北京智障一中"}
	xiaohong := Student{Human{"xiaohong", 10, 172.0, 102.0}, "大陆傻叉二中"}

	liming := Singer{Human{"liming", 55, 181, 150}, "宝丽金"}
	xiaoge := Singer{Human{"xiaoge", 65, 175, 165}, "滚石"}

	var mines []Mine

	mines = append(mines, &xiaoming)
	mines = append(mines, &xiaohong)
	mines = append(mines, &liming)
	mines = append(mines, &xiaoge)

	for index, mine := range mines {
		fmt.Printf("index = %v, json = %v \n", index, jsonParse(mine))
		mine.SayHi()
		mine.Sing("song_" + strconv.Itoa(index))
		fmt.Println("------------------------------------------------")
	}

}

func (h *Human) SayHi()  {
	fmt.Printf("Hi everyone, I am %v \n", h.Name)
}

func (h *Human) Sing(song string)  {
	fmt.Printf("I am %v, the song is %v \n", h.Name, song)
}

func jsonParse(i interface{}) string {
	marshal, err := json.Marshal(i)

	if err != nil {
		log.Fatal(err)
	}

	return string(marshal)
}